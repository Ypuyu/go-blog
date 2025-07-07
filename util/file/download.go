package file

import (
	"context"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"go-blog/logger"
)

// Download 文件下载
func Download(fromURL, path string) error {
	ctx := context.Background()
	schema := "http://"
	httpsSchema := "https://"
	if !strings.HasPrefix(fromURL, schema) && !strings.HasPrefix(fromURL, httpsSchema) {
		fromURL = schema + fromURL
	}
	resp, err := http.Get(fromURL)
	if err != nil {
		logger.Errorf(ctx, "Download get url err:%v", err)
		return err
	}
	defer resp.Body.Close()

	dir := filepath.Dir(path)
	err = os.MkdirAll(dir, fs.ModeDir)
	if err != nil {
		logger.Errorf(ctx, "MkdirAll err:%v", err)
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		logger.Errorf(ctx, "Download open file err:%v", err)
		return err
	}
	defer file.Close()
	if _, err := io.Copy(file, resp.Body); err != nil {
		logger.Errorf(ctx, "Download file read err:%v", err)
		return err
	}
	if err = file.Sync(); err != nil {
		logger.Errorf(ctx, "file Sync err:%v", err)
		return err
	}
	return nil
}
