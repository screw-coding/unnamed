package lexer

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/screw-coding/yaml/token"
	"io"
	"strings"
)

// Tokenizer 词法解析器
type Tokenizer struct {
	reader      *bufio.Reader
	currentLine []rune // 使用 run 类型 而不是 byte 类型，因为希望兼容 unicode
	line        int    // 所读取的文档行数

	position     int  // 所输入字符串中的当前位置（指向当前字符）
	readPosition int  // 所输入字符串中的当前读取位置（指向当前字符之后的一个字符）
	ch           rune // 当前正在查看的字符

	tokens []*token.Token // 保存所有的 token
}

func New(reader io.Reader) *Tokenizer {
	t := &Tokenizer{reader: bufio.NewReader(reader)}
	err := t.readLine()
	if err != nil {
		panic("error read")
	}
	return t
}

// 读取一行数据，并且去掉结束的换行符
func (t *Tokenizer) readLine() error {
	for {
		line, err := t.reader.ReadBytes('\n')
		if err != nil {
			return err
		}
		line = bytes.TrimRight(line, "\r\n")
		// line + 1
		t.line += 1
		t.currentLine = []rune(string(line))
		if !t.isBlankLine() {
			break
		}
	}

	// 重置 position
	t.position = 0
	t.readPosition = 0
	t.readRune()
	return nil
}

// 读取下一个字符
func (t *Tokenizer) readRune() {
	if t.readPosition >= len(t.currentLine) {
		t.ch = 0 // NUL字符的ASCII编码，用来表示“尚未读取任何内容”或“文件结尾”
	} else {
		t.ch = t.currentLine[t.readPosition]
	}
	t.position = t.readPosition
	t.readPosition += 1
}

// 查看下一个字符
func (t *Tokenizer) peekRune() rune {
	if t.readPosition >= len(t.currentLine) {
		return 0
	} else {
		return t.currentLine[t.readPosition]
	}
}

func (t *Tokenizer) getPosition() *token.Position {
	// 前面的空白字符数量
	pos := 0
	for t.currentLine[pos] == ' ' {
		pos++
	}

	return &token.Position{
		Line:      t.line,
		OffSet:    t.position,
		IndentNum: pos,
	}
}

func (t *Tokenizer) nextToken() (*token.Token, error) {
	var tok *token.Token

	// 什么时候换行
	if t.ch == 0 {
		err := t.readLine()
		if err != nil {
			if err == io.EOF {
				// 结尾用给了
				tok = token.DocumentEnd(nil)
				t.appendToken(tok)
				return tok, err
			} else {
				return nil, err
			}
		}

	}

	currentLineString := string(t.currentLine)
	fmt.Println(string(t.currentLine))

	// 读取到了 -- , 判断是文档开头还是结束
	// TODO 如果读到了文档最后，也需要标记 documentEnd
	if strings.HasPrefix(string(t.currentLine), "---") {
		// 如果是 第一行
		if t.line == 1 {
			tok = token.DocumentStart(t.getPosition())
			t.skipRunes(3)
		} else {
			if t.tokens[len(t.tokens)-1].Type == token.DocumentEndType {
				tok = token.DocumentStart(t.getPosition())
				t.skipRunes(3)
			} else {
				tok = token.DocumentEnd(t.getPosition())
			}
		}
		t.appendToken(tok)
		return tok, nil
	} else {
		// 如果第一行不是 --- 并且还没有任何token
		if len(t.tokens) == 0 {
			tok = token.DocumentStart(t.getPosition())
			t.appendToken(tok)
			return tok, nil
		}
	}
	// 如果 # 开头
	switch t.ch {
	case '#':
		// 读取 # 一直到 换行符之间的所有内容
		pos := t.getPosition()
		tok = token.Comment(t.readUntilLineBreak(), pos)
	case ':':
		pos := t.getPosition()
		tok = token.MappingValue(pos)
		t.readRune()
	default:
		// 如果当前行有 : 那前面的一定是一个 key
		if strings.ContainsRune(currentLineString, ':') && !strings.Contains(currentLineString, "- ") {
			// 当前指针是否在开头
			if t.position == 0 {
				t.skipBlankUntilLetter()
				// 然后读取出 : 之前的 key
				position := t.getPosition()
				mappingKey := t.readUntilMappingValueCharacter()
				tok = token.MappingKey(mappingKey, position)
				// 如果不在开头，那么前面是不是 :
			} else {
				// 读取到第一个非空格的字符
				t.skipBlankUntilNoneBlank()
				position := t.getPosition()
				tok = token.ScalarValue(t.readUntilBlank(), position)
				t.skipBlankUntilNoneBlank()
			}
		}
		// 如果当前行有 -, 那么说明是一个 sequence
		if strings.Contains(currentLineString, "- ") && !strings.ContainsRune(currentLineString, ':') {
			// 当前指针是否在开头
			if t.position == 0 {
				t.skipBlankUntilDash()
				position := t.getPosition()
				tok = token.SequenceEntry(position)
				t.skipRunes(2)
			} else {
				// 读取到第一个非空格的字符
				t.skipBlankUntilNoneBlank()
				position := t.getPosition()
				tok = token.ScalarValue(t.readUntilBlank(), position)
				t.skipBlankUntilNoneBlank()
			}
		}

		// 都有
		if strings.Contains(currentLineString, "- ") && strings.ContainsRune(currentLineString, ':') {
			if t.position == 0 {
				t.skipBlankUntilDash()
				position := t.getPosition()
				tok = token.SequenceEntry(position)
				t.skipRunes(2)
			} else if t.tokens[len(t.tokens)-1].Type == token.SequenceEntryType {
				t.skipBlankUntilLetter()
				// 然后读取出 : 之前的 key
				position := t.getPosition()
				mappingKey := t.readUntilMappingValueCharacter()
				tok = token.MappingKey(mappingKey, position)
			} else {
				// 读取到第一个非空格的字符
				t.skipBlankUntilNoneBlank()
				position := t.getPosition()
				tok = token.ScalarValue(t.readUntilBlank(), position)
				t.skipBlankUntilNoneBlank()
			}
		}

	}

	if tok == nil {
		tok = token.Unknown(t.getPosition())
	}

	t.appendToken(tok)
	return tok, nil
}

