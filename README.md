## host-local-tools
this is for resolving IPAM on type host-local, when deploying /etc/cni/multus/net.d/xxx.cnf on every
host. To ensure the range ip being not crossing, use etcd distribute lock to allocate IP ranges. And
write ip ranges to /etc/cni/multus/net.d/xxx.cnf.


读取 模板文件 resource/host-local-template.json

#### 根据 app.ini 配置文件，进行各项配置

#### 使用ETCD 分布式锁以及etcd k-v 存储引擎 进行记录已分配IP 

#### 根据模板 和 分配的IP 生成配置文件
