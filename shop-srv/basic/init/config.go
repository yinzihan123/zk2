package init

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"shop-srv/basic/globals"
	"strings"
)

func InitConfig() {
	viper.SetConfigFile("../config/config.yaml")
	viper.ReadInConfig()
	viper.Unmarshal(&globals.AppConfig)
	fmt.Println(globals.AppConfig)
	nacosConfig := globals.AppConfig.Nacos
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: nacosConfig.Host,
			Port:   uint64(nacosConfig.Port),
		},
	}
	// 客户端配置
	clientConfig := constant.ClientConfig{
		NamespaceId:         nacosConfig.NameSpace, // 如果不需要命名空间，可以留空
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// 创建配置客户端
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(err)
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: nacosConfig.DataId,
		Group:  nacosConfig.Group,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
	viper.Reset()
	viper.SetConfigType("yaml")
	err = viper.ReadConfig(strings.NewReader(content))
	if err != nil {
		panic(err)
	}
	viper.Unmarshal(&globals.AppConfig)
	fmt.Println(globals.AppConfig)
}
