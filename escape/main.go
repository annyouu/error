package main

import "fmt"

// escape 型を定義
type escape struct {}

func f() {
	g()  // g()を呼び出す
}

func g() {
	panic(escape{})  // escape型のパニックを発生させる。
}

func main() {
	// deferでパニックで回復
	defer func() {
		// panicが発生した場合にその内容を回収する
		if r := recover(); r != nil {
			if _, ok := r.(escape); ok {
				fmt.Println("Escaped") // escape型のパニックが発生した場合
			} else {
				panic(r)  // 他のパニックが発生した場合は再度panicを発生させる。
			}
		}
	}()
	f() // f()を呼び出す。
}