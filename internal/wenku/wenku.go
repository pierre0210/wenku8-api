package wenku

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/pierre0210/wenku8-api/internal/util"
)

type Volume struct {
	Title       string    `json:"title"`
	Vid         int       `json:"vid"`
	ChapterList []Chapter `json:"chapterList"`
}

type Chapter struct {
	Title   string   `json:"title"`
	Cid     int      `json:"cid"`
	Content string   `json:"content"`
	Urls    []string `json:"urls"`
}

func GetVolumeList(aid int) (string, string, []Volume) {
	var volumeList []Volume
	var novelTitle string
	var novelAuthor string
	c := colly.NewCollector()
	c.OnHTML("#title", func(h *colly.HTMLElement) {
		titleByte := util.GbkToUtf8([]byte(h.Text))
		twTitle := util.Simplified2TW(string(titleByte))
		novelTitle = twTitle
	})
	c.OnHTML("#info", func(h *colly.HTMLElement) {
		authorByte := util.GbkToUtf8([]byte(h.Text))
		novelAuthor = string(authorByte)
	})
	c.OnHTML("td", func(h *colly.HTMLElement) {
		if h.DOM.HasClass("vcss") {
			var vol Volume
			vid, err := strconv.Atoi(h.Attr("vid"))
			if err != nil {
				log.Println(err.Error())
				return
			}
			titleByte := util.GbkToUtf8([]byte(h.Text))
			twTitle := util.Simplified2TW(string(titleByte))
			vol.Title = twTitle
			vol.Vid = vid
			volumeList = append(volumeList, vol)
		} else if h.DOM.HasClass("ccss") && h.ChildAttr("a", "href") != "" {
			var ch Chapter
			titleByte := util.GbkToUtf8([]byte(h.ChildText("a")))
			twTitle := util.Simplified2TW(string(titleByte))
			ch.Title = twTitle
			ch.Cid, _ = strconv.Atoi(strings.Split(h.ChildAttr("a", "href"), "&cid=")[1])
			volumeList[len(volumeList)-1].ChapterList = append(volumeList[len(volumeList)-1].ChapterList, ch)
		}
	})

	err := c.Visit(fmt.Sprintf("https://www.wenku8.net/modules/article/reader.php?aid=%d", aid))
	util.ErrorHandler(err, false)

	return novelTitle, novelAuthor, volumeList
}
