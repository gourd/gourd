// Code generated by go-bindata.
// sources:
// data/general.tpl
// data/rest/gorilla/pat.tpl
// data/service/upperio.tpl
// DO NOT EDIT!

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _generalTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x8e\xc1\x6e\x85\x20\x10\x45\xd7\xf5\x2b\x6e\xde\xea\x99\xb4\xba\xef\x0f\x34\xdd\xb8\xec\x9e\xca\x80\xa4\xca\x18\x1c\x6a\x8c\xf1\xdf\x0b\x95\xa6\x79\xb2\x1a\xe0\x9c\x03\xfb\x0e\x67\xd0\x7c\x50\xc0\x71\xb4\x2d\xde\xc8\x53\x50\x42\x1a\x9f\x1b\x2c\xc7\xa0\x71\xff\xa6\xb0\x38\xf6\x48\x6c\x01\xeb\x34\x92\xd7\x69\xaa\x4a\xa0\xe3\xf5\x1a\x50\xf2\x6b\x9c\x37\xff\x42\x62\x3a\x16\x7a\xc5\xbb\xc1\xc6\x11\xab\xf2\x02\x61\x04\x7a\xb1\xc5\x85\x0c\x6e\x81\x71\x23\xc1\xf9\xb4\x21\x98\x28\x31\xd0\x73\x96\xcf\xa5\x19\x9e\x05\xfd\xa0\xbc\x4d\x94\x34\x55\xfe\x89\xd0\x34\x8f\x39\x70\x9b\x55\xff\xa5\x2c\xdd\xd0\xe4\x37\x2b\x37\xcd\x1c\x04\xf7\x47\xe8\x3c\x5d\x0a\x54\x5f\x12\x3d\xeb\x3f\xff\xe9\x27\x00\x00\xff\xff\x8c\x75\x27\x05\x27\x01\x00\x00")

func generalTplBytes() ([]byte, error) {
	return bindataRead(
		_generalTpl,
		"general.tpl",
	)
}

