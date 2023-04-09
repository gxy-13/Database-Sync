package cmd

import (
	"databasesync/conf"
	"databasesync/logger"
	"github.com/gin-gonic/gin"

	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = cobra.Command{
	Use:   "DBS",
	Short: "This is DBS.",
	Long:  "DataBase Syncer is a tool synchronization tool used for syncing between MySQL and SQL Server.",
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.Default()
		// Add your Gin routes here
		if err := logger.Init(conf.Conf.LogConf, conf.Conf.Mode); err != nil {
			log.Printf("logger.Init() failed, err:%v\n", err)
		}
		r.Run(":8080")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			log.Printf("homedir.Dir() failed, err:%v\n", err)
			os.Exit(1)
		}
		configsPath := filepath.Join(home, "conf")
		viper.AddConfigPath(configsPath)
		viper.SetConfigName("configs")
	}
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		os.Exit(1)
	}
	if err := viper.Unmarshal(conf.Conf); err != nil {
		log.Printf("viper.Unmarshal() failed, err:%v\n", err)
		os.Exit(1)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Println("configs.toml has changed.....")
		if err := viper.Unmarshal(conf.Conf); err != nil {
			log.Printf("viper.Unmarshal() failed, err:%v\n", err)
			os.Exit(1)
		}
	})
}
