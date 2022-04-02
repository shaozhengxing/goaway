package fileCache

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type fileCache struct {
}

type structure struct {
	Data   string `json:"data"`
	Expire int    `json:"expire"`
}

func NewFileCache() *fileCache {
	return &fileCache{}
}

func (f fileCache) Put(key string, value string, ttl time.Duration) error {
	file, err := os.OpenFile("storage/"+fmt.Sprintf("%x", md5.Sum([]byte(key))), os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	var v structure
	v.Data = value
	if ttl == 0 {
		v.Expire = 0
	} else {
		v.Expire = int(time.Now().Add(ttl).Unix())
	}

	jsonStr, err := json.Marshal(v)
	if err != nil {
		return err
	}

	_, err = file.Write(jsonStr)

	return err
}

func (f fileCache) Get(key string) (string, error) {
	filePath := "storage/" + fmt.Sprintf("%x", md5.Sum([]byte(key)))

	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	var v structure
	err = json.Unmarshal(data, &v)
	if err != nil {
		return "", err
	}

	now := time.Now().Unix()
	if v.Expire != 0 && (now > int64(v.Expire)) {
		defer func(f *os.File, ff string) {
			f.Close()

			os.Remove(ff)
		}(file, filePath)

		return "", errors.New("缓存已超时")
	}

	return v.Data, err
}
