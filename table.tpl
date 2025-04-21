package plusgorm
type {{.NameCamel}} {{"struct{"}}
 {{range $i,$v:=.Fields}}
   {{- $v.NameCamel}} {{$v.Type}} `json:"{{$v.NameUnderline}}" gorm:"column:{{$v.Name}};{{- $v.GormExtraTag -}}"`
 {{end}}
{{"}"}}