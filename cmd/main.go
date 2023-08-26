package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/theluckiesthuman/lexiquery/internal/config"
	"github.com/theluckiesthuman/lexiquery/internal/usecase"
	"github.com/theluckiesthuman/lexiquery/internal/utils"
)

func main() {
	fmt.Println("Welcome to the Lexi Query!")

	dq := usecase.NewDictionaryQuery(config.GetCollegiateURL(), config.GetAPIKey())
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter a word (or 'exit' to quit): ")
		word, _ := reader.ReadString('\n')
		word = utils.SanitizeForURL(word)

		if word == "" {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		if word == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		definition, err := dq.Query(word)
		if err == nil {
			fmt.Println("Definition:", definition)
		} else {
			fmt.Println("Error:", err)
		}
	}
}
