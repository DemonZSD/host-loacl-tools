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
	"strings"
	"net"
)
var logger *logrus.Logger
const hasAllocate = "/lock/allocated/ipam"
func init() {
	logger = inilog.GetLog()
	//var cfg = config.Appcfg
	//var conf = clientv3.Config{
	//	Endpoints:   strings.Split(cfg.EtcdAddr, " "),
	//	DialTimeout: 5 * time.Second,
	//}
	//allocateMutex := &etcd.EtcdMutex{
	//	Conf:conf,
	//	Ttl:10,
	//	Key: hasAllocate,
	//}
	//allocateMutex.DeleteKey(allocateMutex.Key)
}
func main() {
	var lockKey = "/lock/allocate/ipam"
	var cfg = config.Appcfg
	vfInfo := &types.VFInfo{
		Count: 0,
		Master: cfg.VFName,
	}
	path1 := "resource/host-local-template.json"
	initHostLocal, err := utils.ReadJsonFile(path1)
	if err != nil {
		logger.Infoln(fmt.Sprintf("read template failed: %v", err))
	}
	vfNum, err := vfInfo.ReadVFNum()

	if err != nil {
		logger.Errorln(fmt.Sprintf("read vf file failed：%v", err))
		return
	}
	//go DoTask(lockKey, vfNum, initHostLocal, AllocateIp)
	//go DoTask(lockKey, vfNum, initHostLocal, AllocateIp)
	//go DoTask(lockKey, vfNum, initHostLocal, AllocateIp)
	//go DoTask(lockKey, vfNum, initHostLocal, AllocateIp)
	//go DoTask(lockKey, vfNum, initHostLocal, AllocateIp)
	//go DoTask(lockKey, vfNum, initHostLocal, AllocateIp)
	DoTask(lockKey, vfNum, initHostLocal, AllocateIp)
	time.Sleep(20 * time.Second)
}

func AllocateIp(hostlocal *types.HostLocal) error{
	var cfg = config.Appcfg
	hostlocal.Ipam.SetSubnet(cfg.Subnet)
	hostlocal.Ipam.SetGateway(cfg.Subnet)
	hostlocal.Master=cfg.VFName
	hostlocal.Name=cfg.Name
	hostlocal.Type=cfg.Type
	hostlocal.Mode = cfg.Mode
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
	for {
		err := allocateMutex.Lock()
		if err != nil{
			logger.Errorln(fmt.Sprintf("get lock failed: %v", err))
			time.Sleep(1 * time.Second)
		}else{
			logger.Infoln("get lock success")
			break
		}
	}

	originIp, err := allocateMutex.GetValue(hasAllocate)
	if err != nil{
		logger.Errorln(fmt.Sprintf("get value failed : %v", err))
		return
	}
	if originIp == "" {
		if originIp, err = types.GetInitIpFromSubset(cfg.Subnet); err != nil {
			logger.Errorln(fmt.Sprintf("get init ip from subset failed : %v", err))
			return
		}
		if err := allocateMutex.Update(hasAllocate, originIp); err != nil{
			logger.Errorln(fmt.Sprintf("set value failed: %v", err))
			return
		}
	}
	startIp, err := utils.OffsetIPRange(1, net.ParseIP(originIp), cfg.Subnet)
	if err != nil {
		logger.Errorln(fmt.Sprintf("offset ip range failed: %v", err))
	}
	endIp, err := utils.OffsetIPRange(vfNum, startIp, cfg.Subnet)
	if err != nil {
		logger.Errorln(fmt.Sprintf("offset ip range failed: %v", err))
	}
	hostlocal.Ipam.SetIpRanges(startIp, endIp)
	allocateIP(hostlocal)
	if err := allocateMutex.Update(hasAllocate, endIp.String()); err != nil{
		logger.Errorln(fmt.Sprintf("set value failed: %v", err))
		return
	}
	logger.Infoln(fmt.Sprintf("success allocate ip [%s - %s]", startIp, endIp))
	defer allocateMutex.UnLock()
}
