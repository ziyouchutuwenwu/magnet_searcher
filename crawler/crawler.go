package crawler

import (
	"github.com/shiyanhui/dht"
)

type Crawler struct {
	downloader *dht.Wire

	config *dht.Config
	dht    *dht.DHT
}

func (crawler *Crawler) onAnnouncePeer(infoHash, ip string, port int) {
	crawler.downloader.Request([]byte(infoHash), ip, port)
}

func NewCrawler(downloader *dht.Wire) *Crawler {
	crawler := new(Crawler)

	config := dht.NewCrawlConfig()
	config.OnAnnouncePeer = crawler.onAnnouncePeer

	dhtInstance := dht.New(config)

	crawler.config = config
	crawler.dht = dhtInstance
	crawler.downloader = downloader

	return crawler
}

func (crawler *Crawler) Run() {
	crawler.dht.Run()
}
