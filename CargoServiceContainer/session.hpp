#ifndef SESSION_H
#define SESSION_H

#include <QThread>
#include <QWebSocket>
#include <QDebug>
#include <QMap>
#include <QList>
#include "gen/rpc.pb.h"

class Session : public QThread
{
    QWebSocket *socket;

     Q_OBJECT
public:
    QString id;
    explicit Session(QWebSocket* socket, QObject *parent = 0);
    ~Session();
    void run();

signals:
    void end(QString);
    void messageReceived(const QByteArray& data, QString sessionId);

private Q_SLOTS:
    void processBinaryMessage(QByteArray);
    void disconnected();

public Q_SLOTS:
    /**
     * @brief sendMessage Utility function to send a message to other end of the socket
     * @param msg
     */
    void sendMessage(const QByteArray& data, QString sessionId);

};

#endif // SESSION_H
