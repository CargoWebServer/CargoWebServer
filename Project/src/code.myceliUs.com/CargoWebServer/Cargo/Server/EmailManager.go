package Server

import (
	"errors"
	"log"

	"code.myceliUs.com/CargoWebServer/Cargo/Config/CargoConfig"
	"code.myceliUs.com/CargoWebServer/Cargo/Persistence/CargoEntities"
	"code.myceliUs.com/CargoWebServer/Cargo/Utility"

	pop3 "github.com/bytbox/go-pop3"
	gomail "gopkg.in/gomail.v1"
)

type EmailManager struct {

	// info about connection on smtp server...
	m_infos map[string]CargoConfig.SmtpConfiguration
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

	smtpManager := new(EmailManager)

	return smtpManager
}

/**
 * Do intialysation stuff here.
 */
func (this *EmailManager) Initialyze() {
	this.m_infos = make(map[string]CargoConfig.SmtpConfiguration, 0)
	smtpConfigurations := GetServer().GetConfigurationManager().GetSmtpConfigurations()

	// Smtp server configuration...
	for i := 0; i < len(smtpConfigurations); i++ {
		this.m_infos[smtpConfigurations[i].M_id] = smtpConfigurations[i]
	}

	// Pop configuration...
	//this.ReceiveMailFunc("pop.gmail.com:995", "dave.courtois60@gmail.com", "400zm89a")

}

func (this *EmailManager) Start() {
	log.Println("--> Start EmailManager")
}

func (this *EmailManager) Stop() {
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
func (this *EmailManager) ReceiveMailFunc(address string, user string, pass string) {

	client, err := pop3.DialTLS(address)

	if err != nil {
		log.Fatalf("Error: %v\n", err)
	} else {
		err = client.Auth("dave.courtois60@gmail.com", "400zm89a")
		if err != nil {
			log.Fatalf("Error: %v\n", err)
		} else {
			msgs, sizes, err := client.ListAll()
			if err != nil {
				log.Fatalf("Error: %v\n", err)
			} else {
				for i := 0; i < len(msgs); i++ {
					if err != nil {
						log.Fatalf("Error: %v\n", err)
					} else {
						log.Println("msg:", msgs[i], "size:", sizes[i])
						msgStr, err := client.Retr(msgs[i])
						if err != nil {
							log.Fatalf("Error: %v\n", err)
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

/**
 * Send mail... The server id is the authentification id...
 */
func (this *EmailManager) SendEmail(id string, from string, to []string, cc []*CarbonCopy, subject string, body string, attachs []*Attachment, bodyType string, messageId string, sessionId string) {

	errObj := this.sendEmail(id, from, to, cc, subject, body, attachs, bodyType)
	if errObj != nil {
		GetServer().reportErrorMessage(messageId, sessionId, errObj)
		return
	}
	// Wrote message success here.
	log.Println("Message was send to ", to, " by ", from)
	for i := 0; i < len(cc); i++ {
		log.Println("--> cc :", cc[i].Mail)
	}

	return
}
