
ServiceContainer::~ServiceContainer(){
    // close the connections.
    this->close();
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

QObject* ServiceContainer::getObjectByTypeName(QString typeName){
    return this->objects.value(typeName);
}

void ServiceContainer::loadPluginObjects(){

    QDir pluginsDir(QCoreApplication::applicationDirPath());

#if defined(Q_OS_WIN)
    if (pluginsDir.dirName().toLower() == "debug" || pluginsDir.dirName().toLower() == "release")
        pluginsDir.cdUp();
#elif defined(Q_OS_MAC)
    if (pluginsDir.dirName() == "MacOS") {
        pluginsDir.cdUp();
        pluginsDir.cdUp();
        pluginsDir.cdUp();
    }
#endif
    pluginsDir.cd("plugins");
    foreach (QString fileName, pluginsDir.entryList(QDir::Files)) {
        //qDebug() << "---------> " << pluginsDir.absoluteFilePath(fileName);
        QPluginLoader pluginLoader(pluginsDir.absoluteFilePath(fileName));
        QString iid =  pluginLoader.metaData().value("IID").toString();
        QJsonObject metaData = pluginLoader.metaData().value("MetaData").toObject();
        QObject *plugin = pluginLoader.instance();
        if(plugin != NULL){
            QStringList values = iid.split(".");
            QString className = values.at(values.size()-1);

            this->objects.insert(className, plugin);

            // Keep meta infos...
            this->metaInfos.insert(iid,metaData);

            // Append the plugin object.
            qDebug() << "Load object: " << className;

        }
    }
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
            QString serverSideMethodName = QString(packageName + "." + className + "_" +(*it).toObject()["name"].toString()).replace(".", "_");

            QString serverCode = "function " + serverSideMethodName + "(";


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
            QString serverCodeFunctionCall = className + "." +(*it).toObject()["name"].toString() + "(";
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

                   // Append comma to serverCode.
                   serverCode += name;
                   serverCodeFunctionCall+= name;
                   if (it_ != --parameters.constEnd()){
                       serverCode += ", ";
                       serverCodeFunctionCall+= ", ";
                   }
                   serverCodeFunctionCall += ")";
            }

            serverCode += "){\n";
            // Here Is the server side code...
            serverCode += " return " + serverCodeFunctionCall + "\n";

            serverCode += "}\n";

            // Save the code in the map.
            this->serverCodes.insert(serverSideMethodName, serverCode);

            // Now the execute js call...
            clientCode +=  "    executeJsFunction(\"" + serverSideMethodName + "\", params ";

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

        qDebug() << "Get client code: " << (*it).toObject()["IID"];
    }

    return clientCode;
}

QJsonArray ServiceContainer::GetActionInfos(){
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
    qDebug() << "Ping received!";
    return "pong";
}

QVariantList ServiceContainer::ExecuteJsFunction(QVariantList params){
    // first of all i will create a new engine...
    QScriptEngine engine;

    // Now I will put the plugin objects in the engine context.
    for(int i=0; i < this->objects.keys().length(); i++){
        QScriptValue objectValue = engine.newQObject(this->objects.value(this->objects.keys()[i]));
        engine.globalObject().setProperty(this->objects.keys()[i], objectValue);
    }

    // I will now evaluate the script function...
    QString function = params[0].toString();

    // If I have the name of the function only and not it's code.
    if( function.indexOf("function") == -1){
        // I will find the server side code...
        QMap<QString, QString>::const_iterator it = this->serverCodes.find(function);
        if(it != this->serverCodes.end()){
            function = (*it);
        }
    }

    QScriptValue object = engine.evaluate("({toEvaluate:" + function + "})");
    QScriptValue toEvaluate = object.property("toEvaluate");

    QScriptValueList params_;
    // Now I will set the function parameters...
    for(int i= 1; i < params.length(); i++){
        params_.append(engine.newVariant(params.at(i)));
    }

    QScriptValue result = toEvaluate.call(object, params_);
    QVariantList results;
    results.push_back(result.toVariant());

    return results;
}
