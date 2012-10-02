package main

import (
	"os"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
	"encoding/csv"
	"fmt"
)

func getConfig(configFile string) ([][]string) {
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
	file, e := os.Open(contentFileName)
	if e != nil {
		panic(e)
	}
	defer file.Close()
	fileInfo, e := file.Stat()
	if e != nil {
		panic(e)
	}
	resp.ContentLength = fileInfo.Size()
	resp.Body = file
	return resp
}

func CreateDefault(listen string, configFile string, verbose bool) {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = verbose

	records := getConfig(configFile)
	fmt.Println("Configuration:")
	for i, config := range records {
		if i == 0 {
			// header line
			continue
		}
		fmt.Println("- replace", config[0], "by", config[2])
		proxy.OnRequest(goproxy.UrlIs(config[0])).DoFunc(
			func(r *http.Request, ctx *goproxy.ProxyCtx)(*http.Request, *http.Response) {
				return r,nil
			})
	}
	fmt.Println("")
	fmt.Println("^C to stop")
	fmt.Println("")

	log.Fatal(http.ListenAndServe(listen, proxy))
}

func main() {
	CreateDefault(":8888", "config.csv", false)
}
