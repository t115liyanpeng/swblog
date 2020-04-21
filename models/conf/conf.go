package conf

//Servercfg 服务配置信息
type Servercfg struct {
	WebName string `json:"webname"`
	Port 	string `json:"port"`
}

//Databasecfg 数据库配置
type Databasecfg struct {
	IPAddress string `json:"address"`
	Port      int    `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
	DbName    string `json:"dbname"`
}

//Config 程序的配置
type Config struct {
	Server   *Servercfg   `json:"servercfg"`
	Database *Databasecfg `json:"databasecfg"`
}
