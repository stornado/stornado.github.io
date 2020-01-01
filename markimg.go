// 下载markdown中图片到本地
// 并替换图片链接为本地地址

package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
)

const MARKDWON_IMAGE_URL_PATTERN = `!\[(.+?)\]\((.+?)(\s+'(.+?)'|\s+"(.+?)")?\)`

func genMd5(text string) (md5hash string) {
	m := md5.New()
	io.WriteString(m, text)
	return hex.EncodeToString(m.Sum(nil))
}

func getFilename(filepath string) (filename string) {
	filenameWithExt := path.Base(filepath)
	return strings.TrimSuffix(filenameWithExt, path.Ext(filenameWithExt))
}

func download(url string, dir string) (host, title, dst string, err error) {
	local := path.Join(dir, genMd5(url)+path.Ext(url))
	if _, err = os.Stat(local); os.IsExist(err) {
		log.Fatalf("%s aready saved at %s\n", url, local)
		return
	}

	log.Printf("fetching %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("fetch %s %v\n", url, err)
		return
	}

	defer resp.Body.Close()

	f, err := os.Create(local)
	if err != nil {
		log.Fatalf("create local file %v\n", err)
		return
	}
	_, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}

	host = strings.ToUpper(strings.TrimPrefix(resp.Request.Host, "www."))
	title = getFilename(url)
	dst = local
	return
}

type MarkdownImgUrl struct {
	MarkImg string
	ImageUrl
}

type ImageUrl struct {
	AltText string
	Url     string
	Title   string `json:"omitempty"`
}

func findImgUrl(filepath string) (imgUrls []MarkdownImgUrl, err error) {
	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("open markdwon %s %s\n", filepath, err)
		return
	}

	content := string(f)

	imgUrlRe := regexp.MustCompile(MARKDWON_IMAGE_URL_PATTERN)
	urls := imgUrlRe.FindAllStringSubmatch(content, -1)

	for _, url := range urls {
		imgUrls = append(imgUrls, MarkdownImgUrl{url[0], ImageUrl{AltText: url[1], Url: url[2], Title: url[3]}})
		log.Printf("find %s %s %s", url[1], url[2], url[3])
	}

	log.Println(imgUrls)
	return
}

func replaceMarkdownImgUrl(filepath string, replace map[string]string) (output []byte, err error) {
	f, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("open markdwon %s %s\n", filepath, err)
		return
	}
	defer f.Close()
	reader := bufio.NewReader(f)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return nil, err
		}
		for mkUrl, hexoImgUrl := range replace {
			if ok, _ := regexp.Match(mkUrl, line); ok {
				reg := regexp.MustCompile(mkUrl)
				newByte := reg.ReplaceAll(line, []byte(hexoImgUrl))
				output = append(output, newByte...)
				output = append(output, []byte("\n")...)
			} else {
				output = append(output, line...)
				output = append(output, []byte("\n")...)
			}
		}
	}
	return
}

func writeToFile(filepath string, output []byte) (err error) {
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		return
	}
	writer := bufio.NewWriter(f)
	_, err = writer.Write(output)
	if err != nil {
		return
	}
	writer.Flush()
	return
}

var markdownFilepath = flag.String("markdown", "", "输入markdown文件绝对路径")
var imageSaveDir = flag.String("dest", "", "请输入图片保存目录(默认为markdown同名目录)")
var replaceWtihHexoAssetImg = flag.Bool("replace", false, "是否使用HEXO {% asset_img %}替换![]()")

func main() {

	flag.Parse()

	if *markdownFilepath == "" {
		flag.Usage()
		log.Fatalf("\n\n请输入markdown文件地址")

	}

	if _, err := os.Stat(*markdownFilepath); os.IsNotExist(err) {
		log.Fatalln("%s 不存在\n", *markdownFilepath)
	}

	if *imageSaveDir == "" {
		assetDir, _ := path.Split(*markdownFilepath)
		*imageSaveDir = path.Join(assetDir, getFilename(*markdownFilepath))
	}

	if _, err := os.Stat(*imageSaveDir); os.IsNotExist(err) {
		//log.Fatalf("%s 不存在\n", *imageSaveDir)
		err := os.MkdirAll(*imageSaveDir, os.ModeDir)
		if err != nil {
			log.Fatalf("创建 %s %v\n", *imageSaveDir, err)
		}
	}

	imgUrls, err := findImgUrl(*markdownFilepath)
	if err != nil {
		log.Fatalln(err)
	}

	if len(imgUrls) == 0 {
		log.Fatalln("%s 无图片需要下载替换")
	}

	ch := make(chan MarkdownImgUrl)

	for _, imgUrl := range imgUrls {
		go func(imgUrl MarkdownImgUrl) {
			ch <- imgUrl
		}(imgUrl)
	}

	saved := make(map[string]string)

	for imgUrl := range ch {
		go func(imgUrl MarkdownImgUrl) {
			host, title, dst, err := download(imgUrl.Url, *imageSaveDir)
			if err != nil {
				log.Fatalln(err)
			}

			saved[imgUrl.MarkImg] = fmt.Sprintf("{%% asset_img %s From: %s %%}", path.Base(dst), host)

			log.Printf("saved %s %s %s %s %s %s\n", imgUrl.AltText, imgUrl.Url, imgUrl.Title, host, title, dst)
		}(imgUrl)
	}
	close(ch)

	if *replaceWtihHexoAssetImg && len(saved) > 0 {
		output, err := replaceMarkdownImgUrl(*markdownFilepath, saved)
		if err != nil {
			log.Fatalln(err)
		}

		err = writeToFile(*markdownFilepath, output)
		if err != nil {
			log.Fatalln(err)
		}
	}

	log.Println("all done")
}
