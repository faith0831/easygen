{{- /* @lang csharp */ -}}
{{- /* @env Module 模块名称 */ -}}
namespace Digua.Module.{{.ENV.Module}}.Domain
{ 
    public class {{ .Table.Name | pascal }}
    {
        {{- range .Table.Columns }}
        /// <summary>
        /// {{ if gt (len .Comment) 0 }}{{ .Comment }}{{ else }}{{ .Name }}{{ end }}
        /// </summary>
        public {{ .LangDataType }} {{ .Name | pascal }} { get; set; }
        {{- end }}
    }
}