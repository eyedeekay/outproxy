package tgtun

import (
	"fmt"
	"strconv"
)

//TelegramOption is a TelegramOutProxy Option
type TelegramOption func(*TelegramOutProxy) error

//SetTelegramFilePath sets the path to save the config file at.
func SetTelegramFilePath(s string) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		c.Forwarder.Config().FilePath = s
		return nil
	}
}

//SetTelegramType sets the type of the forwarder server
func SetTelegramType(s string) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		if s == "http" {
			c.Forwarder.Config().Type = s
			return nil
		} else {
			c.Forwarder.Config().Type = "server"
			return nil
		}
	}
}

//SetTelegramSigType sets the type of the forwarder server
func SetTelegramSigType(s string) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		if s == "" {
			c.Forwarder.Config().SigType = ""
		} else if s == "DSA_SHA1" {
			c.Forwarder.Config().SigType = "DSA_SHA1"
		} else if s == "ECDSA_SHA256_P256" {
			c.Forwarder.Config().SigType = "ECDSA_SHA256_P256"
		} else if s == "ECDSA_SHA384_P384" {
			c.Forwarder.Config().SigType = "ECDSA_SHA384_P384"
		} else if s == "ECDSA_SHA512_P521" {
			c.Forwarder.Config().SigType = "ECDSA_SHA512_P521"
		} else if s == "EdDSA_SHA512_Ed25519" {
			c.Forwarder.Config().SigType = "EdDSA_SHA512_Ed25519"
		} else {
			c.Forwarder.Config().SigType = "EdDSA_SHA512_Ed25519"
		}
		return nil
	}
}

//SetTelegramSaveFile tells the router to save the tunnel's keys long-term
func SetTelegramSaveFile(b bool) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		c.Forwarder.Config().SaveFile = b
		return nil
	}
}

//SetTelegramHost sets the host of the service to forward
func SetTelegramHost(s string) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		c.Forwarder.Config().TargetHost = s
		return nil
	}
}

//SetTelegramPort sets the port of the service to forward
func SetTelegramPort(s string) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid TCP Server Target Port %s; non-number ", s)
		}
		if port < 65536 && port > -1 {
			c.Forwarder.Config().TargetPort = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetTelegramSAMHost sets the host of the TelegramOutProxy's SAM bridge
func SetTelegramSAMHost(s string) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		c.Forwarder.Config().SamHost = s
		return nil
	}
}

//SetTelegramSAMPort sets the port of the TelegramOutProxy's SAM bridge using a string
func SetTelegramSAMPort(s string) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid SAM Port %s; non-number", s)
		}
		if port < 65536 && port > -1 {
			c.Forwarder.Config().SamPort = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetTelegramName sets the host of the TelegramOutProxy's SAM bridge
func SetTelegramName(s string) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		c.Forwarder.Config().TunName = s
		return nil
	}
}

//SetTelegramInLength sets the number of hops inbound
func SetTelegramInLength(u int) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		if u < 7 && u >= 0 {
			c.Forwarder.Config().InLength = u
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetTelegramOutLength sets the number of hops outbound
func SetTelegramOutLength(u int) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		if u < 7 && u >= 0 {
			c.Forwarder.Config().OutLength = u
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel length")
	}
}

//SetTelegramInVariance sets the variance of a number of hops inbound
func SetTelegramInVariance(i int) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		if i < 7 && i > -7 {
			c.Forwarder.Config().InVariance = i
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetTelegramOutVariance sets the variance of a number of hops outbound
func SetTelegramOutVariance(i int) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		if i < 7 && i > -7 {
			c.Forwarder.Config().OutVariance = i
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel variance")
	}
}

//SetTelegramInQuantity sets the inbound tunnel quantity
func SetTelegramInQuantity(u int) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		if u <= 16 && u > 0 {
			c.Forwarder.Config().InQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel quantity")
	}
}

//SetTelegramOutQuantity sets the outbound tunnel quantity
func SetTelegramOutQuantity(u int) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		if u <= 16 && u > 0 {
			c.Forwarder.Config().OutQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel quantity")
	}
}

//SetTelegramInBackups sets the inbound tunnel backups
func SetTelegramInBackups(u int) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		if u < 6 && u >= 0 {
			c.Forwarder.Config().InBackupQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel backup quantity")
	}
}

//SetTelegramOutBackups sets the inbound tunnel backups
func SetTelegramOutBackups(u int) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		if u < 6 && u >= 0 {
			c.Forwarder.Config().OutBackupQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel backup quantity")
	}
}

//SetTelegramEncrypt tells the router to use an encrypted leaseset
func SetTelegramEncrypt(b bool) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		if b {
			c.Forwarder.Config().EncryptLeaseSet = true
			return nil
		}
		c.Forwarder.Config().EncryptLeaseSet = false
		return nil
	}
}

//SetTelegramLeaseSetKey sets the host of the TelegramOutProxy's SAM bridge
func SetTelegramLeaseSetKey(s string) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		c.Forwarder.Config().LeaseSetKey = s
		return nil
	}
}

//SetTelegramLeaseSetPrivateKey sets the host of the TelegramOutProxy's SAM bridge
func SetTelegramLeaseSetPrivateKey(s string) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		c.Forwarder.Config().LeaseSetPrivateKey = s
		return nil
	}
}

