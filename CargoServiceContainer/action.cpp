#include "action.h"
#include "WS/serviceContainer.h"

// The Qt stuff here...
#include <QDebug>
#include <QThread>
#include <QMetaObject>
#include <QMetaMethod>
#include <QVariantList>
#include <QJsonArray>
#include <QJsonObject>
#include <QJsonDocument>

QVariant CallMethod(QObject* object, QMetaMethod metaMethod, QVariantList args)
{
    // Convert the arguments
    QVariantList converted;

    // We need enough arguments to perform the conversion.
   QList<QByteArray> methodTypes = metaMethod.parameterTypes();
    if (methodTypes.size() < args.size()) {
        qDebug() << "Insufficient arguments to call" << metaMethod.name();
        return QVariant();
    }

    for (int i = 0; i < methodTypes.size(); i++) {
        const QVariant& arg = args.at(i);

        QByteArray methodTypeName = methodTypes.at(i);
        QByteArray argTypeName = arg.typeName();

        QVariant::Type methodType = QVariant::nameToType(methodTypeName);

        QVariant copy = QVariant::fromValue(arg);

        // If the types are not the same, attempt a conversion. If it
        // fails, we cannot proceed.
        if (copy.type() != methodType) {
            if (copy.canConvert(methodType)) {
                if (!copy.convert(methodType)) {
                    qWarning() << "Cannot convert" << argTypeName
                               << "to" << methodTypeName;
                    return QVariant();
                }
            }
        }

        converted << copy;
    }

    QList<QGenericArgument> arguments;
    for (int i = 0; i < converted.size(); i++) {

        // Notice that we have to take a reference to the argument, else
        // we'd be pointing to a copy that will be destroyed when this
        // loop exits.
        QVariant& argument = converted[i];

        // A const_cast is needed because calling data() would detach
        // the QVariant.
        QGenericArgument genericArgument(
            QMetaType::typeName(argument.userType()),
            const_cast<void*>(argument.constData())
        );

        arguments << genericArgument;
    }

    QVariant returnValue(QMetaType::type(metaMethod.typeName()),
        static_cast<void*>(NULL));

    QGenericReturnArgument returnArgument(
        metaMethod.typeName(),
        const_cast<void*>(returnValue.constData())
    );

    // Perform the call
    bool ok = metaMethod.invoke(
        object,
        Qt::DirectConnection,
        returnArgument,
        arguments.value(0),
        arguments.value(1),
        arguments.value(2),
        arguments.value(3),
        arguments.value(4),
        arguments.value(5),
        arguments.value(6),
        arguments.value(7),
        arguments.value(8),
        arguments.value(9)
    );

    if (!ok) {
        return QVariant();
    } else {
        return returnValue;
    }
}

Action::Action(const QString& id_, const QString& name_) :
    name(name_),
    id(id_)
{
}

Action::~Action(){
    // Clear the memory associated with params.
    for(int i=0; i < this->params.size(); i++){
        delete this->params.at(i);
    }
}

void Action::appendParam(QString name, QVariant value, QString typeName){
    data* param = new data();
    param->name = name;
    param->value = value;
    param->typeName = typeName;
    this->params.append(param);
}

void Action::run()
{
    // When the thread is ready it will execute the fucntion...
    // run the command here...
    // http://doc.qt.io/qt-5/qmetaobject.html
    QVariantList list;

    // Here I will recreate the action prototype from it list of parameter...
    QString prototype;

    if(this->name == "ExecuteJsFunction") {
        prototype = "ExecuteJsFunction(QVariantList)";
        // Append the function parameters here...
        QVariantList arg;
        for(int i=0; i < this->params.size(); i++){
            arg.append(this->params[i]->value);
        }
        list.push_back(arg);
    }else{
        prototype = this->name;
        prototype += "(";
        // no more than 10 paremeter are allow...!!!
        for(int i=0; i < this->params.size(); i++){
            list.append(this->params[i]->value);
            prototype += this->params[i]->typeName;
            if(i < this->params.size() - 1 && this->params.size() > 1){
                prototype += ",";
            }
        }
        prototype += ")";
    }

    // Test object here...
    QObject* obj = ServiceContainer::getInstance();

    // Retreive the object function.
    int index = obj->metaObject()->indexOfSlot(prototype.toStdString().c_str());
    QMetaMethod metaMethod = obj->metaObject()->method(index);
    QVariant retVal = CallMethod(obj, metaMethod, list);

    // TODO test retVal for error and report error instead of response in that case.


    // Wait for the answer...
    com::mycelius::message::Message* result = new com::mycelius::message::Message();
    result->set_type(com::mycelius::message::Message_MessageType_RESPONSE);
    result->set_index(-1);
    result->set_total(1);

    // I will create the response...
    com::mycelius::message::Response*  rsp = new com::mycelius::message::Response();
    rsp->set_id(this->id.toStdString());

    // So here I will create the response and send it back to the caller...
    if(retVal.isValid()){
        com::mycelius::message::Data* d = rsp->add_results();;
        d->set_name("result");
        if(retVal.type() == QMetaType::QStringList){
            // The type is a string list...
            d->set_type(::com::mycelius::message::Data_DataType_JSON_STR);
            QJsonDocument doc;
            doc.setArray(::QJsonArray::fromStringList(retVal.toStringList()));

            // So here I will
            d->set_databytes(doc.toJson().toStdString());
        }else if(retVal.type() == QMetaType::QVariantList){
            // The type is a string list...
            d->set_type(::com::mycelius::message::Data_DataType_JSON_STR);
            QJsonDocument doc;
            doc.setArray(::QJsonArray::fromVariantList(retVal.toList()));

            // So here I will
            d->set_databytes(doc.toJson().toStdString());

        } else if(retVal.type() == QMetaType::Int){
            // The type is a integer...
            d->set_type(::com::mycelius::message::Data_DataType_INTEGER);
            d->set_databytes(retVal.toString().toStdString());
        }else if(retVal.type() == QMetaType::Double){
            // The type is a float...
            d->set_type(::com::mycelius::message::Data_DataType_DOUBLE);
            d->set_databytes(retVal.toString().toStdString());
        }else if(retVal.canConvert(QMetaType::QString)){
            // The type is a string...
            d->set_type(::com::mycelius::message::Data_DataType_STRING);
            d->set_databytes(retVal.toString().toStdString());
        }else if(retVal.canConvert(QMetaType::QJsonObject)){
            // The type is a json object...
            d->set_type(::com::mycelius::message::Data_DataType_JSON_STR);
            QJsonDocument doc;
            doc.setObject(retVal.toJsonObject());

            // So here I will
            d->set_databytes(doc.toJson().toStdString());

        }else if(retVal.canConvert(QMetaType::QJsonArray)){
            // The type is a json array...
            d->set_type(::com::mycelius::message::Data_DataType_JSON_STR);
            QJsonDocument doc;
            doc.setArray(retVal.toJsonArray());

            // So here I will
            d->set_databytes(doc.toJson().toStdString());
        }
    }else{
        qDebug() << "The result is void";
    }

    result->set_allocated_rsp(rsp);

    // The result will be send back as a signal...
    emit done(result);
}
