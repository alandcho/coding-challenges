package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestCalculateStats(t *testing.T) {
    // Test case 1: Empty input
    reader := bufio.NewReader(strings.NewReader(""))
    stats := calculateStats(reader, "test.txt")
    if stats.bytes != 0 || stats.lines != 0 || stats.words != 0 || stats.runes != 0 || stats.filename != "test.txt" {
        t.Errorf("Test case 1 failed")
    }

    // Test case 2: Single line with no spaces
    reader = bufio.NewReader(strings.NewReader("HelloWorld"))
    stats = calculateStats(reader, "test.txt")
    if stats.bytes != 10 || stats.lines != 0 || stats.words != 1 || stats.runes != 10 || stats.filename != "test.txt" {
        t.Errorf("Test case 2 failed")
    }

    // Test case 3: Multiple lines with spaces
    reader = bufio.NewReader(strings.NewReader("Hello\nWorld\n"))
    stats = calculateStats(reader, "test.txt")
    if stats.bytes != 12 || stats.lines != 2 || stats.words != 2 || stats.runes != 12 || stats.filename != "test.txt" {
        t.Errorf("Test case 3 failed")
    }

    // Test case 4: Unicode characters
    reader = bufio.NewReader(strings.NewReader("こんにちは"))
    stats = calculateStats(reader, "test.txt")
    if stats.bytes != 15 || stats.lines != 0 || stats.words != 1 || stats.runes != 5 || stats.filename != "test.txt" {
        t.Errorf("Test case 4 failed")
    }
}