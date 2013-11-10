package takarabago

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"path"
	"runtime"
	"time"
)

type Tails []string
type Heads []string

func Takarabako() string {
	return (GetHead() + GetTail())
}

func GetHead() string {
	heads := ReadJson(path.Join(GetCurrentDir(), "../dict/head.json"))
	return heads[RandInt(len(heads))]
}

func GetTail() string {
	tails := ReadJson(path.Join(GetCurrentDir(), "../dict/tail.json"))
	return tails[RandInt(len(tails))]
}

func ReadJson(filePath string) []string {
	readJson, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	items := []string{}
	err = json.Unmarshal([]byte(readJson), &items)
	if err != nil {
		panic(err)
	}
	return items
}

func RandInt(length int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(length)
}

func GetCurrentDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
