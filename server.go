package main

import (
	"fmt"
	"flag"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type CustomHandler struct {
	Debug bool
}

func (hl *CustomHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	client := http.DefaultClient
	rpURL, err := url.Parse(req.URL.String())
	if err != nil {
		log.Fatal("URL Error %s", err)
	}
	rpURL.Scheme = "http"
	rpURL.Host = req.Host

	req2 := &http.Request{}
	req2.Method = req.Method
	req2.URL = rpURL
	req2.Proto = req.Proto
	req2.ProtoMajor = req.ProtoMajor
	req2.ProtoMinor = req.ProtoMinor
	req2.Header = req.Header
	req2.Body = req.Body
	req2.ContentLength = req.ContentLength
	req2.TransferEncoding = req.TransferEncoding
	req2.Host = req.Host
	req2.Form = req.Form
	req2.PostForm = req.PostForm
	req2.MultipartForm = req.MultipartForm
	req2.Trailer = req.Trailer
	req2.RemoteAddr = req.RemoteAddr
	req2.TLS = req.TLS

	response, err := client.Do(req2)
	defer response.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		log.Fatal(err)
	}
	if (hl.Debug) {
		log.Printf("%q", dump)
	}

	for key, vals := range response.Header {
		for _, val := range vals {
			rw.Header().Set(key, val)
		}
	}
	io.Copy(rw, response.Body)

}

func main() {
	var host = flag.String("host", "localhost", "Host to bind")
	var port = flag.Int("port", 8080, "Port to bind")
	var debug = flag.Bool("debug", false, "Debug")
	flag.Parse()

	handler := &CustomHandler{}
	handler.Debug = *debug
	server := &http.Server {
		Addr: fmt.Sprintf("%s:%d", *host, *port),
		Handler: handler,
	}
	defer server.Close()

	log.Fatal(server.ListenAndServe())
}
