package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"

	"encoding/binary"

	"github.com/spf13/cobra"
)

var (
	knockTargetPort    int
	knockPacketsNumber int
	knockKey           string
	knockPortHidden    int
)

func knockRun(cmd *cobra.Command, args []string) error {
	knockKey, err := knockPassword()
	if err != nil {
		return err
	}
	hostname := args[0]
	params := knockParams(knockKey, knockPacketsNumber, knockTargetPort)
	address := fmt.Sprintf("%s:%d", hostname, params.KnockPort)
	udpaddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return err
	}
	sock, err := net.DialUDP("udp", nil, udpaddr)
	if err != nil {
		return err
	}
	defer sock.Close()
	for _, knock := range params.Knocks {
		bytes := make([]byte, 0, 4)
		bytes = binary.BigEndian.AppendUint32(bytes, knock)
		sock.Write(bytes)
	}
	if knockPortHidden != 0 {
		knockProxy(hostname)
	}
	return nil
}

func knockProxy(hostname string) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", hostname, knockPortHidden))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		io.Copy(conn, os.Stdin)
		wg.Done()
	}()
	go func() {
		io.Copy(os.Stdout, conn)
		wg.Done()
	}()
	wg.Wait()
}

func knockCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "knock [hostname]",
		Short: "execute port knock",
		Args:  cobra.MinimumNArgs(1),
		RunE:  knockRun,
	}
	cmd.Flags().IntVarP(&knockPacketsNumber, "num", "n", 10, "number of packets to knock with")
	cmd.Flags().IntVarP(&knockPortHidden, "port-hidden", "P", 0, "after knocking start proxying to this hidden port")
	cmd.Flags().IntVarP(&knockTargetPort, "port-target", "T", 2222, "upd port to send knocks (default is 2222)")
	return cmd
}
