package snowlookupd

import (
	"fmt"
	"log"
	"net"
	"os"
	"snow/internal/protocol"
	"snow/internal/util"
	"snow/internal/version"
	"sync"
)

type SnowLookupd struct {
	sync.RWMutex
	opts         *Options
	tcpListener  net.Listener
	httpListener net.Listener
	waitGroup    util.WaitGroupWrapper
	DB           *RegistrationDB
}

func New(opts *Options) *SnowLookupd {
	if opts.Logger == nil {
		opts.Logger = log.New(os.Stderr, opts.LogPrefix, log.Ldate|log.Ltime|log.Lmicroseconds)
	}
	n := &SnowLookupd{
		opts: opts,
		DB:   NewRegistrationDB(),
	}
	n.logf(version.String("snowlookupd"))
	return n
}

func (l *SnowLookupd) logf(f string, args ...interface{}) {
	l.opts.Logger.Output(2, fmt.Sprintf(f, args...))
}

func (l *SnowLookupd) Main() {
	ctx := &Context{l}
	tcpListener, err := net.Listen("tcp", l.opts.TCPAddress)
	if err != nil {
		l.logf("FATAL:listen (%s) failed - %s", l.opts.TCPAddress, err)
		os.Exit(1)
	}
	l.Lock()
	l.tcpListener = tcpListener
	l.Unlock()
	tcpServer := &tcpServer{ctx: ctx}
	l.waitGroup.Wrap(func() {
		protocol.TCPServer(tcpListener, tcpServer, l.opts.Logger)
	})

	/*
		httpListener, err := net.Listen("tcp", l.opts.HTTPAddress)
		if err != nil {
			l.logf("FATAL:listen (%s) failed - %s", l.opts.HTTPAddress, err)
		}
		l.Lock()
		l.httpListener = httpListener
		l.Unlock()
		httpServer := newHTTPServer(ctx)
		l.waitGroup.Wrap(func() {
			http_api.Server(httpListener, httpServer, "http", l.opts.Logger)
		})
	*/
}

func (l *SnowLookupd) RealTCPAddr() *net.TCPAddr {
	l.RLock()
	defer l.RUnlock()
	return l.tcpListener.Addr().(*net.TCPAddr)
}

func (l *SnowLookupd) RealHTTPAddr() *net.TCPAddr {
	l.RLock()
	defer l.RUnlock()
	return l.httpListener.Addr().(*net.TCPAddr)
}

func (l *SnowLookupd) Exit() {
	if l.tcpListener != nil {
		l.tcpListener.Close()
	}
	if l.httpListener != nil {
		l.httpListener.Close()
	}
	l.waitGroup.Wait()
}
