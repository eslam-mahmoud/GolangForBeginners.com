package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// type with exported field & annotations
// annotations used mostly in JSON and DB
// check https://www.youtube.com/watch?v=XbDtxIbJpNE
type human struct {
	Name      string   `json:"name"`
	Expertise []string `json:"expertise"`
}

func main() {
	// values that we want to save to file
	h := human{
		Name:      "Eslam",
		Expertise: []string{"GoLang", "PHP", "JS"},
	}
	// convert struct to []bytes
	hBytes, err := json.Marshal(h)
	panicOnError(err)
	_ = hBytes
	{
		// err = ioutil.WriteFile("./example1.json", hBytes, os.ModePerm)
		// panicOnError(err)
		// log.Println("example1.json saved")
	}
	{
		// 	f, err := os.Create("./example2.txt")
		// 	panicOnError(err)
		// 	defer f.Close()

		// 	// we can "write" byte slices
		// 	writtenBytes, err := f.Write([]byte("First line"))
		// 	panicOnError(err)
		// 	fmt.Printf("wrote %v \n", writtenBytes)

		// 	// or we can "WriteString"
		// 	writtenBytes, err = f.WriteString("\nSecond line")
		// 	panicOnError(err)
		// 	fmt.Printf("wrote %v \n", writtenBytes)

		// 	log.Println("example2.txt saved")
	}
	{
		// f, err := os.Create("./example3.json")
		// panicOnError(err)
		// defer f.Close()

		// w := bufio.NewWriter(f)
		// // WriteString slpit the string if it is bigger than
		// // max buffer size
		// writtenBytes, err := w.WriteString(string(hBytes))
		// // we can use w.Write() and send []byte
		// panicOnError(err)
		// fmt.Printf("wrote %d\n", writtenBytes)
		// w.Flush()
		// log.Println("example3.json saved")
		// // Flush writes any buffered data to the underlying io.Writer.
	}
	{
		// f, err := os.Create("./example4.json")
		// panicOnError(err)
		// defer f.Close()

		// linesToWrite := []string{
		// 	"This is an example",
		// 	"how to write to a file",
		// 	"line by line.",
		// }

		// w := bufio.NewWriter(f)
		// for _, line := range linesToWrite {
		// 	writtenBytes, err := w.WriteString(line)
		// 	panicOnError(err)
		// 	fmt.Printf("Bytes Written: %d\n", writtenBytes)
		// 	fmt.Printf("Available: %d\n", w.Available())
		// 	fmt.Printf("Buffered : %d\n", w.Buffered())
		// }
		// w.Flush()
		// log.Println("example4.json saved")
		// // Flush writes any buffered data to the underlying io.Writer.
	}
	{
		// f, err := os.Create("./example5.txt")
		// panicOnError(err)
		// defer f.Close()

		// linesToWrite := []string{
		// 	"This is an example\n",
		// 	"how to write to a file\n",
		// 	"line by line.\n",
		// }

		// w := bufio.NewWriterSize(f, 10)
		// for _, line := range linesToWrite {
		// 	writtenBytes, err := w.WriteString(line)
		// 	panicOnError(err)
		// 	fmt.Printf("Bytes Written: %d\n", writtenBytes)
		// 	fmt.Printf("Available: %d\n", w.Available())
		// 	fmt.Printf("Buffered : %d\n", w.Buffered())
		// }
		// // Flush writes any buffered data to the underlying io.Writer.
		// w.Flush()
		// log.Println("example5.txt saved")
	}
	{
		// create file
		err := ioutil.WriteFile("./example6.txt", []byte("Creating example6"), os.ModePerm)
		panicOnError(err)
		log.Println("example6.txt saved")

		// open file to append data to it
		f, err := os.OpenFile("./example6.txt", os.O_APPEND|os.O_WRONLY, 0777)
		panicOnError(err)
		defer f.Close()

		_, err = f.WriteString("\nUpdating example6")
		panicOnError(err)
		log.Println("example6.txt updated")
	}
}

func panicOnError(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
