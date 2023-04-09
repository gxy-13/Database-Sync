package conf

var Conf = new(AppConf)

type AppConf struct {
	Name   string `mapstructure:"name"`
	Author string `mapstructure:"author"`
	Mode   string `mapstructure:"mode"`
	*LogConf
	*MysqlConf
	*SQLServer
}

type MysqlConf struct {
	Host     string `mapstructure:"host"`
	Database string `mapstructure:"database"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
}

type SQLServer struct {
	Host     string `mapstructure:"host"`
	Database string `mapstructure:"database"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
}

type LogConf struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackUps int    `mapstructure:"max_backups"`
}
