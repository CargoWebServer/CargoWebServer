#ifndef SESSION_H
#define SESSION_H

#include <QThread>
#include <QWebSocket>
#include <QDebug>
#include <QMap>
#include <QList>
#include <QVector>
#include "../gen/rpc.pb.h"

class Listener;

class Session : public QThread
{
    static int MAX_MESSAGE_SIZE;
    QWebSocket *socket;
    QMap<QString, QList<com::mycelius::message::Message*> > pending;
    QMap<QString, QVector<QByteArray> > pendingMsgChunk;

     Q_OBJECT
public:
    QString id;
    explicit Session(QWebSocket* socket, QObject *parent = 0);
    ~Session();
    void run();

signals:
    void end(QString);
        void onEvent(QString, int, QMap<QString, QVariant>);

private Q_SLOTS:
    void processBinaryMessage(QByteArray);
    void disconnected();

public Q_SLOTS:

    /**
     * @brief completeProcessMessageData Send back the answer to the client when the
     * action thread has finish processing the data.
     */
    void completeProcessMessageData(com::mycelius::message::Message*);

private:

    /**
     * @brief processPendingMessage When a message is larger than the MAX_MESSAGE_SIZE
     * the message is split into multiple smaller messages. Each message chung is process
     * by this function.
     *
     * @param messageId The original message id, each pending message have the same id...
     */
    void processPendingMessage(QString messageId);

    /**
     * @brief processIncommingMessage
     * @param msg
     */
    void processIncommingMessage(com::mycelius::message::Message& msg);

    /**
     * @brief sendMessage Utility function to send a message to other end of the socket
     * @param msg
     */
    void sendMessage(com::mycelius::message::Message *msg);
};

#endif // SESSION_H
