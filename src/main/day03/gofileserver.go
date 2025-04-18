package main

import (
	"fmt"
    "io"
    "mime"
    "net/http"
    "os"
    "path/filepath"
)

func main() {
    fs := http.FileServer(http.Dir("/Users/chao/goworkspace/golearning/src/main/day03"))
    http.Handle("/", fs)

    http.HandleFunc("/upload", uploadHandler)
    http.HandleFunc("/download", downloadHandler)

    fmt.Println("Listening on :8080...")
    http.ListenAndServe(":8080", nil)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        // 文件上传
		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()
		f, err := os.Create(header.Filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		if _, err := io.Copy(f, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
    } else {
        http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
    }
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        // 文件下载
		file := r.URL.Path[1:]
		if _, err := os.Stat(file); err == nil {
			f, err := os.Open(file)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer f.Close()
			w.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(file)))
			w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%!q(MISSING)", filepath.Base(file)))
			if _, err := io.Copy(w, f); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
    } else {
        http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
    }
}
