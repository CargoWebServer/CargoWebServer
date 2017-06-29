#include "serviceContainer.h"
#include "session.h"
#include <QDir>
#include <QThreadPool>
#include <QPluginLoader>
#include <QDebug>
#include <QCoreApplication>
#include <QtScript/QtScript>

#include "../serviceContainer.cpp"

// Does variable will be set from the Cargo server.
ServiceContainer* ServiceContainer::instance = 0;

ServiceContainer *ServiceContainer::getInstance()
{
    if(ServiceContainer::instance == 0){
        ServiceContainer::instance = new ServiceContainer(QStringLiteral("serviceContainer"), QWebSocketServer::NonSecureMode);
    }

    // Connect new connction.
    connect(ServiceContainer::instance, &QWebSocketServer::newConnection, ServiceContainer::instance, &ServiceContainer::onNewConnection);

    // end of services.
    connect(ServiceContainer::instance, &QWebSocketServer::closed, ServiceContainer::instance, &ServiceContainer::closed);

    return ServiceContainer::instance;
}

ServiceContainer::ServiceContainer(const QString &serverName, SslMode secureMode, QObject *parent) :
    QWebSocketServer(serverName,secureMode, parent)
{
    // load the server plugins.
    this->loadPluginObjects();
}

// This function is called by QTcpServer when a new connection is available.
void ServiceContainer::onNewConnection(){

    // We have a new connection
    QWebSocket *socket = this->nextPendingConnection();
    if(socket != NULL){
        // Every new session will be run in a newly created thread
        Session *session = new Session(socket, this);

        // connect signal/slot
        // once a thread is not needed, it will be beleted later
        connect(session, SIGNAL(finished()), session, SLOT(deleteLater()));

        // Start the session...
        session->start();
    }
}

