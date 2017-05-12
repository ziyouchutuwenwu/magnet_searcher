package downloader

import (
	"github.com/shiyanhui/dht"
	"github.com/ziyouchutuwenwu/objective_go/thread/basic_thread"
	"magnet_searcher/parser"
)

type ParseFinishedCallBack func(bt *parser.BitTorrent)

type Downloader struct {
	dhtDownloader *dht.Wire

	//回调函数
	OnParseFinish ParseFinishedCallBack
}

func (downloader *Downloader) onParse(thread *basic_thread.Thread, argObject interface{}) {

	for resp := range downloader.dhtDownloader.Response() {

		metadata, err := dht.Decode(resp.MetadataInfo)
		if err != nil {
			continue
		}
		bt := parser.Parse(metadata, resp.InfoHash)

		if nil != downloader.OnParseFinish {
			downloader.OnParseFinish(bt)
		}
	}
}

func (parser *Downloader) Prepare() {

	thread := basic_thread.Create()
	thread.Init()
	thread.Tag = "parser_thread"

	//回调自带阻塞
	thread.SetCallBack(parser.onParse)
	thread.Start()

	go parser.dhtDownloader.Run()
}

func NewDownloader(wire *dht.Wire) *Downloader {
	downloader := new(Downloader)
	downloader.dhtDownloader = wire

	return downloader
}