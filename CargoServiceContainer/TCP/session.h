#ifndef SESSION_H
#define SESSION_H

#include <QThread>
#include <QTcpSocket>
#include <QDebug>
#include <QMap>
#include <QList>
#include <QVector>
#include "../gen/rpc.pb.h"

class Session : public QThread
{
    Q_OBJECT

public:
    QString id;
    explicit Session(QTcpSocket* socket, QObject *parent = 0);
    ~Session();
    void run();

signals:
    void error(QTcpSocket::SocketError socketerror);
    void end(QString);
    void messageReceived(const QByteArray& data, QString sessionId);

private slots:
    void readyRead();
    void disconnected();

public slots:
    /**
     * @brief sendMessage Utility function to send a message to other end of the socket
     * @param msg
     */
    void sendMessage(const QByteArray& data, QString sessionId);

private:
    QTcpSocket *socket;
};

#endif // SESSION_H
