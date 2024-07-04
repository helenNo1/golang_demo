package thread

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/helenNo1/xueyi/poc"
	"github.com/helenNo1/xueyi/poc/base"
)

type Thread struct {
	ThreadWg    *sync.WaitGroup
	ThreadChan  chan struct{}
	ToCheckChan chan string
}

func NewThread(queueChan chan string) *Thread {
	return &Thread{
		ThreadWg:    &sync.WaitGroup{},
		ThreadChan:  make(chan struct{}),
		ToCheckChan: queueChan,
	}
}

func (nt *Thread) Threadfunc() {
	defer nt.ThreadWg.Done()
	if strings.TrimSpace(base.PocNamesStr) == "" {
		log.Fatal("PocNameStr nil")
	}
	pocname_list := strings.Split(base.PocNamesStr, ",")
	time.Sleep(3 * time.Second)
	defer close(nt.ThreadChan)

	for {
		if url, ok := <-nt.ToCheckChan; ok {
			nt.ThreadChan <- struct{}{}
			nt.ThreadWg.Add(1)
			go func() {
				defer nt.ThreadWg.Done()
				defer func() {
					<-nt.ThreadChan
				}()

				poclist_item := &poc.PocList{}
				poclist_item.NameList = pocname_list

				poclist_item.Url = url
				//log.Println(poclist_item.NameList, poclist_item.Url)
				poclist_item.Build()
				poclist_item.Run()
			}()
		} else {
			break
		}
	}
}
