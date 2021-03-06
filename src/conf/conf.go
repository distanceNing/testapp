package conf

type DbConf struct {
	Addr     string `yaml:"addr"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type MongoDbConf struct {
	Addr     string `yaml:"addr"`
	Database string `yaml:"database"`
}

type RedisConf struct {
	Addr     string `yaml:"addr"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type AppConf struct {
	Name      string `yaml:"name"`
	Addr      string `yaml:"addr"`
	ImagePath string `yaml:"image_path"`
	CdnPath   string `yaml:"cdn_path"`
}

type ServerConf struct {
	AppConf     AppConf     `yaml:"appconf"`
	DbConf      DbConf      `yaml:"dbconf"`
	MongoDbConf MongoDbConf `yaml:"mongodbconf"`
	RedisConf   RedisConf   `yaml:"redisconf"`
}
