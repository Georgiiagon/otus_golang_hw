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

func (c Client) Connect() (err error) {
	dialer := &net.Dialer{}
	ctx, _ := context.WithTimeout(context.Background(), c.timeout)
	c.conn, err = dialer.DialContext(ctx, "tcp", c.address)

	return err
}

func (c Client) Close() error {
	return c.conn.Close()
}

func (c Client) Send() error {
	b, err := io.ReadAll(c.in)
	if err != nil {
		return err
	}
	_, err = c.conn.Write(b)
	if err != nil {
		return err
	}
	return err
}

func (c Client) Receive() error {
	b, err := io.ReadAll(c.conn)
	if err != nil {
		return err
	}
	c.out.Write(b)
	return nil
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	dialer := &net.Dialer{}
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	conn, _ := dialer.DialContext(ctx, "tcp", address)

	return Client{address: address, timeout: timeout, in: in, out: out, conn: conn}
}
