# Package {{.Package}} provides base level structs and validation for
# the protocol.
#
# The code in this file is auto-generated. Do not edit it by hand as it will
# be overwritten when code is regenerated.

{{range .Messages}}
{{.Comment}}

class Action_{{.Name}}({{range .Fields}}{{if eq .Name "Version"}}Versioned{{end}}{{end}}ActionBase):
    ActionPrefix = '{{.Code}}'
{{range .Fields}}{{if eq .Name "Version"}}    ActiveVersion = 0
{{end}}{{end}}
    schema = {
        {{range $i, $v := .Fields}}{{if gt $i 1}}{{if gt $i 2}},
        {{end}}'{{ $v.Name }}': {{ $v.Padding }} [{{ minus $i 2 }}, DAT_{{ $v.Type }}, {{ $v.Length }}]{{end}}{{end}}
    }

    rules = {
        'contractFee': 5000,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
{{range $i, $v := .Fields}}{{if gt $i 2}}        {{if eq $v.Name "Version"}}self.Version = self.ActiveVersion{{else}}self.{{ $v.Name }} = None{{end}}
{{end}}{{end}}
{{end}}

ActionClassMap = {
    {{range $i, $v := .Messages}}{{if $i}},
    {{end}}'{{$v.Code}}': Action_{{$v.Name}}{{end}}
}