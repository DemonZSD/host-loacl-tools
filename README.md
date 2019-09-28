## host-local-tools
this is for resolving IPAM on type host-local, when deploying /etc/cni/multus/net.d/xxx.cnf on every
host. To ensure the range ip being not crossing, use etcd distribute lock to allocate IP ranges. And
write ip ranges to /etc/cni/multus/net.d/xxx.cnf.
