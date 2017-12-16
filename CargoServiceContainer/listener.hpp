#ifndef LISTENER_H
#define LISTENER_H

#include <QObject>
#include <QVariant>
#include <QMap>
#include <QString>
#include <QStack>
#include <QThread>
#include <QMutex>
#include <QMutexLocker>
#include <QWaitCondition>

#include "event.hpp"

class Listener : public QThread
{
    // Contain event to process.
    QStack<Event> events;
    QMutex mutex;
    QMutex mutex_;
    QWaitCondition stackNotEmpty;
    bool isRuning;
    void stop();
    void run();

    bool hasEventNext();
    Event getNextEvent();
    void popEvent();

    Q_OBJECT
public:
    explicit Listener(QObject *parent = nullptr);
    ~Listener();
    virtual QStringList getChannelIds()=0;

protected:
    virtual void  processEvent(const Event& evt) = 0;

public slots:
    void onEvent(const Event& evt);
};

#endif // LISTENER_H
