package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	flag "github.com/spf13/pflag"
)


func main() {

    var isByteCounts *bool = flag.BoolP("bytes", "c", false, "print the byte counts")
    var isCountLines *bool = flag.BoolP("lines", "l", false, "print the newline counts")
    var isCountWords *bool = flag.BoolP("words", "w", false, "print the word counts")
    var isCountLocale *bool = flag.BoolP("chars", "m", false, "print the character counts")
    help := flag.BoolP("help", "h", false, "help message")
    
    flag.Parse()

    if *help {
        flag.Usage()
        os.Exit(0)
    }

    if flag.NArg() != 1 {
        fmt.Println("Error: exactyle one parameter required")
        fmt.Println("Usage: ad_wc [-clwm] parameter")
        os.Exit(1);
    }


    filename := flag.Arg(0)
    inputContent, err := readFile(filename)

    if err != nil {
        fmt.Printf("Error when trying to open file: %v\n", filename)
        os.Exit(1);
    }
   
    var result string; 

    if (*isCountLines) {
        result += " " + strconv.Itoa(countLines(inputContent));
    }

    if (*isCountWords) {
        result += " " + strconv.Itoa(countWords(inputContent));
    }

    if (*isCountLocale) {
        result += " " + strconv.Itoa(countLocale(inputContent));
    }

    if (*isByteCounts) {
        result += " " +strconv.Itoa(countBytes(inputContent));
    }


    fmt.Println(result + " " + filename)
}

func readFile(filename string) (string, error){
    dat, err := os.ReadFile(filename)
    
    if err != nil {
        return "", err
    }

    return string(dat), nil;
}

func countBytes(stringToCount string) int {
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
