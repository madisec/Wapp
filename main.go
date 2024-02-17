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
	url_flag := flag.String("u", "", "Get url for run tool")
	silent := flag.Bool("silent", false, "Enable silent mnode")
	flag.Parse()
	if *url_flag != "" && !*silent {
		fig()
		wapp(*url_flag)
	} else if *url_flag != "" {

		wapp(*url_flag)
	} else {
		figure.NewFigure("Wappalizer", "larry3d", true).Print()
		fmt.Println("\tWappalizer tool - Powered by MadiSec")
		fmt.Println("\tVersion: 2.0.1")
		fmt.Println("\tFor using script run script with -h flag")
		fmt.Println("")

	}
}

func fig() {
	figure.NewFigure("Wappalizer", "larry3d", true).Print()
	fmt.Println("\tWappalizer tool - Powered by MadiSec")
	fmt.Println("")
}
func wapp(url string) {
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	data, _ := ioutil.ReadAll(resp.Body) // Ignoring error for example

	wappalyzerClient, err := wappalyzer.New()
	fingerprints := wappalyzerClient.Fingerprint(resp.Header, data)
	fmt.Printf("%v\n", fingerprints)
}
