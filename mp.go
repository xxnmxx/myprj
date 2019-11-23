package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	i := scanner.Text()
	ii := strings.Split(i, " ")
	sum := 0
	for _, v := range ii {
		num, _ := strconv.Atoi(v)
		sum = sum + num
	}
	fmt.Println(sum)
}

func sum(i ...int) int {
	sum := 0
	for _, v := range i {
		sum = sum + v
	}
	return sum
}
