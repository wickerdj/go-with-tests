package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// Countdown ...
func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintf(out, strconv.Itoa(i)+"\n")
	}
	fmt.Fprintf(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
