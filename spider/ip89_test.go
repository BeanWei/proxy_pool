package spider

import (
	"fmt"
	"github.com/phpgao/proxy_pool/model"
	"testing"
	"time"
)

func TestIp89_Fetch(t *testing.T) {
	newProxyChan := make(chan *model.HttpProxy, 100)
	spider := ip89{Spider{ch: newProxyChan}}
	spider.Run()
	timeout := time.After(30 * time.Second)
	for {
		select {
		case proxy := <-newProxyChan:
			fmt.Println(proxy)
			go func(p *model.HttpProxy) {
				flag := p.SimpleTcpTest()
				fmt.Println(flag)
			}(proxy)
		case <-timeout:
			fmt.Println("There's no more time to this. Exiting!")
			return
		}
	}
}