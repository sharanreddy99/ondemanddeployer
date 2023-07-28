package utils

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/beego/beego/v2/client/cache"
)

var lock *sync.Mutex
var rdb cache.Cache

func ReadFromFile(fileName string) []byte {
	lock.Lock()
	defer lock.Unlock()
	dat, err := os.ReadFile(fileName)
	if err != nil {
		Log("Error while reading file data: ", err.Error())
		return []byte{}
	}

	return dat
}

func WriteToFile(fileName string, data []byte) {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		Log("Error while opening file for writing: ", err.Error())
	}

	if _, err = f.WriteString(string(data)); err != nil {
		Log("Error while writing data to file: ", err.Error())
	}
}

func ReadFromCache(fileName string) interface{} {
	lock.Lock()
	defer lock.Unlock()
	if data, err := rdb.Get(context.Background(), fileName); err != nil {
		Log("Error while reading file data: ", err.Error())
		return nil
	} else {
		return data
	}
}

func WriteToCache(fileName string, data interface{}) {
	lock.Lock()
	defer lock.Unlock()
	if err := rdb.Put(context.Background(), fileName, data, 24*time.Hour); err != nil {
		Log("Error while opening file for writing: ", err.Error())
	}
}

func init() {
	lock = &sync.Mutex{}
	rdb, _ = cache.NewCache("memory", `{"interval":86400}`)
}
