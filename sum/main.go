package main

import (
	"errors"
	"fmt"
)

type CustomError struct  {
	Message string
	Err error
}

// CustomErrorに対しエラーを実装
func (e *CustomError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap()メソッドを実装して、ラップされたエラーを返す
func (e *CustomError) Unwrap() error {
	return e.Err
}

// Walk関数: 再帰的にラップされたエラーを処理する
func Walk(err error, f func(err error)) {
	if err == nil {
		return
	}
	//Unwrap()メソッドを使って再帰的にエラーを展開
	for err != nil {
		f(err)
		if unwrapped, ok := err.(interface{ Unwrap() error }); ok {
			err = unwrapped.Unwrap()
		} else {
			break
		}
	}

	// switch err := err.(type) {
	// case interface{ Unwrap() error }:
	// 	// Unwrap()メソッドを使って再帰的にエラーを展開
	// 	Walk(err.Unwrap(), f)
	// default:
	// 	//最終的に到達したエラーを処理
	// 	f(err)
	// }
}

func main() {
	// エラーのラッピング
	err1 := errors.New("the base error")
	err2 := &CustomError{
		Message: "custom error 1",
		Err: err1,
	}
	err3 := &CustomError {
		Message: "custom error 2",
		Err: err2,
	}

	// Walk関数を使ってエラーを再帰的に処理
	Walk(err3, func(err error){
		//エラーがあるときにそのメッセージを表示
		fmt.Println("Error:", err)
	})

	
}