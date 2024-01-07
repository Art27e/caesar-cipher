package main

import (
	"io/ioutil"
	"os"
	//"regexp"
	"strings"
)

var str string
var result string
var lines []string

func cipher() {
	resultSlice2 := []rune(str)
	for i := 0; i < len(resultSlice2); i++ {
		if resultSlice2[i] >= 65 && resultSlice2[i] <= 67 || resultSlice2[i] >= 97 && resultSlice2[i] <= 99 {
			resultSlice2[i] = resultSlice2[i] + 23
		} else if resultSlice2[i] >= 68 && resultSlice2[i] <= 90 || resultSlice2[i] >= 100 && resultSlice2[i] <= 122 {
			resultSlice2[i] = resultSlice2[i] - 3
		}
		result = result + string(resultSlice2[i])
	}
}

func reader() {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	input := string(data)
	lines = append(lines, input)
	str = strings.Join(lines, " ")
}

func writer() {
	writer, _ := os.Create(os.Args[2])
	defer writer.Close()
	writer.WriteString(result)
}

func main() {
	reader()
	cipher()
	writer()
}
