TEMPLATE = app
CONFIG += console
CONFIG -= app_bundle
DESTDIR  = ../WebApp/Cargo/bin
QT += websockets qml #script
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


unix:!macx:INCLUDEPATH += /usr/local/include
win32:INCLUDEPATH += C:/msys64/mingw64/include

win32: LIBS += -LC:/usr/local/lib/ -lprotobuf.dll
unix:!macx: LIBS += -lprotobuf


