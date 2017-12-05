TEMPLATE = app
CONFIG += console
CONFIG -= app_bundle
DESTDIR  = ../WebApp/Cargo/bin
QT += network qml #script
CONFIG += c++11

HEADERS += \
    action.h \
    TCP/serviceContainer.h \
    TCP/session.h

SOURCES += \
    action.cpp \
    TCP/serviceContainer.cpp \
    TCP/session.cpp \
    gen/rpc.pb.cc \
    main.cpp \
    serviceContainer.cpp

DEFINES += PORT_NUMBER=9595

unix:!macx:INCLUDEPATH += /usr/local/include
win32:INCLUDEPATH += C:/msys64/mingw64/include

win32: LIBS += -LC:/usr/local/lib/ -lprotobuf.dll
unix:!macx: LIBS += -lprotobuf

