package main //このコードはmainてパッケージだぜ的な

import (
	"fmt"
	"math"
	"math/rand" //importしたパッケージは使わないとエラー吐くぽい
) //複数の要素の列挙をカッコでグループ化するのをfactoredとかいうらしい

//shift演算子が直感的に理解できてない、やばいね

const Pi = 3.14 //定数の設定はこう　defineではないよ

func add(x float64, y float64) float64 { //足し算するやつ
	return x + y
}

func swap(x, y string) (string, string) { //戻り値は複数返せる
	return y, x
}

func split(sum int) (x, y int) {
	y = sum % 10
	x = (sum - y) / 10
	return
}

func main() {
	fmt.Println("rand num is", rand.Intn((10)))
	fmt.Println(math.Pi) //ここはpythonと書き方おんなじ！
	fmt.Println(Pi)
	fmt.Println(add(42.1, 13.5)) //関数の書き方は変な感じ　型が後ろで戻り値も後ろ
	fmt.Println(multiply(4, 8))  //関数の宣言は前でも後ろでもいい。ここは気持ち悪いな。。

	var num01 int = 25 //変数の宣言はvar
	var num02 = 49     //初期化子があれば型は省略可能
	//初期化しなかったときはゼロ値が与えられるぞ、boolがfalse,数値がnullではなく0,文字列は空
	var (
		test1 bool = true
		test2 int  = 24
		test3      = 4.5
	) //importと同じでfactoredな宣言が可能
	a, b := swap("hello", "world") //:=で暗黙的な型宣言　var+初期化子といっしょ　関数内でしか使えない

	fmt.Println(a, b)

	fmt.Println(split(num01))
	fmt.Println(split(num02))
	fmt.Printf("Type: %T\n", test1) //型は%Tで表示　type()のpythonとは違うね
	fmt.Printf("Type: %T\n", test2)
	fmt.Printf("Type: %T\n", test3)
}

func multiply(x, y int8) int8 { //掛け算するやつ 型の宣言はまとめられる
	return x * y
}
