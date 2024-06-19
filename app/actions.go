package app

import (
	"fmt"
	"log"
	"net"
	"webinspect/utils"

	"github.com/urfave/cli/v2"
)

// SearchIps prints given host's IPv4 and IPv6 addresses
func SearchIps(c *cli.Context) error {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
	return nil
}

// SearchDNSTxt prints TXT record for given domain name
func SearchDNSTxt(c *cli.Context) error {
	domain := c.String("domain")

	txts, err := net.LookupTXT(domain)
	if err != nil {
		log.Fatal(err)
	}

	for _, txt := range txts {
		fmt.Println(txt)
	}
	return nil
}

// SearchServers prints server names for the given domain name
func SearchServers(c *cli.Context) error {
	domain := c.String("domain")

	servers, err := net.LookupNS(domain)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
	return nil
}

// HealthCheck prints the connection status with the given destination
func HealthCheck(c *cli.Context) error {
	port := c.String("port")
	if port == "" {
		port = "80"
	}

	status := utils.CheckDomainConnection(c.String("domain"), port)
	fmt.Println(status)
	return nil
}
