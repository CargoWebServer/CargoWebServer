#include <exception>
#include <QJSEngine>
#include <event.hpp>
#include <QThreadPool>

ServiceContainer::~ServiceContainer(){
    // close the connections.
    this->close();
}


/**
 * @brief serializeToByteArray Serialyse the message to an array of bytes...
 * @param msg The proto message.
 * @return
 */
QByteArray serializeToByteArray(google::protobuf::Message *msg){
    QByteArray ra;
    ra.resize(msg->ByteSize());
    msg->SerializeToArray(ra.data(),ra.size());
    return ra;
}

void ServiceContainer::startServer()
{
    if(!this->listen(QHostAddress::Any, this->port))
    {
        qDebug() << "Could not start server";
    }
    else
    {
        qDebug() << "Listening to port " << this->port << "...";
    }
}

void  ServiceContainer::setPort(quint16 port) {
    this->port = port;
}

QMap<QString, QObject*> ServiceContainer::loadPluginObjects(){

    QDir pluginsDir(QCoreApplication::applicationDirPath());
    pluginsDir.cd("plugins");

    // Object define by plugin...
    QMap<QString, QObject*> objects;

    foreach (QString fileName, pluginsDir.entryList(QDir::Files)) {
        if(fileName.indexOf(".dll.a") == -1){ // Do not load the lib but the dll int that particular case.
            QPluginLoader pluginLoader(pluginsDir.absoluteFilePath(fileName));
            QString iid =  pluginLoader.metaData().value("IID").toString();
            QJsonObject metaData = pluginLoader.metaData().value("MetaData").toObject();
            QObject *plugin = pluginLoader.instance();
            if(plugin != NULL){
                QStringList values = iid.split(".");
                if(values.size() > 0){
                    QString className = values.at(values.size()-1);
                    objects.insert(className, plugin);

                    // Keep meta infos...
                    this->metaInfos.insert(iid,metaData);

                    // Append the plugin object.
                    qDebug() << "Load object: " << className;

                }else{
                    qDebug() << pluginLoader.errorString();
                }
            }
        }
    }

    return objects;
}

