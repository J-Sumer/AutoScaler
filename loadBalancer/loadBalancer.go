package loadBalancer

import (
	"fmt"
	"encoding/json"
	// "io/ioutil"
	"os"
	"log"
	// "net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	// "time"
)

type Config struct {
    Proxy    Proxy     `json:"proxy"`
    Backends []Backend `json:"backends"`
}

type Proxy struct {
    Port string `json:"port"`
}

type Backend struct {
    URL    string `json:"url"`
    IsDead bool
    mu     sync.RWMutex
}

var cfg Config

var mu sync.Mutex
var idx int = 0

// lbHandler is a handler for loadbalancing
func lbHandler(w http.ResponseWriter, r *http.Request) {
    maxLen := len(cfg.Backends)
    // Round Robin
    mu.Lock()
    // currentBackend := cfg.Backends[idx%maxLen]
    targetURL, err := url.Parse(cfg.Backends[idx%maxLen].URL)
    if err != nil {
        log.Fatal(err.Error())
    }
    idx++
    mu.Unlock()
    reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)
    reverseProxy.ServeHTTP(w, r)
}

func Serve() {
	data, err := os.ReadFile("./config.json")
    if err != nil {
		fmt.Println("Error while reading configuration file")
        log.Fatal(err.Error())
    }
    json.Unmarshal(data, &cfg)

    // director := func(request *http.Request) {
    //     request.URL.Scheme = "http"
    //     request.URL.Host = ":8081"
    // }

    // rp := &httputil.ReverseProxy{
    //     Director: director,
    // }

    // s := http.Server{
    //     Addr:    ":8080",
    //     Handler: rp,
    // }

	s := http.Server{
        Addr:    ":" + cfg.Proxy.Port,
        Handler: http.HandlerFunc(lbHandler),
    }

    if err := s.ListenAndServe(); err != nil {
        log.Fatal(err.Error())
    }
}
