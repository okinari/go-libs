package golibs

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/sclevine/agouti"
)

type WebScraping struct {
	driver *agouti.WebDriver
	page   *agouti.Page
}

func NewWebScraping() *WebScraping {

	ws := new(WebScraping)

	client := new(http.Client)
	chromeOptions := agouti.ChromeOptions(
		"args",
		[]string{
			// "--headless",
			"--disable-gpu", // 暫定的に必要らしい…?
			// "–incognito", // シークレットモード
		},
	)
	ws.driver = agouti.ChromeDriver(
		// httpクライアントを設定する(設定していない場合は、内部的にhttp.DefaultClientが使われる)
		agouti.HTTPClient(client),
		// // なんのタイムアウトなのか不明
		// agouti.Timeout(190),
		// SSL証明書が無効な場合は拒否する(デフォルトではすべて受け入れる)
		agouti.RejectInvalidSSL,
		chromeOptions,
	)

	err := ws.driver.Start()
	FailOnError(err)

	ws.page, err = ws.driver.NewPage(agouti.Browser("chrome"))
	FailOnError(err)

	// ポップアップの無効化
	err = ws.page.CancelPopup()
	FailOnError(err)

	return ws
}

func (ws *WebScraping) Close() {
	ws.driver.Stop()
}

func (ws *WebScraping) GetPage() *agouti.Page {
	return ws.page
}

func (ws *WebScraping) SetValueByID(key string, value string) {
	ws.page.FindByID(key).Fill(value)
}

func (ws *WebScraping) ExecJavaScript(script string, variable map[string]interface{}) {
	ws.page.RunScript(script, variable, nil)
}

func (ws *WebScraping) Sample(requestUrl string) string {
	ws.page.Navigate(requestUrl)
	// fmt.Println(page.HTML())

	// curContentsDom, err := ws.page.HTML()
	// FailOnError(err)

	// return curContentsDom

	// readerCurContents := strings.NewReader(curContentsDom)

	// // 現状ブラウザで開いているページのDOMを取得
	// contentsDom, err := goquery.NewDocumentFromReader(readerCurContents)
	// FailOnError(err)

	// // selector部分にセレクトボックスのセレクタを入れる。セレクトボックス内の子要素を全部取得
	// listDom := contentsDom.Find(" selector ").Children()

	// // セレクトボックスの子要素の個数を取得
	// listLen := listDom.Length()

	// for i := 1; i <= listLen; i++ {
	// 	iStr := strconv.Itoa(i)

	// 	// セレクタの属性（ここではoption）のバリューで繰り返しクリックする
	// 	ws.page.Find(" selector > option:nth-child(" + iStr + ")").Click()

	// 	//適宜ブラウザが反応するための間を取る
	// 	time.Sleep(2 * time.Second)
	// }

	// JavaScriptを実行する
	// var number int
	ws.page.RunScript("document.querySelector('#input-field-1').value='sample'", map[string]interface{}{}, nil)
	// ws.page.FindByID("input-field-1").Fill("takashi@k-wineclub.net")
	// ws.page.FindByID("input-field-2").Fill("takashi0971")
	// ws.page.FirstByClass("signin-button").Click()

	time.Sleep(10 * time.Second)

	aaa, _ := ws.page.HTML()

	return aaa
}

func (ws *WebScraping) GetFileByWeb(fileName string, url string) {

	f, err := os.Create(fileName)
	FailOnError(err)

	res, err := http.Get(url)
	FailOnError(err)
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	FailOnError(err)

	f.Write(data)
}

func (ws *WebScraping) ReadFile(filePath string) *goquery.Document {
	data, err := ioutil.ReadFile(filePath)
	stringReader := strings.NewReader(string(data))
	doc, err := goquery.NewDocumentFromReader(stringReader)
	FailOnError(err)
	return doc
}
