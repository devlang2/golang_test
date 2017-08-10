package main

import (
	//	"bufio"
	//	"fmt"
	//	"io/ioutil"
	"os"
	"time"
)

const (
	LOGFILETIME = "20060102_150"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	data := []byte("hello\ngo\n")
	//	err := ioutil.WriteFile("dat1", d1, 0644)
	//	check(err)

	//    ioutil.WriteFile()

	//    f, err := os.Create("dat2")
	//	check(err)
	//	defer f.Close()
	//	d2 := []byte{115, 111, 109, 101, 10}
	//	n2, err := f.Write(d2)
	//	check(err)
	//	fmt.Printf("wrote %d bytes\n", n2)
	//	n3, err := f.WriteString("writes\n")
	//	fmt.Printf("wrote %d bytes\n", n3)
	//	f.Sync()
	//	w := bufio.NewWriter(f)
	//	n4, err := w.WriteString("buffered\n")
	//	fmt.Printf("wrote %d bytes\n", n4)
	//	w.Flush()

	fp := time.Now().Format(LOGFILETIME) + "0.log"
	f, err := os.OpenFile(fp, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	check(err)
	defer f.Close()

	//	fmt.Println(filename)
	f.Write(data)

}
