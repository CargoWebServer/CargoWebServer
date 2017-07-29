package Server

import (
	"errors"
	"log"

	"code.myceliUs.com/CargoWebServer/Cargo/Entities/CargoEntities"
	"code.myceliUs.com/CargoWebServer/Cargo/Entities/Config"
	"code.myceliUs.com/Utility"

	pop3 "github.com/bytbox/go-pop3"
	gomail "gopkg.in/gomail.v1"
)

type EmailManager struct {

	// info about connection on smtp server...
	m_infos map[string]*Config.SmtpConfiguration
}

var emailManager *EmailManager

func (this *Server) GetEmailManager() *EmailManager {
	if emailManager == nil {
		emailManager = newEmailManager()
	}
	return emailManager
}

/**
 * Carbon copy list...
 */
type CarbonCopy struct {
	Mail string
	Name string
}

/**
 * Attachment file, if the data is empty or nil
 * that means the file is on the server a the given path.
 */
type Attachment struct {
	FileName string
	FileData []byte
}

/**
 * Singleton that return reference to the Smtp service.
 */
func newEmailManager() *EmailManager {

	// Here I will register the type Attachement and Carbon Copy...
	Utility.RegisterType((*Attachment)(nil))
	Utility.RegisterType((*CarbonCopy)(nil))

	emailManager := new(EmailManager)

	return emailManager
}

////////////////////////////////////////////////////////////////////////////////
// Service functions
////////////////////////////////////////////////////////////////////////////////

/**
 * Do intialysation stuff here.
 */
func (this *EmailManager) initialize() {

	log.Println("--> Initialize EmailManager")
	// Create the default configurations
	GetServer().GetConfigurationManager().setServiceConfiguration(this.getId(), -1)

	this.m_infos = make(map[string]*Config.SmtpConfiguration, 0)
	smtpConfigurations := GetServer().GetConfigurationManager().getActiveConfigurationsEntity().GetObject().(*Config.Configurations).GetSmtpConfigs()

	// Smtp server configuration...
	for i := 0; i < len(smtpConfigurations); i++ {
		this.m_infos[smtpConfigurations[i].M_id] = smtpConfigurations[i]
	}

}

func (this *EmailManager) getId() string {
	return "EmailManager"
}

func (this *EmailManager) start() {
	log.Println("--> Start EmailManager")
}

func (this *EmailManager) stop() {
	log.Println("--> Stop EmailManager")
}

/**
 * Validate email
 */
func (this *EmailManager) ValidateEmail(email string) (isValid bool) {
	// TODO implement it...
	isValid = true

	return
}

/**
 * Send mail... The server id is the authentification id...
 */
func (this *EmailManager) sendEmail(id string, from string, to []string, cc []*CarbonCopy, subject string, body string, attachs []*Attachment, bodyType string) (errObj *CargoEntities.Error) {

	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to...)

	// Attach the multiple carbon copy...
	var cc_ []string
	for i := 0; i < len(cc); i++ {
		cc_ = append(cc_, msg.FormatAddress(cc[i].Mail, cc[i].Name))
	}

	if len(cc_) > 0 {
		msg.SetHeader("Cc", cc_...)
	}

	msg.SetHeader("Subject", subject)
	msg.SetBody(bodyType, body)

	for i := 0; i < len(attachs); i++ {
		f, err := gomail.OpenFile(GetServer().GetConfigurationManager().GetApplicationDirectoryPath() + attachs[i].FileName)
		if err == nil {
			msg.Attach(f)
		} else {
			errObj = NewError(Utility.FileLine(), EMAIL_ATTACHEMENT_FAIL_ERROR, SERVER_ERROR_CODE, errors.New("The file '"+GetServer().GetConfigurationManager().GetApplicationDirectoryPath()+attachs[i].FileName+"' could not be attached to email with id '"+id+"'."))
		}

	}

	config := this.m_infos[id]

	mailer := gomail.NewMailer(config.M_hostName, config.M_user, config.M_pwd, config.M_port)

	if err := mailer.Send(msg); err != nil {
		errObj = NewError(Utility.FileLine(), EMAIL_ERROR, SERVER_ERROR_CODE, errors.New("Email with id '"+id+"' failed to send with error '"+err.Error()+"'."))
	}

	return
}

/**
 * That function read the content of the mailbox
 */
func (this *EmailManager) receiveMailFunc(address string, user string, pass string) {

	client, err := pop3.DialTLS(address)

	if err != nil {
		log.Println("Error: %v\n", err)
	} else {
		err = client.Auth(user, pass)
		if err != nil {
			log.Println("Error: %v\n", err)
		} else {
			msgs, sizes, err := client.ListAll()
			if err != nil {
				log.Println("Error: %v\n", err)
			} else {
				for i := 0; i < len(msgs); i++ {
					if err != nil {
						log.Println("Error: %v\n", err)
					} else {
						log.Println("msg:", msgs[i], "size:", sizes[i])
						msgStr, err := client.Retr(msgs[i])
						if err != nil {
							log.Println("Error: %v\n", err)
						} else {
							log.Println(msgStr)
						}

					}
				}
			}
		}
	}

	defer func() {
		client.Quit()
	}()

}

//////////////////////////////////////////////////////////////////////////////////
// Api
//////////////////////////////////////////////////////////////////////////////////

// @api 1.0
// Event handler function.
// @param {interface{}} values The entity to set.
// @scope {public}
// @src
//EmailManager.prototype.onEvent = function (evt) {
//    EventHub.prototype.onEvent.call(this, evt)
//}
func (this *EmailManager) OnEvent(evt interface{}) {
	/** empty function here... **/
}

// @api 1.0
// This function is use to send a email message to a given addresse.
// @param {string} id The server connection id
// @param {string} from The email of the sender
// @param {string} to The destination email's list
// @param {[]interface{}} cc The carbon copy's.
// @param {string} title The email title.
// @param {string} msg The message to send must be html format.
// @param {[]interface} attachs The list of local file to upload to the server and attach to the message.
// @param {string} bodyType text/html or text/plain
// @param {string} messageId The request id that need to access this method.
// @param {string} sessionId The user session.
// @scope {public}
// @param {callback} progressCallback The function is call when chunk of response is received.
// @param {callback} successCallback The function is call in case of success and the result parameter contain objects we looking for.
// @param {callback} errorCallback In case of error.
func (this *EmailManager) SendEmail(id string, from string, to []string, cc []interface{}, subject string, body string, attachs []interface{}, bodyType string, messageId string, sessionId string) {

	// Initialyse the parameters object of not already intialyse.
	var cc_ []*CarbonCopy
	values, err := Utility.InitializeStructures(cc, "[]*Server.CarbonCopy")
	if err == nil {
		cc_ = values.Interface().([]*CarbonCopy)
	}

	var attachs_ []*Attachment
	values, err = Utility.InitializeStructures(attachs, "[]*Server.Attachment")
	if err == nil {
		attachs_ = values.Interface().([]*Attachment)
	}

	var errObj *CargoEntities.Error
	errObj = GetServer().GetSecurityManager().canExecuteAction(sessionId, Utility.FunctionName())
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}

	errObj = this.sendEmail(id, from, to, cc_, subject, body, attachs_, bodyType)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}

	// Wrote message success here.
	log.Println("Message was send to ", to, " by ", from)
	for i := 0; i < len(cc); i++ {
		log.Println("--> cc :", cc_[i].Mail)
	}

	return
}
