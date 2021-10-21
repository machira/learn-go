//package mocking
package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep(seconds int)
}

type RealSleeper struct {

}

func (r* RealSleeper) Sleep(i int) {
	//oneSecond, _ := time.ParseDuration("1s")
	//time.Sleep(1 * time.Second)
	time.Sleep(time.Duration(i) * time.Second)
}

func Countdown(out io.Writer, s Sleeper) {
	//oneSecond, _ := time.ParseDuration("1s")
	for i := 3; i > 0; i-- {
		//time.Sleep(oneSecond)
		s.Sleep(1)
		_, _ = fmt.Fprintln(out, i)
	}
	//time.Sleep(oneSecond)
	s.Sleep(1)
	_,_= fmt.Fprint(out, "Go!")
}

func main() {
	r := &RealSleeper{}
	Countdown(os.Stdout, r)
}