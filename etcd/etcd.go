package etcd

import (
	"fmt"
	"github.com/signmem/etcd3pool/etcd/etcdv3"
)

func NewEtcd3() ( clientV3 *etcdv3.EtcdClient, err error ) {

	client, err := etcdv3.EtcdKVControl()   // *EtcdClient , err

	if err != nil {
		fmt.Printf("[ERROR] etcd init false: %v\n", err)
		return clientV3, err
	}

	clientV3 = client
	return  clientV3, nil
}
