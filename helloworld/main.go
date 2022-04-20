package main

import (
	"encoding/json"
	"fmt"
	"helloworld/algorithm"
	"helloworld/web"
	"io"
	"io/ioutil"
	"net/http"
)

type Param[T float64 | float32 | int] struct {
	Arr []T       `json:"arr"`
	A   []float64 `json:"a"`
	B   []float64 `json:"b"`
}

var handler = map[string]func(w http.ResponseWriter, r *http.Request){

	"/tree": func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		if method == "OPTIONS" {
			return
		}
		fmt.Println(method)
		fmt.Println(r.ContentLength)
		bytes, err := ioutil.ReadAll(r.Body)
		fmt.Println("as follow:", string(bytes))

		param := &Param[float64]{
			Arr: nil,
			A:   nil,
			B:   nil,
		}
		_ = json.Unmarshal(bytes, param)

		fmt.Println(param)
		s, res, m := algorithm.Solution(param.A, param.B)
		fmt.Println(res)
		algorithm.PrintMat(s)
		algorithm.PrintMat(m)
		tree := algorithm.Treeify(param.Arr, s, 0, len(param.A)-1)

		marshal, _ := json.Marshal(tree)
		format := string(marshal)
		fmt.Println(format)
		_, err = io.WriteString(w, format)
		if err != nil {
			return
		}

	},
}

// 控制台输入go env -w GOARCH=amd64，修改环境参数，才可以调试
// 修改运行配置中添加环境GOARCH即可 👆
// 配置环境GOOS=linux可以编译linux可执行文件
func main() {
	web.Start(handler)
}
