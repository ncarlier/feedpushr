package plugins

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
	"sync/atomic"
	"time"

	"github.com/ncarlier/feedpushr/v3/pkg/format"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

const defaultEmailSubjectFormat = "[feedpushr] - {{.Title}}"
const defaultEmailBodyFormat = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
	<meta name="viewport" content="width=device-width"/>
	<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
	<title>{{.Title}}</title>
</head>
<body>
	<dl>
		<dt>Title</dt>
		<dd>{{.Title}}</dd>
		<dt>Publication date</dt>
		<dd><time datetime="2018-07-07">{{.Published}}</time></dd>
		<dt>Origin</dt>
		<dd><a href="{{.Link}}">{{.Link}}</a></dd>
	</dl>
	{{if .Content }}
		{{.Content}}
    {{else}}
        <p>{{.Text}}</p>
    {{end}}
</body>
</html>
`

var supportedSMTPConnTypes = map[string]string{
	"plain":        "Plain",
	"tls":          "TLS",
	"tls-insecure": "TLS (insecure)",
}

var emailSpec = model.Spec{
	Name: "email",
	Desc: "New articles are sent by email.\n\nYou can customize the title and the body using the [template engine](https://github.com/ncarlier/feedpushr#output-format).",
	PropsSpec: []model.PropSpec{
		{
			Name: "host",
			Desc: "SMTP host",
			Type: model.Text,
		},
		{
			Name:    "conn",
			Desc:    "SMTP connection type",
			Type:    model.Select,
			Options: supportedSMTPConnTypes,
		},
		{
			Name: "username",
			Desc: "SMTP username",
			Type: model.Text,
		},
		{
			Name: "password",
			Desc: "SMTP password",
			Type: model.Password,
		},
		{
			Name: "from",
			Desc: "From",
			Type: model.Text,
		},
		{
			Name: "to",
			Desc: "To",
			Type: model.Text,
		},
		{
			Name: "subject",
			Desc: "Subject format (by default: \"[feedpushr] {{.Title}}\"",
			Type: model.Text,
		},
		{
			Name: "format",
			Desc: "Body format (by default: HTML page with content)",
			Type: model.Textarea,
		},
	},
}

// EmailOutputPlugin is the STDOUT output plugin
type EmailOutputPlugin struct{}

// Spec returns plugin spec
func (p *EmailOutputPlugin) Spec() model.Spec {
	return emailSpec
}

// Build creates output provider instance
func (p *EmailOutputPlugin) Build(def *model.OutputDef) (model.Output, error) {
	hostname, port, err := net.SplitHostPort(def.Props.Get("host"))
	if err != nil {
		return nil, err
	}
	if port == "" {
		port = "25"
	}
	host := hostname + ":" + port
	conn := def.Props.Get("conn")
	if conn == "" {
		return nil, fmt.Errorf("missing Conn property")
	}
	from := def.Props.Get("from")
	if from == "" {
		return nil, fmt.Errorf("missing From property")
	}
	to := def.Props.Get("to")
	if to == "" {
		return nil, fmt.Errorf("missing To property")
	}
	formatValue := defaultEmailBodyFormat
	if formatProp, ok := def.Props["format"]; ok && formatProp != "" {
		formatValue = fmt.Sprintf("%v", formatProp)
	}
	formatter, err := format.NewTemplateFormatter(def.Hash(), formatValue)
	if err != nil {
		return nil, err
	}
	titleFormatValue := defaultEmailSubjectFormat
	if formatProp, ok := def.Props["subject"]; ok && formatProp != "" {
		titleFormatValue = fmt.Sprintf("%v", formatProp)
	}
	titleFormatter, err := format.NewTemplateFormatter(def.Hash()+"-title", titleFormatValue)
	if err != nil {
		return nil, err
	}
	definition := *def
	definition.Spec = emailSpec

	return &EmailOutputProvider{
		definition:     definition,
		formatter:      formatter,
		titleFormatter: titleFormatter,
		host:           host,
		username:       def.Props.Get("username"),
		password:       def.Props.Get("password"),
		conn:           conn,
		from:           from,
		to:             to,
	}, nil
}

// EmailOutputProvider STDOUT output provider
type EmailOutputProvider struct {
	definition     model.OutputDef
	formatter      format.Formatter
	titleFormatter format.Formatter
	host           string
	username       string
	password       string
	conn           string
	from           string
	to             string
}

func (op *EmailOutputProvider) buildEmailPayload(subject, body string) string {
	// Build email headers
	headers := make(map[string]string)
	headers["From"] = op.from
	headers["To"] = op.to
	headers["Subject"] = subject
	headers["MIME-version"] = "1.0"
	headers["Content-Type"] = `text/html; charset="UTF-8"`
	headers["Date"] = time.Now().Format(time.RFC1123Z)

	// Build email payload
	payload := ""
	for k, v := range headers {
		payload += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	payload += "\r\n" + body
	return payload
}

// Send article to STDOUT.
func (op *EmailOutputProvider) Send(article *model.Article) (bool, error) {
	t, err := op.titleFormatter.Format(article)
	if err != nil {
		atomic.AddUint32(&op.definition.NbError, 1)
		return false, err
	}
	b, err := op.formatter.Format(article)
	if err != nil {
		atomic.AddUint32(&op.definition.NbError, 1)
		return false, err
	}

	// Build email payload
	payload := op.buildEmailPayload(t.String(), b.String())

	// Dial connection
	conn, err := net.DialTimeout("tcp", op.host, 5*time.Second)
	if err != nil {
		return false, err
	}
	// Connect to SMTP server
	hostname, _, _ := net.SplitHostPort(op.host)
	client, err := smtp.NewClient(conn, hostname)
	if err != nil {
		return false, err
	}

	if op.conn == "tls" || op.conn == "tls-insecure" {
		// TLS config
		tlsConfig := &tls.Config{
			InsecureSkipVerify: op.conn == "tls-insecure",
			ServerName:         hostname,
		}
		if err := client.StartTLS(tlsConfig); err != nil {
			return false, err
		}
	}

	// Set auth if needed
	if op.username != "" {
		if err := client.Auth(smtp.PlainAuth("", op.username, op.password, hostname)); err != nil {
			return false, err
		}
	}

	// Set the sender and recipient first
	if err := client.Mail(op.from); err != nil {
		return false, err
	}
	if err := client.Rcpt(op.to); err != nil {
		return false, err
	}

	// Send the email body.
	wc, err := client.Data()
	if err != nil {
		return false, err
	}

	_, err = wc.Write([]byte(payload))
	if err != nil {
		return false, err
	}
	err = wc.Close()
	if err != nil {
		return false, err
	}
	err = client.Quit()
	if err != nil {
		return false, err
	}

	atomic.AddUint32(&op.definition.NbSuccess, 1)
	return true, nil
}

// GetDef return output provider definition
func (op *EmailOutputProvider) GetDef() model.OutputDef {
	return op.definition
}
