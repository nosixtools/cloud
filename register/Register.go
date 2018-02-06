package register

import (
	"cloud/common/zk"
	"cloud/config"
	log "code.google.com/p/log4go"
	myzk "github.com/samuel/go-zookeeper/zk"
	"cloud/common/constant"
	"fmt"
	"time"
	"cloud/common/url"
	url2 "net/url"
)

type Register struct {
	RegisterConfig *config.RegisterConfig
	conn           *myzk.Conn
}

func NewRegister(RegisterConfig *config.RegisterConfig) (*Register, error) {
	register := &Register{RegisterConfig,nil}
	err := register.init()
	return  register, err
}

func (register *Register) init() error {
	var err error
	register.conn, err = zk.Connect(register.RegisterConfig.Address, register.RegisterConfig.Timeout)
	if err != nil {
		log.Error(err.Error())
	}
	time.Sleep(time.Second * 1)
	return err
}

func (register *Register) RegisterService(config *config.ProviderConfig) error {
	filePath := constant.RPC_PREFIX + constant.RPC_SEPERATE + config.ServiceName + constant.RPC_SEPERATE + config.Config.Group + constant.RPC_SEPERATE +  config.Config.Version
	exists, state, err := register.conn.Exists(filePath)
	if err != nil {
		log.Error(err.Error())
	}
	fmt.Println(exists, state)
	url := &url.Url{Protocol:"cloud", ServiceName:"HelloService", Config:config.Config}
	data := url.GetUrl()
	fmt.Println(url2.ParseQuery(data))
	if(!exists) {
		log.Info(filePath)
		err := zk.Create(register.conn, filePath)
		fmt.Println(err)
		if err != nil {
			log.Error(err.Error())
		}
	}
	err = zk.RegisterTemp(register.conn, filePath, []byte(data))
	fmt.Println(err)

	return nil
}
