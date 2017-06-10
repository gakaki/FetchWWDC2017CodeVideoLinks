package fetchAppleWWDC2017

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"runtime"
	"strings"
	"sync"
)

func exportVideosData() {

	videos := readJsonAndDeserialize("output_detail.json")

	//所有sd link
	var allSDLinks []string
	for _, v := range videos {
		if v.VideoSD != "" {
			allSDLinks = append(allSDLinks, v.VideoSD)
		}
	}

	//所有hd link
	var allHDLinks []string
	for _, v := range videos {
		if v.VideoHD != "" {
			allHDLinks = append(allHDLinks, v.VideoHD)
		}
	}
	//所有resource link
	var allResourcesLink []string
	for _, v := range videos {
		for _, w := range v.Resources {
			if w.URL != "" && w.Type != "link" {
				//println(v.ID, v.Title, w.Type, w.URL)
				allResourcesLink = append(allResourcesLink, w.URL)
			}
			//else {
			//	println(w.Type, "", w.URL)
			//}
		}
		println("\n", getUrlFileName(v.VideoSD))
	}

	//最后txt 写入
	writeLines(allSDLinks, "links_sd.txt")
	writeLines(allHDLinks, "links_hd.txt")
	writeLines(allResourcesLink, "links_resources.txt")

	print("all sd links count is ", len(allSDLinks))
}

func batchFetchVideoDetails() []Video {
	//json 解析之后
	videos := readJsonAndDeserialize("output.json")
	//go chan 并发
	maxWorkerCount := 20
	queue := make(chan Video, maxWorkerCount)
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := sync.WaitGroup{}

	var videosNew []Video
	for i := 0; i < maxWorkerCount; i++ {
		go func() {
			defer wg.Done()
			wg.Add(1)
			for v := range queue {
				v = fetchVideoDetail(v)
				fmt.Println(v.VideoSD)
				videosNew = append(videosNew, v)

			}
		}()
	}

	for _, v := range videos {
		queue <- v
	}
	close(queue)
	wg.Wait()

	writeToJSON(videosNew, "output_detail.json")
	return videosNew
}
func fetchVideoDetail(v Video) Video {

	url := v.DetailUrl
	doc, e := getContentFromUrl(url)

	if e != nil {
		fmt.Fprintf(os.Stderr, ">>>>>>>network Error: %s\n", e)
		return Video{}
	}

	v.Desc = doc.Find(".details p").Eq(0).Text()

	link_node := doc.Find(".links").Eq(0)

	link_node.Find("li.video a").Each(func(j int, node *goquery.Selection) {
		href := node.AttrOr("href", "")

		if strings.Contains(href, "_hd_") {
			v.VideoHD = href
		}
		if strings.Contains(href, "_sd_") {
			v.VideoSD = href
		}
	})

	link_node.Find("li.document,li.download").Each(func(j int, node *goquery.Selection) {
		documentA := node.Find("a").Eq(0)
		href := documentA.AttrOr("href", "")
		text := documentA.Text()

		resource := Resource{}
		resource.Title = text
		resource.URL = href

		var typeS = "link"

		if strings.Contains(href, ".pdf") {
			typeS = "pdf"
		} else if strings.Contains(href, ".zip") {
			typeS = "code"
		} else {
			typeS = "link"
		}
		resource.Type = typeS

		v.Resources = append(v.Resources, resource)
	})

	return v
}
