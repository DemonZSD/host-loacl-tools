// Copyright 2019 Weshzhu
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// author: Weshzhu
package utils

import (
	"go.etcd.io/etcd/clientv3"
	"context"
	inilog "config-writer/log"
	"github.com/sirupsen/logrus"
	"fmt"
	"config-writer/types"
	"time"
)

type EtcdMutex struct {
	Ttl int64 // 租约时间
	Conf clientv3.Config // ETCD 集群配置
	Key string
	cancel context.CancelFunc
	lease clientv3.Lease //
	leaseId clientv3.LeaseID
	txn clientv3.Txn
}
var logger *logrus.Logger
func (em * EtcdMutex)init() error {

	logger = inilog.GetLog()
	client, err := clientv3.New(em.Conf)
	if err != nil {
		logger.Errorln(fmt.Sprintf("init etcd client failed: %v", err))
		return err
	}
	em.txn = clientv3.NewKV(client).Txn(context.TODO())
	em.lease = clientv3.NewLease(client)
	//设置租约过期时间
	leaseResp,err := em.lease.Grant(context.TODO(),em.Ttl )
	if err != nil {
		logger.Errorln(fmt.Sprintf("grant lease failed : %v", err))
	}
	ctx, em.cancel = context.WithCancel(context.TODO())
	em.leaseId = leaseResp.ID

	return err
}

func DoTask(allocateIP func(_ *types.HostLocal) error, em EtcdMutex, client *clientv3.Client){
	//启动一个协程去监听
	leaseRespChan, err := em.lease.KeepAlive(ctx, em.leaseId)  //自动续租
	go listenLeaseChan(leaseRespChan)
	//业务处理
	kv := clientv3.NewKV(client)
	//创建事务
	txn := kv.Txn(context.TODO())
	txn.If(clientv3.Compare(clientv3.CreateRevision("/cron/lock/job9"),"=",0)).
		Then(clientv3.OpPut("/cron/lock/job9","xxx",clientv3.WithLease(leaseId))).
		Else(clientv3.OpGet("/cron/lock/job9"))//否则抢锁失败

	//提交事务
	if txtResp,err :=txn.Commit();err != nil {
		fmt.Println(err)
		return
	} else {
		//判断是否抢锁
		if !txtResp.Succeeded {
			fmt.Println("锁被占用：",string(txtResp.Responses[0].GetResponseRange().Kvs[0].Value))
			return
		}
	}

	fmt.Println("处理任务")



	//释放锁（停止续租，终止租约）
	defer cancleFunc()//函数退出取消自动续租
	defer lease.Revoke(context.TODO(),leaseId) //终止租约（去掉过期时间）

	time.Sleep(10 * time.Second)
}


func listenLeaseChan(leaseRespChan <-chan *clientv3.LeaseKeepAliveResponse) {
	var (
		leaseKeepResp *clientv3.LeaseKeepAliveResponse
	)
	for {
		select {
		case leaseKeepResp = <-leaseRespChan:
			if leaseKeepResp == nil {
				fmt.Println("租约失效了")
				goto END
			} else {
				fmt.Println(leaseKeepResp.ID)

			}

		}
	}
END:
}