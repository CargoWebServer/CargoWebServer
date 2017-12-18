#ifndef LISTENER_H
#define LISTENER_H

#include <QObject>
#include "event.hpp"

class Listener : public QObject
{
    Q_OBJECT
public:
    explicit Listener(QObject *parent = nullptr);
    ~Listener();
    virtual QStringList getChannelIds()=0;

public slots:
    virtual void onEvent(const Event& evt) = 0;
};

#endif // LISTENER_H
