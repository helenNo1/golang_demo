package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"sync"

	"github.com/helenNo1/xueyi/command_line"
	"github.com/helenNo1/xueyi/poc/base"
	"github.com/helenNo1/xueyi/thread"
	"github.com/helenNo1/xueyi/util"
)

var addtarget_wg sync.WaitGroup
var target_queue_chan = make(chan string, 100000)

func setLogLineFile() {
	log.SetFlags(log.LstdFlags | log.Lshortfile) // 设置日志格式，包括时间和文件行号
}

func cidr2StrList(file string) {
	defer addtarget_wg.Done()
	defer close(target_queue_chan)
	ips := util.ReadLinesFromFile(file)
	for _, ip := range ips {
		list_ := util.Cidr2IPs(ip)
		for _, v := range list_ {
			target_queue_chan <- v
		}
	}
}

func file2StrList(file string) {
	defer addtarget_wg.Done()
	defer close(target_queue_chan)
	f, _ := os.Open(file)
	defer f.Close()
	reader := bufio.NewReader(f)

	for {
		inputString, _, readerError := reader.ReadLine()
		if readerError == io.EOF {
			println("io.eof")
			break
		}
		//log.Printf("push tocheck chan: %s\n", inputString)
		if string(inputString) != "" {
			target_queue_chan <- string(inputString)
		}
	}

}

func main() {
	setLogLineFile()

	cl := command_line.NewCommandLine()
	flag.StringVar(&cl.Mode, "m", "a", "mode default ")
	flag.StringVar(&cl.Src, "s", "src", "src filename")
	flag.StringVar(&base.PocSuccFile, "d", "dst", "Dst filename")
	flag.IntVar(&cl.ThreadNum, "t", 1, "thread num")
	flag.StringVar(&base.PocNamesStr, "p", "", "pocname split by ,")
	flag.StringVar(&util.ProxyStr, "proxy", "socks5://127.0.0.1:10808", "proxystr socks5://127.0.0.1:10808")
	// flag.BoolVar(&cl.Email, "e", false, "send mail")
	flag.Parse()

	//删除目标文件
	util.ClearDst(base.PocSuccFile)
	//读取文件url到chan ，异步
	addtarget_wg.Add(1)
	if _, err := os.Lstat(cl.Src); err != nil {
		log.Fatal(err, "c1.Src nil")
	}
	if cl.Mode == "a" {
		go file2StrList(cl.Src)
	} else {
		go cidr2StrList(cl.Src)
	}

	t := thread.NewThread(target_queue_chan)
	t.ThreadChan = make(chan struct{}, cl.ThreadNum)
	t.ThreadWg.Add(1)
	go t.Threadfunc()
	t.ThreadWg.Wait()

	addtarget_wg.Wait()
}
