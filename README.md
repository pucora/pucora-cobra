Velonetics Cobra
====

An adapter of the [cobra](http://github.com/spf13/cobra) lib for the [Velonetics](http://velonetics.io) framework

Package cmd defines the cobra command structs and an execution method for adding an improved CLI to
Velonetics based api gateways

## Basic example

```
package main

import (
	"os"

	"github.com/velonetics/velonetics-cobra/v2"
	"github.com/velonetics/velonetics-viper/v2"
	"github.com/velonetics/lura/v2/config"
	"github.com/velonetics/lura/v2/logging"
	"github.com/velonetics/lura/v2/proxy"
	veloneticsgin "github.com/velonetics/lura/v2/router/gin"
)

func main() {

	cmd.Execute(viper.New(), func(serviceConfig config.ServiceConfig) {
		logger, _ := logging.NewLogger("DEBUG", os.Stdout, "")
		veloneticsgin.DefaultFactory(proxy.DefaultFactory(logger), logger).New().Run(serviceConfig)
	})

}
```

## Available commands

The `cmd` package includes four commands: `check`, `check-plugin`, `help` and `run`.

1. *check* validates the received config file.
2. *check-plugin* validates the dependencies shared between the binary and a plugin.
3. *help* displays details about any command.
4. *run* executes the passed executor once the received flags overwrite the parsed config.

```
$ ./velonetics
 ╓▄█                          ▄▄▌                               ╓██████▄µ
▐███  ▄███╨▐███▄██H╗██████▄  ║██▌ ,▄███╨ ▄██████▄  ▓██▌█████▄  ███▀╙╙▀▀███╕
▐███▄███▀  ▐█████▀"╙▀▀"╙▀███ ║███▄███┘  ███▀""▀███ ████▀╙▀███H ███     ╙███
▐██████▌   ▐███⌐  ,▄████████M║██████▄  ║██████████M███▌   ███H ███     ,███
▐███╨▀███µ ▐███   ███▌  ,███M║███╙▀███  ███▄```▄▄` ███▌   ███H ███,,,╓▄███▀
▐███  ╙███▄▐███   ╙█████████M║██▌  ╙███▄`▀███████╨ ███▌   ███H █████████▀
                     ``                     `'`
Version: undefined

The API Gateway builder

Usage:
  velonetics [command]

Available Commands:
  check        Validates that the configuration file is valid.
  check-plugin Checks your plugin dependencies are compatible.
  help         Help about any command
  run          Runs the Velonetics server.

Flags:
  -h, --help   help for velonetics

Use "velonetics [command] --help" for more information about a command.

```
