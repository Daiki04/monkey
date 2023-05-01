package token

// トークンの種類を表す型
type TokenType string

// トークンの構造体
type Token struct {
	Type TokenType
	Literal string
}

// トークンタイプの種類
// Goでは定数の型指定は不要．定数が使用されるときに適切な型の定数として扱ってくれる．
const (
	ILLEGAL = "ILLEGAL" // トークンや文字が未知であることを表す
	EOF = "EOF" // ファイル終端（停止してよいことを示す）

	// 識別子 + リテラル（区別不要）
	IDENT = "IDENT" // 識別子（変数名）：add, foobar, x, y, ...
	INT = "INT" // 整数リテラル：1234567890

	// 演算子
	ASSIGN = "=" // 代入演算子
	PLUS = "+" // 加算演算子

	// デリミタ：データを区切る区切り文字
	COMMA = "," // カンマ
	SEMICOLON = ";" // セミコロン

	LPAREN = "(" // 左括弧
	RPAREN = ")" // 右括弧
	LBRACE = "{" // 左波括弧
	RBRACE = "}" // 右波括弧

	// キーワード：特別な意味を持つ識別子
	FUNCTION = "FUNCTION" // 関数
	LET = "LET" // 変数宣言
)