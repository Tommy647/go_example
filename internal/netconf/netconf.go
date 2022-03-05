package netconf

import (
	"bufio"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"

	"golang.org/x/crypto/ssh"

	"github.com/phersus/go-netconf/netconf"
)

var (
	host         = flag.String("host", "R1", "Hostname")
	username     = flag.String("username", "", "Username")
	key          = flag.String("key", os.Getenv("HOME")+"/.ssh/id_rsa", "SSH private key file")
	passphrase   = flag.String("passphrase", "", "SSH private key passphrase (cleartext)")
	nopassphrase = flag.Bool("nopassphrase", true, "SSH private key does not contain a passphrase")
	pubkey       = flag.Bool("pubkey", false, "Use SSH public key authentication")
	agent        = flag.Bool("agent", false, "Use SSH agent for public key authentication")
)

// SystemInformation provides a representation of the system-information container
type SystemInformation struct {
	HardwareModel string `xml:"system-information>hardware-model"`
	OsName        string `xml:"system-information>os-name"`
	OsVersion     string `xml:"system-information>os-version"`
	SerialNumber  string `xml:"system-information>serial-number"`
	HostName      string `xml:"system-information>host-name"`
}

// BuildConfig captures information from the console to build a SSH Client Config
func BuildConfig() *ssh.ClientConfig {

	var config *ssh.ClientConfig

	var pass string
	if *pubkey {
		if *agent {
			var err error
			config, err = netconf.SSHConfigPubKeyAgent(*username)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			if *nopassphrase {
				pass = "\n"
			} else {
				if *passphrase != "" {
					pass = *passphrase
				} else {
					var readpass []byte
					var err error
					fmt.Printf("Enter Passphrase for %s: ", *key)
					readpass, err = terminal.ReadPassword(syscall.Stdin)
					if err != nil {
						log.Fatal(err)
					}
					pass = string(readpass)
					fmt.Println()
				}
			}
			var err error
			config, err = netconf.SSHConfigPubKeyFile(*username, *key, pass)
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		fmt.Printf("Enter Password: ")
		bytePassword, err := terminal.ReadPassword(syscall.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println()

		config = netconf.SSHConfigPassword(*username, string(bytePassword))
	}
	return config
}

func Conn() {

	flag.Parse()

	if *username == "" {
		fmt.Printf("Enter a valid username: ")
		r := bufio.NewScanner(os.Stdin)
		r.Scan()
		*username = r.Text()
	}
	config := BuildConfig()
	s, err := netconf.DialSSH(*host, config)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	// show system information | display xml
	reply, err := s.Exec(netconf.RawMethod("<get-system-information/>"))
	if err != nil {
		panic(err)
	}
	var q SystemInformation

	err = xml.Unmarshal([]byte(reply.RawReply), &q)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("hostname: %s\n", q.HostName)
	fmt.Printf("model: %s\n", q.HardwareModel)
	fmt.Printf("version: %s\n", q.OsVersion)
	fmt.Printf("serial: %s\n", q.SerialNumber)
}
