package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"net"
	"sync"
	"time"
)

var Wg sync.WaitGroup

func init() {
	flag.DurationVar(&Timeout, "timeout", defaultTimeout, "timeout server connection")
	flag.StringVar(&Host, "host", "localhost", "host to connect")
	flag.StringVar(&Port, "port", "443", "port to connect")
}

var (
	defaultTimeout = 10 * time.Second
	Timeout        time.Duration
	Host           string
	Port           string
)

func main() {
	flag.Parse()

	address := net.JoinHostPort(Host, Port)
	in := &bytes.Buffer{}
	out := &bytes.Buffer{}

	client := NewTelnetClient(address, Timeout, ioutil.NopCloser(in), out)
	client.Connect()
	Wg.Wait()
}
