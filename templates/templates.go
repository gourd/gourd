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

var _endpointsEndpointsTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x58\x4f\x4f\xec\x36\x10\x3f\x67\x3f\x85\x9b\x8a\xca\x41\xdb\x6c\xcf\x5b\x71\xa8\x80\xbe\x56\xa2\x40\x1f\xbc\x13\xe2\x10\x92\x49\xb0\x48\xec\x60\x3b\xfc\xd1\x6a\xbf\x7b\x67\x6c\x67\x37\xec\x52\xe8\x83\xae\xf4\x78\xda\x03\x90\xcc\x8c\xc7\xbf\xf9\xf3\x9b\xd8\xcc\x66\xac\x80\x52\x48\x60\x71\x9b\xe5\x37\x59\x05\x31\x9b\xcf\x47\x51\x78\x61\xa8\x4f\x4f\x6f\x2a\x92\xe1\x23\xc8\x82\x9e\x46\x83\x55\xa2\x69\x95\xb6\xc6\xaf\x8a\x2b\x61\xaf\xbb\xab\x34\x57\xcd\xa4\x52\x3f\xdf\x08\x3b\xa1\x1f\x5c\xd6\x2a\x21\x6d\x3c\x8a\x2a\xd5\xe9\x22\xb7\x0f\xec\xa9\x29\x0a\x9d\x65\xae\xa4\x85\x07\x32\x7c\x5e\x6f\xac\xd2\xe0\xb4\xaa\xce\x64\x95\x2a\x5d\x4d\x1e\x26\x12\x06\x2b\x51\x59\x36\xce\x43\xad\xaa\xf8\x5f\x50\xe7\xaa\xf0\x81\x8e\x26\x13\x17\xe3\x19\x39\x4e\x8f\xb3\x06\x50\x7a\x18\xf0\x1a\xa6\xc1\x76\x5a\xb2\xfd\x2f\x9f\x0f\x18\x2c\xa4\xa5\xd2\xeb\x8b\x46\x65\x27\xf3\x17\x7c\x71\xa9\x3a\x39\x66\xf4\xbb\x65\xc6\x6a\x21\xab\x84\xf1\xa5\xd3\x26\x6b\x2f\xbc\xf8\xb2\x17\xa6\xfd\xe2\x84\xcd\x30\x2e\x84\x7a\x97\x69\x91\x5d\xd5\x60\x98\x55\xac\x33\xc0\xea\xcc\x82\x1e\x45\x59\x5d\xab\xfc\x50\x5a\x61\x1f\x8f\x84\xb1\x6c\xba\xc7\x08\x0e\x4f\xd8\xee\xc5\x25\x61\x3a\x7f\x6c\x09\x0d\x9b\xf5\x21\xfd\xf4\x44\x3e\x9b\x33\x2c\x5f\x05\xd6\x61\xa7\xe5\x9f\xc0\xae\x87\x18\xb9\xfc\xbb\x37\x34\x89\xd7\x0c\x62\x8f\xd2\x59\x0d\xf2\x75\x0d\x1a\x9c\xe2\xfc\xe4\xe0\x64\x8a\x91\x3e\xb2\xeb\xec\x0e\x98\x84\x7b\xca\x44\x97\x5b\x0a\xc7\xad\x1a\x45\xcb\x65\x7b\x68\x79\x03\xfc\xc5\xc4\x8c\x06\x0b\x2e\xe2\x5c\x03\xe6\x23\xbe\x64\x21\x7c\x6a\xb4\xd0\x18\xe9\xbe\xff\x3b\x66\xc0\xd0\x18\x74\x99\xe5\x30\x9b\x63\x09\x34\x66\x73\x20\x41\x03\xad\xe9\x47\x69\x9f\x76\x02\x8e\x99\xe9\x1d\xa1\x2d\xd6\xbf\xc9\xac\x50\x12\x95\x9a\x32\xd1\x77\x75\xfa\xc7\xf9\xf9\xe9\x67\xb8\xed\xc0\x58\xda\x3c\x41\x03\x51\x32\xcd\xf6\xf6\x98\x14\x35\xba\x8b\xa2\xc8\x90\x7f\x5c\xe4\xe2\x4d\x0f\x69\xa3\x3f\x69\x7b\x99\xd5\xbd\x3a\x3d\x03\x7d\x07\xfa\x2f\x53\x61\x24\x71\x23\x8c\xc1\xe8\xb1\x72\xce\x31\x02\x60\x4b\x9e\x44\x11\xb9\x43\x6f\xf8\x87\xde\x7c\x79\xf1\x69\x3e\x80\x1e\x52\x1b\x19\x1f\x1c\x21\x0e\xa5\xe6\x3a\x60\x24\xf9\x0f\xef\x40\x89\x9c\x4b\xcf\x5a\xac\x92\x2d\x79\xec\xb2\xc7\xd4\x95\xcd\x84\x24\xe4\x3b\x26\xf4\x04\xdf\x31\x49\x3c\x66\x8b\x36\x72\x78\x92\x17\xa3\x88\x22\x24\x2d\x68\x66\xd2\xfd\x5a\x19\xe0\x49\x88\xcb\xd7\x1a\xbb\x8c\x9a\x1e\x45\xc8\xf6\xf4\x34\xec\xef\x75\x53\xb6\xf3\xe3\x1d\xee\x06\xb4\x41\xf0\x9f\xee\x3b\x15\xc7\x38\x83\xe2\x7f\x8e\x9d\x0c\xa2\x90\x00\x07\xc3\xc7\xef\xb0\x8c\x03\xda\x1e\x98\x9f\x08\xe4\xcd\xed\xc1\x93\x00\xe9\xb5\x92\x82\xa4\xf9\x85\xfd\x60\x5a\x25\x0d\x15\x96\x7a\x78\x6f\x38\x42\x06\x0d\xed\x42\x72\x63\x67\xba\xca\xfa\x5d\x48\xf9\xee\x40\x90\xcc\xc7\x21\xe7\x61\xcf\x91\xdb\x73\xc0\x30\x94\x6b\x01\x77\xaf\x71\x6c\xd9\xaa\x5f\xc5\xb4\x5b\xca\x7a\x58\x9b\x72\x9f\xfe\xbf\x3b\xd0\x8f\xc9\x96\x87\xdf\x2a\x0f\xdd\xa7\x87\x98\xd8\x40\xa3\xf4\xa3\xfb\x38\x2e\x3e\x53\x44\xbc\x9a\xe0\xaf\x7c\xa0\x16\xab\xfb\x7e\x1a\x30\xf4\x0c\x32\x9d\x5f\xf3\xdb\x24\xfd\xad\xae\x39\xd4\x9b\x4b\x8c\x71\x3b\x2d\x08\x6a\xe2\xb1\x63\xef\x82\x96\x6f\xe3\x22\x62\x35\xe9\x11\x48\x82\x4e\x0d\xf7\x8b\xc7\x1b\xfc\x2c\xd1\x1e\x2b\xfb\x3b\xee\x54\xac\xbb\xfd\xaf\x6c\x86\xfa\x35\xba\xd6\x98\xeb\x2d\x55\xb7\x54\x7d\x37\x55\xb1\x1f\xba\x1a\xcf\x64\xd3\x27\x0c\x45\x45\x8e\xad\x68\x17\x39\x0a\x66\xd8\x5d\x28\xe5\x9b\xcb\x93\xdb\xf5\x1d\xc4\xf5\xba\x1e\xed\x47\x98\x33\xfe\xe8\x2c\x01\x0a\x3a\x2b\x97\xe2\x81\x29\xdc\xb2\xac\xd5\x3d\x0b\xad\x53\xb2\x16\x6f\x6c\x4f\x2b\xfa\x75\x93\x24\xc2\x3b\x60\x85\x56\xf1\x34\x44\x7b\x0c\xf7\xa7\xe4\x93\x27\xa9\x43\x7c\x06\xf6\x5c\xd9\xac\xe6\xe8\x86\xbb\x1a\x24\x4b\xcd\x91\x68\x84\x75\x9a\xdb\xf4\x53\xff\x9a\xe0\x99\x66\x21\x3a\x29\x4b\x03\x24\x4b\x56\xe6\xd6\xea\xd8\xea\xda\xe2\xf5\x73\xfc\x1b\x07\x17\xce\xa1\xf6\xc9\xec\x7a\x3e\x3b\xc9\x62\xc6\x91\x3e\xbe\xa5\xd1\x16\x5f\xae\x4c\x3a\xac\xda\xd2\xc4\x9f\xed\xe2\x4b\x47\x0a\xbc\x72\xa2\xc2\x45\x8d\x80\x0b\xc3\xb7\x63\xf1\x9b\x1d\x8b\xbe\xdb\x96\x37\x89\x10\x03\x4d\xba\x2f\x4e\xc5\xa9\x9e\x74\x36\xff\x75\x63\x57\x05\x77\x8c\x08\x01\xd2\x47\xfb\x4d\x63\x62\x53\xb7\x80\x55\x7a\x16\x50\xc3\xa6\xe8\xf9\xee\x13\xa5\x6f\x0f\xc7\x57\x9a\x89\xa2\x78\xf9\xb0\xb2\x25\xeb\xc7\x22\x6b\x29\xb0\x5a\xf6\x1a\x7c\x2e\xa4\xa5\x1a\xd3\xab\xab\xf3\xc7\xbc\x42\x78\x3a\x3d\x3b\x7f\x0e\x9c\xca\xcd\x9f\xef\x65\xf8\xac\x5d\x5a\xdc\xf2\xf0\xec\xff\x4b\x1b\xfe\x61\xfb\x4f\x00\x00\x00\xff\xff\xb7\x49\x30\x54\x9c\x16\x00\x00")

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

	info := bindataFileInfo{name: "endpoints/endpoints.tpl", size: 5788, mode: os.FileMode(436), modTime: time.Unix(1448977595, 0)}
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

	info := bindataFileInfo{name: "general.tpl", size: 295, mode: os.FileMode(436), modTime: time.Unix(1448977595, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _restGorillaPatTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x5a\x5f\x6f\xe3\xb8\x11\x7f\xb6\x3f\x05\xab\x62\x0f\x54\xea\x2a\x5b\xa0\xe8\x43\x8a\x14\xd8\xdb\xcd\x5d\xb7\xd8\xec\xba\x71\xd2\x3e\x04\x41\xa1\x48\xb4\xad\x86\x16\x15\x92\x72\x12\xe4\xf2\xdd\x3b\x43\x52\x12\x65\x53\x67\x67\x93\x00\x6d\x17\xc8\xda\xe2\x9f\x1f\x67\xe6\x37\x1c\x0e\x47\x7e\x7c\x24\x39\x9b\x17\x25\x23\x51\x95\x66\x37\xe9\x82\x45\xe4\xe9\x69\x3c\x72\x0f\x04\xfa\x93\xe9\xcd\x02\xdb\xe0\x2b\x2b\x73\xfc\x36\xf6\x66\x15\xab\x4a\x48\xad\xec\xac\x68\x51\xe8\x65\x7d\x9d\x64\x62\x75\xb8\x10\xbf\xbf\x29\xf4\x21\xfe\xc1\xb4\x4a\x14\xa5\x8e\xc6\xa3\xa5\xd6\x95\x96\x69\xa9\x70\x16\x19\x18\xdf\x0e\x38\xc4\xe1\x30\x6b\x21\x6a\x99\x67\xfa\x7e\x63\x02\x34\x9a\xf1\x99\x28\x35\xbb\x47\xf8\x70\x7f\xc5\xe4\xca\xad\xad\x98\x5c\x17\x19\x1b\x00\x72\xbd\xcd\xb2\x03\x83\xb4\x90\xcc\xf4\x0a\x9e\x96\x8b\x44\xc8\xc5\xe1\xfd\x61\xc9\x3c\x39\xa0\x93\x95\x99\xc8\x8b\x72\x71\xf8\x6f\x25\x4a\x1c\x3d\x5f\x19\x01\xb9\x58\xe0\x07\x0e\x6f\x56\x51\x5a\xc2\xcc\xb5\xfb\x0a\x73\x54\x34\x60\x6c\x80\xb4\xfc\x8c\xe7\x75\x99\x19\x72\x66\x28\x0d\x34\xcd\xac\xe8\x8a\x56\xa9\x5e\x2a\xe2\xe9\x9a\x4c\xb1\x65\x42\x1a\x16\x14\x59\xa5\xd5\xa5\x5d\xe9\xaa\x69\x4c\x4e\xdc\x97\x98\xd0\x65\x5a\xe6\x9c\xc9\x3e\x48\x83\x1f\x93\x47\x50\xef\xf0\x90\xac\x53\x59\xa4\xd7\x9c\x29\xa2\x05\xa9\x15\x23\x3c\xd5\x4c\x02\x57\x4c\x5b\x99\x8e\x8e\xc9\xcf\x4c\xfb\x32\x8e\x47\xc6\x76\x5f\xd3\x95\xe9\x8d\xfc\x3e\xd0\xbe\x14\x75\x89\xed\x46\x83\xe4\x2b\x3c\xd1\xd8\xae\xe5\xf4\x87\x8f\xb4\xe6\x9a\xac\x8a\x1c\x04\xbc\x4b\x25\x53\xe3\x11\xc8\x41\x2a\xc9\x2a\x78\xfa\x28\x19\xc8\xd0\x2a\x9a\x9c\xb6\xe3\xc8\x31\x41\x8b\xd1\xa2\x2c\x99\x24\x01\xa5\xb7\x9a\x40\xcd\xd1\x48\x32\x5d\xcb\xd2\x4c\x25\x14\xfd\xcf\x31\x9c\x7c\xb4\x9f\x13\x22\xd9\x6d\xcd\x94\x26\x30\x83\xc9\x79\x9a\xb1\xc7\x27\xb0\x20\x08\x56\x09\xa0\xcf\x6b\x05\xfb\x4b\x89\x7f\x42\xc6\x06\x1b\xf5\xaa\x38\xf4\x2d\x05\xcf\x99\x3c\x22\x69\xf9\xa0\x97\xc0\x09\x79\x10\x35\xb9\x4b\x41\x04\x30\x6c\x2e\xc8\x1d\x38\x21\xd1\x4b\xd4\x4b\x17\xfa\xc1\x4d\xed\xfd\xbb\x66\x73\xb4\x62\x5a\x55\xe8\x34\x38\x2d\xd5\xe9\x75\xaa\x18\x0e\x76\x4a\x18\xd5\x51\x89\x56\xe8\x18\x7a\x81\x13\xf0\x26\xdf\x88\x17\x55\xfe\x66\x46\x7c\xbe\x0d\xc1\xaf\x86\x8d\x68\xb4\x03\x67\x46\xa7\x71\x20\x09\xf5\x9c\xdb\x07\x1c\x3b\xbb\x81\x7b\x36\x12\x00\x32\x98\x6d\x95\xea\x42\x94\x06\x0a\x71\x9a\x50\x93\xfc\xf5\xfc\x7c\x7a\x66\x41\x51\x6c\x34\xd6\xa8\x98\x13\x49\x8e\x8f\x49\x59\x70\xcb\xe1\x48\xa1\x40\x30\xcd\x38\x76\x72\x82\x92\x7d\xc6\x55\xcb\x94\xb7\xfd\x66\xef\x30\x79\xaa\x16\x60\xc1\x68\x55\x28\x85\x2c\x77\x5a\x93\x2e\x7e\xc1\x3f\x04\x04\x3c\xf8\x30\x8f\xd6\x7a\xf8\xf5\xc9\xa8\xc0\x38\x2e\xf7\xc3\xe5\x15\x6e\x9e\xf3\x87\x0a\xf7\xce\xe3\x13\xf6\xdc\x1a\x33\xa0\xfa\x11\x20\xcb\x87\xe8\x2a\xa1\x56\xac\xbf\xe3\x63\xcf\x02\xa6\x1d\x9f\x95\x35\x29\x2a\xee\xf6\x2d\x95\x8d\xaa\xd8\xf1\x9b\x17\x29\x0b\x31\x2f\x99\x55\xc0\x85\x9e\xd3\xc8\xd0\x46\xc4\xb5\x4e\x8b\x12\x0d\xf0\x4e\x59\x20\x42\xdf\xa9\x38\x9a\x90\x36\x36\x18\x91\xe2\x5d\xc6\x80\xff\x20\x1e\x80\x23\xaa\xe4\x23\x17\x8a\xd1\x56\x41\x88\x15\xb9\xd9\x30\xc6\xae\xe0\x85\x62\x6e\x1e\x8b\x7c\xdc\x41\x82\x9c\xa9\xcc\x96\xf4\x36\x4e\x3e\x70\x4e\x19\x7f\x53\xad\x95\x59\xcc\x6a\x7d\x04\x7f\xd1\xc4\x4c\x35\x21\x2f\x99\x41\x7b\xcd\x53\x49\xe3\xfd\x34\x77\x6a\x6a\xc6\xb9\xd5\x0b\xf7\xa3\x93\x9e\xb3\x92\x1e\x80\x32\xe4\x2f\xe4\xbd\x93\xdf\xba\x04\xec\xee\x75\x74\x05\x90\xa6\xfb\xf2\xfd\x55\x0f\xeb\x95\x03\x51\x6d\x63\xc8\xee\x40\x04\xa2\x85\xa3\xd0\x97\x02\xb6\xc6\x70\x0c\x22\xff\x3d\x41\xc8\x28\x66\x47\xd9\x9e\xe3\x81\x50\x1b\x74\xae\x4d\x5e\xd7\x6d\x38\xb3\x90\xc3\xf1\x0c\x06\x73\x34\x12\x0c\xc6\x49\x97\xc6\x95\xa6\xbc\x96\x29\xa7\x31\x6c\xfd\x83\x5e\x88\x68\x04\x30\x53\x5c\x04\xfb\xe5\x17\x72\xe0\x3f\x5b\x89\x5c\x13\x24\x09\x37\x8c\xf6\x30\x26\xe4\x7d\xdc\xee\xbd\xc0\x9a\x30\x09\xe7\xbe\xb6\x4f\x19\x79\x1a\x7b\x78\x7e\x84\x12\x4c\x50\xf0\xce\x81\x60\xea\x9d\x04\x03\x5a\xe7\x68\xc9\x33\x8b\x34\xc9\x43\x25\x85\x16\x99\xe0\x3d\x77\x9b\xba\xc6\xff\x0d\x97\x33\xf6\x6f\x03\xf7\x0b\xbc\x4d\x81\x65\xb2\x25\x59\x27\x54\x03\xc5\xce\x9b\x33\xd8\xae\x24\xec\x75\x47\x0e\xc5\x49\xd8\xc6\xc4\xfb\x0a\x52\xc6\x33\xd7\x4c\xd7\x83\x4e\x1b\xbb\xa8\x8d\x44\x0c\x60\x7d\x65\x77\x2d\x50\xcf\xbd\x26\x64\x1d\x77\x92\xb7\x9a\x74\xcc\x2f\x18\xd8\x01\xc2\x8e\xea\x88\xc7\xc4\x1f\x8f\x5b\x01\x67\xec\x92\x65\x37\x40\x5f\x97\x3b\x8e\x47\xa6\x6d\x0a\x63\x7e\xb4\x61\xeb\xa8\x61\xda\x9b\x67\xb5\x88\x83\x8e\xb1\x95\x22\x3e\xd3\x43\xde\xd8\x45\x76\xfb\x48\xd0\x49\x3c\x2f\x71\xc6\x1e\xad\x4c\x5a\x0e\x46\x49\x20\xa3\x3f\xad\xef\xdb\x84\xc8\x1d\x53\x2b\x3c\x43\xc5\x9d\x5d\xa0\x33\x5e\xbb\x58\xc3\xdd\xbe\xeb\x79\x5e\xb1\xf6\x1d\xb7\x89\x3e\x61\xd2\xad\xa9\xf6\xe5\xfc\xc3\x1c\x6c\xf7\xff\x44\xf9\x6b\x10\xb5\x2f\x4b\xa3\x5d\xf7\x8a\x3e\x4d\x86\xaa\x63\xf8\x67\x1e\xdc\x95\x0e\x2f\xb3\x9f\x3f\x41\x8a\x6d\x49\x2c\xb2\x94\xf3\x07\xd7\x9e\x93\xa3\x22\x87\x44\x8e\xf1\xdc\x8c\xa6\x77\x42\xde\x28\x1b\xc8\xe1\x7a\x48\x30\xa5\xc8\xc9\xc5\xd9\x17\x22\x45\xad\x81\xad\x89\xf9\x82\xd9\x60\xc5\xb2\x62\x5e\x64\xb1\x8d\xef\xfe\x32\xbd\xda\x43\xf2\xc9\xf4\xb8\x0c\xff\x27\xe4\xd2\x5d\x6f\x24\x39\xc0\x91\x89\xeb\x32\xc6\xdf\xa2\x64\xcb\xf6\x20\x2e\x1e\xdb\x09\xc8\x64\xf3\x6d\x1a\x23\x07\x34\x02\x45\xa2\x98\xe0\xc9\x54\x40\xbe\x96\xc1\xe5\x7a\x81\xa7\x58\x86\xf7\xc3\x23\x2f\xea\x81\x03\xe4\x0a\x26\x7d\xc8\x73\x1a\xc1\x9c\x09\xe4\xab\xb1\xf1\x36\xbb\xb8\x37\xb4\xc1\x9f\x31\x6d\x67\x21\x58\x3c\xee\x38\x7a\xf2\x6d\x8c\xa9\x14\xaa\x6f\x9f\x88\xb9\x1b\x10\x88\x77\xf6\x60\x6d\xdc\xd4\x37\x57\x33\xe3\x0d\x0d\x36\x76\xf7\x95\x4d\x9d\xb0\x03\x73\x87\x54\x82\xaf\x2b\xac\x11\xc1\x57\xb8\x13\x98\xd2\xc2\x08\x1b\x66\x5a\x5a\x4b\xff\x04\x77\xb7\x7f\xa4\xbc\x66\x34\x52\xa6\x04\x85\x16\x00\xd7\x6d\x06\x81\xfb\x46\x91\xf5\x5e\xd3\x6f\x97\x33\xf5\x15\x48\xcf\x79\xa1\xa9\x1b\x39\x21\xd1\xc4\x4c\x1e\xa1\x55\xfe\x35\xb1\xeb\xe2\x1a\xc8\x15\xb1\x93\xed\x2e\xb8\x4d\x66\xf0\x64\x26\x7a\x5e\xee\xcb\x5c\xa5\x0b\xff\x76\xe7\x4b\x2f\xe6\x73\xc5\x60\x87\xf3\x62\x55\xe8\x26\xf2\x04\xac\x27\x60\x08\xa9\xc1\x76\x7f\xfa\xa3\xdb\xd7\x42\x05\x94\xb6\x70\x56\x70\x1e\x1a\x60\xd6\x89\x9a\x64\xc0\x60\x78\x36\x31\x6d\xba\x3d\x21\x5c\x11\x2a\x99\xa2\x16\x17\xb0\x38\xc5\x09\x13\xf2\x87\xf7\x13\x02\x72\xfc\xd9\xa6\xb6\xbd\x88\x20\xc0\x01\x84\xee\x22\xc2\x53\x93\x63\x06\x56\xe2\xbf\xb6\x12\xdf\xb5\x12\xc7\xd4\x72\x73\xa5\x2e\x05\xc0\x2b\xab\xe5\x00\xda\x64\xc1\xd6\xb8\xc1\x6e\x71\x7b\x7c\x33\x36\xa2\xd6\x54\x71\xd3\xfa\x05\x0d\x43\x8d\x79\xcc\xc4\x6e\x8b\xdd\x0e\x6d\xa2\xbf\xcd\xbe\x7d\xc5\x2d\x61\x3b\x15\x49\xc9\xf6\x76\xd0\x4b\x88\x4c\x6e\x97\x39\x48\x83\x00\x2a\x0a\x02\xc1\x4d\x64\x70\x50\xe5\x64\xc5\x56\x02\x76\x20\xe8\x5c\x67\x00\xc6\xfc\x9d\xd7\x2c\xf3\xa6\xa1\x0a\x44\x6a\xa4\xe9\xd2\xee\xce\x06\x3f\xf4\x2b\x0b\x76\xc2\xa6\x5a\x90\xc5\x65\x48\x26\x56\x37\x71\xfb\x5a\x11\x25\x95\xc9\x8f\x22\x7f\x40\x4b\xdb\x03\x07\x86\x39\xf1\xa9\x77\x3a\x04\x4d\xec\x0a\x4f\xcf\xb6\x70\x67\x3d\x87\xf0\xa6\xc6\x6b\xea\x4d\xe6\x7a\xf4\x2b\xb5\x26\x7b\xeb\xb6\xe6\x8d\xae\x9a\x9b\x61\x8f\x65\x5b\x69\x09\x1c\xb6\x5e\x72\xdb\x22\xb9\x92\x4e\x1f\xc8\x9c\x6b\xcf\x82\x69\xcb\x00\x78\x6f\xea\xd1\x8e\xfd\x21\x6a\x4c\xe5\xb5\xcb\xa0\xf0\x04\x2b\x80\x1e\x31\x47\x27\x22\x67\x27\xb3\xf3\x79\xdd\xdd\x9b\xd4\x78\xd4\x96\x8f\xb7\x8d\x74\x10\xa8\x28\xa3\xb5\xb8\x58\x24\x53\x57\x30\xc1\xea\xef\x11\x79\xf7\x5b\x48\x5e\x33\x53\xcf\x85\x93\xd0\x56\x84\x9b\x9b\x40\xdc\xad\x71\x19\xb9\x31\xa8\x92\x0f\x0e\x2e\x89\x56\x76\x6b\x50\x50\xac\x8f\xe1\x95\xc3\x3b\x8c\x20\x70\x72\xca\xf4\x52\xe4\xa8\xce\xe5\x95\xd5\xe3\x31\x9a\x7e\x9b\x9d\x47\x4f\xc1\xe1\xd6\xdf\x9c\xa3\xf5\xf8\x0e\xa3\x77\xb5\x6c\x73\xf4\xfb\x4a\x9c\xfe\xb3\xb9\x97\x4e\x36\x2f\xaa\x03\xa2\xee\x00\x33\x10\x93\x7e\xb9\xfc\x7b\x90\x3e\x63\xce\x87\x25\xac\x8d\x1b\x14\x75\x00\x24\xfa\xdd\x46\x65\x2b\x46\x9a\xbb\x75\x9a\x38\xbd\x37\x6f\xbd\x1a\x59\xc7\x5c\x87\x13\x87\xe1\x83\xec\xfd\x7c\xb2\x41\x9e\x37\x3e\x40\x9f\xd9\x65\x43\xf0\xaf\x40\xdf\x33\xe0\xfa\x04\x62\x9e\xf6\x7d\x48\xdb\x04\x9a\xeb\x10\x6d\x11\xc2\x0c\xfa\x6b\xd9\x1a\xdf\x4b\x09\x6c\x50\xe2\x10\x74\x78\xeb\x5d\x6c\x90\xd7\x8e\x0e\x50\x67\x4f\x84\x30\xf6\x2b\x30\xb7\x37\x58\x9f\x37\x2b\xd5\xf7\x20\x0d\x6f\x3c\x57\x72\xdd\x49\x1b\x26\xfd\x2f\x8b\x96\x16\x21\xde\x06\xdd\x6f\xaf\xb9\xb1\x01\xb2\xdc\xb5\x23\x04\xfc\x0a\x5c\xed\x09\xb5\x63\x87\xed\x85\x32\xb8\xbb\xcc\x8d\x6b\x27\x45\x39\xe3\xec\xe5\x3b\xab\x41\x89\x43\xd0\x41\xaa\x3e\x9d\x7c\x39\x39\x3f\xe9\xb3\xd5\x4e\xd8\x8f\xaf\x0e\xff\x15\x18\xdb\x13\x6c\x78\x4f\x58\x80\x81\xc3\xc8\x25\x3a\x90\xe7\x40\x9a\xe3\xbf\x89\x3e\xc3\x7c\xe8\xba\x80\x7b\xb5\x7b\x3f\x05\xd9\x3b\xd6\x1c\x6c\x81\x61\xfb\xcd\x3b\x8e\xa7\x72\xde\xe3\xea\xcc\x8c\x45\x6b\xb9\xe4\x25\xf4\x42\x1e\x3a\x40\x64\x45\x92\x24\x09\xbd\x69\x9f\x62\xb7\xbd\x2b\xfb\xf9\x11\x66\x5c\xc4\x25\x49\x2a\x90\x1b\xf5\xca\x62\xe4\xe3\xd9\xc5\x27\x3f\x39\xeb\x7e\x04\x00\x39\xac\xaf\x45\x53\xa5\x52\x9d\x63\xd9\x37\xf0\x3d\xff\xea\x75\x0c\xad\xd9\xfc\xc6\x22\x67\x2a\x93\x45\x85\xaf\x57\x09\xed\x38\x9c\x6c\x27\xf6\x40\xbb\x9b\xb4\x25\x57\xff\x77\x0d\x9e\x87\x7b\x73\x12\x63\x2a\xea\xec\x09\xe6\x74\x12\x19\xc6\x88\xbd\x7c\x35\x52\x81\x0d\x5c\xaa\x7c\x74\xdc\xb6\x59\xbe\x80\x44\x77\xfb\xf4\xb2\xe8\x2a\x2d\x8b\x8c\xda\xd7\x72\xe0\x2b\xf6\xf7\x18\xee\xa7\x19\xff\x09\x00\x00\xff\xff\xff\x8a\xf3\xa5\x3d\x23\x00\x00")

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

	info := bindataFileInfo{name: "rest/gorilla/pat.tpl", size: 9021, mode: os.FileMode(436), modTime: time.Unix(1448977595, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _storeUpperioTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x58\xcd\x6e\xdb\x38\x10\x3e\xdb\x4f\xc1\x0a\xd8\x82\x4a\x03\xba\x87\xa2\x87\x14\x39\x74\xe3\x76\x11\x20\x6d\xb3\x49\xbb\x97\x20\x28\x64\x89\xb2\x89\x48\xa2\x42\x52\x69\x82\xc0\xef\xbe\x33\x43\x4a\x96\x6d\xb9\x6e\x8a\x5c\x76\x0b\x34\x96\xc8\xe1\xfc\x7e\xf3\x23\x3e\x3e\xb2\x4c\xe6\xaa\x92\x2c\xaa\x93\xf4\x26\x99\xcb\x88\x2d\x97\xe3\x51\x78\x61\xb0\x2f\xce\x6f\xe6\xb8\x06\x8f\xb2\xca\xf0\x69\xdc\x3b\xa5\xca\x5a\x1b\x67\xfd\xa9\xa8\x92\x6e\xb2\x70\xae\x8e\xc6\xa3\xb9\x6e\x4c\x96\xba\x7b\x16\xcd\x95\x5b\x34\x33\x91\xea\x72\x42\x8b\x93\x1b\xe5\x26\xa9\xae\x9c\xbc\x77\x40\x38\xbc\x6f\x9d\x36\xf2\xe7\xbb\x93\xa6\xae\xa5\x51\x1a\xa8\x40\x21\x95\x33\x71\x9a\x89\x53\x7b\xe9\x8c\xaa\x50\xe3\xfe\x51\x9b\xc0\x09\x05\x1c\x44\xd3\xa8\x2c\xea\xd9\x32\x7c\x16\x04\xcb\x2a\xd5\x19\xbc\x4d\x66\x89\x95\x6f\xdf\xa0\x2e\x96\xb6\xad\x17\x18\x18\x8c\xa2\xbc\x24\x33\x0a\x3d\xc7\x1f\x52\x4a\x28\x3d\xc9\x66\xd1\x0e\x9f\x01\x5f\xef\xe6\x71\xde\x54\x29\x53\x95\x72\x3c\x66\x8f\xe3\xd1\x64\xd2\xd2\x90\x81\xac\x36\xfa\x4e\x65\xd2\xb0\x1f\x60\x09\xbe\xdd\x3f\x8c\x47\xb4\x25\xce\xc3\x96\x15\x53\x3a\xf1\x11\x38\x71\xb4\x4b\x7c\x7d\xa8\xa5\xf8\x9c\x94\x12\x5d\x70\xc8\x48\x04\x37\xec\x00\xe3\x22\x2e\xe4\x6d\x23\xad\x8b\x19\xb7\x5e\x86\xb8\xc4\xbf\x87\x4c\x1a\x83\xff\xb5\x21\x45\x46\x46\xba\xc6\x54\xec\x2f\xe9\x36\x59\x12\x3d\x37\xf1\x78\xb4\x8c\xc7\x60\x02\xe8\xbc\x8b\xaa\xd5\xdf\x32\x93\xfc\x60\x83\x24\xde\x03\xbb\xc5\x0c\xa8\x7d\x30\x48\xba\x61\x01\xf9\x52\xcf\x5c\xa2\x2a\x96\x25\x2e\xc1\x18\x8e\x47\xd9\xcc\x93\x1d\x1d\xb3\x80\x1d\xf1\xa5\x96\x15\x37\x87\x2c\x02\xc7\x27\x4d\xe1\x22\x30\x0c\xe0\x80\x54\x2f\x8e\x59\xa5\x8a\x9e\x3b\xc0\xe4\x35\xc6\x56\x5a\xab\x74\xc5\x4e\xa7\x8c\x57\xda\xc9\x23\x56\x49\x99\xb1\x4c\xa7\x4d\x29\x2b\x97\x38\xd8\x04\x76\x00\x8c\x73\x03\x41\xba\x47\xb9\x51\x34\x26\x01\x78\x16\xce\xc1\x4a\x9b\x28\x02\xbc\x10\xec\x3c\x9d\x82\x83\xdf\xb5\x34\x2f\xf0\x14\xa9\xb1\xe2\x04\x4b\x57\x11\x7b\xd5\x92\xbc\x62\xd1\x35\x8b\x3a\xfd\xd6\x40\x94\x00\x00\x5b\xfd\x2d\x1c\x7c\x39\xe8\xbf\x47\xf4\x4d\xc7\x1f\x60\x6d\xc5\x99\x9e\xf3\x68\x57\x68\xd0\x4f\x81\xab\x07\xc1\x30\x02\xac\x34\x77\x10\xff\xb9\xac\xc0\xdb\x29\x3b\xf9\x76\x31\x65\xb9\x36\xcc\x01\xe5\xd6\x11\x8f\x25\xa0\x4c\x1c\xb8\x71\xf6\xe0\x5d\xc3\x4e\xce\x4e\x99\xd3\xba\x18\x0f\x1e\x0a\x72\x9c\x69\x52\x87\x4e\x9a\xce\x58\xf8\x97\xcd\xc4\xb4\x0b\xfd\xca\x75\x3e\x89\x83\xd6\x27\x46\x82\x30\x96\x6c\xb1\x85\xac\x64\x6e\x21\x3b\xf0\x1c\x32\x9d\xd3\x42\x9d\x18\x88\xad\xc7\xed\x4e\x34\xc6\x81\x31\x1f\x8f\xa0\xd0\x65\x21\xd5\x4e\xe0\xd1\x02\x02\xeb\xf0\xfe\xa1\x72\xca\x3d\x9c\x3b\x40\x2c\xdf\x46\xef\x5c\x3a\x96\xea\xa2\x90\x29\xe2\x08\x19\x15\x45\x07\x5f\x0b\xcc\x8a\x82\xff\x0a\x58\x93\xba\x2e\x1e\x20\x03\xab\x4c\x97\x0c\xab\x5f\xf0\x00\xf8\xb4\x7d\x52\xd9\xce\x32\x88\xf4\x98\x2e\xf0\x2b\x3e\xcb\x1f\xff\xbc\x41\x99\x12\x97\x64\x2d\x38\x58\xbf\x32\x7e\xb9\xc4\x2d\x81\x0e\x01\x2e\xad\x1f\x8f\x83\x10\x2b\xbe\x1a\x55\x5e\xa8\xf9\xc2\x71\x5f\x51\xc5\xb7\x8b\xb3\x0f\xa1\xc8\x0a\x7a\x90\x5f\xb5\x17\xcd\x41\xdc\xd5\xd1\x75\x0c\x79\x79\x8c\x48\xeb\xd7\x6b\x30\xe9\x53\x62\xec\x22\x29\x28\x1e\xca\xc9\xf2\x10\x55\xaf\x35\x64\xe3\xac\x90\x44\xc1\x6f\x1b\x95\xde\x30\x8c\x37\xc2\x2d\x64\x3b\x16\x24\xa0\x28\x7d\x3d\x05\x7c\x04\x46\xd2\x78\x47\x96\x18\xe6\x9b\xd6\xb8\xb5\xfd\x77\xb8\x81\xde\x95\xb5\x0f\xc2\x31\x50\xb7\xfb\xd3\x3f\xd1\x2b\x03\xa1\xe8\x62\x01\xc1\xe8\xe2\x91\x65\xa4\xb8\xa4\xe8\x63\x14\xfa\x51\xde\x11\x86\xef\xad\x50\xa4\x15\xef\xc1\x9c\x2a\xe3\xb2\x0e\xae\x29\x2c\xa5\xcf\x48\x65\x1d\x40\x06\xe9\x42\xc7\xda\xd6\x93\x92\x3d\xe7\xd1\x07\x44\x20\x4b\x11\xbb\x28\x79\x13\xdb\x47\xec\x0f\x1b\x91\x08\x41\x94\x3c\x46\xb3\xbd\x62\x01\xd2\xb8\x7c\x0a\xad\xdd\x54\x49\xb1\x81\xc5\x41\xd3\xfa\xea\x77\x60\x45\xf7\xdc\x48\xf2\xcd\xca\x53\x4f\x85\x5d\x58\x40\x3a\x58\xe0\x80\x5f\xae\x2a\xf7\xf6\x4d\xbc\x81\xa7\xb5\x2a\x76\x29\x13\x93\x2e\x86\xea\x01\x94\x23\xe5\x2c\xc3\x74\x56\x18\x2a\x6e\xe3\xbd\x45\xc0\x73\x83\x22\x70\x1b\xdc\xf3\x77\x23\xcd\x43\x1c\x5e\x2e\xa4\x85\x86\x43\xe9\x1e\xda\x6d\xdb\x94\x20\xd1\xfc\x26\x47\x09\x30\x1b\x70\x03\x45\x14\xf0\xe8\x57\xb7\xba\xf5\x40\xb5\xd8\x55\x2e\xf6\x80\xd4\x33\x83\x05\xa3\xe4\x5d\xf0\xbc\x02\xd9\x58\x8c\xd5\x9d\xac\xd8\x2d\x5a\xb0\xf2\x82\x25\x49\x50\xd6\xfa\x3d\x95\xea\x1c\xbf\xc5\x86\xe6\x1f\xe3\x20\xd8\x53\x1e\xaf\x89\xb6\x2d\xa8\x3f\x2a\x80\x2a\x12\x2e\x3d\x22\x06\xf7\x89\x43\xdc\x53\x15\xb3\xc9\xc2\x04\x4a\x65\xac\x82\x5c\x2f\xa9\xed\x52\x45\x48\xaa\x87\x71\xcb\x02\xfe\x8a\x4b\xa0\xe3\xad\x92\xf4\x72\x1b\x0b\x21\xe2\xc0\x6a\x01\x25\xb2\xc0\x02\x3f\xc7\xfe\x40\x0a\x93\x0d\x5f\xf2\xdc\x4a\x9c\xd0\xc0\x67\xaf\xfb\x6a\x11\xcf\x1b\x55\x43\xb5\xaa\x1c\x5f\xa3\x8d\xbd\x8e\x2b\x26\x67\xaa\x54\xbb\x78\xf8\xbd\x15\x93\x40\x1b\x77\x76\x76\x59\x04\xaa\x7a\x9c\x7e\x81\xf6\xee\x57\x2d\xa5\x48\xae\x8c\x75\xdb\xa0\x05\x67\xa4\x0b\xf9\x44\xd4\x02\x6f\xec\x5b\xbf\xd5\xb4\x3a\xe4\x18\x42\xaa\x65\xb9\x81\xae\xb3\x9a\xc0\x0a\xc4\xc9\xcb\xab\xeb\x4d\xd1\x8f\x4b\x4c\x92\xa3\xb6\x8c\x40\x02\x50\xaa\xf0\x58\x5c\xb6\x20\x4a\xe3\x30\xdc\x34\x65\xdd\xb1\x07\x9f\x41\x55\xd7\x0a\x2b\x0e\x76\xe8\x32\xa9\xd9\x24\x4c\x03\xe3\xb6\x32\x89\x90\x87\x10\xed\xf7\x90\x05\xc5\xaf\x74\x4d\x20\x80\x91\x0e\x9a\x47\x53\x41\x51\x35\x12\xbf\x72\xe8\x58\x01\xf3\xe2\x41\x11\x23\x8a\x7d\x20\xb7\xca\xdf\x67\xed\x3e\xe2\xb1\x81\x56\x0c\x1d\x6a\xee\xe7\x8a\xbb\xa4\x68\x24\x6a\xec\xd3\x8a\x4c\x20\x9a\x50\xf3\x7c\x40\x5b\x77\x66\x9e\x7e\x3c\xe2\x07\xa1\xfe\xad\xbb\x2f\x06\x7d\x18\xaa\x75\xf5\xfa\xba\xab\x26\x60\x59\x40\xcb\xb7\x3a\xc3\x29\x67\x0b\x1e\x30\xbb\x3e\x09\x19\x9e\xcf\x6f\x83\xe3\xb9\x26\x1a\xe4\x33\xeb\x15\x21\x1a\x71\xc1\xb1\x38\x08\x52\xbf\x40\xa9\x7e\xf2\x3a\x64\xdf\xa9\x1b\x62\x5a\x7d\x4a\x6a\x4e\x53\xab\xed\x1a\x24\xd5\x14\xa8\xab\x68\x07\xd5\x96\x38\xfe\x1f\x4c\x18\x8d\x0f\x37\xea\xee\x0b\x40\xd6\x15\xf2\x36\x27\xb0\xec\x84\x68\xd2\x74\xb0\x6f\x24\x20\x96\xcf\x32\x12\x2c\x37\x3a\xee\x54\x16\xf2\x39\xb0\xe9\xf9\x6c\x62\xf3\xbf\x8e\x43\x23\x4b\x7d\xb7\x3f\x94\x17\x44\x36\xac\xf4\x7a\x20\x33\x74\xd3\xf3\x06\xb2\x57\x68\xa0\xbc\xea\xd4\x17\x01\x96\xe0\x33\xe2\xb0\x04\xe5\x60\x66\xc0\x2c\x49\xaa\x76\x98\xdb\x17\xce\x1e\x27\x1e\x6f\x56\x17\xb4\x2b\xc8\xde\xfa\x9a\x85\x4e\xb2\xa5\xcb\x99\x82\x5a\xfa\x53\x7d\x58\x01\x24\x4f\x51\x0a\x59\x6e\x28\x86\x4b\x1b\xca\x0d\x36\x3a\xaf\xde\x19\x54\x7d\x55\xd9\x1a\x40\x48\xe1\x85\xbe\x32\x87\x92\xa1\xf3\xa7\x2a\x05\x8c\x78\x5d\x0c\x69\x12\x33\x1a\x78\x51\x21\x49\x8d\xb7\x2e\xa0\x77\x6c\xeb\xd4\x7d\xca\x7b\x7a\x4e\x2d\x4e\x16\x71\x7b\xbf\x83\x09\x11\xa6\x0d\x52\x15\x6f\x73\xda\x6b\xae\x7e\x22\xed\xfd\x26\xa6\xc4\x62\x1c\x8f\x30\x42\x7b\x7b\x74\x73\xa2\x0d\xc9\x85\x82\x86\x13\x15\xf3\x74\xda\xe7\x80\x77\x5f\xf4\xba\x5c\x0e\x5f\xe1\x6c\x7c\xe2\xe8\xaa\xc2\x83\x90\x08\x2b\x01\xac\x63\xe1\x93\x01\xcb\xeb\xef\x17\x36\x10\x87\x17\x2a\x7e\x3e\x2b\xa5\xb5\x78\x9f\x4a\x6d\xa1\xbd\x37\x82\xcf\xee\xbd\xc1\xd5\x73\x5e\xda\x79\xf8\x86\x26\xdf\x00\x4f\x71\x6e\x70\x58\xb4\x62\x75\xab\xf1\x8a\x01\x59\xbc\x12\x9d\x3f\x8b\xec\xbc\x27\xfc\x90\xdd\x31\x98\x96\x69\xd4\xca\x93\x54\x3e\x2e\x49\x1d\x7f\x4f\x94\x97\x4e\x5c\xd6\xa8\x15\x1d\x01\x5a\x1c\xac\x3b\x00\x15\xda\xca\xb5\xeb\x94\x4e\x0f\xb7\x48\x06\xe6\x56\x65\x59\x63\x71\x0a\xdf\x0b\x29\xe4\x0c\x98\x22\xe8\xac\xd4\xd9\xba\x08\xbd\xf4\x2d\xc1\x53\xaf\xee\xae\x02\x8a\xfc\x32\x2a\xdb\xbb\xb1\xfd\x37\x00\x00\xff\xff\xae\x7d\x93\x52\x1b\x17\x00\x00")

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

	info := bindataFileInfo{name: "store/upperio.tpl", size: 5915, mode: os.FileMode(436), modTime: time.Unix(1448980824, 0)}
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

