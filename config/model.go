package config

type MysqlConf struct {
	Mysql struct{
		Dbname 		string 		`json:"dbname"`
		Password 	string 		`json:"password"`
		Host 		string 		`json:"host"`
		Port 		string 		`json:"port"`
		Username 	string 		`json:"username"`
	} `json:"mysql"`
}
