package outproxy

import (
	"log"
	"net/http"

	"github.com/eyedeekay/sam-forwarder/interface"
	"github.com/eyedeekay/sam-forwarder/tcp"
)

// HttpOutProxy is a SAM-based SOCKS outproxy you connect to with a regular TCP
// tunnel
type HttpOutProxy struct {
    *samforwarder.SAMForwarder
	Prox      *Proxy
	up        bool
}

func (f *HttpOutProxy) GetType() string {
	return "outproxyhttp"
}

func (f *HttpOutProxy) ServeParent() {
	log.Println("Starting eepsite server", f.Base32())
	if err = f.SAMForwarder.Serve(); err != nil {
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

func (s *HttpOutProxy) Load() (samtunnel.SAMTunnel, error) {
	if !s.up {
		log.Println("Started putting tunnel up")
	}
	f, e := s.SAMForwarder.Load()
	if e != nil {
		return nil, e
	}
	s.SAMForwarder = f.(*samforwarder.SAMForwarder)

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
	s.SAMForwarder = &samforwarder.SAMForwarder{}
	s.Prox = &Proxy{}
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
	return l.(*HttpOutProxy), nil
}
