package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de cli pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "App cli"
	app.Usage = "Busca Ips e nomes de servidor na internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca Ips de endereço na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "salazarmanagement.com",
				},
			},
			Action: buscarIps,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	// net
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
