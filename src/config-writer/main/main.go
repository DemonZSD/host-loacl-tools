package main

import (
	"config-writer"
	"config-writer/config"
	"fmt"
	"net"
	"config-writer/utils"
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
	var origIp = "188.188.0.1"
	var cidr = "188.188.0.1/16"
	startIp, err := utils.OffsetIPRange(1, net.ParseIP(origIp), cidr)
	endIp, err := utils.OffsetIPRange(5, startIp, cidr)
	hostlocal.Ipam.SetIpRanges(startIp, endIp)
	hostlocal.Ipam.SetSubnet("188.188.0.1/16")
	hostlocal.Ipam.SetGateway("188.188.0.1")
	hostlocal.Master="enp24s0"
	hostlocal.Name="sriov-cnf"
	hostlocal.Type="sriov"
	hostlocal.Mode = "bridge"
	config_writer.WriteJsonToFile(cfg.SavePath, hostlocal)
	fmt.Println()
}
