package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func processFile(inputFile string) {
	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		content, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		result, err := convert(string(content))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}
}

func main() {
	inputFile := ""
	flag.StringVar(&inputFile, "file", "", "input file")
	flag.Parse()

	if inputFile != "" {
		processFile(inputFile)
	} else {
		for i := 1; i < len(flag.Args()); i++ {
			result, err := convert(flag.Args()[i])
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
		}
	}
}

func convert(str string) (string, error) {
	reg := regexp.MustCompile(`[\d\.]+`)
	result := reg.FindAllString(str, -1)
	if len(result) == 0 {
		return "", fmt.Errorf("invalid format")
	}

	if len(result) != 4 {
		return "", fmt.Errorf("invalid rgba format")
	}

	r, _ := strconv.ParseInt(result[0], 10, 64)
	g, _ := strconv.ParseInt(result[1], 10, 64)
	b, _ := strconv.ParseInt(result[2], 10, 64)
	a, _ := strconv.ParseFloat(result[3], 64)

	return fmt.Sprintf("#%02X%02X%02X%02X", r, g, b, int(a*255)), nil
}
