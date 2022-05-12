package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type data struct {
	sync.Mutex
	text   string
	buffer *[]byte
}

func HASH(inc chan *data) {
	for elme := range inc {
		if elme == nil {
			return
		}
		fmt.Println("s")
		elme.Unlock()
	}
}

func Write(outc chan *data) {
	for elme := range outc {
		if elme == nil {
			return
		}
		elme.Lock()
		fmt.Println("s")
		elme.Unlock()
	}
}

func main() {
	var wait sync.WaitGroup
	var inc chan *data
	var outc chan *data

	wait.Add(1)
	go func() {
		defer wait.Done()
		HASH(inc)
	}()

	wait.Add(1)
	go func() {
		defer wait.Done()
		Write(outc)
	}()

	f, _ := os.Open("text.txt")
	defer f.Close()

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		line := sc.Text()
		data := data{}
		data.text = line
		data.Lock()
		inc <- &data
		outc <- &data
	}

	wait.Wait()

}
