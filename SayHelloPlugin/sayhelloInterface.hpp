#ifndef SAYHELLO_HPP
#define SAYHELLO_HPP

class SayHelloInterface
{
public:
    virtual ~SayHelloInterface() {}

public slots:
    // Slot are callable from JS
    virtual QString sayHelloTo(const QString &message) = 0;
};

#define sayHelloInterface_iid "com.mycelius.SayHelloInterface"

Q_DECLARE_INTERFACE(SayHelloInterface, sayHelloInterface_iid)

#endif // SAYHELLO_HPP
