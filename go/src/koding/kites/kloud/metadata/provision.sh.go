// Code generated by go-bindata.
// sources:
// provision.sh
// DO NOT EDIT!

package metadata

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

var _provisionSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x52\x6d\x6b\x2a\x47\x14\xfe\xbc\xf3\x2b\x4e\xd7\xd0\x26\xb4\xb3\xab\x42\x09\xb4\x54\x28\x89\x14\x49\x6a\xc0\x24\x50\x68\xcb\x32\xce\x1e\xdd\xb9\xce\xce\x19\xce\xcc\x9a\xb7\xeb\x7f\xbf\xac\x46\xbd\x57\x93\x1b\xbf\xcd\xf8\xbc\xcd\xb3\x4f\xe7\x87\x7c\x6a\x5c\x3e\x55\xa1\x12\xa2\x03\x57\x54\x1a\x37\x07\x4f\x21\x4a\xcf\xb4\x34\xc1\x90\x83\xa0\xd9\xf8\x08\x8c\xc1\x93\x0b\x66\x6a\x11\x66\xc4\x50\xa2\xb7\xf4\xd4\xe2\x15\x2c\xac\x41\x17\x21\x20\x2f\x8d\xc6\x4c\x74\x44\x07\x2e\xc8\x3f\xb1\x99\x57\x11\x4e\x2f\xce\xa0\xdf\xed\xf5\x65\xbf\xdb\x3b\xdf\x9a\x8c\x9c\xce\x7e\x01\x65\x2d\xac\x41\xa1\xd5\x47\x5e\x62\x99\x09\x11\x30\x82\xc4\x86\xc0\x1b\x8f\x33\x65\x6c\x1b\xee\xae\x6a\x7d\xad\xa5\x87\x96\xbe\x54\x6c\xd4\xd4\x62\x00\xc5\x08\x5e\x85\x80\x25\x2c\x8d\x82\x88\xcc\x6a\x46\x5c\xff\x14\x76\x20\x98\x5a\xd2\x8b\x4c\xe0\xa3\x27\x8e\x70\x75\x73\x39\x1a\xff\x55\xdc\xdf\x0e\x27\xe3\x3f\xff\x1e\xfe\x91\x9e\x9c\xbc\x1c\xdc\xfd\x26\x4f\x5e\x96\x8a\xb3\xc5\x3a\x6c\xa1\xb4\xa6\xc6\xc5\xc2\x33\xcd\x8c\xc5\xc2\x19\xbd\x70\xaa\xc6\xd5\x2a\xdd\x89\x5e\x8f\x86\xe3\xbb\xe2\x7e\x72\xbd\xd1\xdb\x1d\x0f\xa4\x36\x55\x15\x0d\xdb\xaf\xc8\xb7\x17\x93\xe1\x70\xbc\x23\xef\x8f\x07\xe4\xa0\x19\xd1\x6d\xc9\x22\xb2\xf2\x90\xa2\xae\x08\x8a\xab\xcb\xe2\xf2\x66\x3c\x2c\xe0\x33\x44\x44\x90\x0a\xf2\xa5\xe2\xdc\xd2\x3c\xd7\x96\x9a\x52\x1a\x67\xa2\xa4\x26\xfa\x26\x66\x96\xe6\x29\x0c\xff\x19\xdd\x09\x51\x2b\xe3\x4e\xcf\xe0\x45\x24\x6b\x9d\xb4\xd7\x3f\xcf\xba\x59\x37\xeb\xc1\x71\x2b\xab\x14\x06\x03\xc8\x31\xea\xbc\xa2\x10\x83\x10\x49\xa4\x46\x57\x7b\xab\xcd\xeb\x5a\x7d\xf8\x4f\x24\xb0\xfe\x7d\x3f\x87\x10\x49\xbd\x28\x0d\x83\xf4\x90\x93\x8f\xf9\xc2\x44\x7c\xd5\x11\x89\x6e\xd8\x82\x94\x96\xb4\x8a\xed\x16\xa5\x0c\xc6\xb6\x53\x93\x32\x54\xf4\x20\x91\x99\x18\xa4\x64\x8c\xfc\x04\xbf\xc2\xb7\xd5\xaf\x52\x90\xaf\x56\x90\xc7\xda\x6f\xe3\xcd\x9f\x45\x32\x7f\x36\x1e\xa4\x2c\x51\x53\xed\x19\x43\x00\x29\x67\xc4\x1a\x5b\xe9\x58\xd2\x11\x05\x06\x47\xf9\xf6\x31\xab\x9a\x4a\xf8\xf9\xf1\x5d\x44\x0b\xa1\x07\x07\x72\x02\x6f\xac\x6d\x95\xb6\x6d\x25\x87\xe4\xcd\xe5\xdb\xd5\x26\x1f\xb6\xfa\x4e\x14\x90\x35\x46\x55\xaa\xa8\x64\x13\x90\xdf\x89\xb3\x07\xb5\x7b\x7f\xfd\x84\x66\x9a\x6f\x76\x98\x6f\xff\xcd\x3e\x05\x72\xc0\x8d\x13\x22\x31\x33\xf8\x17\xe4\xe3\x11\xb6\x75\x91\x6b\x70\xa8\xe0\xff\xdf\x21\x56\xe8\x76\xf9\xdf\x86\x89\x64\x66\xc4\x6a\xb3\x4d\xf8\x71\x30\xf8\xe0\xad\x5f\x02\x00\x00\xff\xff\x6d\xf2\xee\x5c\xc6\x04\x00\x00")

func provisionShBytes() ([]byte, error) {
	return bindataRead(
		_provisionSh,
		"provision.sh",
	)
}

func provisionSh() (*asset, error) {
	bytes, err := provisionShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "provision.sh", size: 1222, mode: os.FileMode(420), modTime: time.Unix(1470666525, 0)}
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
	"provision.sh": provisionSh,
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
	"provision.sh": {provisionSh, map[string]*bintree{}},
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
