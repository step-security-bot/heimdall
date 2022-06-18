package template

import (
	"bytes"
	"net/url"
	"text/template"

	"github.com/Masterminds/sprig/v3"

	"github.com/dadrus/heimdall/internal/heimdall"
	"github.com/dadrus/heimdall/internal/pipeline/subject"
)

type Template string

func (t Template) Render(ctx heimdall.Context, sub *subject.Subject) (string, error) {
	tmpl, err := template.New("Subject").
		Funcs(sprig.TxtFuncMap()).
		Funcs(template.FuncMap{"urlenc": url.QueryEscape}).
		Parse(string(t))
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, sub)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}