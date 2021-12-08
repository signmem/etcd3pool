package etcdv3

import (
	"errors"
	"fmt"
	"github.com/signmem/etcd3pool/g"
	"log"
	"time"
	"context"
	"go.etcd.io/etcd/clientv3"
	"strings"
)

func KeyNotFound(s string) bool {
	return strings.HasPrefix(s, "100: key not found")
}

func (c *EtcdClient) GetMultiKeyValues(keyList []string) (map[string]string, error) {

	// params keyList ["key1", "key2']
	// return {["key1":"value"], ["key2":"value"]}

	//  key

	vars := make(map[string]string)
	requestTimeOut :=  3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeOut)

	for _,  key :=  range keyList {
		// resp, err := c.Client.Get(ctx, key, clientv3.WithPrefix())
		resp, err := c.Client.Get(ctx, key)
		if err != nil {
			fmt.Printf("[WARN] etcd get values key err: %s\n", err)
			continue
		}


		if len(resp.Kvs) > 0 {
			for _, kv := range resp.Kvs {
				kvKEY := string(kv.Key)
				kvVALUE := string(kv.Value)
				vars[kvKEY] = kvVALUE
			}
		}
	}
	cancel()
	return vars, nil
}

func (c *EtcdClient) GetkeyWithRecursive(key string) (map[string]string, error) {

	// use to range keyList
	// recursive : true
	// Traverse all info in key

	vars := make(map[string]string)
	requestTimeOut :=  3 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeOut)

	resp, err := c.Client.Get(ctx, key, clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	cancel()
	if err != nil {
		return vars, err
	}

	if len(resp.Kvs) > 0 {
		for _, kv := range resp.Kvs {
			kvKEY := string(kv.Key)
			kvVALUE := string(kv.Value)
			vars[kvKEY] = kvVALUE
		}
	}

	return vars, nil
}



func (c *EtcdClient) GetOneKeyValue(key string) (val string, err error) {
	// get one keys value
	// resursive : false

	requestTimeOut :=  3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeOut)
	resp, err := c.Client.Get(ctx, key,
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	cancel()

	if err != nil {
		return "", err
	}

	// how to verify key not exists or key is one directory ???
	// to be continue
	if len(resp.Kvs) > 0 {
		val = string(resp.Kvs[0].Value)
	}
	return val, nil
}



func (c *EtcdClient) IsADirectory(key string) (status bool, err error) {
	requestTimeOut :=  3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeOut)
	resp, err := c.Client.Get(ctx, key,
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	cancel()
	if err != nil {
		return false, err
	}

	if resp.Count == 0 {
		return true, nil
	}
	return false, nil
}

func (c *EtcdClient) KeyIsExsit(key string) (status bool, err error) {

	_, err = c.GetOneKeyValue(key)
	if err != nil {
		if KeyNotFound(err.Error()) {
			return false, nil
		} else {
			log.Printf("[WARN] etcd key is exists err: %s", err)
			return true, err
		}
	}
	return true , nil
}

func (c *EtcdClient) SetKeyValue(key string, val string) (string, error) {

	// input key value
	var KeyIsNull = errors.New("etcd: key can not be null")

	if key == "" {
		return "", KeyIsNull
	}

	requestTimeOut :=  3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeOut)
	_, err := c.Client.Put(ctx, key, val)
	cancel()
	if err != nil {
		return "", err
	}

	if g.Config().Debug {
		fmt.Printf("[INFO] etcd set key %s, val: %s\n", key, val)
	}

	return val, nil
}

func (c *EtcdClient) DelKey(key string, recursive bool) (status bool, err error) {
	requestTimeOut :=  3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeOut)

	var ops clientv3.OpOption
	if recursive == true {
		ops = clientv3.WithPrefix()
	} else {
		ops = clientv3.WithKeysOnly()
	}

	_, err = c.Client.Delete(ctx, key, ops)
	cancel()

	if err != nil {
		fmt.Printf("DelKey() error: %s\n", err)
		return false,  err
	}
	if g.Config().Debug {
		fmt.Printf("etcd del key:%s, recursive: %t\n", key, recursive)
	}

	return  true,nil
}
