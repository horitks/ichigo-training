package main

import (
	"fmt"
	"math/rand"
	"time"
)

// TODO: 定数ではなくデッキから減らしていけるようにするとか。
func getStack() []int {
	return []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 11, 11, 12, 12, 13, 14, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}
}

func getNumber() []int {
	stack := getStack()
	rand.Seed(time.Now().UnixNano())

	res := []int{}
	for i := 0; i < 5; i++ {
		num := rand.Intn(len(stack))
		res = append(res, stack[num])
		// 取ったカードはデッキから削除
		stack = append(stack[:num], stack[num+1:]...)
	}

	return res
}

func main() {
	fmt.Println("Pick up 5 card! Please make 15 with four arithmetic operations.")
	fmt.Println("")
	for _, v := range getNumber() {
		fmt.Print("[")
		fmt.Print(v)
		fmt.Print("]")
	}
	fmt.Println("")
}
