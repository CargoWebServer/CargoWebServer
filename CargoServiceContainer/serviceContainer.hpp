#ifndef TRIPLESERVER_H
#define TRIPLESERVER_H

// Qt stuff here...
#include <QtWebSockets/QWebSocketServer>
#include <QSettings>
#include <QStringList>
#include <QMap>
#include <QJsonArray>
#include <QJsonObject>
#include <QJSEngine>
#include <QMutex>
#include <QMutexLocker>

#include "gen/rpc.pb.h"
#include "messageprocessor.hpp"

class Session;

/**
 * @brief Service Container is a TCP server. It's use to interface
 * c++ class functionality over a network.
 */
class ServiceContainer : public QWebSocketServer
{
    // The port
    quint16 port;

    // Contain metadata informations.
    QMap<QString, QJsonObject> metaInfos;

    // The instance to the server itself...
    static ServiceContainer* instance;

    // The server side functions.
    QMap<QString, QString> serverCodes;

    // That contain the list engines assciated with their
    // session id.
    QMap<QString, QJSEngine*> engines;

    // The list of listeners.
    QStringList listeners;

    // The message processor.
    MessageProcessor* messageProcessor;

    // Use it to protect engines map access.
    QMutex mutex;

    // plugins...
    QMap<QString, QObject*> loadPluginObjects();

    // Set the listener.
    void setListeners(Session* session);

    Q_OBJECT
public:
    explicit ServiceContainer(const QString &serverName, SslMode secureMode,
                              QObject *parent = Q_NULLPTR);
    void startServer();
    virtual ~ServiceContainer();

    /**
     * A singleton to the server
     **/
    static ServiceContainer *getInstance();

    /**
     * Set port.
     */
    void setPort(quint16 port);

    /**
     * Get an object with his typeName.
     **/
    QObject* getObjectByTypeName(QString typeName);


private Q_SLOTS:
    /**
     * @brief onNewConnection function called when a connection is open.
     */
    void onNewConnection();
    /**
     * @brief onSessionEnd function called when the session is terminated.
     */
    void onSessionEnd(QString);

public Q_SLOTS:
    //////////////////////////////////////////////
    // Service Container api.
    /////////////////////////////////////////////

    /**
     * @brief Ping
     * @return
     */
    QString Ping();

    /**
     * @brief ExecuteJsFunction
     * That function is use to run JS script. It return a list of QObject.
     * @return A list of object generated by the script.
     */
    QVariantList ExecuteJsFunction(QVariantList);


    /**
     * @brief GetServicesClientCode
     * Return the client side source code to inject in the VM to be able to
     * make remote methode call.
     * @return
     */
    QString GetServicesClientCode();


    /**
     * @brief GetActionInfos
     * Return plugin and array of JSON object of the form:
     * [{"IID":"plugin_1",
     *  "actions":[
     *              { "name":"action_1","doc":"@api 1.0... ",
     *                "parameters":[{"name":"p0", "type":"string", "isArray":"false"}, ...],
     *                "results":[{"name":"r0", "type":"string", "isArray":"false"}, ...]
     *              ]}, ...
     *             ]}
     * ]
     * @return
     */
    QJsonArray GetActionInfos();

    /**
     * @brief GetListeners Return the list of channel to listen at.
     * @return
     */
    QStringList GetListeners();

};

#endif // TRIPLESERVER_H
