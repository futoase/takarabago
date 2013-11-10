package takarabago_test

import (
	"encoding/json"
	tg "github.com/futoase/takarabago/libs"
	"io/ioutil"
	"path"
	"runtime"
	"testing"
)

func TestHead(t *testing.T) {
	heads, err := ioutil.ReadFile(path.Join(GetCurrentDir(), "../dict/head.json"))

	if err != nil {
		t.Error(err)
	}

	items := []string{}
	err = json.Unmarshal([]byte(heads), &items)
	if err != nil {
		t.Error(err)
	}
	if stringInSlice(tg.GetHead(), items) == false {
		t.Error("Not head in heads.json!")
	}
}

func TestTail(t *testing.T) {
	heads, err := ioutil.ReadFile(path.Join(GetCurrentDir(), "../dict/tail.json"))

	if err != nil {
		t.Error(err)
	}

	items := []string{}
	err = json.Unmarshal([]byte(heads), &items)
	if err != nil {
		t.Error(err)
	}
	if stringInSlice(tg.GetTail(), items) == false {
		t.Error("Not tail in tails.json!")
	}
}

func stringInSlice(val string, list []string) bool {
	for _, v := range list {
		if val == v {
			return true
		}
	}
	return false
}

func GetCurrentDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
