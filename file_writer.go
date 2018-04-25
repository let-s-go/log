package log

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

type fileWriter struct {
	fileName string
	fileSize int64
	maxFile  int

	locker sync.Mutex
	size   int64
	file   *os.File
}

func newFileWriter(fileName string, fileSize int64, maxFile int) *fileWriter {
	return &fileWriter{
		fileName: fileName,
		fileSize: fileSize,
		maxFile:  maxFile,
	}
}

func (f *fileWriter) Write(p []byte) (int, error) {
	if f.fileSize > 0 && f.size > f.fileSize {
		oldName := f.fileName + ".log"
		newName := f.fileName + time.Now().Format("_20060102150405.log")
		f.file.Close()
		os.Rename(oldName, newName)
		f.file = nil
		if f.maxFile > 0 {
			go f.remove()
		}
	}
	if f.file == nil {
		oldName := f.fileName + ".log"
		var err error
		f.file, err = os.OpenFile(oldName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
		f.size = 0
		fs, _ := f.file.Stat()
		if fs != nil {
			f.size = fs.Size()
		}
		f.file.Seek(2, 0)
	}
	n, err := f.file.Write(p)
	f.size += int64(n)
	return n, err
}

func (f *fileWriter) remove() {
	f.locker.Lock()
	defer f.locker.Unlock()

	dir, name := path.Split(f.fileName)

	infos, _ := ioutil.ReadDir(dir)
	ns := make([]string, 0)
	for _, info := range infos {
		if !info.IsDir() {
			n := info.Name()
			if n != name+".log" && strings.HasPrefix(n, name) {
				ns = append(ns, n)
			}
		}
	}

	sort.Strings(ns)
	for i := 0; i < len(ns)-f.maxFile; i++ {
		os.Remove(filepath.Join(dir, ns[i]))
	}
}
