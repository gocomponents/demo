package context

import (
	"context"
	"fmt"
	"net"
	"time"
)

func Cancel()  {
	ctx, cancel := context.WithCancel(context.Background())
	go handleCtx(ctx)
	go func(cancel context.CancelFunc) {
		time.Sleep(1*time.Minute)
		cancel()
	}(cancel)
}

func handleCtx(ctx context.Context)  {
	ln, err := net.Listen("tcp", ":8080")
	if nil!=err {
		fmt.Println(err)
		return
	}
	defer ln.Close()
	connChan:=make(chan net.Conn)
	go func(ln net.Listener,connChan chan net.Conn) {
		for {
			conn, err := ln.Accept()
			if nil!=err {
				fmt.Println(err)
				break
			}
			connChan<-conn
		}
	}(ln,connChan)
	for {
		select {
		case conn := <-connChan:
			go handleConn(ctx,conn)
		case <-ctx.Done():
			return
		}
	}
}

func handleConn(ctx context.Context,conn net.Conn)  {
	defer conn.Close()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			buf := make([]byte, 100)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println(err)
				return
			}
			msg:=fmt.Sprintf("hello,%s",string(buf[0:n]))
			conn.Write([]byte(msg))
		}
	}
}
