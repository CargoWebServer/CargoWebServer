#include "serviceContainer.h"
#include "session.h"
#include <QDir>
#include <QThreadPool>
#include <QPluginLoader>
#include <QDebug>

// Does variable will be set from the Cargo server.
const QString ServiceContainer::defaultApplicationName = "ServiceContainer";
const QString ServiceContainer::defaultOrganizationName = "com.myceliUs";
const unsigned int ServiceContainer::defaultPort = 1234;

ServiceContainer* ServiceContainer::instance = 0;

ServiceContainer *ServiceContainer::getInstance()
{
    if(ServiceContainer::instance == 0){
        ServiceContainer::instance = new ServiceContainer();
    }
    return instance;
}

ServiceContainer::ServiceContainer(QObject* parent) :
    QTcpServer(parent),
    port(ServiceContainer::defaultPort)
{
  // load the server settings.
  this->loadSettings();

  // load the plugins ojects.
  this->loadPluginObjects();
}

ServiceContainer::~ServiceContainer(){
   delete this->settings;
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
void ServiceContainer::incomingConnection(qintptr socketDescriptor){
    // We have a new connection
    qDebug() << socketDescriptor << " Connecting...";

    // Every new session will be run in a newly created thread
    Session *session = new Session(socketDescriptor, this);

    // connect signal/slot
    // once a thread is not needed, it will be beleted later
    connect(session, SIGNAL(finished()), session, SLOT(deleteLater()));

    // Start the session...
    session->start();
}


void ServiceContainer::saveSettings(){

}

void ServiceContainer::loadSettings(){
    // Here I will initialyse the application settings...
    this->settings = new QSettings(QSettings::IniFormat, QSettings::UserScope, organizationName, applicationName);

    if (!settings->contains("port")){
        this->settings->setValue("port", QVariant(ServiceContainer::defaultPort));
    }else{
        this->port = this->settings->value("port").toInt();
    }

    if (!settings->contains("applicationName")){
        this->settings->setValue("applicationName", QVariant(ServiceContainer::defaultApplicationName));
    }else{
        this->applicationName = this->settings->value("applicationName").toString();
    }

    if (!settings->contains("organizationName")){
        this->settings->setValue("organizationName", QVariant(ServiceContainer::defaultOrganizationName));
    }else{
        this->organizationName = this->settings->value("organizationName").toString();
    }
}

void ServiceContainer::setApplicationPath(QString path){
    this->applicationPath = path;
}

QObject* ServiceContainer::getObjectByTypeName(QString typeName){
    return this->objects.value(typeName);
}

void ServiceContainer::loadPluginObjects(){
    QDir pluginsDir(this->applicationPath);

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
