// 正規表現

package main

import (
	"fmt"
	"regexp"
)

// プログラム起動時に正規表現をコンパイル(エラーなら即パニック)
var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

// パニックを起こす例
// var invalidRegex = regexp.MustCompile(`^([a-z]+)`)

func main() {
	//事前にコンパイルされたvalidIDを使う
	fmt.Println(validID.MatchString("adam[23]")) // true

	//関数内でコンパイルする場合は、エラー処理を行う
	validID2, err := regexp.Compile(`^[a-z]+\[[0-9]+\]$`)
	if err != nil {
		fmt.Println("正規表現のコンパイルに失敗:", err)
		return
	}
	fmt.Println(validID2.MatchString("adam[23]")) //true
}