package lexer

import (
	"errors"
	"strconv"
	"unicode/utf16"

	"github.com/bamchoh/json_parser/token"
)

type Lexer struct {
	input        []rune
	position     int
	readPosition int
	ch           rune
}

func New(input string) *Lexer {
	l := &Lexer{input: []rune(input)}
	l.readChar()
	return l
}

func newToken(tokenType token.TokenType, ch rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) NextToken() (token.Token, error) {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case ':':
		tok = newToken(token.COLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '-':
		l.readChar()
		if isDigit(l.ch) {
			tok.Type = token.NUMBER
			num := l.readNumber()
			lit := make([]rune, len(num)+1)
			lit[0] = '-'
			for i := 1; i < len(lit); i++ {
				lit[i] = num[i-1]
			}
			tok.Literal = string(lit)
			return tok, nil
		}
		return tok, errors.New("minus character should be followed by a number part")
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	case '"':
		ret, err := l.readString()
		if err != nil {
			return tok, err
		}
		tok.Type = token.STRING
		tok.Literal = ret
	default:
		if isLetter(l.ch) {
			tok.Literal = string(l.readIdentifier())
			tok.Type = token.LookupIdent(tok.Literal)
			return tok, nil
		} else if isDigit(l.ch) {
			tok.Type = token.NUMBER
			tok.Literal = string(l.readNumber())
			return tok, nil
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok, nil
}

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) readNumber() []rune {
	position := l.position
	if isDigit1to9(l.ch) {
		l.readChar()
		for isDigit(l.ch) {
			l.readChar()
		}
		if l.ch == '.' {
			l.readChar()
			for isDigit(l.ch) {
				l.readChar()
			}
		}
		if l.ch == 'e' {
			l.readChar()
			if l.ch == '+' || l.ch == '-' || isDigit(l.ch) {
				l.readChar()
				for isDigit(l.ch) {
					l.readChar()
				}
			}
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() (string, error) {
	var ret []rune
eos:
	for {
		l.readChar()
		switch l.ch {
		case '"':
			break eos
		case 0:
			break eos
		case '\\':
			l.readChar()
			switch l.ch {
			case 'b':
				ret = append(ret, '\b')
				continue
			case 'f':
				ret = append(ret, '\f')
				continue
			case 'r':
				ret = append(ret, '\r')
				continue
			case 'n':
				ret = append(ret, '\n')
				continue
			case 't':
				ret = append(ret, '\t')
				continue
			case 'u':
				var r1, r2 rune
				r1, err := l.readUnicode()
				if err != nil {
					return "", err
				}
				if utf16.IsSurrogate(r1) {
					l.readChar()
					if l.ch == '\\' && l.peekChar() == 'u' {
						l.readChar()
						r2, err = l.readUnicode()
						if err != nil {
							return "", err
						}
						r1 = utf16.DecodeRune(r1, r2)
					} else {
						return "", errors.New("Unicode surrogate piar is wrong")
					}
				}
				ret = append(ret, r1)
				continue
			}
			ret = append(ret, l.ch)
			continue
		default:
			ret = append(ret, l.ch)
		}
	}
	return string(ret), nil
}

func (l *Lexer) readUnicode() (ret rune, err error) {
	pos := l.position + 1
	l.readChar()
	l.readChar()
	l.readChar()
	l.readChar()
	num, err := strconv.ParseInt(string(l.input[pos:pos+4]), 16, 32)
	if err != nil {
		return 0, err
	}
	return rune(num), nil
}

func (l *Lexer) skipWhiteSpace() {
	for isWhiteSpace(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() []rune {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch rune) bool {
	return !isDigit(ch) && !isWhiteSpace(ch) && !isSymbol(ch)
}

func isSymbol(ch rune) bool {
	return (ch == 0x5B ||
		ch == 0x7B ||
		ch == 0x5D ||
		ch == 0x7D ||
		ch == 0x3A ||
		ch == 0x2C)
}

func isWhiteSpace(ch rune) bool {
	return ch == 0x20 || ch == 0x09 || ch == 0x0A || ch == 0x0D
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isDigit1to9(ch rune) bool {
	return '1' <= ch && ch <= '9'
}
