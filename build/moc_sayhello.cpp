/****************************************************************************
** Meta object code from reading C++ file 'sayhello.h'
**
** Created by: The Qt Meta Object Compiler version 67 (Qt 5.8.0)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include "../SayHelloPlugin/sayhello.h"
#include <QtCore/qbytearray.h>
#include <QtCore/qmetatype.h>
#include <QtCore/qplugin.h>
#if !defined(Q_MOC_OUTPUT_REVISION)
#error "The header file 'sayhello.h' doesn't include <QObject>."
#elif Q_MOC_OUTPUT_REVISION != 67
#error "This file was generated using the moc from 5.8.0. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

QT_BEGIN_MOC_NAMESPACE
QT_WARNING_PUSH
QT_WARNING_DISABLE_DEPRECATED
struct qt_meta_stringdata_SayHello_t {
    QByteArrayData data[4];
    char stringdata0[29];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_SayHello_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_SayHello_t qt_meta_stringdata_SayHello = {
    {
QT_MOC_LITERAL(0, 0, 8), // "SayHello"
QT_MOC_LITERAL(1, 9, 10), // "sayHelloTo"
QT_MOC_LITERAL(2, 20, 0), // ""
QT_MOC_LITERAL(3, 21, 7) // "message"

    },
    "SayHello\0sayHelloTo\0\0message"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_SayHello[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
       1,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       0,       // signalCount

 // slots: name, argc, parameters, tag, flags
       1,    1,   19,    2, 0x0a /* Public */,

 // slots: parameters
    QMetaType::QString, QMetaType::QString,    3,

       0        // eod
};

void SayHello::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        SayHello *_t = static_cast<SayHello *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: { QString _r = _t->sayHelloTo((*reinterpret_cast< const QString(*)>(_a[1])));
            if (_a[0]) *reinterpret_cast< QString*>(_a[0]) = _r; }  break;
        default: ;
        }
    }
}

const QMetaObject SayHello::staticMetaObject = {
    { &QObject::staticMetaObject, qt_meta_stringdata_SayHello.data,
      qt_meta_data_SayHello,  qt_static_metacall, Q_NULLPTR, Q_NULLPTR}
};


const QMetaObject *SayHello::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *SayHello::qt_metacast(const char *_clname)
{
    if (!_clname) return Q_NULLPTR;
    if (!strcmp(_clname, qt_meta_stringdata_SayHello.stringdata0))
        return static_cast<void*>(const_cast< SayHello*>(this));
    if (!strcmp(_clname, "SayHelloInterface"))
        return static_cast< SayHelloInterface*>(const_cast< SayHello*>(this));
    if (!strcmp(_clname, "com.mycelius.SayHelloInterface"))
        return static_cast< SayHelloInterface*>(const_cast< SayHello*>(this));
    return QObject::qt_metacast(_clname);
}

int SayHello::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 1)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 1;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 1)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 1;
    }
    return _id;
}

QT_PLUGIN_METADATA_SECTION const uint qt_section_alignment_dummy = 42;

#ifdef QT_NO_DEBUG

