package assetgo

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Writer write go source file
type Writer struct {
	bytes.Buffer
}

// WritePackage write package name
func (w *Writer) WritePackage(name string) error {
	_, err := w.Buffer.WriteString("package " + name + "\n\n")
	return err
}

// WriteImport write correct imports
func (w *Writer) WriteImport() error {
	_, err := w.Buffer.WriteString(`
import(
    "time"
)
`)
	return err
}

// WriteInitBegin write global variabales and init function block beginning
func (w *Writer) WriteInitBegin() error {
	_, err := w.Buffer.WriteString(`
var assetData = make(map[string]*Asset)

type Asset struct {
	File     string
	Data     string
	Size     int64
	ModTime  time.Time
	rawBytes []byte
}

func init(){
`)
	return err
}

// WriteInitEnd write init function block ending
func (w *Writer) WriteInitEnd() error {
	_, err := w.Buffer.WriteString(`
}
`)
	return err
}

var assetFileTemplate = `
    assetData["%s"] = &Asset{
        File:"%s",
        Data:"%s",
        Size:%d,
        ModTime:time.Unix(%d,0),
    }
`

// WriteAssetFile write asset file to struct to global variable
func (w *Writer) WriteAssetFile(p string, info os.FileInfo, data []byte) error {
	p = filepath.ToSlash(p)
	str := fmt.Sprintf(assetFileTemplate, p, p, data, info.Size(), info.ModTime().Unix())
	_, err := w.Buffer.WriteString(str)
	return err
}

// ToFile flush writer bytes to file
func (w *Writer) ToFile(file string) error {
	os.MkdirAll(filepath.Dir(file), os.ModePerm)
	return ioutil.WriteFile(file, w.Buffer.Bytes(), os.ModePerm)
}
