package outproxy

import (
	"context"
	"github.com/armon/go-socks5"
	"golang.org/x/time/rate"
	"net"
)

type SocksRuleSet struct {
	Ports   []int
	Domains []string
	IPs     []net.IP
	Rate    map[string]*rate.Limiter
	Limit   float64
	Burst   int

	Default bool
}

func (s *SocksRuleSet) Allow(ctx context.Context, req *socks5.Request) (context.Context, bool) {
	if s.Limit > 0 && s.Rate[req.DestAddr.IP.String()] == nil {
		s.Rate[req.DestAddr.IP.String()] = rate.NewLimiter(rate.Limit(s.Limit), s.Burst)
	}
	if s.Limit > 0 && s.Rate[req.DestAddr.IP.String()].Allow() == false {
		return ctx, false
	}
	for _, v := range s.Ports {
		if v == req.DestAddr.Port {
			return ctx, !s.Default
		}
	}
	for _, v := range s.Domains {
		if v == req.DestAddr.FQDN {
			return ctx, !s.Default
		}
	}
	for _, v := range s.IPs {
		if v.String() == req.DestAddr.IP.String() {
			return ctx, !s.Default
		}
	}
	return ctx, s.Default
}
