package main

import (
	"bufio"
	"fmt"
	"github.com/ledbutter/musicSource"
	"os"
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
}
