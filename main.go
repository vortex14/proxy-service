package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/urfave/cli/v2"

	"github.com/fatih/color"
	"github.com/vortex14/gotyphoon/extensions/servers/gin/domains/proxy"
	"github.com/vortex14/gotyphoon/log"
)

func init() {
	log.InitD()

	if len(os.Getenv("PROXY_LIST")) == 0 {
		color.Red("Not found proxy list!")
		os.Exit(1)
	}
}

func main() {
	app := &cli.App{
		Name: "Proxy service",
		UsageText: `
			run proxy server: pserver -r localhost`,
		Description: "Run proxy proxy server",
		Usage:       "pserver -r localhost",
		Commands: []*cli.Command{
			{
				Name:  "run",
				Usage: "Run proxy",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "host",
						Aliases: []string{"r"},
						Value:   "localhost",
						Usage:   "redis host from arg",
					},
					&cli.StringFlag{
						Name:    "prefix",
						Aliases: []string{"pr"},
						Value:   "prefix",
						Usage:   "prefix for redis key",
					},
					&cli.StringFlag{
						Name:    "domain",
						Aliases: []string{"d"},
						Value:   "https://httpbin.org/ip",
						Usage:   "main domain for check",
					},
					&cli.IntFlag{
						Name:    "CheckBlockedTime",
						Aliases: []string{"cb"},
						Value:   90,
						Usage:   "interval ( in sec) for check blocked proxy",
					},
					&cli.IntFlag{
						Name:    "CheckTime",
						Aliases: []string{"ct"},
						Value:   10,
						Usage:   "interval ( in sec) for check locked proxy",
					},
					&cli.IntFlag{
						Name:    "ConcurrentCheck",
						Aliases: []string{"cc"},
						Value:   20,
						Usage:   "interval ( in sec) for check locked proxy",
					},
					&cli.IntFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Value:   1329,
						Usage:   "port of service",
					},
				},
				Action: func(c *cli.Context) error {
					redisHost := c.String("host")
					CheckBlockedTime := c.Int("CheckBlockedTime")
					CheckTime := c.Int("CheckTime")
					BlockedTime := c.Int("BlockedTime")
					ConcurrentCheck := c.Int("ConcurrentCheck")
					Port := c.Int("port")
					Prefix := c.String("prefix")
					Domain := c.String("domain")

					color.Yellow(
						fmt.Sprintf(`Domain: %s redis: %s, CheckBlockedTime: %d, CheckTime: %d, BlockedTime: %d, ConcurrentCheck: %d, Prefix: %s, `,
							Domain, redisHost, CheckBlockedTime, CheckTime, BlockedTime,
							ConcurrentCheck, Prefix,
						),
					)

					_ = proxy.Constructor(&proxy.Settings{
						RedisHost:        redisHost,
						CheckBlockedTime: CheckBlockedTime,
						CheckTime:        CheckTime,
						BlockedTime:      BlockedTime,
						ConcurrentCheck:  ConcurrentCheck,
						Port:             Port,
						CheckHosts:       []string{Domain},
						PrefixNamespace:  Prefix,
					}).Run()

					return nil
				},
			},
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		color.Red("Err: ", err)
	}

}
