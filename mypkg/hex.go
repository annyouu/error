package mypkg

import "fmt"

// Hex型のインターフェースを定義
type Hex int

// Stringメソッド (16進数に変換)
func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}