{{- /* @lang csharp */ -}}
{{- /* @env Module 功能模块 */ -}}
namespace Digua.Mall.Module.{{.Module}}.Domain
{ 
    public class {{ .Table.Name | camel }}
    {
        {{- range .Table.Columns }}
        /// <summary>
        /// {{ if gt (len .Comment) 0 }}{{ .Comment }}{{ else }}{{ .Name }}{{ end }}
        /// </summary>
        public {{ .LangDataType }} {{ .Name | camel }} { get; set; }
        {{- end }}
    }
}