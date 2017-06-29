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

win32:CONFIG(release, debug|release): LIBS += -L$$PWD/OpenSSL/VC_Static/ -llibeay32MD
else:win32:CONFIG(debug, debug|release): LIBS += -L$$PWD/OpenSSL/VC_Static/ -llibeay32MDd

INCLUDEPATH += $$PWD/OpenSSL/VC_Static
DEPENDPATH += $$PWD/OpenSSL/VC_Static

win32-g++:CONFIG(release, debug|release): PRE_TARGETDEPS += $$PWD/OpenSSL/VC_Static/liblibeay32MD.a
else:win32-g++:CONFIG(debug, debug|release): PRE_TARGETDEPS += $$PWD/OpenSSL/VC_Static/liblibeay32MDd.a
else:win32:!win32-g++:CONFIG(release, debug|release): PRE_TARGETDEPS += $$PWD/OpenSSL/VC_Static/libeay32MD.lib
else:win32:!win32-g++:CONFIG(debug, debug|release): PRE_TARGETDEPS += $$PWD/OpenSSL/VC_Static/libeay32MDd.lib


win32:CONFIG(release, debug|release): LIBS += -L$$PWD/OpenSSL/VC_Static/ -lssleay32MD
else:win32:CONFIG(debug, debug|release): LIBS += -L$$PWD/OpenSSL/VC_Static/ -lssleay32MDd

INCLUDEPATH += $$PWD/OpenSSL/VC_Static
DEPENDPATH += $$PWD/OpenSSL/VC_Static

win32-g++:CONFIG(release, debug|release): PRE_TARGETDEPS += $$PWD/OpenSSL/VC_Static/libssleay32MD.a
else:win32-g++:CONFIG(debug, debug|release): PRE_TARGETDEPS += $$PWD/OpenSSL/VC_Static/libssleay32MDd.a
else:win32:!win32-g++:CONFIG(release, debug|release): PRE_TARGETDEPS += $$PWD/OpenSSL/VC_Static/ssleay32MD.lib
else:win32:!win32-g++:CONFIG(debug, debug|release): PRE_TARGETDEPS += $$PWD/OpenSSL/VC_Static/ssleay32MDd.lib


win32: LIBS += -L$$PWD/../../../../../Qt/5.8/msvc2015_64/lib/ -llibprotobuf

INCLUDEPATH += $$PWD/../../../../../Qt/5.8/msvc2015_64/include
DEPENDPATH += $$PWD/../../../../../Qt/5.8/msvc2015_64/include
