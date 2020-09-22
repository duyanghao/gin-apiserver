package util

import (
	"github.com/pkg/sftp"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ScpFile(client *sftp.Client, srcPath, dstPath string, forward bool) error {
	if !forward {
		dstFile, err := client.OpenFile(dstPath, os.O_RDWR)
		if err != nil {
			return err
		}
		defer dstFile.Close()
		// Delete file firstly
		os.Remove(srcPath)
		srcFile, err := os.Create(srcPath)
		if err != nil {
			return err
		}
		defer srcFile.Close()
		_, err = dstFile.WriteTo(srcFile)
		if err != nil {
			return err
		}
	} else {
		// Delete file firstly
		client.Remove(dstPath)
		dstFile, err := client.OpenFile(dstPath, os.O_RDWR|os.O_CREATE)
		if err != nil {
			return err
		}
		defer dstFile.Close()
		srcFile, err := os.Open(srcPath)
		if err != nil {
			return err
		}
		defer srcFile.Close()
		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func ScpDir(client *sftp.Client, srcDir, dstDir string, forward bool) error {
	var files []os.FileInfo
	var err error
	if !forward {
		files, err = client.ReadDir(dstDir)
	} else {
		files, err = ioutil.ReadDir(srcDir)
	}
	if err != nil {
		return err
	}
	for _, f := range files {
		err = ScpFile(client, filepath.Join(srcDir, f.Name()), filepath.Join(dstDir, f.Name()), forward)
		if err != nil {
			return err
		}
	}
	return nil
}
