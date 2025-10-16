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
	crashSite := Room{
		Name:        "The Crash Site",
		Description: "Empty for now.",
		Exits:       make(map[string]*Exit),
	}
	baseExterior := Room{
		Name:        "The Base's Exterior",
		Description: "Empty for now.",
		Exits:       make(map[string]*Exit),
	}
	solarArray := Room{
		Name:        "The Solar Array",
		Description: "Still empty.",
		Items:       make(map[string]*Item),
		Exits:       make(map[string]*Exit),
	}
	solarArray.Items["battery"] = &Item{
		Name:        "Battery",
		Description: "A battery. Can server to power electrical devices for a short time.",
	}
	crashSite.Exits["east"] = &Exit{
		Description: "Empty for noww",
		Destination: &baseExterior,
	}
	baseExterior.Exits["south"] = &Exit{
		Description: "Emptyyy",
		Destination: &solarArray,
	}
	baseExterior.Exits["west"] = &Exit{
		Description: "Empty",
		Destination: &crashSite,
	}
	solarArray.Exits["north"] = &Exit{
		Description: "Emptyed",
		Destination: &baseExterior,
	}
	player := Player{
		CurrentRoom: &crashSite,
		Inventory:   make(map[string]*Item),
	}
	for {
		fmt.Println("")
		fmt.Print(">")
		input, _ := reader.ReadString('\n')
		cleanInput := strings.TrimSpace(input)
		fieldsCommand := strings.Fields(cleanInput)
		command := fieldsCommand[0]
		arg1 := fieldsCommand[1]
		fmt.Printf("You entered the command :%s\n", cleanInput)
		switch game.GameMode {
		case ModeRoom:
			//stuff for when outside
			switch command {
			case "go":
				// stuff
				if exit, ok := player.CurrentRoom.Exits[arg1]; ok {
					fmt.Println("")
					fmt.Println("***************************")
					fmt.Println("")
					fmt.Println(exit.Description)
					fmt.Println("***************************")
					fmt.Println("")
					player.CurrentRoom = exit.Destination
					fmt.Println(player.CurrentRoom.Description)
				}
			case "look":
				//stuff
			case "take":
				//stuff
			case "use":
				//stuff
			default:
				fmt.Println("Unknown command in room mode.")
			}

		case ModeGrid:
			//stuff for when no light and when inside in general

		}
	}
}
