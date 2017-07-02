
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
        QPluginLoader pluginLoader(pluginsDir.absoluteFilePath(fileName));
        QObject *plugin = pluginLoader.instance();
        if(plugin != NULL){
            this->objects.insert(plugin->metaObject()->className(), plugin);
            // Append the plugin object.
            qDebug() << "Load object: " << plugin->metaObject()->className();
        }
    }
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
    QScriptValue object = engine.evaluate("({toEvaluate:" + params[0].toString() + "})");
    QScriptValue toEvaluate = object.property("toEvaluate");

    QScriptValueList params_;
    // Now I will set the function parameters...
    for(int i= 1; i < params.length(); i++){
        qDebug() << params.at(i);
        params_.append(engine.newVariant(params.at(i)));
    }

    QScriptValue result = toEvaluate.call(object, params_);
    QVariantList results;
    results.push_back(result.toVariant());

    return results;
}
