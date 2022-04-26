package main

import (
	"fmt"
	"math/rand"
	"time"

	tm "github.com/buger/goterm"
	randomUserAgent "github.com/corpix/uarand"
	"github.com/valyala/fasthttp"
)

var (
	client fasthttp.Client

	errors  int    = 0
	sent    int    = 0
	rpm     int    = 0
	itemID  string = "7077018440971685162"
	threads int    = 1000
)

func addShare(itemID string) {
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	req.Header.SetMethod("POST")
	req.SetRequestURI(generateURL())
	req.SetBody([]byte(fmt.Sprintf("item_id=%v&share_delta=1", itemID)))

	req.Header.Set("User-Agent", randomUserAgent.GetRandom())
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	if err := client.Do(req, res); err != nil {
		errors++
		return
	}

	sent++
}

func rpmCounter() {
	for {
		before := sent
		time.Sleep(100 * time.Millisecond)
		after := sent

		rpm = (after - before) * 600
	}
}

func statusPrinter() {
	for {
		fmt.Printf("[+] Sent: %v | Requests per minute: %v | Errors: %v\r", sent, rpm, errors)
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	tm.Clear()

	fmt.Print("[+] github.com/qoft\n\n")

	fmt.Print("\n\n")

	go rpmCounter()
	go statusPrinter()

	for i := 0; i < threads; i++ {
		go func() {
			for {
				addShare(itemID)
			}
		}()
	}

	select {}
}
