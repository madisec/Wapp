package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/common-nighthawk/go-figure"
	wappalyzer "github.com/projectdiscovery/wappalyzergo"
)

func main() {
	figure.NewFigure("Wappalizer tool", "rectangles", true).Print()
	fmt.Println(" 			Wappalizer tool - Powered by MadiSec")
	fmt.Println("")
	url_flag := flag.String("u", "", "Get url for run tool")
	flag.Parse()
	resp, err := http.DefaultClient.Get(*url_flag)
	if err != nil {
		log.Fatal(err)
	}
	data, _ := ioutil.ReadAll(resp.Body) // Ignoring error for example

	wappalyzerClient, err := wappalyzer.New()
	fingerprints := wappalyzerClient.Fingerprint(resp.Header, data)
	fmt.Printf("%v\n", fingerprints)

	// Output: map[Acquia Cloud Platform:{} Amazon EC2:{} Apache:{} Cloudflare:{} Drupal:{} PHP:{} Percona:{} React:{} Varnish:{}]
}
