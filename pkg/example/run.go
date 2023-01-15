package example

import "fmt"

func RunCommandRun(stringParameter string) error {
	fmt.Printf("Run command executed with string parameter set to %s\n.", stringParameter)
	return nil
}
