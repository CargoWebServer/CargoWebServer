#include "serviceContainer.h"
#include "session.h"
#include <QDir>
#include <QThreadPool>
#include <QPluginLoader>
#include <QDebug>
#include <QUuid>
#include <QCoreApplication>
#include <QtScript/QtScript>
#include "listener.hpp"

// common tcp/ws code.
#include "../serviceContainer.cpp"

// Does variable will be set from the Cargo server.
ServiceContainer* ServiceContainer::instance = 0;

ServiceContainer *ServiceContainer::getInstance()
{
    if(ServiceContainer::instance == 0){
        ServiceContainer::instance = new ServiceContainer();
    }
    return instance;
}

ServiceContainer::ServiceContainer(QObject *parent) :
    QTcpServer(parent)
{
}


// This function is called by QTcpServer when a new connection is available.
void ServiceContainer::incomingConnection(qintptr socketDescriptor){

    // We have a new connection
    qDebug() << socketDescriptor << " Connecting...";

    // I will create the socket here.
    QTcpSocket* socket = new QTcpSocket();
    // set the ID
    if(!socket->setSocketDescriptor(socketDescriptor))
    {
        // something's wrong, we just emit a signal
        qDebug() << "error encounter! TCP/serviceContainer.cpp ln 44";
        return;
    }

    // Every new session will be run in a newly created thread
    Session *session = new Session(socket, this);

    // connect signal/slot
    // once a thread is not needed, it will be beleted later
    connect(session, SIGNAL(finished()), session, SLOT(deleteLater()));

    // Need it to remove session engine from the map.
    connect(session, SIGNAL(end(QString)), this, SLOT(onSessionEnd(QString)));

    // Start the session...
    session->start();

    this->setListeners(session);

}
