package lexer

import "monkey/token"

// Goの変数はゼロ値という初期化がされる
type Lexer struct {
	input string // 入力文字列
	position int // 入力文字列における現在の位置（現在の文字を指し示す）
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
	l.position = l.readPosition // 現在の位置を更新
	l.readPosition += 1 // 読み込み位置を更新
}

// 1文字を読む（peek：覗き見），更新はしない
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) { // 終端に達した場合
		return 0 // NUL文字
	} else {
		return l.input[l.readPosition] // 現在の文字のUnicode番号（10進数）を取得
	}
}

// 現在検査中の文字を認識して，トークンを生成
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace() // 空白文字をスキップ

	switch l.ch {
		case '=':
			if l.peekChar() == '=' { // 次の文字が'='の場合
				ch := l.ch // 現在の文字を保存
				l.readChar() // 次の文字を読み込み，現在の位置と次の読み込み位置を更新
				literal := string(ch) + string(l.ch) // 現在の文字と次の文字を連結
				tok = token.Token{Type: token.EQ, Literal: literal} // トークンを生成，literalがstring型なので，newToken()は使えないため，直接Token構造体を生成
			} else {
			tok = newToken(token.ASSIGN, l.ch)
			}
		case '+':
			tok = newToken(token.PLUS, l.ch)
		case '-':
			tok = newToken(token.MINUS, l.ch)
		case '!':
			if l.peekChar() == '=' { // 次の文字が'='の場合
				ch := l.ch // 現在の文字を保存
				l.readChar() // 次の文字を読み込み，現在の位置と次の読み込み位置を更新
				literal := string(ch) + string(l.ch) // 現在の文字と次の文字を連結
				tok = token.Token{Type: token.NOT_EQ, Literal: literal} // トークンを生成，literalがstring型なので，newToken()は使えないため，直接Token構造体を生成
			} else {
			tok = newToken(token.BANG, l.ch)
			}
		case '/':
			tok = newToken(token.SLASH, l.ch)
		case '*':
			tok = newToken(token.ASTERISK, l.ch)
		case '<':
			tok = newToken(token.LT, l.ch)
		case '>':
			tok = newToken(token.GT, l.ch)
		case ';':
			tok = newToken(token.SEMICOLON, l.ch)
		case '(':
			tok = newToken(token.LPAREN, l.ch)
		case ')':
			tok = newToken(token.RPAREN, l.ch)
		case ',':
			tok = newToken(token.COMMA, l.ch)
		case '{':
			tok = newToken(token.LBRACE, l.ch)
		case '}':
			tok = newToken(token.RBRACE, l.ch)
		case 0:
			tok.Literal = ""
			tok.Type = token.EOF
		default:
			if isLetter(l.ch) { // アルファベットの場合
				tok.Literal = l.readIdentifier() // 識別子を読み込み．内部でreadChar()が呼ばれるため，現在の位置と次の読み込み位置が更新されているのでreturnで早期に抜ける．
				tok.Type = token.LookupIdent(tok.Literal) // 識別子の種類を判定
				return tok
			} else if isDigit(l.ch) { // 数字の場合
				tok.Type = token.INT
				tok.Literal = l.readNumber() // 数字を読み込み．内部でreadChar()が呼ばれるため，現在の位置と次の読み込み位置が更新されているのでreturnで早期に抜ける．
				return tok
			} else {
				tok = newToken(token.ILLEGAL, l.ch) // 未知のトークン
			}
	}

	l.readChar() // 次の文字を読み込み，現在の位置と次の読み込み位置を更新
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)} // トークンを生成
}

// 識別子の読み込み：アルファベットの間は読み込み位置を進める
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) { // アルファベットの間は読み込み位置を進める
		l.readChar()
	}
	return l.input[position:l.position] // 読み込み位置を進めた範囲が識別子
}

// 数字の読み込み：数字の間は読み込み位置を進める
func (l *Lexer) readNumber() string { // 数字の読み込み
	position := l.position
	for isDigit(l.ch) { // 数字の間は読み込み位置を進める
		l.readChar()
	}
	return l.input[position:l.position] // 読み込み位置を進めた範囲が数字
}

// アルファベットかどうかの判定
func isLetter(ch byte) bool { // アルファベットかどうかの判定
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// 数字かどうかの判定
func isDigit(ch byte) bool { // 数字かどうかの判定
	return '0' <= ch && ch <= '9'
}

// 空白文字の読み飛ばし
func (l *Lexer) skipWhitespace() { // 空白文字をスキップ
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' { // 空白文字の間は読み込み位置を進める
		l.readChar()
	}
}