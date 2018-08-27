package jurge

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
	//	"github.com/mailto/sendmail"
)

var currentTimeData = time.Now().Format("2006-01-02")

var dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))

func Openfile() ([]string, map[string]string) {
	filename := dir + "/" + "people.txt"
	//	fmt.Println(filename)
	people, f_err := os.Open(filename)
	if f_err != nil {
		fmt.Printf("Error: %s\n", f_err)
	}
	defer people.Close()
	//namelist := []string{}
	br := bufio.NewReader(people)
	receive := []string{}
	m := make(map[string]string)
	for {
		n, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		//fmt.Println(string(n))
		if len(n) != 0 {
			//			Jurge(string(n))
			fields := strings.Fields(string(n))
			birthday := fields[len(fields)-1]
			if birthday == currentTimeData {
				receive = append(receive, fields[0])
				m[fields[0]] = fields[1]
			}
		}
		// jurge people birthday is today
	}
	return receive, m
}
