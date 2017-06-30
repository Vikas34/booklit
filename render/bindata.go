// Code generated by go-bindata.
// sources:
// render/html/paragraph.html
// render/html/section.html
// render/html/sequence.html
// render/html/string.html
// DO NOT EDIT!

package render

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _renderHtmlParagraphHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x29\xb0\xab\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xcc\x4b\x49\xad\xd0\x51\x50\x29\x4e\xcd\x2b\x49\xcd\x4b\x4e\x55\xb0\xb2\x55\xd0\xab\xad\xad\xae\xce\x4c\x83\x4a\xd6\xd6\x2a\x54\x57\xa7\xe6\xa5\x80\x04\x11\xca\x6a\x14\x8a\x52\xf3\x52\x52\x8b\x40\xa2\x60\x49\x1b\xfd\x02\x3b\x2e\x40\x00\x00\x00\xff\xff\xd3\x3e\x01\x8c\x5a\x00\x00\x00")

func renderHtmlParagraphHtmlBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlParagraphHtml,
		"render/html/paragraph.html",
	)
}

func renderHtmlParagraphHtml() (*asset, error) {
	bytes, err := renderHtmlParagraphHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/paragraph.html", size: 90, mode: os.FileMode(420), modTime: time.Unix(1498827174, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlSectionHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\xd6\x0b\xc9\x2c\xc9\x49\x55\xa8\x51\x28\x4a\xcd\x4b\x49\x2d\x52\xa8\x51\x28\x4e\x4d\x2e\xc9\xcc\xcf\xf3\x48\x4d\x04\xf1\xf5\x6a\x6b\xb9\xb8\xaa\xab\xf5\x9c\xf2\x53\x2a\xe1\xaa\x6a\x6b\xb9\x00\x01\x00\x00\xff\xff\x07\xdf\xfd\x32\x3a\x00\x00\x00")

func renderHtmlSectionHtmlBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlSectionHtml,
		"render/html/section.html",
	)
}

func renderHtmlSectionHtml() (*asset, error) {
	bytes, err := renderHtmlSectionHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/section.html", size: 58, mode: os.FileMode(420), modTime: time.Unix(1498825583, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlSequenceHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x2e\x4a\xcc\x4b\x4f\x55\xd0\x73\xce\xcf\x2b\x49\xcd\x2b\x29\xae\xad\xad\xae\xd6\x53\xa8\x51\x28\x4a\xcd\x4b\x49\x2d\x02\xf1\x52\xf3\x52\x6a\x6b\xb9\x00\x01\x00\x00\xff\xff\xfc\xbe\x50\x06\x29\x00\x00\x00")

func renderHtmlSequenceHtmlBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlSequenceHtml,
		"render/html/sequence.html",
	)
}

func renderHtmlSequenceHtml() (*asset, error) {
	bytes, err := renderHtmlSequenceHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/sequence.html", size: 41, mode: os.FileMode(420), modTime: time.Unix(1498830402, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlStringHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x56\xd0\x0b\x2e\x29\xca\xcc\x4b\x57\xd0\xad\xad\xe5\x02\x04\x00\x00\xff\xff\x11\x3f\xd1\x9c\x0f\x00\x00\x00")

func renderHtmlStringHtmlBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlStringHtml,
		"render/html/string.html",
	)
}

func renderHtmlStringHtml() (*asset, error) {
	bytes, err := renderHtmlStringHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/string.html", size: 15, mode: os.FileMode(420), modTime: time.Unix(1498796001, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"render/html/paragraph.html": renderHtmlParagraphHtml,
	"render/html/section.html": renderHtmlSectionHtml,
	"render/html/sequence.html": renderHtmlSequenceHtml,
	"render/html/string.html": renderHtmlStringHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"render": &bintree{nil, map[string]*bintree{
		"html": &bintree{nil, map[string]*bintree{
			"paragraph.html": &bintree{renderHtmlParagraphHtml, map[string]*bintree{}},
			"section.html": &bintree{renderHtmlSectionHtml, map[string]*bintree{}},
			"sequence.html": &bintree{renderHtmlSequenceHtml, map[string]*bintree{}},
			"string.html": &bintree{renderHtmlStringHtml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

