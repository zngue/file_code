package code

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

/*
*@Author Administrator
*@Date 9/4/2021 12:17
*@desc 
 */
func ControllerFile(SelectTableName string)  {
	model := new(FileNameChange).Case2Camel(SelectTableName)
	all := strings.ReplaceAll(controllerTemp, "{{model}}", model)
	path := viper.GetString("temp.path")
	modepath := viper.GetString("temp.modepath")
	controllerpath := viper.GetString("temp.controllerpath")
	controllerpath=path+"/"+controllerpath
	CreateMutiDir(controllerpath)
	modelFile:=controllerpath+"/"+SelectTableName+".go"
	all = strings.ReplaceAll(all, "{{path}}", modepath+"/"+path)
	f, _ := os.Create(modelFile)
	f.Write([]byte(all))
}