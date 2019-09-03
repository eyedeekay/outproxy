package outproxy

import (
	"fmt"
	"strconv"
)

//HttpOption is a HttpOutProxy Option
type HttpOption func(*HttpOutProxy) error

//SetHttpFilePath sets the path to save the config file at.
func SetHttpFilePath(s string) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		c.Forwarder.Config().FilePath = s
		return nil
	}
}

//SetHttpType sets the type of the forwarder server
func SetHttpType(s string) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		if s == "http" {
			c.Forwarder.Config().Type = s
			return nil
		} else {
			c.Forwarder.Config().Type = "server"
			return nil
		}
	}
}

//SetHttpSigType sets the type of the forwarder server
func SetHttpSigType(s string) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
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

//SetHttpSaveFile tells the router to save the tunnel's keys long-term
func SetHttpSaveFile(b bool) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		c.Forwarder.Config().SaveFile = b
		return nil
	}
}

//SetHttpHost sets the host of the service to forward
func SetHttpHost(s string) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		c.Forwarder.Config().TargetHost = s
		return nil
	}
}

//SetHttpPort sets the port of the service to forward
func SetHttpPort(s string) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
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

//SetHttpSAMHost sets the host of the HttpOutProxy's SAM bridge
func SetHttpSAMHost(s string) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		c.Forwarder.Config().SamHost = s
		return nil
	}
}

//SetHttpSAMPort sets the port of the HttpOutProxy's SAM bridge using a string
func SetHttpSAMPort(s string) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
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

//SetHttpName sets the host of the HttpOutProxy's SAM bridge
func SetHttpName(s string) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		c.Forwarder.Config().TunName = s
		return nil
	}
}

//SetHttpInLength sets the number of hops inbound
func SetHttpInLength(u int) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		if u < 7 && u >= 0 {
			c.Forwarder.Config().InLength = u
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetHttpOutLength sets the number of hops outbound
func SetHttpOutLength(u int) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		if u < 7 && u >= 0 {
			c.Forwarder.Config().OutLength = u
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel length")
	}
}

//SetHttpInVariance sets the variance of a number of hops inbound
func SetHttpInVariance(i int) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		if i < 7 && i > -7 {
			c.Forwarder.Config().InVariance = i
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetHttpOutVariance sets the variance of a number of hops outbound
func SetHttpOutVariance(i int) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		if i < 7 && i > -7 {
			c.Forwarder.Config().OutVariance = i
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel variance")
	}
}

//SetHttpInQuantity sets the inbound tunnel quantity
func SetHttpInQuantity(u int) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		if u <= 16 && u > 0 {
			c.Forwarder.Config().InQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel quantity")
	}
}

//SetHttpOutQuantity sets the outbound tunnel quantity
func SetHttpOutQuantity(u int) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		if u <= 16 && u > 0 {
			c.Forwarder.Config().OutQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel quantity")
	}
}

//SetHttpInBackups sets the inbound tunnel backups
func SetHttpInBackups(u int) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		if u < 6 && u >= 0 {
			c.Forwarder.Config().InBackupQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel backup quantity")
	}
}

//SetHttpOutBackups sets the inbound tunnel backups
func SetHttpOutBackups(u int) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		if u < 6 && u >= 0 {
			c.Forwarder.Config().OutBackupQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel backup quantity")
	}
}

//SetHttpEncrypt tells the router to use an encrypted leaseset
func SetHttpEncrypt(b bool) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		if b {
			c.Forwarder.Config().EncryptLeaseSet = true
			return nil
		}
		c.Forwarder.Config().EncryptLeaseSet = false
		return nil
	}
}

//SetHttpLeaseSetKey sets the host of the HttpOutProxy's SAM bridge
func SetHttpLeaseSetKey(s string) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		c.Forwarder.Config().LeaseSetKey = s
		return nil
	}
}

//SetHttpLeaseSetPrivateKey sets the host of the HttpOutProxy's SAM bridge
func SetHttpLeaseSetPrivateKey(s string) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		c.Forwarder.Config().LeaseSetPrivateKey = s
		return nil
	}
}

//SetHttpLeaseSetPrivateSigningKey sets the host of the HttpOutProxy's SAM bridge
func SetHttpLeaseSetPrivateSigningKey(s string) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		c.Forwarder.Config().LeaseSetPrivateSigningKey = s
		return nil
	}
}

