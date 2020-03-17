package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

var (
	inFilePath  = flag.String("in", "", "input file path")
	outFilePath = flag.String("out", "", "output file path")
)

func main() {

	flag.Parse()
	input_dir := *inFilePath
	output_dir := *outFilePath

	fmt.Println(input_dir + "\n")
	fmt.Println(output_dir + "\n")

	files, _ := ioutil.ReadDir(input_dir)
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			input_file := input_dir + file.Name()
			output_file := output_dir + file.Name()
			fmt.Println(output_file)

			in, err := os.Open(input_file)
			if err != nil {
				log.Fatal(err)
			}
			defer in.Close()

			out, err := os.Create(output_file)
			if err != nil {
				log.Fatal(err)
			}
			defer out.Close()

			reader := transform.NewReader(in, traditionalchinese.Big5.NewDecoder())
			scanner := bufio.NewScanner(reader)
			for scanner.Scan() {
				s := scanner.Text()
				if s[len(s)-1] == '\\' {
					out.WriteString(s[:len(s)-1])
				} else {
					out.WriteString(s + "\n")
				}
			}

			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