func generalTpl() (*asset, error) {
	bytes, err := generalTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "general.tpl", size: 295, mode: os.FileMode(436), modTime: time.Unix(1440989564, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _restGorillaPatTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\x5d\x6f\xdb\xc6\x12\x7d\x96\x7e\xc5\x84\x17\x01\xa8\x1b\x81\xca\x05\x2e\xfa\xa0\xc2\x0f\xa9\xed\x14\x6d\x5c\xdb\xb5\x62\xf4\x21\x28\x02\x9a\x1c\x49\x84\x97\x1f\xde\x5d\xda\x15\x0c\xfd\xf7\xce\xcc\x2e\x25\x4a\x96\x1c\x3b\x9f\x4e\xa0\x00\xb1\x29\x72\x66\x67\xce\xd9\xd9\x33\xcb\x95\x6f\x6f\x21\xc5\x71\x56\x20\x04\x55\x9c\x5c\xc6\x13\x0c\x60\x3e\xef\x76\xfc\x07\xa0\xe7\xd1\xe9\xe5\x84\xef\xd1\x25\x16\x29\x5f\x75\x5b\x5e\x59\x5e\x95\xda\x1a\xe7\x15\x4c\x32\x3b\xad\x2f\xa2\xa4\xcc\x07\x93\x52\x67\x4a\xc5\x83\x2a\xb6\xc1\xfa\x93\x5a\xa7\x83\xa4\x4c\x31\xd9\xf8\xa4\x42\x9d\x6f\x7c\x60\x50\x5f\x67\x09\xf2\x33\x55\x4e\xf8\x57\x81\x76\x30\xb5\xb6\xe2\x6b\x63\x75\x52\x16\xd7\xfe\x32\x2b\x26\x26\xd8\x92\x33\x87\x96\x84\xbb\x83\x81\x20\x7c\x3b\xab\x30\x3a\x8e\x73\xa4\x9b\x67\x68\x2c\x5c\x64\x45\x6a\xc0\xc7\x03\x5b\x02\xa1\x00\x5d\xd6\x16\x75\x77\x5c\x17\xc9\x46\xaf\x50\xc3\x7f\xc9\x2e\x3a\x13\xbb\x3e\x5c\xc4\x06\xfb\x50\x94\x75\xe1\x7e\x56\xe0\xf2\xea\xc1\x6d\xb7\xdb\xa1\xd0\x3e\x1f\xf2\x99\x1a\xa2\x1c\x86\x7b\xe2\x03\x2f\x20\x18\x04\xf4\x53\x9c\xba\x1d\xdc\xfc\x44\x3e\xdc\x66\xe9\x3c\x60\x93\x6d\xde\x12\xe8\x26\x9e\x31\x88\x2c\xc5\xc2\x66\x63\xba\x9e\x22\x5c\xe2\x0c\x88\xaf\x34\xb3\x59\x59\x40\x38\x2e\x35\xd4\x55\x1a\x5b\x4a\x59\x23\xe5\x89\xd7\x08\x31\x71\x97\xa2\x42\x8b\xbd\x6e\x67\x82\xf6\x0d\xce\xf6\xc9\x85\x43\x31\x0b\x0c\x98\xd9\x8f\xce\xf0\xaa\x26\x02\x7a\x10\xf2\x88\x0d\x6f\x11\xdb\x1a\x46\xdb\xe1\x24\xde\x9e\x1c\x9c\x0c\x61\x82\x05\x6a\x8a\x22\x39\x90\xb5\xa5\x94\xa0\x1c\xd3\xc7\xcc\xc8\xa0\x9c\x8e\x73\x70\xff\xd2\x59\x11\xe7\x59\x12\x2b\x35\x83\x1b\xaa\x08\x90\x5a\x20\x8b\x4c\xf2\xd0\xd1\xf9\xd9\x51\xf4\x67\x8d\x7a\x16\xf6\xa2\x5f\xd1\x86\xc1\x30\x4b\x83\x1e\x30\x6c\x2a\x40\x48\xa6\x71\x31\x41\xb2\x27\x54\xb5\x2e\x16\xc9\x1d\xe3\x8d\xe4\x47\x5e\xaf\xd2\x34\x0c\xc8\xa9\x4f\x0c\x11\xd0\xf9\x76\xd2\xae\x38\x8e\x23\x4b\x65\x04\x58\x58\x91\xe0\x0f\xe5\x44\x8c\x1d\x27\x8f\x05\x70\xc5\xe6\xad\xf4\x1b\x1f\x4a\x9f\x91\xb4\x20\x38\xfe\xaa\x58\x53\x39\x18\x5a\x9e\x7c\x49\x75\xca\xd5\xdb\xe9\xf0\x8d\x91\xd5\xcb\xd8\xaf\x4b\x9d\xbb\xc0\x46\x96\x72\xcf\x1b\x19\x89\xe7\x96\x52\x34\xaa\x54\x66\x43\xef\xdb\x87\xa0\x2f\x66\x4c\xc4\xfb\xbe\x8b\xc1\xe3\x71\xa6\xe0\x7c\x19\x61\xe7\x8a\x93\x1b\xd1\x67\xf1\x64\x8f\xf9\x4a\x72\x55\x3c\xa1\xd1\xa9\xe0\x84\xab\x95\x34\xcb\xf1\xd8\xa0\xed\x13\xcf\x79\x66\x1b\x76\x61\x03\xbd\x25\xd9\x40\x9d\x15\xf6\xa7\xff\x3b\x5e\x3b\xa5\x69\xe0\x2d\xa1\xb9\xe1\x24\xe9\x8e\xda\xf0\x5c\xc2\xb8\xc7\xd9\x18\x64\x84\xbd\x3d\x08\x02\x37\xa2\xdc\xa3\x64\x50\x6b\x4f\x0a\x4b\x4d\x74\xca\x20\xce\x29\x74\xc8\x0e\x7d\xf8\xdf\xcb\x3e\x50\x16\x3f\x8b\x1d\xb9\x17\x99\xf2\xfe\x9d\x12\xf6\x68\x04\xb9\x9e\x77\x9b\x1f\x34\xaa\xda\x10\x49\x7d\x52\x24\x45\x91\xd4\x7a\x24\x57\xfc\xcc\x7f\xa8\x9b\xfa\x68\x96\x39\x97\x56\x34\x42\x7b\x22\x14\x85\x8e\xa9\x5e\x73\xf7\x88\x89\x09\x85\x1e\x71\xf4\xcb\xe8\x6a\xb1\x4e\xa8\x3c\x53\x45\x53\x49\xaa\x9d\x19\xc3\x62\xa2\x51\xd1\x0a\x4f\x81\x3e\xd6\x34\x3a\x3f\x79\xa5\x54\x79\xb3\x6d\x8d\xb0\xe2\x98\xea\x90\xa6\x57\x9a\x42\x44\x57\xf4\x9b\x50\xb6\xc6\x74\x75\x48\xe5\x5d\x8c\x4b\x88\xa2\x88\xa8\x40\x3d\x8e\x13\xbc\x9d\xf7\x98\x04\x2a\x44\xc6\x9f\x73\x0c\x76\xe3\x59\xfd\xa3\xfe\x87\xd1\x2e\x72\xce\x23\x49\x23\x5c\x19\xd9\x0d\x49\x23\xfa\x95\x4f\x9d\x25\x3a\xa5\x58\x76\x1c\x06\x67\x87\xa3\xb7\x22\xcf\x43\x78\x6e\x68\x6d\x55\x3d\x07\x79\x5f\x23\x01\xec\x76\x74\x74\x4a\xd3\x11\x56\x7d\x07\xeb\x06\x3c\x2a\x53\x95\x85\xc1\xbf\x74\x26\x7d\xe0\x4e\xc9\x72\xa2\xd7\xb1\x96\xb9\x93\xd4\xfd\x84\xc4\x94\xd0\xa4\xa0\x8e\x25\xf0\xbd\xfa\xca\xb5\x60\xb8\xb2\x07\x98\xf4\xe1\x3d\x43\x74\x44\x11\xc6\x03\x67\x70\x72\xd9\x20\x15\x22\xd7\xad\x3c\xa3\x64\x75\x43\xf9\x34\x05\x40\xda\xd5\xe8\x09\x2f\xf9\x45\xd1\x91\xc3\x7a\x7f\x73\xa3\x53\x6d\xb2\xc9\xb3\x56\xbd\xf9\x80\x7e\xce\xc2\x3c\xae\xde\xb9\xa9\xfa\xbb\x35\x43\xae\x32\xa9\x2d\xc7\xb6\x36\xc1\x10\x20\x10\xd8\x41\xdf\xdd\x97\x7e\x3c\x64\xb1\x17\x9e\x46\x62\xf6\x1b\xbb\x17\xb1\x1a\x51\x82\xa8\x0f\xd9\xde\x9b\xe7\x68\x0c\x6f\x54\x86\x10\x34\x46\xe0\xac\xe0\x70\x39\xec\xdc\x2d\xf5\xd6\x6c\xca\x43\x28\x2f\x6c\x9c\x15\xac\x3a\xeb\x18\x1b\x2e\xfc\x64\x53\x8a\x91\xb8\x84\xbd\xde\xea\x12\xa2\xff\xd4\xba\x29\x9c\x89\xf6\x55\x69\x30\x6c\x08\xa5\x26\x55\x26\xdc\xdc\xb8\x65\xd8\x19\xdd\x44\x59\xc4\x52\x77\xc9\xa1\xdc\x5c\x18\xbb\xa9\x6d\xb4\x8f\x6d\x79\x29\x83\x9f\xe7\xc8\x4d\x6c\x88\x5b\x78\xbf\x0b\xec\xf7\xd1\xc9\x31\x9c\x17\x39\x29\xc5\x34\x56\x5b\x41\x7c\xa1\xe9\xfa\x25\x4e\x9b\xd5\x7c\x67\x96\xf6\xe3\xa2\x28\xed\x1a\xe0\xf6\x2c\xb5\xa8\x75\xdc\x24\x53\x4c\x2e\x5b\x8b\x74\xc9\x81\x5b\xdd\x8b\x75\xbc\xa8\xf7\x20\x91\x35\x09\xc1\x0b\xb7\xe1\x42\xaf\x8d\x6d\xce\x38\x7c\x1f\x72\x33\x69\x37\x52\x51\x56\x47\x11\xd9\x7f\x01\x92\x24\xea\x3a\x27\x94\xc4\xbd\xf8\x1d\x98\x65\x19\x09\x72\x2a\x37\xb9\x1f\x12\x22\x46\xf8\xd0\xca\x10\x2f\xae\xf8\xe7\xc6\x97\x85\xe7\xe8\x29\x14\xc7\xeb\x38\x53\xd4\x2b\x68\xa7\xb5\x02\xfa\xde\xf2\x70\x0a\x29\x93\xcf\x42\xdb\x7d\x4c\xea\xcb\xcc\x03\x53\x27\x09\x25\xe2\x42\x2d\x33\x6f\x25\x7e\xf2\x46\x9e\xc9\x56\x5a\x40\xbd\x5b\x19\x0f\xe7\xfc\x98\x93\x9c\xfb\xce\x70\xd6\x6c\x9d\x0d\xc5\x56\xd2\x22\x78\x7b\x81\x5f\xa9\x45\x7c\x65\xf1\xff\xac\xea\xfa\xa4\xfb\xc8\x63\xc4\x3f\xc7\xbc\xa4\x17\x02\xde\x1b\xd3\x14\x66\xf1\x85\x42\xc3\x6b\x58\xdd\xe9\x05\x47\xf4\x0e\x11\x6e\xd8\x8e\x25\xfe\x25\x6b\xf9\xca\xe5\xa6\xa0\x91\x81\x11\xc6\x3a\x99\xca\x8b\x05\x31\xa9\x1e\x3c\x3b\x46\xfc\x9e\xbe\x12\xd0\x6b\x71\xfa\x71\x3a\x40\x34\x98\xe8\x08\x8b\x90\x58\xe1\x7d\xf1\xcb\x4f\xd8\xa5\x6c\x81\xd3\x42\x73\x5c\xda\xd7\xc4\x5f\xba\xcc\x71\x4e\xd3\x41\xef\x35\x1f\x6a\x56\xaa\x8c\xd3\x65\xab\x52\xdf\x67\xaf\x72\x50\x3f\x81\xdf\x15\xfd\xbd\x57\x80\x5b\x0a\x8c\xaa\xc5\xf6\x06\xe1\xe5\xf7\xf2\xa5\xec\xee\x74\x77\xa7\xbb\x0f\xd4\x5d\x77\xb8\xb3\xaa\xbe\xee\x88\x45\xa6\x61\xe5\xbd\xb0\x5d\x6c\xcf\xe8\x1f\x3c\xff\xcf\x35\x91\xcb\x9e\x32\xec\x4e\xa8\x1f\x20\xd4\xf7\x2b\x24\x31\xfb\x83\x28\xe4\x67\xed\x49\x1f\xa5\x99\xb2\x6b\xbd\x9d\xdf\x69\x53\x4f\x44\xbb\xcf\xe5\xdc\x59\x8e\x53\xea\xaf\xb7\x57\x7e\x92\xc7\x29\x3b\x65\x7f\x8c\xb2\x6f\x3a\x5c\xf9\xb0\xe0\xef\x0e\x5e\xb6\x0b\xb5\x08\xfa\x9d\xaf\x66\x48\xc3\xd2\xdd\x6b\xc9\x83\xbb\x9d\x95\xef\x32\x1e\x73\x78\xe5\xbe\x79\xfb\xee\xfb\x9d\xc7\xef\xd1\x2c\x4e\xaf\x4c\xe4\x14\xbe\xa9\x8c\xde\x37\x3c\xc6\x11\x50\x80\xeb\xe7\x36\x07\xf2\x2d\x27\xb7\x20\x77\xb5\x3b\xb1\xd9\xf5\x97\x8f\xed\x2f\x3b\x0d\xfd\x26\x1a\xea\xfe\x50\xe1\x47\xd1\x50\x8f\xa6\xa5\xa1\x5e\x98\x16\xef\x98\x4f\x47\x40\xdd\x5f\xd7\xf8\x3f\xb4\xf9\x37\x00\x00\xff\xff\xd0\x8e\x52\x0b\x52\x24\x00\x00")

func restGorillaPatTplBytes() ([]byte, error) {
	return bindataRead(
		_restGorillaPatTpl,
		"rest/gorilla/pat.tpl",
	)
}

func restGorillaPatTpl() (*asset, error) {
	bytes, err := restGorillaPatTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rest/gorilla/pat.tpl", size: 9298, mode: os.FileMode(436), modTime: time.Unix(1440988415, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _serviceUpperioTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x58\x4f\x6f\xdb\xb8\x12\x3f\x4b\x9f\x82\x35\xf0\x0a\x29\x08\xe4\x1e\x8a\x1e\x02\xf8\x50\xc4\xed\x43\x80\xbc\xa6\x2f\x49\xf7\x52\x14\x85\x2c\x8d\x65\x22\x34\xa9\x92\x94\xbb\x46\xe0\xef\xbe\x33\x24\xa5\xc8\x96\x9c\x26\xdd\xec\x61\x0b\x34\x96\xc9\xe1\xfc\xf9\xcd\x6f\x66\x28\xdf\xdf\xb3\x12\x96\x5c\x02\x9b\xd4\x79\x71\x97\x57\x30\x61\xbb\x5d\x1c\x85\x2f\x0c\xf7\xb3\xcf\x77\x15\xad\xe1\x23\xc8\x92\x9e\xe2\xde\x29\xbe\xae\x95\xb6\xc6\x9f\x9a\x48\xb0\xd3\x95\xb5\xf5\x04\x9f\x2b\x6e\x57\xcd\x22\x2b\xd4\x7a\x5a\xa9\x46\x97\x53\x03\x7a\xc3\x0b\x78\x6c\x6f\xda\xd4\x35\x68\xae\x50\x06\x6d\xf0\x25\xcb\x2e\xca\xec\xc2\xdc\x58\xcd\x65\xe5\x4d\xf4\x8e\x9a\xdc\x2a\xcd\x51\x43\xd6\x34\xbc\x24\xbd\x20\x0b\x55\xa2\xe8\x74\x91\x1b\x78\xf7\x96\x96\x8c\x3b\x6b\xbc\xc6\x10\x40\x34\x11\xaa\xa2\x4d\x67\x2e\xe3\x6a\x5a\x2e\x26\x47\x02\x44\x85\x1e\x93\x78\xd9\xc8\x82\x71\xc9\x6d\x92\xb2\xfb\x38\x9a\x4e\x5b\x99\xe0\x3c\xab\xb5\xda\xf0\x12\x34\xfb\x89\x3e\xd2\xb7\x3f\xb7\x71\x14\x36\xb3\xcf\x61\xd3\x64\x73\x77\xea\x23\x6a\x4b\x26\x84\xef\xed\xb6\x86\xec\x53\xbe\x06\xb4\x32\x39\x65\xce\x4c\xa2\xd9\x09\x01\x99\x5d\xc3\x8f\x06\x8c\x4d\x59\x62\x5a\x3b\xd9\x8d\xff\x3c\x65\xa0\x35\xfd\x57\xda\x39\x14\x69\xb0\x8d\x96\xec\xbf\x60\x0f\xd5\x86\x13\x89\x4e\xe3\x68\x97\xc6\x18\x0c\x7a\x7f\x5c\xae\x8d\xc4\x30\x9d\xff\x64\x47\x84\x3c\x1e\x8f\x19\x1b\x09\xe1\xe4\x88\xf0\x41\x2c\x0e\x5d\xb5\xb0\x39\x97\xac\xcc\x6d\x4e\xe9\x8c\xa3\x72\xe1\xc5\xce\x66\x2c\xf0\x24\xbb\xaa\x41\x26\xfa\x94\x4d\x30\x15\x79\x23\xec\x04\x03\x44\xda\x90\xd4\xab\x19\x93\x5c\xf4\x80\xc1\xd0\xe3\xb1\xb4\xe5\x98\xf4\x56\xc2\xb0\x19\x7b\x7d\xc4\xc7\xfb\x72\x81\xd4\x09\x92\x1e\xc2\x63\xf8\x91\x6a\x44\xaf\x02\x89\x5e\x16\xec\xfc\xcb\xf5\x9c\x2d\x95\x66\x16\x65\x07\x87\x7c\x2e\x50\x32\xb7\x50\xb2\xc5\x96\xb9\x92\x60\xe7\x97\x17\xcc\x2a\x25\xe2\xd1\x43\x9d\x25\xab\x9b\xc2\x52\x94\xf3\x05\x2b\x17\xd9\xbc\x45\xcb\x3b\x78\xae\x01\xb5\xb2\x7c\x70\x1e\x89\xcc\xec\x0a\x3a\x74\x4f\x99\x5a\xba\x85\x3a\xd7\x20\xad\x4f\xee\x23\x09\x4b\x83\xea\x24\x8e\x0a\x85\x00\xb6\xdc\x3c\xc7\x2f\x06\xd3\x54\x77\x2b\x1f\xa4\xe5\x76\xfb\xd9\x62\x62\x93\x61\x92\x2b\xb0\xac\x50\x42\x40\x61\xb9\x92\xa4\x4c\x88\x2e\xcb\x06\xd5\x09\x91\x3c\x25\xa7\x79\x5d\x8b\x2d\xd2\x55\x96\x6a\xcd\xa8\x1d\x30\x5f\xf8\x08\x61\xfb\xc4\xcb\xa3\x5d\x85\xe4\x89\x55\xf8\x99\x7d\x82\x9f\x7f\xbc\x25\x9b\x40\x4b\x50\x67\x09\x62\xf0\x00\xc1\x6e\x47\x5b\x19\xc1\x82\x5a\x5a\x34\x67\xc1\x88\xc9\x6e\x35\x5f\x5f\xf3\x6a\x65\x13\xdf\x83\xb2\x2f\xd7\x97\x1f\x42\x5b\xca\xdc\x03\xdc\x2a\x6f\x3a\x41\x73\x5f\xcf\xbe\xa5\x48\xdf\x19\x11\xb7\xd7\x80\x28\xa6\xdb\xab\xf9\xd5\x19\x82\x23\x37\xa0\x09\x24\xdc\xe2\x12\xc3\xf1\x19\xba\x83\x2d\x31\xd7\x05\xef\xfe\x81\x44\x86\x21\x23\x50\x82\xf2\x08\x0e\xf6\x00\x4e\x59\xf6\xd6\x48\xa2\x0f\xf9\x11\x4c\xbe\xfb\x34\xcc\x9c\x6c\xf6\x1e\x0b\x4e\x96\x09\xd4\xc1\x4f\x61\x1c\x75\x23\x5e\x76\xd9\x1a\x95\x0b\x0d\x77\x98\x3f\xec\xc0\xd8\x13\x31\xa0\x65\x32\xf9\x40\x94\x60\x05\x11\x8a\xac\x1f\x52\xee\x8c\xfd\xc7\x4c\x9c\x99\xcc\x49\x26\x29\x2a\x8f\xbc\x73\x1d\xcb\x68\xe3\x42\x5a\xd0\x32\x17\x07\xf4\x18\x0d\xb0\x1f\x44\xc7\x1f\x02\x89\x80\xdd\xc7\xf0\x99\x4c\x08\x0b\x24\x87\x0b\x09\x52\x2a\xc1\x30\xdf\xbd\x4d\xf7\x53\xbc\xdf\x49\x6e\x20\xd7\xc5\x6a\xac\x50\xb1\x21\x70\x6b\x5c\xfe\x39\x25\x2c\x31\xe9\x13\xaa\xd3\xeb\xc3\xea\xfc\xd1\x41\xf4\xff\x06\xf4\xf6\x94\x89\xc3\xd2\xbc\xe4\xc6\xfe\xd3\xe5\x49\x7a\x04\xda\x79\x08\xc3\xb5\x5d\x5e\x49\xa5\xc1\x63\x4d\x76\x7d\x33\x39\x65\xdf\x1d\x9f\x32\x1c\x2e\xff\xcb\x6b\x34\xe0\x74\xa0\x4a\xcd\x61\x83\xdd\x4c\x08\xd6\x60\x0c\x26\x8e\x36\xb9\xc6\x75\x43\x9d\xef\x1a\x0c\x0e\x00\xe7\x8b\xc0\xa1\x40\x8a\x52\x36\x9b\xb1\x37\xc1\x1f\xd3\x52\xf9\x23\x47\x82\xd2\x20\xf4\x04\x18\xdb\x6d\xc7\x8b\xeb\x66\xc9\x0f\x72\xc4\x3f\xa6\x94\xc4\x10\xd3\x0a\x23\x10\xd4\x30\x2b\x64\x94\xb3\xeb\x04\xaf\x96\x4b\x03\x74\x47\x78\xb5\x6f\x1b\xff\x66\x37\x77\xbc\xc6\xaa\x97\x36\xd9\x13\xf5\x4a\x3b\x0d\x97\x7c\xcd\x8f\x28\xf0\x5b\x0f\x1a\x82\x68\xcf\x2b\x42\x9a\x00\xd2\x0e\x0e\xe3\xd6\xba\x4e\x62\xdc\xf5\x04\xd9\x8d\xf0\xb1\x9a\xd7\x20\x68\x10\xd6\xb9\xa5\xba\x89\x43\x49\x91\x9d\xf7\x98\x57\x51\x8f\x66\xf6\xd1\xba\x7b\xe0\x35\xc9\x07\x6e\x5f\xa1\x0d\xbf\x68\x5c\xaa\x97\x5c\x23\x13\x06\x44\x5f\xe7\xb6\x58\xc1\xb3\x99\x8e\xda\x69\x08\xfd\xf6\x04\xea\x68\xb5\xd4\x38\x3a\x1e\x6e\x1b\x82\x38\xf8\xfa\xeb\xb7\x43\xc3\xf7\xbb\x16\x28\xcc\xa7\xaf\xb2\x02\xab\xea\x29\x55\x80\x02\x52\x59\xbc\x09\x34\x44\x72\x0d\x74\x75\xee\x08\x7b\x22\x7a\x74\x1d\x01\xf9\x93\xb2\x1f\xe9\xe0\xc8\xec\x33\x06\x0b\xc9\x41\xbb\xc9\x45\x03\x34\xca\x2b\xbe\x01\xc9\x6a\x85\x4c\x71\x32\xa1\xa3\x79\xe8\xdb\x90\x4b\x2f\x1f\x47\xc9\x49\xe8\x6e\xfb\xa1\xa6\xe8\x11\x23\xc7\xbe\xbe\xf9\x36\x92\xd7\x2f\x75\x49\x97\x8b\x41\x22\xb1\xb8\x9f\x99\x43\xaf\xe9\x6f\xa5\xf1\x25\x3b\x15\xb6\xdc\xdf\xeb\x53\xae\x52\xcf\xfa\x9d\x04\x1b\x13\x45\xe2\xdb\x51\x68\x64\x8d\xc7\x8d\xb4\x79\xce\x97\x7e\xd2\x70\x30\xfd\x1a\x0c\xa0\xc0\x78\x19\x0e\xe7\xa7\x53\xfb\x42\xf3\xf3\xf0\xa2\x3b\x07\x01\x2f\x93\x6a\xaf\x69\x98\xea\x7f\x7b\x52\x35\xac\xd5\xe6\xd7\x49\xbd\x76\x62\xe3\x4e\x0f\x53\x5a\x12\x58\x2f\x9d\xd2\x5e\x05\x63\x9b\x57\x85\xaf\x2b\x1a\x1a\xaa\x20\x5e\xae\xd1\x45\xbd\x75\xef\x2b\xb9\x6c\xef\x40\xbf\x4e\x6c\x4f\x17\xce\xae\x41\xd1\x52\x84\xc1\xfe\xe0\x05\x0b\x5b\xea\xc0\x1f\xba\x95\x3c\xee\x93\xbb\x50\x3c\xcf\x31\x52\x3a\x70\x2e\xdc\x7f\xfa\x0e\x8e\x76\x7d\xef\xe2\x25\xb6\x55\x2e\x4d\x8d\x94\x74\xc9\xc6\xd6\x5d\xe1\xdb\x3e\xb6\xdc\x67\x3b\x86\xaa\x92\x5a\x1c\xbb\x8d\xb9\x3b\x23\x39\x05\x6e\x12\xd5\x02\x1b\xf4\xd0\xaf\xb4\x73\xda\xc9\x27\x6e\x92\x80\x48\xdb\x97\x7c\x2a\x91\x30\x7c\x9d\xbb\xf4\x42\xdf\xfe\xea\xd1\x2f\xad\x27\xbc\xef\xb9\x62\x63\x09\x1d\x62\xae\x02\xda\xc3\x87\x3f\x44\x84\x82\x23\x53\xe3\xc5\x4b\xb5\x3b\xef\x6b\xa0\x1f\x42\xdc\xd7\xdd\x6e\xfc\x0d\x7e\xe4\x6d\x41\x49\x49\x87\xb1\x38\x1e\x8c\xb0\x4e\x8d\x2f\x10\x3c\x19\xfd\x5e\x8d\xb4\xf0\x09\x65\x60\xef\x45\x19\x8f\xe2\xb8\x55\x84\x66\x3e\x72\x89\xe1\x06\xef\x56\x74\x1f\x7c\x02\xa0\xa4\x1b\x11\x75\xc0\xf5\xd8\xe7\xb1\xf1\x9b\xe4\x46\xef\x87\xa9\xbf\x02\x00\x00\xff\xff\x1a\xef\xae\x80\xaf\x13\x00\x00")

func serviceUpperioTplBytes() ([]byte, error) {
	return bindataRead(
		_serviceUpperioTpl,
		"service/upperio.tpl",
	)
}

func serviceUpperioTpl() (*asset, error) {
	bytes, err := serviceUpperioTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "service/upperio.tpl", size: 5039, mode: os.FileMode(436), modTime: time.Unix(1440988991, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"general.tpl": generalTpl,
	"rest/gorilla/pat.tpl": restGorillaPatTpl,
	"service/upperio.tpl": serviceUpperioTpl,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"general.tpl": &bintree{generalTpl, map[string]*bintree{
	}},
	"rest": &bintree{nil, map[string]*bintree{
		"gorilla": &bintree{nil, map[string]*bintree{
			"pat.tpl": &bintree{restGorillaPatTpl, map[string]*bintree{
			}},
		}},
	}},
	"service": &bintree{nil, map[string]*bintree{
		"upperio.tpl": &bintree{serviceUpperioTpl, map[string]*bintree{
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

