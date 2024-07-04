package poc

import (
	"github.com/helenNo1/xueyi/poc/base"
	"github.com/helenNo1/xueyi/poc/mm"
)

type PocList struct {
	NameList []string
	Url      string

	poc_list []base.PocInt
}

func (pl *PocList) Build() {
	//log.Println(pl.NameList)
	for _, v := range pl.NameList {
		var p_i base.PocInt
		switch v {
		case "googlechart_rce":
			var p_ mm.GoogleChartRce
			p_.Name = v
			p_.Url = pl.Url
			p_i = &p_
			pl.poc_list = append(pl.poc_list, p_i)
		case "tp":
			var p_ mm.TpRce
			p_.Name = v
			p_.Url = pl.Url
			p_i = &p_
			pl.poc_list = append(pl.poc_list, p_i)
		case "xueyi":
			var p_ mm.XueyiWeak
			p_.Name = v
			p_.Url = pl.Url
			p_i = &p_
			pl.poc_list = append(pl.poc_list, p_i)
		}

	}

}

func (pl *PocList) Run() {
	for _, v := range pl.poc_list {
		v.Run()
	}
}
