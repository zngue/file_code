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
func RouterFile(SelectTableName string)  {
	model := new(FileNameChange).Case2Camel(SelectTableName)
	all := strings.ReplaceAll(routerTemp, "{{model}}", model)
	path := viper.GetString("temp.path")
	modepath := viper.GetString("temp.modepath")
	routerpath := viper.GetString("temp.routerpath")
	routerpath=path+"/"+routerpath
	CreateMutiDir(routerpath)
	modelFile:=routerpath+"/"+SelectTableName+".go"
	all = strings.ReplaceAll(all, "{{path}}", modepath+"/"+path)
	f, _ := os.Create(modelFile)
	f.Write([]byte(all))
}