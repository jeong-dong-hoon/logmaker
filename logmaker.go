package logmaker

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
)

type logmessage struct {
	Message string
}

func MakeLog(title string, message logmessage, filepath string) error {

	path := path.Join(filepath, title)
	mkdirifnotexist(filepath)

	err := ioutil.WriteFile(path, []byte(message.Message), 0644)
	if err != nil {
		return err
	}
	return nil
}

func MakeMessage(Message ...string) logmessage {
	result := logmessage{}
	var str bytes.Buffer
	for i := 0; i < len(Message); i++ {
		str.WriteString(Message[i])
		str.WriteString("\n")
	}
	result.Message = str.String()
	return result
}

func mkdirifnotexist(path string) {
	_, err := os.ReadDir(path)
	if !os.IsExist(err) {

		os.Mkdir(path, 0644)

	}

}
