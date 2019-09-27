package main

import (
	config_writer "config-writer"
	"config-writer/config"
	"fmt"
)

func main(){
	var aa = config.Config
	fmt.Println(aa.SavePath)
	path1 := "src/config-writer/template/host-local.json"
	hostlocal, err := config_writer.ReadJsonFile(path1)
	if err != nil {
		fmt.Println("read template failed")
	}

	hostlocal


}