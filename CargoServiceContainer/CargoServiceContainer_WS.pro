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


win32: LIBS += -L$$PWD/lib/ -llibprotobuf


INCLUDEPATH += $$PWD/../../../../../Qt/5.8/msvc2015_64/include $$PWD/include/

win32:!win32-g++: PRE_TARGETDEPS += $$PWD/../../../../../Qt/5.8/msvc2015_64/lib/libprotobuf.lib
else:win32-g++: PRE_TARGETDEPS += $$PWD/../../../../../Qt/5.8/msvc2015_64/lib/liblibprotobuf.a

unix:!macx: LIBS += -lprotobuf

win32:CONFIG(release, debug|release): LIBS += -L$$PWD/lib/ -llibeay32MD
else:win32:CONFIG(debug, debug|release): LIBS += -L$$PWD/lib/ -llibeay32MDd

INCLUDEPATH += $$PWD/OpenSSL/VC_Static
DEPENDPATH += $$PWD/OpenSSL/VC_Static

win32-g++:CONFIG(release, debug|release): PRE_TARGETDEPS += $$PWD/lib/liblibeay32MD.a
else:win32-g++:CONFIG(debug, debug|release): PRE_TARGETDEPS += $$PWD/lib/liblibeay32MDd.a
else:win32:!win32-g++:CONFIG(release, debug|release): PRE_TARGETDEPS += $$PWD/lib/libeay32MD.lib
else:win32:!win32-g++:CONFIG(debug, debug|release): PRE_TARGETDEPS += $$PWD/lib/libeay32MDd.lib


win32:CONFIG(release, debug|release): LIBS += -L$$PWD/lib/ -lssleay32MD
else:win32:CONFIG(debug, debug|release): LIBS += -L$$PWD/lib/ -lssleay32MDd

INCLUDEPATH += $$PWD/OpenSSL/VC_Static
DEPENDPATH += $$PWD/OpenSSL/VC_Static

win32-g++:CONFIG(release, debug|release): PRE_TARGETDEPS += $$PWD/lib/libssleay32MD.a
else:win32-g++:CONFIG(debug, debug|release): PRE_TARGETDEPS += $$PWD/lib/libssleay32MDd.a
else:win32:!win32-g++:CONFIG(release, debug|release): PRE_TARGETDEPS += $$PWD/lib/ssleay32MD.lib
else:win32:!win32-g++:CONFIG(debug, debug|release): PRE_TARGETDEPS += $$PWD/lib/ssleay32MDd.lib
