package main

import (
	//	"fmt"

	"github.com/mailto/jurge"
	"github.com/mailto/sendmail"
)

func main() {
	receive, m := jurge.Openfile()
	//	fmt.Println(receive, m)
	if len(receive) != 0 {
		sendmail.Sendmail(receive, m)
	}

}
