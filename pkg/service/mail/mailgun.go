package mail

import (
	"context"
	"log"
	"time"

	"github.com/and07/boilerplate-go/pkg/utils"
	"github.com/hashicorp/go-hclog"
	"github.com/mailgun/mailgun-go/v4"
)

// SGMailService is the sendgrid implementation of our MailService.
type MGMailService struct {
	logger  hclog.Logger
	configs *utils.Configurations
}

// NewMGMailService returns a new instance of MGMailService
func NewMGMailService(logger hclog.Logger, configs *utils.Configurations) *MGMailService {
	return &MGMailService{logger, configs}
}

// CreateMail takes in a mail request and constructs mail type.
func (ms *MGMailService) CreateMail(mailReq *Mail) []byte {
	return []byte{}
}

// SendMail creates a sendgrid mail from the given mail request and sends it.
func (ms *MGMailService) SendMail(mailReq *Mail) error {

	mg := mailgun.NewMailgun(ms.configs.MailGunDomain, ms.configs.MailGunPrivateAPIKey)

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(mailReq.from, mailReq.subject, mailReq.body, mailReq.to...)

	if mailReq.mtype == MailConfirmation {
		message.SetTemplate(ms.configs.MailVerifTemplate)
	} else if mailReq.mtype == PassReset {
		message.SetTemplate(ms.configs.PassResetTemplate)
	}

	err := message.AddTemplateVariable("Username", mailReq.data.Username)
	if err != nil {
		log.Fatal(err)
	}
	err = message.AddTemplateVariable("Code", mailReq.data.Code)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)
	if err != nil {
		return err
	}
	ms.logger.Info("mail sent successfully", resp, id)
	return nil
}

// NewMail returns a new mail request.
func (ms *MGMailService) NewMail(from string, to []string, subject string, mailType Type, data *Data) *Mail {
	return &Mail{
		from:    from,
		to:      to,
		subject: subject,
		mtype:   mailType,
		data:    data,
	}
}
