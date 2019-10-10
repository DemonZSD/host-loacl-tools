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
package etcd

import (
	"go.etcd.io/etcd/clientv3"
	"context"
	inilog "config-writer/log"
	"github.com/sirupsen/logrus"
	"fmt"
	"github.com/pkg/errors"
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
	var ctx context.Context
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
	_, err = em.lease.KeepAlive(ctx, em.leaseId)  //自动续租
	return err
}
func(em *EtcdMutex)Lock()error{
	err := em.init()
	if err != nil{
		return err
	}
	//LOCK:
	em.txn.If(clientv3.Compare(clientv3.CreateRevision(em.Key),"=",0)).
		Then(clientv3.OpPut(em.Key,"",clientv3.WithLease(em.leaseId))).
		Else()
	txnResp,err := em.txn.Commit()
	if err != nil{
		return err
	}
	if !txnResp.Succeeded{   //判断txn.if条件是否成立
		return errors.New("抢锁失败")
	}
	return nil
}

func(em *EtcdMutex)UnLock(){
	em.cancel()
	em.lease.Revoke(context.TODO(),em.leaseId)
	fmt.Println("释放了锁")
}


