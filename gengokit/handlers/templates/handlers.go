package templates

const HandlerMethods = `
{{ with $te := .}}
		{{range $i := .Methods}}
		// {{.Name}} implements Service.
		func (s {{ToLower $te.ServiceName}}Service) {{.Name}}(ctx context.Context, in *{{GoName .RequestType.Name}}) (*{{GoName .ResponseType.Name}}, error){
			var resp {{GoName .ResponseType.Name}}
			resp = {{GoName .ResponseType.Name}}{
				{{range $j := $i.ResponseType.Message.Fields -}}
					// {{GoName $j.Name}}:
				{{end -}}
			}
			return &resp, nil
		}
		{{end}}
{{- end}}
`

const Handlers = `
package {{ToLower .Service.Name}}

import (
	"context"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() {{GoName .Service.Name}}Server {
	return {{ToLower .Service.Name}}Service{}
}

type {{ToLower .Service.Name}}Service struct{}

{{with $te := . }}
	{{range $i := $te.Service.Methods}}
		// {{$i.Name}} implements Service.
		func (s {{ToLower $te.Service.Name}}Service) {{$i.Name}}(ctx context.Context, in *{{GoName $i.RequestType.Name}}) (*{{GoName $i.ResponseType.Name}}, error){
			var resp {{GoName $i.ResponseType.Name}}
			resp = {{GoName $i.ResponseType.Name}}{
				{{range $j := $i.ResponseType.Message.Fields -}}
					// {{GoName $j.Name}}:
				{{end -}}
			}
			return &resp, nil
		}
	{{end}}
{{- end}}
`
