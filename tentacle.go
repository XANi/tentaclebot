package main

import (
	"github.com/op/go-logging"
	"github.com/urfave/cli"
	"os"
	"strings"
	"github.com/zerosvc/go-zerosvc"
)

var version string
var log = logging.MustGetLogger("main")
var stdout_log_format = logging.MustStringFormatter("%{color:bold}%{time:2006-01-02T15:04:05.0000Z-07:00}%{color:reset}%{color} [%{level:.1s}] %{color:reset}%{shortpkg}[%{longfunc}] %{message}")

func main() {
	stderrBackend := logging.NewLogBackend(os.Stderr, "", 0)
	stderrFormatter := logging.NewBackendFormatter(stderrBackend, stdout_log_format)
	logging.SetBackend(stderrFormatter)
	logging.SetFormatter(stdout_log_format)

	app := cli.NewApp()
	app.Name = "foobar"
	app.Description = "do foo to bar"
	app.Version = version
	app.HideHelp = true
	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "help, h", Usage: "show help"},
		cli.StringFlag{
			Name:   "url",
			Value:  "tcp://127.0.0.1:1883",
			Usage:  "It's an url. Encode password as user:pass@host",
			EnvVar: "MQ_URL",
		},
	}
	app.Action = MainLoop
	// to sort do that
	//sort.Sort(cli.FlagsByName(app.Flags))
	//sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)

}


func MainLoop(c *cli.Context) error {
	if c.Bool("help") {
		cli.ShowAppHelp(c)
		os.Exit(1)
	}
	node := zerosvc.NewNode("tena")
	log.Infof("Starting app version: %s", version)
	log.Infof("var example %s", c.GlobalString("url"))
	return nil
}
