#include <QJsonDocument>
#include <QJsonObject>
#include <QJsonArray>
#include "gen/rpc.pb.h"
#include "listener.hpp"

/**
 * @brief Session::MAX_MESSAGE_SIZE The size must be the same on both side of the socket.
 */
int Session::MAX_MESSAGE_SIZE = 17740;

void Session::processIncommingMessage(com::mycelius::message::Message& msg){
    // Now i will determine if the message is a request, a response or an event...
    if(msg.type() == com::mycelius::message::Message_MessageType_ERROR){
        // The message is an error

    }else if(msg.type() == com::mycelius::message::Message_MessageType_REQUEST){
        // Now I will call process message from the store.
        QString methodName = QString::fromStdString(msg.rqst().method());
        Action* action = new Action(QString::fromStdString(msg.rqst().id()), methodName, this->id);

        // Now I will append the parameters...
        const ::google::protobuf::RepeatedPtrField< ::com::mycelius::message::Data >&params = msg.rqst().params();

        for(::google::protobuf::RepeatedPtrField< ::com::mycelius::message::Data >::const_iterator it = params.cbegin();
            it != params.cend(); it++){
            ::com::mycelius::message::Data param = *it;
            QVariant var;
            if(param.type() == ::com::mycelius::message::Data_DataType_DOUBLE){
                var = QVariant(param.databytes().c_str()).toFloat();
                action->appendParam(QString::fromStdString(param.name()), var, "double");
            }else if(param.type() == ::com::mycelius::message::Data_DataType_INTEGER){
                var = QVariant(param.databytes().c_str()).toInt();
                action->appendParam(QString::fromStdString(param.name()), var, "int");
            }else if(param.type() == ::com::mycelius::message::Data_DataType_BOOLEAN){
                var = QVariant(param.databytes().c_str()).toBool();
                action->appendParam(QString::fromStdString(param.name()), var, "bool");
            }else if(param.type() == ::com::mycelius::message::Data_DataType_BYTES){
                var = QVariant(QByteArray(param.databytes().c_str(), param.databytes().length()));
                action->appendParam(QString::fromStdString(param.name()), var, "QByteArray");
            }else if(param.type() == ::com::mycelius::message::Data_DataType_JSON_STR){
                // JSON object found here. It can be array or a map...
                QString jsonStr = QVariant(QByteArray(param.databytes().c_str(), param.databytes().length())).toString();
                QJsonDocument jsonDoc = QJsonDocument::fromJson(jsonStr.toUtf8());
                // From the jsonDoc...
                if(jsonDoc.isObject()){
                     QJsonObject jsonObject = jsonDoc.object();
                     var = jsonObject;
                     action->appendParam(QString::fromStdString(param.name()), var, "QJsonObject");

                }else if(jsonDoc.isArray()){
                    QJsonArray jsonArray = jsonDoc.array();
                    var = jsonArray;
                    action->appendParam(QString::fromStdString(param.name()), var, "QJsonArray");
                }else if(jsonDoc.isEmpty() || jsonDoc.isNull()){
                    //var = NULL;
                }

            }else if(param.type() == ::com::mycelius::message::Data_DataType_STRING){
                var = QVariant(param.databytes().c_str());
                action->appendParam(QString::fromStdString(param.name()), var, "QString");
            }
            //qDebug() << "Type name:" << var.typeName() << " Param type: " << param.type() << " name " <<  QString::fromStdString(param.name()) << " value:" << var;
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
        //qDebug() << " process pending message with id " << messageId;
        if(this->pending.contains(messageId)){
             //qDebug() << " number of pending message for " << messageId << " is " << this->pending.find(messageId)->size();
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
                // All chuck are received here...
                QByteArray originalMessageData;
                for(QVector<QByteArray>::iterator it = array.begin(); it != array.end(); ++it){
                    originalMessageData = originalMessageData + *it;
                }
                this->pendingMsgChunk.remove(messageId);

                // Here I will recreate the original message from the assembled data array...
                com::mycelius::message::Message originalMessage;

                bool isCorrect = originalMessage.ParseFromArray(originalMessageData.constData(), originalMessageData.size());
                if(isCorrect){
                    this->processIncommingMessage(originalMessage);
                }else{
                    qDebug() << "Fail to recreate the original message: ";
                    qDebug() << originalMessageData.constData();
                }
            }
        }else{
            // Here I will store the pending message...
            QVector<QByteArray> container;
            container.resize(total);
            container[0] = QByteArray(msg.data().c_str(), msg.data().size());
            this->pendingMsgChunk.insert(messageId, container);
        }

        // Here I will send back the response...
        com::mycelius::message::Message *responseMsg = new com::mycelius::message::Message();
        responseMsg->set_type(com::mycelius::message::Message_MessageType_RESPONSE);
        responseMsg->set_id(messageId.toStdString());
        responseMsg->set_index(-1);
        responseMsg->set_total(1);

        com::mycelius::message::Response *rsp = new com::mycelius::message::Response();
        rsp->set_id(messageId.toStdString());
        responseMsg->set_allocated_rsp(rsp);

        // Response was sent to the server.
        this->sendMessage(responseMsg);
        delete responseMsg; // Remove the response explicitely here.

    }else if(msg.type() == com::mycelius::message::Message_MessageType_EVENT){
        QString channelId = QString::fromStdString(msg.evt().name());
        int evtNumber = msg.evt().code();

        QMap<QString, QVariant> evtDataMap;
        const ::google::protobuf::RepeatedPtrField< ::com::mycelius::message::Data >&params =  msg.evt().evtdata();
        for(::google::protobuf::RepeatedPtrField< ::com::mycelius::message::Data >::const_iterator it = params.cbegin();
            it != params.cend(); it++){
            ::com::mycelius::message::Data param = *it;
            QVariant var;
            if(param.type() == ::com::mycelius::message::Data_DataType_DOUBLE){
                var = QVariant(param.databytes().c_str()).toFloat();
            }else if(param.type() == ::com::mycelius::message::Data_DataType_INTEGER){
                var = QVariant(param.databytes().c_str()).toInt();
            }else if(param.type() == ::com::mycelius::message::Data_DataType_BOOLEAN){
                var = QVariant(param.databytes().c_str()).toBool();
            }else if(param.type() == ::com::mycelius::message::Data_DataType_BYTES){
                var = QVariant(QByteArray(param.databytes().c_str(), param.databytes().length()));
            }else if(param.type() == ::com::mycelius::message::Data_DataType_JSON_STR){
                // JSON object found here. It can be array or a map...
                QString jsonStr = QVariant(QByteArray(param.databytes().c_str(), param.databytes().length())).toString();
                QJsonDocument jsonDoc = QJsonDocument::fromJson(jsonStr.toUtf8());
                // From the jsonDoc...
                if(jsonDoc.isObject()){
                     QJsonObject jsonObject = jsonDoc.object();
                     var = jsonObject;
                }else if(jsonDoc.isArray()){
                    QJsonArray jsonArray = jsonDoc.array();
                    var = jsonArray;
                }else if(jsonDoc.isEmpty() || jsonDoc.isNull()){
                    //var = NULL;
                }

            }else if(param.type() == ::com::mycelius::message::Data_DataType_STRING){
                var = QVariant(param.databytes().c_str());
            }
            evtDataMap[QString::fromStdString(param.name())] = var;
        }

        // dispatch to the listeners.
        Event evt(channelId, evtNumber, evtDataMap);

        emit this->onEvent(evt);
    }
}

