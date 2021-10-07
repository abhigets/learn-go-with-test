package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

const finalWord = "Go!"
const countDown = 3

func Countdown(writer io.Writer, sleeper *Sleeper) {
	for i := countDown; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(writer, i)
	}

	time.Sleep(1 * time.Second)
	fmt.Fprint(writer, finalWord)
}

func main() {
	fmt.Fprintf(os.Stdout, "3")
}
