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
    // load the server plugins.
    this->loadPluginObjects();
}


// This function is called by QTcpServer when a new connection is available.
void ServiceContainer::incomingConnection(qintptr socketDescriptor){
    // We have a new connection
    qDebug() << socketDescriptor << " Connecting...";

    // Every new session will be run in a newly created thread
    Session *session = new Session(socketDescriptor, this);

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
        QObject* object = objects.value(objects.keys()[i]);
        QJSValue objectValue = engine->newQObject(object);
        engine->globalObject().setProperty(objects.keys()[i], objectValue);
        // Now with a dynamic cast I will try to convert the object as a listener...
        Listener* listener = reinterpret_cast<Listener*>(object);
        if(listener != NULL){
            session->registerListener(listener);
        }
    }
    // Keep the reference to the engine.
    this->engines[session->id] = engine;
}
