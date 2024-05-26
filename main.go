package main

import (
	"fmt"
	"os"

	"github.com/kaklikOf13/KLL/kll"
)

func PrintHelp(command string) {
	switch command {
	case "interpret":
		fmt.Println(`
interpret
	- This Command Is Used For Interpret Code. you can use better errors manipulations`)
	case "version":
		fmt.Println(`
version
	- Use this to show current version of kll`)
	case "tokenizer":
		fmt.Println(`
tokenizer
	- Use this To show tokens of a file
	- Tokens Is the words
	- Like myword`)
	default:
		fmt.Println(`
Commands:
	- interpret <mainFile> // Use this command to interpret a code.
	- tokenizer <script>  // Use this To show tokens of a file
	- help <command> // Use this command to show more of other commands
	- version // Use this to show current version of kll`)
		fmt.Println()
	}

}

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "help":
			if len(os.Args) > 2 {
				PrintHelp(os.Args[2])
			} else {
				PrintHelp("")
			}
		case "tokenizer":
			if len(os.Args) > 2 {
				f, err := os.ReadFile(os.Args[2])
				if err != nil {
					fmt.Println(err)
				}
				toks, kllerr := kll.Tokenizer(string(f))
				if !kllerr.Is(kll.Success) {
					fmt.Println(kllerr)
				}
				line := 1
				show := fmt.Sprintf("Line %v:", line)
				for _, tok := range toks {
					switch tok.Type {
					case kll.TT_NEWLINE:
						line++
						show += "\n" + fmt.Sprintf("Line %v:", line)
					default:
						show += " " + tok.String()
					}
				}
				fmt.Println(show)
			} else {
				fmt.Println("Dont Have A File")
			}
		default:
			PrintHelp("")
		}
	} else {
		PrintHelp("")
	}
}
