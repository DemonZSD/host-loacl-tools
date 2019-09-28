package main

import (
	"config-writer"
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
	hostlocal.Ipam.GenerateIpRanges("188.188.0.100", "188.188.0.150")
	hostlocal.Ipam.SetSubnet("188.188.0.1/16")
	hostlocal.Ipam.SetGateway("188.188.0.1")
	hostlocal.Master="enp24s0"
	hostlocal.Name="sriov-cnf"
	hostlocal.Type="sriov"
	hostlocal.Mode = "bridge"
	config_writer.WriteJsonToFile(cfg.SavePath, hostlocal)


	fmt.Println(config_writer.IncrementIP("188.188.0.1","188.188.0.1/16"))
}
