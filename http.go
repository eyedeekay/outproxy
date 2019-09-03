package outproxy

import (
	"log"
	"net/http"

	"github.com/eyedeekay/sam-forwarder/config"
	"github.com/eyedeekay/sam-forwarder/interface"
	"github.com/eyedeekay/sam-forwarder/tcp"
	"github.com/eyedeekay/sam3/i2pkeys"
)

// HttpOutProxy is a SAM-based SOCKS outproxy you connect to with a regular TCP
// tunnel
type HttpOutProxy struct {
	Forwarder samtunnel.SAMTunnel
	Prox      *Proxy
	up        bool
}

func (f *HttpOutProxy) Config() *i2ptunconf.Conf {
	return f.Forwarder.Config()
}

func (f *HttpOutProxy) ID() string {
	return f.Config().ID()
}

func (f *HttpOutProxy) Keys() i2pkeys.I2PKeys {
	return f.Forwarder.Keys()
}

func (f *HttpOutProxy) Cleanup() {
	f.Forwarder.Cleanup()
}

func (f *HttpOutProxy) GetType() string {
	return f.Forwarder.GetType()
}

/*func (f *HttpOutProxy) targetForPort443() string {
	if f.TargetForPort443 != "" {
		return "targetForPort.4443=" + f.TargetHost + ":" + f.TargetForPort443
	}
	return ""
}*/

func (f *HttpOutProxy) Props() map[string]string {
	return f.Forwarder.Props()
}

func (f *HttpOutProxy) Print() string {
	return f.Forwarder.Print()
}

func (f *HttpOutProxy) Search(search string) string {
	return f.Forwarder.Search(search)
}

// Target returns the host:port of the local service you want to forward to i2p
func (f *HttpOutProxy) Target() string {
	return f.Forwarder.Target()
}

//Base32 returns the base32 address where the local service is being forwarded
func (f *HttpOutProxy) Base32() string {
	return f.Forwarder.Base32()
}

//Base32Readable returns the base32 address where the local service is being forwarded
func (f *HttpOutProxy) Base32Readable() string {
	return f.Forwarder.Base32Readable()
}

//Base64 returns the base64 address where the local service is being forwarded
func (f *HttpOutProxy) Base64() string {
	return f.Forwarder.Base64()
}

func (f *HttpOutProxy) ServeParent() {
	log.Println("Starting eepsite server", f.Base32())
	if err = f.Forwarder.Serve(); err != nil {
		f.Cleanup()
	}
}

//Serve starts the SAM connection and and forwards the local host:port to i2p
func (f *HttpOutProxy) Serve() error {
	go f.ServeParent()
	if f.Up() {
		log.Println("Starting HTTP/S proxy", f.Target())
		if err := http.ListenAndServe(f.Target(), f.Prox); err != nil {
			log.Fatal("ListenAndServe:", err)
		}
	}
	return nil
}

func (f *HttpOutProxy) Up() bool {
	if f.Forwarder.Up() {
		return true
	}
	return false
}

//Close shuts the whole thing down.
func (f *HttpOutProxy) Close() error {
	return f.Forwarder.Close()
}

func (s *HttpOutProxy) Load() (samtunnel.SAMTunnel, error) {
	if !s.up {
		log.Println("Started putting tunnel up")
	}
	f, e := s.Forwarder.Load()
	if e != nil {
		return nil, e
	}
	s.Forwarder = f.(*samforwarder.SAMForwarder)

	s.up = true
	log.Println("Finished putting tunnel up")
	return s, nil
}

//NewHttpOutProxyd makes a new SAM forwarder with default options, accepts host:port arguments
func NewHttpOutProxyd(host, port string) (*HttpOutProxy, error) {
	return NewHttpOutProxydFromOptions(SetHttpHost(host), SetHttpPort(port))
}

//NewHttpOutProxydFromOptions makes a new SAM forwarder with default options, accepts host:port arguments
func NewHttpOutProxydFromOptions(opts ...func(*HttpOutProxy) error) (*HttpOutProxy, error) {
	var s HttpOutProxy
	s.Forwarder = &samforwarder.SAMForwarder{}
    s.Prox = &Proxy{}
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
	return l.(*HttpOutProxy), nil
}
