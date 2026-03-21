package config

type DatabaseConfig struct {
	DbType       string
	Addr         string
	MaxOpenConn  int
	MaxIdealConn int
	MaxIdealTime string
}
