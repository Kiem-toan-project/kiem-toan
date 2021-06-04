package main

import (
	"backend/common/httpreq"
	"crypto/tls"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"time"
)
var (
	flConfigFile = ""
	flConfigYaml = ""
	flExample    = false
	flNoEnv      = false
)

func InitFlags() {
	flag.StringVar(&flConfigFile, "config-file", "", "Path to config file")
	flag.StringVar(&flConfigYaml, "config-yaml", "", "Config as yaml string")
	flag.BoolVar(&flNoEnv, "no-env", false, "Don't read config from environment")
	flag.BoolVar(&flExample, "example", false, "Print example config then exit")
}

func ParseFlags() {
	flag.Parse()
}

func main() {
	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	rcfg := httpreq.RestyConfig{Client: client}
	rClient := httpreq.NewResty(rcfg)
	resp, err := rClient.R().Get("http://google.com")
	fmt.Printf("\nResponse Body: %v", resp)
	InitFlags()
	ParseFlags()
	// load config
	cfg, err := Load(false)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("cfg: ", cfg)
}

type ConfigPostgres struct {
	Protocol string `yaml:"protocol"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	SSLMode  string `yaml:"sslmode"`
	Timeout  int    `yaml:"timeout"`

	MaxOpenConns    int `yaml:"max_open_conns"`
	MaxIdleConns    int `yaml:"max_idle_conns"`
	MaxConnLifetime int `yaml:"max_conn_lifetime"`

	GoogleAuthFile string `yaml:"google_auth_file"`
}

type DBConfig struct {
	Postgres          ConfigPostgres `yaml:"postgres"`
}

// Config ...
type Config struct {
	Databases    DBConfig        `yaml:",inline"`
}

// Default ...
func Default() Config {
	cfg := Config{
		Databases: DBConfig{
			Postgres: DefaultPostgres(),
		},
	}
	return cfg
}

// Load loads config from file
func Load(isTest bool) (Config, error) {
	var cfg, defCfg Config
	defCfg = Default()
	err := LoadWithDefault(&cfg, defCfg)
	if err != nil {
		return cfg, err
	}

	//cfg.Databases.MustLoadEnv()
	return cfg, err
}

// DefaultPostgres ...
func DefaultPostgres() ConfigPostgres {
	return ConfigPostgres{
		Protocol:       "",
		Host:           "postgres",
		Port:           5432,
		Username:       "postgres",
		Password:       "postgres",
		Database:       "test",
		SSLMode:        "",
		Timeout:        15,
		GoogleAuthFile: "",
	}
}

func LoadWithDefault(v, def interface{}) (err error) {
	defer func() {
		if flExample {
			if err != nil {
				//ll.Fatal("Error while loading config", l.Error(err))
			}
			//PrintExample(v)
			os.Exit(2)
		}
	}()

	if (flConfigFile != "") && (flConfigYaml != "") {
		return errors.New("must provide only -config-file or -config-yaml")
	}
	if flConfigFile != "" {
		err = LoadFromFile(flConfigFile, v)
		if err != nil {
			//ll.S.Errorf("can not load config from file: %v (%v)", flConfigFile, err)
		}
		return err
	}
	if flConfigYaml != "" {
		return LoadFromYaml([]byte(flConfigYaml), v)
	}
	reflect.ValueOf(v).Elem().Set(reflect.ValueOf(def))
	return nil
}

// LoadFromFile loads config from file
func LoadFromFile(configPath string, v interface{}) (err error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	return LoadFromYaml(data, v)
}

func LoadFromYaml(input []byte, v interface{}) (err error) {
	return yaml.Unmarshal(input, v)
}