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

const (
	MARKDOWN_IMAGE_URL_PATTERN = `!\[(.+?)\]\((.+?)(\s+'(.+?)'|\s+"(.+?)")?\)`
	IMAGE_CONTENT_TYPE_PATTERN = `.*image/(\w+).*`
)

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
		log.Fatalf("%s already saved at %s\n", url, local)
		return
	}

	// log.Printf("fetching %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("fetch %s %v\n", url, err)
		return
	}

	defer resp.Body.Close()

	conentType := string(resp.Header.Get("Content-Type"))
	imgSuffixRe := regexp.MustCompile(IMAGE_CONTENT_TYPE_PATTERN)
	if path.Ext(url) == "" && conentType != "" {
		if imgSuffix := imgSuffixRe.FindStringSubmatch(conentType); imgSuffix[1] != "" {
			local += "." + imgSuffix[1]
		}
	}

	// log.Printf("%s will save at %s\n", url, local)

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

	log.Printf("%s saved at %s\n", url, dst)
	return
}

type MarkdownImgURL struct {
	MarkImg string
	ImageURL
}

type ImageURL struct {
	AltText string
	Link    string
	Title   string `json:"omitempty"`
}

func findImgURL(filepath string) (imgURLs []MarkdownImgURL, err error) {
	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("open markdown %s %s\n", filepath, err)
		return
	}

	content := string(f)

	imgURLRe := regexp.MustCompile(MARKDOWN_IMAGE_URL_PATTERN)
	urls := imgURLRe.FindAllStringSubmatch(content, -1)

	for _, url := range urls {
		imgURLs = append(imgURLs, MarkdownImgURL{url[0], ImageURL{AltText: url[1], Link: url[2], Title: url[3]}})
		log.Printf("find %s %s %s", url[1], url[2], url[3])
	}

	return
}

func replaceMarkdownImgURL(filepath string, replace map[string]string) (output []byte, err error) {
	f, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("open markdown %s %s\n", filepath, err)
		return
	}
	defer f.Close()
	reader := bufio.NewReader(f)

	replaced := false
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return output, nil
			}
			return nil, err
		}
		replaced = false
		for mkURL, assetImgURL := range replace {
			if ok, _ := regexp.Match(mkURL, line); ok {
				reg := regexp.MustCompile(mkURL)
				newByte := reg.ReplaceAll(line, []byte(assetImgURL))
				output = append(output, newByte...)
				output = append(output, []byte("\n")...)
				replaced = true
			} else if strings.Contains(string(line), mkURL) {
				newByte := []byte(strings.ReplaceAll(string(line), mkURL, assetImgURL))
				output = append(output, newByte...)
				output = append(output, []byte("\n")...)
				replaced = true
			}
		}

		if !replaced {
			output = append(output, line...)
			output = append(output, []byte("\n")...)
		}
	}
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

func saveMarkdownImg(markdownImgURLs <-chan MarkdownImgURL, quit <-chan int, saved *map[string]string) {

	for {
		select {
		case imgURL := <-markdownImgURLs:
			host, _, dst, err := download(imgURL.Link, *imageSaveDir)
			if err != nil {
				log.Fatalln(err)
			}

			(*saved)[imgURL.MarkImg] = fmt.Sprintf("{%% asset_img %s From: %s %%}", path.Base(dst), host)

			log.Printf("saved %s %s %s\n", imgURL.AltText, imgURL.Link, dst)
		case <-quit:
			return
		}
	}
}

var markdownFilepath = flag.String("markdown", "", "输入markdown文件绝对路径")
var imageSaveDir = flag.String("dest", "", "请输入图片保存目录(默认为markdown同名目录)")
var replaceWithAssetImg = flag.Bool("replace", false, "是否使用HEXO {% asset_img %}替换![]()")

func main() {

	flag.Parse()

	if *markdownFilepath == "" {
		flag.Usage()
		log.Fatalf("\n\n请输入markdown文件地址")
	}

	if _, err := os.Stat(*markdownFilepath); os.IsNotExist(err) {
		log.Fatalf("%s 不存在\n", *markdownFilepath)
	}

	if *imageSaveDir == "" {
		assetDir, _ := path.Split(*markdownFilepath)
		*imageSaveDir = path.Join(assetDir, getFilename(*markdownFilepath))
	}

	if _, err := os.Stat(*imageSaveDir); os.IsNotExist(err) {
		err := os.MkdirAll(*imageSaveDir, os.ModeDir)
		if err != nil {
			log.Fatalf("创建 %s %v\n", *imageSaveDir, err)
		}
	}

	imgURLs, err := findImgURL(*markdownFilepath)
	if err != nil {
		log.Fatalln(err)
	}

	if len(imgURLs) == 0 {
		log.Fatalf("%s 无图片需要下载替换", *markdownFilepath)
	}

	markdownImgURLs := make(chan MarkdownImgURL)
	quit := make(chan int)

	go func() {
		for _, imgURL := range imgURLs {
			markdownImgURLs <- imgURL
		}
		quit <- 0
	}()

	log.Println("waiting for download")

	saved := make(map[string]string)

	saveMarkdownImg(markdownImgURLs, quit, &saved)

	if *replaceWithAssetImg && len(saved) > 0 {
		output, err := replaceMarkdownImgURL(*markdownFilepath, saved)
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
