package gateway

import (
	"fmt"
	"github.com/nightwolf93/gostream-server/rtmp"
	log "github.com/sirupsen/logrus"
	"net"
)

// Gateway represent a streaming gateway
type Gateway struct {
	config GatewayConfig
}

// NewGateway return a new Gateway
func NewGateway(config GatewayConfig) *Gateway {
	gateway := &Gateway{
		config: config,
	}
	return gateway
}

func (gateway *Gateway) Listen() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", gateway.config.Port))
	if err != nil {
		log.Fatalf("gateway can't listen: %v", err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Errorf("can't accept gateway client: %v", err)
		}
		log.Infof("new connection incoming on the gateway from %s", conn.RemoteAddr().String())
		srvRmtpConn := rtmp.NewSrcRtmpConn(&conn)
		_ = srvRmtpConn
	}
}
