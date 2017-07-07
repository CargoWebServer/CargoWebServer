TEMPLATE = app
CONFIG += console
CONFIG -= app_bundle
DESTDIR  = ../WebApp/Cargo/bin
QT += websockets script
CONFIG += c++11

HEADERS += \
    action.h \
    WS/serviceContainer.h \
    WS/session.h

SOURCES += \
    action.cpp \
    WS/serviceContainer.cpp \
    WS/session.cpp \
    gen/rpc.pb.cc \
    main.cpp

DEFINES += PORT_NUMBER=9494 WS


INCLUDEPATH += $$PWD/include
DEPENDPATH += $$PWD/lib

LIBS += -L$$PWD/lib/ -llibprotobuf
