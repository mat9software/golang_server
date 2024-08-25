package main

import (
 //"fmt"
 "log"
 "net/http"
 //"os"
 "encoding/json"
 "io"
 "fmt"
)

// --------------------------------------------------
// MAIN
// --------------------------------------------------
func main() {
  HttpServerInit()
  HttpServerStart()
}

// --------------------------------------------------
// HTTP SERVER
// --------------------------------------------------

// --------------------------------------------------
// Setup
func HttpServerInit() {
  fs := http.FileServer(http.Dir("public"))
  http.Handle("/", fs)

  http.HandleFunc("/test-json", TestHandlerJson)
  http.HandleFunc("/test-proxy", TestHandlerProxy)
}

func HttpServerStart() {
 httpPort := "8888"

 log.Print("Listening on port " + httpPort)
 log.Fatal(http.ListenAndServe(":" + httpPort, nil))
}

// --------------------------------------------------
// Handlers
type Response struct {
    Message string `json:"message"`
}

func TestHandlerJson(w http.ResponseWriter, r *http.Request) {
    response := Response{Message: "Hello, World!"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func TestHandlerProxy(w http.ResponseWriter, r *http.Request) {
    // Create a new request based on the incoming request
    fmt.Println("Debug: " + r.URL.Path)
    //mdtmp req, err := http.NewRequest(r.Method, "https://target-server.com"+r.URL.Path, r.Body)
//mdtmp https://query2.finance.yahoo.com/v8/finance/chart/AAPL;?period1=1724176800&period2=1724349600&interval=1m
    req, err := http.NewRequest(r.Method, "https://query2.finance.yahoo.com/v8/finance/chart/AAPL;?period1=1724176800&period2=1724349600&interval=1m", r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Copy headers from the original request
    for name, values := range r.Header {
        for _, value := range values {
            req.Header.Add(name, value)
        }
    }

    // Send the request to the target server
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    // Copy the response headers and body to the original response
    for name, values := range resp.Header {
        for _, value := range values {
            w.Header().Add(name, value)
        }
    }
    w.WriteHeader(resp.StatusCode)
    io.Copy(w, resp.Body)
}
