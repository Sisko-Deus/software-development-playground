package main

import "fmt"

type KeyValueKeeper interface {
	Set(key int, value int) error
	Get(key int) (int, error)
	Check(key int) bool
	Remove(key int) error
}

type Storage struct {
	keeper KeyValueKeeper
}

func newStorage(keeper KeyValueKeeper) *Storage {
	return &Storage{keeper: keeper}
}

func main() {
	mapStorage := newMapStorage()
	storage1 := newStorage(mapStorage)
	storage1.keeper.Set(1, 10)
	fmt.Println("Check (1): ", storage1.keeper.Check(1))
	fmt.Println("Check (2): ", storage1.keeper.Check(2))
	v, err := storage1.keeper.Get(1)
	if err == nil {
		fmt.Println("Get (1): ", v)
	}

}
