package main

import (
	"fmt"
	"github.com/azryve/locknock/internal/locknock"

	"github.com/spf13/cobra"
)

var (
	rulesetPacketsNumber int
	rulesetPort          int
	rulesetKey           string
)

func rulesetRun(cmd *cobra.Command, args []string) error {
	rulesetKey, err := knockPassword()
	if err != nil {
		return err
	}
	params := knockParams(rulesetKey, rulesetPacketsNumber)
	iptRend := locknock.IPTablesRulesRenderer{
		Params: params,
	}
	fmt.Print(iptRend.Render())
	return nil
}

func rulesetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ruleset",
		Short: "generate port knocking ruleset",
		RunE:  rulesetRun,
	}
	cmd.Flags().IntVarP(&rulesetPort, "port", "P", 22, "Port number to lock (default is 22)")
	cmd.Flags().IntVarP(&rulesetPacketsNumber, "num", "n", 100, "number of packets to knock with")
	return cmd
}
