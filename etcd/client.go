package etcd

import (
	"fmt"
	"log"
	"strconv"
	"github.com/signmem/etcd3pool/g"
	"time"
)

var (
	WRITE = 0
	READ = 0
)


func WriteTest() {
	client, err := NewEtcd3()
	if err != nil {
		fmt.Printf("[ERROR] connection to etcd Server err:%s", err)
		return
	}

	num := 1
	endLine := g.Config().TestLine
	for {
		if num >= endLine {
			num = 1
		}
		timeNow := time.Now().Unix()
		timeNowStr := strconv.FormatInt(timeNow, 10)
		numStr := strconv.Itoa(num)
		hostname := "/" + "falcon-agent-test-" + numStr
		_, err = client.SetKeyValue(hostname, timeNowStr )
		if err != nil {
			log.Printf("[ERROR] put host %s err:%s", hostname, err)
		}
		num += 1
		WRITE +=1
	}
}


