#include "serviceContainer.h"
#include "session.h"
#include <QDir>
#include <QThreadPool>
#include <QPluginLoader>
#include <QDebug>
#include <QCoreApplication>
#include <QtScript/QtScript>

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

ServiceContainer::~ServiceContainer(){
    // close the connections.
    this->close();
}

void ServiceContainer::startServer()
{
    if(!this->listen(QHostAddress::Any, this->port))
    {
        qDebug() << "Could not start server";
    }
    else
    {
        qDebug() << "Listening to port " << this->port << "...";
    }
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

void  ServiceContainer::setPort(quint16 port) {
    this->port = port;
}

QObject* ServiceContainer::getObjectByTypeName(QString typeName){
    return this->objects.value(typeName);
}

void ServiceContainer::loadPluginObjects(){

    QDir pluginsDir(QCoreApplication::applicationDirPath());

#if defined(Q_OS_WIN)
    if (pluginsDir.dirName().toLower() == "debug" || pluginsDir.dirName().toLower() == "release")
        pluginsDir.cdUp();
#elif defined(Q_OS_MAC)
    if (pluginsDir.dirName() == "MacOS") {
        pluginsDir.cdUp();
        pluginsDir.cdUp();
        pluginsDir.cdUp();
    }
#endif
    pluginsDir.cd("plugins");
    foreach (QString fileName, pluginsDir.entryList(QDir::Files)) {
        QPluginLoader pluginLoader(pluginsDir.absoluteFilePath(fileName));
        QObject *plugin = pluginLoader.instance();
        if(plugin != NULL){
            this->objects.insert(plugin->metaObject()->className(), plugin);
            // Append the plugin object.
            qDebug() << "Load object: " << plugin->metaObject()->className();
        }
    }
}

QString ServiceContainer::Ping(){
    return "pong";
}

QVariantList ServiceContainer::ExecuteJsFunction(QVariantList params){
    // first of all i will create a new engine...
    QScriptEngine engine;

    // Now I will put the plugin objects in the engine context.
    for(int i=0; i < this->objects.keys().length(); i++){
        QScriptValue objectValue = engine.newQObject(this->objects.value(this->objects.keys()[i]));
        engine.globalObject().setProperty(this->objects.keys()[i], objectValue);
    }

    // I will now evaluate the script function...
    QScriptValue object = engine.evaluate("({toEvaluate:" + params[0].toString() + "})");
    QScriptValue toEvaluate = object.property("toEvaluate");

    QScriptValueList params_;
    // Now I will set the function parameters...
    for(int i= 1; i < params.length(); i++){
        qDebug() << params.at(i);
        params_.append(engine.newVariant(params.at(i)));
    }

    QScriptValue result = toEvaluate.call(object, params_);
    QVariantList results;
    results.push_back(result.toVariant());

    return results;
}
