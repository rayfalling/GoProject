// test project main.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var address = "millionRandom.txt"
	fin, err := os.Open(address)
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(fin)
	m := make(map[string]int)
	for {
		data, Err1 := buf.ReadString('\n')
		if Err1 != nil {
			if Err1 == io.EOF {
				break
			}
		}
		elem, exist := m[data]
		if exist == false {
			m[data] = 1
		} else {
			m[data] = elem + 1
		}
	}
	for k, v := range m {
		fmt.Print(k, " : ", v, "\r\n")
	}
}
