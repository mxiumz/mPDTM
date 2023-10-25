/*
シ MXIUM SYSTEMS シ
printhelp.go
Prints the help test
*/

package files

import "fmt"

func FuncHelp() {
	Banmain()
	var helpText = `
...................,'...................
...............'...;:lll'...............
..............'..'.,,:;lxx'.............  Built by *mxium systems* - R&D
.............; .'..,,:;,cdO'............  Founded By - https://github.com/m0ham3dx  
............:, ... '';,,:lkc............    
...........,'..       .';ldd............	
...........:.'   ....':,.';:;...........  Helper function for PDTM that displays 
............;,  . .':..;.';,:...........  a menu with desriptions of tools
...........,'.......:,;:',c.,..........'  which you select for installation
...........'.,.'.  .,'....:..:......''''   
............... ,.  ........;;;c'''.....  
............    ..  ..... .;':.l:',,''''  PDTM Repo Official- 
.........';      .. .. ..''','';clc;,,'.  https://github.com/projectdiscovery/pdtm
.......,'..,.       .......'';ccllddlc;'
.....,,.   '...   .......;l::c::;;;,;:lo
.,;:,. .    ..............;llc;,,,,;;;cl
.,.....:..       ....';;,;,'.,,'',:::;'.  No guarantee this program will work as expected !
;.,....'.'.  ....'..,,;:;,',,...,,;''.;l
.... .. ... . ........'.',,'. ..'....';c  FREE PALESTINE ! 
  .    ..... ...........'.'.. .....  ...
	`
	fmt.Println(helpText)
}
