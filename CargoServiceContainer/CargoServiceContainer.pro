TEMPLATE = app
CONFIG += console
CONFIG -= app_bundle

QT += network
QMAKE_CXXFLAGS += -Wl,--stack,4194304000
CONFIG += c++11

HEADERS += \
    action.h \
    qcompressor.h \
    serviceContainer.h \
    session.h

SOURCES += \
    action.cpp \
    main.cpp \
    qcompressor.cpp \
    serviceContainer.cpp \
    session.cpp \
    gen/rpc.pb.cc

INCLUDEPATH += $$PWD/ ../include include
DEPENDPATH += $$PWD/lib
unix:!macx: LIBS += -lprotobuf -lprotobuf-lite -lprotoc -lz
win32: LIBS += -L$$PWD/lib -lprotobuf -lprotobuf-lite -lprotoc -lz

