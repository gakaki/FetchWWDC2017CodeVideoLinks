package fetchAppleWWDC2017

import (
	"fmt"
	"testing"
)
func TestBatchFetchVideoDetails(t *testing.T) {
	batchFetchVideoDetails()

}
func TestFetchVideoDetail(t *testing.T) {

	var videos = readJsonAndDeserialize("output.json")
	videos[12].DetailUrl = "https://developer.apple.com/videos/play/wwdc2017/602/"
	v := fetchVideoDetail(videos[12])
	fmt.Println(v)
}


func TestPrintDetails(t *testing.T) {
	exportVideosData()

}
