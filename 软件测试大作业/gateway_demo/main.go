package main

import (
	"bufio"
	"fmt"
	"github.com/anaskhan96/soup"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//爬取指定url的网页内容
func HttpGet(url string) (result string, err error) {
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")
	req.Header.Set("Referer", "https://www.mzitu.com/all/")

	resp, err1 := (&http.Client{}).Do(req)

	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	//读取网页当前的内容
	buf := make([]byte, 4*1024)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		result += string(buf[:n])
	}
	return
}

//解析爬取的地址信息
//使用正则表达式解析当前内的地址
func SplitHomePage(htmlString string) (urls []string) {
	doc := soup.HTMLParse(htmlString)
	ps := doc.FindAll("p","class","url")
	for _,p := range ps {
		a := p.Find("a")
		urls = append(urls,a.Attrs()["href"])
	}
	return
}

//爬取一个妹子的每一张图片
func SpiderOneGirl(url, picPath string) bool {
	//获取当前页面的title信息
	fmt.Println("单页网址为:", url)
	time.Sleep(1000)
	result, err := HttpGet(url)
	if err != nil {
		fmt.Printf("url= %s 获取信息失败\n", url)
		return false
	}
	re := regexp.MustCompile(`<title>(?s:(.*?))</title>`)
	if re == nil {
		fmt.Println("regexp.MustCompile err")
	}

	title := re.FindStringSubmatch(result)
	fmt.Println("标题为:", title[1])
	//循环获取网址的信息,当网址请求访问返回值不为404时,进行图片地址解析
	i := 0
	for {
		if !SpiderOneGirlPicture(url, picPath+title[1]+"/", i) {
			break
		}
		time.Sleep(2000)
		i++
	}
	return false
}

//爬取某一个妹子的分页图片实际地址,并按照标题保存
func SpiderOneGirlPicture(url, picPath string, i int) bool {
	requestUrl := url + "/" + strconv.Itoa(i)
	picResult, err := HttpGet(requestUrl)

	if err != nil {
		//如果发生错误直接返回 不进行任何匹配
		return true
	}

	re := regexp.MustCompile(`><img src="(?s:(.*?))"`)
	if re == nil {
		fmt.Println("regexp.MustCompile err")
	}

	picUrl := re.FindStringSubmatch(picResult)

	if len(picUrl) > 0 {
		fmt.Println("需要保存的图片地址页为:", picUrl[1])
		//下载图片
		saveGirlPicture(picUrl[1]+"?width=700&height=1050", picPath, requestUrl)
		time.Sleep(2000)
		return true
	}

	return false
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

//保存网络图片
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

func DoWork() bool {

	fmt.Printf("准备获取妹子图")
	//获取全部页面的数据
	url := "https://www.mzitu.com/all/"
	//地址
	picPath := "D:/downloads/pic/mzitu/"
	//开始获取全部页面内容
	result, err := HttpGet(url)

	if err != nil {
		fmt.Println("HttpGet Error = ", err)
		return false
	}

	//获取到全部的妹子图图片网址切片
	urls := SplitHomePage(result)

	for _, url := range urls {
		if !SpiderOneGirl(url, picPath) {
			continue
		}
	}

	return true
}

func main() {
	//工作函数
	if DoWork() {
		fmt.Println("获取图片完毕")
	}
}