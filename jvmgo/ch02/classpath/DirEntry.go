package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		// 发生异常 panic 函数会使程序挂掉，然后运行时会打印出调用栈
		// 但是并不是立即向上传递 panic, 而是等到 defer 那里，等 defer 执行完之后在进行传递
		panic(err)
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	// 拼接完整路径
	fileName := filepath.Join(self.absDir, className)
	// 读取文件二级制
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}
