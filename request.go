package golibs

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func RequestPost(requestUrl string, params map[string]string, headers map[string]string) string {

	// パラメータを作る
	values := url.Values{}
	for key, val := range params {
		values.Add(key, val)
	}

	// リクエストを作る
	// 例えばPOSTでパラメータつけてUser-Agentを指定する
	req, err := http.NewRequest("POST", requestUrl, strings.NewReader(values.Encode()))
	FailOnError(err)

	return requestCommon(req, headers)
}

func RequestGet(requestUrl string, params map[string]string, headers map[string]string) string {

	// パラメータを作る
	paramArray := make([]string, len(params))
	for key, val := range params {
		paramArray = append(paramArray, fmt.Sprintf("%s=%s", key, url.QueryEscape(val)))
	}
	paramString := strings.Join(paramArray, "&")

	// リクエストを作る
	req, err := http.NewRequest("GET", requestUrl+"?"+paramString, nil)
	FailOnError(err)

	return requestCommon(req, headers)
}

func requestCommon(req *http.Request, headers map[string]string) string {

	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for key, val := range headers {
		req.Header.Set(key, val)
	}

	// リスエスト投げてレスポンスを得る
	client := &http.Client{}
	resp, err := client.Do(req)
	FailOnError(err)

	// レスポンスをNewDocumentFromResponseに渡してドキュメントを得る
	doc, err := goquery.NewDocumentFromResponse(resp)
	FailOnError(err)

	ret, err := doc.Html()
	FailOnError(err)

	return ret
}
