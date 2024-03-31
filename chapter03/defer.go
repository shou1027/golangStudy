package main

import "fmt"

func main() {
	fmt.Println("開始")

	defer play()

	fmt.Println("終了")
}

func play() {
	fmt.Println("delayが呼び出されました")
}
