package fetchAppleWWDC2017

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"strings"
)

//先抓顶部分类
//在分类获取详细页链接

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func fetchVideoList() (videos []Video) {
	urlApplePrefix := "https://developer.apple.com"

	url := urlApplePrefix + "/videos/wwdc2017/"
	doc, e := getContentFromUrl(url)

	if e != nil {
		log.Print(e, " 出错了 系列页访问出错")
	}

	//var category_map = make(map[string]string)
	doc.Find(".collection-focus-group").Each(func(i int, node *goquery.Selection) {

		category_id := node.AttrOr("id", "")
		category_title := node.Find("span.focus-group-link span.font-bold").Text()
		//category_map[category_id] = category_title

		doc.Find(".collection-item").Each(func(i int, node *goquery.Selection) {

			imgNode := node.Find(".col-30 img").Eq(0)
			imageUrl := imgNode.AttrOr("src", "")

			aNode := node.Find(".col-70 a").Eq(0)
			detailUrl := aNode.AttrOr("href", "")
			detailUrl = urlApplePrefix + detailUrl
			title := aNode.Find("h4").Eq(0).Text()

			sessionName := node.Find(".col-70 .video-tags .event span.smaller").Eq(0).Text()
			tags := node.Find(".col-70 .video-tags .focus span.smaller").Eq(0).Text()

			c := Category{}
			c.ID = category_id
			c.Title = category_title

			v := Video{}
			v.Category = c
			v.CategoryID = category_id
			v.CategoryTitle = category_title

			v.Title = title
			v.SessionName = sessionName
			v.ID = strings.Replace(sessionName, "Session ", "", 1)
			v.TAGS = tags
			v.Image = imageUrl
			v.DetailUrl = detailUrl

			videos = append(videos, v)
		})
	})

	for _, v := range videos {
		fmt.Println(">>>>>>", v)
	}
	videosJson, _ := json.MarshalIndent(videos, "", " ")
	ioutil.WriteFile("output.json", videosJson, 0644)
	return videos
}
