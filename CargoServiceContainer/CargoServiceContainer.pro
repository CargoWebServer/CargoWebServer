TEMPLATE = app
CONFIG += console
CONFIG -= app_bundle
DESTDIR  = ../WebApp/Cargo/bin
QT += websockets qml #script
CONFIG += c++11

HEADERS += \
    action.hpp \
    event.hpp \
    listener.hpp \
    messageprocessor.hpp \
    serviceContainer.hpp \
    session.hpp \
    gen/rpc.pb.h

SOURCES += \
    action.cpp \
    listener.cpp \
    main.cpp \
    messageprocessor.cpp \
    serviceContainer.cpp \
    session.cpp \
    gen/rpc.pb.cc

DEFINES += PORT_NUMBER=9494 WS

unix:!macx:INCLUDEPATH += /usr/local/include
win32:INCLUDEPATH += C:/msys64/mingw64/include

win32: LIBS += -LC:/usr/local/lib/ -lprotobuf.dll
unix:!macx: LIBS += -lprotobuf


