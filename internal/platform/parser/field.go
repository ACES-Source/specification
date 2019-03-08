package parser

import (
	"fmt"
	"strings"
)

type Field struct {
	Name         string
	Label        string
	Description  string
	Type         string
	Size         int64
	Required     bool
	ExampleValue string `yaml:"example_value"`
	ExampleHex   string `yaml:"example_hex"`
	Notes        string
}

func (f Field) FieldName() string {
	return strings.Replace(f.Name, " ", "", -1)
}

func (f Field) FieldBytes() string {
	return fmt.Sprintf("%v", f.Size)
}

func (f Field) FormType() string {
	if f.Name == "Payload" {
		return "json.RawAction"
	}

	if f.Type == "[]byte" || f.Type == "byte" {
		return "string"
	}

	return f.Type
}

func (f Field) Tags() string {
	return fmt.Sprintf("`json:\"%s,omitempty\"`", f.JSONFriendly())
}

func (f Field) JSONFriendly() string {
	snake := matchFirstCap.ReplaceAllString(f.Name, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func (f Field) IsData() bool {
	return f.Type == "bin16"
}

func (f Field) IsBytes() bool {
	return f.GoType() == "[]byte" && !f.IsData()
}

func (f Field) IsNumeric() bool {
	s := strings.ToLower(f.Type)

	return strings.HasPrefix(s, "uint") ||
		strings.HasPrefix(s, "float") ||
		strings.HasPrefix(s, "int") ||
		s == "time"
}

func (f Field) IsFloat() bool {
	s := strings.ToLower(f.Type)

	return strings.HasPrefix(f.Type, s)
}

func (f Field) Length() int {
	return int(f.Size)
}

func (f Field) IsNvarchar() bool {
	s := strings.ToLower(f.Type)

	return strings.HasPrefix(s, "nvarchar")
}

func (f Field) GoTypeSingular() string {
	return strings.Replace(f.GoType(), "[]", "", 1)
}

func (f Field) GoType() string {
	s := strings.ToLower(f.Type)

	prefix := ""
	if strings.HasSuffix(s, "[]") {
		prefix = "[]"
		s = f.Type[:len(s)-2]
	}

	switch s {
	case "string":
		if f.Length() > 1 {
			return "[]byte"
		}

		return "byte"

	case "opcode":
		return "byte"

	case "sha",
		"sha256",
		"bin[var]",
		"bin",
		"pushdata_length",
		"text",
		"payload":
		return "[]byte"

	case "time", "timestamp":
		return "uint64"

	case "uint", "int", "float":
		return fmt.Sprintf("%v%s%v", prefix, s, f.Size*8)

	case "nvarchar8":
		return prefix + "Nvarchar8"

	case "nvarchar16":
		return prefix + "Nvarchar16"

	case "nvarchar32":
		return prefix + "Nvarchar32"

	case "header":
		return "Header"

	}

	return fmt.Sprintf("%s%s", prefix, s)
}

func (f Field) IsInternalTypeArray() bool {
	return f.IsInternalType() && strings.HasSuffix(f.Type, "[]")
}

func (f Field) IsNativeTypeArray() bool {
	return !f.IsInternalType() && strings.HasSuffix(f.Type, "[]")
}

func (f Field) IsInternalType() bool {
	s := strings.Replace(f.Type, "[]", "", 1)

	switch s {
	case "Header",
		"VotingSystem",
		"Registry",
		"KeyRole",
		"NotableRole",
		"Amendment",
		"Entity",
		"TargetAddress",
		"Address",
		"TokenSender",
		"TokenReceiver":

		return true
	}

	return strings.HasPrefix(s, "nvarchar")
}

func (f Field) Trimmable() bool {
	if f.Name == "Payload" {
		return false
	}

	if f.Type == "STRING" && f.Length() > 1 {
		return true
	}

	if f.Type == "SHA" {
		return true
	}

	// if f.IsData() {
	// 	return false
	// }

	return false
}
