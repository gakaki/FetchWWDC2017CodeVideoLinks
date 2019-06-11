package fetchAppleWWDC2019

import (
	"testing"
)

func TestFetchVideoList(t *testing.T) {

	videos := fetchVideoList()
	print(videos)

}
