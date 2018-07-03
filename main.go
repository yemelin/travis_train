package main

import "fmt"

import "github.com/yemelin/travis_train/mprovider"

func main() {
	m := mprovider.New()
	fmt.Println(m.Message())
}