QString ServiceContainer::GetServicesClientCode(){
    // Here I will generate the code for the client side.
    QJsonArray actionsInfo = this->GetActionInfos();
    QString clientCode;
    QMap<QString, QString> packageNames;

    for(QJsonArray::const_iterator it = actionsInfo.constBegin(); it != actionsInfo.constEnd(); ++it){
        // The IID is written like, com.cargo.AnalyseurCSP_Interface
        QStringList values = (*it).toObject()["IID"].toString().split(".");
        QString packageName;
        QString className;

        for(int i=0; i < values.length(); i++){
            if(i == values.length() - 1){
                // The value is a class
                className = values[i];
                clientCode +=  packageName + "." + className + " = function(service){\n";
                clientCode +=  "    this.service = service;\n"; // Needed to call the function service.executeJsFunction()...
            }else{
                // value is a namespace
                if(i==0){
                    packageName =  values[i];
                }else{
                    packageName += "." + values[i];
                }
                if(!packageNames.contains(packageName)){
                    packageNames.insert(packageName, packageName);
                    if(i==0){
                        clientCode +="var ";
                    }
                    clientCode += packageName + " = {};\n";
                }
            }
        }

        clientCode += " return this;\n";
        clientCode += "}\n\n";

        // Now the function.
        QJsonArray functions = (*it).toObject()["actions"].toArray();

        for(QJsonArray::const_iterator it=functions.constBegin(); it != functions.constEnd(); ++it){
            // Now the callback's...
            QStringList callbacks;
            QJsonArray docs =  (*it).toObject()["doc"].toArray();

            for(QJsonArray::const_iterator it_=docs.constBegin(); it_ != docs.constEnd(); ++it_){
                QString line = (*it_).toString();
                clientCode += "// " + line + "\n";
                if(line.indexOf("@param {callback} successCallback") != -1){
                    callbacks.append("successCallback");
                }else if(line.indexOf("@param {callback} progressCallback") != -1){
                    callbacks.append("progressCallback");
                }else if(line.indexOf("@param {callback} errorCallback") != -1){
                    callbacks.append("errorCallback");
                }
            }

            // The server side methode name.
            clientCode += packageName + "." + className + ".prototype." +(*it).toObject()["name"].toString() + "=function(";

            // Now the parameters...
            QJsonArray parameters = (*it).toObject()["parameters"].toArray();

            for(QJsonArray::const_iterator it_=parameters.constBegin(); it_ != parameters.constEnd(); ++it_){
                clientCode += (*it_).toObject()["name"].toString();
                if( it_  < --parameters.constEnd()){
                    clientCode += ", ";
                }
            }

            for(int i=0; i < callbacks.size(); i++){
                clientCode += ", " + callbacks[i];
            }

            clientCode += ", caller){\n";
            // Now the function body...
            clientCode +=  "    var params = []\n";
            QString serverSideMethodName = className + "." +(*it).toObject()["name"].toString();
            for(QJsonArray::const_iterator it_=parameters.constBegin(); it_ != parameters.constEnd(); ++it_){
                   QString name = (*it_).toObject()["name"].toString();
                   QString typeName = (*it_).toObject()["type"].toString();
                   bool isArray = (*it_).toObject()["isArray"].toBool();
                   if(isArray){
                       clientCode +=  "    params.push(createRpcData(" + name + ", \"JSON_STR\", \"" + name + "\", \""+ typeName +"\"))\n";
                   }else if(typeName == "string"){
                       clientCode +=  "    params.push(createRpcData(" + name + ", \"STRING\", \"" + name + "\", \""+ typeName +"\"))\n";
                   }else if(typeName == "double"){
                       clientCode +=  "    params.push(createRpcData(" + name + ", \"DOUBLE\", \"" + name + "\", \""+ typeName +"\"))\n";
                   }else if(typeName == "[]int8"){
                       clientCode +=  "    params.push(createRpcData(" + name + ", \"BYTES\", \"" + name + "\", \""+ typeName +"\"))\n";
                   }else if(typeName == "int"){
                       clientCode +=  "    params.push(createRpcData(" + name + ", \"INTEGER\", \"" + name + "\", \""+ typeName +"\"))\n";
                   }else if(typeName == "bool"){
                       clientCode +=  "    params.push(createRpcData(" + name + ", \"BOOLEAN\", \"" + name + "\", \""+ typeName +"\"))\n";
                   }else{
                       clientCode +=  "    params.push(createRpcData(" + name + ", \"JSON_STR\", \"" + name + "\", \""+ typeName +"\"))\n";
                   }
            }

            // Now the execute js call...
            clientCode +=  "    service.executeJsFunction(\"" + serverSideMethodName + "\", params ";

            // I will now append the callback...
            if(callbacks.contains("progressCallback")){
                clientCode +=  ", function(index, total, caller){ //Progress Callback\n";
                clientCode +=  "        caller.progressCallback(index, total, caller.caller)\n";
                clientCode +=  "     }\n";
            }else{
                clientCode +=  ", undefined // Progress Callback\n";
            }

            QString caller = "{";
            for(int i = 0; i < callbacks.size(); i++){
                caller +=  "\"" + callbacks[i] + "\" : " + callbacks[i];
                if(i < callbacks.size() - 1){
                    caller += ",";
                }
            }
            caller += ", \"caller\":caller}";

            if(callbacks.contains("successCallback")){
                clientCode +=  "    , function(results, caller){ //Success Callback\n";
                bool isArray = false;

                // Must contain only one return value, this is c++
                QJsonArray parameters = (*it).toObject()["results"].toArray();
                for(QJsonArray::const_iterator it_=parameters.constBegin(); it_ != parameters.constEnd(); ++it_){
                    bool isArray = (*it_).toObject()["isArray"].toBool();
                }

                if(isArray){
                    clientCode +=  "        caller.successCallback(results, caller.caller)\n";
                }else{
                    clientCode +=  "        caller.successCallback(results[0], caller.caller)\n";
                }
                clientCode +=  "     }\n";
            }else{
                clientCode +=  "    , undefined // Success Callback\n";
            }

            if(callbacks.contains("errorCallback")){
                clientCode +=  "    , function(errObj, caller){ //Error Callback\n";
                clientCode +=  "        caller.errorCallback(errObj, caller.caller)\n";
                clientCode +=  "     }\n";
            }else{
                clientCode +=  "    , undefined // Error Callback\n";
            }

            clientCode +=  "    , " + caller + "\n";
            clientCode += " , this.service.conn.id)\n";

            clientCode += "}\n";

        }

       // qDebug() << "Get client code: " << (*it).toObject()["IID"];
    }
    return clientCode;
}

