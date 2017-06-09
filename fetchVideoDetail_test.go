package fetchAppleWWDC2017

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestFetchVideoDetail(t *testing.T) {
	//var url = "https://developer.apple.com/videos/play/wwdc2017/709/" //单个sku页面

	buf, err := ioutil.ReadFile("output.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	var videos []Video
	json.Unmarshal(buf, &videos)
	fmt.Println(videos, err)

	videos[12].DetailUrl = "https://developer.apple.com/videos/play/wwdc2017/602/"
	v := fetchVideoDetail(videos[12])
	fmt.Println(v)
}

func TestBatchFetchVideoDetails(t *testing.T) {
	batchFetchVideoDetails()



}
