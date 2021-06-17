package calculator

import (
	"github.com/aerogear/charmil/pkg/core"
	"github.com/namit-chandwani/calculator-charmil-plugin/pkg/cmdloader"
	"github.com/spf13/cobra"
)

type CalculatorPlugin struct{}

func init() {
	core.Register("calculator", &CalculatorPlugin{})
}

func (p CalculatorPlugin) CreateRootCmd( /*Add config arg here*/ ) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "calculator",
		Short:        "Calculator CLI",
		Long:         `A basic calculator CLI built in Golang using the Cobra library.`,
		SilenceUsage: true,
	}

	cmdloader.AddCommands(cmd)

	return cmd
}
