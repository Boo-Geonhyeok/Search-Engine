package scraper

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/Boo-Geonhyeok/playground/models"
)

func GetBookList(ch chan string) {
	for i := 0; i < 10; i++ {
		pageNum := strconv.Itoa(i + 1)
		fmt.Println("Scraping Page:" + pageNum)

		req, err := http.NewRequest("GET", "https://product.kyobobook.co.kr/api/gw/pdt/category/all?page="+pageNum+"&per=50&saleCmdtDvsnCode=KOR&saleCmdtClstCode=0101&sort=new", nil)
		if err != nil {
			log.Fatal(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bookList := &models.BookList{}
		err = json.Unmarshal(data, bookList)
		if err != nil {
			log.Fatal(err)
		}
		writeBookDetail(*bookList)
		ch <- "Page " + pageNum + " Done"
	}
	close(ch)
}

func writeBookDetail(bookList models.BookList) {
	file, err := os.OpenFile("book.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()

	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	for _, book := range bookList.Data.TabContents {
		NbookDetail := book.InbukCntt
		bookDetail := strings.ReplaceAll(NbookDetail, "\n", "")
		row := []string{book.CmdtName, book.SaleCmdtID, bookDetail}
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

	// run bookMorpheme.py
	c := exec.Command("python3", "./python/bookMorpheme.py")
	if err := c.Run(); err != nil {
		log.Fatal("Can't run bookMorpheme.py ", err)
	}
}
