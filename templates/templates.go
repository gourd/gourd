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

var _restGorillaPatTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5b\x5f\x6f\xdb\x38\x12\x7f\xb6\x3f\x05\xa3\x43\x17\xf2\xad\x57\xde\x03\x0e\xf7\x90\x43\x1f\xb6\x89\xdb\xdd\xbb\x6e\x92\x8b\xd3\xbb\x87\xa0\x38\x28\x12\x6d\xeb\x22\x8b\x0a\x45\xc5\x09\x82\x7c\xf7\x9b\x3f\xa4\xfe\x45\x76\xd2\x20\x69\xb3\x45\x16\xd8\xd6\x12\xc9\xe1\xcc\x8f\x33\xbf\x19\x92\xea\xcd\x8d\x88\xe5\x3c\xc9\xa4\xf0\xf2\x30\x3a\x0f\x17\xd2\x13\xb7\xb7\xc3\x81\x7d\x10\xd0\x1e\x1c\x9d\x2f\xf0\x1d\xfc\x94\x59\x8c\xbf\x86\x8d\x51\xc9\x2a\x57\xda\x14\x3c\xca\x5b\x24\x66\x59\x9e\x05\x91\x5a\x4d\x16\xea\xa7\xf3\xc4\x4c\xf0\x7f\x18\x96\xab\x24\x33\xde\x70\xb0\x34\x26\x37\x3a\xcc\x0a\x1c\x25\x36\xf4\xaf\x3a\x4c\xb0\xbb\xd7\x95\xab\x93\x34\x0d\x27\x79\x88\xf2\x16\xaa\xd4\x71\x64\xae\x3a\xa2\xe0\xe5\x24\x52\x99\x91\x57\xd8\x69\x32\xe9\x6b\x8d\x65\x74\x47\x34\xb6\xe4\x52\xaf\x7a\x1b\x0a\xa9\x2f\x93\x48\x52\x9b\x4a\xc3\x6c\x11\x28\xbd\x98\x5c\x4d\x32\x69\xea\xc9\xa0\x51\x66\x20\x3c\xc9\x16\x93\xff\x15\x2a\xc3\xde\xf3\x15\x6a\xe1\xa5\x6a\x81\x7f\x61\x77\x67\x57\x61\x34\x8c\xbc\xb4\x3f\x61\x4c\xe1\x6d\xc0\x19\xf5\x25\x90\x87\x93\x09\xad\xca\xc9\x75\x2e\x83\x83\x70\x25\xe1\xe5\xb1\x2c\x8c\x38\x4b\xb2\xb8\x10\x56\x49\x61\x94\x00\x84\x84\x56\xa5\x91\x7a\x38\x2f\xb3\xa8\x77\x94\xaf\xc5\x9f\xa1\x5f\x70\x4c\xfd\xc6\xe2\x2c\x2c\xe4\x58\x64\xaa\xcc\xf8\xcf\x5c\xb0\x5e\x23\x71\x33\x44\x20\xc5\x65\xa8\x93\xf0\x2c\x95\x05\xce\x50\x16\x52\xa4\x21\x4e\x30\x58\x48\x33\xb3\x53\xef\xbe\x15\x1f\xa4\xe9\xce\x36\x1c\x58\xd5\xe8\x19\xfa\x78\xdd\x1e\x80\xc2\x19\xc0\x7a\xde\x90\xf3\x43\xb7\xcf\x0d\xc8\x09\xd3\x54\x45\xd3\xcc\x24\xe6\x1a\xfb\x34\xc7\x04\xbf\xd4\x6d\xad\x8e\x1f\x13\x40\x68\x4b\x67\x6c\x67\x03\x65\x36\x57\x1a\x66\x97\x3c\x41\xae\x15\x38\x04\xfc\x38\x93\xf0\x5e\x8a\x48\x4b\x30\x18\x02\x44\xcb\x3c\xd4\x72\x8f\x1e\x51\x32\x42\x8c\x68\xe2\xd2\x06\xc7\xf2\xa2\x04\x74\xc7\x42\xba\x05\x09\x78\x9e\x23\xa3\x47\xc2\x97\x5a\x0b\xf8\x5f\x69\x84\x75\x80\xb3\xe6\x69\x18\xc9\xa5\x4a\x63\xa9\x77\x45\x98\x5d\x9b\x25\x80\x2e\xae\x55\x29\xd6\x61\x66\x10\xeb\x58\x89\x35\x78\xa4\x30\x4b\xa7\x1b\x8f\x6c\xfd\x67\x95\x0c\xf3\x1c\x5d\x08\x47\x85\x26\xc4\x35\x85\xbe\x5a\x9a\x52\x67\xc3\xc1\xed\x83\xec\x2c\xf3\xb8\x69\xe7\x27\x7a\xac\xec\xec\x31\x0b\x6c\x4d\x3b\x6f\x11\xd4\xaf\x63\x30\x6b\xbb\xdd\x60\x6b\x89\x33\x38\xc5\x15\x77\xe6\xb5\x9d\x64\xc3\x52\xbe\x6c\xf3\x96\x61\x16\xa7\x52\x20\x7b\x25\x45\x91\xa8\x4c\x68\x89\xb1\x19\x0b\x78\x2c\x71\x25\xa1\x05\x3d\x7e\xbd\xd9\xc4\xc6\x60\x0e\xfb\xb1\x48\xc0\x4d\x44\x10\x04\x40\xdf\x52\xcf\xc1\xa8\x9b\xdb\x11\xdb\x4a\xa6\xae\x50\x18\x0e\x0b\x20\xe4\x7f\x2f\xaf\x7c\x3d\xaa\x54\x13\x2b\x8a\xb0\xb5\xaf\x9b\x92\x59\x24\x48\x1c\x55\xaa\x03\x13\x03\xbb\xfd\xb6\x0f\x7a\x88\x85\xcc\xa4\x4e\x22\x08\xdd\x6b\xfb\x3e\x16\xbb\x49\x2c\xe6\x89\x4c\x63\xea\xed\xaf\x95\x3e\x2f\x18\x3c\xa4\x38\x84\x23\x16\x9f\x8e\x3f\x12\xd9\x21\x57\x0d\x07\xc0\x52\x2d\xa9\xad\xb4\x13\xec\x53\x8b\xb5\xfa\x3d\x52\x63\x3f\x20\xb0\xb2\x9a\x7f\x89\x86\xf9\x63\xd1\x59\x6e\xd0\x0e\x40\xd0\x01\xa8\x10\xfc\xab\x94\xfa\xda\x1f\x21\x1a\xbe\x07\x7a\x7b\x23\x01\x2a\xaf\x21\x63\x89\x08\x56\x68\x81\x0b\x07\x8c\x4f\x23\x9c\x37\x1d\xc8\xf5\x1e\xbc\x2a\x60\xd8\x2f\x71\xec\x7b\x30\x0a\x40\x8a\x19\x48\x9e\xbe\xd5\xd9\xcd\x31\x93\x86\xc7\xa1\xc0\xd1\x5d\x8f\x60\x00\xd0\x4b\x11\x02\x7e\x12\x17\x38\x58\x80\x57\x51\x00\x08\x97\x9a\x9b\x90\xb9\x11\xcf\x08\x1a\x28\x7b\xd1\x81\xc0\x5a\x35\xb4\xf1\x13\x6a\xc8\x2d\x05\x16\x09\xf0\x13\xb8\x9f\x92\xcc\x00\x5f\xcc\x8c\x66\xbc\xdf\x2b\xbd\xfa\x77\x98\x96\xd2\xf7\x0a\xaa\x41\x10\x83\x64\x2e\x5c\xa7\x1d\xc8\x31\x1e\xad\x10\x8d\x2b\x68\x42\xce\xb2\xc1\x2c\x4f\x13\xe3\xdb\x9e\x63\xe1\x8d\x69\xf0\x00\x71\xf9\xef\x98\xe7\xc5\x39\x70\xc5\x04\x0f\x26\x39\x83\x8b\x60\x06\x4f\x34\x90\x06\x40\x3a\x22\xb8\x6b\x9d\xf3\x70\x81\xa1\xee\x30\x68\x6a\xaf\xe6\xf3\x42\x42\x98\xa5\xc9\x2a\xd9\x48\x34\x80\x9f\x82\x2e\xa2\x04\xf4\xfe\xf6\x57\xf6\xb0\x81\x2a\x7a\x8c\x66\x71\xac\x78\xda\xd7\x81\xe6\xe1\x76\x80\x85\x64\x34\x30\xa1\x77\x86\x17\x86\x91\xc1\x52\x24\x38\x42\x2b\x3e\xc1\xe4\x3e\x0e\x18\x8b\xbf\xfc\x3c\x16\xa0\xc7\xdf\xa9\xdf\xdb\xb7\x22\x4b\x52\x3b\x7e\xa0\xc0\x05\x94\xa1\xdf\xb7\x0e\x0c\x94\x9a\xf6\xcc\x94\x6e\x9b\x29\xbd\x6f\xa6\x14\x66\x4a\xbb\x33\x39\x77\x1f\xdc\x22\xeb\xf0\x1a\xc0\x3b\x9d\xc8\x4b\x0c\xb3\x0b\x0c\x90\x43\xc2\xc8\x67\xa8\x46\xee\xed\x47\x04\xc6\x27\x78\x68\x60\x1d\x66\x17\x9b\xc2\xe8\x1f\xb3\xc3\x03\x0c\x0a\x6e\x2c\x44\x28\xee\x06\x84\x59\x02\x1d\xd9\x38\xb3\x22\x49\x02\x98\xa8\x04\x15\x23\xc4\xc7\x2b\xb9\x52\x10\x83\x60\x73\x19\x81\x30\xd9\x8c\x3d\x37\xcd\xb3\x12\x16\xa8\xe4\xb4\xa9\x53\x4f\x8d\x41\xa3\x6c\xaa\xe2\xb1\x6b\xd5\x00\x5e\xe0\x5a\x62\x89\x8b\xd1\xcb\x1a\x6a\x5f\x07\xef\x54\x7c\x8d\x40\xd3\x22\xe2\x38\xab\xbd\xd3\x6c\x23\x51\xd9\x12\xe3\x8b\x01\xae\xc1\xb3\x12\x9e\x15\x3b\xbd\x0a\x73\x34\x7c\x15\x9e\x4b\x1f\x7e\x9f\x32\xa3\x7c\x6e\xe6\xc7\xa1\xed\x77\xea\x31\xba\xde\x67\x16\xf3\xb6\xbd\xc8\x9c\x2b\x21\x36\xb0\x6d\xa7\xe1\xef\xb5\x63\xd7\x92\x88\xb7\xbb\x82\x28\xb9\x3d\x54\x4c\xb5\xbe\x28\xb0\xaf\x1e\xac\x55\x2b\x72\xfb\x58\x88\x45\x72\x29\xb1\x94\x00\x30\x33\x60\x37\x72\x65\xec\x34\x1c\x74\x06\x38\x3a\x5b\x0b\x8b\x2a\x8f\xf8\x8f\x4e\x68\x63\xd1\x94\x50\x17\x12\xdd\xca\x09\x64\x36\xbd\x6a\x9a\xb1\x57\xad\x6b\x87\x82\x1e\x01\xbf\xf6\x9d\xc8\x1e\x8f\x42\xad\xa6\x28\xd6\x4a\x10\xf2\x2a\x0f\x71\x7b\xc4\xe6\x70\xfd\x02\x96\xd8\x9d\xc0\x94\x9f\x97\xd8\x44\x03\x44\x65\xe5\x1d\x49\xf7\xd8\xd9\x4d\x73\xa0\xcc\x45\x99\x44\xe7\x50\xc0\x5c\x51\xda\x5d\x28\xd8\xe2\x42\xd1\x12\x57\x29\x62\xad\xa1\x5e\xc7\x9c\x01\x35\x38\xec\xac\x56\x98\xe7\xa0\xba\x89\x96\x28\x28\xf0\x0d\x6c\x7f\x18\x9d\x08\x2a\x9d\x8e\x77\xbf\x0b\x63\xeb\xbf\xa4\xe4\x2e\xae\xbb\x05\x0a\xc7\x6e\xed\x3c\x0a\xe0\x2f\xe7\x1e\x5f\x84\x7c\x55\x07\x13\xaa\x24\x0c\x57\x72\x54\x95\x74\xb4\x08\x3f\xe1\x7f\xf4\x64\xa3\x9c\x76\xb2\x50\xb2\x2d\x8b\xe1\x00\x37\x97\x47\xf0\x93\x36\x65\x68\xd7\x8f\xc2\x9b\x78\xf0\x27\x36\xd0\xc3\x4d\x12\xe3\x96\x30\x4f\x4b\x1d\xa6\x07\xdb\xba\x83\x37\xc3\xde\x3a\x38\x82\x40\x34\x73\xdf\x3b\x9e\xce\x4e\x68\x9a\x5d\xf1\xa6\x80\x52\xaa\x2d\x61\xc4\xda\x14\x46\xd1\x5e\x80\xcb\x9f\x42\x2c\x25\x12\x31\x34\x9c\x1c\xee\x1f\xee\x42\x78\x5f\x43\x45\x7d\x29\x45\x26\xd7\x96\xa9\xd1\x27\x68\x14\xba\xbe\x1b\xd6\xc3\x04\xae\x11\xc0\xe2\x1f\x38\x63\x35\xe2\xd4\xe3\x2d\xa3\xf7\xd9\xf1\x10\x9e\x5e\xd8\xe3\x83\x60\x8f\xff\xc6\x7d\x63\x3b\x4e\xc0\xd7\xb7\x97\x53\xa0\x39\x6c\xc1\x9d\x20\xaa\xb1\x21\xce\x0d\x14\xdc\x18\x1d\x63\xa1\xce\x51\x57\x77\x5c\x12\xfc\x7a\x72\x72\x64\x7d\x01\x15\xb0\x0c\xb2\x03\xbd\xb8\x5e\x72\xa9\xda\x2d\x34\xce\xf4\x1b\xce\x9f\x85\xa9\xeb\x10\x60\xf4\x48\xfd\x7b\xb1\x00\x53\x3c\x2a\xef\x1b\x85\x4f\x92\x89\xfa\x04\xc6\xf9\x10\x0e\xbb\x4b\x4d\x8d\xdd\x28\x6f\xaf\xc9\xd4\x85\xe6\xc4\xd4\xda\x69\xe3\x5e\x42\x8e\x1a\x06\x5b\x0d\x31\x66\xaa\x02\xa3\x3e\x8b\xd8\xc8\x8d\x5f\x6e\xe0\x7c\x65\xa0\x6a\xb4\x2e\xc6\x04\xa2\xce\x4c\x98\x64\x68\xf4\x9b\xfa\xe0\xc5\x7f\x53\x8c\xc0\xe9\x1a\x87\x1d\xa4\xd7\x68\x2b\x08\x98\x4f\xe7\x40\x2e\x45\xb0\x97\xaa\x42\x56\x69\x37\x5a\x4a\xa0\x8e\x7a\xf3\x54\x1b\xc3\xbb\xae\x6a\x7f\xe5\x59\xe0\xbc\x1f\xf9\xec\x46\xda\x32\xaa\x69\x73\x35\x7b\x6d\x31\x14\x8b\x67\x49\x1c\xcb\xac\x77\x51\xac\xcc\xaa\x44\x68\x06\xd9\xe9\xce\xce\xce\x67\xb7\x97\x86\xc8\xe0\xbe\x10\x71\x7f\xba\xf4\x68\x89\xaa\xf9\x02\xbb\x72\xa0\x87\x6d\x78\xf2\xf5\xa0\xc2\xd0\x2e\x0a\x29\xc2\x6b\x42\xda\x8c\xad\x92\x4e\x35\x8b\x0f\x48\x63\xde\x1a\x59\xa5\xee\xf7\x50\x5b\x76\x30\xdf\x53\xce\x29\x04\x86\x7f\x5f\x0d\x40\x46\x79\x85\x09\x4d\x59\x78\xbb\x9c\x2a\x66\xf4\x74\xf8\xcf\x31\xb5\xd1\x81\xde\x2e\xec\xf1\xef\xb6\x11\xa7\x61\x93\x38\x6d\xc9\x94\xb7\xe3\x3a\x9b\x93\x7a\xcc\xb7\x0d\x72\x39\xb6\xf5\xef\x3d\xf4\xd2\x53\xe9\x3c\x84\x64\x68\xcf\x66\xc7\x06\x55\x16\xa0\x9d\xdb\xe8\x3b\x20\xa1\xef\x9f\x4f\xaa\xb2\xdf\xee\x41\xb0\x20\xa9\x4e\x74\x31\x64\x53\x34\xa2\x73\x6e\xea\xf7\x6c\xae\x5c\x6c\xcf\x64\xa8\xa3\xa5\x7f\x31\xc2\xa3\x1e\x5f\xa6\xcf\x09\x4f\x41\x73\x55\xa1\x0d\xa9\x9d\xe2\xbe\x0a\xe8\xc7\x45\x31\x1e\x18\x04\x1f\x65\x86\xca\xe3\xae\xf3\xe7\x8d\x64\x79\xa0\xcc\x7b\x98\x2b\x6e\x0b\x16\x32\xc5\xaa\xb6\x9f\x95\x53\x05\x75\x5e\xc5\xc9\x69\x0f\x29\xf7\x22\xd3\xa2\xe5\x3b\xd0\x34\xb8\x6b\x9b\xc5\x5d\xc5\xba\x27\x9d\x94\x4a\x1f\xac\xd2\xa3\xb8\xd8\xdd\x84\xa0\x3f\xe3\x71\xd3\xa3\x56\xed\x2b\xd0\xac\x4c\xc7\x36\x6c\x36\xf1\x2a\x02\xf6\xca\xa9\xaf\x9c\xfa\x08\x4e\xc5\xf3\x09\xe0\xd5\x86\x13\x88\x0d\x14\xfb\x5d\x91\x6a\x2f\x1d\xe2\x91\xf3\x2b\x1d\x3e\x9a\x0e\x2b\x1b\x7a\x33\x16\x1e\x92\xfe\x40\x15\xe3\xcd\xed\x8b\x22\x4f\x3e\x97\x7b\x1e\xfa\x7c\xd2\x00\x74\x87\x7b\x15\x1d\x6f\x3a\xe0\x73\xb4\xdd\x3c\x9a\xbb\x43\xde\xb0\x24\x75\x27\x77\x12\x38\xac\xef\x7f\x2e\xf0\x96\xc8\xde\xfc\xbc\x32\xfd\x1f\x80\xe9\xe7\x09\x5e\xaa\x2f\x25\xa3\x92\x19\xa1\xe6\xf4\x98\xc4\x7f\x68\xf2\xe6\x93\x17\x7b\xc7\xdc\x3c\x79\x69\x31\x25\x07\xb1\x2f\x89\x25\x9f\xc3\xa6\x26\x53\x1e\xd1\x9c\x68\x18\x5d\xf3\xe2\xe3\xfe\x59\x65\xa4\xf8\x22\x03\x0d\x5d\xc3\x3d\xe0\x08\xc5\x22\xf0\x8d\xf2\x93\xd5\xd6\x2a\x51\x9d\xb9\x54\xba\x16\x81\x5d\x01\xa4\x8f\xfe\x03\x9e\x17\x94\xa9\xbe\x52\xee\xe9\x3f\x1f\x69\x5f\x07\x34\x12\xd1\xbe\x4c\xe5\x37\x4c\x44\xb2\xb3\xb9\x26\x87\xd8\xbe\xe5\x76\x3c\xc6\xd7\xf6\xc0\x37\xc4\x35\x5b\xf7\x0c\xaf\x09\xe6\x35\xc1\xbc\xa4\x04\xf3\x60\xfe\x8d\x29\x3a\xbf\x31\xff\x5a\x25\xfa\xf8\x97\xd9\x83\xbf\xb6\x79\x25\xdf\x7b\xc9\x97\xd8\xb7\x79\xf3\xb7\x82\x35\x4a\xe5\x1a\xbf\xc1\x8b\x96\x10\x49\xc4\x8f\x40\x7c\xf5\x2d\xdc\x70\xb0\xc2\x66\xba\x56\xab\xae\xd1\xf6\xb0\x2f\x02\x54\x11\xd3\x5e\x0a\x6e\xf9\x81\xbf\x03\xb6\xd7\x79\x7b\xee\xb3\x4c\x55\x98\x5f\xe9\x4b\x38\x5a\x8c\xf6\xfd\xe7\x81\x5c\x33\xfe\x2d\x69\x78\xdb\xb9\xca\x91\x8d\xd1\x00\x56\xc0\xef\xb9\xaf\xa3\xe6\xd6\x05\x3e\xbe\x68\xdf\x82\xe3\x9b\xf6\x9c\x3c\xe1\x3b\xfa\x8e\xcf\xaf\xe6\xfc\x54\x48\xa7\xff\xa6\x21\xd3\xc6\xa5\xb3\xdf\xbd\x85\xc6\x9b\x56\xcd\x5f\x8b\x1d\x1d\xce\x4e\xee\x5c\x72\xc2\x73\x0d\x84\x85\xc8\x5d\x0d\x08\xa4\xe0\x54\xd2\x27\xbb\x4f\x0c\x55\x7d\xfb\xd0\x00\x8b\x3e\x52\x78\x11\x50\x7d\x98\x9e\xd8\xe2\x91\x31\xaa\x01\xe8\x42\xc4\x1f\x87\xe2\x9f\x4f\x8c\x10\x9f\x23\x36\xd0\xb1\x1f\xdb\xbd\x20\x7c\xba\x9e\xd4\x40\xc1\xc2\x64\xbf\xb3\x71\x7e\x94\x97\x4f\x8d\x92\x3b\x30\x68\xe0\xc4\xaf\x5e\x04\x4c\x47\x9f\xda\x6e\x54\xdb\x6f\xf1\xe1\x44\x51\xe1\xc3\x29\xe5\x89\x21\x72\xa5\xec\x8b\x0c\xb4\xfd\xe9\xc7\xe9\xc9\xb4\x05\x52\x0b\x04\xc0\x89\xff\x75\x83\xfd\x87\x0e\xff\x0f\x00\x00\xff\xff\x9f\xc9\x1c\x3b\x86\x32\x00\x00")

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

	info := bindataFileInfo{name: "rest/gorilla/pat.tpl", size: 12934, mode: os.FileMode(436), modTime: time.Unix(1446145778, 0)}
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

