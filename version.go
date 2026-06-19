package cmd

import (
	"github.com/pucora/lura/v2/core"
	"github.com/spf13/cobra"
)

func versionFunc(cmd *cobra.Command, _ []string) {
	cmd.Println("Pucora Version:", core.PucoraVersion)
	cmd.Println("Go Version:", core.GoVersion)
	cmd.Println("Glibc Version:", core.GlibcVersion)
}
