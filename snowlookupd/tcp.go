package snowlookupd

import (
	"net"
)

type tcpServer struct {
	ctx *Context
}

func (p *tcpServer) Handle(clientConn net.Conn) {

}
