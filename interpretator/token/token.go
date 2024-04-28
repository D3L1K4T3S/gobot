package token

type TokenType string

const (
	// Операторы
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	DENIAL   = "!"
	ASSIGN   = "="

	// Операторы сравнения
	EQ     = "=="
	NOT_EQ = "!="
	LT     = "<"
	GT     = ">"

	// Разделители
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Служебные
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	// Литералы
	INT    = "INT"
	STRING = "STRING"
	IDENT  = "IDENT"

	// Таблица символов
	IF       = "IF"
	ELSE     = "ELSE"
	LET      = "LET"
	RETURN   = "RETURN"
	WHILE    = "WHILE"
	FUNCTION = "FUNCTION"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"if":     IF,
	"else":   ELSE,
	"while":  WHILE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
	"func":   FUNCTION,
	"let":    LET,
}

func GetKeyWord(id string) TokenType {
	if token, ok := keywords[id]; ok {
		return token
	}
	return IDENT
}
