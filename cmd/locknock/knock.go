package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

var (
	knockPacketsNumber int
	knockKey           string
	knockPortProxy     int
)

func knockRun(cmd *cobra.Command, args []string) error {
	knockKey, err := knockPassword()
	if err != nil {
		return err
	}
	hostname := args[0]
	params := knockParams(knockKey, knockPacketsNumber)
	for _, port := range params.KnockPorts {
		address := fmt.Sprintf("%s:%d", hostname, port)
		udpaddr, err := net.ResolveUDPAddr("udp", address)
		if err != nil {
			return err
		}
		sock, err := net.DialUDP("udp", nil, udpaddr)
		if err != nil {
			return err
		}
		defer sock.Close()
		sock.Write([]byte(""))
	}
	if knockPortProxy != 0 {
		knockProxy(hostname)
	}
	return nil
}

func knockProxy(hostname string) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", hostname, knockPortProxy))
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
	cmd.Flags().IntVarP(&knockPacketsNumber, "num", "n", 100, "number of packets to knock with")
	cmd.Flags().IntVarP(&knockPortProxy, "port-proxy", "P", 0, "after knocking start proxying to this port")
	return cmd
}
