package main

import (
	"fmt"
	"os"
	"strconv"

	zsend "github.com/blacked/go-zabbix"
	docopt "github.com/docopt/docopt-go"
	hierr "github.com/reconquest/hierr-go"
)

var (
	version = "[manual build]"
)

func main() {
	usage := `zabbix-agent-extension-kannel

Usage:
  zabbix-agent-extension-kannel [options] --discovery
  zabbix-agent-extension-kannel [options]

Options:
  -k --kannel <url>           Kannel url
                                [default: http://127.0.0.1:13000/status.txt].
  -z --zabbix <zabbix>        Hostname or IP address of zabbix server
                                [default: 127.0.0.1].
  -p --port <port>            Port of zabbix server [default: 10051].
  --prefix <prefix>           Add part of your prefix for key [default: None].

Commands:
  --discovery                 Run low-level discovery.

Other:
  -h --help                   Show this screen.
  -v --version                Show version.
`
	args, err := docopt.Parse(usage, nil, true, version, false)
	if err != nil {
		fmt.Println(hierr.Errorf(err, "can't parse docopt"))
		os.Exit(1)
	}

	zabbix := args["--zabbix"].(string)
	port, err := strconv.Atoi(args["--port"].(string))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	prefix := args["--prefix"].(string)
	if prefix == "None" {
		prefix = "kannel"
	} else {
		prefix = fmt.Sprintf("%s.%s", prefix, "kannel")
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	kannel := args["--kannel"].(string)

	stats, err := getKannelStats(kannel)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if args["--discovery"].(bool) {
		err = discovery(stats)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}

	var metrics []*zsend.Metric

	metrics = createMetrics(
		hostname,
		stats,
		metrics,
		prefix,
	)

	packet := zsend.NewPacket(metrics)
	sender := zsend.NewSender(
		zabbix,
		port,
	)
	sender.Send(packet)
	fmt.Println("OK")
}
