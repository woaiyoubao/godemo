package download

import (
	"strings"
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"io"
	"bytes"
)

type Download struct {
	Url string
}

const(
	savePath2 = "E://temp//goimg//"
)

func (download *Download) GetImg2() (n int64, err error) {
	path := strings.Split(download.Url, "/")
	var name string
	if len(path) > 1 {
		name = savePath2 + path[len(path)-1]
	}
	fmt.Println(name)
	out, err := os.Create(name)
	defer out.Close()
	resp, err := http.Get(download.Url)
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	n, err = io.Copy(out, bytes.NewReader(pix))
	return

}

func GetImg3(url string) (n int64, err error) {
	path := strings.Split(url, "/")
	var name string
	if len(path) > 1 {
		name = savePath2 + path[len(path)-1]
	}
	fmt.Println(name)
	out, err := os.Create(name)
	defer out.Close()
	resp, err := http.Get(url)
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	//n, err = io.Copy(out, bytes.NewReader(pix))
	out.Write(pix)
	return

}
