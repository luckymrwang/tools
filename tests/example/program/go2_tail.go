package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	lines := int64(5)
	file, err := os.Open("/Users/sino/Downloads/1_1539779451004.csv")
	if err != nil {
		log.Println(err)
		return
	}
	fileInfo, _ := file.Stat()
	buf := bufio.NewReader(file)
	offset := fileInfo.Size() % 8192
	data := make([]byte, 8192) // 一行的数据
	totalByte := make([][][]byte, 0)
	readLines := int64(0)
	for i := int64(0); i <= fileInfo.Size()/8192; i++ {
		readByte := make([][]byte, 0) // 读取一页的数据
		file.Seek(fileInfo.Size()-offset-8192*i, io.SeekStart)
		data = make([]byte, 8192)
		n, err := buf.Read(data)
		if err == io.EOF {
			if strings.TrimSpace(string(bytes.Trim(data, "\x00"))) != "" {
				readLines++
				readByte = append(readByte, data)
				totalByte = append(totalByte, readByte)
			}
			if readLines > lines {
				break
			}
			continue
		}
		if err != nil {
			log.Println("Read file error:", err)
			return
		}
		strs := strings.Split(string(data[:n]), "\n")
		if len(strs) == 1 {
			b := bytes.Trim([]byte(strs[0]), "\x00")
			if len(b) == 0 {
				continue
			}
		}
		if (readLines + int64(len(strs))) > lines {
			strs = strs[int64(len(strs))-lines+readLines:]
		}
		for j := 0; j < len(strs); j++ {
			readByte = append(readByte, bytes.Trim([]byte(strs[j]+"\n"), "\x00"))
		}
		readByte[len(readByte)-1] = bytes.TrimSuffix(readByte[len(readByte)-1], []byte("\n"))
		totalByte = append(totalByte, readByte)
		readLines += int64(len(strs))

		if readLines >= lines {
			break
		}
	}
	totalByte = ReverseByteArray(totalByte)
	log.Println(ByteArrayToString(totalByte))
}

func ReverseByteArray(s [][][]byte) [][][]byte {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
	return s
}

func ByteArrayToString(buf [][][]byte) string {
	str := make([]string, 0)
	for _, v := range buf {
		for _, vv := range v {
			str = append(str, string(vv))
		}
	}
	return strings.Join(str, "")
}
