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
	var ctx context.Context
	logger = inilog.GetLog()
	client, err := clientv3.New(em.Conf)
	if err != nil {
		logger.Errorln(fmt.Sprintf("init etcd client failed: %v", err))
		return err
	}
	em.txn = clientv3.NewKV(client).Txn(context.TODO())
	em.lease = clientv3.NewLease(client)
	leaseResp,err := em.lease.Grant(context.TODO(),em.Ttl )
	if err != nil {
		logger.Errorln(fmt.Sprintf("grant lease failed : %v", err))
	}
	ctx, em.cancel = context.WithCancel(context.TODO())
	em.leaseId = leaseResp.ID
	_, err = em.lease.KeepAlive(ctx, em.leaseId)
	return err
}