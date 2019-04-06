package templating

import (
	"bytes"
	"dddl/deployment"
	"io/ioutil"
	"os"
	"text/template"
)

var tmpl *template.Template

// Setup sets up for templating the docker-compose file.
func Setup() error {
	tmpl = template.New("tmpl")

	rawTmpl, err := ioutil.ReadFile("./docker-compose-template.yaml")
	if err != nil {
		return err
	}

	_, err = tmpl.Parse(string(rawTmpl))
	if err != nil {
		return err
	}

	return nil
}

// Perform performs templating with the specified deployment payload.
func Perform(payload deployment.Payload) error {
	buf := &bytes.Buffer{}

	err := tmpl.Execute(buf, payload)
	if err != nil {
		return err
	}

	f, err := os.Create("./docker-compose.yaml")
	if err != nil {
		return err
	}
	defer f.Close()

	f.Write(buf.Bytes())
	return nil
}
