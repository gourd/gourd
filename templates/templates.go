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

	info := bindataFileInfo{name: "general.tpl", size: 295, mode: os.FileMode(436), modTime: time.Unix(1445626823, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _restGorillaPatTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5b\xeb\x6f\xe3\xb8\x11\xff\x2c\xff\x15\x3c\x15\x5b\x48\x7b\x3e\x79\x0b\x14\xfd\x90\x22\x05\x6e\x93\xec\xde\x15\xfb\x48\xf3\x68\x3f\x2c\x82\x42\xb1\x68\x5b\x0d\x2d\x2a\x14\xed\xc4\xc8\xe5\x7f\xef\xcc\x90\x94\x28\x5b\x7e\x64\x37\xc1\x6d\xda\x04\xd8\xac\xc4\xc7\x70\xf8\x9b\x27\x47\xcc\xdd\x1d\xcb\xf8\x28\x2f\x38\x0b\xcb\x74\x78\x95\x8e\x79\xc8\xee\xef\x7b\x81\x7d\x61\xd0\x9f\x1c\x5f\x8d\xb1\x0d\x1e\x79\x91\xe1\x53\xcf\x9b\x95\x4f\x4b\xa9\x74\x65\x66\x85\xe3\x5c\x4f\x66\x97\xc9\x50\x4e\x07\x63\xf9\xd3\x55\xae\x07\xf8\x0f\xa6\x95\x32\x2f\x74\xd8\x0b\x26\x5a\x97\x5a\xa5\x45\x85\xb3\xd8\x9a\xf1\xf5\x80\x01\x0e\x0f\x97\xe9\xaa\x5c\x88\x74\x50\xa6\x48\x6f\x2c\x67\x2a\x1b\xea\xdb\x25\x52\xd0\x38\x18\xca\x42\xf3\x5b\xbd\x32\x1d\xfb\x4a\xae\xa6\x9d\x1d\x15\x57\xf3\x7c\xc8\xa9\x4f\x8a\xb4\x18\x27\x52\x8d\x07\xb7\x83\x82\xeb\x86\x20\x74\xf2\x62\x28\xb3\xbc\x18\x0f\xfe\x53\xc9\x02\x47\x8f\xa6\xb4\x92\x90\x63\xfc\x0f\x87\x3b\xde\x2b\xad\x60\xe6\xdc\x3e\xc2\x9c\x2a\x5c\x83\x25\x90\x34\xf0\xf7\x06\x03\x42\xfe\xd4\x70\x93\x7c\x4a\xa7\x1c\xda\x8f\x2c\x8e\x15\x53\x5c\xcf\x54\xc1\x0e\xce\x4f\x0e\x19\xaf\x5b\x47\x52\x75\x4d\xeb\x8d\x66\xc5\x70\x23\xbd\xa8\x90\xb3\xa2\xcf\xf0\x77\xc9\x0c\x93\x31\x8b\x1a\xc2\xd3\xb4\xfc\x62\x9a\x2f\x5c\x63\xe2\x26\xc7\xec\x0e\x00\x01\x86\xe7\xa9\xca\xd3\x4b\xc1\x2b\xa6\x25\x9b\x55\x9c\x89\x54\x73\xd5\x0b\x52\x21\xe4\xf0\xa8\xd0\xb9\x5e\xb0\xbd\x7d\x86\xcc\x44\x31\x7b\x8d\xfc\x9c\x2d\x4a\xe4\x83\xdd\xb9\x0d\xfd\xd1\x6b\xbd\xbb\x67\xf7\xad\xd9\x1f\xf2\x4a\xfb\x14\xbe\x5c\x74\xd3\x68\xb5\x1b\x2a\x63\xae\xed\xde\x91\xc0\x7b\xae\xbb\x60\x0a\xac\xf0\xe9\x1d\x86\x85\x1d\x83\x42\xb3\xd7\x4a\x4b\xc5\x3d\xe4\x27\x5c\x71\xea\x38\xfb\x7c\xf8\x79\x0f\xf0\x5a\xb0\x49\x3a\xe7\xac\xe0\x37\x88\xe7\x6c\xa8\x11\x14\x9a\xd5\x0b\x9a\x69\xfb\x30\xf2\x8a\x47\x1b\xe1\xed\x79\x13\xbe\x84\x43\xc5\x01\xd5\xf0\x82\x59\x18\x50\xf5\xad\x5e\x26\x07\xe6\xff\x3e\xe3\x0c\x06\x73\x35\x4a\x87\xfc\xee\x1e\x04\xa9\x40\x26\x5e\x0b\x0c\x50\x0a\xff\x49\x65\x84\x87\x8c\x03\x42\x8e\x10\x8c\x05\x4d\x9a\xa6\x3a\x97\x05\x74\xaa\x3e\x93\x57\x88\x87\xb3\xb5\xe4\x97\xb3\xb3\xe3\x13\x7e\x3d\xe3\x95\x46\x06\x62\x18\x94\x8f\xd8\x0f\x30\xea\x0e\x1e\x11\x46\x85\xe3\x2d\x9c\xc9\x11\xae\xf4\x2b\xae\x5f\xa4\xc2\x0d\x20\x60\xb9\xfa\x58\x8d\x61\x2b\xe1\x34\xaf\x2a\xd8\x3e\x88\x90\xa8\x02\x07\xac\x31\xdf\x20\x40\x82\x44\x4f\xe1\x9b\x91\x33\x3c\xdd\x7b\xbc\xdb\xc5\xa0\xa1\x32\xfb\x43\x86\x6b\xa9\x47\xca\x32\x89\x3d\x3f\xec\xb3\x22\x17\x5f\xcb\x2b\x18\x7a\x72\x5a\x82\xb0\xf4\x28\x0a\x09\x44\x26\x2f\x75\x9a\x17\xc8\xff\xab\xca\x51\x62\xd1\xab\x2a\x0e\xfb\xcc\xd3\x29\xe2\x2b\xde\xb8\x9f\x20\x00\x5f\xc0\x15\xab\x92\x03\x21\x2b\x1e\xc5\x76\x87\x46\xec\xa0\x70\x68\x07\xd0\x04\x6e\x26\x39\xb6\x3c\x98\xbe\x3d\xf6\xea\x0f\x73\x58\x8f\xe3\x02\x96\x7e\x72\x40\x5d\x11\xec\xd6\x76\x3c\x3a\x02\x38\x20\xb0\x30\x10\x23\x06\x05\xe2\xa6\x6f\xf9\x75\xac\x19\x27\x83\xd4\x68\x8d\x28\xb6\x4c\x6d\x13\x2f\xf9\x5a\x0e\xba\x01\x31\xa1\xa8\x50\xc4\xa8\xd0\xfb\xec\xcb\x85\xa7\xd4\x77\xfc\xbe\x57\xcf\xed\xd1\x5c\xcf\x6c\xa0\x5d\xe5\x7c\xbe\xcd\x70\x1a\xf5\x7b\x90\xf9\x5c\x23\x7e\x76\x6e\x12\x39\x20\xff\x31\xe3\x6a\x11\xbf\x98\xd7\x33\x30\x2f\x0a\x32\x68\x60\x53\x3e\x95\x6a\x41\x81\xb4\x0e\x67\x68\x4f\x02\x37\xb1\x14\x8a\xea\xd9\x4e\xb9\x3c\xc3\x3b\xe5\xa9\x1a\x4e\xa2\xeb\x38\xf9\x59\x88\x88\x8b\xa7\x84\xa7\xa2\xb5\x6a\xbb\xab\xc2\x3e\x19\x65\x6d\x6d\x5f\x67\x62\xc0\x6d\x95\x7c\xe0\x05\x32\xcf\xf6\xf7\xd9\x1b\xc3\x71\x4d\xa7\xe1\xf7\x93\xd4\xef\x60\xad\x6c\x95\xb0\x31\x53\x2e\x36\x18\xa6\x00\x20\x5f\x8c\xf2\xc5\x28\x1f\xd7\x28\x9f\xa5\x19\xee\x60\x2d\xb3\x32\xdb\x9e\xfd\x7d\xa5\xbd\x80\xf2\x97\x2d\x93\xf1\xb2\x52\x9f\x54\x6d\x5a\xd8\x1f\x5e\xa3\x3d\x85\x17\x2b\x06\x06\x5b\x6d\x06\x99\x34\x20\xbc\x80\x56\xe0\x36\xc3\x8e\xeb\x04\x92\x70\x60\x39\xab\xa2\x17\x6b\x7c\x0e\xd6\x68\x54\xaf\xc9\x40\xed\x4e\xd0\xca\xce\xa9\x2b\x42\xd1\x62\x46\xf7\xd7\x27\x4c\x31\xdd\xb1\x1b\xb7\x89\xa1\xe3\x1b\x0c\x6d\x7d\xf6\xb8\x6c\x75\x19\x17\xfc\xa9\xac\x6e\x17\x3f\xb7\xe4\xe6\xa2\x78\x6b\x46\xe2\x34\x80\xac\x93\xc9\x11\xcb\xb3\x6d\x11\xf1\xc5\x34\x9f\xad\x69\x8e\x72\x10\x9c\x9e\x70\x83\x4a\xa1\x51\xe0\xf8\x4a\x42\x7f\x96\xa1\x10\x36\x65\x8c\xae\xd3\xdf\x1c\x52\x17\xf9\x9b\xe7\xef\x6c\x8c\xb7\x31\x2f\x41\x6f\x6d\xe9\xef\x04\x15\xff\x12\x04\xdd\xe8\x90\x96\xac\x4c\x35\x53\x72\x86\x65\xb6\x75\x35\x3e\x9c\x18\x29\xf6\x1a\x86\x26\x27\x34\xb4\xcf\x2e\xd3\x8a\xbb\xe3\x78\xbb\xe6\xf7\x7f\x58\xca\xe3\xe8\xcf\x86\x4e\xd3\x58\xa9\x64\xc9\x15\x3c\x5c\xf2\x11\xd6\xf8\x4c\x75\xa5\x17\x00\x24\xd0\xc7\xcb\x54\xf1\x03\x57\x8c\xb1\xa5\xba\x8f\x79\x96\x09\x7e\x03\x3d\x2e\x42\xe4\x45\xc1\x15\xeb\x28\x95\xae\x34\x91\xce\xda\xdd\x92\x08\x1f\x18\x5d\x4a\x74\xda\xeb\x23\x4c\x40\xb6\x54\x0a\xe8\x9b\x48\x91\x71\xb5\xc7\xd2\x62\xa1\xc9\x36\x17\x72\xc6\x6e\xd2\x82\x2a\x93\x99\x64\x37\xb9\x9e\x90\xdb\xa8\x4d\x0e\xa7\xb6\x7e\x2c\x24\x69\x59\x62\xf1\x1a\xa7\xa5\x3a\x45\x65\x6a\xd4\x9e\xd1\xd6\x71\x13\x35\xd3\xb1\x71\x5e\xf7\x3b\xc1\x6d\xd2\x8c\x16\xdc\xe7\x2e\xf3\x78\x14\xb8\x59\x1b\xef\x87\xc3\x5d\xf1\xcd\x11\x7d\xf7\x44\xda\x42\xbc\x3e\xb8\xee\x16\x5d\x5b\xe1\x75\xbb\xfb\x7b\x78\x80\x5d\xf2\x6f\x8d\x83\x33\x1e\x6e\x6d\x2e\x12\xec\x74\x52\xf0\x50\x68\xc2\xf4\x86\x38\xdd\x15\xb5\xbe\x66\xd7\x8f\x11\xaa\x37\x21\x13\x74\x86\xeb\xad\xf1\x7a\x63\xc0\x7e\xc2\xbd\xaf\x89\xd9\x4b\x91\x6e\x9b\x2a\xc0\xee\x34\x17\xc2\x6c\x07\xad\xd2\x32\x2d\x78\x11\xbd\xc6\x4a\xd2\xdf\x5c\x21\x29\x30\x5a\x01\x36\x3e\xa7\xcc\x9a\xba\xbf\xbc\xb9\x68\xd1\x7a\x64\xc7\x65\xcf\x30\xdb\x1d\x17\xb0\xd6\xf6\x5a\xd6\x15\x51\x1a\xd0\x72\x4e\x14\xc8\xd6\xbb\x26\xf6\x0d\xa1\xe0\x91\x5d\x13\xed\xd4\x8c\xea\xdb\x54\xaa\xdb\x57\x77\x2a\xd9\xb2\xa0\x85\x0d\xe0\x8e\x64\x12\xb5\x63\xb8\x23\x44\xe3\xf6\x0d\xa5\xdf\x7e\x63\xaf\xfd\x77\x43\xd9\x36\x99\x2f\x61\x2d\x1a\x7d\xf6\x26\x7e\x12\x75\xa0\x15\xbd\x5a\xa7\x43\x1d\xdb\xfb\xc8\x5a\x4b\xf6\x37\x0a\xfc\xb9\x11\x63\x0d\x33\x2d\x02\xe6\x9d\xce\x84\xc6\x18\xa6\xe5\x50\x8a\x96\x62\x1c\xdb\xc6\xe7\xa1\x1c\x08\xc2\xbc\x76\xb9\xdf\xa0\x17\xf5\x5a\x8d\x3f\xfa\xc4\x6f\x4e\x6c\x2b\x7d\x67\x2e\xfb\x6c\x1e\x7b\x89\x72\x0b\xec\x31\x87\xa5\xc1\x48\xab\x7a\x2f\xf8\xa9\x1e\xe3\x93\x84\xa0\x34\xe1\xc3\x2b\x00\x6c\x5a\x23\x89\xe9\x60\x61\xe3\xe1\x31\x0c\x3c\xb0\x23\xf6\x1c\xc2\xde\x6c\x97\xe3\x76\x09\x64\x25\x0b\x7b\xa0\x64\x9e\x58\x34\x5b\xd2\x84\x1d\xf3\x84\x76\xa2\xb0\x43\xdc\xf8\x8a\x54\x61\x39\x40\x78\x0a\x62\x35\x24\x98\xe2\xa2\x28\x17\xac\x33\x7c\x9c\xdd\x46\xad\xd0\x32\xc5\x70\x27\x6f\x22\xd8\x52\x23\xbb\xb6\x1e\x76\x2a\x62\xc7\x42\x1b\x33\x52\x8a\xd0\xdd\x7a\x67\x45\xb2\x8b\xe2\x99\xa1\xff\xcb\x9a\xb7\xd5\x2b\xec\x2a\x8d\x17\x15\x36\xeb\x39\xef\xb7\xbb\x16\xd7\x1e\x75\xee\xbb\xdc\x15\x1d\x26\x3d\xde\x87\x1f\x7a\xe9\x99\x22\x0a\x7e\xd6\xfb\xf5\x10\x60\x34\x1a\x9e\x0f\x21\x4d\x5f\xd8\xf6\x8c\xed\xe5\x19\x64\xa3\x5c\x64\x34\x3a\xba\x91\xea\xaa\x32\xb1\x0d\x2b\x0b\x98\x20\x65\xec\xfc\xe4\x03\xd5\x18\x50\x83\x4d\x84\xf3\xa9\xb6\xee\x73\x25\x87\xd4\x63\x85\xf6\x0e\xd5\xd9\x9e\xd0\x14\x7b\x8d\x23\x13\xdb\x45\xfa\xb7\xa2\x95\x2b\x09\x4b\x4e\xf5\x48\x95\x00\x0b\xe6\xb4\x10\xc5\x08\x78\x14\x02\xdf\x61\xcc\x30\x36\xe7\x90\x6c\x0e\x27\x69\x31\xc6\x38\xee\x2a\x98\x5e\xe4\xb1\x75\xcc\xe4\xe7\x2c\x8b\x42\x98\xd5\x87\x2c\x3b\x26\x93\x33\xcb\xb7\x06\xbb\x35\x4e\x5d\xfd\x93\x4a\x4c\xbd\x76\x61\xb8\x86\x15\x53\x3f\x84\xc0\xbc\xd9\x5a\x2b\x16\x6f\x85\x9f\x13\xfa\x90\xb9\x19\x4f\x08\x5a\xcf\x9e\xb9\x56\x77\x65\xeb\x6a\x90\x9b\x80\x2e\x55\x78\xfb\x0e\x1e\xe1\x34\x43\x85\x9d\x00\x1b\x4e\xb5\x32\x78\xbf\x03\xcb\xfc\x67\x2a\x66\x3c\x0a\x2b\xba\xdc\x67\x8b\x85\x6e\x10\xa8\x6c\x18\xda\x6a\x1b\xf6\xd3\x82\xe6\x6a\x1b\x1c\x2b\x44\xae\x23\x3b\xb2\xcf\xc2\x3e\x4d\x0e\x10\x97\x7f\xf7\xcd\xba\xb8\x06\x4a\x8c\x99\xc9\x46\xf3\xaf\x93\x53\x78\xa3\x89\x9e\x7b\xf6\x79\x2e\xd3\xb1\x6f\xb6\x3e\xf7\x72\x34\xaa\x38\x58\x96\xc8\xa7\x79\x53\x57\x5a\xc5\x4f\xc2\x10\x36\x03\xf4\xfe\xf2\x67\x9b\x12\xcb\xaa\x63\xd3\x86\x9c\x61\x5c\x74\x0d\xa0\x75\x42\x97\x14\x11\x0d\x0f\x13\x6a\xd3\xb5\xdf\xb4\xf7\xff\x92\x63\xdc\xc5\x39\x2c\x1e\xe1\x84\x3e\xfb\xd3\x9b\x3e\x03\x3e\x4c\x05\xb3\x95\x12\x07\x12\x54\x40\x6a\xeb\x05\xdc\x2f\xcc\xa6\x3b\x56\x12\x9b\x56\x12\xdb\x56\x12\xb0\x92\x58\x5e\xa9\xa9\x60\xa2\x63\xeb\x2d\x5f\x72\xb8\x46\x03\xf9\x4c\x18\x45\x06\xaa\xd8\xb5\x7e\x40\x60\x22\x82\x27\xee\xb5\xcc\xec\x7a\x9d\x19\xfd\xfd\xf4\xf3\x27\x34\x0a\xd3\x59\xb1\x94\xad\x1a\x84\x9e\x80\x3b\xb2\x76\x66\x49\x12\x05\xd8\xa2\xac\xbf\x9f\x64\xee\x03\x8a\xb9\x70\x37\x53\xdc\xb7\x3d\xb7\xcc\x93\x3a\x2c\xff\x6b\x4e\x7d\xfc\x68\x30\x68\x7f\xc4\xe9\x05\x0d\x08\xcd\xae\x02\x68\x40\x59\xe2\xbd\x52\xb4\x5e\xc3\xa1\x8a\x54\xf2\x56\x66\x8b\xe6\x72\x17\x0c\xb3\xdc\x47\x5e\x34\xee\x44\xd8\x16\xd0\x1e\x0c\x70\x03\x9e\xa5\xf0\xa4\xd8\xb9\xb2\xd9\xf2\x8d\xc8\xe5\x92\x59\xfb\x93\xb2\x3b\xca\xb6\x84\xbc\xf6\x83\xce\x52\x69\xde\xaf\x4b\xb5\x09\x51\x70\x7b\x10\x99\xba\x90\x81\xc7\xc7\x96\xd4\xb1\x7f\x55\x34\xe6\xc6\x8d\x61\xb8\x2a\xed\x6b\xc5\xc6\xf9\x9c\x17\xcc\x4f\xce\x24\xc3\x41\xf8\x41\xb2\x35\xc1\x39\xb9\x1b\x66\xb1\x36\x33\xfe\xa5\x72\x2a\xf1\x77\xa5\x77\x78\xa5\xb7\x8d\x39\xd0\xf4\x75\xed\xa8\x30\xba\x76\xd3\xa8\x19\x8c\x48\x4c\x73\x9d\x31\x76\xe8\x19\x72\x45\xb9\x97\xa5\xc0\xf8\x6d\x99\xe2\xb7\x0a\xb3\x1d\x53\x61\x82\x9d\xd8\x42\xde\x91\x79\x9f\x60\x97\xb9\x78\x54\xef\x72\x85\xd2\x96\x7d\x76\x7c\x4f\xbd\x9e\xe5\xc3\x2b\x48\x6b\x6e\x29\x18\x8f\xe5\x55\x8e\xa9\x4c\x56\x07\x0e\x3c\xcf\x97\x18\x49\xe0\xdc\x7e\x29\xf8\x14\xa3\x1f\xe4\x3c\xc3\x09\x5d\x4d\x8c\xf4\xa2\xe4\x06\x9d\x21\xe4\x3f\x4b\x3a\xff\x36\xcd\xac\x56\x13\x93\x7b\xbd\x3a\xe9\xa3\xb9\x1b\x07\xc7\x98\x9e\x3a\xa5\x79\x10\xf2\x75\x7a\x4b\xa8\x12\x31\x94\x64\x1c\xb7\x52\xbf\x9f\xf0\xc7\xcf\xfd\xe8\x52\x39\x24\x72\x93\xaa\x47\x65\xbc\x63\x78\xc4\x45\x31\xaf\x63\x3f\xb2\x70\x10\xc2\x6f\xec\xa0\x97\xbb\x3c\xbb\x87\x04\xb8\x14\x33\x95\x8a\x4f\x9b\x86\x83\x36\xfb\xf7\x4f\x4f\x8e\x4e\xcf\x68\x19\x53\x36\x64\x6d\x0a\xad\xfb\xcb\x48\x6e\xd7\x7b\xe8\xb1\xd9\xc6\x81\xfd\xfe\x52\x42\xd0\xfc\x05\x76\x2f\x8c\x4e\xb4\x91\x06\x04\x4d\x8a\x8f\xdf\xef\xea\x03\x04\xe2\x3a\x2d\xd1\xd7\xf6\x09\x71\x7b\x7c\x3a\x98\xa4\x79\x41\x1f\xfa\xea\x91\x07\x82\xa7\xea\xbd\xf9\xc3\x02\xaa\x7b\x2e\x15\x71\xfc\x36\xc3\x10\xb5\x74\xd6\x1e\xdc\x85\x5c\x16\xfe\x88\xfb\x88\xe3\xa8\xe3\xfa\x36\x71\xd4\xf2\x59\x86\x45\xdf\xc4\xb1\xa5\xbd\x4d\xb3\xc7\xb7\x54\xcc\x8c\x6a\xe6\xcf\x2b\x6e\x59\x8f\xd7\x4d\xf1\x2d\x2a\x5a\x36\x31\x54\x23\x65\x12\xe4\xe3\xcf\xa7\x67\x2b\x12\x84\xf7\x06\x7b\x2b\x95\x13\x9b\x0d\x30\x3c\x49\x09\x3a\x0d\x7f\x1f\xd2\xc1\x14\xbb\x91\xcd\xca\xf1\x1c\x52\x37\x09\xae\xa0\x4b\x34\xcd\x15\x61\x4f\x38\x14\x07\xbe\x0b\xd1\xbc\x3f\x3a\xb3\x57\xa8\x8d\x4c\x1a\xc0\x97\x45\x62\x6a\xd4\xf8\xfb\x99\x48\x04\x8f\x4b\x5d\x12\x31\x77\x43\x3d\x69\xd8\xf3\xd3\x77\x24\x8f\x65\x4b\xf1\x50\xb7\x62\xb1\xa9\x93\xb3\x93\x72\xf6\x7d\x48\xc5\xb0\xb5\xc9\x8b\xd9\x8f\x25\x5d\x82\x71\xd7\x10\x3d\xd1\x34\xf4\x7e\x77\xc9\x1c\x9f\xb7\x2d\xa5\x81\xdc\x8a\xc4\xdc\x1f\xa9\x45\x62\x6e\x9a\xfc\xbe\x52\x59\x23\x03\x7b\x09\xa6\x4b\x06\xee\x52\xda\x77\xe9\xac\x0e\x8f\x3e\x1c\x9d\x1d\xb5\xa4\xd0\x42\x19\x04\x61\xfe\xf2\xcd\xfe\x11\xdc\x7f\x03\x00\x00\xff\xff\x4e\x79\xd7\xa2\x86\x38\x00\x00")

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

	info := bindataFileInfo{name: "rest/gorilla/pat.tpl", size: 14470, mode: os.FileMode(436), modTime: time.Unix(1446441660, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _serviceUpperioTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x58\x4f\x6f\xdb\x3a\x12\x3f\x5b\x9f\x82\x6b\x60\x0b\x2a\x08\xe4\x1e\x8a\x1e\xb2\xf0\xa1\x1b\xb7\x8b\x00\xe9\x9f\x4d\xda\x77\x29\x8a\x42\x96\xc6\x36\x11\x8a\x54\x48\x2a\x6d\x10\xf8\xbb\xbf\x19\x92\x92\x15\x4b\x4e\x93\xbe\x5c\xde\x3b\xb4\x91\xc8\xe1\xfc\xfd\xcd\x6f\x68\xdd\xdd\xb1\x12\x56\x42\x01\x9b\xd6\x79\x71\x95\xaf\x61\xca\xb6\xdb\x64\x12\x5f\x18\xee\x67\x9f\xae\xd6\xb4\x86\x8f\xa0\x4a\x7a\x4a\x7a\xa7\x44\x55\x6b\xe3\x6c\x38\x35\x55\xe0\x66\x1b\xe7\xea\x29\x3e\xaf\x85\xdb\x34\xcb\xac\xd0\xd5\x6c\xad\x1b\x53\xce\x2c\x98\x1b\x51\xc0\x43\x7b\xb3\xa6\xae\xc1\x08\x8d\x32\x68\x43\xac\x58\x76\x56\x66\x67\xf6\xd2\x19\xa1\xd6\xc1\x44\xef\xa8\xcd\x9d\x36\x02\x35\x64\x4d\x23\x4a\xd2\x0b\xaa\xd0\x25\x8a\xce\x96\xb9\x85\xd7\xaf\x68\xc9\xfa\xb3\x36\x68\x8c\x01\x4c\xa6\x52\xaf\x69\xd3\x9b\xcb\x84\x9e\x95\xcb\xe9\x81\x00\x51\x61\xc8\x49\xb2\x6a\x54\xc1\x84\x12\x8e\xa7\xec\x2e\x99\xcc\x66\xad\x4c\x74\x9e\xd5\x46\xdf\x88\x12\x0c\xfb\x81\x3e\xd2\xdb\xcf\xdb\x64\x12\x37\xb3\x4f\x71\xd3\x66\x0b\x7f\xea\x1d\x6a\xe3\x53\xca\xef\xe7\xdb\x1a\xb2\x0f\x79\x05\x68\x65\x7a\xcc\xbc\x19\x6e\xd8\x11\x25\x32\xbb\x80\xeb\x06\xac\x4b\x19\xb7\xad\x9d\xec\x32\xfc\x3d\x66\x60\x0c\xfd\xd3\xc6\x3b\x34\x31\xe0\x1a\xa3\xd8\xff\xc0\xed\xab\x8d\x27\xb8\x49\x93\xc9\x36\x4d\x30\x18\xf4\xfe\xb0\x5c\x1b\x89\x65\x26\xff\xc1\x0e\x08\x85\x7c\x3c\x64\x6c\x24\x84\xa3\x03\xc2\x7b\xb1\xf8\xec\xea\xa5\xcb\x85\x62\x65\xee\x72\x2a\x67\x32\x29\x97\x41\xec\x64\xce\x22\x4e\xb2\x8f\x35\x28\x6e\x8e\xd9\x14\x4b\x91\x37\xd2\x4d\x31\x40\x84\x0d\x49\xfd\x6b\xce\x94\x90\xbd\xc4\x60\xe8\xc9\x58\xd9\x72\x2c\x7a\x2b\x61\xd9\x9c\xbd\x38\xe0\xe3\x5d\xb9\x44\xe8\x44\xc9\x90\xc2\x43\xf9\x23\xd5\x98\xbd\x35\x28\xf4\xb2\x60\xa7\x5f\x2e\x16\x6c\xa5\x0d\x73\x28\x3b\x38\x14\x6a\x81\x92\xb9\x83\x92\x2d\x6f\x99\x6f\x09\x76\x7a\x7e\xc6\x9c\xd6\x32\x19\x3d\xd4\x59\x72\xa6\x29\x1c\x45\xb9\x58\xb2\x72\x99\x2d\xda\x6c\x05\x07\x4f\x0d\xa0\x56\x96\x0f\xce\x23\x90\x99\xdb\x40\x97\xdd\x63\xa6\x57\x7e\xa1\xce\x0d\x28\x17\x8a\xfb\x40\xc1\xd2\xa8\x9a\x27\x93\x42\x63\x02\x5b\x6c\x9e\xe2\x8b\xc5\x32\xd5\xdd\xca\x5b\xe5\x84\xbb\xfd\xe4\xb0\xb0\x7c\x58\xe4\x35\x38\x56\x68\x29\xa1\x70\x42\x2b\x52\x26\x65\x57\x65\x8b\xea\xa4\xe4\x8f\xa9\x69\x5e\xd7\xf2\x16\xe1\xaa\x4a\x5d\x31\xa2\x03\x16\x1a\x1f\x53\xd8\x3e\x89\xf2\x20\xab\x90\x3c\xa1\x0a\xff\x66\x1f\xe0\xc7\x1f\xaf\xc8\x26\xd0\x12\xd4\x19\xc7\x1c\xec\x52\xb0\xdd\xd2\x56\x46\x69\x41\x2d\x6d\x36\xe7\xd1\x88\xcd\x3e\x1b\x51\x5d\x88\xf5\xc6\xf1\xc0\x41\xd9\x97\x8b\xf3\xb7\x91\x96\x32\xff\x00\x9f\x75\x30\xcd\xd1\xdc\xd7\x93\x6f\x29\xc2\x77\x4e\xc0\xed\x11\x10\x85\xf4\x3e\x37\x76\x93\x4b\x5f\x15\xe1\xa0\x3a\x26\xd7\x6b\x6d\xad\x58\x4a\xf0\x12\xfc\xba\x11\xc5\x15\x5b\x89\x9f\x1e\x5d\xb1\x29\xa8\x7b\x51\xa2\x0a\x34\x84\x98\x88\x8a\xc0\x84\x44\x56\x54\xec\xab\x36\xb8\x7b\xfb\xff\xa1\x0d\xca\x2e\xd4\xa1\x08\x73\x94\x6e\xf7\x17\xff\xa5\xac\x8c\x94\xa2\xab\x05\x16\xa3\xab\x47\x59\x7a\xc7\xc1\x57\x9f\xaa\xd0\xaf\xf2\x81\x32\x7c\x6f\x8d\x92\x6c\xf6\x06\xc3\x51\x25\x87\x3a\xa6\x46\x5a\xdf\x2d\x13\x51\x76\x00\x19\x95\x8b\x1c\x3f\xf4\x13\x49\x1f\x69\x58\x28\xb7\xe2\xd3\xb7\x84\x42\x56\x10\x86\xc9\xfa\x3e\xca\x4f\xd8\xbf\xed\xd4\x9b\xc9\xbc\x24\x4f\x29\xf4\xe0\x5c\x07\x6c\xda\x38\x53\x0e\x8c\xca\xe5\x1e\x22\x47\x03\xec\x07\xd1\x41\x96\x92\x74\x05\x3e\x43\xbb\x7c\x3d\x15\x7c\x71\x81\xe4\x70\x81\x23\x8a\x39\x86\xf9\xfa\x55\xba\x87\xaa\x7b\xe4\x75\x09\xb9\x29\x36\x63\xdc\x80\x1c\x24\x9c\x65\xd4\xd8\x82\x0a\xc6\x6d\xfa\x08\x42\x08\xfa\x90\x10\xae\xbb\x14\xfd\xbf\x01\x73\x9b\x76\xaf\x17\x60\x91\xa3\x7d\xeb\xc7\x59\xd5\xf2\x38\x36\x5d\xd8\xe4\x64\x07\x07\x2c\x37\xc8\x9f\x88\xcd\xb0\x3a\x18\x75\x23\xcc\x71\x88\x3a\x7e\x01\xd8\xa0\x0c\x17\x8c\x80\x9b\x98\x7f\x81\xb6\x89\x87\xc5\x0d\x28\x76\x4d\x31\xec\x72\x61\xbd\x25\x24\xb9\xfe\x18\xf2\xac\xc7\xaf\x33\x9c\x86\xe1\x31\x8d\x86\x83\xe4\xfc\x9e\x69\xdb\x02\xfc\x9d\x40\xd8\x92\xe0\x36\xe0\x62\x74\xdf\x6b\x48\x7b\xae\x6e\x90\xe3\x24\xf1\xf4\x1a\x51\x15\xac\x78\xc3\x1f\x57\x2b\x0b\x74\x37\xc1\x40\x5f\xf6\x75\xe1\xff\xd9\xe5\x95\xa8\x91\x6e\x94\xe3\xf7\x64\xd3\xa0\x78\xa7\xe4\x5c\x54\xe2\x90\x8e\xb0\xb7\x53\x12\x65\xd3\xce\xb9\xae\x01\xd2\x24\x42\xec\x23\x8e\xd9\xb0\x6a\x3d\xba\x57\xc2\x58\x37\xc4\x5b\x95\xbb\x62\x03\x4f\x06\x1c\x6a\xa7\xf1\xf3\xdb\xb3\xa7\x2b\xba\xf1\x20\xb3\x6c\x65\x70\x78\xec\xee\x1b\x92\x4a\xfc\xe2\xeb\xb7\x7d\x07\xee\xb6\x84\xf1\x93\x1d\x13\x20\x7a\x3d\xd2\x79\x8a\x77\xb3\x88\x80\x22\x8d\x97\x8d\xa6\xaa\x3b\x03\x98\x3b\xa4\x67\x2d\x88\x34\x68\xe0\x56\x79\xcd\x66\x71\x86\x27\x2d\xbd\x64\xb1\x91\xae\xd3\xec\x0d\x42\x58\x3e\x66\xfc\xa1\x80\xd2\x0e\xa7\x40\xa3\x90\x1d\x0d\xd0\xa5\xdc\x1f\x93\x78\x3f\x3a\x92\x29\x41\x30\x14\x74\x84\xc3\x3e\x68\xf7\x8e\x0e\x8e\x4c\x55\x1c\x36\xeb\x70\x51\xb8\xc9\x65\x03\xe4\x73\xe8\x0a\x1f\x84\x97\x89\xc4\x15\x4a\xdb\xa6\xb4\x0c\xf2\xc9\x84\x1f\x45\x12\xbb\x9f\xc2\x14\x3d\x62\xe4\xd8\xd7\x97\xdf\x3a\x32\xc0\xd8\x22\x6e\xbe\xd4\x25\x5d\x5b\x06\x40\xd1\xea\xa9\x18\x09\x9a\xfe\x12\x4c\x9e\xeb\x8a\x42\x7a\x96\x3d\x26\xf1\x17\x4f\x4c\xaf\x36\x10\xa8\x9f\xac\x86\xeb\xd4\x31\xfb\xee\xc7\x1b\xb5\xd9\xfb\xbc\x26\xf5\xd4\x87\x27\x7d\x62\x40\x72\xa4\x48\x3c\x41\xa4\xe9\x3f\xe0\xca\xd0\x84\xa2\x93\xef\x81\x10\xca\x8e\x8d\xdb\xde\x20\x1a\x8a\x15\xf5\xe3\xfe\x31\x33\xde\xab\x7d\xa6\x19\xbf\x7f\xff\x5f\x80\x84\xe7\xc1\x69\xd0\x34\xc4\xe9\xdf\x1d\x91\x06\x2a\x7d\xf3\xeb\xa2\x5e\x78\xb1\x71\xa7\x87\x25\x2d\x29\x59\xcf\x5d\xd2\x1e\xfd\x20\xed\xea\x22\x90\x02\xcb\xe9\x99\x70\x59\xa1\x8b\x78\x11\xa0\xae\xc9\x55\x7b\x4f\xfb\x75\x61\x7b\xba\x78\x3a\x64\x1c\x8a\x30\xda\x1f\xfc\xee\xc4\x39\x33\xf0\xe7\x5c\x20\xcb\x3e\xe8\x13\x93\x28\xf2\x34\xc7\x48\xe9\xc0\x39\x5a\xdc\x73\x70\x74\x14\x06\x17\xcf\x71\x26\x08\x65\x6b\x84\xa4\x2f\x36\xce\x9d\x35\x52\x09\xce\x8b\x27\x3b\x86\xaa\x78\x2d\xc7\xbd\x49\x99\xbf\xd7\x92\x53\xe0\xc7\x73\x2d\x71\xba\x0c\xfd\x4a\x3b\xa7\xbd\x3c\xf7\x63\x10\x64\xda\x7e\xfb\xa0\x16\x89\x37\x13\xef\x2e\x7d\xe7\x68\x3f\x06\xf5\x5b\xeb\x11\x3f\x83\x7d\xb3\x31\x4e\x87\x98\xef\x80\xf6\xf0\xfe\xa5\x35\x36\x1c\x99\x1a\x6f\x5e\xea\xdd\x45\x5f\x03\x7d\x1f\xf2\xaf\xdb\xed\xf8\x87\x8d\x91\x5f\x34\x5a\x29\x3a\x8c\xcd\xb1\x33\xc2\x3a\x35\xa1\x41\x88\x80\x7f\xaf\x47\xda\xf4\x49\x6d\xe1\xde\xf7\x03\x3c\x8a\x53\x46\x53\x36\xf3\x91\x1b\x9e\xb0\xac\xb1\x74\x5f\x7d\x44\x42\x49\x37\x66\xd4\x27\xae\x87\xbe\x90\x9b\xb0\x49\x6e\xf4\xbe\xd7\xfd\x19\x00\x00\xff\xff\x5c\x8b\x37\xa3\xc6\x14\x00\x00")

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

	info := bindataFileInfo{name: "service/upperio.tpl", size: 5318, mode: os.FileMode(436), modTime: time.Unix(1445795740, 0)}
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
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"general.tpl": &bintree{generalTpl, map[string]*bintree{}},
	"rest": &bintree{nil, map[string]*bintree{
		"gorilla": &bintree{nil, map[string]*bintree{
			"pat.tpl": &bintree{restGorillaPatTpl, map[string]*bintree{}},
		}},
	}},
	"service": &bintree{nil, map[string]*bintree{
		"upperio.tpl": &bintree{serviceUpperioTpl, map[string]*bintree{}},
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

