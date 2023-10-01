package main

import (
	"fmt"
)

type Number interface {
	int | float64
}

type StorageInter[T any] interface {
	GetItem() T
	StoreItem(T)
}

type Storage[T any] struct {
	Item T
}

func (s *Storage[T]) GetItem() T {
	return s.Item
}

func (s *Storage[T]) StoreItem(item T) {
	s.Item = item
}

var _ StorageInter[int] = &Storage[int]{}

func printIt[T any](item T) {
	fmt.Println(item)
}

func bigger[T Number, K any](a T, b T, prefix K) {
	if a > b {
		fmt.Println(prefix, a)
		return
	}
	fmt.Println(prefix, b)

}

func main() {
	printIt(1)       // 1
	printIt("Hello") // Hello

	bigger(1, 2, "The bigger integer is:")   // The bigger integer is: 2
	bigger(3.0, 2.0, "The bigger float is:") // The bigger float is: 3

	var intStorage StorageInter[int] = &Storage[int]{}
	intStorage.StoreItem(789)
	fmt.Println(intStorage.GetItem())

	var stringStorage StorageInter[string] = &Storage[string]{}
	stringStorage.StoreItem("This is a string")
	fmt.Println(stringStorage.GetItem())
}
