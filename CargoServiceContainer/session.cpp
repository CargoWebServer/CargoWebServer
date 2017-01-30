#include "session.h"
#include "action.h"
#include "serviceContainer.h"
#include "gen/rpc.pb.h"
#include <QCoreApplication>
#include <QThreadPool>

/**
 * @brief serializeToByteArray Serialyse the message to an array of bytes...
 * @param msg The proto message.
 * @return
 */
QByteArray serializeToByteArray(google::protobuf::Message *msg){
    QByteArray ra;
    ra.resize(msg->ByteSize());
    msg->SerializeToArray(ra.data(),ra.size());
    return ra;
}

/**
 * @brief writelen Write the len of the flowing message to be read in the buffer.
 * @param soc The socket where to write the message in
 * @param len The calculated length of the message.
 */
void writelen(QAbstractSocket *soc,uint32_t len){
    soc->write((char*)&len,4);
}

/**
 * @brief Session::MAX_MESSAGE_SIZE The size must be the same on both side of the socket.
 */

int Session::MAX_MESSAGE_SIZE = 17739;

Session::Session(QWebSocket* socket, QObject *parent) :
    QThread(parent)
{
    this->socket = socket;

    // connect socket and signal
    // note - Qt::DirectConnection is used because it's multithreaded
    //        This makes the slot to be invoked immediately, when the signal is emitted.
    connect(this->socket, &QWebSocket::binaryMessageReceived, this, &Session::processBinaryMessage, Qt::DirectConnection);
    connect(this->socket, &QWebSocket::disconnected, this, &Session::disconnected);

    // Move the socket to the main thread so it will be accessible
    // from inside the slot...
    this->socket->setParent(NULL);
    this->socket->moveToThread(QCoreApplication::instance()->thread());
}

Session::~Session(){
    disconnect(this->socket, &QWebSocket::binaryMessageReceived, this, &Session::processBinaryMessage);
    disconnect(this->socket, &QWebSocket::disconnected, this, &Session::disconnected);
    qDebug() << "session is now closed!";
}

void Session::run()
{
    // make this thread a loop,
    // thread will stay alive so that signal/slot to function properly
    // not dropped out in the middle when thread dies
    exec();
}

void Session::processBinaryMessage(QByteArray data)
{
    // get the information
    com::mycelius::message::Message msg;
    msg.ParseFromArray(data, data.size());
    this->processIncommingMessage(msg);
}

