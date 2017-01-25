#include <QCoreApplication>
#include "serviceContainer.h"

using namespace std;

int main(int argc, char *argv[])
{
    QCoreApplication a(argc, argv);

    // Make a server and starts it
    ServiceContainer server;

    // Set the application path...
    server.setApplicationPath(a.applicationDirPath());

    server.startServer();

    return a.exec();
}

