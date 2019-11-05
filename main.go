package main
import "github.com/blais0u/username/twitter"
import "github.com/blais0u/username/github"
import "fmt"

func main() {
	fmt.Println(twitter.Validate("Blais0u"))

	fmt.Println(github.Validate("Blais0u"))
}
