#ifndef EVENT_HPP
#define EVENT_HPP
#include <QString>
#include <QMap>
#include <QVariant>

const QString AccountEvent                   = "AccountEvent";
const int AccountRegisterSuccessEvent    = 0;
const int AccountConfirmationSucessEvent = 1;
const QString SessionEvent                   = "SessionEvent";
const int LoginEvent                     = 4;
const int LogoutEvent                    = 5;
const int StateChangeEvent               = 6;
const QString EntityEvent                    = "EntityEvent";
const int NewEntityEvent                 = 7;
const int UpdateEntityEvent              = 8;
const int DeleteEntityEvent              = 9;
const int OpenEntityEvent                = 10;
const int CloseEntityEvent               = 11;
const QString FileEvent                      = "FileEvent";
const int NewFileEvent                   = 12;
const int DeleteFileEvent                = 13;
const int UpdateFileEvent                = 14;
const int OpenFileEvent                  = 15;
const int CloseFileEvent                 = 16;
const int FileEditEvent                  = 17;
const QString DataEvent                      = "DataEvent";
const int DeleteRowEvent                 = 18;
const int NewRowEvent                    = 19;
const int UpdateRowEvent                 = 20;
const int NewDataStoreEvent              = 21;
const int DeleteDataStoreEvent           = 22;
const QString SecurityEvent                  = "SecurityEvent";
const int NewRoleEvent                   = 23;
const int DeleteRoleEvent                = 24;
const int UpdateRoleEvent                = 25;
const QString PrototypeEvent                 = "PrototypeEvent";
const int NewPrototypeEvent              = 26;
const int UpdatePrototypeEvent           = 27;
const int DeletePrototypeEvent           = 28;
const QString ProjectEvent                   = "ProjectEvent";
const QString EmailEvent                     = "EmailEvent";
const QString ServiceEvent                   = "ServiceEvent";
const QString ConfigurationEvent             = "ConfigurationEvent";
const int NewTaskEvent                   = 29;
const int UpdateTaskEvent                = 30;
const QString EventEvent                     = "EventEvent";
const QString LdapEvent                      = "LdapEvent";
const QString OAuth2Event                    = "OAuth2Event";
const QString SchemaEvent                    = "SchemaEvent";
const QString WorkflowEvent                  = "WorkflowEvent";
const int NewBpmnDefinitionsEvent        = 31;
const int DeleteBpmnDefinitionsEvent     = 32;
const int UpdateBpmnDefinitionsEvent     = 33;
const int StartProcessEvent              = 34;

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
