package main

import (
	"fmt"

	"github.com/azryve/locknock/internal/locknock"

	"github.com/spf13/cobra"
)

var (
	rulesetTargetPort    int
	rulesetPacketsNumber int
	rulesetHiddenPort    int
	rulesetKey           string
)

func rulesetRun(cmd *cobra.Command, args []string) error {
	rulesetKey, err := knockPassword()
	if err != nil {
		return err
	}
	params := knockParams(rulesetKey, rulesetPacketsNumber, rulesetTargetPort)
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
	cmd.Flags().IntVarP(&rulesetPacketsNumber, "num", "n", 10, "number of packets to knock with")
	cmd.Flags().IntVarP(&rulesetHiddenPort, "port-hidden", "P", 22, "tcp port number to lock (default is 22)")
	cmd.Flags().IntVarP(&rulesetTargetPort, "port-target", "T", 2222, "upd port to send knocks (default is 2222)")
	return cmd
}
