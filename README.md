## VoidECho
This is a text-based Go game, like my previous project TextHackventure.
 This was made for siege.hackclub.com. It follows this week's theme, 'signals', because the goal of the game is to read the space base's radio center, and send an emergency message trought the long-range frequency.
## Gameplay
This game feautre two distinct gameplay modes that change as you progess throught the story : 
- 'Room' based : when you're outside the base, you move around the location. You move around with commands like ```go north``` if your current room have an exit to the north. It's full commands list is the following :   
```go [north/south/east/west] - go in the specified direction, if a path exist.
look - Describe the surrounding, the items around, the different paths, etc.
look [item/feature name] - Look at the specified item or feature, if it exists in the current room.
take [item name] - Take the specified item name, if it exists in the current location.
use [item] - Use the specified item if a) it exists in the inventory and b) it can be used in the current location.
```  

- 'Grid' based : When you're in the base, the whole place is like a big grid : you move with commands like ```go north 10```. It's full command list is the following : 
```
help - display this help menu
ping - use your suit's sensors to scan the surrounding in a range of 3 units for obstacles.
go [north/south/east/west] (distance) - move in the specified direction, as logn as there is no wall/obstacles.
look - use your suit's sensors to scan for any object of interest in a range of 3 units.
map - when power is back on, show a full map of the base.
use [feature/item] - use the specified item, as long as you're at maximum 3 units of it.
```
## Installation 
You can either download the correct pre-built binary for your OS & CPU architecture, or build it yourself :

- ```git clone https://github.com/Hash-AK/VoidEcho```
- ```cd VoidEcho```
- ```go run .```
