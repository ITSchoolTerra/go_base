package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fInp, err := os.Open("09File/info.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	scan := bufio.NewScanner(fInp)
	fOut, err := os.Create("09File/infoOut.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for scan.Scan() {
		scan.Text()
		s := ""
		fmt.Scan(&s)
		fOut.WriteString(scan.Text() + s + "\n")
	}
	fInp.Close()
	fOut.Close()
}