QJsonArray ServiceContainer::GetActionInfos(){
    QMutexLocker ml(&this->mutex);
    QJsonArray actionInfos;
    QMapIterator<QString, QJsonObject> i(this->metaInfos);
    while (i.hasNext()) {
        i.next();
        QJsonObject info;
        info["IID"] = i.key();
        info["actions"] = i.value().value("actions").toArray();
        actionInfos.append(info);
    }
    return actionInfos;
}

QString ServiceContainer::Ping(){
    return "pong";
}

void  ServiceContainer::onSessionEnd(QString sessionId){
    QMutexLocker ml(&this->mutex);
    delete this->engines[sessionId]; // Clear memory
    this->engines.remove(sessionId); // remove from the map.
}

QVariantList ServiceContainer::ExecuteJsFunction(QVariantList params){
    QMutexLocker ml(&this->mutex);

    // first of all i will create a new engine...
    QVariantList results;

    // I will now evaluate the script function...
    QString sessionId = params[0].toString();
    QJSEngine* engine = this->engines[sessionId];

    // Now I will put the plugin objects in the engine context.
    QString function = params[1].toString();
    QJSValue toEvaluate = engine->evaluate(function);
    QJSValueList args;

    for(int i= 2; i < params.length(); i++){
        args.append(engine->toScriptValue<QVariant>(params.at(i)));
    }
    try {
        QJSValue result = toEvaluate.call(args);
        results.push_back(result.toVariant());
    }
    catch (std::exception & e) {
       // deal with it
       qDebug()<< "Script error found!!!!" << e.what();
       // Here I will get the exception information.
       QJsonObject errObj;
       errObj["TYPENAME"] = "CargoEntities.Error";
       errObj["M_errorPath"] = "Line 313 serviceContainer.cpp";
       errObj["M_code"] = "EXECUTE_JS_FUNCTION_ERROR";
       errObj["M_body"] = QString(e.what());
       results.push_back(errObj);
    }
    return results;
}

void ServiceContainer::setListeners(Session* session){
    QMutexLocker ml(&this->mutex);
    // Here I will append the js engine for that session and put object on it.
    QJSEngine *engine = new QJSEngine();
    QMap<QString, QObject*> objects = this->loadPluginObjects();
    for(int i=0; i < objects.keys().length(); i++){
        QJSValue objectValue = engine->newQObject(objects.value(objects.keys()[i]));
        engine->globalObject().setProperty(objects.keys()[i], objectValue);
        // Now with a dynamic cast I will try to convert the object as a listener...
        Listener* listener = reinterpret_cast<Listener*>(objects.value(objects.keys()[i]));
        connect(session, SIGNAL(onEvent(const Event&)), listener, SLOT(onEvent(const Event&)));

        // Register the listener
        QStringList channelIds = listener->getChannelIds();
         for(int i=0; i < channelIds.length(); i++){
             if(!this->listeners.contains(channelIds[i])){
                 this->listeners.push_back(channelIds[i]);
             }
         }
    }

    // Keep the reference to the engine.
    this->engines[session->id] = engine;
}

QStringList ServiceContainer::GetListeners(){
    return this->listeners;
}
