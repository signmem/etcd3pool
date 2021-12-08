package cron

import (
	"github.com/signmem/etcd3pool/etcd"
	"sync"
	"time"
)

var (
	m *sync.RWMutex
)

func ResetMetric() {
	for {
		etcd.READ = etcd.WRITE
		etcd.WRITE = 0
		time.Sleep( 60 * time.Second )
	}
}