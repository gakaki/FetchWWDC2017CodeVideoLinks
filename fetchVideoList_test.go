package fetchAppleWWDC2017

import (
	"testing"
)

func TestFetchVideoList(t *testing.T) {

	videos := fetchVideoList()
	print(videos)

}
