package register

import (
	"cloud/common/zk"
	"cloud/config"
	log "code.google.com/p/log4go"
	myzk "github.com/samuel/go-zookeeper/zk"
	"cloud/common/constant"
	"time"
	"strconv"
	"sync"
	"strings"
	"math/rand"
)

var (
	waitNodeDelaySecond   = time.Second * 1
	waitNodeDelay         = 1
)
type Register struct {
	RegisterConfig *config.RegisterConfig
	conn           *myzk.Conn
	lock      	   sync.Mutex
	addes          []string
}

func NewRegister(RegisterConfig *config.RegisterConfig) (*Register, error) {
	register := &Register{RegisterConfig:RegisterConfig}
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
	filePath := constant.RPC_PREFIX + constant.RPC_SEPERATE_ONE + config.ServiceName + constant.RPC_SEPERATE_ONE + config.Config.Group + constant.RPC_SEPERATE_ONE +  config.Config.Version
	nodeInfo := config.Config.Host + constant.RPC_SEPERATE_TWO + strconv.Itoa(config.Config.Port) + constant.RPC_SEPERATE_TWO + config.Config.Weight
	path := filePath + constant.RPC_SEPERATE_ONE + nodeInfo
	exists, _, err := register.conn.Exists(filePath)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	if(!exists) {
		err := zk.Create(register.conn, filePath)
		if err != nil {
			log.Error(err.Error())
			return err
		}

	}
	err = zk.CreateEphemeral(register.conn, path)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (register *Register) DiscoverServices(config *config.ReferenceConfig) error {
	filePath := constant.RPC_PREFIX + constant.RPC_SEPERATE_ONE + config.ServiceName + constant.RPC_SEPERATE_ONE + config.Config.Group + constant.RPC_SEPERATE_ONE +  config.Config.Version
	go func(){
		for {
			nodes, watch, err := zk.GetNodesW(register.conn, filePath)
			if err == zk.ErrNodeNotExist {
				log.Warn("zk don't have node \"%s\", retry in %d second", filePath, waitNodeDelay)
				time.Sleep(waitNodeDelaySecond)
				continue
			} else if err == zk.ErrNoChild {
				log.Warn("zk don't have any children in \"%s\", retry in %d second", filePath, waitNodeDelay)
				time.Sleep(waitNodeDelaySecond)
				continue
			} else if err != nil {
				log.Error("getNodes error(%v), retry in %d second", err, waitNodeDelay)
				time.Sleep(waitNodeDelaySecond)
				continue
			}
			register.lock.Lock()
			register.addes = nodes
			register.lock.Unlock()
			event := <-watch
			log.Info("zk path: \"%s\" receive a event %v", filePath, event)
		}
	}()
	time.Sleep(time.Second * 2)
	return nil
}

func (register *Register) Seletor() string {
	if register.addes != nil && len(register.addes) > 0 {
		index := rand.Intn(len(register.addes))
		service := register.addes[index]
		return service[:strings.LastIndex(service, ":")]
	} else {
		return ""
	}
}

