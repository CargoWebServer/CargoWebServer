#ifndef EVENT_HPP
#define EVENT_HPP
#include <QString>
#include <QMap>
#include <QVariant>

const QString AccountEvent = "AccountEvent";
const int AccountRegisterSuccessEvent = 0;
const int AccountConfirmationSucessEvent = 1;
const QString SessionEvent = "SessionEvent";
const int LoginEvent = 4;
const int LogoutEvent = 5;
const int StateChangeEvent = 6;
const QString BpmnEvent = "BpmnEvent";
const int NewProcessInstanceEvent = 7;
const int UpdateProcessInstanceEvent = 8;
const int NewDefinitionsEvent = 9;
const int DeleteDefinitionsEvent = 10;
const int UpdateDefinitionsEvent = 11;
const QString EntityEvent = "EntityEvent";
const int NewEntityEvent = 12;
const int UpdateEntityEvent = 13;
const int DeleteEntityEvent = 14;
const int OpenEntityEvent = 15;
const int CloseEntityEvent  = 16;
const QString FileEvent = "FileEvent";
const int NewFileEvent = 17;
const int DeleteFileEvent = 18;
const int UpdateFileEvent = 19;
const int OpenFileEvent = 20;
const int CloseFileEvent = 21;
const QString DataEvent = "DataEvent";
const int DeleteRowEvent = 22;
const int NewRowEvent = 23;
const int UpdateRowEvent = 24;
const int NewDataStoreEvent = 25;
const int DeleteDataStoreEvent = 26;
const QString SecurityEvent = "SecurityEvent";
const int NewRoleEvent = 27;
const int DeleteRoleEvent = 28;
const int UpdateRoleEvent = 29;
const QString PrototypeEvent = "PrototypeEvent";
const int NewPrototypeEvent = 30;
const int UpdatePrototypeEvent = 31;
const int DeletePrototypeEvent = 32;
const QString ProjectEvent = "ProjectEvent";
const QString EmailEvent = "EmailEvent";
const QString ServiceEvent = "ServiceEvent";
const QString ConfigurationEvent = "ConfigurationEvent";
const int NewTaskEvent = 33;
const int UpdateTaskEvent = 34;
const QString EventEvent = "EventEvent";
const QString LdapEvent = "LdapEvent";
const QString OAuth2Event = "OAuth2Event";
const QString SchemaEvent = "SchemaEvent";
const QString WorkflowEvent = "WorkflowEvent";

/**
 * @brief The Event struct Event that came from the network.
 */
struct Event {
  // The event name is a channel id.
  QString name;

  // The event number, ex. NewTaskEvent...
  int number;

  // The event data
  QMap<QString, QVariant> data;

  // The empty constructor.
  Event(){

  }

  // The event default constructor.
  Event(QString name, int number, const QMap<QString, QVariant> &data) :
      name(name),
      number(number),
      data(data)
  {

  }

  Event(const Event& other):
    name(other.name),
    number(other.number),
    data(other.data)
  {
  }

  Event& operator=(const Event& other)
  {
      if(&other == this)
             return *this;
      *this = Event(other);
  }

};

#endif // EVENT_HPP
