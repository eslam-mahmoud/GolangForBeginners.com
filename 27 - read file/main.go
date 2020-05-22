package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	{
		// Read full file in the memory
		// do not do this in production app
		bytes, err := ioutil.ReadFile("./test.json")
		if err != nil {
			panic(err)
		}
		fmt.Print(string(bytes))
	}
	{
		// read file in chunks
		f, err := os.Open("./test.json")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		reader := bufio.NewReader(f)
		b := make([]byte, 5)
		for {
			numberOfBytesRead, err := reader.Read(b)
			if err != nil {
				if !errors.Is(err, io.EOF) {
					fmt.Println("Error reading file:", err)
				}
				break
			}
			fmt.Println(string(b[0:numberOfBytesRead]))
			// fmt.Println(string(b))
		}
	}
	{
		// read file line by line
		f, err := os.Open("./test.json")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		s := bufio.NewScanner(f)
		for s.Scan() {
			fmt.Println(s.Text())
		}
		// s.Err() does not return io.EOF
		err = s.Err()
		if err != nil {
			panic(err)
		}
	}
}
