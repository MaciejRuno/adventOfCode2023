package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	file, err := os.Open("./data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		vals := []int{}
		chars := []rune(scanner.Text())
		for i := range chars {
			char := string(chars[i])

			i, err := strconv.Atoi(char)
			if err == nil {
				vals = append(vals, i)
			}
		}
		sum = sum + vals[0] * 10 + vals[len(vals) - 1]
    }

	fmt.Printf("sum of values: %d", sum)
}