void Session::processIncommingMessage(com::mycelius::message::Message& msg){

    // Now i will determine if the message is a request, a response or an event...
    if(msg.type() == com::mycelius::message::Message_MessageType_ERROR){
        // The message is an error
    }else if(msg.type() == com::mycelius::message::Message_MessageType_REQUEST){
        // Now I will call process message from the store.
        QString methodName = QString::fromStdString(msg.rqst().method());
        Action* action = new Action(QString::fromStdString(msg.rqst().id()), methodName);

        // Now I will append the parameters...
        const ::google::protobuf::RepeatedPtrField< ::com::mycelius::message::Data >& params = msg.rqst().params();

        for(::google::protobuf::RepeatedPtrField< ::com::mycelius::message::Data >::const_iterator it = params.cbegin();
            it != params.cend(); it++){

            ::com::mycelius::message::Data param = *it;
            QVariant var;

            if(param.type() == ::com::mycelius::message::Data_DataType_DOUBLE){

            }else if(param.type() == ::com::mycelius::message::Data_DataType_INTEGER){

            }else if(param.type() == ::com::mycelius::message::Data_DataType_STRING){
                var = QVariant(param.databytes().c_str());
                qDebug() << "Type name:" << var.typeName() << " value:" << var;
                action->appendParam(QString::fromStdString(param.name()), var, "QString");

            }else if(param.type() == ::com::mycelius::message::Data_DataType_BYTES){
                var = QVariant(QByteArray(param.databytes().c_str(), param.databytes().length()));
                action->appendParam(QString::fromStdString(param.name()), var, "QByteArray");
            }else if(param.type() == ::com::mycelius::message::Data_DataType_JSON_STR){

            }
        }

        // Connect the slot whit the signal...
        connect(action, SIGNAL(done(com::mycelius::message::Message*)),
                this, SLOT(completeProcessMessageData(com::mycelius::message::Message*)),
                Qt::AutoConnection);


        // In that case the action will be execute...
        QThreadPool::globalInstance()->start(action);


    }else if(msg.type() == com::mycelius::message::Message_MessageType_RESPONSE){
        // The message is a response...
        QString messageId = QString::fromStdString(msg.rsp().id());
        if(this->pending.contains(messageId)){
            qDebug() << " process next pending message, message id is " << messageId;
            this->processPendingMessage(messageId);
        }

    }else if(msg.type() == com::mycelius::message::Message_MessageType_TRANSFER){

        int total = msg.total();
        int index = msg.index();
        QString messageId = QString::fromStdString(msg.id());

        if(this->pendingMsgChunk.contains(messageId) == true){
            // First I will insert the message inside the vector...
            QVector<QByteArray>& array = *this->pendingMsgChunk.find(messageId);
            array[index] = QByteArray(msg.data().c_str(), msg.data().size());

            if( index == total - 1){
                QByteArray originalMessageData;
                for(QVector<QByteArray>::iterator it = array.begin(); it != array.end(); ++it){
                    originalMessageData = originalMessageData + *it;
                }

                this->pendingMsgChunk.remove(messageId);

                // Here I will recreate the original message from the assembled data array...
                com::mycelius::message::Message originalMessage;
                originalMessage.ParseFromArray(originalMessageData.constData(), originalMessageData.size());
                this->processIncommingMessage(originalMessage);
            }
        }else{
            // Here I will store the pending message...
            QVector<QByteArray> container;
            container.resize(total);
            container[0] = QByteArray(msg.data().c_str(), msg.data().size());
            this->pendingMsgChunk.insert(messageId, container);
        }

        // Here I will send back the response...
        com::mycelius::message::Message* responseMsg = new com::mycelius::message::Message();
        responseMsg->set_type(com::mycelius::message::Message_MessageType_RESPONSE);
        responseMsg->set_id(messageId.toStdString());
        responseMsg->set_index(-1);
        responseMsg->set_total(1);

        com::mycelius::message::Response* rsp = new com::mycelius::message::Response();
        rsp->set_id(messageId.toStdString());
        responseMsg->set_allocated_rsp(rsp);
        this->sendMessage(responseMsg);

    }else if(msg.type() == com::mycelius::message::Message_MessageType_EVENT){

    }
}

void Session::sendMessage(com::mycelius::message::Message *msg){
    // Send messsage back.
    this->socket->sendBinaryMessage(serializeToByteArray(msg));
}

void Session::completeProcessMessageData(com::mycelius::message::Message * msg){

    if( msg->ByteSize() < Session::MAX_MESSAGE_SIZE){
        this->sendMessage(msg);
    }else{

        int count = msg->ByteSize() / Session::MAX_MESSAGE_SIZE;

        // Round up...
        if(msg->ByteSize() % Session::MAX_MESSAGE_SIZE > 0){
            count++;
        }

        QByteArray messageData = serializeToByteArray(msg);
        QString messageId = QString::fromStdString(msg->rsp().id());

        this->pending.insert(messageId, QList<com::mycelius::message::Message*>() );

        for(int i=0; i<=count; i++){
            QByteArray bytesSlice;
            int startIndex = i * Session::MAX_MESSAGE_SIZE;
            if(i<count){
                bytesSlice = messageData.mid(startIndex, Session::MAX_MESSAGE_SIZE );
            } else {
                bytesSlice = messageData.mid(startIndex);
            }

            // Now I will create a transfert message...
            com::mycelius::message::Message* transferMsg = new com::mycelius::message::Message();
            transferMsg->set_type(com::mycelius::message::Message_MessageType_TRANSFER);
            transferMsg->set_id(messageId.toStdString());
            transferMsg->set_index(i);
            transferMsg->set_total(count);

            // Set the data...
            transferMsg->set_data(bytesSlice.toStdString());

            // Append to the pending message.
            this->pending.find(messageId)->push_back(transferMsg);
        }

        // Start the message transfer...
        this->processPendingMessage(messageId);
    }
}

void Session::processPendingMessage(QString messageId){
    if(this->pending.find(messageId)->length() > 0){
        // Here I will get the first message...
        com::mycelius::message::Message* msg = this->pending.find(messageId)->at(0);

        // Serialyse the message
        this->sendMessage(msg);

        if(msg->type() == com::mycelius::message::Message_MessageType_RESPONSE){
            this->pending.remove(messageId);
        }else{
            this->pending.find(messageId)->pop_front();
        }

        delete msg;
        msg = NULL;
    }
}

void Session::disconnected()
{
    socket->deleteLater();
    this->exit(0);
}
