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

        // Need it to remove session engine from the map.
        connect(session, SIGNAL(end(QString)), this, SLOT(onSessionEnd(QString)));

        // Start the session...
        session->start();

        // Here I will append the js engine for that session and put object on it.
        QJSEngine *engine = new QJSEngine();
        QMap<QString, QObject*> objects = this->loadPluginObjects();
        for(int i=0; i < objects.keys().length(); i++){
            QJSValue objectValue = engine->newQObject(objects.value(objects.keys()[i]));
            engine->globalObject().setProperty(objects.keys()[i], objectValue);
            // Now with a dynamic cast I will try to convert the object as a listener...
            Listener* listener = reinterpret_cast<Listener*>(objects.value(objects.keys()[i]));
            if(listener != NULL){
                session->registerListener(listener);
            }
        }

        // Keep the reference to the engine.
        this->engines[session->id] = engine;
    }
}

