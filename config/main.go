package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"runtime"
	"path/filepath"
)

var (
	Yml *YmlFile
)

type (
	YmlFile struct {
		Database	DB	`yaml:"database"`
	}

	DB struct {
		Host     	string 		`yaml:"host"`
		Port  		string 		`yaml:"port"`
		User 		string 		`yaml:"user"`
		Password 	string		`yaml:"password"`
		DbName		string		`yaml:"dbname"`
	}
)

func init() () {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	source, err := ioutil.ReadFile(basePath + "/../config.yml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(source, &Yml)
	if err != nil {
		panic(err)
	}
}