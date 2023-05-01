package lexer

import "monkey/token"

// Goの変数はゼロ値という初期化がされる
type Lexer struct {
	input string // 入力文字列
	poition int // 入力文字列における現在の位置（現在の文字を指し示す）
	readPosition int // これから読み込む位置（現在の文字の次）
	ch byte // 現在検査中の文字のUnicode番号（10進数))
}

func New(input string) *Lexer {
	l := &Lexer{input: input} // 入力文字列をもとにlexerを生成
	l.readChar() // 先頭の文字読み込み，現在の位置と次の読み込み位置を更新
	return l
}

// 1文字読み込み，文字列の位置を更新
// レシーバ付き関数（メソッド）
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) { // 終端に達した場合
		l.ch = 0 // NUL文字
	} else {
		l.ch = l.input[l.readPosition] // 現在の文字のUnicode番号（10進数）を取得
	}
	l.poition = l.readPosition // 現在の位置を更新
	l.readPosition += 1 // 読み込み位置を更新
}

// 現在検査中の文字がデリミタかどうかを判定
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
		case '=':
			tok = newToken(token.ASSIGN, l.ch)
		case ';':
			tok = newToken(token.SEMICOLON, l.ch)
		case '(':
			tok = newToken(token.LPAREN, l.ch)
		case ')':
			tok = newToken(token.RPAREN, l.ch)
		case ',':
			tok = newToken(token.COMMA, l.ch)
		case '+':
			tok = newToken(token.PLUS, l.ch)
		case '{':
			tok = newToken(token.LBRACE, l.ch)
		case '}':
			tok = newToken(token.RBRACE, l.ch)
		case 0:
			tok.Literal = ""
			tok.Type = token.EOF
	}

	l.readChar() // 次の文字を読み込み，現在の位置と次の読み込み位置を更新
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)} // トークンを生成
}