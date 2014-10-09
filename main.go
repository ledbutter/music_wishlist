package main

import (
	//	"net/http"
	"bufio"
	"fmt"
	//	"github.com/ledbutter/newmath"
	"github.com/ledbutter/musicSource"
	"os"
	//	"strings"
	//	"io/ioutil"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What band are you interested in: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading stdin")
	}
	fmt.Println(text)
	var crit musicSource.SearchCriteria
	crit.ArtistName = text[:len(text)-1]
	mbs := new(musicSource.MusicBrainzSearcher)
	mbs.Criteria = crit
	result := mbs.Search()
	fmt.Println(result)
	//crit := musicSource.SearchCriteria(text, "")
	//fmt.Println(crit)
	//temp := newmath.Sqrt(2)
	//fmt.Println(temp)
	/*crit := Criteria(artistName, "")
	mbs := make(MusicBrainzSearcher)
	mbs.criteria = crit
	result = mbs.search()
	fmt.println(result)*/
	/*resp, err := http.Get("http://www.google.com")
	if err != nil {
		println("Uh Oh")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	contents := string(body)
	println(contents)
	*/
	//http.HandleFunc("/", hello)
	//http.ListenAndServe(":8080", nil)
}

/*func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}*/