// 跳过空行
func (t *Tokenizer) skipWhiteLine() {
	for t.ch == rune(token.SpaceCharacter) ||
		t.ch == rune(token.TabCharacter) ||
		t.ch == rune(token.LineBreakCharacter) ||
		t.ch == rune(token.ReturnCharacter) {
		t.readRune()
	}
}

// 跳过读取几个字符
func (t *Tokenizer) skipRunes(num int) {
	for i := 0; i < num; i++ {
		t.readRune()
	}
}

func (t *Tokenizer) appendToken(tok *token.Token) {
	t.tokens = append(t.tokens, tok)
}

func (t *Tokenizer) isBlankLine() bool {
	for _, r := range t.currentLine {
		if r != rune(token.SpaceCharacter) &&
			r != rune(token.TabCharacter) &&
			r != rune(token.LineBreakCharacter) &&
			r != rune(token.ReturnCharacter) {
			return false
		}
	}
	return true
}

// 从当前位置读取，一直到换行为止
func (t *Tokenizer) readUntilLineBreak() string {
	position := t.position
	for t.ch != 0 {
		t.readRune()
	}
	return string(t.currentLine[position:t.position])
}

func (t *Tokenizer) readUntilMappingValueCharacter() string {
	position := t.position
	for t.ch != rune(token.MappingValueCharacter) {
		t.readRune()
	}
	return strings.TrimSpace(string(t.currentLine[position:t.position]))
}

func (t *Tokenizer) readUntilBlank() string {
	position := t.position
	for t.ch != rune(token.SpaceCharacter) {
		if t.ch == 0 {
			break
		}
		t.readRune()
	}
	return strings.TrimSpace(string(t.currentLine[position:t.position]))
}

// 跳过空格字符移植到读取到当前字符为 字母为止
func (t *Tokenizer) skipBlankUntilLetter() {
	for isLetter(t.ch) && t.ch == ' ' {
		t.readRune()
	}
}

func (t *Tokenizer) skipBlankUntilDash() {
	for t.ch != '-' {
		t.readRune()
	}
}

// 跳过空格字符移植到读取到当前字符为 字母为止
func (t *Tokenizer) skipBlankUntilNoneBlank() {
	for t.ch == ' ' {
		if t.ch == 0 {
			break
		}
		t.readRune()
	}
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}
