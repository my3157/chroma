package lexers

import (
	. "github.com/alecthomas/chroma" // nolint
)

// Cap'N'Proto Proto lexer.
var CapNProto = Register(MustNewLexer(
	&Config{
		Name:      "Cap&#x27;n Proto",
		Aliases:   []string{"capnp"},
		Filenames: []string{"*.capnp"},
		MimeTypes: []string{},
	},
	Rules{
		"root": {
			{`#.*?$`, CommentSingle, nil},
			{`@[0-9a-zA-Z]*`, NameDecorator, nil},
			{`=`, Literal, Push("expression")},
			{`:`, NameClass, Push("type")},
			{`\$`, NameAttribute, Push("annotation")},
			{`(struct|enum|interface|union|import|using|const|annotation|extends|in|of|on|as|with|from|fixed)\b`, Keyword, nil},
			{`[\w.]+`, Name, nil},
			{`[^#@=:$\w]+`, Text, nil},
		},
		"type": {
			{`[^][=;,(){}$]+`, NameClass, nil},
			{`[[(]`, NameClass, Push("parentype")},
			Default(Pop(1)),
		},
		"parentype": {
			{`[^][;()]+`, NameClass, nil},
			{`[[(]`, NameClass, Push()},
			{`[])]`, NameClass, Pop(1)},
			Default(Pop(1)),
		},
		"expression": {
			{`[^][;,(){}$]+`, Literal, nil},
			{`[[(]`, Literal, Push("parenexp")},
			Default(Pop(1)),
		},
		"parenexp": {
			{`[^][;()]+`, Literal, nil},
			{`[[(]`, Literal, Push()},
			{`[])]`, Literal, Pop(1)},
			Default(Pop(1)),
		},
		"annotation": {
			{`[^][;,(){}=:]+`, NameAttribute, nil},
			{`[[(]`, NameAttribute, Push("annexp")},
			Default(Pop(1)),
		},
		"annexp": {
			{`[^][;()]+`, NameAttribute, nil},
			{`[[(]`, NameAttribute, Push()},
			{`[])]`, NameAttribute, Pop(1)},
			Default(Pop(1)),
		},
	},
))
