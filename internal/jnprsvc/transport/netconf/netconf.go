package netconf

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/Juniper/go-netconf/netconf"
	"golang.org/x/crypto/ssh/terminal"

	"golang.org/x/crypto/ssh"
)

const (
	envUsername = `USERNAME` // Username to log in on the juniper device
	envPassword = `PASSWORD` // Password to use to log in on the juniper device
)

func conConfig() *ssh.ClientConfig {
	flag.Parse()
	var config *ssh.ClientConfig

	jnprUsername := os.Getenv(envUsername)
	jnprPassword := os.Getenv(envPassword)

	if strings.EqualFold(jnprUsername, "") {
		fmt.Print("juniper username: ")
		r := bufio.NewScanner(os.Stdin)
		r.Scan()
		jnprUsername = r.Text()
	}
	if strings.EqualFold(jnprPassword, "") {
		fmt.Print("juniper password: ")
		bytePassword, err := terminal.ReadPassword(syscall.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println()
		config = netconf.SSHConfigPassword(jnprUsername, string(bytePassword))
	}
	return config
}

func Conn(jnprHost string) (*netconf.Session, error) {
	config := conConfig()
	return netconf.DialSSH(jnprHost, config)
}
