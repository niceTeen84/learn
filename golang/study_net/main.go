package main

import (
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"github.com/valyala/fasthttp"
)

func SatrtServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	fasthttp.ListenAndServe(":7000", r)
}

func init() {
	client, err := api.NewClient(&api.Config{Address: "192.168.0.150:8500", Scheme: "http"})
	if err != nil {
		log.Fatal("connet to consul failed ", err.Error())
	}

	kv := client.KV()
	p := &api.KVPair{Key: "REDIS_MAXCLIENTS", Value: []byte("1000")}
	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)
	}

	// Lookup the pair
	pair, _, err := kv.Get("REDIS_MAXCLIENTS", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)
}

func GetLocalIp() net.IP {
	// use alibaba dns server
	conn, err := net.Dial("udp", "223.5.5.5:80")
	if err != nil {
		log.Fatal("can not connected to server ", err.Error())
	}

	defer conn.Close()
	local := conn.LocalAddr().(*net.UDPAddr)
	return local.IP
}

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal("query ip address failed ", err.Error())
	}

	for _, addr := range addrs {
		// fmt.Println(idx, addr)
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}

	fmt.Println(GetLocalIp())
}
