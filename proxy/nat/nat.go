//go:build linux
// +build linux

package nat

import (
	"context"

	"gitlab.atcatw.org/atca/community-edition/trojan-go/common"
	"gitlab.atcatw.org/atca/community-edition/trojan-go/config"
	"gitlab.atcatw.org/atca/community-edition/trojan-go/proxy"
	"gitlab.atcatw.org/atca/community-edition/trojan-go/proxy/client"
	"gitlab.atcatw.org/atca/community-edition/trojan-go/tunnel"
	"gitlab.atcatw.org/atca/community-edition/trojan-go/tunnel/tproxy"
)

const Name = "NAT"

func init() {
	proxy.RegisterProxyCreator(Name, func(ctx context.Context) (*proxy.Proxy, error) {
		cfg := config.FromContext(ctx, Name).(*client.Config)
		if cfg.Router.Enabled {
			return nil, common.NewError("router is not allowed in nat mode")
		}
		ctx, cancel := context.WithCancel(ctx)
		serverStack := []string{tproxy.Name}
		clientStack := client.GenerateClientTree(cfg.TransportPlugin.Enabled, cfg.Mux.Enabled, cfg.Websocket.Enabled, cfg.Shadowsocks.Enabled, false)
		c, err := proxy.CreateClientStack(ctx, clientStack)
		if err != nil {
			cancel()
			return nil, err
		}
		s, err := proxy.CreateServerStack(ctx, serverStack)
		if err != nil {
			cancel()
			return nil, err
		}
		return proxy.NewProxy(ctx, cancel, []tunnel.Server{s}, c), nil
	})
}

func init() {
	config.RegisterConfigCreator(Name, func() any {
		return new(client.Config)
	})
}
