package proxyserver

import (
	"io"
	"net/http"
)

func (ps *ProxyServer) StartServer() error {
	http.Handle("/", ps)

	if err := http.ListenAndServe(ps.addr, nil); err != nil {
		ps.logger.Fatal(err)
		return err
	}

	return nil
}

func (ps *ProxyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestC := *r
	requestC.RequestURI = ""

	ps.logger.Println("Incoming request for : ", requestC.URL.String())

	res, err := ps.dialer.Do(&requestC)

	if err != nil {
		ps.logger.Println("ps client Do err ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(res.StatusCode)

	copyHeaders(w, &res.Header)

	resBody := res.Body
	defer resBody.Close()
	if _, err := io.Copy(w, resBody); err != nil {
		ps.logger.Printf("copy error:%v\n", err)
	}
}

func copyHeaders(dst http.ResponseWriter, src *http.Header) {
	for k, vs := range *src {
		for _, v := range vs {
			dst.Header().Add(k, v)
		}
	}
}
