package assetgo

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// WalkDirectory walk directory to generate compress data
func WalkDirectory(w *Writer, dir string) error {
	return filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		data, err := ioutil.ReadFile(p)
		if err != nil {
			return err
		}
		compressedData, err := Compress(data)
		if err != nil {
			return err
		}
		return w.WriteAssetFile(p, info, compressedData)
	})
}
