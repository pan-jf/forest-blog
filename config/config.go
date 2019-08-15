package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const ConfigFilePath = "app.json"


type Config struct {
	SiteName string `json:"siteName"`

	Author string `json:"author"`
	
	Icp string `json:"icp"`
	
	Port string `json:"port"`

	PageSize int `json:"pageSize"`

	DocumentPath string `json:"documentPath"`

	HtmlKeywords string `json:"htmlKeywords"`

	HtmlDescription string `json:"htmlDescription"`

	CategoryListFileNumber int `json:"categoryListFileNumber"`
}




var Cfg Config

var CurrentDir string

func init()  {
	var pwdErr error
	CurrentDir,pwdErr = os.Getwd()

	if pwdErr != nil {
		panic(pwdErr)
	}

	configFile,err := ioutil.ReadFile(ConfigFilePath)

	if err != nil {
		panic(err)
	}

	jsonErr := json.Unmarshal(configFile,&Cfg)

	if jsonErr != nil {
		panic(err)
	}
	fmt.Println("init config...")
}