void Session::disconnected()
{
    qDebug() << "session closed!";
    if(socket != NULL){
        socket->deleteLater();
    }
    exit(0);
}
