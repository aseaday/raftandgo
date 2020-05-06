package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
)

func main() {
	toWrite := ""
	for i := 0; i < 100000; i++ {
		number := rand.Int31()
		toWrite = toWrite + strconv.FormatInt(int64(number), 10) + "\n"
	}
	err := ioutil.WriteFile("data", []byte(toWrite), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
