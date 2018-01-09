#include "serviceContainer.h"
#include "session.h"
#include <QDir>
#include <QThreadPool>
#include <QPluginLoader>
#include <QDebug>
#include <QCoreApplication>
#include <QtScript/QtScript>
#include "listener.hpp"

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

}

// This function is called by QTcpServer when a new connection is available.
void ServiceContainer::onNewConnection(){
    // We have a new connection
    QWebSocket *socket = this->nextPendingConnection();
    if(socket != NULL){
        // Block other thread that want ot use the engines maps.

        // Every new session will be run in a newly created thread
        Session *session = new Session(socket, this);

        // connect signal/slot
        // once a thread is not needed, it will be beleted later
        connect(session, SIGNAL(finished()), session, SLOT(deleteLater()));

        // Need it to remove session engine from the map.
        connect(session, SIGNAL(end(QString)), this, SLOT(onSessionEnd(QString)));

        // Here I will connect the message processor and the session to make message processor process the received message.
        connect(session, SIGNAL(messageReceived(const QByteArray&, QString)), this->messageProcessor, SLOT(processIncommingMessage(const QByteArray&, QString)));

        // When the message is process a response will be emit.
        // Here I will connect the message processor and the session to make message processor process the received message.
        connect(this->messageProcessor, SIGNAL(sendResponse(const QByteArray&, QString)), session, SLOT(sendMessage(const QByteArray&, QString)));

        // Start the session...
        session->start();

        this->setListeners(session);
    }
}

