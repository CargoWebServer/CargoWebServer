#ifndef LISTENER_H
#define LISTENER_H

#include <QObject>
#include <QVariant>
#include <QMap>
#include <QString>
#include "event.hpp"

class Listener : public QObject
{
    Q_OBJECT
public:
    explicit Listener(QObject *parent = nullptr);
    virtual QStringList getChannelIds()=0;

public slots:
    virtual void onEvent(QString, int evtNumber, const QMap<QString, QVariant> &evtData) = 0;
};

#endif // LISTENER_H