void Session::completeProcessMessageData(com::mycelius::message::Message * msg){
    // Now i can remove the action...
    QString messageId = QString::fromStdString(msg->rsp().id());

    if( msg->ByteSize() < Session::MAX_MESSAGE_SIZE){
        this->sendMessage(msg);
    }else{
        QByteArray messageData = serializeToByteArray(msg);

        int count = int(double(msg->ByteSize() / Session::MAX_MESSAGE_SIZE));
        if(msg->ByteSize() % Session::MAX_MESSAGE_SIZE > 0){
            count++;
        }

        // If not fit exactly inside the message.
        this->pending.insert(messageId, QList<com::mycelius::message::Message*>() );
        for(int i=0; i<count; i++){
            QByteArray bytesSlice;
            int startIndex = i * Session::MAX_MESSAGE_SIZE;
            if(startIndex + Session::MAX_MESSAGE_SIZE < messageData.size()){
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

    // Remove the incomming message to be process, it can be sent directly or chunked...
    delete msg;
    msg = NULL;
}

void Session::processPendingMessage(QString messageId){
    if(this->pending.find(messageId)->length() > 0){
        // Here I will get the first message...
        com::mycelius::message::Message* msg = this->pending.find(messageId)->at(0);
        this->pending.find(messageId)->pop_front();
        if(this->pending.find(messageId)->length() == 0){
             this->pending.remove(messageId);
        }

        // Send the next message pending message.
        this->sendMessage(msg);
        delete msg;
        msg = NULL;
    }
}

void Session::disconnected()
{
    qDebug() << "session closed!";
    if(socket != NULL){
        socket->deleteLater();
    }
    exit(0);
}
