package fetchAppleWWDC2017

type Video struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	SessionName string   `json:"sessionName"`
	Category    Category `json:"category"`

	CategoryID    string `json:"categoryId"`
	CategoryTitle string `json:"categoryTitle"`

	TAGS      string `json:"tags"`
	DetailUrl string `json:"detailUrl"`
	Image     string `json:"image"`
	Desc      string `json:"desc"`

	VideoSD string `json:"videoSD"`
	VideoHD string `json:"videoHD"`

	Resources []Resource `json:"resources"`
}

type Category struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
type Resource struct {
	URL   string `json:"url"`
	Title string `json:"title"`
	Type  string `json:"type"`
}
