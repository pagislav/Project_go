package main

// "time"
import (
	"fmt"
	"os"
	"os/exec"
)

const a = 30
const b = 30

var arr [a][b]string

func mapp() {
	for i := 0; i < b; i++ {
		arr[0][i] = "#"
	}
	for i := 1; i < a-1; i++ {
		for j := 0; j < b; j++ {
			if j == 0 || j == b-1 {
				arr[i][j] = "#"
			} else {
				arr[i][j] = " "
			}
		}
	}
	for i := 0; i < b; i++ {
		arr[a-1][i] = "#"
	}
}

func vivod() {
	for i := 0; i < a; i++ {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		for j := 0; j < b; j++ {
			fmt.Printf(arr[i][j])
		}
		fmt.Println(" ")

	}
}

func main() {
	mapp()
	for true {
		vivod()
		//		time.Sleep(50 * time.Millisecond)
	}

}
