package html5

import (
	"bytes"
	"strings"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

var stringTmpl = newTextTemplate("string element", "{{ escape . }}",
	texttemplate.FuncMap{
		"escape": EscapeString,
	})

func renderStringElement(ctx renderer.Context, str types.StringElement) ([]byte, error) { //nolint: unparam
	buf := bytes.NewBuffer(nil)
	err := stringTmpl.Execute(buf, str.Content)
	if err != nil {
		return []byte{}, errors.Wrapf(err, "unable to render string")
	}
	result := convert(buf.String(), ellipsis, copyright, trademark, registered)
	return []byte(result), nil
}

func ellipsis(source string) string {
	return strings.Replace(source, "...", "&#8230;&#8203;", -1)
}

func copyright(source string) string {
	return strings.Replace(source, "(C)", "&#169;", -1)
}

func trademark(source string) string {
	return strings.Replace(source, "(TM)", "&#153;", -1)
}

func registered(source string) string {
	return strings.Replace(source, "(R)", "&#174;", -1)
}

type converter func(string) string

func convert(source string, converters ...converter) string {
	result := source
	for _, convert := range converters {
		result = convert(result)
	}
	return result
}
