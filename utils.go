package fetchAppleWWDC2017

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func getColorId(src string) string {
	var res = ""

	pat := `[0-9]{3}`

	r, err := regexp.Compile(pat)
	if err != nil {
		fmt.Fprintf(os.Stderr, "regex compile error: %s\n", err)
		log.Fatal(err)
	}

	result_slice := r.FindStringSubmatch(src)
	if len(result_slice) > 0 {
		fmt.Println(regexp.MatchString(pat, src))
		res = result_slice[0]
		fmt.Printf("%v", res)
	}
	return res
}
func eur2Rmb(eurmoney string) string {
	if strings.Contains(eurmoney, "€") {
		/* 1欧元 =7.38035716 人民币 */
		moneystr := strings.Replace(eurmoney, "€", "", 1)
		money, err := strconv.ParseFloat(moneystr, 64)

		if err != nil {
			fmt.Fprintf(os.Stderr, "eurmoney convert error: %s\n", err)
			log.Print("出错了 注意 ", err)
			return "0"
		} else {
			res_float := money * 7.38035716
			res := fmt.Sprintf("%.2f", res_float)
			return res
		}

	} else {
		return "0"
	}
}

func getCategoryLevelShopId(url string) (string, string, string) {
	splitStrs := strings.Split(url, "/")

	var categoryLevel1 = splitStrs[4]
	var categoryLevel2 = splitStrs[5]
	var shopId = splitStrs[6]
	if len(splitStrs) < 7 {
		categoryLevel1 = splitStrs[4]
		categoryLevel2 = splitStrs[5]
		shopId = splitStrs[6]
	}

	println(categoryLevel1, categoryLevel2, shopId)
	return categoryLevel1, categoryLevel2, shopId
}

func getContentFromUrl(url string) (*goquery.Document, error) {
	debug := true

	httpClient := gorequest.New()
	httpClient.Debug = debug

	//超时2秒 并且重试三次
	request := gorequest.New().Timeout(20 * time.Second)
	resp, body, errs := request.Get(url).Retry(3, 2*time.Second, http.StatusBadRequest, http.StatusInternalServerError, http.StatusGatewayTimeout).
		End()

	if resp == nil {
		resp, body, errs = request.Get(url).Retry(3, 2*time.Second, http.StatusBadRequest, http.StatusInternalServerError, http.StatusGatewayTimeout).
			End()
	}
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
	}

	if !debug {
		log.Println(body)
	}

	//go query get the page data
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Print(err)
	}
	return doc, err
}

func getContentFromFile(filePath string) (*goquery.Document, error) {

	f, e := os.Open(filePath)
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()

	doc, e := goquery.NewDocumentFromReader(f)
	if e != nil {
		log.Fatal(e)
	}
	return doc, e
}

func AppendStringToFile(path, text string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}
