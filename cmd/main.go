package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/blais0u/username"
	_ "github.com/blais0u/username/github"
	_ "github.com/blais0u/username/twitter"
)

const pseudo string = "blais0u91"

func main() {
	sn := username.SocialNetWorks
	var wg sync.WaitGroup
	ch := make(chan string, 20)

	//boucle sur les slice de social net
	// je boucle sur les tets validate et isavialbale
	for _, s := range sn {
		wg.Add(1)
		go check(s, &wg, ch)
	}

	for msg := range ch {
		fmt.Println(msg)
	}
	wg.Wait()
	close(ch)
}

func check(s username.SocialNetwork, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()

	validate := s.Validate(pseudo)

	if validate {
		available, err := s.IsAvailable(pseudo)

		if err != nil {
			log.Fatal("Error network")
		}
		ch <- ("Available :" + strconv.FormatBool(available))
	}
}
