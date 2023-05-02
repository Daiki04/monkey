package ast

import "monkey/token"

// ASTのノード（文、式）のインターフェース
type Node interface {
	TokenLiteral() string // トークンのリテラル値を返す
}

// 文のためのインターフェース
type Statement interface {
	Node
	statementNode() // ダミーメソッド
}

// 式のためのインターフェース
type Expression interface {
	Node
	expressionNode() // ダミーメソッド
}

// すべてのASTのルートノード
type Program struct {
	Statements []Statement // 文の配列
}

// トークンのリテラル値を返す
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 { // 文が存在する場合
		return p.Statements[0].TokenLiteral() // 最初の文のトークンのリテラル値を返す
	} else {
		return ""
	}
}

// let文：Statementインターフェースを満たす
type LetStatement struct {
	Token token.Token // token.LETトークン
	Name  *Identifier // 識別子
	Value Expression  // 式
}

// インターフェースを満たすためにインターフェースと同じ名前のメソッドを定義
// インターフェイスの中にある関数と同名の関数を定義することで，自動的にインターフェースを実装することができる
func (ls *LetStatement) statementNode() {} // ダミーメソッド
func (ls *LetStatement) TokenLiteral() string {return ls.Token.Literal}

// 識別子：Expressionインターフェースを満たす
// x = 5; y = x; というような場合に，前者のxは文であるが，後者の識別子xは値を生成する式(Expression)であるためExpressionインターフェースで実装する
type Identifier struct {
	Token token.Token // token.IDENTトークン
	Value string      // 識別子の名前
}

// インターフェースを満たすためにインターフェースと同じ名前のメソッドを定義
// インターフェースの中にある関数と同名の関数を定義することで，自動的にインターフェースを実装することができる
func (i *Identifier) expressionNode() {} // ダミーメソッド
func (i *Identifier) TokenLiteral() string {return i.Token.Literal}