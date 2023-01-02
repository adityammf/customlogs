//TO EDIT

package utility

import (
	"os"
	"path/filepath"
	"time"
)

var (
	version      string
	buildVersion string
)

//not all the flags used yet
type configuration struct {
	Country      string `long:"country" short:"c" description:"Country Name" env:""`
	Currency     string `long:"currency" short:"" description:"Currency Used" env:""`
	BinDir       string
	DebugParam   bool   `long:"debug" short:"d" description:"Debug Flag" env:""`
	ConfigPath   string `long:"conf" short:"" description:"configuration path" env:""`
	V            bool   `long:"version" short:"v" description:"Get Version of binary" env:""`
	Version      string
	BuildVersion string
	Host         string `long:"host" short:"" description:"Hosted On" env:""`
	Port         string `long:"port" short:"" description:"Port Used" env:""`
	StartTime    time.Time
}

var binDir string

//Configuration : Variable of configuration struct type populated with default values
var Configuration configuration

func init() {
	binDir, _ = getBinDir()

	// Configuration = configuration{
	// 	DebugParam:   false,
	// 	ConfigPath:   binDir + "/conf/custom_logs.yml",
	// 	StartTime:    time.Now(),
	// 	Version:      version,
	// 	BuildVersion: buildVersion,
	// }
}

//ReadConfig : Read configuration file into the structure overwriting the default values with viper
// func ReadConfig(filename string) (*viper.Viper, error) {
// 	defaults := map[string]interface{}{
// 		"mongoDetails": map[string]string{
// 			"mongoUrl":        "mongodb://127.0.0.1:27017",
// 			"mongodbUsername": "",
// 			"mongodbPassword": "",
// 		},
// 		"country":  "",
// 		"currency": "",
// 		"host":     "",
// 		"port":     "",
// 	}

// 	v := viper.New()
// 	for key, value := range defaults {
// 		v.SetDefault(key, value)
// 	}

// 	configDir, err := filepath.Abs(filepath.Dir(filename))
// 	if err != nil {
// 		return nil, err
// 	}
// 	configFile := strings.Split(filepath.Base(filename), ".")[0]

// 	v.SetConfigName(configFile)
// 	v.AddConfigPath(configDir)
// 	v.AutomaticEnv()
// 	v.ReadInConfig()

// 	return v, nil

// }

func getBinDir() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ".", err
	}
	return dir, err
}
