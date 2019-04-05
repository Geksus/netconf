package main

import (
	"fmt"
	"github.com/exsver/netconf"
	"github.com/exsver/netconf/helpers/junos"
	"log"
	"strings"
)

func main() {
	netconf.LogLevel.Messages()

	sw, err := junos.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	ARPTable, err := sw.GetARPTableInformation()

	if err != nil {
		log.Fatalf("%s", err)
	}

	for _, ARP := range ARPTable.ARPTableEntries {
		fmt.Printf("%s %s %s\n", strings.TrimSpace(ARP.MACAddress), strings.TrimSpace(ARP.IPAddress), strings.TrimSpace(ARP.InterfaceName))
	}

	fmt.Printf("Total entries: %v", ARPTable.ARPEntryCount)
}
