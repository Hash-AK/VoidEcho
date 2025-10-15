package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	ModeRoom = "room"
	ModeGrid = "grid"
)

type Game struct {
	World    [][]int
	Player   Player
	PowerOn  bool
	KnownMap [][]bool
	GameMode string
}
type Player struct {
	X           int
	Y           int
	Inventory   map[string]*Item
	CurrentRoom *Room
}
type Item struct {
	Name        string
	Description string
}
type Exit struct {
	Description string
	Destination *Room
}
type Room struct {
	Name        string
	Description string
	Exits       map[string]*Exit
	Items       map[string]*Item
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	game := Game{
		GameMode: ModeRoom,
	}
	for {
		fmt.Println("")
		fmt.Println(">")
		input, _ := reader.ReadString('\n')
		cleanInput := strings.TrimSpace(input)
		fmt.Printf("You entered the command :%s\n", cleanInput)
		switch game.GameMode {
		case ModeRoom:
			//stuff for when outside
		case ModeGrid:
			//stuff for when no light and when inside in general
		}
	}
}
