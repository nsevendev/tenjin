package mailer

type Mail struct {
	To      string            `json:"to"`
	Subject string            `json:"subject"`
	Body    string            `json:"body"`
	Type    string            `json:"type,omitempty"`
	Context map[string]string `json:"context,omitempty"`
}

type Mailer struct {
	Host     string
	Port     string
	User     string
	Password string
	From     string
}