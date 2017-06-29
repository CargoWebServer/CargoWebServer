#include <QCoreApplication>
#include <QDebug>

#ifdef WS {
    #include "WS/serviceContainer.h"
#else
    #include "TCP/serviceContainer.h"
#endif

#include <iostream>

using namespace std;

int main(int argc, char *argv[])
{
    QCoreApplication a(argc, argv);

    // Now I will set it port number...
    if(argc == 2){
        int port = atoi(argv[1]);
        ServiceContainer::getInstance()->setPort(port);
    }else {
        ServiceContainer::getInstance()->setPort(PORT_NUMBER);
    }

    // Set the application path...
    ServiceContainer::getInstance()->startServer();

    return a.exec();
}

