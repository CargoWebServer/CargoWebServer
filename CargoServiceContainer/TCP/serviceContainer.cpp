#include "serviceContainer.h"
#include "session.h"
#include <QDir>
#include <QThreadPool>
#include <QPluginLoader>
#include <QDebug>
#include <QCoreApplication>
#include <QtScript/QtScript>

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
        QJSValue objectValue = engine->newQObject(objects.value(objects.keys()[i]));
        engine->globalObject().setProperty(objects.keys()[i], objectValue);
    }
    // Keep the reference to the engine.
    this->engines[session->id] = engine;
}
