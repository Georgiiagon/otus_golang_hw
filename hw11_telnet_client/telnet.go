package main

import (
	"context"
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type Client struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	conn    net.Conn
}

func (c *Client) Connect() (err error) {
	dialer := &net.Dialer{}
	ctx, _ := context.WithTimeout(context.Background(), c.timeout)
	c.conn, err = dialer.DialContext(ctx, "tcp", c.address)
	return err
}

func (c *Client) Close() error {
	err := c.conn.Close()
	return err
}

func (c *Client) Send() error {
	b := make([]byte, 0)
	_, err := c.in.Read(b)
	if err != nil {
		return err
	}
	_, err = c.conn.Write(b)
	return err
}

func (c *Client) Receive() error {
	return nil
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return Client{address: address, timeout: timeout, in: in, out: out}
}
