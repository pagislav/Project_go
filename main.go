package main

// 0.1 "бублик", нет хвоста,нет еды,нет kbhit() и  getch(), хоть что-то
import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

const a = 30

var dir string
var xcor int = 15
var ycor int = 15

func vivod() {
	//cmd := exec.Command("cmd", "/c", "cls") // for windows
	cmd := exec.Command("clear") // for Linux and MacOS
	cmd.Stdout = os.Stdout
	cmd.Run()

	for i := 0; i < a; i++ {
		fmt.Print("#")
	}
	fmt.Println(" ")
	for i := 1; i < a-1; i++ {
		fmt.Print("#")
		for j := 1; j < a-1; j++ {
			if i == ycor && j == xcor {
				fmt.Print("0")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println("#")
	}
	for i := 0; i < a; i++ {
		fmt.Print("#")
	}
	time.Sleep(time.Millisecond * 10 * 10)
}

func vvod() {
	r := bufio.NewReader(os.Stdin)
	c, err := r.ReadByte()
	if err != nil {
		panic(err)
	}
	dir = string(c)
	switch dir {
	case "w":
		ycor -= 1
	case "s":
		ycor += 1
	case "a":
		xcor -= 1
	case "d":
		xcor += 1
	}
	if xcor == 29 {
		xcor = 1
	} else if xcor == 0 {
		xcor = 28
	}
	if ycor == 29 {
		ycor = 1
	} else if ycor == 0 {
		ycor = 28
	}
}
func move() {

}

func main() {
	gameOver := false

	for !gameOver {
		vvod()
		vivod()
		move()
	}

}
