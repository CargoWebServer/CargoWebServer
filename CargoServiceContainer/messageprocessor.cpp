#include "messageprocessor.hpp"
#include <QJsonDocument>
#include <QJsonObject>
#include <QJsonArray>
#include <QThreadPool>
#include <QDebug>
#include "gen/rpc.pb.h"
#include "listener.hpp"
#include "action.hpp"

/**
 * @brief Session::MAX_MESSAGE_SIZE The size must be the same on both side of the socket.
 */
int MessageProcessor::MAX_MESSAGE_SIZE = 17740;

QByteArray serializeToByteArray(const google::protobuf::Message& msg){
    QByteArray ra;
    ra.resize(msg.ByteSize());
    msg.SerializeToArray(ra.data(),ra.size());
    return ra;
}

MessageProcessor::MessageProcessor(QObject *parent) : QObject(parent)
{

}

void MessageProcessor::processIncommingMessage(const QByteArray& data, QString sessionId){
    com::mycelius::message::Message msg;
    if(!msg.ParseFromArray(data, data.size())){
        qDebug() << "fail to parse protobuffer message from byte array!";
        qDebug() << msg.DebugString().c_str();
        return;
    }

    // Keep message and session associated.
    this->messageSession.insert(QString::fromStdString(msg.id()), sessionId);

    // Now i will determine if the message is a request, a response or an event...
    if(msg.type() == com::mycelius::message::Message_MessageType_ERROR){
        // The message is an error

    }else if(msg.type() == com::mycelius::message::Message_MessageType_REQUEST){
        // Now I will call process message from the store.
        QString methodName = QString::fromStdString(msg.rqst().method());
        Action* action = new Action(QString::fromStdString(msg.rqst().id()), methodName, sessionId);

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
                    originalMessageData.append(*it);
                }

                this->pendingMsgChunk.remove(messageId);
                // qDebug() << "Remove message: " << messageId << " pending count: " << this->pendingMsgChunk.count();

                // Here I will recreate the original message from the assembled data array...
                com::mycelius::message::Message originalMessage;
                if(originalMessage.ParseFromArray(originalMessageData.constData(), originalMessageData.size())){
                    this->processIncommingMessage(originalMessageData, sessionId);
                }else{
                    qDebug() << "Fail to recreate the original message: ";
                    qDebug() << originalMessageData.constData();
                }
            }
        }else{
            // qDebug() << "Append message " << messageId << index << ":" << total;
            // Here I will store the pending message...
            QVector<QByteArray> container;
            container.resize(total);
            container[index] = QByteArray(msg.data().c_str(), msg.data().size());
            this->pendingMsgChunk.insert(messageId, container);
        }

        // The last message does not need to send respond back.
        if( index != total - 1){
            // Here I will send back the response...
            com::mycelius::message::Message responseMsg;
            responseMsg.set_type(com::mycelius::message::Message_MessageType_RESPONSE);
            responseMsg.set_id(messageId.toStdString());
            responseMsg.set_index(-1);
            responseMsg.set_total(1);

            // allocated message is a pointer manage by the protobuffer.
            com::mycelius::message::Response *rsp = new com::mycelius::message::Response();
            rsp->set_id(messageId.toStdString());
            responseMsg.set_allocated_rsp(rsp);

            this->sendMessage(responseMsg);
        }

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

void MessageProcessor::completeProcessMessageData(com::mycelius::message::Message* msg){
    // Now i can remove the action...
    QString messageId = QString::fromStdString(msg->rsp().id());

    if(msg->ByteSize() < MessageProcessor::MAX_MESSAGE_SIZE){
        this->sendMessage(*msg);
    }else{
        QByteArray messageData = serializeToByteArray(*msg);

        int count = int(double(msg->ByteSize() / MessageProcessor::MAX_MESSAGE_SIZE));
        if(msg->ByteSize() % MessageProcessor::MAX_MESSAGE_SIZE > 0){
            count++;
        }

        // If not fit exactly inside the message.
        this->pending.insert(messageId, QList<com::mycelius::message::Message>() );
        for(int i=0; i<count; i++){
            QByteArray bytesSlice;
            int startIndex = i * MessageProcessor::MAX_MESSAGE_SIZE;
            if(startIndex + MessageProcessor::MAX_MESSAGE_SIZE < messageData.size()){
                bytesSlice = messageData.mid(startIndex, MessageProcessor::MAX_MESSAGE_SIZE );
            } else {
                bytesSlice = messageData.mid(startIndex);
            }

            // Now I will create a transfert message...
            com::mycelius::message::Message transferMsg;
            transferMsg.set_type(com::mycelius::message::Message_MessageType_TRANSFER);
            transferMsg.set_id(messageId.toStdString());
            transferMsg.set_index(i);
            transferMsg.set_total(count);

            // Set the data...
            transferMsg.set_data(bytesSlice.toStdString());

            // Append to the pending message.
            this->pending.find(messageId)->push_back(transferMsg);
        }

        // Start the message transfer...
        this->processPendingMessage(messageId);
    }

    // Done with the message
    delete msg;
    msg = NULL;
}

void MessageProcessor::processPendingMessage(QString messageId){
    if(this->pending.find(messageId)->length() > 0){
        // Here I will get the first message...
        com::mycelius::message::Message msg = this->pending.find(messageId)->at(0);
        this->pending.find(messageId)->pop_front();
        if(this->pending.find(messageId)->length() == 0){
            this->pending.remove(messageId);
        }
        // Send the next message pending message.
        this->sendMessage(msg);
    }
}

void MessageProcessor::sendMessage(const com::mycelius::message::Message& msg){
    QByteArray data =  serializeToByteArray(msg);
    QString sessionId = this->messageSession.value(QString::fromStdString(msg.id()));
    this->messageSession.remove(QString::fromStdString(msg.id()));
    emit this->sendResponse(data, sessionId);
}
