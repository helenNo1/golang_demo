package util

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ClearDst(filename string) {
	if _, err := os.Lstat(filename); err != nil {
		_, err = os.Create(filename)
		if err != nil {
			log.Fatal(err, filename)
		}
	}
}

func Writeline2file(dst_name string, line string) {
	// log.Println(dst_name, line)
	bs, err := os.ReadFile(dst_name)
	if err != nil {
		log.Println(err)
		return
	}
	if strings.Contains(string(bs), line) {
		return
	}
	f, err := os.OpenFile(dst_name, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	_, err2 := io.WriteString(f, line+"\n")

	if err2 != nil {
		log.Println(err2)
		return

	}

}
func ReadAllFromFile(src_name string) string {
	filepath := src_name
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return ""
	}
	return string(content)
}

func ReadLinesFromFile(src_name string) []string {
	domain_list := []string{}
	f, err := os.Open(src_name)
	if err != nil {
		return domain_list
	}
	buf := bufio.NewReader(f)
	for {
		b, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		// log.Println(string(b))
		if !(strings.TrimSpace(string(b)) == "") {
			domain_list = append(domain_list, strings.TrimSpace(string(b)))
		}
	}
	return domain_list
}

func FileContainsStr(file, str string) bool {
	filestr := ReadAllFromFile(file)
	return strings.Contains(filestr, str)
}
