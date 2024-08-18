package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    fmt.Println("Hello, Go!")

    var args []string = os.Args

    if (len(args) < 2) {
        fmt.Println("please provide the filename");
        return;
    }

    var filename string= args[len(args) - 1];
    var isCountCharacter *bool = flag.Bool("c", false, "count number of character");
    var isCountLine *bool = flag.Bool("l", false, "count number of character");
    flag.Parse();

    fmt.Println("Filename: ", filename, "count :", *isCountCharacter);
    var content string = readFile(filename);
    var result string; 
    if (*isCountCharacter) {
        result += " " +strconv.Itoa(countCharacter(content));
    }

    if (*isCountLine) {
        result += " " + strconv.Itoa(countLines(content));
    }

    fmt.Println(result + " " + filename);
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