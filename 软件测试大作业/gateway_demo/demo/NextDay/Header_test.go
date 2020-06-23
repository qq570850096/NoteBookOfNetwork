package NextDay

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"testing"
)

func TestGetWebsite(t *testing.T) {
	var url string
	fmt.Scanf("%s",url)
	GetWebsite("https://e-hentai.org/g/1658163/1623728b58/")
	test()
}

func test(){
	html := `<body>

				<div id="div1">DIV1</div>
				<div class="name">DIV2</div>
				<span>SPAN</span>

			</body>
			`

	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(html))
	if err!=nil{
		log.Fatalln(err)
	}

	dom.Find(".name").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}
