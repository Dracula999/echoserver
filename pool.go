package main

import (
	"net"
)

struct connPool interface{
	Get() (net.Conn, error)
	Size() int
	Close() 
}
