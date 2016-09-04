package main

import (
	"github.com/codegangsta/cli"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "t, transmission",
			Usage: "Transmission RPC interface location",
		},
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
		torrentUri := c.String("location")
		username := c.String("username")
		password := c.String("password")
		transmissionUrl := c.String("transmission");

		/**
		GangstaCli doesn't like having mandatory flags, lengthy discussion:
		https://github.com/urfave/cli/issues/85
		 */
		if (transmissionUrl == "" || torrentUri == "") {
			cli.ShowAppHelp(c)
			os.Exit(1)
		}

		client := NewTransmissionClient(transmissionUrl, username, password)
		client.addTorrent(torrentUri)
	}

	app.Run(os.Args)
	time.Sleep(3 * time.Second)

}
