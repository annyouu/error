package main

import (
	"html/template"
	"os"
)

// テンプレートは、変数の値の埋め込みができる。 Parse("<h1>{{.Title}}</h1>")で、Parseの中身の部分をテンプレートしてつかえるように置き換える


var tmpl = template.Must(template.New("example").Parse("<h1>{{.Title}}</h1>"))


//③os.Stdoutに出力ってどういうこと？
// os.Stdoutは、標準出力を表す。つまり、tmpl.Execute(os.Stdout, data)を実行すると、ターミナルに結果が表示される。

func main() {
	data := struct{ Title string }{"Hello, World!"}
	tmpl.Execute(os.Stdout, data)
}





