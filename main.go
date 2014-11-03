package main

import (
	"bufio"
	"fmt"
	"github.com/ledbutter/musicSource"
	"github.com/ledbutter/musicStorage"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What band are you interested in: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading stdin")
	}
	var crit musicSource.SearchCriteria
	crit.ArtistName = text[:len(text)-1]
	fmt.Print("What album: ")
	text, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading stdin")
	}
	crit.AlbumName = text[:len(text)-1]
	mbs := new(musicSource.MusicBrainzSearcher)
	mbs.Criteria = crit
	result := musicSource.Album{}
	result, err = mbs.Search()
	if err != nil {
		fmt.Println("Error searching")
	}
	fmt.Println(result)
	//now push it to our storage
	storageAlbum := musicStorage.SavedAlbum{}
	storageAlbum.Key = result.Id
	storageAlbum.Title = result.Title
	storageAlbum.Artist = result.Artist
	musicStorage.AddAlbum(storageAlbum)
	fmt.Println("Saved album")

	allAlbums := musicStorage.ListAlbums()
	for _, a := range allAlbums {
		fmt.Println(a)
	}

}
