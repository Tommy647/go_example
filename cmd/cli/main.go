package main

import "github.com/Tommy647/go_example/internal/netconf"

func main() {
	// [gorozco::chicky @ go_example] >> >> >> $ go run cmd/cli/main.go -host=FW1 -pubkey=true
	// [gorozco::chicky @ go_example] >> >> >> $ ssh -i ~/.ssh/id_rsa.pub gorozco@r1 -p 830 -s netconf
	netconf.Conn()
}
