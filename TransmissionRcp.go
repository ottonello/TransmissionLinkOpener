package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"log"
	"time"
)

const transmissionUrl = "http://192.168.1.166:9091/transmission/rpc"

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "location, l",
			Usage: "Specify a torrent/magnet URI to download",
		},
		cli.StringFlag{
			Name:  "username, u",
			Usage: "Basic Auth username",
		},
		cli.StringFlag{
			Name:  "password, p",
			Usage: "Basic Auth password",
		},
	}

	app.Name = "Transmission URI adder"
	app.Usage = "Call Transmission RCP to add a torrent"

	app.Action = func(c *cli.Context) {
		uri := c.String("location")
		username := c.String("username")
		password := c.String("password")

		client := NewTransmissionClient(username, password)
		if (uri != "") {
			fmt.Printf("Trying to add specified URI to Transmission rpc at '%s'\n", transmissionUrl)
			client.addTorrent(uri)
		} else {
			/**
			GangstaCli doesn't like having mandatory flags, lengthy discussion:
			https://github.com/urfave/cli/issues/85
			 */
			log.Fatalln("uri is required")
		}
	}

	app.Run(os.Args)
	time.Sleep(3  * time.Second)

}
