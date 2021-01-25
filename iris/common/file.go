package common

import (
	"bufio"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/BurntSushi/graphics-go/graphics"
)

func AbsolutePath(path string) string {
	var pathPrefix string
	file, err := exec.LookPath(os.Args[0])
	if err == nil {
		path, err := filepath.Abs(file)
		if err == nil {
			index := strings.LastIndex(path, string(os.PathSeparator))
			pathPrefix = path[:index] + "/"
		}
	}
	return pathPrefix + path
}

func ReadFile(path string) string {
	return string(ReadFileBytes(path))
}

func ReadFileBytes(path string) []byte {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		Error("ReadFile:", err.Error())
	}
	return c
}

func WriteFile(path string, body []byte) error {
	err := ioutil.WriteFile(path, body, 0644)
	if err != nil {
		Error("WriteFile:", err.Error())
		return err
	}
	return nil
}

func AppendFile(file string, text string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		Error("AppendFile:", err.Error())
	}
	if _, err = f.WriteString(text); err != nil {
		Error("AppendFile:", err.Error())
	}
}

func DeleteFile(path string) {
	err := os.Remove(path)
	if err != nil {
		Error("DeleteFile:", err.Error())
	}
}

func ReadLine(fileName string, limit int, offset int) (r string, e error) {
	f, err := os.Open(fileName)
	if err != nil {
		e = err
		return
	}
	buf := bufio.NewReader(f)
	for i := 0; i < offset+limit; i++ {
		line, err := buf.ReadString('\n')
		if i >= offset {
			r = r + line
		}
		if err != nil {
			if err == io.EOF {
				return
			}
			return
		}
	}
	return
}

func ReplaceLine(fileName string, line int, with string) (string, error) {
	if line < 1 {
		return "", errors.New(JoinStr("无效的行号", line))
	}
	if runtime.GOOS == "windows" {
		return "", nil
	}
	return Exec(fmt.Sprintf(`sed -i '' '%vs/.*/%v/' %v`, line, with, fileName))
}

func Pathinfo(url string) P {
	p := P{}
	url = strings.Replace(url, "\\", "/", -1)
	if strings.Index(url, "/") < 0 {
		url = JoinStr("./", url)
	}
	re := regexp.MustCompile("(.*)/([^/]*)\\.([^.]*)")
	match := re.FindAllStringSubmatch(url, -1)
	if len(match) > 0 {
		m0 := match[0]
		if len(m0) == 4 {
			p["basename"] = m0[0]
			p["dirname"] = m0[1]
			p["filename"] = m0[2]
			p["extension"] = strings.ToLower(m0[3])
		}
	}
	return p
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func Mkdir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func DirTree(path string, ext string, limit int) (files []P) {
	files = []P{}
	i := 0
	filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if i >= limit {
			return errors.New("reach limit")
		}
		i++
		if f != nil && !f.IsDir() {
			if strings.HasSuffix(f.Name(), ext) {
				files = append(files, P{"file": path})
			}
		}
		return nil
	})
	return
}

func FileRemoveLine(file string, start int, lines int) {
	cmd := fmt.Sprintf("sed -i '%v,%vd' %v", start, start+lines-1, file)
	if runtime.GOOS == "windows" {
		return
	}
	Exec(cmd)
}

func RemoveSpaceLine(file string, filter interface{}) {
	cmd := fmt.Sprintf("sed -i '/%v/d' %v", filter, file)
	if runtime.GOOS == "windows" {
		return
	}
	Exec(cmd)
}

func FileInsertLine(file string, start int, txt string) {
	cmd := fmt.Sprintf("sed -i '%vi %v' %v", start, txt, file)
	if runtime.GOOS == "windows" {
		return
	}
	Exec(cmd)
}

func ResizeImage(file string, width int) error {
	src, err := LoadImage(file)
	if err != nil {
		return err
	}
	bound := src.Bounds()
	dx := bound.Dx()
	dy := bound.Dy()
	// 缩略图的大小
	dst := image.NewRGBA(image.Rect(0, 0, width, width*dy/dx))
	// 产生缩略图,等比例缩放
	err = graphics.Scale(dst, src)
	if err != nil {
		return err
	}
	//保存文件
	err = SaveImage(file, dst)
	if err != nil {
		return err
	}
	return nil
}

func LoadImage(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, err
}

func SaveImage(path string, img image.Image) (err error) {
	imgfile, err := os.Create(path)
	defer imgfile.Close()
	err = png.Encode(imgfile, img)
	if err != nil {
		log.Fatal(err)
	}
	return
}

//返回拷贝的数据总量，以及拷贝过程中产生的错误
func CopyFile(destName string, srcName string) (int64, error) {
	srcFile, err := os.Open(srcName)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()
	destFile, err := os.OpenFile(destName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return 0, err
	}
	defer destFile.Close()
	return io.Copy(destFile, srcFile)
}
