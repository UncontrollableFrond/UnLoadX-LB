package main

import (
  "net/http"
  "net/url"
  "encoding/json"
  "loadbalancer"
  "io"
  "log"
)

func updateIpTables(w http.ResponseWriter, r *http.Request) {
  // parse the JSON array of IP/ports from the request body
  type message struct {
    Ip, Port, Application string
  }

  var servers []message
  dec := json.NewDecoder(r.Body)
  for {
    if err := dec.Decode(&servers); err == io.EOF {
      break
    } else if err != nil {
      log.Fatal(err)
    }
  }

  serverURLs := make([]url.URL, 0)
  serverPointers := make([]*url.URL, 0)

  for _, element := range servers {
    server := url.URL{
      Scheme: "http",
      Host: element.Ip + ":" + element.Port,
    }
    serverURLs = append(serverURLs, server)
    serverPointers = append(serverPointers, &server)
  }

  // start the load balancer, passing in the array
  // this works, but for some reason it causes the above call to
  // WriteHeader to be ignored
  loadbalancer.LoadBalance(loadbalancer.RoundRobin, serverPointers)
  w.WriteHeader(http.StatusOK)

}


// listens for a POST request of IPs and ports from the API server
func main() {
  http.HandleFunc("/iptables", updateIpTables)
  http.ListenAndServe(":9000", nil)
}