//SetTelegramLeaseSetPrivateSigningKey sets the host of the TelegramOutProxy's SAM bridge
func SetTelegramLeaseSetPrivateSigningKey(s string) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		c.Forwarder.Config().LeaseSetPrivateSigningKey = s
		return nil
	}
}

//SetTelegramMessageReliability sets the host of the TelegramOutProxy's SAM bridge
func SetTelegramMessageReliability(s string) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		c.Forwarder.Config().MessageReliability = s
		return nil
	}
}

//SetTelegramAllowZeroIn tells the tunnel to accept zero-hop peers
func SetTelegramAllowZeroIn(b bool) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		if b {
			c.Forwarder.Config().InAllowZeroHop = true
			return nil
		}
		c.Forwarder.Config().InAllowZeroHop = false
		return nil
	}
}

//SetTelegramAllowZeroOut tells the tunnel to accept zero-hop peers
func SetTelegramAllowZeroOut(b bool) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		if b {
			c.Forwarder.Config().OutAllowZeroHop = true
			return nil
		}
		c.Forwarder.Config().OutAllowZeroHop = false
		return nil
	}
}

//SetTelegramCompress tells clients to use compression
func SetTelegramCompress(b bool) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		if b {
			c.Forwarder.Config().UseCompression = true
			return nil
		}
		c.Forwarder.Config().UseCompression = false
		return nil
	}
}

//SetTelegramFastRecieve tells clients to use compression
func SetTelegramFastRecieve(b bool) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		if b {
			c.Forwarder.Config().FastRecieve = true
			return nil
		}
		c.Forwarder.Config().FastRecieve = false
		return nil
	}
}

//SetTelegramReduceIdle tells the connection to reduce it's tunnels during extended idle time.
func SetTelegramReduceIdle(b bool) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		if b {
			c.Forwarder.Config().ReduceIdle = true
			return nil
		}
		c.Forwarder.Config().ReduceIdle = false
		return nil
	}
}

//SetTelegramReduceIdleTime sets the time to wait before reducing tunnels to idle levels
func SetTelegramReduceIdleTime(u int) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		c.Forwarder.Config().ReduceIdleTime = 300000
		if u >= 6 {
			c.Forwarder.Config().ReduceIdleTime = (u * 60) * 1000
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in minutes) %v", u)
	}
}

//SetTelegramReduceIdleTimeMs sets the time to wait before reducing tunnels to idle levels in milliseconds
func SetTelegramReduceIdleTimeMs(u int) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		c.Forwarder.Config().ReduceIdleTime = 300000
		if u >= 300000 {
			c.Forwarder.Config().ReduceIdleTime = u
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetTelegramReduceIdleQuantity sets minimum number of tunnels to reduce to during idle time
func SetTelegramReduceIdleQuantity(u int) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		if u < 5 {
			c.Forwarder.Config().ReduceIdleQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid reduce tunnel quantity")
	}
}

//SetTelegramCloseIdle tells the connection to close it's tunnels during extended idle time.
func SetTelegramCloseIdle(b bool) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		if b {
			c.Forwarder.Config().CloseIdle = true
			return nil
		}
		c.Forwarder.Config().CloseIdle = false
		return nil
	}
}

//SetTelegramCloseIdleTime sets the time to wait before closing tunnels to idle levels
func SetTelegramCloseIdleTime(u int) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		c.Forwarder.Config().CloseIdleTime = 300000
		if u >= 6 {
			c.Forwarder.Config().CloseIdleTime = (u * 60) * 1000
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in minutes) %v", u)
	}
}

//SetTelegramCloseIdleTimeMs sets the time to wait before closing tunnels to idle levels in milliseconds
func SetTelegramCloseIdleTimeMs(u int) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		c.Forwarder.Config().CloseIdleTime = 300000
		if u >= 300000 {
			c.Forwarder.Config().CloseIdleTime = u
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetTelegramAccessListType tells the system to treat the accessList as a whitelist
func SetTelegramAccessListType(s string) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		if s == "whitelist" {
			c.Forwarder.Config().AccessListType = "whitelist"
			return nil
		} else if s == "blacklist" {
			c.Forwarder.Config().AccessListType = "blacklist"
			return nil
		} else if s == "none" {
			c.Forwarder.Config().AccessListType = ""
			return nil
		} else if s == "" {
			c.Forwarder.Config().AccessListType = ""
			return nil
		}
		return fmt.Errorf("Invalid Access list type(whitelist, blacklist, none)")
	}
}

//SetTelegramAccessList tells the system to treat the accessList as a whitelist
func SetTelegramAccessList(s []string) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		if len(s) > 0 {
			for _, a := range s {
				c.Forwarder.Config().AccessList = append(c.Forwarder.Config().AccessList, a)
			}
			return nil
		}
		return nil
	}
}

//SetTelegramTargetForPort sets the port of the TelegramOutProxy's SAM bridge using a string
/*func SetTelegramTargetForPort443(s string) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid Target Port %s; non-number ", s)
		}
		if port < 65536 && port > -1 {
			c.Forwarder.Config().TargetForPort443 = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}
*/

//SetTelegramKeyFile sets
func SetTelegramKeyFile(s string) func(*TelegramOutProxy) error {
	return func(c *TelegramOutProxy) error {
		c.Forwarder.Config().KeyFilePath = s
		return nil
	}
}
