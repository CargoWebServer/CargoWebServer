#ifndef ACTION_H
#define ACTION_H

#include <QRunnable>
#include <QObject>
#include <QVariant>
#include <QList>
#include "gen/rpc.pb.h"

/**
 * @brief The data struct That structure contain the parameters
 * data.
 */
struct data {
    QString name;
    QVariant value;
    QString typeName;
};

/**
 * @brief Action This class is use to execute an action in it's own thread.
 */
class Action:  public QObject, public QRunnable
{
    Q_OBJECT

public:
    Action(const QString& id, const QString& name, const QString& sessionId);
    ~Action();
    void run();
    void appendParam(QString name, QVariant value, QString typeName);
    QString getId(){return this->id;}

signals:
    /** That signal is emit when the action is completed **/
    void done(com::mycelius::message::Message*);

private:
    /**
     * @brief sessionId The id who generate the action.
     */
    QString sessionId;

    /**
    * @brief id Must be the id of the request message.
    */
    QString id;
    /**
     * @brief name The name of the action.
     */
    QString name;

    /**
     * @brief params The array of parameters.
     */
    QList<data*> params;

};

#endif // ACTION_H
