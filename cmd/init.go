package cmd

import (
	"ecommerce/common/log"
	"ecommerce/internal/sqlclient"
	"ecommerce/repository"
	"fmt"
	"os"
	"time"

	"github.com/caarlos0/env"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Dir  string `env:"CONFIG_DIR" envDefault:"config/config.json"`
	Port string
	DB   bool
}

var config Config

func init() {
	loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		log.Fatal(err)
	}
	time.Local = loc

	if err := env.Parse(&config); err != nil {
		log.Error("Get environment values fail")
		log.Fatal(err)
	}

	viper.SetConfigFile(config.Dir)
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err.Error())
		panic(err)
	}
	cfg := Config{
		Dir:  config.Dir,
		Port: viper.GetString(`main.port`),
		DB:   viper.GetBool(`main.database`),
	}
	if cfg.DB {
		sqlClientConfig := sqlclient.SqlConfig{
			Host:         viper.GetString(`database.host`),
			Database:     viper.GetString(`database.database`),
			Username:     viper.GetString(`database.username`),
			Password:     viper.GetString(`database.password`),
			Port:         viper.GetInt(`database.port`),
			DialTimeout:  20,
			ReadTimeout:  30,
			Timeout:      30,
			PoolSize:     10,
			MaxIdleConns: 10,
			MaxOpenConns: 10,
		}
		repository.InitRepo(sqlclient.NewSqlClient(sqlClientConfig))
	}
	config = cfg
}

func Execute() {
	var rootCmd = &cobra.Command{Use: "ecommerce"}
	rootCmd.AddCommand(cmdAPI)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("error:")
		fmt.Println(err)
		os.Exit(1)
	}
}
