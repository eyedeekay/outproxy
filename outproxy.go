package outproxy

import (
	"log"
    "net"
    "golang.org/x/time/rate"

	"github.com/armon/go-socks5"
	"github.com/eyedeekay/sam-forwarder/config"
	"github.com/eyedeekay/sam-forwarder/interface"
	"github.com/eyedeekay/sam-forwarder/tcp"
	"github.com/eyedeekay/sam3/i2pkeys"
)

// OutProxy is a SAM-based SOCKS outproxy you connect to with a regular TCP
// tunnel
type OutProxy struct {
	Forwarder samtunnel.SAMTunnel
	Conf      *socks5.Config
	Socks     *socks5.Server
	up        bool
}

var err error

func (f *OutProxy) Config() *i2ptunconf.Conf {
	return f.Forwarder.Config()
}

func (f *OutProxy) ID() string {
	return f.Config().ID()
}

func (f *OutProxy) Keys() i2pkeys.I2PKeys {
	return f.Forwarder.Keys()
}

func (f *OutProxy) Cleanup() {
	f.Forwarder.Cleanup()
}

func (f *OutProxy) GetType() string {
	return f.Forwarder.GetType()
}

/*func (f *OutProxy) targetForPort443() string {
	if f.TargetForPort443 != "" {
		return "targetForPort.4443=" + f.TargetHost + ":" + f.TargetForPort443
	}
	return ""
}*/

func (f *OutProxy) Props() map[string]string {
	return f.Forwarder.Props()
}

func (f *OutProxy) Print() string {
	return f.Forwarder.Print()
}

func (f *OutProxy) Search(search string) string {
	return f.Forwarder.Search(search)
}

// Target returns the host:port of the local service you want to forward to i2p
func (f *OutProxy) Target() string {
	return f.Forwarder.Target()
}

//Base32 returns the base32 address where the local service is being forwarded
func (f *OutProxy) Base32() string {
	return f.Forwarder.Base32()
}

//Base32Readable returns the base32 address where the local service is being forwarded
func (f *OutProxy) Base32Readable() string {
	return f.Forwarder.Base32Readable()
}

//Base64 returns the base64 address where the local service is being forwarded
func (f *OutProxy) Base64() string {
	return f.Forwarder.Base64()
}

func (f *OutProxy) ServeParent() {
	log.Println("Starting eepsite server", f.Base32())
	if err = f.Forwarder.Serve(); err != nil {
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

func (f *OutProxy) Up() bool {
	if f.Forwarder.Up() {
		return true
	}
	return false
}

//Close shuts the whole thing down.
func (f *OutProxy) Close() error {
	return f.Forwarder.Close()
}

func (s *OutProxy) Load() (samtunnel.SAMTunnel, error) {
	if !s.up {
		log.Println("Started putting tunnel up")
	}
	f, e := s.Forwarder.Load()
	if e != nil {
		return nil, e
	}
	s.Forwarder = f.(*samforwarder.SAMForwarder)
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
	s.Forwarder = &samforwarder.SAMForwarder{}
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
	s.Forwarder.Config().SaveFile = true
	log.Println("Options loaded", s.Print())
	l, e := s.Load()
	if e != nil {
		return nil, e
	}
	return l.(*OutProxy), nil
}
