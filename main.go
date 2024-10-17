package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"io"
	"bytes"
)

func getNumberOfCharacters(r io.Reader) (int) {
	scanner := bufio.NewScanner(r)
    scanner.Split(bufio.ScanRunes) // Split by words

    // Count the words
    CharCount := 0
    for scanner.Scan() {
        CharCount++
    }

    // Check for errors during scanning
    if err := scanner.Err(); err != nil {
        fmt.Println("Error scanning file:", err)
        return 0
    }
	return CharCount
}

func getNumberOfWords(r io.Reader) (int) {
	scanner := bufio.NewScanner(r)
    scanner.Split(bufio.ScanWords) // Split by words

    // Count the words
    wordCount := 0
    for scanner.Scan() {
        wordCount++
    }

    // Check for errors during scanning
    if err := scanner.Err(); err != nil {
        fmt.Println("Error scanning file:", err)
        return 0
    }
	return wordCount
}

func getNumberOfLines(r io.Reader) (int) {
	reader := bufio.NewReader(r)
	lineCount := 0
	for {
		_, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		lineCount++
	}
	return lineCount
}

func getFileSize(r io.Reader) (int) {
	n, err := io.Copy(io.Discard, r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
		os.Exit(1)
	}
	
	return int(n)
}

func main() {
	numberOfBytesFlag := flag.Bool("c", false, "outputs the number of bytes in a file")
	numberOfLinesFlag := flag.Bool("l", false, "outputs the number of lines in a file")
	numberOfWordsFlag := flag.Bool("w", false, "outputs the number of Words in a file")
	numberOfCharFlag := flag.Bool("m", false, "outputs the number of Characters in a file")
	flag.Parse()
	args := flag.Args()

	if len(args) != 0 {
		filePath := args[0]
		_, err := os.Stat(filePath)
		if os.IsNotExist(err) {
			fmt.Printf("file does not exist: %s", filePath)
			return
		} else if err != nil {
			fmt.Printf("error accessing the file: %v", err)
			return
		}

		file, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		os.Stdin = file

	} 
	
	var buf bytes.Buffer
	_, err := io.Copy(&buf, os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
		os.Exit(1)
	}

	reader1 := bytes.NewReader(buf.Bytes())

	if (*numberOfBytesFlag) {
		fileSize := getFileSize(reader1)
		if len(args) != 0 {
			fmt.Printf("%d\t%s\n", fileSize, args[0])
			return
		}
		fmt.Printf("%d\n", fileSize)
		return
	}

	if (*numberOfLinesFlag ) {
		lineCount := getNumberOfLines(reader1)
		if len(args) != 0 {
			fmt.Printf("%d\t%s\n", lineCount, args[0])
			return
		}
		fmt.Printf("%d\n", lineCount)
		return
	}

	if (*numberOfWordsFlag ) {
		wordCount := getNumberOfWords(reader1)
		if len(args) != 0 {
			fmt.Printf("%d\t%s\n", wordCount, args[0])
			return
		}
		fmt.Printf("%d\n", wordCount)
		return
	}

	if (*numberOfCharFlag ) {
		CharCount := getNumberOfCharacters(reader1)
		if len(args) != 0 {
			fmt.Printf("%d\t%s\n", CharCount, args[0])
			return
		}
		fmt.Printf("%d\n", CharCount)
		return
	}

	fileSize := getFileSize(reader1)

	reader2 := bytes.NewReader(buf.Bytes())
	lineCount := getNumberOfLines(reader2)

	reader3 := bytes.NewReader(buf.Bytes())
	wordCount := getNumberOfWords(reader3)
	if len(args) != 0 {
		fmt.Printf("%d\t%d\t%d\t%s\n", lineCount, wordCount, fileSize, args[0])
		return
	}
	fmt.Printf("%d\t%d\t%d\n", lineCount, wordCount, fileSize)
}
