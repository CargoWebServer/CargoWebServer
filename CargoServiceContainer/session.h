#ifndef SESSION_H
#define SESSION_H

#include <QThread>
#include <QTcpSocket>
#include <QDebug>
#include <QMap>
#include <QVector>
#include "gen/rpc.pb.h"

class Session : public QThread
{
    static int MAX_MESSAGE_SIZE;

     Q_OBJECT
public:
    explicit Session(qintptr ID, QObject *parent = 0);
    void run();

signals:
    void error(QTcpSocket::SocketError socketerror);

public slots:
    void readyRead();
    void disconnected();

    /**
     * @brief completeProcessMessageData Send back the answer to the client when the
     * action thread has finish processing the data.
     */
    void completeProcessMessageData(com::mycelius::message::Message*);

private:
    QTcpSocket *socket;
    qintptr socketDescriptor;
    QMap<QString, QList<com::mycelius::message::Message*> > pending;
    QMap<QString, QVector<QByteArray> > pendingMsgChunk;

    /**
     * @brief processPendingMessage When a message is larger than the MAX_MESSAGE_SIZE
     * the message is split into multiple smaller messages. Each message chung is process
     * by this function.
     *
     * @param messageId The original message id, each pending message have the same id...
     */
    void processPendingMessage(QString messageId);


    void processIncommingMessage(com::mycelius::message::Message& msg);

    /**
     * @brief sendMessage Utility function to send a message to other end of the socket
     * @param msg
     */
    void sendMessage(com::mycelius::message::Message *msg);
};

#endif // SESSION_H
