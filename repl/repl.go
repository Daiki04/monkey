package repl

import (
	"bufio"
	"fmt"
	"io"

	"monkey/lexer"
	"monkey/token"
)

// REPL（Read-Eval-Print-Loop）

// 1. ユーザーが入力した文字列を受け取る
// 2. 入力文字列を字句解析し，トークンを生成
// 3. 生成したトークンを表示
const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in) // 標準入力からの読み込みを行うScannerを生成

	for {
		// 1. ユーザーが入力した文字列を受け取る
		fmt.Printf(PROMPT) // プロンプトを表示
		scanned := scanner.Scan() // 標準入力から1行読み込み，改行を除いた文字列を返す
		if !scanned { // EOF（End Of File）の場合
			return
		}

		line := scanner.Text() // 読み込んだ文字列を取得
		l := lexer.New(line) // 入力文字列をもとにlexerを生成

		// 2. 入力文字列を字句解析し，トークンを生成
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() { // 入力文字列の終端に達するまでトークンを生成
			// 3. 生成したトークンを表示
			fmt.Printf("%+v\n", tok) // トークンを表示
		}
	}
}