QT_PLUGIN_METADATA_SECTION
static const unsigned char qt_pluginMetaData[] = {
    'Q', 'T', 'M', 'E', 'T', 'A', 'D', 'A', 'T', 'A', ' ', ' ',
    'q',  'b',  'j',  's',  0x01, 0x00, 0x00, 0x00,
    0xd0, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00,
    0xbc, 0x00, 0x00, 0x00, 0x1b, 0x03, 0x00, 0x00,
    0x03, 0x00, 'I',  'I',  'D',  0x00, 0x00, 0x00,
    0x1e, 0x00, 'c',  'o',  'm',  '.',  'm',  'y', 
    'c',  'e',  'l',  'i',  'u',  's',  '.',  'S', 
    'a',  'y',  'H',  'e',  'l',  'l',  'o',  'I', 
    'n',  't',  'e',  'r',  'f',  'a',  'c',  'e', 
    0x1b, 0x09, 0x00, 0x00, 0x09, 0x00, 'c',  'l', 
    'a',  's',  's',  'N',  'a',  'm',  'e',  0x00,
    0x08, 0x00, 'S',  'a',  'y',  'H',  'e',  'l', 
    'l',  'o',  0x00, 0x00, 0x1a, 0x00, 0xa1, 0x00,
    0x07, 0x00, 'v',  'e',  'r',  's',  'i',  'o', 
    'n',  0x00, 0x00, 0x00, 0x11, 0x00, 0x00, 0x00,
    0x05, 0x00, 'd',  'e',  'b',  'u',  'g',  0x00,
    0x15, 0x10, 0x00, 0x00, 0x08, 0x00, 'M',  'e', 
    't',  'a',  'D',  'a',  't',  'a',  0x00, 0x00,
    '<',  0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00,
    '8',  0x00, 0x00, 0x00, 0x14, 0x03, 0x00, 0x00,
    0x04, 0x00, 'K',  'e',  'y',  's',  0x00, 0x00,
    ' ',  0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00,
    0x1c, 0x00, 0x00, 0x00, 0x0e, 0x00, 's',  'a', 
    'y',  'h',  'e',  'l',  'l',  'o',  'p',  'l', 
    'u',  'g',  'i',  'n',  0x8b, 0x01, 0x00, 0x00,
    0x0c, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00,
    'p',  0x00, 0x00, 0x00, '8',  0x00, 0x00, 0x00,
    'd',  0x00, 0x00, 0x00, 'T',  0x00, 0x00, 0x00
};

#else // QT_NO_DEBUG

QT_PLUGIN_METADATA_SECTION
static const unsigned char qt_pluginMetaData[] = {
    'Q', 'T', 'M', 'E', 'T', 'A', 'D', 'A', 'T', 'A', ' ', ' ',
    'q',  'b',  'j',  's',  0x01, 0x00, 0x00, 0x00,
    0xd0, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00,
    0xbc, 0x00, 0x00, 0x00, 0x1b, 0x03, 0x00, 0x00,
    0x03, 0x00, 'I',  'I',  'D',  0x00, 0x00, 0x00,
    0x1e, 0x00, 'c',  'o',  'm',  '.',  'm',  'y', 
    'c',  'e',  'l',  'i',  'u',  's',  '.',  'S', 
    'a',  'y',  'H',  'e',  'l',  'l',  'o',  'I', 
    'n',  't',  'e',  'r',  'f',  'a',  'c',  'e', 
    0x15, 0x09, 0x00, 0x00, 0x08, 0x00, 'M',  'e', 
    't',  'a',  'D',  'a',  't',  'a',  0x00, 0x00,
    '<',  0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00,
    '8',  0x00, 0x00, 0x00, 0x14, 0x03, 0x00, 0x00,
    0x04, 0x00, 'K',  'e',  'y',  's',  0x00, 0x00,
    ' ',  0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00,
    0x1c, 0x00, 0x00, 0x00, 0x0e, 0x00, 's',  'a', 
    'y',  'h',  'e',  'l',  'l',  'o',  'p',  'l', 
    'u',  'g',  'i',  'n',  0x8b, 0x01, 0x00, 0x00,
    0x0c, 0x00, 0x00, 0x00, 0x9b, 0x12, 0x00, 0x00,
    0x09, 0x00, 'c',  'l',  'a',  's',  's',  'N', 
    'a',  'm',  'e',  0x00, 0x08, 0x00, 'S',  'a', 
    'y',  'H',  'e',  'l',  'l',  'o',  0x00, 0x00,
    '1',  0x00, 0x00, 0x00, 0x05, 0x00, 'd',  'e', 
    'b',  'u',  'g',  0x00, 0x1a, 0x00, 0xa1, 0x00,
    0x07, 0x00, 'v',  'e',  'r',  's',  'i',  'o', 
    'n',  0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00,
    '8',  0x00, 0x00, 0x00, 0x84, 0x00, 0x00, 0x00,
    0xa0, 0x00, 0x00, 0x00, 0xac, 0x00, 0x00, 0x00
};
#endif // QT_NO_DEBUG

QT_MOC_EXPORT_PLUGIN(SayHello, SayHello)

QT_WARNING_POP
QT_END_MOC_NAMESPACE
