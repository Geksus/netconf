package main

import (
	"fmt"
	"github.com/exsver/netconf/helpers/comware"
	"log"
)

func main() {

	// Creating new device
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf", "netconf")
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Preparing set of commands
	commands := `
interface GigabitEthernet1/0/3
 default
 description to_switch_205
 port link-type trunk
 undo port trunk permit vlan 1
 port trunk permit vlan 300 to 302
 quit
display current-configuration interface GigabitEthernet1/0/3
`
	// Executing commands in device
	output, err := sw.RunCLICommand(commands, true)
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Printing output
	fmt.Printf("%s", comware.CorrectNewLines(output))
}
