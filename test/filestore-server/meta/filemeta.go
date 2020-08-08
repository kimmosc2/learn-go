package meta

// FileMeta:文件元信息结构
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

// UpdateFileMeta:新增/更新文件元信息
func UpdateFileMeta(meta FileMeta) {
	fileMetas[meta.FileSha1] = meta
}

// GetFileMeta:通过文件sha1获取文件元信息,if exists, return file and true
// else,return false
func GetFileMeta(fileSha1 string) (filemeta FileMeta, ok bool) {
	filemeta, ok = fileMetas[fileSha1]
	return filemeta, ok
}

// RemoveFileMeta 移除信息
func RemoveFileMeta(fileSha1 string) {
	delete(fileMetas,fileSha1)
}