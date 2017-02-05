#ifndef SAYHELLO_IMPL_H
#define SAYHELLO_IMPL_H

#include <QObject>
#include "sayhelloInterface.hpp"

class SayHello: public QObject, SayHelloInterface
{
    Q_OBJECT
    Q_PLUGIN_METADATA(IID "com.mycelius.SayHelloInterface" FILE "sayhelloplugin.json")
    Q_INTERFACES(SayHelloInterface)

public slots:
    // Slot are callable from JS
    QString sayHelloTo(const QString &message) override;
};

#endif // SAYHELLO_IMPL_H
