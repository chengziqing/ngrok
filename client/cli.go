package client

import (
	"flag"
	"fmt"
	"github.com/chengziqing/ngrok/version"
	"os"
)

const usage1 string = `Usage: %s [OPTIONS] <local port or address>
Options:
`

const usage2 string = `
Examples:
	guoqiangti 80
	guoqiangti -subdomain=example 8080
	guoqiangti -proto=tcp 22
	guoqiangti -hostname="example.com" -httpauth="user:password" 10.0.0.1


Advanced usage: guoqiangti [OPTIONS] <command> [command args] [...]
Commands:
	guoqiangti start [tunnel] [...]    Start tunnels by name from config file
	guoqiangti help                    Print help
	guoqiangti version                 Print guoqiangti version

Examples:
	guoqiangti start www api blog pubsub
	guoqiangti -log=stdout -config=guoqiangti.yml start ssh
	guoqiangti version

`

type Options struct {
	config    string
	logto     string
	authtoken string
	httpauth  string
	hostname  string
	protocol  string
	subdomain string
	command   string
	args      []string
}

func ParseArgs() (opts *Options, err error) {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage1, os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, usage2)
	}

	config := flag.String(
		"config",
		"",
		"Path to guoqiangti configuration file. (default: $HOME/.guoqiangti)")

	logto := flag.String(
		"log",
		"none",
		"Write log messages to this file. 'stdout' and 'none' have special meanings")

	authtoken := flag.String(
		"authtoken",
		"",
		"Authentication token for identifying an guoqiangti.com account")

	httpauth := flag.String(
		"httpauth",
		"",
		"username:password HTTP basic auth creds protecting the public tunnel endpoint")

	subdomain := flag.String(
		"subdomain",
		"",
		"Request a custom subdomain from the guoqiangti server. (HTTP only)")

	hostname := flag.String(
		"hostname",
		"",
		"Request a custom hostname from the guoqiangti server. (HTTP only) (requires CNAME of your DNS)")

	protocol := flag.String(
		"proto",
		"http",
		"The protocol of the traffic over the tunnel {'http','tcp'} (default: 'http')")

	flag.Parse()

	opts = &Options{
		config:    *config,
		logto:     *logto,
		httpauth:  *httpauth,
		subdomain: *subdomain,
		protocol:  *protocol,
		authtoken: *authtoken,
		hostname:  *hostname,
		command:   flag.Arg(0),
	}

	switch opts.command {
	case "start":
		opts.args = flag.Args()[1:]
	case "version":
		fmt.Println(version.MajorMinor())
		os.Exit(0)
	case "help":
		flag.Usage()
		os.Exit(0)
	case "":
		err = fmt.Errorf("Error: Specify a local port to tunnel to, or " +
			"an ngrok command.\n\nExample: To expose port 80, run " +
			"'ngrok 80'")
		return

	default:
		if len(flag.Args()) > 1 {
			err = fmt.Errorf("You may only specify one port to tunnel to on the command line, got %d: %v",
				len(flag.Args()),
				flag.Args())
			return
		}

		opts.command = "default"
		opts.args = flag.Args()
	}

	return
}
