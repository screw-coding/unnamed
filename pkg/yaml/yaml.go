package yaml

import (
	"io"
	"os"
	"strings"
)

// Unmarshal 从某一个 reader 中读取 yaml 格式的数据，解析后赋值给一个结构体对象
func Unmarshal(in io.Reader, out interface{}) (err error) {
	return nil
}

func UnmarshalFromString(in string, out interface{}) (err error) {
	return Unmarshal(strings.NewReader(in), out)
}

func UnmarshalFromFile(filePath string, out interface{}) (err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	return Unmarshal(file, out)
}
