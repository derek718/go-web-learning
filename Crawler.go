package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
)

type ThreadItem struct {
	url     string   //帖子路径
	content string   //帖子内容
	imgs    []string //帖子图片
}

var (
	//帖子url匹配
	threadItemExp = regexp.MustCompile(`"thread/[0-9]+"`)

	//图片地址
	threadImageExp = regexp.MustCompile(`src="//i\.4cdn\.org/asp/[0-9]+s\.jpg"`)
)

//找到贴子连接
func findThreads(url string) []ThreadItem {
	var threads = make([]ThreadItem, 0)

	content, err := httpGet(url)
	if err != 200 {
		return threads
	}

	//内容获取到
	urlList := threadItemExp.FindAllStringSubmatch(content, 10000)
	var sliceUrl = make([]string, 0)

	//把多余的"引号去掉，返回的是一个二维数组
	for _, v := range urlList {
		var yesRep = strings.Replace(v[0], "\"", "", -1)
		//把去重的字符放入到新建的sliceUrl
		sliceUrl = append(sliceUrl, yesRep)
	}

	//去重
	sort.Strings(sliceUrl)
	sliceUrl = unequal(sliceUrl)

	//把获取到的url参数进行拼装

	for _, v := range sliceUrl {
		threads = append(threads, ThreadItem{url: "http://boards.4chan.org/asp/" + v})
	}
	return threads

}

//请求一个地址
func httpGet(url string) (content string, statusCode int) {
	resp, err := http.Get(url)
	if err != nil {
		statusCode = -100
		return
	}

	defer resp.Body.Close()
	//写入网页所有内容io
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		statusCode = -200
		return
	}

	statusCode = resp.StatusCode
	content = string(data)
	return
}

//把重复的URL字符去除
func unequal(sliceUrl []string) []string {
	l := len(sliceUrl)
	var ret = make([]string, 0)
	for i := 0; i < l; i++ {
		if i > 0 && sliceUrl[i-1] == sliceUrl[i] {
			continue
		}
		ret = append(ret, sliceUrl[i])
	}

	return ret
}

//获取一个新贴子的内容
func (t *ThreadItem) getContents() *ThreadItem {
	content, err := httpGet(t.url)
	if err != 200 {
		t.content = ""
		return t
	}

	t.content = string(content)
	return t
}

//获取一个贴子的图片地址
func (t *ThreadItem) getImages() *ThreadItem {
	imgs := threadImageExp.FindAllStringSubmatch(t.content, 10000)
	var iL = make([]string, 0)
	for _, v := range imgs {
		//		fmt.Println(v[0])
		iL = append(iL, v[0])
	}
	t.imgs = iL
	return t
}

func (t *ThreadItem) download() {
	//	fmt.Println(t.url)
	last := strings.LastIndex(t.url, "/")
	dir := "./download" + string(t.url[last:len(t.url)])
	fmt.Println("create dir:", dir)
	err := os.Mkdir(dir, os.ModeDir)

	if err != nil {
		fmt.Println(err)
		return
	}
	os.Chmod(dir, 0777)
	downLoadImgs(t.imgs, dir)
	/*for _, v := range t.imgs {
		//"/"最后出现的位置
		//src="//i.4cdn.org/u/1505393100498s.jpg"
		pos := strings.LastIndex(v, "/")
		fileName := string(v[pos : len(v)-1])

		files, err := os.Create(dir + fileName)
		if err != nil {
			fmt.Println("Error for create file.")
			fmt.Println(err)
			continue
		}
		defer files.Close()

		data, err1 := downloadImg("http:" + v[5:len(v)-1])
		if err1 != 200 {
			fmt.Println("图片下载失败：", err1)
			continue
		}
		files.Write(data)
	}*/
}

func downLoadImgs(imgList []string, loadSrc string) {
	count := len(imgList)
	wg := sync.WaitGroup{}
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func(tmpStr string, l string) {
			defer wg.Done()
			pos := strings.LastIndex(tmpStr, "/")
			fileName := string(tmpStr[pos : len(tmpStr)-1])

			//得到准备创建的文件url
			file, err := os.Create(l + fileName)
			defer file.Close()
			if err != nil {
				fmt.Println("Error for create file.")
				fmt.Println(err)
				os.Exit(1)
			}

			//准备下载图片
			data, err1 := downloadImg("http:" + tmpStr[5:len(tmpStr)-1])
			if err1 != 200 {
				fmt.Println("图片下载失败：", err1, "地址：", l, fileName)
				//				os.Exit(1)
			}
			file.Write(data)
		}(imgList[i], loadSrc)
	}
	wg.Wait()
}

//图片字节下载
func downloadImg(url string) (content []byte, statusCode int) {
	url = strings.Replace(url, "s.", ".", -1)
	resp, err1 := http.Get(url)
	if err1 != nil {
		statusCode = -100
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		url = strings.Replace(url, ".jpg", ".png", -1)
		resp, err1 = http.Get(url)
		if err1 != nil {
			statusCode = -100
			return
		}
	}
	content, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		statusCode = -200
		return
	}
	statusCode = resp.StatusCode
	return
}

func work(url string) {
	fmt.Println("准备要爬取的URL:", url)
	var threads = findThreads(url)

	for _, v := range threads {
		(&v).getContents().getImages().download()
	}
}

func mains() {
	//	pages := []string{"","2", "3", "4", "5", "6", "7", "8", "9", "10"}
	pages := []string{""}

	for _, v := range pages {
		work("http://boards.4chan.org/asp/" + v + "/")
	}
}
