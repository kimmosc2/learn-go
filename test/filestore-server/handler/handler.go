package handler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"learn-go/test/filestore-server/meta"
	"learn-go/test/filestore-server/util"
	"log"
	"net/http"
	"os"
	"time"
)

// UploadHandler 上传文件接口
func UploadHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			log.Printf("read file error:%v", err)
			io.WriteString(writer, "internal server error")
		}
		io.WriteString(writer, string(data))
		// 返回上传HTML
	} else if request.Method == http.MethodPost {
		// 接收文件流及存储到本地目录
		file, header, err := request.FormFile("file")
		if err != nil {
			log.Printf("failed to get data:%v", err)
			return
		}
		defer file.Close()
		fileMeta := meta.FileMeta{
			FileName: header.Filename,
			Location: "./tmp/" + header.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		create, err := os.Create(fileMeta.Location)
		if err != nil {
			log.Printf("create file error:%v", err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer create.Close()
		fileMeta.FileSize, err = io.Copy(create, file)
		if err != nil {
			log.Printf("copy file error:%v", err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		create.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(create)
		meta.UpdateFileMeta(fileMeta)

		http.Redirect(writer, request, "/file/upload/suc", http.StatusFound)
	}
}

// UploadSucHandler 上传完成
func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload finished")
}

// GetFileInfoHandler 查询文件接口(获取文件元信息)
func GetFileInfoHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filehash := r.Form["filehash"][0]
	fileMeta, exist := meta.GetFileMeta(filehash)
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	marshal, err := json.Marshal(fileMeta)
	if err != nil {
		log.Printf("json marshal error:%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(marshal)
}

// DownloadHandler 文件下载接口
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filehash := r.Form.Get("filehash")
	fileMeta, exist := meta.GetFileMeta(filehash)
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	open, err := os.Open(fileMeta.Location)
	if err != nil {
		log.Printf("hit the cache but no such file:%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer open.Close()

	all, err := ioutil.ReadAll(open)
	// _, err = io.Copy(w, open)
	if err != nil {
		log.Printf("read file error:%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("Content-Descrption", "attachment;filename=\""+fileMeta.FileName+"\"")
	_, err = w.Write(all)
	if err != nil {
		log.Printf("write fail:%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

// // FileQueryHandler : 查询批量的文件元信息
// func FileQueryHandler(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
//
// 	limitCnt, _ := strconv.Atoi(r.Form.Get("limit"))
// 	username := r.Form.Get("username")
// 	//fileMetas, _ := meta.GetLastFileMetasDB(limitCnt)
// 	// userFiles, err := dblayer.QueryUserFileMetas(username, limitCnt)
// 	// if err != nil {
// 	// 	w.WriteHeader(http.StatusInternalServerError)
// 	// 	return
// 	// }
//
// 	data, err := json.Marshal(userFiles)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	w.Write(data)
// }

// FileMetaUpdateHandle 修改文件接口
func FileMetaUpdateHandle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	opType := r.Form.Get("op")
	fileSha1 := r.Form.Get("filehash")
	newFileName := r.Form.Get("filename")
	if opType != "0" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	curFileMeta, exist := meta.GetFileMeta(fileSha1)
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("not such file from hash %v", fileSha1)
		return
	}
	curFileMeta.FileName = newFileName
	meta.UpdateFileMeta(curFileMeta)
	marshal, err := json.Marshal(curFileMeta)
	if err != nil {
		log.Printf("json marshal error:%v",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(marshal)
}


// FileDeleteHandler 删除文件
func FileDeleteHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fileSha1 := r.Form.Get("filehash")

	fileMeta, ok := meta.GetFileMeta(fileSha1)
	if !ok {
		log.Printf("no such file from hash %v", fileSha1)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err := os.Remove(fileMeta.Location)
	if err != nil {
		log.Printf("remove error:%v",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	meta.RemoveFileMeta(fileSha1)
	w.WriteHeader(http.StatusOK)
}