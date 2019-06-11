package fetchAppleWWDC2019

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

//先抓顶部分类
//在分类获取详细页链接

func fetchVideoList() (videos []Video) {

	urlApplePrefix := "https://developer.apple.com"
	//doc, e := getContentFromUrl(urlApplePrefix + "/videos/wwdc2019/")

	doc, e := getContentFromFile("./list.html")

	if e != nil {
		log.Print(e, " 出错了 系列页访问出错")
	}

	var videos_in_categories [][]Video
	//var category_map = make(map[string]string)
	doc.Find(".collection-focus-group").Each(func(i int, node *goquery.Selection) {

		category_id := node.AttrOr("id", "")
		category_title := node.Find("span.focus-group-link span.font-bold").Text()
		//category_map[category_id] = category_title

		var videos_in_category []Video
		node.Find(".collection-item").Each(func(j int, node_sub *goquery.Selection) {

			imgNode := node_sub.Find(".video-image").Eq(0)
			imageUrl := imgNode.AttrOr("src", "")

			aNode := node_sub.Find("a").Eq(1)
			detailUrl := aNode.AttrOr("href", "")
			detailUrl = urlApplePrefix + detailUrl
			title := aNode.Find("h4").Eq(0).Text()

			videoTagsNode := node_sub.Find(".video-tags li")
			sessionName := videoTagsNode.Eq(0).Text()
			tags := videoTagsNode.Eq(2).Text()
			//no fetch detail in detail page list page desc is not complete
			//desc := node_sub.Find(".description").Text()




			v := Video{}
			v.CategoryID = category_id
			v.CategoryTitle = category_title

			c := Category{}
			c.ID = category_id
			c.Title = category_title
			v.Category = c

			v.Title = title
			v.SessionName = sessionName
			v.ID = strings.Replace(sessionName, "Session ", "", 1)
			v.TAGS = tags
			v.Image = imageUrl
			v.DetailUrl = detailUrl
			//v.Desc		= desc

			videos_in_category = append(videos_in_category, v)
		})

		videos_in_categories = append(videos_in_categories, videos_in_category)
	})

	for _, videos_in_category := range videos_in_categories {
		for _, v := range videos_in_category {
			videos = append(videos, v)
		}
	}

	for _, v := range videos {
		fmt.Println(">>>>>>", v)
	}
	fmt.Println(">>>>>> 获得视频数量为:", len(videos))

	writeToJSON(videos, "output.json")

	return videos
}
