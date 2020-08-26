package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

var (
	// mGet  = "GET"
	mPost = "POST"
)

type mux struct {
	handlers map[string]func(http.ResponseWriter, *http.Request)
}

func newMux() *mux {
	return &mux{
		handlers: make(map[string]func(http.ResponseWriter, *http.Request)),
	}
}

func (m *mux) GET(path string, handler http.HandlerFunc) {
	m.handlers[fmt.Sprintf("%s:%s", mGet, path)] = handler
}

func (m *mux) POST(path string, handler http.HandlerFunc) {
	m.handlers[fmt.Sprintf("%s:%s", mPost, path)] = handler
}

func (m *mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f, ok := m.handlers[fmt.Sprintf("%s:%s", r.Method, r.URL.Path)]
	if !ok {
		notFound(w)
		return
	}

	f(w, r)
}

func notFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func getHostName() string {
	name, err := os.Hostname()
	if err != nil {
		return ""
	}
	return name
}

func main() {
	var addr string
	flag.StringVar(&addr, "addr", ":3000", "Address of HTTP server")
	flag.Parse()

	router := newMux()
	router.GET("/whoami", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("GET: /whoami - Client: %s\n", r.RemoteAddr)
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		response := map[string]interface{}{
			"time": time.Now(),
			"host": getHostName(),
			"ip":   getLocalIP(),
		}
		json.NewEncoder(w).Encode(response)
	})

	router.POST("/whoami", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("POST: /whoami - Client: %s\n", r.RemoteAddr)
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		response := map[string]interface{}{
			"time": time.Now(),
			"host": getHostName(),
			"ip":   getLocalIP(),
		}
		json.NewEncoder(w).Encode(response)
	})

	server := http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         addr,
		Handler:      router,
	}

	log.Println("HTTP server is serving at: ", addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
