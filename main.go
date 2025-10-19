package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

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
type GridFeature struct {
	Name        string
	Description string
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
func typeWrite(text string, delay int, textColor ...color.Attribute) {
	if len(textColor) > 0 {
		color.Set(textColor...)
		defer color.Unset()
	}
	for _, char := range text {
		fmt.Printf("%c", char)
		totalDelay := delay * int(time.Millisecond)
		time.Sleep(time.Duration(totalDelay))
	}
}
func parseMap(mapString string) (world [][]int, startX int, startY int) {
	lines := strings.Split(strings.TrimSpace(mapString), "\n")
	world = make([][]int, len(lines))
	startX, startY = -1, -1
	for y, line := range lines {
		trimmedline := strings.TrimSpace(line)
		world[y] = make([]int, len(trimmedline))
		for x, char := range trimmedline {
			var tile = 0
			switch char {
			case '#':
				tile = 1
			case 'D':
				tile = 2

			case '@':
				startX, startY = x, y
				tile = 0
			default:
				tile = 0
			}
			world[y][x] = tile
		}

	}
	return world, startX, startY
}
func main() {
	baseMapString := `
######################################################
## ##               #                 #        (N2)  #
##@##    (N1)                         # (3)          #
## #########        #                 ##### ###################
##         #        #                 D              #        #
########## #        #                 #####         (D)   (T3)#
#      ### ##########   (T1)          # (T2)         #        #
#  (1)     #        #######           #########################
########## #      ################ #########        #
#      ### #        D                      #        #
#                 #####                    D    (2) #
#####################################################
	`
	worldGrid, startX, startY := parseMap(baseMapString)
	reader := bufio.NewReader(os.Stdin)
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
		Description: "A battery. Can serve to power electrical devices for a short amount of time.",
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
		X:           startX,
		Y:           startY,
	}
	game := Game{
		GameMode: ModeRoom,
		Player:   player,
		World:    worldGrid,
		PowerOn:  true,
	}
	gridFeatures := make(map[string]GridFeature)
	gridFeatures["7,4"] = GridFeature{Name: "lever", Description: "A power lever. Maybe actionning it could bring back power?"}

	baseExterior.Features["airlock"] = "The base's airlock."
	typeWrite("[*] ENGINE FAILURE\n", 40, color.FgRed)
	typeWrite("[*] INITING EMERGENCY PROCEDURES\n", 40, color.FgRed)
	typeWrite("[*] ACTIVATING EMERGENCY TERMAL SHIELDS\n", 40, color.FgRed)
	typeWrite("[*] ENTERING ATMOSPHERE\n", 40, color.FgRed)
	typeWrite("[*] PREPARING FOR THE IMPACT\n", 40, color.FgRed)
	typeWrite("[*] IMPACT IN :", 40, color.FgRed)
	typeWrite("  3, 2, 1.....\n", 500, color.FgRed)
	typeWrite("***************************", 40, color.FgRed)
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
					fmt.Println("[*] SYSTEM REPORT INCOMING :")
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
						break

					} else {
						fmt.Println("[*] SYSTEM REPORT : THERE'S NOTHING ELSE TO SEE HERE.")
					}
				}
				target := arg1
				if item, ok := player.Inventory[target]; ok {
					fmt.Println("In your inventory : ")
					fmt.Println(" * ", item.Name, " - ", item.Description)
					break

				}
				if item, ok := player.CurrentRoom.Items[target]; ok {
					fmt.Println("Items around you : ")
					fmt.Println(" * ", item.Name, " - ", item.Description)
					break
				}
				if feature, ok := player.CurrentRoom.Features[target]; ok {
					fmt.Println("Features around you : ")
					fmt.Println(" * ", feature)
					break

				}
				fmt.Println("[*] SYSTEM ERROR : NOTHING NAMED : ", target, " IN THE SURROUNDING.")
			case "take":
				if item, ok := player.CurrentRoom.Items[arg1]; ok {
					fmt.Println("")
					fmt.Println("***************************")
					fmt.Println("")
					fmt.Println("Taking : ", item.Name)
					player.Inventory[arg1] = item
					delete(player.CurrentRoom.Items, arg1)
					fmt.Println("You took", arg1, "and put it in your inventory.")
				} else if len(fieldsCommand) > 1 {
					fmt.Println("[*] SYSTEM ERROR : ITEM NOT FOUND : ", arg1)
				} else {
					fmt.Println("[*] SYSTEM ERROR : PLEASE SPECIFY AN ITEM TO TAKE")
				}
			case "use":
				if arg1 == "battery" {
					if _, ok := player.Inventory["battery"]; ok {
						if player.CurrentRoom == &baseExterior {
							fmt.Println("You open the control pannel, plug the two terminal of the battery, and then SHWOOSH! The airlock open wide! As you clicky enter inside, the lock close behidn you, and all the light's goes down : the battery didn't last long. You're now compeltly in the dark, and can't go behind. You will need to use your suit's sensors to move around and bring back the power...")
							delete(player.Inventory, "battery")
							game.PowerOn = false
							game.GameMode = ModeGrid
						} else {
							fmt.Println("[*] SYSTEM ERROR : NOT IN THE CURRENT ROOM.")
						}
					} else {
						fmt.Println("[*] SYSTEM ERROR : ITEM NOT IN INVENTORY.")
					}
				} else {
					fmt.Println("[*] SYSTEM ERROR : CANNOT USE THIS ITEM.")
				}
			case "help":
				fmt.Println("")
				fmt.Println("[*] Help menu :")
				fmt.Println("go [north/south/east/west] - go in the specified direction, if a path exist.")
				fmt.Println("look - Describe the surrounding, the items around, the different paths, etc.")
				fmt.Println("look [item/feature name] - Look at the specified item or feature, if it exists in the current room.")
				fmt.Println("take [item name] - Take the specified item name, if it exists in the current location.")
				fmt.Println("use [item] - Use the specified item if a) it exists in the inventory and b) it can be used in the current location.")

			default:
				fmt.Println("[*] SYSTEM ERROR : UNKNOWN COMMAND IN ROOM MODE.")

			}

		case ModeGrid:
			//stuff for when no light and when inside in general
			switch command {
			case "ping":
				// North pinggging
				for dist := 1; ; dist++ {
					checkY := game.Player.Y - dist
					if checkY < 0 || game.World[checkY][game.Player.X] == 1 {
						safeStep := dist - 1
						fmt.Printf("[*] SENSOR REPORT : Path clear for %d units to the North.\n", safeStep)
						break
					}

				}
				// South pinging
				for dist := 1; ; dist++ {
					checkY := game.Player.Y + dist
					if checkY >= len(game.World) || game.World[checkY][game.Player.X] == 1 {
						safeStep := dist - 1
						fmt.Printf("[*] SENSOR REPORT : Path Clear for %d units to the South.\n", safeStep)
						break

					}
				}
				// East piginging
				for dist := 1; ; dist++ {
					checkX := game.Player.X + dist
					if checkX >= len(game.World[game.Player.Y]) || game.World[game.Player.Y][checkX] == 1 {
						safeStep := dist - 1
						fmt.Printf("[*] SENSOR REPORT : Path clear for %d units to the East.\n", safeStep)
						break
					}
				}
				for dist := 1; ; dist++ {
					checkX := game.Player.X - dist
					if checkX < 0 || game.World[game.Player.Y][checkX] == 1 {
						safeStep := dist - 1
						fmt.Printf("[*] SENSOR REPORT : Path clear for %d units to the West.\n", safeStep)
						break
					}
				}

			case "go":
				direction := arg1
				distance := arg2
				if len(direction) > 0 {
					if len(distance) > 0 {
						distanceInt, err := strconv.Atoi(distance)
						if err != nil {
							fmt.Println("[*] SYSTEM ERROR : distance IS NOT A NUMBER.")
						}
						fmt.Println(distanceInt)
						stepTaken := 0
						for i := 0; i < distanceInt; i++ {
							currentX, currentY := game.Player.X, game.Player.Y
							nextX, nextY := currentX, currentY
							switch direction {
							case "north":
								nextY--
							case "south":
								nextY++
							case "east":
								nextX++
							case "west":
								nextX--
							default:
								fmt.Println("[*] SYSTEM ERROR : UNKNOWN DIRECTION.")
								goto endGoCommand
							}
							if nextY < 0 || nextY >= len(game.World) || nextX < 0 || nextX >= len(game.World[nextY]) {
								fmt.Println("[*] SENSOR REPORT : IMPACT IMMINENT. You reached the edge of the structure. Movement halted.")
								break
							}
							nextTile := game.World[nextY][nextX]
							if nextTile == 1 {
								fmt.Println("[*] SENSORT REPORT : IMPACT IMMINENT. Wall detected. Movement halted.")
								break
							}
							if nextTile == 2 {
								fmt.Println("[*] SENSORT REPORT : MOVEMENT HALTED. Airlock unlocking procedure required.")
								break
							}
							game.Player.X = nextX
							game.Player.Y = nextY
							stepTaken++

						}
						if stepTaken > 0 {
							fmt.Printf("[*] MOVEMENT REPORT : Moved %s %d steps. New coordinates: (%d, %d)\n", direction, stepTaken, game.Player.X, game.Player.Y)
						}
					endGoCommand:
						break
					} else {
						fmt.Println("[*] SYSTEM ERROR : USE THE COMMAND 'go' WITH THE ADDITIONAL ARGUMENT : distance.")
					}
				} else {
					fmt.Println("[*] SYSTEM ERROR : USE THE COMMAND 'go' WITH THE FOLLOWING ARGUMENTS : direction and distance.")
				}
			case "look":
				fmt.Println("[*] Proximity scanner activated...")
				foundSomething := false
				for coordStr, feature := range gridFeatures {
					var featureX, featureY int
					fmt.Sscanf(coordStr, "%d,%d", &featureX, &featureY)
					distX := game.Player.X - featureX
					if distX < 0 {
						distX = -distX
					}
					distY := game.Player.Y - featureY
					if distY < 0 {
						distY = -distY
					}
					totalDist := distX + distY
					if totalDist <= 3 {
						fmt.Printf("[*] SENSOR REPORT : Object of interest detected in a range of %d units : %s\n", totalDist, feature.Description)
						foundSomething = true
					}

				}
				if !foundSomething {
					fmt.Println("[*] SENSOR REPORT : No object of interest detected in the immediate vicinity.")
				}
			case "use":
				itemToUse := arg1
				if len(itemToUse) > 0 {
					var foundFeature *GridFeature
					var featureCoord string
					for coord, f := range gridFeatures {
						if f.Name == itemToUse {
							var featureX, featureY int
							fmt.Sscanf(coord, "%d,%d", &featureX, featureY)
							distX := game.Player.X - featureX
							if distX < 0 {
								distX = -distX
							}
							distY := game.Player.Y - featureY
							if distY < 0 {
								distY = -distY
							}
							if (distX + distY) <= 3 {
								foundFeature = &f
								featureCoord = coord
								break
							}
						}
						if foundFeature == nil {
							fmt.Println("[*] SYSTEM ERROR: Cannot use '", itemToUse, "'. It is not in the immediate vicinity.")
							break

						}
						switch foundFeature.Name {
						case "lever":
							if game.PowerOn {
								fmt.Println("[*] SYSTEM REPORT : Main pwoer already online.")

							} else {

							}
						}

					}
				} else {
					fmt.Println("[*] SYSTEM ERROR : Please specify what item to use.")
				}
			}

		default:
			fmt.Println("[*] SYSTEM ERROR : UNKNOWN COMMAND IN GRID MODE.")

		}
	}
}
