package context

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"syscall"
)

type Config struct {
	Debug  bool        `yaml:"debug"`
	Server *ServerConf `yaml:"server"`

	Fabric *FabricGWOption `yaml:"fabric"`

	Chaincode *ChaincodeOption `yaml:"chaincode"`

	// system initial admin
	// ca registrar
	Admin *AdminOption `json:"admin" yaml:"admin"`

	// couchdb options
	Couchdb *CouchdbOption `json:"couchdb" yaml:"couchdb"`

	// JWT
	JWT *JWTOption `json:"jwt" yaml:"jwt"`
}

type ServerConf struct {
	Host string `json:"host" yaml:"host"`
	Port string `json:"port" yaml:"port"`
}

func (s *ServerConf) GetServerAddr() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

type FabricGWOption struct {
	CCPath     string `json:"ccPath" yaml:"ccPath"`         // connection config file path
	WalletPath string `json:"walletPath" yaml:"walletPath"` // file type fabricwallet path
	OrgName    string `json:"orgName" yaml:"orgName"`
	MSPID      string `json:"mspId" yaml:"mspId"`
}

type ChaincodeOption struct {
	Channel   string `yaml:"channel"`
	Chaincode string `yaml:"chaincode"`
}

type AdminOption struct {
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}

type CouchdbOption struct {
	// HostPort is host:port format e.g 127.0.0.1:5984
	HostPort string `json:"hostPort" yaml:"hostPort"`
	User     string `json:"user" yaml:"user"`
	Passwd   string `json:"passwd" yaml:"passwd"`
	Protocol string `json:"protocol" yaml:"protocol"`
}

type JWTOption struct {
	Secret      string `json:"secret" yaml:"secret"`
	ExpireHours int    `json:"expireHours" yaml:"exporeHours"`
}

func (c *Config) LoadConf(filePath string) error {
	var err error
	viper.SetConfigFile(filePath)
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = viper.Unmarshal(c)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

// minPermission:
// 4 -> only check if can read
// 4 + 2 = 6 -> check if can read and write
func CheckPathExist(path string, permission int, desc string) error {
	// first check if exist
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			// Logger.Errorf("", "%s[%s] doesn't exist", desc, path)
		} else {
			// Logger.Errorf("", "%s[%s] failed, %s", desc, path, err.Error())
		}
		return err
	}

	// then check if can read or read/write
	var bit uint32 = syscall.O_RDWR
	if permission < 6 {
		bit = syscall.O_RDONLY
	}

	err := syscall.Access(path, bit)
	if err != nil {
		// Logger.Errorf("", "%s[%s] cannot access, %s", desc, path, err.Error())
		return err
	}

	return nil
}
