package lexer

import (
	"github.com/screw-coding/yaml/token"
	"strings"
	"testing"
)

func TestNextToken_comment_and_multi_documents(t *testing.T) {
	intput := `---
# this is comment

# comment2
---
# comment3
---
# comment3
---
`
	tests := []struct {
		expectedType token.TokenType
	}{
		{token.DocumentStartType},
		{token.CommentType},
		{token.CommentType},
		{token.DocumentEndType},
		{token.DocumentStartType},
		{token.CommentType},
		{token.DocumentEndType},
		{token.DocumentStartType},
		{token.CommentType},
		{token.DocumentEndType},
	}

	l := New(strings.NewReader(intput))
	for i, test := range tests {
		tok, _ := l.nextToken()
		if test.expectedType != tok.Type {
			t.Fatalf("tests[#[%d] - tokentype wrong: expected=%q got=%q", i, test.expectedType, tok.Type)
		}
	}

}

func TestNextToken_mapping(t *testing.T) {
	intput := `
a: b # comment a = b
c: d
e:
 f:
  g: xx # comment xxxx

`
	tests := []struct {
		expectedType token.TokenType
	}{
		{token.DocumentStartType},
		{token.MappingKeyType},
		{token.MappingValueType},
		{token.ScalarType},
		{token.CommentType},
		{token.MappingKeyType},
		{token.MappingValueType},
		{token.ScalarType},
		{token.MappingKeyType},   // e
		{token.MappingValueType}, // :
		{token.MappingKeyType},   // f
		{token.MappingValueType}, // :
		{token.MappingKeyType},   // g
		{token.MappingValueType}, // :
		{token.ScalarType},       // xx
		{token.CommentType},
		{token.DocumentEndType},
	}

	l := New(strings.NewReader(intput))
	for i, test := range tests {
		tok, _ := l.nextToken()
		if test.expectedType != tok.Type {
			t.Fatalf("tests[#[%d] - tokentype wrong: expected=%q got=%q", i, test.expectedType, tok.Type)
		}
	}
}

func TestNextToken_sequence(t *testing.T) {
	intput := `---
a:
 - a1
 - a2
b:
 - name: a1
   age: 18
 - name: a2
   age: 20
`
	tests := []struct {
		expectedType token.TokenType
	}{
		{token.DocumentStartType},
		{token.MappingKeyType},
		{token.MappingValueType},

		{token.SequenceEntryType}, // '-'
		{token.ScalarType},        // a1
		{token.SequenceEntryType},
		{token.ScalarType},

		{token.MappingKeyType},   // b
		{token.MappingValueType}, //:

		{token.SequenceEntryType}, // '-'

		{token.MappingKeyType},   // name
		{token.MappingValueType}, //:
		{token.ScalarType},       // a1

		{token.MappingKeyType},   // age
		{token.MappingValueType}, //:
		{token.ScalarType},       // 18

		{token.SequenceEntryType}, // '-'

		{token.MappingKeyType},   // name
		{token.MappingValueType}, //:
		{token.ScalarType},       // a2

		{token.MappingKeyType},   // age
		{token.MappingValueType}, //:
		{token.ScalarType},       // 20

		{token.DocumentEndType},
	}

	l := New(strings.NewReader(intput))
	for i, test := range tests {
		tok, _ := l.nextToken()
		if test.expectedType != tok.Type {
			t.Fatalf("tests[#[%d] - tokentype wrong: expected=%q got=%q", i, test.expectedType, tok.Type)
		}
	}
}
