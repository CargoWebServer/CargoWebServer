#include "session.h"
#include "action.h"
#include "serviceContainer.h"
#include "gen/rpc.pb.h"
#include <QCoreApplication>
#include <QThreadPool>

// Common ws/tcp code.
#include "../session.cpp"

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
    qDebug() << "session is deleted!";
}

void Session::run()
{
    // make this thread a loop,
    // thread will stay alive so that signal/slot to function properly
    // not dropped out in the middle when thread dies
    exec();
}

void Session::sendMessage(com::mycelius::message::Message *msg){
    // Send messsage back.
    this->socket->sendBinaryMessage(serializeToByteArray(msg));
}

void Session::processBinaryMessage(QByteArray data)
{
    // get the information
    com::mycelius::message::Message msg;
    msg.ParseFromArray(data, data.size());
    this->processIncommingMessage(msg);
}


/*if(this->socket != NULL){
    disconnect(this->socket, &QWebSocket::binaryMessageReceived, this, &Session::processBinaryMessage);
    disconnect(this->socket, &QWebSocket::disconnected, this, &Session::disconnected);
}*/
