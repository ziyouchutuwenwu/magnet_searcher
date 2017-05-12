package main

import (
	"fmt"
	"github.com/shiyanhui/dht"
	"magnet_searcher/crawler"
	"magnet_searcher/downloader"
	"magnet_searcher/link"
	"magnet_searcher/parser"
)

func main() {
	wire := dht.NewWire()

	dhtDowdloader := downloader.NewDownloader(wire)
	dhtDowdloader.OnParseFinish = onParseFinished
	dhtDowdloader.Prepare()

	dhtCrawler := crawler.NewCrawler(wire)
	dhtCrawler.Run()
}

func onParseFinished(bt *parser.BitTorrent) {
	fmt.Println("**********************************************************************")

	fmt.Println("magnet link is", link.MagnetLink(bt.InfoHash))

	fmt.Println("name", bt.Name)
	fmt.Println("infoHash", bt.InfoHash)
	fmt.Println("total length", bt.Length)

	for _, file := range bt.Files {

		for index, path := range file.Path {
			fileName := path.(string)
			fmt.Println("loop files index", index, "path name", fileName)
		}
		fmt.Println("loop files length", file.Length)
	}
	fmt.Println("**********************************************************************")
}
