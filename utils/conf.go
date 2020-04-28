package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

const (
	APPModeDEV = "develop"
	APPModePro = "product"
)

var confReader *ini.File

func init()  {
	source := fmt.Sprintf("%s.ini", APPModeDEV)
	var err error
	confReader, err = ini.Load(source)
	if err != nil {
		os.Exit(1)
	}
}

func GetString(section string, key string) string {
	return confReader.Section(section).Key(key).String()
}

func GetInt(section string, key string) int {
	return confReader.Section(section).Key(key).MustInt(0)
}

func GetBool(section string, key string) bool {
	return confReader.Section(section).Key(key).MustBool(false)
}