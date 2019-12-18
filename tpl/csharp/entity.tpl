public class {{ .Table.Name }}
{
    {{- range .Table.Columns }}
    public {{ .LangDataType }} {{ .Name }} { get; set; }
    {{- end }}
}