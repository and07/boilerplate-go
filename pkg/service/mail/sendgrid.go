package mail

import (
	"github.com/and07/boilerplate-go/pkg/utils"
	"github.com/hashicorp/go-hclog"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SGMailService is the sendgrid implementation of our MailService.
type SGMailService struct {
	logger  hclog.Logger
	configs *utils.Configurations
}

// NewSGMailService returns a new instance of SGMailService
func NewSGMailService(logger hclog.Logger, configs *utils.Configurations) *SGMailService {
	return &SGMailService{logger, configs}
}

// CreateMail takes in a mail request and constructs a sendgrid mail type.
func (ms *SGMailService) CreateMail(mailReq *Mail) []byte {

	m := mail.NewV3Mail()

	from := mail.NewEmail("bookite", mailReq.from)
	m.SetFrom(from)

	if mailReq.mtype == MailConfirmation {
		m.SetTemplateID(ms.configs.MailVerifTemplate)
	} else if mailReq.mtype == PassReset {
		m.SetTemplateID(ms.configs.PassResetTemplate)
	}

	p := mail.NewPersonalization()

	tos := make([]*mail.Email, 0)
	for _, to := range mailReq.to {
		tos = append(tos, mail.NewEmail("user", to))
	}

	p.AddTos(tos...)

	p.SetDynamicTemplateData("Username", mailReq.data.Username)
	p.SetDynamicTemplateData("Code", mailReq.data.Code)

	m.AddPersonalizations(p)
	return mail.GetRequestBody(m)
}

// SendMail creates a sendgrid mail from the given mail request and sends it.
func (ms *SGMailService) SendMail(mailReq *Mail) error {

	request := sendgrid.GetRequest(ms.configs.SendGridAPIKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var Body = ms.CreateMail(mailReq)
	request.Body = Body
	response, err := sendgrid.API(request)
	if err != nil {
		ms.logger.Error("unable to send mail", "error", err)
		return err
	}
	ms.logger.Info("mail sent successfully", "sent status code", response.StatusCode)
	return nil
}

// NewMail returns a new mail request.
func (ms *SGMailService) NewMail(from string, to []string, subject string, mailType Type, data *Data) *Mail {
	return &Mail{
		from:    from,
		to:      to,
		subject: subject,
		mtype:   mailType,
		data:    data,
	}
}
