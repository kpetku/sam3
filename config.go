package sam3

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"

	"github.com/eyedeekay/sam3/i2pkeys"
)

// I2PConfig is a struct which manages I2P configuration options
type I2PConfig struct {
	SamHost string
	SamPort string
	TunName string

	SamMin string
	SamMax string

	Fromport string
	Toport   string

	Style   string
	TunType string

	DestinationKeys i2pkeys.I2PKeys

	SigType                   string
	EncryptLeaseSet           string
	LeaseSetKey               string
	LeaseSetPrivateKey        string
	LeaseSetPrivateSigningKey string
	LeaseSetKeys              i2pkeys.I2PKeys
	InAllowZeroHop            string
	OutAllowZeroHop           string
	InLength                  string
	OutLength                 string
	InQuantity                string
	OutQuantity               string
	InVariance                string
	OutVariance               string
	InBackupQuantity          string
	OutBackupQuantity         string
	FastRecieve               string
	UseCompression            string
	MessageReliability        string
	CloseIdle                 string
	CloseIdleTime             string
	ReduceIdle                string
	ReduceIdleTime            string
	ReduceIdleQuantity        string
	//Streaming Library options
	AccessListType string
	AccessList     []string
}

func (f *I2PConfig) Sam() string {
	host := "127.0.0.1"
	port := "7656"
	if f.SamHost != "" {
		host = f.SamHost
	}
	if f.SamPort != "" {
		port = f.SamPort
	}
	return host + ":" + port
}

func (f *I2PConfig) SetSAMAddress(addr string) {
	hp := strings.Split(addr, ":")
	if len(hp) == 1 {
		f.SamHost = hp[0]
	} else if len(hp) == 2 {
		f.SamPort = hp[1]
		f.SamHost = hp[0]
	}
	f.SamPort = "7656"
	f.SamHost = "127.0.0.1"
}

func (f *I2PConfig) ID() string {
	if f.TunName == "" {
		b := make([]byte, 12)
		for i := range b {
			b[i] = "abcdefghijklmnopqrstuvwxyz"[rand.Intn(len("abcdefghijklmnopqrstuvwxyz"))]
		}
		f.TunName = string(b)
	}
	return " ID=" + f.TunName + " "
}

func (f *I2PConfig) Leasesetsettings() (string, string, string) {
	var r, s, t string
	if f.LeaseSetKey != "" {
		r = " i2cp.leaseSetKey=" + f.LeaseSetKey + " "
	}
	if f.LeaseSetPrivateKey != "" {
		s = " i2cp.leaseSetPrivateKey=" + f.LeaseSetPrivateKey + " "
	}
	if f.LeaseSetPrivateSigningKey != "" {
		t = " i2cp.leaseSetPrivateSigningKey=" + f.LeaseSetPrivateSigningKey + " "
	}
	return r, s, t
}

func (f *I2PConfig) FromPort() string {
	if f.samMax() < 3.1 {
		return ""
	}
	if f.Fromport != "0" {
		return " FROM_PORT=" + f.Fromport + " "
	}
	return ""
}

func (f *I2PConfig) ToPort() string {
	if f.samMax() < 3.1 {
		return ""
	}
	if f.Toport != "0" {
		return " TO_PORT=" + f.Toport + " "
	}
	return ""
}

func (f *I2PConfig) SessionStyle() string {
	if f.Style != "" {
		return " STYLE=" + f.Style + " "
	}
	return " STYLE=STREAM "
}

func (f *I2PConfig) samMax() float64 {
	i, err := strconv.Atoi(f.SamMax)
	if err != nil {
		return 3.1
	}
	return float64(i)
}

func (f *I2PConfig) MinSAM() string {
	if f.SamMin == "" {
		return "3.0"
	}
	return f.SamMin
}

func (f *I2PConfig) MaxSAM() string {
	if f.SamMax == "" {
		return "3.1"
	}
	return f.SamMax
}

func (f *I2PConfig) DestinationKey() string {
	if &f.DestinationKeys != nil {
		return " DESTINATION=" + f.DestinationKeys.String() + " "
	}
	return " DESTINATION=TRANSIENT "
}

func (f *I2PConfig) SignatureType() string {
	if f.samMax() < 3.1 {
		return ""
	}
	if f.SigType != "" {
		return " SIGNATURE_TYPE=" + f.SigType + " "
	}
	return ""
}

func (f *I2PConfig) EncryptLease() string {
	if f.EncryptLeaseSet == "true" {
		return " i2cp.encryptLeaseSet=true "
	}
	return ""
}

func (f *I2PConfig) Reliability() string {
	if f.MessageReliability != "" {
		return " i2cp.messageReliability=" + f.MessageReliability + " "
	}
	return ""
}

func (f *I2PConfig) Reduce() string {
	if f.ReduceIdle == "true" {
		return "i2cp.reduceOnIdle=" + f.ReduceIdle + "i2cp.reduceIdleTime=" + f.ReduceIdleTime + "i2cp.reduceQuantity=" + f.ReduceIdleQuantity
	}
	return ""
}

func (f *I2PConfig) Close() string {
	if f.CloseIdle == "true" {
		return "i2cp.closeOnIdle=" + f.CloseIdle + "i2cp.closeIdleTime=" + f.CloseIdleTime
	}
	return ""
}