//SetHttpMessageReliability sets the host of the HttpOutProxy's SAM bridge
func SetHttpMessageReliability(s string) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		c.Forwarder.Config().MessageReliability = s
		return nil
	}
}

//SetHttpAllowZeroIn tells the tunnel to accept zero-hop peers
func SetHttpAllowZeroIn(b bool) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		if b {
			c.Forwarder.Config().InAllowZeroHop = true
			return nil
		}
		c.Forwarder.Config().InAllowZeroHop = false
		return nil
	}
}

//SetHttpAllowZeroOut tells the tunnel to accept zero-hop peers
func SetHttpAllowZeroOut(b bool) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		if b {
			c.Forwarder.Config().OutAllowZeroHop = true
			return nil
		}
		c.Forwarder.Config().OutAllowZeroHop = false
		return nil
	}
}

//SetHttpCompress tells clients to use compression
func SetHttpCompress(b bool) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		if b {
			c.Forwarder.Config().UseCompression = true
			return nil
		}
		c.Forwarder.Config().UseCompression = false
		return nil
	}
}

//SetHttpFastRecieve tells clients to use compression
func SetHttpFastRecieve(b bool) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		if b {
			c.Forwarder.Config().FastRecieve = true
			return nil
		}
		c.Forwarder.Config().FastRecieve = false
		return nil
	}
}

//SetHttpReduceIdle tells the connection to reduce it's tunnels during extended idle time.
func SetHttpReduceIdle(b bool) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		if b {
			c.Forwarder.Config().ReduceIdle = true
			return nil
		}
		c.Forwarder.Config().ReduceIdle = false
		return nil
	}
}

//SetHttpReduceIdleTime sets the time to wait before reducing tunnels to idle levels
func SetHttpReduceIdleTime(u int) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		c.Forwarder.Config().ReduceIdleTime = 300000
		if u >= 6 {
			c.Forwarder.Config().ReduceIdleTime = (u * 60) * 1000
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in minutes) %v", u)
	}
}

//SetHttpReduceIdleTimeMs sets the time to wait before reducing tunnels to idle levels in milliseconds
func SetHttpReduceIdleTimeMs(u int) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		c.Forwarder.Config().ReduceIdleTime = 300000
		if u >= 300000 {
			c.Forwarder.Config().ReduceIdleTime = u
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetHttpReduceIdleQuantity sets minimum number of tunnels to reduce to during idle time
func SetHttpReduceIdleQuantity(u int) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		if u < 5 {
			c.Forwarder.Config().ReduceIdleQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid reduce tunnel quantity")
	}
}

//SetHttpCloseIdle tells the connection to close it's tunnels during extended idle time.
func SetHttpCloseIdle(b bool) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		if b {
			c.Forwarder.Config().CloseIdle = true
			return nil
		}
		c.Forwarder.Config().CloseIdle = false
		return nil
	}
}

//SetHttpCloseIdleTime sets the time to wait before closing tunnels to idle levels
func SetHttpCloseIdleTime(u int) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		c.Forwarder.Config().CloseIdleTime = 300000
		if u >= 6 {
			c.Forwarder.Config().CloseIdleTime = (u * 60) * 1000
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in minutes) %v", u)
	}
}

//SetHttpCloseIdleTimeMs sets the time to wait before closing tunnels to idle levels in milliseconds
func SetHttpCloseIdleTimeMs(u int) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		c.Forwarder.Config().CloseIdleTime = 300000
		if u >= 300000 {
			c.Forwarder.Config().CloseIdleTime = u
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetHttpAccessListType tells the system to treat the accessList as a whitelist
func SetHttpAccessListType(s string) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
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

//SetHttpAccessList tells the system to treat the accessList as a whitelist
func SetHttpAccessList(s []string) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		if len(s) > 0 {
			for _, a := range s {
				c.Forwarder.Config().AccessList = append(c.Forwarder.Config().AccessList, a)
			}
			return nil
		}
		return nil
	}
}

//SetHttpTargetForPort sets the port of the HttpOutProxy's SAM bridge using a string
/*func SetHttpTargetForPort443(s string) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
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

//SetHttpKeyFile sets
func SetHttpKeyFile(s string) func(*HttpOutProxy) error {
	return func(c *HttpOutProxy) error {
		c.Forwarder.Config().KeyFilePath = s
		return nil
	}
}
