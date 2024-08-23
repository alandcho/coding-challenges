package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

type Stats struct {
    bytes int64
    lines int64
    words int64
    runes int64
    filename string
}

func main() {
    var isCountCharacters *bool = flag.Bool("c", false, "count number of characters")
    var isCountLines *bool = flag.Bool("l", false, "count number of lines")
    var isCountWords *bool = flag.Bool("w", false, "count number of words")
    var isCountLocale *bool = flag.Bool("m", false, "count number of words")
    flag.Parse()

    filenames := flag.CommandLine.Args()

    if len(filenames) == 0 { 
        reader := bufio.NewReader(os.Stdin)
        stats := calculateStats(reader, "")
        printStats(*isCountCharacters, *isCountLines, *isCountWords, *isCountLocale, stats)
    } else {
        for _, filename := range filenames {
            stats, err := calculateStatsForFile(filename)
            if err != nil {
                fmt.Println("Error reading file: ", err)
            } else {
                printStats(*isCountCharacters, *isCountLines, *isCountWords, *isCountLocale, stats)
            }
        }
    }
}

func printStats(isCountCharacters bool, isCountLines bool, isCountWords bool, isCountLocale bool, stats Stats) {

    if !isCountCharacters && !isCountLines && !isCountWords && !isCountLocale {
        isCountCharacters = true
        isCountLines = true
        isCountWords = true
        isCountLocale = true
    }

    if isCountCharacters {
        fmt.Printf("%d ", stats.bytes)
    }

    if isCountLines {
        fmt.Printf("%d ", stats.lines)
    }

    if isCountWords {
        fmt.Printf("%d ", stats.words)
    }

    if isCountLocale {
        fmt.Printf("%d ", stats.runes)
    }

    fmt.Printf(stats.filename)
}


func calculateStatsForFile(filename string) (Stats, error) {
    file, err := os.Open(filename)
    if (err != nil) {
        return Stats{}, err
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    return calculateStats(reader, filename), nil
}

func calculateStats(reader *bufio.Reader, filename string) Stats {
    var prevRune rune
    var byteCount int64
    var lineCount int64
    var wordCount int64
    var runeCount int64


    for {
        currentRune, bytes, err := reader.ReadRune()

        if err != nil {
			if err == io.EOF {
				if prevRune != rune(0) && !unicode.IsSpace(prevRune) {
					wordCount++
				}
				break
			}
			log.Fatal(err)
		}

        byteCount += int64(bytes)
        runeCount++

        if currentRune == '\n' {
            lineCount++
        }

        if unicode.IsSpace(currentRune) && !unicode.IsSpace(prevRune) {
            wordCount++
        }

        prevRune = currentRune

    }

    return Stats{bytes: byteCount, lines: lineCount, words: wordCount, runes: runeCount, filename: filename}
}

