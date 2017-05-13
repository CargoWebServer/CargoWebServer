#-------------------------------------------------
#
# Project created by QtCreator 2017-01-25T12:28:07
#
#-------------------------------------------------

TEMPLATE        = lib
CONFIG         += plugin
QT             += widgets core sql
INCLUDEPATH    +=
TARGET          = sayhelloplugin
DESTDIR         = ../WebApp/Cargo/bin/plugins
CONFIG += c++11

HEADERS += \
    sayhelloInterface.hpp \
    sayhello.h

SOURCES += \
    sayhello.cpp
