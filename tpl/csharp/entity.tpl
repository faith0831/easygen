{{- /* @lang csharp */ -}}
{{- /* @env Namespace 命名空间 */ -}}
namespace {{ .Namespace }}
{ 
    public class {{ .Table.Name }}
    {
        {{- range .Table.Columns }}
        public {{ .LangDataType }} {{ .Name }} { get; set; }
        {{- end }}
    }
}