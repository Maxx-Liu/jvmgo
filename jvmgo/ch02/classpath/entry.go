package classpath

import (
	"os"
	"strings"
)

// 路径分隔符";"
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	/**
	负责寻找和加载类
	@param className 相对路径，路径之间用 / 分割，后跟 .class
	@return (字节数据, 最终定位到 class 文件的 Entry, 错误信息)
	 */
	readClass(className string) ([]byte, Entry, error)
	// toString()
	String() string
}
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
