package main

import (
	"io"
	"log"
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

func (c *Client) transferMsg(in io.Reader, out io.Writer) error {
	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	return nil
}

func (c *Client) Connect() error {
	conn, err := net.DialTimeout("tcp", c.address, c.timeout)
	if err != nil {
		log.Printf("Connection error: %s", err)
		return err
	}

	c.conn = conn

	return nil
}

func (c *Client) Send() error { return c.transferMsg(c.in, c.conn) }

func (c *Client) Receive() error { return c.transferMsg(c.conn, c.out) }

func (c *Client) Close() error { return c.conn.Close() }

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &Client{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}
