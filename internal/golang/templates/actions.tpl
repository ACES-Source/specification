package protocol

import (
	"bytes"
	"fmt"
	"strings"
)

const (
{{- range .}}
	{{.CodeNameComment}}
	{{.CodeName}} = "{{.Code}}"
{{ end }}

	// ComplianceActionFreeze identifies a freeze type
	ComplianceActionFreeze = byte('F')

	// ComplianceActionThaw identifies a thaw type
	ComplianceActionThaw = byte('T')

	// ComplianceActionConfiscation identifies a confiscation type
	ComplianceActionConfiscation = byte('C')

	// ComplianceActionReconciliation identifies a reconcilation type
	ComplianceActionReconciliation = byte('R')
)

// TypeMapping holds a mapping of action codes to action types.
func TypeMapping(code string) OpReturnMessage {
	switch(code) {
{{- range . }}
	case Code{{.Name}}:
		result := {{.Name}}{}
		return &result
{{- end }}
	default:
		return nil
	}
}

{{ range $action := . }}

{{comment (print .Name " " .Metadata.Description) "//"}}
type {{.Name}} struct {
{{ range .Fields }}	{{ .FieldName }} {{ .FieldGoType }} {{ .GoTags }} // {{ .FieldDescription }}
{{ end -}}
}

{{ if .Constructor -}}
// New{{.Name}} returns a new {{.Name}} with specified values set.
func New{{.Name}}({{ range $i, $c := .Constructor }}{{if $i}}, {{end}}{{ .ConstructorName }} {{ .ConstructorGoType }}{{ end -}}) *{{.Name}} {
	result := {{.Name}}{}
	{{ range .Constructor -}}
	result.{{ .ConstructorField -}}
	{{ if eq .ConstructorSetMethod "=" }} = {{ .ConstructorName }}{{ else }}.{{ .ConstructorSetMethod }}({{ .ConstructorName }}){{ end }}
	{{ end -}}
	return &result
}
{{ end -}}

{{ range .Functions }}// {{.FunctionName}} {{ .FunctionDescription }}.
func (action *{{ $action.Name }}) {{.FunctionName}}({{ range $i, $c := .FunctionParams }}{{if ne .ParamName "hidden"}}{{if $i}}, {{end}}{{ .ParamName }} {{ .ParamGoType }}{{ end -}}{{ end -}}) {
{{ if eq .FunctionType "set" -}}
	{{ range .FunctionParams -}}
	action.{{ .ParamField -}}
	{{ if eq .ParamSetMethod "=" }} = {{ .ParamValue }}{{ else }}.{{ .ParamSetMethod }}({{ .ParamValue }}){{ end }}
	{{ end -}}
{{ else if eq .FunctionType "append" -}}
	action.{{ .FunctionAppendType }}s = append(action.{{ .FunctionAppendType }}s, *New{{ .FunctionAppendType }}({{ range $i, $param := .FunctionParams }}{{if $i}}, {{end}}{{ .ParamName }}{{ end }}))
{{ end -}}
}
{{ end }}

// Type returns the type identifer for this message.
func (action {{.Name}}) Type() string {
	return Code{{.Name}}
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *{{.Name}}) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *{{.Name}}) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

