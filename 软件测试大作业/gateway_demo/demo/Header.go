package demo

import (
	"bufio"
	"fmt"
	"github.com/anaskhan96/soup"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

var Path = "D:/comic"

func main () {
	var url string
	fmt.Scanf("%s",&url)
	GetWebsite(url)
}

func GetWebsite(url string) {
	urls := make([]string,0)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	resp,err := soup.Get(url)
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	links := doc.FindAll("div","class","gdtm")
	for _, link := range links {
		l := link.Find("a")
		urls = append(urls,l.Attrs()["href"])
	}
	title:=doc.Find("h1")
	fmt.Println(title.Text())
}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func saveGirlPicture(imgUrl, imgPath, referer string) {

	//创建文件夹
	exist, err := PathExists(imgPath)

	if !exist {
		_ = os.Mkdir(imgPath, os.ModePerm)
	}

	fileTemName := path.Base(imgUrl)
	fileName := fileTemName[0:strings.Index(fileTemName, "?")]

	req, _ := http.NewRequest("GET", imgUrl, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")
	req.Header.Set("Referer", referer)

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)
	file, err := os.Create(imgPath + fileName)
	if err != nil {
		panic(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)
	written, _ := io.Copy(writer, reader)
	fmt.Printf("Total length: %d", written)
}
