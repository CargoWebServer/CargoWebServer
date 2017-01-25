#-------------------------------------------------
#
# Project created by QtCreator 2015-09-15T13:07:31
#
#-------------------------------------------------

QT       -= core gui

TARGET = zlib
TEMPLATE = lib
CONFIG += staticlib

unix {
    target.path = /usr/lib
    INSTALLS += target
}

DISTFILES += \
    ../../zlib.3.pdf

HEADERS += \
    ../../zutil.h \
    ../../zlib.h \
    ../../zconf.h \
    ../../trees.h \
    ../../inftrees.h \
    ../../inflate.h \
    ../../inffixed.h \
    ../../inffast.h \
    ../../gzguts.h \
    ../../deflate.h \
    ../../crc32.h

SOURCES += \
    ../../zutil.c \
    ../../uncompr.c \
    ../../trees.c \
    ../../inftrees.c \
    ../../inflate.c \
    ../../inffast.c \
    ../../infback.c \
    ../../gzwrite.c \
    ../../gzread.c \
    ../../gzlib.c \
    ../../gzclose.c \
    ../../deflate.c \
    ../../crc32.c \
    ../../compress.c \
    ../../adler32.c
