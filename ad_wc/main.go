package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    var args []string = os.Args

    var inputContent string = "";
    var filename string = "";
    if (len(args) != 3) {
        content, err := io.ReadAll(os.Stdin)

        inputContent = string(content)

        if err != nil {
            fmt.Fprintln(os.Stderr, "reading standard input:", err)
        }
    } else {
        filename = args[len(args) - 1];
        inputContent =  readFile(filename);
    }

 
    var isCountCharacters *bool = flag.Bool("c", false, "count number of characters");
    var isCountLines *bool = flag.Bool("l", false, "count number of lines");
    var isCountWords *bool = flag.Bool("w", false, "count number of words");
    var isCountLocale *bool = flag.Bool("m", false, "count number of words");
    flag.Parse();

   
    var result string; 
    if (*isCountCharacters) {
        result += " " +strconv.Itoa(countCharacter(inputContent));
    }

    if (*isCountLines) {
        result += " " + strconv.Itoa(countLines(inputContent));
    }

    if (*isCountWords) {
        result += " " + strconv.Itoa(countWords(inputContent));
    }

    if (*isCountLocale) {
        result += " " + strconv.Itoa(countLocale(inputContent));
    }

    fmt.Println(result + " " + filename)
}

func readFile(filename string) string{
    dat, err := os.ReadFile(filename)
    check(err)
    return string(dat);
}

func countCharacter(stringToCount string) int {
    return len(stringToCount);
}

func countLines(stringToCount string) int {
    var count int = 0;
    
    for _, char := range stringToCount {
        if (char == '\n') {
            count++;
        }
    } 

    return count;
}

func countWords(stringToCount string) int {
    return len(strings.Fields(stringToCount));
}

func countLocale(stringToCount string) int {
    return utf8.RuneCountInString(stringToCount);
}
