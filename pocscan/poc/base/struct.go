package base
type PocInt interface {
	Run()
}

type Poc struct {
	Vul      bool
	UserPass string
	Msg      string
	PocStart
	//PocInt
}

type PocStart struct {
	Name         string
	Url          string
	UserPassList []string
}

var PocSuccFile string
var PocNamesStr string
