package configuration

import (
	"time"
)

var Config *Configuration

type Configuration struct {
	App             App             `mapstructure:"app"`
	Log             Log             `mapstructure:"log"`
	DatabaseConfigs DatabaseConfigs `mapstructure:"database"`
	Redis           Redis           `mapstructure:"redis"`
	WsClient        WsClient        `mapstructure:"ws-client"`
	AccountSetup    AccountSetup    `mapstructure:"account-setup"`
}

type App struct {
	Name          string `mapstructure:"port"`
	Port          int    `mapstructure:"port"`
	Version       string `mapstructure:"version"`
	MobileVersion string `mapstructure:"mobile-version"`
	ProjectId     string `mapstructure:"project_id"`
	Env           string `mapstructure:"env"`
}

type Log struct {
	Env   string `mapstructure:"env"`
	Level string `mapstructure:"level"`
}

type DatabaseConfigs struct {
	SampleDB DbConfig `mapstructure:"sample"`
}

type DbConfig struct {
	Host               string        `mapstructure:"host"`
	Port               int           `mapstructure:"port"`
	Username           string        `mapstructure:"username"`
	Password           string        `mapstructure:"password"`
	Schema             string        `mapstructure:"schema"`
	MaxIdleConns       int           `mapstructure:"maxIdleConns"`
	MaxOpenConns       int           `mapstructure:"maxOpenConns"`
	MaxLifeTimeMinutes time.Duration `mapstructure:"maxLifeTimeMinutes"`
}

type Redis struct {
	Host          string        `mapstructure:"host"`
	Port          string        `mapstructure:"port"`
	Database      int           `mapstructure:"database"`
	PoolMaxIdle   int           `mapstructure:"poolMaxIdle"`
	PoolMaxActive int           `mapstructure:"poolMaxActive"`
	PoolTimeOut   time.Duration `mapstructure:"poolTimeout"`
}

type WsClient struct {
	Timeout  int                   `mapstructure:"timeout"`
	Services map[string]WebService `mapstructure:"services"`
}

type WebService struct {
	BaseUrl    string `mapstructure:"base-url"`
	XAuthToken string `mapstructure:"x-auth-token"`
}

type AccountSetup struct {
	CreditLimit          string `mapstructure:"credit-limit"`
	TermLoanAccountLimit int    `mapstructure:"term-loan-account-limit"`
}
