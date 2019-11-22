package main

import (
	"context"
	"log"
	"time"
	
	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/pkg/transport"
)

var (
	defaultTimeout = 10 * time.Second
	Client         *clientv3.Client
)

func init() {
	endpoints := []string{
		"127.0.0.1:2379"
	}
	tlsInfo := transport.TLSInfo{
		CertFile:      "/tmp/test-certs/test-name-1.pem",
		KeyFile:       "/tmp/test-certs/test-name-1-key.pem",
		TrustedCAFile: "/tmp/test-certs/trusted-ca.pem",
	}
	tlsConfig, err := tlsInfo.ClientConfig()
	if err != nil {
		log.Fatal(err)
	}

	Client, err = clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
		TLS:         tlsConfig,
	})
	if err != nil {
		log.Fatal(err)
	}
	// defer Client.Close() // make sure to close the client
}

func main() {
	m,err := GetLockSession("hello","boy")
	if err != nil {
		fmt.Println("m get etcd lock failed:",err)
		return
	}

	// 添加带超时时间的ctx
	ctx,cancel := context.WithTimeout(context.Background(),15*time.Second)
	defer cancel() 

	err = m.Lock(ctx)
	if err !=nil {
		fmt.Println("m lock failed:",err)
		return
	}

	fmt.Println("m lock ...")

	go func() {
		select {
		case <-ctx.Done:
			err = m.Unlock(context.TODO())
			if err != nil {
				fmt.Println("unlock:",err)
				return
			}
			fmt.Println("m unlock ...")
		}
	}() 

	for i := 1; i< 8; i++ {
		time.Sleep(1*time.Second)
		fmt.Println("sleep ",i)
	}
}

func GetLockSession(key string,id string) (*concurrency.Mutex,error) {
	def := "/com/lock"
	lockKey := fmt.Sprintf("%s/%s-%s",def,key,id)
	session,err:=concurrency.NewSession(Client)
	if err != nil {
		return nil,fmt.Errorf("get lock session %v",err)
	}

	return concurrency.NewMutex(session,lockKey),nil
}
