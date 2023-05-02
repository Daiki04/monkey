package lexer

import (
	"testing" // Goのテストフレームワーク

	"monkey/token" 
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	let add = fn(x, y) {
		x + y;
	};

	let result = add(five, ten);
	!-/*5;
	5 < 10 > 5;

	if (5 < 10) {
		return true;
	} else {
		return false;
	}

	10 == 10;
	10 != 9;
	`

	// テストケース（構造体をもったスライス）
	tests := []struct {
		expectedType token.TokenType // 期待されるトークンの種類
		expectedLiteral string // 期待されるトークンのリテラル
	}{
		{token.LET, "let"}, // キーワード: 変数宣言
		{token.IDENT, "five"}, // 識別子: 変数名
		{token.ASSIGN, "="}, // 代入演算子
		{token.INT, "5"}, // 整数リテラル
		{token.SEMICOLON, ";"}, // セミコロン
		{token.LET, "let"}, // キーワード: 変数宣言
		{token.IDENT, "ten"}, // 識別子: 変数名
		{token.ASSIGN, "="}, // 代入演算子
		{token.INT, "10"}, // 整数リテラル
		{token.SEMICOLON, ";"}, // セミコロン
		{token.LET, "let"}, // キーワード: 変数宣言
		{token.IDENT, "add"}, // 識別子: 変数名
		{token.ASSIGN, "="}, // 代入演算子
		{token.FUNCTION, "fn"}, // キーワード: 関数
		{token.LPAREN, "("}, // 左括弧
		{token.IDENT, "x"}, // 識別子: 変数名
		{token.COMMA, ","}, // カンマ
		{token.IDENT, "y"}, // 識別子: 変数名
		{token.RPAREN, ")"}, // 右括弧
		{token.LBRACE, "{"}, // 左波括弧
		{token.IDENT, "x"}, // 識別子: 変数名
		{token.PLUS, "+"}, // 加算演算子
		{token.IDENT, "y"}, // 識別子: 変数名
		{token.SEMICOLON, ";"}, // セミコロン
		{token.RBRACE, "}"}, // 右波括弧
		{token.SEMICOLON, ";"}, // セミコロン
		{token.LET, "let"}, // キーワード: 変数宣言
		{token.IDENT, "result"}, // 識別子: 変数名
		{token.ASSIGN, "="}, // 代入演算子
		{token.IDENT, "add"}, // 識別子: 変数名
		{token.LPAREN, "("}, // 左括弧
		{token.IDENT, "five"}, // 識別子: 変数名
		{token.COMMA, ","}, // カンマ
		{token.IDENT, "ten"}, // 識別子: 変数名
		{token.RPAREN, ")"}, // 右括弧
		{token.SEMICOLON, ";"}, // セミコロン
		{token.BANG, "!"}, // 接頭辞演算子
		{token.MINUS, "-"}, // 減算演算子
		{token.SLASH, "/"}, // 除算演算子
		{token.ASTERISK, "*"}, // 乗算演算子
		{token.INT, "5"}, // 整数リテラル
		{token.SEMICOLON, ";"}, // セミコロン
		{token.INT, "5"}, // 整数リテラル
		{token.LT, "<"}, // 比較演算子
		{token.INT, "10"}, // 整数リテラル
		{token.GT, ">"}, // 比較演算子
		{token.INT, "5"}, // 整数リテラル
		{token.SEMICOLON, ";"}, // セミコロン
		{token.IF, "if"}, // キーワード: 条件分岐
		{token.LPAREN, "("}, // 左括弧
		{token.INT, "5"}, // 整数リテラル
		{token.LT, "<"}, // 比較演算子
		{token.INT, "10"}, // 整数リテラル
		{token.RPAREN, ")"}, // 右括弧
		{token.LBRACE, "{"}, // 左波括弧
		{token.RETURN, "return"}, // キーワード: 戻り値
		{token.TRUE, "true"}, // 真偽値リテラル
		{token.SEMICOLON, ";"}, // セミコロン
		{token.RBRACE, "}"}, // 右波括弧
		{token.ELSE, "else"}, // キーワード: 条件分岐
		{token.LBRACE, "{"}, // 左波括弧
		{token.RETURN, "return"}, // キーワード: 戻り値
		{token.FALSE, "false"}, // 真偽値リテラル
		{token.SEMICOLON, ";"}, // セミコロン
		{token.RBRACE, "}"}, // 右波括弧
		{token.INT, "10"}, // 整数リテラル
		{token.EQ, "=="}, // 比較演算子
		{token.INT, "10"}, // 整数リテラル
		{token.SEMICOLON, ";"}, // セミコロン
		{token.INT, "10"}, // 整数リテラル
		{token.NOT_EQ, "!="}, // 比較演算子
		{token.INT, "9"}, // 整数リテラル
		{token.SEMICOLON, ";"}, // セミコロン
		{token.EOF, ""}, // 終端記号
	}

	l := New(input) // 入力文字列をもとにlexerを生成

	for i, tt := range tests { // テストケースを順に実行
		tok := l.NextToken() // lexerからトークンを取得

		// トークンの種類が期待されるものと異なる場合
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", 
				i, tt.expectedType, tok.Type)
		}

		// トークンのリテラルが期待されるものと異なる場合
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", 
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}