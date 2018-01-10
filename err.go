package law

import (
	"bytes"
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
	"text/template"
)

func newAssertionError(typ, desc string) error {
	var buf bytes.Buffer

	type tctx struct {
		Type    string
		Desc    string
		Frame   runtime.Frame
		Callers []runtime.Frame
	}

	callers := callStack(5, 10)
	if len(callers) > 0 {
		err := errorTemplate.Execute(&buf, tctx{
			Type:    typ,
			Desc:    desc,
			Frame:   callers[0],
			Callers: callers[1:len(callers)],
		})
		if err != nil {
			return err
		}
	} else {
		err := errorTemplate.Execute(&buf, tctx{
			Type: typ,
			Desc: desc,
			Frame: runtime.Frame{
				Function: "???",
				File:     "???",
				Line:     0,
			},
		})
		if err != nil {
			return err
		}
	}

	return errors.New(buf.String())
}

func callStack(skip, depth int) (out []runtime.Frame) {
	pc := make([]uintptr, depth)
	n := runtime.Callers(skip, pc)
	pc = pc[:n] // might get less than depth.

	frames := runtime.CallersFrames(pc)
	for {
		frame, more := frames.Next()
		if !strings.Contains(frame.Function, "runtime.") {
			out = append(out, frame)
		}
		if !more {
			break
		}
	}
	return
}

var errorTempRaw = `
!!! {{.Type}} FAILED in {{.Frame.Function | base}}:{{.Frame.File}}:{{.Frame.Line}}
{{- if .Desc}}
!!!  {{.Desc}}
{{- end}}
{{- if .Callers}}
!!!
!!! Call stack:
{{ range .Callers -}}
!!!   {{.Function | base}} in {{.File}}:{{.Line}}
{{end -}}
!!!
{{end -}}

`

var errorTemplate *template.Template

func init() {
	var err error
	et := template.New("errorTemplate")
	et = et.Funcs(template.FuncMap{
		"base": path.Base,
	})
	if errorTemplate, err = et.Parse(errorTempRaw); err != nil {
		err = errors.New(fmt.Sprintf("in assertion error template compile: %s", err))
		panic(err)
	}
}
