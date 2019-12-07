package distributor

import "net"

type Distributor struct {
	types map[string]func(data interface{}, conn net.Conn)
}

func (d *Distributor) HandleType(typeConn string, handler func(data interface{}, conn net.Conn)) {
	d.types[typeConn] = handler
}

func (d *Distributor) Handle(conn net.Conn) {

}