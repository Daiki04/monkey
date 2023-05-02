package main

import (
	"fmt"
	"os"
	"os/user"

	"monkey/repl"
)

func main() {
	user, err := user.Current() // 現在のユーザー情報を取得
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username) // ユーザー名を表示
	fmt.Printf("Feel free to type in commands\n") // コマンドを入力してください

	repl.Start(os.Stdin, os.Stdout) // REPLを開始
} 