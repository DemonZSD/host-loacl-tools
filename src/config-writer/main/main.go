package main

import (
	config_writer "config-writer"
	"config-writer/config"
	"fmt"
)

func main(){
	var aa = config.Appcfg
	fmt.Println(aa.SavePath)
	path1 := "src/config-writer/template/host-local.json"
	hostlocal, err := config_writer.ReadJsonFile(path1)
	if err != nil {
		fmt.Println("read template failed")
	}
	fmt.Println(hostlocal)
	fmt.Println(aa.EtcdAddr)
	fmt.Println(aa.LogPath)



}