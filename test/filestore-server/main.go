package main

import (
	"learn-go/test/filestore-server/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/meta",handler.GetFileInfoHandler)

	http.HandleFunc("/file/download",handler.DownloadHandler)
	http.HandleFunc("/file/update",handler.FileMetaUpdateHandle)
	http.HandleFunc("/file/delete",handler.FileDeleteHandler)
	http.ListenAndServe(":80", nil)
}
