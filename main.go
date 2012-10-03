package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/elazarl/goproxy"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getConfig(configFile string) [][]string {
	file, e := os.Open(configFile)
	if e != nil {
		panic(e)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, e := csvReader.ReadAll()
	if e != nil {
		panic(e)
	}

	return records
}

func NewResponse(r *http.Request, contentType string, status int, contentFileName string) *http.Response {
	resp := &http.Response{}
	resp.Request = r
	resp.TransferEncoding = r.TransferEncoding
	resp.Header = make(http.Header)
	resp.Header.Add("Content-Type", contentType)
	resp.StatusCode = status

	b, e := ioutil.ReadFile(contentFileName)
	if e != nil {
		panic(e)
	}
	reader := bytes.NewReader(b)
	resp.ContentLength = int64(reader.Len())
	resp.Body = ioutil.NopCloser(reader)
	return resp
}

func CreateDefault(listen string, configFile string, verbose bool) {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = verbose

	records := getConfig(configFile)
	fmt.Println("Configuration from", configFile, ":")
	for i, config := range records {
		if i == 0 {
			// header line
			continue
		}
		defaultUrl := config[0]
		contentType := config[1]
		newFile := config[2]
		fmt.Println("- replace", defaultUrl, "by", newFile)
		proxy.OnRequest(goproxy.UrlIs(defaultUrl)).DoFunc(
			func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
				fmt.Println("Intercept", defaultUrl)
				response := NewResponse(r, contentType, http.StatusOK, newFile)
				return r, response
			})
	}
	fmt.Println("")
	fmt.Println("Listen", listen)
	fmt.Println("^C to stop")
	fmt.Println("")

	log.Fatal(http.ListenAndServe(listen, proxy))
}

func main() {
	var addr string
	var configFile string
	var verbose bool
	var help bool
	flag.StringVar(&addr, "a", ":8888", "Bind to this address:port")
	flag.StringVar(&configFile, "c", "config.csv", "Config file")
	flag.BoolVar(&verbose, "v", false, "Verbose")
	flag.BoolVar(&help, "h", false, "Print this help")

	flag.Parse()
	if help {
		flag.Usage()
		return
	}

	CreateDefault(addr, configFile, verbose)
}
