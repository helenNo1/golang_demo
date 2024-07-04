package util

import (
	"strings"

	"log"
	"net"

	"time"

	"crypto/tls"

	"net/http"

	"net/url"
)

func Url2Netloc(src string) string {
	u, _ := url.Parse(src)
	if u != nil {
		return u.Host
	}
	return ""
}

func GetHttpClient(proxy_str string, timeout int) *http.Client {
	if proxy_str == "" {
		return &http.Client{Timeout: time.Second * time.Duration(timeout),
			Transport: &http.Transport{
				TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
				DisableKeepAlives: true,
			},
		}
	}
	return &http.Client{Timeout: time.Duration(timeout) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
			DisableKeepAlives: true,
			Proxy: func(_ *http.Request) (*url.URL, error) {
				return url.Parse(proxy_str)
			},
		},
	}
}

func GetHttpNoredirectclient(proxy_str string, timeout int) *http.Client {
	if proxy_str == "" {
		return &http.Client{Timeout: time.Duration(timeout) * time.Second,
			Transport: &http.Transport{
				TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
				DisableKeepAlives: true,
			}, CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	}
	return &http.Client{Timeout: time.Duration(timeout) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
			DisableKeepAlives: true,
			Proxy: func(_ *http.Request) (*url.URL, error) {
				// return url.Parse(fmt.Sprintf("socks5://%s", proxy_str))
				return url.Parse(proxy_str)
			},
		}, CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}

func Cidr2IPs(cidr string) []string {
	// C段转ip
	var ips []string

	ipAddr, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		log.Print(err, cidr)
	}

	for ip := ipAddr.Mask(ipNet.Mask); ipNet.Contains(ip); increment(ip) {
		ips = append(ips, ip.String())
	}

	// CIDR too small eg. /31
	if len(ips) <= 2 {
		log.Print("err", cidr)
	}

	return ips
}

func increment(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] != 0 {
			break
		}
	}
}

func Url2netloc(url, protocol string) string {
	//log.Println(url)
	var port_suffix string
	switch protocol {
	case "redis":
		port_suffix = ":" + "6379"
	case "mongodb":
		port_suffix = ":" + "27017"
	case "postgres":
		port_suffix = ":" + "5432"
	case "mysql":

		port_suffix = ":" + "3306"
	case "mssql":
		port_suffix = ":" + "1433"
	}
	var netloc_tmp string
	//log.Println(url)
	var netloc string
	if strings.Contains(url, "http://") || strings.Contains(url, "https://") {
		netloc_tmp = Url2Netloc(url)
		if netloc_tmp == "" {
			return ""
		}
		netloc = strings.Split(netloc_tmp, ":")[0] + port_suffix
	} else {
		netloc = strings.Split(url, ":")[0] + port_suffix
	}

	return netloc
}
