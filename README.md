## host-local-tools

> this is for resolving IPAM on type host-local, when deploying /etc/cni/multus/net.d/xxx.cnf on every
host. To ensure the range ip being not crossing, use etcd distribute lock to allocate IP ranges. And
write ip ranges to /etc/cni/multus/net.d/xxx.cnf.


- 读取 模板文件 resource/host-local-template.json

- 根据 app.ini 配置文件，进行各项配置

    `$ cat app.ini`
    ```yaml
    [PathConfig]
    SavePath = /opt/example.conf   # 写入配置文件的位置
    
    [VFConfig]
    VFName = ifs0     # master name  mallnox PV name
    Name = sriov-cnf  # cni network resource name
    Type = sriov  # cni network type
    Mode = bridge
    
    [IPAllocate]
    Subnet = 188.188.0.1/16  #pre-allocate ip mask
    
    [EtcdConfig]  
    # cluster etcd ips, splited by whitespace " ",
    # such as EtcdAddr = 192.168.0.157:2379  192.168.0.158:2379 
    EtcdAddr = 192.168.0.157:2379  
    
    [LogConfig]
    LogPath = "D:/temp1/logs"   # log path
    ```

- 使用ETCD 分布式锁以及etcd k-v 存储引擎 进行记录已分配IP 

- 根据模板 和 分配的IP 生成配置文件
  
  **note:** 实际存放地址 `/etc/cni/multus/net.d/xxx.cnf`
  
    `$ cat /opt/example.conf`
     ```json
    {
        "name": "sriov-cnf",
        "type": "sriov",
        "if0": "ifs0",
        "mode": "bridge",
        "ipam": {
            "type": "host-local",
            "rangeStart": "188.188.5.198",
            "rangeEnd": "188.188.6.5",
            "subnet": "188.188.0.1/16",
            "gateway": "188.188.0.1",
            "routes": [
                {
                "dst": "0.0.0.0/0"
                }
            ]
        }
    }
    ```