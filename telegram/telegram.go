package tgtun


import (
	"log"
//	"net/http"

	"github.com/eyedeekay/sam-forwarder/config"
	"github.com/eyedeekay/sam-forwarder/interface"
	"github.com/eyedeekay/sam-forwarder/tcp"
	"github.com/eyedeekay/sam3/i2pkeys"
)

// TelegramOutProxy is a SAM-based SOCKS outproxy you connect to with a regular TCP
// tunnel
type TelegramOutProxy struct {
	Forwarder samtunnel.SAMTunnel
	up        bool
}

func (f *TelegramOutProxy) Config() *i2ptunconf.Conf {
	return f.Forwarder.Config()
}

func (f *TelegramOutProxy) ID() string {
	return f.Config().ID()
}

func (f *TelegramOutProxy) Keys() i2pkeys.I2PKeys {
	return f.Forwarder.Keys()
}

func (f *TelegramOutProxy) Cleanup() {
	f.Forwarder.Cleanup()
}

func (f *TelegramOutProxy) GetType() string {
	return f.Forwarder.GetType()
}

/*func (f *TelegramOutProxy) targetForPort443() string {
	if f.TargetForPort443 != "" {
		return "targetForPort.4443=" + f.TargetHost + ":" + f.TargetForPort443
	}
	return ""
}*/

func (f *TelegramOutProxy) Props() map[string]string {
	return f.Forwarder.Props()
}

func (f *TelegramOutProxy) Print() string {
	return f.Forwarder.Print()
}

func (f *TelegramOutProxy) Search(search string) string {
	return f.Forwarder.Search(search)
}

// Target returns the host:port of the local service you want to forward to i2p
func (f *TelegramOutProxy) Target() string {
	return f.Forwarder.Target()
}

//Base32 returns the base32 address where the local service is being forwarded
func (f *TelegramOutProxy) Base32() string {
	return f.Forwarder.Base32()
}

//Base32Readable returns the base32 address where the local service is being forwarded
func (f *TelegramOutProxy) Base32Readable() string {
	return f.Forwarder.Base32Readable()
}

//Base64 returns the base64 address where the local service is being forwarded
func (f *TelegramOutProxy) Base64() string {
	return f.Forwarder.Base64()
}

func (f *TelegramOutProxy) ServeParent() {
	log.Println("Starting eepsite server", f.Base32())
	if err := f.Forwarder.Serve(); err != nil {
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

func (f *TelegramOutProxy) Up() bool {
	if f.Forwarder.Up() {
		return true
	}
	return false
}

//Close shuts the whole thing down.
func (f *TelegramOutProxy) Close() error {
	return f.Forwarder.Close()
}

func (s *TelegramOutProxy) Load() (samtunnel.SAMTunnel, error) {
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

//NewTelegramOutProxyd makes a new SAM forwarder with default options, accepts host:port arguments
func NewTelegramOutProxyd(host, port string) (*TelegramOutProxy, error) {
	return NewTelegramOutProxydFromOptions(SetTelegramHost(host), SetTelegramPort(port))
}

//NewTelegramOutProxydFromOptions makes a new SAM forwarder with default options, accepts host:port arguments
func NewTelegramOutProxydFromOptions(opts ...func(*TelegramOutProxy) error) (*TelegramOutProxy, error) {
	var s TelegramOutProxy
	s.Forwarder = &samforwarder.SAMForwarder{}
	//s.Prox = &Proxy{}
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
	return l.(*TelegramOutProxy), nil
}
