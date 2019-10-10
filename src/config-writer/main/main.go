package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	inilog "config-writer/log"
	"config-writer/utils"
	"config-writer/types"
	"config-writer/config"
	"go.etcd.io/etcd/clientv3"
	"time"
	"config-writer/etcd"
	"net"
	"strings"
)
var logger *logrus.Logger
func init() {
	logger = inilog.GetLog()
}

func main() {
	var lockKey = "/lock/allocate/ipam"
	vfInfo := &types.VFInfo{
		Count: 0,
		Master: "enp24s0",
	}
	path1 := "src/config-writer/template/host-local.json"
	initHostLocal, err := utils.ReadJsonFile(path1)
	if err != nil {
		logger.Infoln(fmt.Sprintf("read template failed: %v", err))
	}

	vfNum, err := vfInfo.ReadVFNum()

	if err != nil {
		logger.Errorln(fmt.Sprintf("read vf file failed：%v", err))
	}
	DoTask(lockKey, vfNum, initHostLocal, AllocateIp)
}


func AllocateIp(hostlocal *types.HostLocal) error{
	var cfg = config.Appcfg

	//fmt.Println(cfg.SavePath)
	hostlocal.Ipam.SetSubnet(cfg.Subnet)
	hostlocal.Ipam.SetGateway(cfg.Subnet)
	hostlocal.Master="enp24s0"
	hostlocal.Name="sriov-cnf"
	hostlocal.Type="sriov"
	hostlocal.Mode = "bridge"
	return utils.WriteJsonToFile(cfg.SavePath, hostlocal)
}

func DoTask(localKey string, vfNum int, hostlocal *types.HostLocal, allocateIP func(*types.HostLocal) error){
	var cfg = config.Appcfg
	var conf = clientv3.Config{
		Endpoints:   strings.Split(cfg.EtcdAddr, " "),
		DialTimeout: 5 * time.Second,
	}
	allocateMutex := &etcd.EtcdMutex{
		Conf:conf,
		Ttl:10,
		Key: localKey,
	}
	err := allocateMutex.Lock()
	if err != nil{
		logger.Errorln(fmt.Sprintf("get lock failed: %v", err))
		return
	}
	logger.Infoln("get lock success")
	var originIp = "188.188.0.1"
	startIp, err := utils.OffsetIPRange(1, net.ParseIP(originIp), cfg.Subnet)
	endIp, err := utils.OffsetIPRange(vfNum, startIp, cfg.Subnet)
	hostlocal.Ipam.SetIpRanges(startIp, endIp)
	allocateIP(hostlocal)
	time.Sleep(10*time.Second)
	defer allocateMutex.UnLock()
}
