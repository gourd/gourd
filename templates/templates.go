// Code generated by go-bindata.
// sources:
// data/endpoints/endpoints.tpl
// data/general.tpl
// data/rest/gorilla/pat.tpl
// data/store/upperio.tpl
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

var _endpointsEndpointsTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x58\x4f\x4f\xec\x36\x10\x3f\x67\x3f\x85\x5f\x2a\x2a\x07\x6d\xb3\x3d\x6f\xb5\x87\x0a\xe8\x6b\x25\x0a\xf4\xc1\x3b\x21\x0e\x21\x99\x04\x8b\xc4\x0e\xb6\xc3\x1f\xad\xf6\xbb\x77\xc6\x76\x76\xc3\x42\xa1\x0f\xba\xd2\xa3\xda\x03\x90\xcc\x8c\xc7\xbf\xf9\xf3\x9b\xd8\xcc\xe7\xac\x80\x52\x48\x60\x71\x9b\xe5\xd7\x59\x05\x31\x5b\x2c\x46\x51\x78\x61\xa8\x4f\x4f\xae\x2b\x92\xe1\x23\xc8\x82\x9e\x46\x83\x55\xa2\x69\x95\xb6\xc6\xaf\x8a\x2b\x61\xaf\xba\xcb\x34\x57\xcd\xa4\x52\x3f\x5d\x0b\x3b\xa1\x1f\x5c\xd6\x2a\x21\x6d\x3c\x8a\x2a\xd5\xe9\x22\xb7\xf7\xec\xb1\x29\x0a\x9d\x65\xae\xa4\x85\x7b\x32\x7c\x5e\x6f\xac\xd2\xe0\xb4\xaa\xce\x64\x95\x2a\x5d\x4d\xee\x27\x12\x06\x2b\x51\x59\x36\xce\x43\xad\xaa\xf8\x1f\x50\xe7\xaa\xf0\x81\x8e\x26\x13\x17\xe3\x29\x39\x4e\x8f\xb2\x06\x50\x7a\x10\xf0\x1a\xa6\xc1\x76\x5a\xb2\xbd\xaf\x5f\xf6\x19\x2c\xa5\xa5\xd2\x4f\x17\x8d\xca\x4e\xe6\x2f\xf8\xe2\x52\x75\x72\xcc\xe8\x77\xcb\x8c\xd5\x42\x56\x09\xe3\x2b\xa7\x4d\xd6\x9e\x7b\xf1\x45\x2f\x4c\xfb\xc5\x09\x9b\x63\x5c\x08\xf5\x36\xd3\x22\xbb\xac\xc1\x30\xab\x58\x67\x80\xd5\x99\x05\x3d\x8a\xb2\xba\x56\xf9\x81\xb4\xc2\x3e\x1c\x0a\x63\xd9\x74\xc6\x08\x0e\x4f\xd8\xee\xf9\x05\x61\x3a\x7b\x68\x09\x0d\x9b\xf7\x21\xfd\xf8\x48\x3e\x5f\x30\x2c\x5f\x05\xd6\x61\xa7\xe5\x9f\xc1\x3e\x0d\x31\x72\xf9\x77\x6f\x68\x12\x3f\x31\x88\x3d\x4a\x67\x35\xc8\xd7\x15\x68\x70\x8a\xb3\xe3\xfd\xe3\x29\x46\xfa\xc0\xae\xb2\x5b\x60\x12\xee\x28\x13\x5d\x6e\x29\x1c\xb7\x6a\x14\xad\x96\xcd\xd0\xf2\x1a\xf8\x8b\x89\x19\x0d\x16\x9c\xc7\xb9\x06\xcc\x47\x7c\xc1\x42\xf8\xd4\x68\xa1\x31\xd2\x3d\xff\x77\xcc\x80\xa1\x31\xe8\x32\xcb\x61\xbe\xc0\x12\x68\xcc\xe6\x40\x82\x06\x5a\xd3\x8f\xd2\x3e\xed\x04\x1c\x33\xd3\x3b\x42\x5b\xac\x7f\x93\x59\xa1\x24\x2a\xf5\x98\xa9\x6b\xca\x46\xdf\xd9\xe9\xef\x67\x67\x27\x5f\xe0\xa6\x03\x63\x09\x40\x82\x46\xa2\x64\x9f\xd0\x6a\x8e\x8f\x91\x21\xf7\x68\xef\xc2\x4d\x0f\x68\x9f\x3f\x68\x77\x99\xd5\xbd\x3a\x3d\x05\x7d\x0b\xfa\x4f\x53\x61\x20\x71\x23\x8c\xc1\xe0\xb1\x70\xce\x27\xee\xcf\x56\x34\x89\x22\x72\x87\xde\xf0\x0f\xbd\xf9\xea\xe2\xd3\x62\x80\x3c\x64\x36\x32\x3e\x36\x02\x1b\x2a\xcd\x75\x80\x47\xf2\x4f\x33\x26\x45\xfd\x36\x94\x48\xb9\xf4\xb4\xc5\x22\xd9\x92\xc7\x2e\x79\x4c\x5d\xda\x4c\x48\x42\xbe\x63\x42\x4b\xf0\x1d\x93\xc4\x63\xb6\xec\x22\x87\x27\x79\x31\x8a\x28\x42\xce\x82\x66\x26\xdd\xab\x95\x01\x9e\x84\xb8\x7c\xa9\xb1\xc9\xa8\xe7\x51\x84\x64\x4f\x4f\xc2\xfe\x5e\x37\x65\x3b\x3f\xdc\xe2\x6e\x40\x1b\x04\xff\xe9\x9e\x53\x71\x8c\x33\x28\xfe\xe3\xd8\xc9\x20\x0a\x09\x70\x30\x7c\xfc\x0e\xcb\x38\xa0\xed\x81\xf9\x81\x40\xde\xdc\x1e\x3c\x09\x90\x5e\x2b\x29\x48\x1a\x5f\xd8\x0f\xa6\x55\xd2\x50\x61\xa9\x85\x67\xc3\x09\x32\xe8\x67\x17\x92\x9b\x3a\xd3\x75\xd2\xef\x42\xca\x77\x07\x82\x64\x31\x0e\x39\x0f\x7b\x8e\xdc\x9e\x03\x82\xa1\x5c\x0b\xb8\x7d\x8d\x62\xab\x56\xfd\x26\xa2\xdd\x50\xd6\xc3\xda\x94\xfb\xf4\xff\xd5\x81\x7e\x48\xb6\x34\xfc\x4e\x69\xe8\x3e\x3c\x44\xc4\x06\x1a\xa5\x1f\xdc\xa7\x71\xf9\x91\x22\xde\xd5\x04\x7f\xed\xf3\xb4\x5c\xdd\xb7\xd3\x80\xa0\xa7\x90\xe9\xfc\x8a\xdf\x24\xe9\xaf\x75\xcd\xa1\xde\x5c\x62\x8c\xdb\x69\xc9\x4f\x13\x8f\x1d\x79\x97\xac\x7c\x1b\x15\x11\xab\x49\x0f\x41\x12\x74\x36\x9b\xb1\x9f\x3d\xde\xe0\x67\x85\xf6\x48\xd9\xdf\x70\xa7\xe2\xa9\xdb\x7f\x4b\x66\xa8\x5f\x63\x6b\x8d\xb9\xde\x32\x75\xcb\xd4\xf7\x32\x15\xdb\xa1\xab\xf1\x40\x36\x7d\x44\x50\x54\xe4\xd8\x89\x76\x99\xa3\x60\x86\xcd\x85\x52\xbe\xb9\x3c\xb9\x5d\xdf\xc1\x5b\xaf\xeb\xd1\x7e\x84\x31\xe3\xcf\xcd\x12\xa0\xa0\x83\x72\x29\xee\x99\xc2\x2d\xcb\x5a\xdd\xb1\xd0\x3a\x25\x6b\xf1\xba\xf6\xb8\xa2\xdf\x36\x48\x22\xbc\x00\x56\x68\x15\x4f\x43\xb4\x47\x70\x77\x42\x3e\x79\x92\x3a\xc4\xa7\x60\xcf\x94\xcd\x6a\x8e\x6e\xb8\xab\x41\xb2\xd2\x1c\x8a\x46\x58\xa7\xb9\x49\x3f\xf7\xaf\x09\x9e\x68\x96\xa2\xe3\xb2\x34\x40\xb2\x64\x6d\x6c\xad\x4f\xad\xae\x2d\x5e\x3f\xc4\xbf\x71\x6e\xe1\x18\x6a\x1f\x8d\xae\xe7\xb3\x93\x2c\x47\x1c\xe9\xe3\x1b\x9a\x6c\xf1\xc5\xda\xa0\xc3\xaa\xad\x4c\xfc\xc9\x2e\xbe\x70\xa4\xc0\xfb\x26\x2a\x5c\xd4\x08\xb8\x30\x7c\x3b\x15\xbf\xd7\xa9\xe8\x9b\x6d\x75\x8d\x08\x31\xd0\xa0\xfb\xea\x54\x9c\xca\x49\x07\xf3\x5f\x36\x76\x4f\x70\x87\x88\x10\x20\x7d\xb2\xdf\x34\x25\x36\x75\x05\x58\x67\x67\x01\x35\x6c\x8a\x9d\xef\x3e\x4f\xfa\xf6\x70\x74\xa5\x91\x28\x8a\x97\x8f\x2a\x5b\xae\x7e\x28\xae\x96\x02\x8b\x65\xaf\xc0\xe7\x42\x5a\x2a\x31\xbd\xba\x32\x7f\xcc\xfb\x83\x67\xd3\xb3\xe3\x67\xdf\xa9\xdc\xf8\xf9\xbf\xcc\x9e\x27\x37\x16\xb7\x3c\x3c\xfb\x7f\xd0\x86\xff\xd5\xfe\x1d\x00\x00\xff\xff\xc3\x3f\xb8\xc5\x97\x16\x00\x00")

