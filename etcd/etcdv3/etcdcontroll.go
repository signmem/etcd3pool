package etcdv3

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"time"
	"go.etcd.io/etcd/clientv3"
	"github.com/signmem/etcd3pool/g"
)



type EtcdClient struct {
	Client clientv3.KV
}

func EtcdKVControl() (*EtcdClient, error) {

	var kapi  clientv3.KV
	endPoints := g.Config().EtcdConfig.Host
	etcdCaFile := g.Config().EtcdSSL.CaFile
	etcdCertFile := g.Config().EtcdSSL.CertFile
	etcdCertKeyFile := g.Config().EtcdSSL.CertKeyFile


	cert, err := tls.LoadX509KeyPair(etcdCertFile, etcdCertKeyFile)
	if err != nil {
		log.Printf("[ERROR] load etcd cert file faile, err: %v", err)
		return  &EtcdClient{kapi}, err
	}

	caData, err := ioutil.ReadFile(etcdCaFile)
	if err != nil {
		log.Printf("[ERROR] load etcd ca file faile, err: %v", err)
		return  &EtcdClient{kapi}, err
	}

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caData)

	_tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      pool,
	}

	cfg := clientv3.Config{
		Endpoints: endPoints,
		TLS:       _tlsConfig,
		DialTimeout:  5 * time.Second,
	}

	cli, err := clientv3.New(cfg)  // ( *Client, error )
	if err != nil {
		log.Printf ("[ERROR] new etcd3 initial error: %v\n", err)
		return &EtcdClient{kapi}, err
	}

	kapi = clientv3.NewKV(cli)

	return &EtcdClient{kapi}, nil
}
