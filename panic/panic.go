package main

import (
	"errors"
	"fmt"
)

func f() (rerr error) {
	//deferでパニックをキャッチ
	defer func() {
		// recover()を使ってパニックから回復
		if r := recover(); r != nil {
			if err, isErr := r.(error); isErr {
				rerr = err
			} else {
				// error型でない場合は、fmt.Errorfでエラーメッセージを作成
				rerr = fmt.Errorf("%v", r)
			}
		}
	}()
	panic(errors.New("error"))
	//panicを使った後のコードに到達することはないため、return nilは削除するか修正する必要がある。
	return nil
}

func main() {
	// f()を呼び出して、エラーが返るか確認
	if err := f(); err != nil {
		fmt.Println("パニックから回復:", err)
	} else {
		fmt.Println("エラーは発生しませんでした。")
	}
}