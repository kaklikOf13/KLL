package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kaklikOf13/KLL/kll"
)

func PrintHelp(command string) {
	switch command {
	case "eval":
		fmt.Println(`eval
    - Use this command to interpret and show result of lastline operation
    -a | --all-lines - Use This To Show result of all lines operation`)
	case "interpret":
		fmt.Println(`interpret
	- This Command Is Used For Interpret Code. you can use better errors manipulations`)
	case "version":
		fmt.Println(`version
	- Use this to show current version of kll`)
	case "tokenizer":
		fmt.Println(`tokenizer
	- Use this To show tokens of a file
	- Tokens Is the words
	- Like myword`)
	case "parse":
		fmt.Println(`tokenizer
	- Use this To show nodes of a file
	- Tokens Is the string
	- Like myword`)
	default:
		fmt.Println(`Commands:
	- interpret <script> // Use this command to interpret a code.
	- eval <script>  // Use this command to eval
	- tokenizer <script>  // Use this To show tokens of a file
	- help <command> // Use this command to show more of other commands
	- version // Use this to show current version of kll`)
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
		case "interpret":
			if len(os.Args) > 2 {
				inter := kll.NewInterpreter(kll.NewScope())
				f, err := os.ReadFile(os.Args[2])
				if err != nil {
					fmt.Println(err)
				}
				inter.Panic(inter.Exec(string(f)))
			} else {
				PrintHelp("")
			}
		case "eval":
			if len(os.Args) > 2 {
				all_lines := false
				file := ""
				for i := 2; i < len(os.Args); i++ {
					switch os.Args[i] {
					case "--all-lines", "-a":
						all_lines = true
					default:
						file = os.Args[i]
					}
				}
				inter := kll.NewInterpreter(kll.NewScope())
				f, err := os.ReadFile(file)
				if err != nil {
					fmt.Println(err)
				}
				res, errl := inter.Eval(string(f), all_lines)
				inter.Panic(errl)
				if all_lines {
					for i, v := range res {
						fmt.Printf("Line %v: %s\n", i+1, v)
					}
				} else {
					fmt.Println(res[0])
				}

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
		case "parse":
			if len(os.Args) > 2 {
				f, err := os.ReadFile(os.Args[2])
				if err != nil {
					fmt.Println(err)
				}
				nodes, kllerr := kll.Parse(string(f))
				if !kllerr.Is(kll.Success) {
					fmt.Println(kllerr)
				}
				line := 1
				show := ""
				maxSize := 1
				for _, node := range nodes {
					if len(node.Callstack.Show) > maxSize {
						maxSize = len(node.Callstack.Show)
					}
				}
				for _, node := range nodes {
					show += fmt.Sprintf("Line %v: \"%s%s\" = ", line, node.Callstack.Show, strings.Repeat(" ", maxSize-len(node.Callstack.Show))) + node.String() + "\n"
					line++
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
