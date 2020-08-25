package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
)

var markets = []string{"en-US", "zh-CN", "ja-JP", "en-AU", "en-UK", "de-DE", "fr-FR", "en-NZ", "en-CA"}

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println("bing wallpapaer")
	dir := flag.String("dir", path.Join(usr.HomeDir, "Pictures/BingWallpaper"), "set place to store downloaded images")
	if _, err := os.Stat(*dir); os.IsNotExist(err) {
		fmt.Printf("Error: path %s is not exist \n", *dir)
		fmt.Println("Creating " + *dir)
		err = os.MkdirAll(*dir, 0755)
		if err != nil {
			fmt.Println("create dir fail")
			return
		}
	}
	for i := range make([]int, 1) {
		imgPrefix := getImagePrefix(getXMLURL(i, ""))
		downloadWallpaper(imgPrefix, *dir)
	}
}

func getXMLURL(idx int, mkt string) string {
	if len(mkt) == 0 {
		mkt = "zh-CN"
	}
	return fmt.Sprintf("http://www.bing.com/HPImageArchive.aspx?format=xml&idx=%d&n=1&mkt=%s", idx, mkt)
}

func getImagePrefix(xmlURL string) string {
	resp, err := http.Get(xmlURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(body))
	re := regexp.MustCompile("<urlBase>(.*)</urlBase>")
	// fmt.Println(re.FindStringSubmatch(string(body)))
	return "http://bing.com" + re.FindStringSubmatch(string(body))[1]
}

// download the highest resolution
func downloadWallpaper(imgPrefix string, dir string) string {
	file := ""
	// resolutions := []string{"_1920x1200", "_1920x1080", "_1366x768", "_1280x768", "_1280x720", "_1024x768"}
	resolutions := []string{"_1920x1080"}

	for _, res := range resolutions {
		url := imgPrefix + res + ".jpg"

		fmt.Println("upstream uri:" + url)

		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		contentLength, _ := strconv.Atoi(resp.Header.Get("Content-Length"))

		// bing will not return 301 for redirect
		if resp.StatusCode == 200 && urlChk(resp, url) {
			file = filepath.Join(dir, filepath.Base(url))
			if _, err := os.Stat(file); os.IsNotExist(err) {
				out, err := os.Create(file)
				errChk(err)

				_, err = io.Copy(out, resp.Body)
				errChk(err)

				out.Sync()
				out.Close()

				if imageChk(file, contentLength) {
					fmt.Println("downloaded to:" + file)
					break
				} else {
					err = os.Remove(file)
					errChk(err)
					file = ""
					continue
				}
			} else {
				if imageChk(file, contentLength) {
					break
				} else {
					err = os.Remove(file)
					errChk(err)
					file = ""
					continue
				}
			}
		}
	}
	return file
}

func urlChk(resp *http.Response, uri string) bool {
	return uriPath(resp.Request.URL.String()) == uriPath(uri)
}
func uriPath(uri string) string {
	re := regexp.MustCompile(`http(s)?:\/\/[^\/]+(.*)`)
	if re.MatchString(uri) {
		return re.FindStringSubmatch(uri)[2]
	}
	return ""
}

func errChk(err error) {
	if err != nil {
		panic(err)
	}
}

func imageChk(image string, length int) bool {
	re := regexp.MustCompile(`^image/`)
	out, err := exec.Command("/usr/bin/file", "-L", "--mime-type", "-b", image).Output()
	errChk(err)
	info, err := os.Stat(image)
	errChk(err)
	return re.MatchString(string(out)) && info.Size() == int64(length)
}
