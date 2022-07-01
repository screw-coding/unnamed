package token

type Character rune

const (
	// SequenceEntryCharacter character for sequence entry
	SequenceEntryCharacter Character = '-'
	// MappingKeyCharacter character for mapping key
	MappingKeyCharacter Character = '?'
	// MappingValueCharacter character for mapping value
	MappingValueCharacter Character = ':'
	// CollectEntryCharacter character for collect entry
	CollectEntryCharacter Character = ','
	// SequenceStartCharacter character for sequence start
	SequenceStartCharacter Character = '['
	// SequenceEndCharacter character for sequence end
	SequenceEndCharacter Character = ']'
	// MappingStartCharacter character for mapping start
	MappingStartCharacter Character = '{'
	// MappingEndCharacter character for mapping end
	MappingEndCharacter Character = '}'
	// CommentCharacter character for comment
	CommentCharacter Character = '#'
	// AnchorCharacter character for anchor
	AnchorCharacter Character = '&'
	// AliasCharacter character for alias
	AliasCharacter Character = '*'
	// TagCharacter character for tag
	TagCharacter Character = '!'
	// LiteralCharacter character for literal
	LiteralCharacter Character = '|'
	// FoldedCharacter character for folded
	FoldedCharacter Character = '>'
	// SingleQuoteCharacter character for single quote
	SingleQuoteCharacter Character = '\''
	// DoubleQuoteCharacter character for double quote
	DoubleQuoteCharacter Character = '"'
	// DirectiveCharacter character for directive
	DirectiveCharacter Character = '%'
	// SpaceCharacter character for space
	SpaceCharacter Character = ' '
	// LineBreakCharacter character for line break
	LineBreakCharacter Character = '\n'
	TabCharacter       Character = '\t'
	ReturnCharacter    Character = '\r'
)

type TokenType int

const (
	// UnknownType reserve for invalid type
	UnknownType TokenType = iota
	// DocumentHeaderType type for DocumentHeader token
	DocumentStartType
	// DocumentEndType type for DocumentEnd token
	DocumentEndType
	// SequenceEntryType type for SequenceEntry token
	SequenceEntryType
	// MappingKeyType type for MappingKey token
	MappingKeyType
	// MappingValueType type for MappingValue token
	MappingValueType
	// MergeKeyType type for MergeKey token
	MergeKeyType
	// CollectEntryType type for CollectEntry token
	CollectEntryType
	// SequenceStartType type for SequenceStart token
	SequenceStartType
	// SequenceEndType type for SequenceEnd token
	SequenceEndType
	// MappingStartType type for MappingStart token
	MappingStartType
	// MappingEndType type for MappingEnd token
	MappingEndType
	// CommentType type for Comment token
	CommentType
	// ScalarType 字面量
	ScalarType
)

func (t TokenType) String() string {
	switch t {
	case DocumentStartType:
		return "DocumentStartType"
	case DocumentEndType:
		return "DocumentEndType"
	case MappingKeyType:
		return "MappingKeyType"
	case MappingValueType:
		return "MappingValueType"
	case CommentType:
		return "CommentType"
	case ScalarType:
		return "ScalarType"

	}
	return ""

}

// Position type for position in YAML document
type Position struct {
	Line      int // 行号
	OffSet    int // 当前行的偏移量
	IndentNum int // 缩进，即前面的空白字符数
}

// Token type for token
type Token struct {
	Type     TokenType
	Value    string // 字面量
	Position *Position
}

// MappingKey create token for MappingKey
func MappingKey(value string, pos *Position) *Token {
	return &Token{
		Type:     MappingKeyType,
		Value:    value,
		Position: pos,
	}
}

// MappingValue create token for MappingValue
func MappingValue(pos *Position) *Token {
	return &Token{
		Type:     MappingValueType,
		Value:    string(MappingValueCharacter),
		Position: pos,
	}
}

// ScalarValue create token for MappingValue
func ScalarValue(source string, pos *Position) *Token {
	return &Token{
		Type:     ScalarType,
		Value:    source,
		Position: pos,
	}
}

func SequenceEntry(pos *Position) *Token {
	return &Token{
		Type:     SequenceEntryType,
		Value:    "- ",
		Position: pos,
	}
}

func DocumentStart(pos *Position) *Token {
	return &Token{
		Type:     DocumentStartType,
		Value:    "---",
		Position: pos,
	}
}

func DocumentEnd(pos *Position) *Token {
	return &Token{
		Type:     DocumentEndType,
		Value:    "",
		Position: pos,
	}
}

func Comment(value string, pos *Position) *Token {
	return &Token{
		Type:     CommentType,
		Value:    value,
		Position: pos,
	}

}

func Unknown(pos *Position) *Token {
	return &Token{
		Type:     UnknownType,
		Value:    "",
		Position: pos,
	}
}
