package url

import (
	"cloud/config"
	"strconv"
)

type Url struct {
	Protocol string
	ServiceName string
	Config *config.Config
}

func (this *Url) GetUrl() string  {
	return this.Protocol+"://"+ this.Config.Host + ":" + strconv.Itoa(this.Config.Port) + "/" + this.Config.Group+"/" + this.ServiceName + "/" + this.Config.Version + "/" + this.Config.Weight
}