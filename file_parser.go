package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	var files []string

	files = append(files, "c:\\users\\alessio\\data.txt")
	files = append(files, "c:\\users\\alessio\\data1.txt")

	var results []string

	for _, file := range files {

		wg.Add(1)

		go func(file string) {

			result, err := parse(file, &wg)

			if err != nil {
				log.Println(err)
			}

			results = append(results, result)

		}(file)

	}

	wg.Wait()

	for _, result := range results {

		fmt.Println(result)
	}
}

func parse(filename string, wg *sync.WaitGroup) (string, error) {

	defer wg.Done()

	file, err := os.Open(filename)

	if err != nil {

		return fmt.Sprintf("Cannot open file %s \n", filename), err
	}

	data, err := ioutil.ReadAll(file)

	result := fmt.Sprintf(string(data))

	return result, nil

}
