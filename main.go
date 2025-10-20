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
##         #        #                                #        #
########## #        #                 #####         (D)   (T3)#
#      ### ##########   (T1)          # (T2)         #        #
#  (1)     #        #######           #########################
########## #      ################ #########        #
#      ### #                               #        #
#                 #####                    D    (2) #
#####################################################
	`
	worldGrid, startX, startY := parseMap(baseMapString)
	reader := bufio.NewReader(os.Stdin)
	crashSite := Room{
		Name:        "The Crash Site",
		Description: "You look around you : you're in the remmeneants of your space capsule. You're still a bit shaked, but you need to get to the base, which is about 500 meters away to the east.",
		Exits:       make(map[string]*Exit),
	}
	baseExterior := Room{
		Name:        "The Base's Exterior",
		Description: "As you reach the base's airlock, you realize that because there is no power, you won't be able to get in. You will need an external source of power to make the airlock open. You recall there's a battery near the Solar Pannel Array, to the south..",
		Exits:       make(map[string]*Exit),
		Features:    make(map[string]string),
	}
	solarArray := Room{
		Name:        "The Solar Array",
		Description: "As you reach the Solar Pannel Array, you find the. There's a locker with the battery in it. A quick glance on it's power level show that it's only at 5%... You will need to take it like that, anyways...",
		Items:       make(map[string]*Item),
		Exits:       make(map[string]*Exit),
	}
	solarArray.Items["battery"] = &Item{
		Name:        "battery",
		Description: "A battery. Can serve to power electrical devices for a short amount of time.",
	}
	crashSite.Exits["east"] = &Exit{
		Description: " ",
		Destination: &baseExterior,
	}
	baseExterior.Exits["south"] = &Exit{
		Description: " ",
		Destination: &solarArray,
	}
	baseExterior.Exits["west"] = &Exit{
		Description: " ",
		Destination: &crashSite,
	}
	solarArray.Exits["north"] = &Exit{
		Description: " ",
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
	gridFeatures["4,7"] = GridFeature{Name: "lever", Description: "A power lever. Maybe actionning it could bring back power?"}
	gridFeatures["25,6"] = GridFeature{Name: "terminal1", Description: "A Computer terminal (terminal1). It's sole purpose is to unlock the door of the Equipement Room (10,43). You can intereact with it with 'use terminal1'. It seems to require a password thought."}
	gridFeatures["39,6"] = GridFeature{Name: "terminal2", Description: "A terminal (terminal2) necessary to open the door of the Radio Station. Require the keycard from the Equipement Room."}
	gridFeatures["9,2"] = GridFeature{Name: "note1", Description: "A note (note1). It reads:\nThe password is 'VOID'."}
	gridFeatures["48,10"] = GridFeature{Name: "keycard", Description: "A shiny access keycard (keycard). It probably serves to unlock a terminal out there..."}
	gridFeatures["61,5"] = GridFeature{Name: "radio", Description: "The long-range communication radio (radio). Permit to contact the Earth."}
	baseExterior.Features["airlock"] = "The base's airlock."
	var equipementRoomLocked = true
	var radioRoomLocked = true
	typeWrite("[*] ENGINE FAILURE\n", 40, color.FgRed)
	typeWrite("[*] INITING EMERGENCY PROCEDURES\n", 40, color.FgRed)
	typeWrite("[*] ACTIVATING EMERGENCY TERMAL SHIELDS\n", 40, color.FgRed)
	typeWrite("[*] ENTERING ATMOSPHERE\n", 40, color.FgRed)
	typeWrite("[*] PREPARING FOR THE IMPACT\n", 40, color.FgRed)
	typeWrite("[*] IMPACT IN :", 40, color.FgRed)
	typeWrite("  3, 2, 1.....\n", 500, color.FgRed)
	typeWrite("***************************", 40, color.FgRed)
	typeWrite("\n\nYou regain consciousness. Your head is pounding, and you can barely remember what happened. You slowly remember the incident : the takeoff from the base HCSW-3 turned dramatic : as your suit's sensors showed, there was an engine failure, which made your whole spaceship go loose altitue and crash onto the base. The wreckage cut the power in the base's main center. You need to get to the base, bring back power, then send an emergency call from the radio station.", 40)
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
								if nextX == 43 && nextY == 10 {
									if equipementRoomLocked {
										fmt.Println("[*] MOVEMENT HALTED. The Equipement Room door is locked.")
										break
									}

								} else if nextX == 53 && nextY == 5 {
									if radioRoomLocked {
										fmt.Println("[*] MOVEMENT HALTED. The Radio Room Door is locked.")
										break
									}
								} else {
									fmt.Println("[*] MOVEMENT HALTED. Door is sealed.")
								}

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
					//fmt.Printf("DEBUG: Attempting to use item: --%s--\n", itemToUse)
					var foundFeature *GridFeature
					var featureCoord string
					for coord, f := range gridFeatures {
						//fmt.Printf("DEBUG: Checking feature with name: --%s--\n", f.Name)
						if f.Name == itemToUse {
							//fmt.Println("DEBUG: Name match SUCCESS!")
							var featureX, featureY int
							fmt.Sscanf(coord, "%d,%d", &featureX, &featureY)
							distX := game.Player.X - featureX
							if distX < 0 {
								distX = -distX
							}
							distY := game.Player.Y - featureY
							if distY < 0 {
								distY = -distY
							}
							if (distX + distY) <= 3 {
								//fmt.Println("DEBUG: Proximity check SUCCESS!")
								tempFeature := f
								foundFeature = &tempFeature
								featureCoord = coord
								break
							}
						}
					}
					if foundFeature == nil {
						fmt.Println("DEBUG: 'foundfeature' is nil after search.")
						fmt.Println("[*] SYSTEM ERROR: Cannot use '", itemToUse, "'. It is not in the immediate vicinity.")
						break

					}
					switch foundFeature.Name {
					case "lever":
						if game.PowerOn {
							fmt.Println("[*] SYSTEM REPORT : Main pwoer already online.")

						} else {
							fmt.Println("You pull the heavy lever. As you reach it's 'On' position, a deep hum resonates trhought the station! All the lights are backup : you can now say correctly everything. You can now use the command 'map' to see the full map of the base.")
							game.PowerOn = true
							originalFeature := gridFeatures[featureCoord]
							originalFeature.Description = "The power lever is now in the 'ON' position."
							gridFeatures[featureCoord] = originalFeature
						}
					case "note1":
						fmt.Println(foundFeature.Description)
					case "terminal1":
						fmt.Println("The terminal screen flickers, asking for a password.")
						fmt.Print("ENTER PASSWORD > ")
						input, _ := reader.ReadString('\n')
						if strings.TrimSpace(input) == "VOID" {
							fmt.Println("ACCESS GRANTED. The door to the Equipement Room slides open.")
							equipementRoomLocked = false
							game.World[10][43] = 0
						} else {
							fmt.Println("ACCESS DENIED.")
						}
					case "keycard":
						fmt.Println("You pick up the keycard and clip it to your suit.")
						game.Player.Inventory["keycard"] = &Item{Name: "keycard", Description: "A standart access keycard."}
						delete(gridFeatures, featureCoord)
					case "terminal2":
						if _, hasKeycard := game.Player.Inventory["keycard"]; hasKeycard {
							fmt.Println("You insert the keycard in the terminal's slot. A green message appear on it : ACCESS GRANTED.")
							fmt.Println("The door to the Radio Room is now unlocked.")
							radioRoomLocked = false
							game.World[5][53] = 0
						} else {
							fmt.Println("ACCESS DENIED. Keycard required.")
						}
					case "radio":
						if radioRoomLocked {
							fmt.Println("You cannot raech the radio, the door is locked. Use terminal2 to unlock it first.")

						} else {
							typeWrite("You power on the radio and tune it to the emergency long-range frequency... After a long silence, a voice crackles back : '...copy that, HCSW-3. We read you. Help is on the way. Hold tight. Over.'\n\n*** YOU WIN! ***", 50, color.FgBlue)
							os.Exit(0)
						}
					default:
						fmt.Println("[*] SYSTEM ERROR: you can't use the '", foundFeature.Name, "'in that way.")

					}

				} else {
					fmt.Println("[*] SYSTEM ERROR : Please specify what item to use.")
				}
			case "map":
				if game.PowerOn == false {
					fmt.Println("[*] SYSTEM ERROR: To use the command 'map' you need the power back on first.")
				} else {
					//stuff for full map
					fmt.Println("\n--- STATION BLUEPRINTS ---")
					for y, row := range game.World {
						for x, tile := range row {
							if game.Player.X == x && game.Player.Y == y {
								color.Set(color.FgGreen)
								fmt.Print("@")
								color.Unset()
								continue
							}
							currentPosKey := fmt.Sprintf("%d,%d", x, y)
							if feature, ok := gridFeatures[currentPosKey]; ok {
								color.Set(color.FgYellow)
								fmt.Print(strings.ToUpper(string(feature.Name[0])))
								color.Unset()
								continue
							}

							switch tile {
							case 1:
								fmt.Print("#")
							case 2:
								color.Set(color.FgBlue)
								fmt.Print("D")
								color.Unset()
							default:
								fmt.Print(".")

							}
						}
						fmt.Println("")

					}

					fmt.Println("-------------------")
					fmt.Println("D = door, # = wall, @ = player position. T = terminal, N = note, L = lever, R = radio, K = keytag")
				}
			case "help":
				fmt.Println("")
				fmt.Println("help - display this help menu")
				fmt.Println("ping - use your suit's sensors to scan the surrounding in a range of 3 units for obstacles.")
				fmt.Println("go [north/south/east/west] (distance) - move in the specified direction, as logn as there is no wall/obstacles.")
				fmt.Println("look - use your suit's sensors to scan for any object of interest in a range of 3 units.")
				fmt.Println("map - when power is back on, show a full map of the base.")
				fmt.Println("use [feature/item] - use the specified item, as long as you're at maximum 3 units of it.")
				fmt.Println("")
			}

		default:
			fmt.Println("[*] SYSTEM ERROR : UNKNOWN COMMAND IN GRID MODE.")

		}
	}
}
