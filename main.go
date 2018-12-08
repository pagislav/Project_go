package main

import "fmt"

//import "time"

var gameOver bool

var ycor int = 15
var xcor int = 15
var arr [30][30]string

// tail_y := make([]int, 1)
// tail_x := make([]int, 1)
// tail_cor_y := make([]int, 1)
// tail_cor_x := make([]int, 1)
// var nap string

func main() {
	gameOver = false
	//	mapp()
	arr[ycor][xcor] = "O"
	fmt.Println(arr[ycor][xcor])
	fmt.Println("chose level of complexity (10 to 1): ")
	var hard int
	fmt.Scanf("%f", &hard)
	hard = hard * 10
	for !gameOver {
		// random_fruit()
		// vivod()
		// shake_move()
		// hvost_move()
		//time.Sleep(time.Millisecond * hard) // тайм.слип не принимает хард,но принимает константы
	}
}
