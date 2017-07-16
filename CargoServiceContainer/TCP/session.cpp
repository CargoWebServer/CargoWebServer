#include "session.h"
#include "action.h"
#include "serviceContainer.h"
#include "gen/rpc.pb.h"
#include <QCoreApplication>
#include <QThreadPool>

// Common ws/tcp code.
#include "../session.cpp"

Session::Session(qintptr ID, QObject *parent) :
    QThread(parent)
{
    this->socketDescriptor = ID;
}

void Session::run()
{
    // thread starts here
    qDebug() << " Thread started";
    socket = new QTcpSocket();

    // set the ID
    if(!socket->setSocketDescriptor(this->socketDescriptor))
    {
        // something's wrong, we just emit a signal
        emit error(socket->error());
        return;
    }

    // connect socket and signal
    // note - Qt::DirectConnection is used because it's multithreaded
    //        This makes the slot to be invoked immediately, when the signal is emitted.
    connect(socket, SIGNAL(readyRead()), this, SLOT(readyRead()), Qt::DirectConnection);
    connect(socket, SIGNAL(disconnected()), this, SLOT(disconnected()));

    // We'll have multiple clients, we want to know which is which
    qDebug() << socketDescriptor << " Client connected";

    // Move the socket to the main thread so it will be accessible
    // from inside the slot...
    socket->setParent(NULL);
    socket->moveToThread(QCoreApplication::instance()->thread());

    // make this thread a loop,
    // thread will stay alive so that signal/slot to function properly
    // not dropped out in the middle when thread dies
    exec();
}

void Session::sendMessage(com::mycelius::message::Message *msg){
    qDebug() << QString::fromStdString(msg->mutable_id()->c_str());
    qDebug() << QString::fromStdString(msg->mutable_rsp()->mutable_id()->c_str());
    writelen(this->socket,msg->ByteSize());
    this->socket->write(serializeToByteArray(msg));
    this->socket->waitForBytesWritten();
}

void Session::readyRead()
{
    // get the information
    while(this->socket->bytesAvailable())
    {
        QByteArray buffer;

        int dataSize;
        this->socket->read((char*)&dataSize, sizeof(int));
        buffer = this->socket->read(dataSize);
        while(buffer.size() < dataSize ) // only part of the message has been received
        {
            this->socket->waitForReadyRead(); // alternatively, store the buffer and wait for the next readyRead()
            buffer.append(this->socket->read(dataSize - buffer.size())); // append the remaining bytes of the message
        }
        com::mycelius::message::Message msg;
        msg.ParseFromArray(buffer, buffer.size());
        this->processIncommingMessage(msg);
    }

}
