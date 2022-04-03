package mail

// Service represents the interface for our mail service.
type Service interface {
	CreateMail(mailReq *Mail) []byte
	SendMail(mailReq *Mail) error
	NewMail(from string, to []string, subject string, mailType Type, data *Data) *Mail
}

type Type int

// List of Mail Types we are going to send.
const (
	MailConfirmation Type = iota + 1
	PassReset
)

// MailData represents the data to be sent to the template of the mail.
type Data struct {
	Username string
	Code     string
}

// Mail represents a email request
type Mail struct {
	from    string
	to      []string
	subject string
	body    string
	mtype   Type
	data    *Data
}
