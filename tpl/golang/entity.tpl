{{- /* @lang golang */ -}}
type {{ .Table.Name }} struct {
    {{- range .Table.Columns }}
    {{ .Name }} {{ .LangDataType }}
    {{- end }}
}