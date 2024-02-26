//go:build linux
// +build linux

package tproxy

import (
	"context"
	"net"

	"gitlab.atcatw.org/atca/community-edition/trojan-go.git/common"
	"gitlab.atcatw.org/atca/community-edition/trojan-go.git/tunnel"
)

type Conn struct {
	net.Conn
	metadata *tunnel.Metadata
}

func (c *Conn) Metadata() *tunnel.Metadata {
	return c.metadata
}

type packetInfo struct {
	metadata *tunnel.Metadata
	payload  []byte
}

type PacketConn struct {
	net.PacketConn
	input  chan *packetInfo
	output chan *packetInfo
	src    net.Addr
	ctx    context.Context
	cancel context.CancelFunc
}

func (c *PacketConn) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	panic("implement me")
}

func (c *PacketConn) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	panic("implement me")
}

func (c *PacketConn) Close() error {
	c.cancel()
	return nil
}

func (c *PacketConn) WriteWithMetadata(p []byte, m *tunnel.Metadata) (int, error) {
	newP := make([]byte, len(p))
	select {
	case c.output <- &packetInfo{
		metadata: m,
		payload:  newP,
	}:
		return len(p), nil
	case <-c.ctx.Done():
		return 0, common.NewError("socks packet conn closed")
	}
}

func (c *PacketConn) ReadWithMetadata(p []byte) (int, *tunnel.Metadata, error) {
	select {
	case info := <-c.input:
		n := copy(p, info.payload)
		return n, info.metadata, nil
	case <-c.ctx.Done():
		return 0, nil, common.NewError("socks packet conn closed")
	}
}
