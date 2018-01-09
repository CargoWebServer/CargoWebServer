#ifndef MESSAGEPROCESSOR_H
#define MESSAGEPROCESSOR_H

#include <QObject>
#include <QList>
#include <QMap>
#include "gen/rpc.pb.h"
#include "event.hpp"
#include "action.h"

/**
 * @brief serializeToByteArray Serialyse the message to an array of bytes...
 * @param msg The proto message.
 * @return
 */
QByteArray serializeToByteArray(const google::protobuf::Message& msg);

class MessageProcessor : public QObject
{
    static int MAX_MESSAGE_SIZE;

    // The map of pending message.
    QMap<QString, QList<com::mycelius::message::Message> > pending;

    // Use for incomming message
    QMap<QString, QVector<QByteArray> > pendingMsgChunk;

    // Keep the association between message and session.
    QMap<QString, QString> messageSession;

    /**
     * @brief processPendingMessage When a message is larger than the MAX_MESSAGE_SIZE
     * the message is split into multiple smaller messages. Each message chung is process
     * by this function.
     *
     * @param messageId The original message id, each pending message have the same id...
     */
    void processPendingMessage(QString messageId);

    /**
     * @brief sendMessage Trigger send message event.
     * @param msg
     */
    void sendMessage(const com::mycelius::message::Message& msg);

    Q_OBJECT
public:
    explicit MessageProcessor(QObject *parent = nullptr);

signals:
    void onEvent(const Event&);
    void sendResponse(const QByteArray& data, QString);

public slots:
    /**
     * @brief processIncommingMessage Process incomming message.
     * @param msg
     */
    void processIncommingMessage(const QByteArray& data, QString sessionId);

    /**
     * @brief completeProcessMessageData Send back the answer to the client when the
     * action thread has finish processing the data.
     */
    void completeProcessMessageData(com::mycelius::message::Message*);
};

#endif // MESSAGEPROCESSOR_H
