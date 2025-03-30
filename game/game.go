package main

import (
	"fmt"
)

func main() {
	var i1 Item
	fmt.Printf("%#v\n", i1)
	fmt.Printf("i1: %#v\n", i1)

	i2 := Item{1, 2}
	fmt.Printf("i2: %#v\n", i2)

	i3 := Item{
		Y: 10,
		X: 20,
	}
	fmt.Printf("i3: %#v\n", i3)
	fmt.Println(NewItem(10, 20))
	fmt.Println(NewItem(10, -20))

	i3.Move(100, 200)
	fmt.Println("i3 (move): %#v\n", i3)

	p1 := Player{
		Name: "Strahinja",
		Item: Item{500, 300},
	}
	fmt.Printf("p1: %#v\n", p1)
	fmt.Printf("p1.X: %#v\n", p1.X)
	p1.Move(400, 600)
	fmt.Printf("p1 (move): %#v\n", p1)

	ms := []mover{
		&i1,
		&p1,
		&i2,
	}

	moveAll(ms, 0, 0)
	for _, m := range ms {
		fmt.Println(m)
	}

	k := Jade
	fmt.Println("k:", k)

	// time.Time import json.Marshal interface
	// json.NewEncoder(os.Stdout).Encode(time.Now())

	p1.FoundKey(Jade)
	fmt.Println(p1.Keys)
	p1.FoundKey(Jade)
	fmt.Println(p1.Keys)
}

// Implement fmt.Stringer interface
func (k Key) String() string {
	switch k {
	case Jade:
		return "jade"
	case Copper:
		return "copper"
	case Crystal:
		return "crystal"
	}

	return fmt.Sprintf("<Key %d>", k)
}

const (
	Jade Key = iota + 1
	Copper
	Crystal
	invalidKey
)

type Key byte

func moveAll(ms []mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

type mover interface {
	Move(x, y int)
	// Move(int,int)
}

func (p *Player) FoundKey(k Key) error {
	if k < Jade || k >= invalidKey {
		return fmt.Errorf("invalid key: %#v", k)
	}
	if !containsKey(p.Keys, k) {
		p.Keys = append(p.Keys, k)
	}
	return nil
}

func containsKey(keys []Key, k Key) bool {
	for _, k2 := range keys {
		if k2 == k {
			return true
		}
	}
	return false
}

type Player struct {
	Name string
	Item // embeded item
	Keys []Key
}

// i is called the reciver
// if you want to mutate, use pointer reciver
func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}

// func NewItem(x,y int) *Item
// func NewItem(x,y int) Item
// func NewItem(x, y int) (Item,error)
func NewItem(x, y int) (*Item, error) {
	if x < 0 || x > 1000 || y < 0 || y > 600 {
		return nil, fmt.Errorf("%d/%d out of bounds %d/%d", x, y, maxX, maxY)
	}

	i := Item{
		X: x,
		Y: y,
	}
	// the go compiler does escape analysis and will allocation i on the heap
	return &i, nil
}

const (
	maxX = 1000
	maxY = 600
)

// item is an item in the game
type Item struct {
	X int
	Y int
}
