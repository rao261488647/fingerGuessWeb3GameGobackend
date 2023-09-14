package setting

type DataBase struct {
	DbType    string
	DbName    string
	Host      string
	Username  string
	Pwd       string
	Charset   string
	ParseTime bool
}

type BlockConfig struct {
	RpcUrl string
}
