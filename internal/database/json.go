package database

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var singleton *JsonDatabase

type JsonDatabase struct {
	file *os.File
}

func NewJsonDatabase(filepath string) {
	file, err := os.OpenFile(filepath, os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	// verify if file is empty, write empty json object
	stat, _ := file.Stat()
	if stat.Size() == 0 {
		file.Write([]byte("{}"))
	}

	singleton = &JsonDatabase{file: file}
}

func GetJsonDatabase() *JsonDatabase {
	return singleton
}

func (j *JsonDatabase) Close() error {
	return j.file.Close()
}

func (j *JsonDatabase) Read() ([]byte, error) {
	stat, err := j.file.Stat()
	if err != nil {
		return nil, err
	}

	buff := make([]byte, stat.Size())
	_, err = j.file.Read(buff)
	if err != nil {
		return nil, err
	}

	return buff, nil
}

func (j *JsonDatabase) Write(data []byte) error {
	return ioutil.WriteFile(j.file.Name(), data, 0755)
}

func (j *JsonDatabase) ReadScan(v any) error {
	j.file.Seek(0, 0)
	content, err := j.Read()
	if err != nil {
		return err
	}

	return json.Unmarshal(content, v)
}
