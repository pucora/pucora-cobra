package main

import (
	"os"

	cmd "github.com/velonetics/velonetics-cobra/v2"
	koanf "github.com/velonetics/velonetics-koanf"
	"github.com/velonetics/lura/v2/config"
	"github.com/velonetics/lura/v2/logging"
	"github.com/velonetics/lura/v2/proxy"
	"github.com/velonetics/lura/v2/router/gin"
)

func main() {
	cmd.Execute(koanf.New(), func(serviceConfig config.ServiceConfig) {
		logger, _ := logging.NewLogger("DEBUG", os.Stdout, "")
		gin.DefaultFactory(proxy.DefaultFactory(logger), logger).New().Run(serviceConfig)
	})
}
