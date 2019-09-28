package main

import (
	config_writer "config-writer"
	"config-writer/config"
	"fmt"
)

func main() {
	var cfg = config.Appcfg
	fmt.Println(cfg.SavePath)
	path1 := "src/config-writer/template/host-local.json"
	hostlocal, err := config_writer.ReadJsonFile(path1)
	if err != nil {
		fmt.Println(fmt.Sprintf("read template failed: %v", err))
	}
	fmt.Println(fmt.Sprintf("hostlocal: %#v", hostlocal))
	fmt.Println(cfg.EtcdAddr)
	fmt.Println(cfg.LogPath)
}
