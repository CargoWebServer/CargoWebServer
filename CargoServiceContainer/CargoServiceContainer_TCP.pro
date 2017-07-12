TEMPLATE = app
CONFIG += console
CONFIG -= app_bundle
DESTDIR  = ../WebApp/Cargo/bin
QT += network script
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


INCLUDEPATH += $$PWD/include
DEPENDPATH += $$PWD/lib

win32: LIBS += -L$$PWD/lib/ -llibprotobuf
unix:!macx: LIBS += -lprotobuf
