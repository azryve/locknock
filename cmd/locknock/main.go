package main

import (
	"fmt"
	"os"

	"github.com/azryve/locknock/internal/locknock"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var Version = "dev"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
	},
}

func main() {
	var rootCmd = &cobra.Command{Use: "locknock"}
	rootCmd.AddCommand(rulesetCmd())
	rootCmd.AddCommand(knockCmd())
	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func knockParams(key string, packetsNumber int) locknock.IPTablesParams {
	knockGen := locknock.KnockGenerator{Key: key}
	knocks := []uint32{}
	for i := 0; i < packetsNumber; i++ {
		knocks = append(knocks, knockGen.Port())
	}
	return locknock.IPTablesParams{
		HiddenPort:              rulesetPort,
		Knocks:                  knocks,
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
