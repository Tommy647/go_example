package main

import (
	"bufio"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Tommy647/go_example/internal/jnprsvc"
	netconf2 "github.com/phersus/go-netconf/netconf"

	"github.com/Tommy647/go_example/internal/jnprsvc/transport/netconf"
)

var (
	jnprHost = flag.String("host", "192.168.123.123", "Target device hostname/IP@")
)

func main() {

	// [gorozco::chicky @ go_example] >> >> >> $ go run cmd/cli/main.go -host=FW1 -pubkey=true
	// [gorozco::chicky @ go_example] >> >> >> $ ssh -i ~/.ssh/id_rsa.pub gorozco@r1 -p 830 -s netconf

	flag.Parse()
	if strings.EqualFold(*jnprHost, "") {
		fmt.Print("hostname/IP@ of the target device: ")
		r := bufio.NewScanner(os.Stdin)
		r.Scan()
		*jnprHost = r.Text()
	}

	s, err := netconf.Conn(*jnprHost)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	r := jnprsvc.NewInstanceInformation()

	reply, _ := s.Exec(netconf2.RawMethod("<get-instance-information><summary/></get-instance-information>"))

	err = xml.Unmarshal([]byte(reply.RawReply), r)
	if err != nil {
		log.Fatal(err)
	}
	for i := range r.InstanceInformation.InstanceCore {
		if !strings.HasPrefix(r.InstanceInformation.InstanceCore[i].InstanceName, "__") {
			fmt.Println(r.InstanceInformation.InstanceCore[i].InstanceName)
		}
	}
}
