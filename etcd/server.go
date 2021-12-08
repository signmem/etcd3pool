package etcd

import (
	"fmt"
	"strconv"
	"time"
)



func ServerControll() {

	client, err := NewEtcd3()
	if err != nil {
		fmt.Printf("[ERROR] connection to Etcd Server err:%s", err)
		return
	}

	for {

		hosts, err := client.GetkeyWithRecursive("/")
		if err != nil {
			fmt.Printf("[ERROR] ServerGetAllEtcdKey() err:%s", err)
			return
		}

		if len(hosts) > 0 {
			timeNow := time.Now().Unix()


			for host, value := range hosts {

				valueInt64, err := strconv.ParseInt(value, 10, 64)

				if err != nil {
					fmt.Printf("[ERROR] key %s, value %s err:%s\n", host, value, err)
					continue
				}

				if timeNow - valueInt64 > 300 {
					//  fmt.Printf("[BINGO] DO SOMETHING TO ALARM terry.tsang!!")
					_, err := client.DelKey(host, false)
					if err != nil {
						fmt.Printf("[ERROR] remove key %s error:%s\n", host, err)
					}
				}
			}
		}
		time.Sleep(60*time.Second)
	}
}