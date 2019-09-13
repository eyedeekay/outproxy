package tgtun

import (
	"log"

	"github.com/9seconds/mtg/config"
	"github.com/9seconds/mtg/proxy"
	"github.com/eyedeekay/sam-forwarder/interface"
	"github.com/eyedeekay/sam-forwarder/tcp"
)

// TelegramOutProxy is a SAM-based SOCKS outproxy you connect to with a regular TCP
// tunnel
type TelegramOutProxy struct {
	*samforwarder.SAMForwarder
	*proxy.Proxy
	Conf *config.Config
	up   bool
}

func (f *TelegramOutProxy) GetType() string {
	return "mtproxy"
}

func (f *TelegramOutProxy) ServeParent() {
	log.Println("Starting eepsite server", f.Base32())
	if err := f.SAMForwarder.Serve(); err != nil {
		f.Cleanup()
	}
}

//Serve starts the SAM connection and and forwards the local host:port to i2p
func (f *TelegramOutProxy) Serve() error {
	go f.ServeParent()
	if f.Up() {
		log.Println("Starting HTTP/S proxy", f.Target())
		/*if err := http.ListenAndServe(f.Target(), f.Prox); err != nil {
			log.Fatal("ListenAndServe:", err)
		}*/
	}
	return nil
}

func (s *TelegramOutProxy) Load() (samtunnel.SAMTunnel, error) {
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

//NewTelegramOutProxyd makes a new SAM forwarder with default options, accepts host:port arguments
func NewTelegramOutProxyd(host, port string) (*TelegramOutProxy, error) {
	return NewTelegramOutProxydFromOptions(SetTelegramHost(host), SetTelegramPort(port))
}

//NewTelegramOutProxydFromOptions makes a new SAM forwarder with default options, accepts host:port arguments
func NewTelegramOutProxydFromOptions(opts ...func(*TelegramOutProxy) error) (*TelegramOutProxy, error) {
	var s TelegramOutProxy
	s.SAMForwarder = &samforwarder.SAMForwarder{}
	//s.Prox = &Proxy{}
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
	return l.(*TelegramOutProxy), nil
}
