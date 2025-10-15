package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Game struct {
	World    [][]int
	Player   Player
	PowerOn  bool
	KnownMap [][]bool
}
type Player struct {
	X         int
	Y         int
	Inventory map[string]*Item
}
type Item struct {
	Name        string
	Description string
}

func main() {
	// 0 mean empty space
	// 1 means wall
	// 2 means locked door
	// 3 means stuf you gotta fidn to unlock the door
	worldMap := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1},
		{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1},
		{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1},
		{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1},
		{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1},
		{1, 0, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}
	game := Game{
		World:   worldMap,
		Player:  Player{X: 2, Y: 2},
		PowerOn: true,
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("")
		fmt.Println(">")
		input, _ := reader.ReadString('\n')
		cleanInput := strings.TrimSpace(input)
		fmt.Printf("You entered the command :%s\n", cleanInput)
	}
}
