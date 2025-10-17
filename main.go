package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
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
	Features    map[string]string
}

func aOrAn(s string) string {
	if s == "" {
		return "a"
	}
	switch strings.ToLower(string(s[0])) {
	case "a", "e", "i", "o", "u":
		return "an"
	default:
		return "a"
	}
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
		Features:    make(map[string]string),
	}
	solarArray := Room{
		Name:        "The Solar Array",
		Description: "Still empty.",
		Items:       make(map[string]*Item),
		Exits:       make(map[string]*Exit),
	}
	solarArray.Items["battery"] = &Item{
		Name:        "battery",
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
	baseExterior.Features["airlock"] = "The base's airlock."
	color.Red("[*] ENGINE FAILURE")
	color.Red("[*] INITING EMERGENCY PROCEDURE")
	color.Red("[*] ENTERING ATMOSPHERE")
	color.Red("[*] PREPARING FOR THE IMPACT")
	color.Red("***************************")
	for {
		fmt.Println("")
		fmt.Print(">")
		input, _ := reader.ReadString('\n')
		cleanInput := strings.TrimSpace(input)
		fieldsCommand := strings.Fields(cleanInput)
		var arg1, arg2 string
		command := fieldsCommand[0]
		if len(fieldsCommand) > 1 {
			arg1 = fieldsCommand[1]
		}
		if len(fieldsCommand) > 2 {
			arg2 = fieldsCommand[2]
		}
		fmt.Printf("You entered the command : %s\n", cleanInput)
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
				} else {
					fmt.Println("[*] SYSTEM ERROR : NO PATH GOING TOWARD THIS WAY.")
				}
			case "look":
				if len(fieldsCommand) == 1 {
					fmt.Println("")
					fmt.Println("You are in : " + player.CurrentRoom.Name)
					fmt.Println(player.CurrentRoom.Description)
					var thingToSee []string
					for _, item := range player.CurrentRoom.Items {
						prefix := aOrAn(item.Name)
						thingToSee = append(thingToSee, prefix+" "+item.Name)
					}
					for featureName := range player.CurrentRoom.Features {
						prefix := aOrAn(featureName)
						thingToSee = append(thingToSee, prefix+" "+featureName)
					}
					for direction, exit := range player.CurrentRoom.Exits {
						thingToSee = append(thingToSee, "an exit toward the "+direction+" leading to "+exit.Destination.Name)
					}
					if len(thingToSee) > 0 {
						fmt.Println("")
						fmt.Println("You also see :")
						for _, thing := range thingToSee {
							fmt.Println(" - ", thing)
						}

					}
				}
			case "take":
				if item, ok := player.CurrentRoom.Items[arg1]; ok {
					fmt.Println("")
					fmt.Println("***************************")
					fmt.Println("")
					fmt.Println("Taking : ", item.Name)
					player.Inventory[arg1] = item
					delete(player.CurrentRoom.Items, arg1)
				} else {
					fmt.Println("[*] SYSTEM ERROR : ITEM NOT FOUND]")
				}
			case "use":
				if arg1 == "battery" {
					if _, ok := player.Inventory["battery"]; ok {
						if player.CurrentRoom == &baseExterior {
							// do shtuff because power
						} else {
							fmt.Println("[*] SYSTEM ERROR : NOT IN THE CURRENT ROOM.")
						}
					} else {
						fmt.Println("[*] SYSTEM ERROR : ITEM NOT IN INVENTORY.")
					}
				} else {
					fmt.Println("[*] SYSTEM ERROR : CANNOT USE THIS ITEM.")
				}
			default:
				fmt.Println("Unknown command in room mode.")

			}

		case ModeGrid:
			//stuff for when no light and when inside in general
			switch command {
			case "go":
				direction := arg1
				distance := arg2
				fmt.Println(direction, distance)
			}
		}
	}
}
