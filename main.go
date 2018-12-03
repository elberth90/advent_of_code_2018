package main

import (
	"fmt"
	"github.com/elberth90/advent_of_code_2018/day_3"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := read("day_3/input.txt")
	if err != nil {
		panic(err)
	}

	result, err := day_3.First(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func read(filename string) ([]string, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return strings.Split(strings.Trim(string(buf), "\n"), "\n"), nil
}