func endpointsEndpointsTplBytes() ([]byte, error) {
	return bindataRead(
		_endpointsEndpointsTpl,
		"endpoints/endpoints.tpl",
	)
}

func endpointsEndpointsTpl() (*asset, error) {
	bytes, err := endpointsEndpointsTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "endpoints/endpoints.tpl", size: 5783, mode: os.FileMode(436), modTime: time.Unix(1447992097, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
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

	info := bindataFileInfo{name: "general.tpl", size: 295, mode: os.FileMode(436), modTime: time.Unix(1446633618, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _restGorillaPatTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x3a\xeb\x6f\xdb\xc8\xf1\x9f\xa5\xbf\x62\x8f\xc0\x1d\x96\xf9\x31\x74\x7e\x40\xd1\x0f\x3e\xb8\x40\xe2\xf8\x72\x77\xc8\xc3\xf5\xa3\xfd\x10\x18\x05\x4d\xae\x24\xd6\x14\x97\x5e\x52\xb2\x0d\x9f\xff\xf7\xce\x63\x97\x5c\x8a\x94\x6d\xa5\x31\xd0\x36\x80\x23\x72\x1f\xf3\x9e\xd9\x99\x59\xde\xdf\x8b\x4c\xcd\xf2\x52\x89\xa0\x4a\xd2\xab\x64\xae\x02\xf1\xf0\x30\x9d\xd8\x17\x01\xf3\xf1\xf1\xd5\x1c\xc7\xe0\x51\x95\x19\x3e\x4d\xbd\x5d\xf9\xb2\xd2\xa6\xa9\x79\x57\x30\xcf\x9b\xc5\xea\x32\x4e\xf5\x72\x6f\xae\x5f\x5f\xe5\xcd\x1e\xfe\xc1\xb6\x4a\xe7\x65\x13\x4c\x27\x8b\xa6\xa9\x1a\x93\x94\x35\xee\x12\x5b\xd6\xb7\x0b\xf6\x70\x79\xb0\x09\xd7\xe4\x45\x91\xec\x55\x09\xc2\x9b\xeb\x95\xc9\xd2\xe6\x76\x03\x14\x0c\x12\xa4\x54\x97\x8d\xba\x6d\x06\x20\xdc\x7c\xa5\xcc\xd2\x52\x55\x2b\xb3\xce\x53\xb5\x05\x90\x9d\x1d\x27\xa8\x5d\xd4\x68\xa3\x68\x56\x17\x49\x39\x8f\xb5\x99\xef\xdd\xee\x95\xca\xa3\x03\x26\x55\x99\xea\x2c\x2f\xe7\x7b\xff\xac\x75\x89\xab\x67\x4b\x22\xb0\xd0\x73\xfc\xc1\xe5\x0e\x0b\xf0\xb8\xc0\xdf\xba\x31\x00\x61\x6d\x1f\x61\x6f\x1d\x6c\x51\x07\x80\x66\x0d\x4e\xf7\xf6\x48\x79\xa7\x48\x13\x0c\xbc\xad\x2a\x58\xfe\x5e\xd5\xa9\x44\xa8\xb5\xf0\x78\x8e\x8f\x71\x24\x14\x09\xad\x11\x4e\x5d\x00\xb4\x4e\x4d\x5e\x01\x04\x84\x76\x03\x2c\xb7\x73\xaf\xeb\x4a\xa5\xf9\x2c\x4f\xc5\x32\xcf\xb2\x42\xdd\x24\x80\x25\x81\xcd\xef\x15\x92\x70\xa2\xae\x57\xaa\x6e\x7e\x59\x95\xa9\x98\x69\x23\x4e\x8e\x4e\xcf\x66\xab\x42\xfc\x7a\x76\x76\x2c\x2c\xd6\xe9\x0c\x67\xb7\xd0\x88\xa8\x7b\x24\xe2\x68\x38\x18\x11\xf7\x20\x51\xa0\x6d\x9d\x98\x3c\xb9\x2c\x54\x2d\x1a\x2d\x56\xb5\x12\x45\xd2\x28\x03\xe6\xa1\x1a\x06\xbe\x7f\x20\x3e\xa8\xc6\x47\x36\x9d\x90\xba\x3e\x27\x4b\x9a\x0d\xfc\xb9\x80\xa1\x5a\xa1\x3a\xe2\x3d\x71\x4c\x58\x86\xb0\x0f\x07\x59\x7e\x32\xec\xed\x82\x9f\x64\x55\x34\x9e\x7c\xea\xe9\x04\xe8\x14\x95\x51\x15\xbc\x1d\x1a\x05\x34\xb6\x02\x8d\x3f\x75\x72\x3c\x10\x28\x1a\x99\x97\xa5\x32\xdd\x82\x23\xfb\x10\x0e\x87\x40\x0c\x93\x89\x51\xcd\xca\x94\xb4\x55\x48\x74\x09\x6b\x74\xf1\x21\xff\x46\xc2\xb0\x56\x04\xec\x50\x66\x96\xa4\xea\xfe\x21\x14\x12\x08\xab\x34\x68\xce\x1b\x8d\x84\x32\x06\xff\xb4\x09\x09\x36\xf2\x55\x15\x30\xb7\xd0\x45\xa6\xcc\x3e\xa8\xfa\xae\x59\x80\x21\x8a\x3b\xbd\x12\x37\x09\x90\x00\x82\xcf\x34\x1b\x49\xb3\x40\xbe\x9a\xbc\xb9\xb3\x5b\x7b\xff\x2e\xd5\x0c\xa5\x6c\x8d\x0d\xb7\x25\x4d\x72\x99\xd4\x0a\x17\x5b\x26\x88\x75\x64\xa2\x25\x3a\x84\x59\xd0\x19\x18\xb6\x2f\xc4\xf3\x2a\x7b\x31\x21\xee\x2e\x43\xb0\xbb\xed\x42\x24\xee\x96\x49\x85\x46\x63\x81\xc4\x12\xde\xbf\xb2\x47\x5f\xf8\x00\xa7\x56\x6e\x60\xbe\x8e\x02\x80\x0c\x62\x5b\x26\x4d\xae\x4b\x02\x15\x09\x7d\x85\xb0\x5c\x04\x8c\xd1\xb7\xac\xdf\x21\xe9\x28\xb0\x49\x3e\x13\x3f\xc0\x32\x52\xe1\xa4\x46\x7a\x60\x07\xd9\x7d\x7c\x84\x84\xfd\x86\x48\xcb\xa4\x68\xe7\xe3\x53\xf0\x2e\x65\x3e\xd5\x73\x10\x60\xb0\xcc\xeb\x1a\x95\xdc\x31\x2d\xba\x88\x0a\xff\x10\xe0\x01\xfa\xb3\xa1\x57\x16\x1e\x3e\x3e\x10\x07\xaa\x40\x74\x3f\x7d\xbd\x40\xdf\x3a\xbb\xab\xd0\xb5\xee\x1f\x70\xe6\x9a\xa4\x80\xdc\x07\x00\xd9\xdc\x05\x17\xb1\x64\xb2\xfe\x8a\xaf\x3d\x01\xd0\x38\xbe\xd7\x2c\x51\xe4\xd9\xba\xb5\x34\x8e\x4b\x9c\xf8\xe1\x40\x94\x79\xf1\xad\xcc\x42\x14\x8e\x4f\x2b\x50\x45\x33\x93\x01\x69\x4d\xe8\xcb\x26\xc9\x4b\x14\xc0\x8f\x35\x03\x12\xf2\xc7\x3a\x0c\x22\xd1\x86\x0e\x22\x29\x7c\x4a\x18\xf0\x1f\x84\x03\xb0\xc3\x3a\x3e\x2c\x74\xad\x64\xcb\x20\x84\x8a\x8c\xfc\x85\xe4\x0a\x46\xa8\x67\xf4\x9a\x67\xd3\x0e\x24\xd0\x99\x98\x74\x21\xaf\xc3\xf8\x6d\x51\x48\x55\xbc\x28\xd7\x35\x21\x63\xae\xf7\xe1\x2f\x88\x68\x2b\x87\xbc\xf8\xb3\x5e\x95\x32\x8c\x4f\x61\x7e\x55\x24\x46\x86\xcf\x93\x80\x65\xb7\x51\x45\xc1\xfc\xa1\x5b\x5a\x2e\x0a\x55\xca\x57\xc0\x94\xf8\x8b\x78\x63\xf9\x60\xd3\x00\x27\x5f\x07\x17\x00\x92\xa6\xbf\xbe\xb9\xe8\xc1\xfa\xce\xf1\x68\xc5\xa1\xe4\xe9\x78\x04\xa4\x8d\x07\xa3\x8f\x39\xb8\xc8\xf6\x50\x24\xfe\x73\x62\x11\x31\xc6\xab\x78\xe6\x60\x4b\xc4\x1d\x35\xb2\x4d\xbd\xae\xdb\xa8\xc6\x20\xb7\x87\x35\x58\x5c\xa0\x90\x60\x31\x6e\xfa\xda\x33\xa9\xe3\x62\x65\x92\x42\x86\x10\x0a\x5e\xf5\x42\x86\x23\x84\xb6\x1e\x30\x25\x7f\xfc\x21\x5e\xf9\xef\x4c\x99\x1d\x12\xcb\xe4\x4a\xc9\x1e\x8c\x48\xbc\x09\x5b\x5f\x7c\x04\x37\x6c\x46\x18\xdf\xdb\xc6\x88\x2e\x27\x1f\xcf\xae\x90\x92\x08\x19\xe8\x0c\x0a\x33\x2d\x03\x02\x65\x63\x69\x95\x49\x48\x5c\x4e\x51\x19\xdd\xe8\x54\x17\x3d\xf3\x3b\xb6\x83\xff\x1d\x26\x48\x7a\x68\x03\xfa\xbf\x61\x7d\x35\x48\x26\x5d\x88\x75\x2c\x1b\x50\xb5\xb5\xee\x14\xdc\x57\x8c\x5b\xe1\xbe\x85\x62\x29\x6c\x63\xe5\x6d\x05\xe9\xeb\x89\x1d\x96\xeb\xad\x46\x1c\xda\x68\x8e\x8a\xd8\x02\xeb\xb3\xba\x69\x01\x8d\x9a\x59\x24\xd6\x61\xc7\x41\xcb\x51\x67\x01\x73\x05\xf2\x80\x70\x54\x77\x06\x80\xa5\x0a\x1e\xc7\x1a\xce\xe0\x85\x4a\xaf\x40\x8d\x5d\x6a\x89\x99\x6e\xe9\x70\x1e\xc3\xca\x43\xbb\x64\xdf\x29\xde\xdb\xce\x4c\x85\xa3\x76\x32\x48\x24\x77\x34\x98\x17\xb6\x98\xa7\x4d\x66\xd4\x66\x3c\xa3\xb1\x32\x9f\x2c\x11\x06\x0a\x25\x86\xba\xe0\xd3\xea\xb6\x4d\x99\xec\x29\xb6\xc4\xa3\x56\xdf\x30\x82\x4e\x78\x2d\x32\xa7\xc2\xe7\xe2\xf3\x8c\x64\xed\xdb\xb1\x0b\x4a\xe3\xba\x67\x51\x3d\x47\xf5\xb4\xf2\x7f\x54\xf3\xdf\x43\x5f\xcf\x55\xd6\xe4\xa9\x22\x64\x43\x5b\xb6\xda\xf3\xaa\xe0\x74\x01\x49\x63\x8d\x89\x5c\x02\xc9\x8e\x2b\x1e\x9d\xd0\xa0\x08\xa4\x92\xf1\x14\xb8\x68\xf7\xc8\x20\xa5\x6a\x10\x32\xcb\x56\xb8\x87\x08\x46\x02\xae\x8d\xf0\x1e\x75\x43\x5c\x42\xe2\xc0\xa8\x09\x38\xa8\x22\xf8\xbf\x2d\xb9\x5b\x88\xd1\x6c\x8c\x1c\x90\x80\xc9\xd5\x7a\x57\x82\x30\x03\x6a\xc9\x19\x04\xa3\x0e\xec\x37\x50\xc4\xe9\xd9\x8e\xf4\x70\x79\xf8\x88\x80\x6c\xd2\xb7\x3b\x39\x78\xa0\x7f\x67\xe1\x50\x8e\xb0\x3b\x25\x99\x2a\xd4\xf3\x05\xb3\x45\x0e\x0c\xe4\x71\xec\x68\xee\x64\xf2\x07\xf0\x8f\x5e\xac\x03\x60\xc3\xe7\xb7\xf7\x00\x96\x43\x57\x9e\x82\xdd\xdf\xd9\xf1\x4c\xec\xe7\x19\x54\x39\xaa\xc8\x68\xb5\xbc\xd1\xe6\xaa\xe6\x6c\x06\x70\x09\xcc\xb3\x33\x71\x7e\xf2\x51\x18\xbd\x6a\x30\x38\x71\x4e\xe3\x43\xed\xf5\x0c\xe3\x61\x83\xc9\x56\xfa\x46\xbc\xc2\x95\xb1\x9d\xa2\xd0\x32\x08\x38\x83\xc8\x02\xd4\x61\xea\x1a\x03\x09\x5c\x7b\x02\xdf\x10\x61\x64\x00\x74\x07\xa1\xa0\xbe\x17\xb8\x31\x38\x75\x39\xc7\xcc\x2d\xc5\x56\xc9\xbe\x77\xd2\x43\x78\xcb\x6a\xd8\xf4\x36\xcb\x64\x00\x7b\x22\xa8\xdd\x42\x8a\xa5\x8c\xdc\x5b\xea\xe0\x83\x06\x79\x17\x02\x0b\xa7\x5d\x04\x7a\xf0\x45\x8a\xf6\x82\xec\xf3\x9b\xa0\x3a\x99\x9a\x69\x85\x5f\x67\xf8\xe2\x72\x3b\x5e\x50\x60\x53\x5b\xbb\x6f\xf2\x84\x13\x98\x2f\x27\x06\x22\x79\x8d\xbd\x5d\x78\x84\xfa\x98\xba\x70\x13\x1c\x38\x6d\x0c\x4b\xfa\x17\x6d\x96\x7f\x4b\x8a\x15\xd8\x6e\x4d\xad\x63\x94\x00\x04\x66\xb7\x08\x82\x73\x10\x70\x6c\xa6\x79\x46\x47\x5d\x4f\x28\x55\x8b\xbc\x91\x76\x65\x24\x82\x88\x36\x4f\x50\x2a\xff\x88\x18\x2f\xe2\x40\x5d\x09\xde\xcc\x31\xfe\x3a\x3e\x85\x37\xda\xe8\xc5\x70\x9f\xe6\x2a\x99\xfb\x9d\x0e\x9f\x7a\x3d\x9b\xd5\x0a\xce\xaf\x22\x5f\xe6\x8d\x3b\x57\x47\xa4\xa7\x61\x89\x58\x81\xec\xfe\xfc\x27\x7b\x6a\xe9\x7a\x84\x69\x06\xc7\x84\x17\x63\x0b\x08\x4f\xe0\x12\x60\x82\xe1\xc9\x84\xc6\x9a\x36\x0d\xb2\xad\xe1\xf8\x18\xb9\x38\x07\xe4\x12\x37\x44\xe2\xff\xdf\x44\x02\xe8\xf8\x99\xcb\xbb\xde\x79\xa7\xc1\x00\x74\xd3\x9d\x77\x0f\xae\xbe\x1a\xc1\x54\x3c\x86\xa9\x78\x0a\x53\x81\xe5\xd4\x26\xa6\x2e\xdd\xc5\xf6\x0d\xeb\xc0\x1d\x0c\x53\x52\x95\x6a\xbe\x90\x8c\x24\x8b\x2a\x74\xa3\x1f\x51\x30\x92\xc4\x43\x1b\x3b\x17\xbb\xde\xe6\x44\xbf\x9f\x7e\xf9\x8c\x2e\xc1\x93\xb5\x48\x46\x1a\xd4\xcd\x22\x69\x9c\x97\x59\x90\x04\x01\x58\xd4\x78\x86\xeb\x14\xce\x88\x4c\x2c\xd5\x52\x83\x07\x02\xcf\xab\x14\x80\x29\xdf\xf3\x1c\x9a\x17\x0d\x55\x40\x92\xa3\xa6\x2b\x35\x3b\x19\xfc\xd4\xef\xb2\xf1\x86\x4d\xb6\xe0\x08\x49\x51\x99\x78\xf7\x80\xee\xcb\x24\x1a\x69\xe2\x77\x3a\xbb\x43\x49\x73\x3a\x05\xcb\x2c\xf9\xd2\xcb\x7d\x46\x45\x6c\x7b\xb0\x3b\x4b\xb8\x93\x9e\x85\xf0\xa2\xc2\x73\xad\x57\x6a\x0d\x3c\xd2\x76\xe5\xce\x13\x8b\x37\xb8\x70\xdd\x91\x9e\x96\xb9\xeb\x38\x92\x4a\x7a\x85\x5c\x0b\xc9\xb6\x37\xfb\x80\xe8\x5c\xdb\x09\x4c\xdb\x0a\xc3\x5e\x41\x4f\xed\x38\xbf\xa1\x1a\x97\x26\xb0\x08\x51\x76\x5e\x82\xd9\x63\x25\x1c\x5f\xec\xa5\x7f\x1e\xc1\x5b\x16\xdb\x54\xa8\x77\x04\x6d\x59\xda\x26\x71\xbe\xde\xb7\xac\x6d\xf3\x9a\x1e\x05\x53\x97\x9c\xe3\x96\xe9\xf0\x9a\xeb\x04\xb2\x08\xa8\x9c\x38\x2b\xaf\xa9\x2b\x83\x89\x05\xb6\x5e\xdd\x85\x15\xbd\xd7\xb8\x11\xdc\xfb\x52\xd1\x75\x51\x5e\xd2\x52\x1f\x92\xa0\x4b\x9d\xe1\x25\x15\x63\x90\x25\xde\x1b\xb1\x0d\x45\xa2\x84\x7c\xa9\x77\x37\x85\x09\x14\x18\xa6\x9b\x77\x75\xd7\xbd\x5f\x32\x6c\xd2\xc5\x57\x4a\x78\x94\x65\xf9\x6c\xa6\x0c\x76\x83\x93\x14\xfb\xfd\x02\x3b\x1b\xdc\x19\x72\xad\x26\x38\x9b\x16\x3a\xe3\x6e\x71\xd2\xae\xeb\x92\x02\xdb\x15\x21\x2a\xd1\xa8\xa8\x21\xe2\xe4\xbf\xef\x95\x6b\x88\x35\xfe\x1d\x36\x49\x64\xa2\xd7\xd7\x0d\xee\xf3\xec\x21\xc0\x87\xe3\xf3\xb3\xa0\x05\xd2\x1a\xc7\xce\x60\x3e\x1c\x79\x60\xac\x7e\x77\x06\xf2\xfe\xe8\xe3\xd1\xd9\x51\x07\x87\xcc\xcf\x87\x42\x7b\xbb\x4e\x4b\x1f\xab\x75\x83\xc7\xd6\x1f\x7f\x39\xc5\x0d\x0f\xad\xa9\x05\x60\x84\x41\x30\x66\x6c\xe8\x81\x97\x39\x64\x72\xf6\x76\x00\x0c\x0a\x93\x5a\xd6\xec\x88\xe9\xe0\xd5\x0c\x84\x2f\x58\x13\x93\x19\xc1\xe9\x59\x71\x51\x0b\x15\x6d\x44\xc9\x30\x1b\x13\xff\x5f\x79\x86\x33\x9d\x14\x7a\x1e\x1f\xdb\x4e\x3d\x56\x95\x24\x2e\x6c\xd0\xef\x61\x8f\xde\xdb\x6c\xef\x21\x21\x63\x85\x22\x01\x4c\x3b\xcb\xd7\x79\xb6\x4a\x0a\xa6\xab\xcb\xbd\xe5\x86\x05\x86\x74\xad\xc9\x84\x51\xc0\xf5\x52\x9d\x8e\x64\x30\x6b\x1a\x74\xb6\x3d\xb8\x47\xe6\xa1\xf8\x57\x42\x0f\xcb\xfd\x20\xec\xb5\x08\x7a\x0e\xb4\xe5\x4a\x1a\x40\x8d\x01\xb3\x61\x92\x52\x6e\x2e\x58\x3e\x91\x3b\x10\xcc\xd0\x42\x23\x08\x6e\x64\x11\xfa\x0a\xdf\x68\xb7\xba\x86\x8b\x38\x3c\x39\x7f\xef\x17\xea\xed\x23\x4a\xc2\x57\xa4\x6b\x7c\xd4\xd2\x53\x56\xb8\x01\x8d\xae\xb0\x65\x57\xaa\x45\xc3\x83\x31\xa4\x1d\xe8\xf2\xe8\xd3\x9b\xe8\xdb\x3b\xe6\x5e\x64\x51\x37\x7c\xdb\x4c\xfa\x06\x3e\x36\x26\xa9\x74\xf3\xa9\xc2\x35\xc3\xf8\x65\x83\xee\x08\xf4\xee\x13\x01\xb7\xe6\xe0\xb1\x4b\x7a\xe4\x7a\x49\x7d\x0f\x00\x35\x2c\x41\xdb\xdb\xc9\xc3\x42\x25\xe6\x03\x7f\xbb\x41\x45\x31\x5a\xfd\x79\xad\xb0\x95\x53\x51\x75\xc9\xa6\x36\x13\xdf\x60\x5d\x91\x60\x12\x46\xfa\x5a\x91\xd8\xf2\x29\xc1\xbd\xcd\x90\xe8\x73\x0c\x7b\x1e\xd6\x95\x7d\xad\xc5\x3c\x5f\xab\x52\xf8\xcd\x29\x2d\x70\x11\x66\x48\xfd\x1d\xce\x4b\x6e\x84\x4d\x4b\x78\xcb\xdf\x4d\x4e\x0e\x3e\xd6\xdf\x02\x1f\xda\x6c\x70\xa9\xb2\x97\x97\x1d\x95\x9c\x97\xdd\x90\xe1\x72\xee\x00\x4b\x62\x1e\x6f\x9b\x66\xe1\x30\x5f\x00\x96\x90\x34\xba\xfb\xb3\x50\x84\xa2\xe6\xb6\x63\x8a\x7d\x08\xf8\x21\x95\x1e\xf1\xdb\x02\x27\x68\xb9\xe8\x38\x1d\x00\x7a\x82\xd7\x41\xbf\x16\x88\xb9\x5e\xe5\xe9\x95\x98\xe5\xb7\x64\xe7\x73\x7d\x95\x63\xd5\x9f\xb5\x95\x16\x5e\x76\x54\x58\x7a\x55\x46\x5f\x16\x6a\xe9\xf5\xf4\xf1\xb2\x72\xd0\xd5\xef\x67\x89\xef\x92\xcc\xfa\x13\xd1\xb9\xef\x75\x04\x69\xf7\xa3\xab\x43\xbc\x21\xed\xfa\xf0\xbb\xa9\xc0\xbf\x36\x20\x68\xa8\xd2\x30\xf4\xd4\xc0\xe1\x16\x9b\x80\x9e\x53\x53\xad\x5a\xda\x8f\x4e\xb8\x50\xed\xc2\x4c\x2b\x34\x48\xf7\x5e\xd3\x2d\x50\xfb\x2d\x4e\xef\xeb\x1b\x31\xf2\xf5\x0d\x11\x59\xb5\x9f\xa4\x7c\xe8\x75\x89\x28\x0c\xca\x16\xd1\x57\x7c\xbf\x08\x9f\xc0\xc6\x91\xac\x18\xe2\x18\x73\xf4\x47\x5c\x7d\xe0\xec\x52\x55\xde\x0d\xbf\x0d\x97\xa8\x29\x7b\x66\x91\x4e\x16\x2e\x36\x75\xea\x03\xa5\xf0\xdd\xf5\x06\x42\x54\xd6\xb2\x6a\xee\x64\xc8\xe8\x54\xc5\xbf\x4e\x10\x5e\x5a\xc9\xe7\x01\xaf\xea\xb9\x31\x8f\xf5\xf1\x31\xb2\x77\x74\x3f\x2c\x5b\x6c\xc0\x87\x65\x2e\xdc\xbe\xc9\xf7\x1a\xb9\xe9\x46\x61\xcb\x3d\x9b\x08\x06\x7f\xcb\x39\x64\x72\xa9\x36\xf8\x79\x98\x4d\x26\x16\xf4\x55\x15\xa7\x7b\xd3\xae\x83\x8d\xc6\x33\x23\x6e\xa2\xde\x27\x48\x78\xd4\xfd\x3c\xec\x71\x83\x91\xe6\xa9\x74\x17\xf5\xae\xfb\xf1\x20\x7b\xe7\x3d\x14\x24\x36\x90\x32\x4c\xa0\x92\x3f\x2e\xb3\xdf\x99\xfd\x2b\x00\x00\xff\xff\xc1\x6b\xa7\x03\x2c\x28\x00\x00")

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

	info := bindataFileInfo{name: "rest/gorilla/pat.tpl", size: 10284, mode: os.FileMode(436), modTime: time.Unix(1447262583, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _storeUpperioTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x58\xcd\x6e\xdb\x38\x10\x3e\x5b\x4f\xc1\x35\xb0\x05\x15\x04\x72\x0f\x45\x0f\x59\xe4\xd0\x8d\xdb\x45\x80\xf4\x67\x93\x76\x2f\x45\x51\xc8\xd2\xd8\x26\x4c\x91\x0a\x49\xa5\x0d\x02\xbf\xfb\xce\x90\x94\xac\x58\x72\xd3\x14\xb9\xec\x1e\x9a\x4a\xe4\x70\x7e\xbf\xf9\x86\xd6\xdd\x1d\x2b\x61\x29\x14\xb0\x69\x9d\x17\x9b\x7c\x05\x53\xb6\xdd\x26\x93\xf8\xc2\x70\x3f\xfb\xb0\x59\xd1\x1a\x3e\x82\x2a\xe9\x29\xe9\x9d\x12\x55\xad\x8d\xb3\xe1\xd4\x54\x81\x9b\xad\x9d\xab\xa7\xf8\xbc\x12\x6e\xdd\x2c\xb2\x42\x57\xb3\x95\x6e\x4c\x39\xdb\x08\x37\xb3\x4e\x1b\xf8\xf1\xee\xac\xa9\x6b\x30\x42\xa3\x14\xda\x11\x4b\x96\x9d\x97\xd9\xb9\xbd\x72\x46\xa8\x55\x30\xd3\x3b\x6c\x73\x3c\x23\x50\x47\xd6\x34\xa2\x24\xcd\xa0\x0a\x5d\xa2\xe8\x6c\x91\x5b\x78\xf9\x82\x96\xac\x3f\x6b\x83\xc6\x18\xc4\x64\x2a\xf5\x8a\x36\xbd\xb9\x4c\xe8\x59\xb9\x98\x1e\x08\x12\x15\x86\xbc\x24\xcb\x46\x15\x4c\x28\xe1\x78\xca\xee\x92\xc9\x6c\xd6\xca\x78\xd7\x59\x6d\xf4\x8d\x28\xc1\xb0\x6f\xe8\x21\xbd\x7d\xbf\x4d\x26\x7e\x2b\xfb\x10\xb7\x6c\x36\xf7\x27\xde\xa0\x26\x3e\xa5\xfc\x7e\xbc\xad\x21\x7b\x97\x57\x80\x16\xa6\xc7\xcc\x9b\xe0\x86\x1d\x51\x22\xb3\x4b\xb8\x6e\xc0\xba\x94\x71\x1b\x6c\x64\x57\xf4\xf7\x98\x81\x31\xf4\x4f\x1b\xef\xc8\xc4\x80\x6b\x8c\x62\x7f\x81\xdb\x57\xe9\xe5\xb9\x49\x93\xc9\x36\x4d\x30\x04\xf4\xf9\x90\x54\xeb\xbf\x65\x26\xff\xc6\x46\x45\x42\x06\x0e\x9b\x19\x71\xfb\x68\x54\x74\x2f\x02\x9f\x4b\xbd\x70\xb9\x50\xac\xcc\x5d\x4e\xc5\x4b\x26\xe5\x22\x88\x9d\x9c\xb2\x88\x8a\xec\x7d\x0d\x8a\x9b\x63\x36\xc5\xc4\xe7\x8d\x74\x53\x0c\x0c\x41\x42\x52\xbf\x9d\x32\x25\x64\x2f\x1d\x18\x72\x32\x2c\x52\x8e\x05\x6e\xf7\x2d\x3b\x65\xcf\x46\xfd\xbb\x2b\x17\x08\x92\x28\x17\xd2\x36\x9e\x33\x0b\xe6\x06\x33\xb6\x02\x85\xfe\x15\xec\xec\xd3\xe5\x9c\x2d\xb5\x61\x0e\x25\x07\x47\x42\xf6\x51\x32\x77\x50\xb2\xc5\x2d\xf3\xe0\x67\x67\x17\xe7\xcc\x69\x2d\x93\xd1\x43\xd1\x8e\x33\x4d\xe1\x28\xba\xf9\x82\x95\x8b\x6c\xde\x66\x29\x38\x77\x66\x00\x75\xb2\x7c\x70\x1a\xe1\xca\xdc\x1a\xba\xac\x1e\x33\xbd\xf4\x0b\x75\x6e\x40\xb9\x50\xd0\x83\x65\x4a\xa3\x62\x9e\x4c\x0a\x8d\x89\x0b\x18\x3c\xc3\x47\x8b\xa5\xa9\xe3\xfb\x6b\xe5\x84\xbb\xfd\xe0\xb0\x94\x7c\x58\xd6\x15\x38\x56\x68\x29\xa1\x70\x42\x2b\x52\x24\x65\x57\x57\x8b\xca\xa4\xe4\x3f\x53\xc5\xbc\xae\xe5\x2d\x42\x53\x95\xba\x62\xd4\xee\x2c\x34\x36\xa6\xae\x7d\x12\xe5\x41\xd6\x20\x79\xc2\x11\xfe\x9f\xbd\x83\x6f\xff\xbc\x20\x9b\x40\x4b\x50\x67\x1c\xa3\xdf\x05\xbf\xdd\xd2\x56\x46\x09\x41\x2d\x6d\x1e\x4f\xa3\x11\x9b\x7d\x34\xa2\xba\x14\xab\xb5\xe3\x81\x63\xb2\x4f\x97\x17\xaf\x23\xed\x64\xfe\x01\x3e\xea\x60\x9a\xa3\xb9\xcf\x27\x5f\x52\x04\xec\x29\x41\xb5\x47\x30\x14\xd2\xdb\xdc\xd8\x75\x2e\x7d\x3d\x84\x83\xea\x98\x5c\xaf\xb5\xb5\x62\x21\xc1\x4b\xf0\xeb\x46\x14\x1b\xb6\x14\xdf\x3d\xaa\x62\x1b\x50\xa7\xa2\x44\x15\x88\x06\xd1\x10\x15\x81\x09\x89\xac\xa8\xcc\x9b\x36\xb8\x7b\xfb\x7f\xd0\x06\x65\x17\xea\x50\x84\x53\x94\x6e\xf7\xe7\x7f\x52\x56\x46\x4a\xd1\xd5\x02\x8b\xd1\xd5\xa3\x2c\xbd\xe3\xe0\xab\x4f\x55\xe8\x57\xf9\x40\x19\xbe\xb6\x46\x49\x36\x7b\x85\xe1\xa8\x92\x43\x1d\x53\x23\xad\xef\x92\x89\x28\x3b\x80\x8c\xca\x45\x0e\x1f\xfa\x89\xa4\x8e\x54\x2b\x94\x5b\xf2\xe9\x6b\x42\x21\x2b\x08\xbf\x64\x7d\x1f\xdf\x27\xec\x77\x3b\xf5\x66\x32\x2f\xc9\x53\x0a\x3d\x38\x17\x61\x4d\xcb\xe7\xca\x81\x51\xb9\xdc\xc3\xe3\x68\x78\xfd\x10\x3a\xc0\x52\x8a\x36\xe0\xf3\xb3\xcb\xd6\x63\xa1\x17\x17\x48\x0e\x17\x38\x62\x98\x63\x90\x2f\x5f\xa4\x7b\x98\xba\x47\x58\x57\x90\x9b\x62\x3d\xc6\x09\xc8\x3c\xc2\x59\x46\x2d\x2d\xa8\x5c\xdc\xa6\x0f\x12\x41\xd0\x86\x44\x70\x1d\xd3\xf3\x77\x03\xe6\x36\x8d\x2f\x97\x60\x91\x8d\x7d\xcb\xc7\x59\xd4\x32\x36\x36\x5b\xd8\xe4\x64\x01\x07\x27\x37\xc8\x97\x88\xc9\xb0\x3a\x18\x65\x23\x8c\x71\x88\x32\x1e\x00\x6a\x50\x86\x0b\x46\xc0\x4d\xcc\xbc\x40\xdb\xc4\xbb\xe2\x06\x14\xbb\xa6\x08\x76\x59\xb0\xde\x12\x52\x5b\x7f\xe0\x78\xae\xe3\xd7\x19\xce\xbc\xf0\x98\x46\xc3\x41\xf2\xf4\x9e\x69\xdb\x02\xfb\x8d\x40\xb8\x92\xe0\x36\x20\x62\x74\xdf\x6b\x48\x7b\xae\xae\x91\xdb\x24\x31\xf3\x0a\xf1\x14\xac\x78\xc3\xef\x97\x4b\x0b\x74\xe7\xc0\x40\x9f\xf7\x75\xe1\xdf\xec\x6a\x23\x6a\xa4\x19\xe5\xf8\x3d\xd9\x34\x28\xde\x29\xb9\x10\x95\x38\xa4\x23\xec\xed\x94\x44\xd9\xb4\x73\xae\x83\x7e\x9a\x44\x70\xbd\xc7\x81\x1a\x56\xad\xc7\xf5\x52\x18\xeb\x86\x48\xab\x72\x57\xac\xe1\x91\x50\x43\xdd\x34\x70\x7e\x69\xda\x74\xe5\x36\x1e\x5e\x96\x2d\x0d\x8e\x8b\xdd\x9d\x42\x52\x71\x9f\x7d\xfe\xb2\x6f\xfa\x6e\x4b\xc8\x3e\x69\x7b\x1f\x51\xeb\xf1\xcd\xd3\xec\xaa\xad\x7c\x91\xc6\xeb\x44\x53\xd5\x9d\x7a\xcc\x19\xd2\xb1\x16\x44\x13\x34\x5a\xab\xbc\x66\xb3\x38\xad\x93\x96\x4e\xb2\xd8\x3c\xd7\x69\xf6\x0a\xa1\x2b\x7f\x66\xdc\xa1\x80\xd2\x0e\x59\xbf\x51\xc8\x86\x06\xe8\xa2\xed\x8f\x49\xbc\x01\x1d\xc9\x94\xa0\x17\x0a\x39\xe0\xac\x77\xda\xbd\xa1\x63\x23\x33\x14\x47\xcb\x2a\x5c\x08\x6e\x72\xd9\x00\x79\x1c\x7a\xc1\x87\xe0\x65\x22\x51\x85\x82\xb6\xe9\x2c\x83\x7c\x32\xe1\x47\x91\xb4\xee\xa7\x2f\x45\x7f\x18\xb9\xf5\xf9\xf9\x97\x8e\x02\x30\xb2\x88\x96\x4f\x75\x49\xd7\x93\x01\x3c\xb4\x7a\x1c\x32\x82\x9e\x5f\x06\xc7\x53\x5d\x45\x48\xcf\xa2\xc7\x1c\xfe\x52\x89\x89\xa5\x8b\x9a\x27\x79\xb2\x1a\xae\x4c\xc7\xec\xab\x1f\x63\xd4\x56\x6f\xf3\x9a\xd4\x53\xdf\x9d\xf4\x89\x00\xc9\x90\xe2\xf0\x84\x90\xa6\xff\x83\xab\x41\x13\xca\x4d\xbe\x07\x02\x28\x3b\xf6\x6d\x7b\x82\x68\x27\x56\xd3\x8f\xf5\x9f\x99\xe5\x5e\xed\x93\xcc\xf2\xfd\xbb\xfd\x1c\x24\x3c\x05\x3e\x83\x9e\x7d\x7c\xfe\xd7\xb1\x68\xa0\xd2\x37\x0f\x97\xf3\xd2\x8b\x8d\x3b\x3d\x2c\x66\x49\xa9\x7a\xda\x62\xf6\x08\x07\x69\x56\x17\x81\x0c\x58\x4e\xcf\x84\xc7\x0a\x1d\xc4\x81\x4f\xdd\x92\xab\xf6\x26\xf6\x50\x49\x7b\x9a\x78\xba\xcf\x32\x14\x5b\xb4\x3d\xf8\x1d\x89\x13\x65\xe0\xcb\x85\x40\x4e\xfd\xa1\x3f\x4c\xa2\xc8\x63\x9c\x22\x95\x7b\x8e\xd1\xd2\x9e\x73\xa3\x03\x2f\xb8\x77\x81\xec\x2f\x94\xad\x11\x88\xbe\xc4\x38\x5f\x56\x48\x1d\x38\x19\x1e\xe9\x14\x2a\xe2\xb5\x1c\xf3\x24\x65\xfe\xb6\x4a\x0e\x81\x1f\xc0\xb5\xc4\x19\x32\xf4\x29\xed\x1c\xf6\xf2\xdc\x8f\x3a\x90\x69\xfb\xe5\x82\x9a\x22\xde\x3a\xbc\xab\xf4\x9d\xa2\xfd\x80\xd3\x6f\xa6\x07\x7f\xd4\xfa\xe6\x62\x9c\x8e\x30\x8f\xf8\xf6\xe8\xfe\x75\x34\x36\x18\x19\x1a\x6f\x56\xea\xd5\x79\x5f\x03\x7d\xd5\xf1\xaf\xdb\xed\xf8\xc7\x89\x91\xdf\x28\x5a\x29\x3a\x8c\xcd\xb0\x33\xc2\x3a\x35\xa1\x21\x88\x6a\x7f\xa5\x27\xda\xd4\x49\x6d\xe1\xde\x97\x00\x66\x01\xa7\x89\xa6\x4c\xe6\x23\x37\x37\x61\x59\x63\xe9\x1e\xfa\x60\x32\x49\x33\x66\xd3\x27\xad\x87\xb9\x90\x97\xb0\x49\x4e\xf4\xbe\xae\xfd\x1b\x00\x00\xff\xff\xa9\x23\x87\x88\x78\x14\x00\x00")

func storeUpperioTplBytes() ([]byte, error) {
	return bindataRead(
		_storeUpperioTpl,
		"store/upperio.tpl",
	)
}

func storeUpperioTpl() (*asset, error) {
	bytes, err := storeUpperioTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "store/upperio.tpl", size: 5240, mode: os.FileMode(436), modTime: time.Unix(1446633715, 0)}
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
	"endpoints/endpoints.tpl": endpointsEndpointsTpl,
	"general.tpl": generalTpl,
	"rest/gorilla/pat.tpl": restGorillaPatTpl,
	"store/upperio.tpl": storeUpperioTpl,
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
	"endpoints": &bintree{nil, map[string]*bintree{
		"endpoints.tpl": &bintree{endpointsEndpointsTpl, map[string]*bintree{}},
	}},
	"general.tpl": &bintree{generalTpl, map[string]*bintree{}},
	"rest": &bintree{nil, map[string]*bintree{
		"gorilla": &bintree{nil, map[string]*bintree{
			"pat.tpl": &bintree{restGorillaPatTpl, map[string]*bintree{}},
		}},
	}},
	"store": &bintree{nil, map[string]*bintree{
		"upperio.tpl": &bintree{storeUpperioTpl, map[string]*bintree{}},
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

