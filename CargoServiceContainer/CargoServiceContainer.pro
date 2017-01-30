TEMPLATE = app
CONFIG += console
CONFIG -= app_bundle
DESTDIR  = ../WebApp/Cargo/bin
QT += websockets script
CONFIG += c++11

HEADERS += \
    action.h \
    serviceContainer.h \
    session.h

SOURCES += \
    action.cpp \
    main.cpp \
    serviceContainer.cpp \
    session.cpp \
    gen/rpc.pb.cc


win32: LIBS += -L$$PWD/../../../../../Qt/5.8/msvc2015_64/lib/ -llibprotobuf

INCLUDEPATH += $$PWD/../../../../../Qt/5.8/msvc2015_64/include
DEPENDPATH += $$PWD/../../../../../Qt/5.8/msvc2015_64/include

win32:!win32-g++: PRE_TARGETDEPS += $$PWD/../../../../../Qt/5.8/msvc2015_64/lib/libprotobuf.lib
else:win32-g++: PRE_TARGETDEPS += $$PWD/../../../../../Qt/5.8/msvc2015_64/lib/liblibprotobuf.a



unix:!macx: LIBS += -lprotobuf
