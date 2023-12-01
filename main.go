package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/common-nighthawk/go-figure"
	wappalyzer "github.com/projectdiscovery/wappalyzergo"
)

func main() {
	figure.NewFigure("Wappalizer tool", "rectangles", true).Print()
	fmt.Println(" 			Wappalizer tool - Powered by MadiSec\n")
	resp, err := http.DefaultClient.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	data, _ := ioutil.ReadAll(resp.Body) // Ignoring error for example

	wappalyzerClient, err := wappalyzer.New()
	fingerprints := wappalyzerClient.Fingerprint(resp.Header, data)
	fmt.Printf("%v\n", fingerprints)

	// Output: map[Acquia Cloud Platform:{} Amazon EC2:{} Apache:{} Cloudflare:{} Drupal:{} PHP:{} Percona:{} React:{} Varnish:{}]
}
