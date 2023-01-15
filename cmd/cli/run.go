package cli

import (
	"errors"
	"github.com/mvazquezc/go-cli-template/pkg/example"
	"github.com/spf13/cobra"
)

var (
	stringParameter      string
	intParameter         int
	stringArrayParameter []string
)

func NewRunCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Exec the run command",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Validate command Args
			err := validateRunCommandArgs()
			if err != nil {
				return err
			}
			// We have the run command logic implemented in our example pkg
			err = example.RunCommandRun(stringParameter)
			if err != nil {
				return err
			}
			return err
		},
	}
	addRunCommandFlags(cmd)
	return cmd
}

func addRunCommandFlags(cmd *cobra.Command) {

	flags := cmd.Flags()
	flags.StringVarP(&stringParameter, "string-parameter", "s", "", "A string parameter")
	flags.IntVarP(&intParameter, "int-parameter", "i", 1, "An int parameter")
	flags.StringArrayVarP(&stringArrayParameter, "string-array", "a", []string{"example"}, "A string array parameter")
	cmd.MarkFlagRequired("string-parameter")
}

// validateCommandArgs validates that arguments passed by the user are valid
func validateRunCommandArgs() error {
	if intParameter != 1 {
		return errors.New("Not valid int-parameter")
	}
	return nil
}
