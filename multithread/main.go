package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"runtime"
	"sync"
)

type data struct {
	count int
	sync.Mutex
	text   string
	buffer []byte
}

func HASH(inc chan *data, w *sync.WaitGroup) {
	defer w.Done()

	for elme := range inc {
		checksum := sha256.Sum256([]byte(elme.text))
		elme.buffer = checksum[:]

		fmt.Println(elme.count)

		elme.Unlock()
	}
}

func Write(outc chan *data, w *sync.WaitGroup) {
	defer w.Done()

	for elme := range outc {
		elme.Lock()

		fmt.Println(elme.count, hex.Dump(elme.buffer))

		elme.Unlock()
	}
}

func main() {
	var wait sync.WaitGroup
	inc := make(chan *data)
	outc := make(chan *data)

	cpus := runtime.NumCPU()

	wait.Add(cpus)
	for i := 0; i < cpus; i++ {
		go HASH(inc, &wait)
	}

	// wait.Add(1)
	// go HASH(inc, &wait)

	wait.Add(1)
	go Write(outc, &wait)

	f, _ := os.Open("text.txt")
	defer f.Close()

	sc := bufio.NewScanner(f)

	i := 0
	for sc.Scan() {
		i++
		// defer close(outc)
		line := sc.Text()
		data := data{
			count:  i,
			text:   line,
			buffer: nil,
			Mutex:  sync.Mutex{},
		}
		data.Lock()
		inc <- &data
		outc <- &data
	}

	close(inc)
	close(outc)

	wait.Wait()

}
