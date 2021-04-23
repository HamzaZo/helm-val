package utils

import (
	"github.com/fatih/color"
	"helm.sh/helm/v3/pkg/cli/output"
	"io"
)

type valuesFormat output.Format

type ValuesWriter struct {
	Values map[string]interface{}
}

func (v ValuesWriter) WriteTable(out io.Writer) error {
	color.New(color.FgBlue).Fprint(out, "PREVIOUS RELEASE VALUES: \n")
	return output.EncodeYAML(out, v.Values)
}

func (v ValuesWriter) WriteJSON(out io.Writer) error {
	color.New(color.FgBlue).Fprint(out, "PREVIOUS RELEASE VALUES: \n")
	return output.EncodeJSON(out, v.Values)
}

func (v ValuesWriter) WriteYAML(out io.Writer) error {
	color.New(color.FgBlue).Fprint(out, "PREVIOUS RELEASE VALUES: \n")
	return output.EncodeYAML(out, v.Values)
}

func ValuesPrinter(defaultValue output.Format, p *output.Format) *valuesFormat {
	*p = defaultValue
	return (*valuesFormat)(p)
}

func (o *valuesFormat) String() string {
	return string(*o)
}

func (o *valuesFormat) Type() string {
	return "format"
}

func (o *valuesFormat) Set(s string) error {
	outfmt, err := output.ParseFormat(s)
	if err != nil {
		return err
	}
	*o = valuesFormat(outfmt)
	return nil
}
