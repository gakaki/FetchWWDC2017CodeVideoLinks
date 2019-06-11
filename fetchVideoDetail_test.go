package fetchAppleWWDC2019

import (
	"fmt"
	"testing"
)

func TestBatchFetchVideoDetails(t *testing.T) {

	batchFetchVideoDetails()

}


func TestFetchSingleVideoDetail(t *testing.T) {
	var videos = readJsonAndDeserialize("output.json")
	//https://developer.apple.com/videos/play/wwdc2019/714/
	for _, v := range videos {
		if v.ID == "714" {
			println("detail url is ", v.DetailUrl)
			v := fetchVideoDetail(v)
			fmt.Println(v)
		}
	}

}

func TestPrintDetails(t *testing.T) {

	exportVideosData()

}

func TestLastPath(t *testing.T) {
	urlstr := "https://devstreaming-cdn.apple.com/videos/wwdc/2017/501fo36iwi2moz2l222/501/501_whats_new_in_audio.pdf"
	fmt.Println(getUrlFileName(urlstr))
}