{{- range $f, $field := .PayloadFields }}
{{- $fieldName := .FieldName }}
	// {{.FieldName}} ({{.FieldGoType}})
{{- if ne (len $field.IncludeIfTrue) 0 }}
	if action.{{ $field.IncludeIfTrue }} {
{{- else if ne (len $field.IncludeIfFalse) 0 }}
	if !action.{{ $field.IncludeIfFalse }} {
{{- else if ne (len $field.IncludeIf.Field) 0 }}
	if {{ range $j, $include := $field.IncludeIf.Values }}{{ if (ne $j 0) }} ||{{ end }} action.{{$field.IncludeIf.Field}} == '{{ $include }}'{{ end }} {
{{- else }}
	{
{{- end }}
{{- if .IsVarChar }}
		if err := WriteVarChar(buf, action.{{.FieldName}}, {{.Length}}); err != nil {
			return nil, err
		}
{{- else if .IsFixedChar }}
		if err := WriteFixedChar(buf, action.{{.FieldName}}, {{.Length}}); err != nil {
			return nil, err
		}
{{- else if .IsVarBin }}
		if err := WriteVarBin(buf, action.{{.FieldName}}, {{.Length}}); err != nil {
			return nil, err
		}
{{- else if .IsInternalTypeArray }}
		if err := WriteVariableSize(buf, uint64(len(action.{{.FieldName}})), {{.Length}}, 8); err != nil {
			return nil, err
		}
		for _, value := range action.{{.FieldName}} {
			b, err := value.Serialize()
			if err != nil {
				return nil, err
			}

			if err := write(buf, b); err != nil {
				return nil, err
			}
		}
{{- else if .IsNativeTypeArray }}
		if err := WriteVariableSize(buf, uint64(len(action.{{.FieldName}})), {{.Length}}, 8); err != nil {
			return nil, err
		}
		for _, value := range action.{{.FieldName}} {
			if err := write(buf, value); err != nil {
				return nil, err
			}
		}
{{- else if .IsInternalType }}
		b, err := action.{{.Name}}.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
{{- else if .IsBytes }}
		if err := write(buf, pad(action.{{.FieldName}}, {{.Length}})); err != nil {
			return nil, err
		}
{{- else}}
		if err := write(buf, action.{{.FieldName}}); err != nil {
			return nil, err
		}
{{- end }}
	}
{{ end }}
	return buf.Bytes(), nil
}

// write populates the fields in {{.Name}} from the byte slice
func (action *{{.Name}}) write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

{{- range $f, $field := .PayloadFields }}
	// {{.FieldName}} ({{.FieldGoType}})
{{- if ne (len $field.IncludeIfTrue) 0 }}
	if action.{{ $field.IncludeIfTrue }} {
{{- else if ne (len $field.IncludeIfFalse) 0 }}
	if !action.{{ $field.IncludeIfFalse }} {
{{- else if ne (len $field.IncludeIf.Field) 0 }}
	if {{ range $j, $include := $field.IncludeIf.Values }}{{ if (ne $j 0) }} ||{{ end }} action.{{$field.IncludeIf.Field}} == '{{ $include }}'{{ end }} {
{{- else }}
	{
{{- end }}
{{- if .IsVarChar }}
		var err error
		action.{{.FieldName}}, err = ReadVarChar(buf, {{.Length}})
		if err != nil {
			return 0, err
		}
{{- else if .IsFixedChar }}
		var err error
		action.{{.FieldName}}, err = ReadFixedChar(buf, {{.Length}})
		if err != nil {
			return 0, err
		}
{{- else if .IsVarBin }}
		var err error
		action.{{.FieldName}}, err = ReadVarBin(buf, {{.Length}})
		if err != nil {
			return 0, err
		}
{{- else if .IsInternalTypeArray }}
		size, err := ReadVariableSize(buf, {{.Length}}, 8)
		if err != nil {
			return 0, err
		}
		action.{{.FieldName}} = make([]{{.SingularType}}, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue {{.SingularType}}
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			action.{{.FieldName}} = append(action.{{.FieldName}}, newValue)
		}
{{- else if .IsNativeTypeArray }}
		size, err := ReadVariableSize(buf, {{.Length}}, 8)
		if err != nil {
			return 0, err
		}
		action.{{.FieldName}} = make({{.FieldGoType}}, size, size)
		if err := read(buf, &action.{{.FieldName}}); err != nil {
			return 0, err
		}
{{- else if .IsInternalType }}
		if err := action.{{.FieldName}}.Write(buf); err != nil {
			return 0, err
		}
{{- else if or .IsBytes .IsData }}
		action.{{.FieldName}} = make([]byte, {{.Length}})
		if err := readLen(buf, action.{{.FieldName}}); err != nil {
			return 0, err
		}
{{- else if .Trimmable }}
		action.{{.FieldName}} = bytes.Trim(action.{{.FieldName}}, "\x00")
{{- else }}
		if err := read(buf, &action.{{.FieldName}}); err != nil {
			return 0, err
		}
{{- end }}
	}
{{ end }}
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action {{.Name}}) PayloadMessage() (PayloadMessage, error) {
{{- if .HasAssetPayload }}
	p := AssetTypeMapping(action.AssetType)
	if p == nil {
		return nil, fmt.Errorf("Undefined asset type : %s", action.AssetType)
	}

	if _, err := p.Write(action.AssetPayload); err != nil {
		return nil, err
	}

	return p, nil
{{- else if .HasMessagePayload }}
	p := MessageTypeMapping(action.MessageType)
	if p == nil {
		return nil, fmt.Errorf("Undefined message type : %s", action.MessageType)
	}

	if _, err := p.Write(action.MessagePayload); err != nil {
		return nil, err
	}

	return p, nil
{{ else }}
	return nil, nil
{{ end -}}
}

func (action {{.Name}}) String() string {
	vals := []string{}
{{ range .Fields -}}
	{{- if eq .Type "STRING" }}
	vals = append(vals, fmt.Sprintf("{{.FieldName}}:\"%v\"", string(action.{{.FieldName}})))
	{{- else if .IsNumeric }}
	vals = append(vals, fmt.Sprintf("{{.FieldName}}:%v", action.{{.FieldName}}))
	{{- else if eq .Type "SHA" }}
	vals = append(vals, fmt.Sprintf("{{.FieldName}}:\"%x\"", action.{{.FieldName}}))
	{{- else if eq .FieldGoType "[]byte" }}
	vals = append(vals, fmt.Sprintf("{{.FieldName}}:%#x", action.{{.FieldName}}))
	{{- else }}
	vals = append(vals, fmt.Sprintf("{{.FieldName}}:%#+v", action.{{.FieldName}}))
	{{- end }}{{ end }}

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
{{ end -}}
