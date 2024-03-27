package main

import (
	"fmt"
	"github.com/azryve/locknock/internal/locknock"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

func main() {
	var rootCmd = &cobra.Command{Use: "locknock"}
	rootCmd.AddCommand(rulesetCmd())
	rootCmd.AddCommand(knockCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func knockParams(key string, packetsNumber int) locknock.IPTablesParams {
	portGen := locknock.PortGenerator{Key: key}
	knockPorts := []int{}
	for i := 0; i < packetsNumber; i++ {
		knockPorts = append(knockPorts, portGen.Port())
	}
	return locknock.IPTablesParams{
		TargetPort:              rulesetPort,
		KnockPorts:              knockPorts,
		TargetReapTimeoutSecs:   30,
		InternalReapTimeoutSecs: 1,
	}
}

func knockPassword() (string, error) {
	password, ok := os.LookupEnv("LOCKNOCK_PASSWORD")
	if ok {
		return password, nil
	}
	fmt.Print("Enter Password: ")
	interactivePass, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", fmt.Errorf("failed to read passowrd")
	}
	return string(interactivePass), nil
}
