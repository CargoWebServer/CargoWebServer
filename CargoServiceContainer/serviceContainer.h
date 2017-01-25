#ifndef TRIPLESERVER_H
#define TRIPLESERVER_H

// Qt stuff here...
#include <QTcpServer>
#include <QSettings>
#include <QStringList>
#include <QMap>

/**
 * @brief Service Container is a TCP server. It's use to interface
 * c++ class functionality over a network.
 */
class ServiceContainer : public QTcpServer
{
    // constant variables...
    static const QString defaultApplicationName;
    static const QString defaultOrganizationName;
    static const unsigned int defaultPort;

    // The instance to the server itself...
    static ServiceContainer* instance;

    Q_OBJECT
public:
    explicit ServiceContainer(QObject *parent = 0);
    void startServer();
    virtual ~ServiceContainer();

    /**
     * A singleton to the server
     **/
    static ServiceContainer *getInstance();

    /**
     * Set path.
     **/
    void setApplicationPath(QString path);

    /**
     * Get an object with his typeName.
     **/
    QObject* getObjectByTypeName(QString typeName);

signals:

protected:
    void incomingConnection(qintptr socketDescriptor);

private:
    // settings...
    QSettings *settings;
    void loadSettings();
    void saveSettings();
    void loadPluginObjects();

    // The port
    int port;
    QString applicationName;
    QString organizationName;

    // The path of the exec...
    QString applicationPath;

    // Object define by plugin...
    QMap<QString, QObject*> objects;

public slots:
	// Do stuff here...

};

#endif // TRIPLESERVER_H
