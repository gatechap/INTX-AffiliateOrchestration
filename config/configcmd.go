package config

import (
	"fmt"

	"th.truecorp.it.dsm.intcom/affiliateorchestration/errormsg"

	"github.com/spf13/viper"
)

const CURRENT_FILE = "configcmd"

func GetService(serviceName string) (svc Service) {

	appConfig, _ := LoadConfig()

	for _, sv := range appConfig.Service {
		if sv.Name == serviceName {
			svc = sv
			return
		}
	}
	return
}

func LoadConfig() (*Config, *errormsg.ErrorHandlerInfo) {
	const CURRENT_FUNCTION = "LoadConfig"

	errHanlerInfo := errormsg.ErrorHandlerInfo{
		ErrorApplication: CURRENT_APPLICATION,
		ErrorModule:      CURRENT_MODULE,
		ErrorFile:        CURRENT_FILE,
		ErrorFunction:    CURRENT_FUNCTION,
	}
	// Set the file name of the configurations file

	viper.SetConfigName("/configmap/application")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	viper.SetConfigType("yml")

	var appConfig Config

	if err := viper.ReadInConfig(); err != nil {
		// fmt.Printf("Error reading config file, %s", err)

		errHanlerInfo.Error = err
		return getDefaultConfigYML(), &errHanlerInfo
	}

	err := viper.Unmarshal(&appConfig)
	if err != nil {
		// fmt.Printf("Unable to decode into struct, %v", err)

		errHanlerInfo.Error = err
		return getDefaultConfigYML(), &errHanlerInfo
	}

	return &appConfig, nil
}

func getDefaultConfigYML() *Config {
	const CURRENT_FUNCTION = "LoadConfig"

	errHanlerInfo := errormsg.ErrorHandlerInfo{
		ErrorApplication: CURRENT_APPLICATION,
		ErrorModule:      CURRENT_MODULE,
		ErrorFile:        CURRENT_FILE,
		ErrorFunction:    CURRENT_FUNCTION,
	}
	// Set the file name of the configurations file

	viper.SetConfigName("./config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	viper.SetConfigType("yml")

	var appConfig Config

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)

		errHanlerInfo.Error = err
		return getDefaultConfig()
	}

	err := viper.Unmarshal(&appConfig)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)

		errHanlerInfo.Error = err
		return getDefaultConfig()
	}

	return &appConfig
}

func getDefaultConfig() *Config {

	application := &Application{
		Name:    APP_NAME,
		Profile: PROFILE,
	}

	server := &Server{
		Port: SERVER_PORT,
	}

	// redis := &Redis{
	// 	Address: REDIS_ADDRESS,
	// }

	// envctrl := &EnvCtrl{
	// 	MaxThread: MAX_THREAD,
	// }

	// getRedisCacheValues := &GetRedisCacheValues{
	// 	MaxInputCacheKeys: MAX_INPUT_GET_CACHE_KEY,
	// }

	// setRedisCacheMultiple := &SetRedisCacheMultiple{
	// 	MaxInputCacheKeys: MAX_INPUT_SET_CACHE_KEY,
	// }

	// delRedisCacheMultiple := &DelRedisCacheMultiple{
	// 	MaxInputCacheKeys: MAX_INPUT_DEL_CACHE_KEY,
	// }

	// api := &Api{
	// 	GetRedisCacheValues:   *getRedisCacheValues,
	// 	SetRedisCacheMultiple: *setRedisCacheMultiple,
	// 	DelRedisCacheMultiple: *delRedisCacheMultiple,
	// }

	return &Config{
		Application: *application,
		Server:      *server,
		// Redis:       *redis,
		// EnvCtrl: *envctrl,
		// Api:     *api,
	}

}
