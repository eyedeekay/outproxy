package outproxy

import (
	"log"
    "net"
    "golang.org/x/time/rate"

	"github.com/armon/go-socks5"
	"github.com/eyedeekay/sam-forwarder/interface"
	"github.com/eyedeekay/sam-forwarder/tcp"
)

// OutProxy is a SAM-based SOCKS outproxy you connect to with a regular TCP
// tunnel
type OutProxy struct {
    *samforwarder.SAMForwarder
	Conf      *socks5.Config
	Socks     *socks5.Server
	up        bool
}

var err error

func (f *OutProxy) GetType() string {
	return "outproxy"
}

func (f *OutProxy) ServeParent() {
	log.Println("Starting eepsite server", f.Base32())
	if err = f.SAMForwarder.Serve(); err != nil {
		f.Cleanup()
	}
}

//Serve starts the SAM connection and and forwards the local host:port to i2p
func (f *OutProxy) Serve() error {
	go f.ServeParent()
	if f.Up() {
		log.Println("Starting SOCKS proxy", f.Target())
		if err := f.Socks.ListenAndServe("tcp", f.Target()); err != nil {
			panic(err)
		}
	}
	return nil
}

//Close shuts the whole thing down.
func (f *OutProxy) Close() error {
	return f.SAMForwarder.Close()
}

func (s *OutProxy) Load() (samtunnel.SAMTunnel, error) {
	if !s.up {
		log.Println("Started putting tunnel up")
	}
	f, e := s.SAMForwarder.Load()
	if e != nil {
		return nil, e
	}
	s.SAMForwarder = f.(*samforwarder.SAMForwarder)
	s.Socks, err = socks5.New(s.Conf)
	if err != nil {
		return nil, err
	}
	s.up = true
	log.Println("Finished putting tunnel up")
	return s, nil
}

//NewOutProxyd makes a new SAM forwarder with default options, accepts host:port arguments
func NewOutProxy(host, port string) (*OutProxy, error) {
	return NewOutProxyFromOptions(SetHost(host), SetPort(port))
}

//NewOutProxydFromOptions makes a new SAM forwarder with default options, accepts host:port arguments
func NewOutProxyFromOptions(opts ...func(*OutProxy) error) (*OutProxy, error) {
	var s OutProxy
	s.SAMForwarder = &samforwarder.SAMForwarder{}
	s.Conf = &socks5.Config{
		Rules: &SocksRuleSet{
			Ports:     []int{80, 443},
			Domains:   []string{""},
			IPs:       []net.IP{},
			Limit:     -1,
            Rate:      make(map[string]*rate.Limiter),
			Burst:     -1,
			Default:   true,
		},
	}
	log.Println("Initializing outproxy")
	for _, o := range opts {
		if err := o(&s); err != nil {
			return nil, err
		}
	}
	s.SAMForwarder.Config().SaveFile = true
	log.Println("Options loaded", s.Print())
	l, e := s.Load()
	if e != nil {
		return nil, e
	}
	return l.(*OutProxy), nil
}
