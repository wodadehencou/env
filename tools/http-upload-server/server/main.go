package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

var port = flag.Int("port", 9090, "listen port of file upload server")

func init() {
	flag.Parse()
}

func main() {
	fmt.Printf("file upload http server, server ip is %s, listen at %d \n", GetOutboundIP().String(), *port)
	http.HandleFunc("/", uploadHandler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		panic("http server listen fail")
	}
}

var uploadTemplate = `
<html>
	<head>
		<title>Upload file</title>
	</head>
	<body>
		<form enctype="multipart/form-data" action="http://%s:%d/" method="post">
			<input type="file" name="uploadfile" />
			<input type="hidden" name="token" value="{{.}}"/>
			<input type="submit" value="upload" />
		</form>
	</body>
</html>
`

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.New("uploadfile").Parse(fmt.Sprintf(uploadTemplate, GetOutboundIP().String(), *port))
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
