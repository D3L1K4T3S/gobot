package lexer

import "gobot/interpretator/token"

type Lexer struct {
	input      string
	currentPos int
	nextPos    int
	char       byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) skipWhiteSpace() {
	for l.char == ' ' ||
		l.char == '\t' ||
		l.char == '\n' ||
		l.char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readChar() {
	if l.nextPos >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.nextPos]
	}
	l.currentPos = l.nextPos
	l.nextPos += 1
}

func (l *Lexer) peekChar() byte {
	if l.nextPos >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPos]
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.currentPos
	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[position:l.currentPos]
}

func (l *Lexer) readNumber() string {
	position := l.currentPos
	for isDigit(l.char) {
		l.readChar()
	}
	return l.input[position:l.currentPos]
}

func (l *Lexer) readString() string {
	position := l.currentPos + 1
	for {
		l.readChar()
		if l.char == '"' || l.char == 0 {
			break
		}
	}
	return l.input[position:l.currentPos]
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func createToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.skipWhiteSpace()

	switch l.char {
	case '=':
		if l.peekChar() == '=' {
			char := l.char
			l.readChar()
			literal := string(char) + string(l.char)
			t = token.Token{Type: token.EQ, Literal: literal}
		} else {
			t = createToken(token.ASSIGN, l.char)
		}
	case '+':
		t = createToken(token.PLUS, l.char)
	case '-':
		t = createToken(token.MINUS, l.char)
	case '!':
		if l.peekChar() == '=' {
			char := l.char
			l.readChar()
			literal := string(char) + string(l.char)
			t = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			t = createToken(token.DENIAL, l.char)
		}
	case '/':
		t = createToken(token.SLASH, l.char)
	case '*':
		t = createToken(token.ASTERISK, l.char)
	case '<':
		t = createToken(token.LT, l.char)
	case '>':
		t = createToken(token.GT, l.char)
	case ';':
		t = createToken(token.SEMICOLON, l.char)
	case ':':
		t = createToken(token.COLON, l.char)
	case ',':
		t = createToken(token.COMMA, l.char)
	case '{':
		t = createToken(token.LBRACE, l.char)
	case '}':
		t = createToken(token.RBRACE, l.char)
	case '(':
		t = createToken(token.LPAREN, l.char)
	case ')':
		t = createToken(token.RPAREN, l.char)
	case '"':
		t.Type = token.STRING
		t.Literal = l.readString()
	case '[':
		t = createToken(token.LBRACKET, l.char)
	case ']':
		t = createToken(token.RBRACKET, l.char)
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isLetter(l.char) {
			t.Literal = l.readIdentifier()
			t.Type = token.GetKeyWord(t.Literal)
			return t
		} else if isDigit(l.char) {
			t.Type = token.INT
			t.Literal = l.readNumber()
			return t
		} else {
			t = createToken(token.ILLEGAL, l.char)
		}
	}

	l.readChar()
	return t
}
