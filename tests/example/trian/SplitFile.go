package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
)

func main() {

	fileToBeChunked := "/Users/sino/Downloads/wrk/IMG_2243.JPG"

	file, err := os.Open(fileToBeChunked)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()

	var fileSize int64 = fileInfo.Size()
	fmt.Println("fileSize:", fileSize)

	const fileChunk = 3 * (1 << 20) // 1 MB, change this to your requirement

	// calculate total number of parts the file will be chunked into

	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

	fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

	for i := uint64(0); i < totalPartsNum; i++ {

		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		fmt.Println("partSize:", partSize)
		partBuffer := make([]byte, partSize)

		file.Read(partBuffer)

		// write to disk
		fileName := "somebigfile_" + strconv.FormatUint(i, 10)
		_, err := os.Create(fileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// write/save buffer to disk
		err = ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Split to : ", fileName)
	}
}
