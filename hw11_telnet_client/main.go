package main

import (
	"flag"
	"net"
	"time"
)

var (
	defaultTimeout = 10 * time.Second
	Timeout        time.Duration
	Host           string
	Port           string
)

func main() {
	flag.DurationVar(&Timeout, "timeout", defaultTimeout, "timeout server connection")
	flag.StringVar(&Host, "host", "localhost", "host to connect")
	flag.StringVar(&Port, "port", "443", "port to connect")

	address := net.JoinHostPort(Host, Port)
	client := NewTelnetClient(address, Timeout)
}