func (f *I2PConfig) DoZero() string {
	r := ""
	if f.InAllowZeroHop == "true" {
		r += " inbound.allowZeroHop=" + f.InAllowZeroHop + " "
	}
	if f.OutAllowZeroHop == "true" {
		r += " outbound.allowZeroHop= " + f.OutAllowZeroHop + " "
	}
	if f.FastRecieve == "true" {
		r += " " + f.FastRecieve + " "
	}
	return r
}
func (f *I2PConfig) Print() []string {
	lsk, lspk, lspsk := f.Leasesetsettings()
	return []string{
		//f.targetForPort443(),
		"inbound.length=" + f.InLength,
		"outbound.length=" + f.OutLength,
		"inbound.lengthVariance=" + f.InVariance,
		"outbound.lengthVariance=" + f.OutVariance,
		"inbound.backupQuantity=" + f.InBackupQuantity,
		"outbound.backupQuantity=" + f.OutBackupQuantity,
		"inbound.quantity=" + f.InQuantity,
		"outbound.quantity=" + f.OutQuantity,
		f.DoZero(),
		//"i2cp.fastRecieve=" + f.FastRecieve,
		"i2cp.gzip=" + f.UseCompression,
		f.Reduce(),
		f.Close(),
		f.Reliability(),
		f.EncryptLease(),
		lsk, lspk, lspsk,
		f.Accesslisttype(),
		f.Accesslist(),
	}
}

func (f *I2PConfig) Accesslisttype() string {
	if f.AccessListType == "whitelist" {
		return "i2cp.enableAccessList=true"
	} else if f.AccessListType == "blacklist" {
		return "i2cp.enableBlackList=true"
	} else if f.AccessListType == "none" {
		return ""
	}
	return ""
}

func (f *I2PConfig) Accesslist() string {
	if f.AccessListType != "" && len(f.AccessList) > 0 {
		r := ""
		for _, s := range f.AccessList {
			r += s + ","
		}
		return "i2cp.accessList=" + strings.TrimSuffix(r, ",")
	}
	return ""
}

func NewConfig(opts ...func(*I2PConfig) error) (*I2PConfig, error) {
	var config I2PConfig
	config.SamHost = "127.0.0.1"
	config.SamPort = "7656"
	config.SamMin = "3.0"
	config.SamMax = "3.2"
	config.TunName = ""
	config.TunType = "server"
	config.Style = "STREAM"
	config.InLength = "3"
	config.OutLength = "3"
	config.InQuantity = "2"
	config.OutQuantity = "2"
	config.InVariance = "1"
	config.OutVariance = "1"
	config.InBackupQuantity = "3"
	config.OutBackupQuantity = "3"
	config.InAllowZeroHop = "false"
	config.OutAllowZeroHop = "false"
	config.EncryptLeaseSet = "false"
	config.LeaseSetKey = ""
	config.LeaseSetPrivateKey = ""
	config.LeaseSetPrivateSigningKey = ""
	config.FastRecieve = "false"
	config.UseCompression = "true"
	config.ReduceIdle = "false"
	config.ReduceIdleTime = "15"
	config.ReduceIdleQuantity = "4"
	config.CloseIdle = "false"
	config.CloseIdleTime = "300000"
	config.MessageReliability = "none"
	for _, o := range opts {
		if err := o(&config); err != nil {
			return nil, err
		}
	}
	return &config, nil
}

// options map
type Options map[string]string

// obtain sam options as list of strings
func (opts Options) AsList() (ls []string) {
	for k, v := range opts {
		ls = append(ls, fmt.Sprintf("%s=%s", k, v))
	}
	return
}

// Config is the config type for the sam connector api for i2p which allows applications to 'speak' with i2p
type Config struct {
	Addr    string
	Opts    Options
	Session string
	Keyfile string
}

// create new sam connector from config with a stream session
func (cfg *Config) StreamSession() (session *StreamSession, err error) {
	// connect
	var s *SAM
	s, err = NewSAM(cfg.Addr)
	if err == nil {
		// ensure keys exist
		var keys i2pkeys.I2PKeys
		keys, err = s.EnsureKeyfile(cfg.Keyfile)
		if err == nil {
			// create session
			session, err = s.NewStreamSession(cfg.Session, keys, cfg.Opts.AsList())
		}
	}
	return
}

// create new sam datagram session from config
func (cfg *Config) DatagramSession() (session *DatagramSession, err error) {
	// connect
	var s *SAM
	s, err = NewSAM(cfg.Addr)
	if err == nil {
		// ensure keys exist
		var keys i2pkeys.I2PKeys
		keys, err = s.EnsureKeyfile(cfg.Keyfile)
		if err == nil {
			// determine udp port
			var portstr string
			_, portstr, err = net.SplitHostPort(cfg.Addr)
			if err == nil {
				var port int
				port, err = strconv.Atoi(portstr)
				if err == nil && port > 0 {
					// udp port is 1 lower
					port--
					// create session
					session, err = s.NewDatagramSession(cfg.Session, keys, cfg.Opts.AsList(), port)
				}
			}
		}
	}
	return
}
