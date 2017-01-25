#include "testobject.h"
#include <QDebug>
TestObject::TestObject(QObject *parent) : QObject(parent)
{

}

QString TestObject::SayHelloTo(QString to){
    qDebug() << "Hello " << to;
}
