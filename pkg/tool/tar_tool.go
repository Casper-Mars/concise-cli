package tool

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
)

//UnTarGz 解压缩gzip算法压缩的tar压缩包到指定的路径
func UnTarGz(tarPackage string, dstPath string) error {
	fr, err := os.Open(tarPackage)
	if err != nil {
		return err
	}
	defer fr.Close()
	// gzip read
	gr, err := gzip.NewReader(fr)
	if err != nil {
		return err
	}
	defer gr.Close()
	// tar read
	tr := tar.NewReader(gr)
	// 读取文件
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		// 显示文件
		log.Println(h.Name)
		// 判断是文件夹还是文件
		if h.Typeflag == tar.TypeDir {
			os.MkdirAll(dstPath+h.Name, os.FileMode(h.Mode))
		} else {
			// 打开文件
			fw, err := os.OpenFile(dstPath+h.Name, os.O_CREATE|os.O_WRONLY, os.FileMode(h.Mode))
			if err != nil {
				return err
			}
			defer fw.Close()
			// 写文件
			_, err = io.Copy(fw, tr)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
