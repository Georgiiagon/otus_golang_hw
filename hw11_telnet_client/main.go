package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/pflag"
)

func init() {
	pflag.DurationVar(&timeout, "timeout", defaultTimeout, "timeout server connection")
	pflag.Parse()
}

var (
	defaultTimeout = 10 * time.Second
	timeout        time.Duration
	host           string
	port           string
)

func main() {
	if len(pflag.Args()) < 2 {
		log.Fatal("should be 2 arguments")
	}
	host = pflag.Arg(0)
	port = pflag.Arg(1)
	address := net.JoinHostPort(host, port)

	client := NewTelnetClient(address, timeout, os.Stdin, os.Stdout)
	if err := client.Connect(); err != nil {
		log.Fatal(err)
	}
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGKILL)
	ctx, cancel := context.WithTimeout(ctx, timeout)

	defer client.Close()

	go func() {
		defer cancel()
		err := client.Send()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	go func() {
		defer cancel()
		err := client.Receive()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	<-ctx.Done()
}
