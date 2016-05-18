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

var _endpointsEndpointsTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x58\x4d\x53\xe3\x46\x13\x3e\xcb\xbf\x62\x56\x6f\xf1\x96\x44\x39\x72\xce\x4e\x71\x48\xc1\xe6\xa3\x42\x80\x2c\xec\x89\xe5\x20\xa4\x96\xad\x42\x9e\x11\x33\x23\xc0\xe5\xf2\x7f\x4f\x77\xcf\x48\x96\x6d\x16\xc2\x52\x4e\x2d\x29\x0e\x80\xd4\xd3\xdd\xd3\x1f\xcf\xd3\x9a\x61\xb1\x10\x39\x14\xa5\x04\x11\xd6\x69\x76\x93\x4e\x20\x14\xcb\xe5\x20\xf0\x2f\x02\xd7\x93\xb3\x9b\x09\xc9\xf0\x11\x64\x4e\x4f\x83\x9e\x55\x39\xab\x95\xb6\xc6\x59\x85\x93\xd2\x4e\x9b\xeb\x24\x53\xb3\xd1\x44\xfd\x70\x53\xda\x11\xfd\xa0\x59\xad\x4a\x69\xc3\x41\x30\x51\x8d\xce\x33\xfb\x20\xd6\x55\x51\xc8\x9a\x99\x92\x16\x1e\x48\x71\x6a\x6d\x6d\x40\xdf\x95\x19\x7c\x45\xd7\xaf\x8e\x48\x33\xdc\xdc\xbb\x53\xb2\x4a\x03\xaf\xaa\x2a\x95\x93\x44\xe9\xc9\xe8\x61\x24\xa1\xb7\x15\x2e\x16\x33\xfc\xfb\x78\x7e\x99\xca\x5d\x49\x06\xa3\x11\x57\xe3\x9c\x3c\x26\x27\xe9\x0c\x50\xfa\xd1\x67\x66\x84\x06\xdb\x68\x29\x0e\x3f\x7f\x3a\x12\xd0\x49\x0b\xa5\xb7\x8d\x06\x45\x23\xb3\x27\x7c\x45\x52\x35\x72\x28\xe8\x77\x2d\x8c\xd5\xa5\x9c\xc4\x22\x5a\x39\x9d\xa5\xf5\xa5\x13\x5f\xb5\xc2\xa4\x35\x8e\xc5\x02\x13\xc2\x50\xef\x52\x5d\xa6\xd7\x15\x18\x61\x95\x68\x0c\x88\x2a\xb5\xa0\x07\x41\x5a\x55\x2a\xfb\x28\x6d\x69\xe7\xc7\xa5\xb1\x62\x7c\x20\x28\x9c\x28\x16\xfb\x97\x57\x14\xd3\xc5\xbc\xa6\x68\xc4\xa2\x4d\xe9\xff\x6b\xf2\xc5\x52\x60\xa3\xb9\xac\x7f\xc0\x9c\xcc\xbb\x44\xe8\x9d\x50\x30\x01\xcb\xef\x9d\x6f\xea\xb7\x2f\x77\x72\xe8\xfe\x62\x42\x46\xec\x6f\xd5\x60\x28\x40\x6b\xfa\x51\x9a\x32\x09\x02\x9d\xde\x3b\x19\xfa\xe2\x4d\x93\x5f\xc1\x92\xc3\xa1\x68\x63\x88\x51\xad\x2c\x58\xe9\xc3\x81\x90\x65\xc5\x86\x81\x8b\x1e\x9f\xb0\x75\x41\x60\x86\x42\xdd\x90\x13\x74\x98\x44\xdb\x1b\x7b\x27\x1f\x50\x89\xad\xc9\x1b\x46\x3f\xc3\xca\x52\x30\x45\x14\xae\x76\xff\x12\x6e\x99\x7f\x09\x63\x91\x2b\x2c\xb6\x54\xb6\xad\xdb\xf6\x26\x61\xbc\x11\xd8\xea\x65\xe9\xba\xc6\x9b\xf4\xf0\x33\x05\x0d\xbc\x70\x71\x7a\x74\x3a\xc6\xce\xcf\xc5\x34\xbd\x03\x21\xe1\x9e\x90\xd1\x64\x96\xda\xcb\x56\x83\x60\x65\x76\x80\x9a\x37\x10\x3d\x09\x94\x41\xcf\xe0\x32\xcc\x34\x20\x3e\xc2\x2b\xf1\xf5\x96\x61\x1f\x04\x2a\x83\x2e\xd2\x0c\x16\x4b\xec\xa0\xc6\x84\x7b\x92\x8d\xe6\x61\x76\x18\x38\x82\xa1\x75\x84\xba\xc8\x87\x59\x6a\x4b\x45\xe9\x73\x4f\xdb\x79\x90\xfc\x76\x71\x71\xf6\x09\x6e\x1b\x30\xdc\x5e\xdf\x0f\x6c\x42\xaf\xa5\x66\x0d\x08\xdc\x98\xdf\x69\x7b\x99\x56\xed\x72\x72\x8e\x53\x01\xf4\x9f\x66\x82\x99\x84\xb3\xd2\x18\xcc\x1e\x3b\xc2\x8e\x31\x00\xb1\x9a\x30\x6d\x93\xc9\x6c\x1b\x30\x3e\x74\x5f\x5a\xc2\x8f\xdf\xbc\x45\x77\x2f\xca\x4d\xe8\xbd\x34\x4e\x82\xd9\x79\x8d\x7d\xb2\x88\x33\xae\x9f\x50\xd7\x36\x2d\x25\xc5\xbe\x67\x3c\x2a\xa2\x3d\x13\x87\x2b\xd8\x73\x40\xf1\x93\x69\x04\x01\x4e\x31\xd0\xc2\x24\x87\x95\x32\x10\xc5\x3e\x31\xd7\x6c\x84\x19\x4d\x81\x41\xe7\x20\x39\x64\x79\x84\x79\xa0\xf3\x1d\xe4\x46\x0a\x81\x4f\x90\x63\x70\xf9\x8d\xc5\xde\xff\xee\x86\x3e\x1c\x7e\x09\xdd\xf4\xe3\x14\xdd\x1e\x51\xec\x43\x7a\xae\x69\x20\x69\x62\x63\xc7\x4d\xad\xa4\x01\xa6\x98\x23\x44\xc7\x85\x1e\x64\x39\x25\x1e\xb4\xe3\xcd\x39\xb7\x0f\x7e\x4e\x78\x41\xbc\x1c\xae\x73\xd6\x91\xb6\xc7\x21\x94\xeb\x12\xee\x9e\x63\xd1\x0a\x8c\x2f\xe2\x92\x41\x76\xf0\x00\x73\xe6\x18\x5c\xef\x13\x99\x78\xea\x50\x85\x58\x8b\xb4\x93\xbf\x1a\xd0\xf3\x77\x1e\x7e\xb7\x3c\xe4\x6f\x31\x31\x71\x06\x33\xa5\xe7\x7c\x5a\xe8\xbe\xdb\xc4\xcb\x8a\xa2\xdf\xf8\x62\x77\xd6\x2d\xdc\x7a\x04\x3e\x87\x54\x67\xd3\xe8\x36\x4e\x7e\xae\xaa\x08\xaa\xdd\xd5\xc5\xf0\x4e\x1d\x7f\x4d\x38\x64\x72\x77\xac\xfd\x36\xaa\x62\xac\x26\x39\x06\x49\xa1\x13\xe2\x7e\xec\x7f\x8c\x7b\xd1\x9e\x28\xfb\x0b\xee\x94\x6f\xbb\xfd\xa7\x64\x87\xea\x39\x36\x57\x58\xeb\x77\x26\xbf\x33\x79\xd7\x4c\x46\xb8\x34\x15\x9e\xd9\xc6\x6b\x04\xc6\x85\x0c\x91\x6a\xbb\x22\x79\x35\x04\x1f\x4a\xa3\xdd\x95\x89\x77\x7d\x05\xaf\xdd\x5a\x1b\xed\x5b\x18\x43\xee\x68\x2d\x01\x72\x3a\x4b\x17\xe5\x83\x50\xb8\x65\x51\xa9\x7b\xe1\x91\x53\x88\x1a\xef\xc2\xeb\x1d\x7d\xd9\xa0\x09\xf0\x76\x3d\x41\xad\x70\xec\xb3\x3d\x81\xfb\x33\xf2\x19\xc5\x09\x47\x7c\x0e\xf6\x42\xd9\xb4\x8a\xd0\x4d\xc4\x3d\x88\x57\x2b\xc7\xe5\xac\xb4\xbc\x72\x4b\x17\x10\xf7\x1a\xe3\x89\xa8\x13\x9d\x16\x85\x01\x92\xc5\xc3\x47\x2e\x16\xbd\xa9\xd6\xd4\xf9\xf3\xe7\xfc\x7f\x7b\xae\x61\x8b\x3a\xc9\x59\x3a\xaf\x54\x9a\x33\xfe\xf1\x36\x8e\x62\x4e\x10\x63\xcb\x4d\x47\xb7\xf7\x11\xf8\xfd\x8d\x40\x87\xac\xd5\xa5\xc2\xa7\x40\x53\xed\x33\x2f\x45\xd4\x50\x3a\xc5\xff\xb4\xb3\x4b\x05\x9f\x28\x7c\x7e\xf4\xfd\xfe\xa6\x91\xb0\xab\xfb\xc2\x26\x15\x73\xa8\x60\x57\x54\x7c\xf5\xe1\xd2\xa1\xe3\x96\xe8\x49\xf3\xaf\xcc\x5f\xc7\xef\x77\x2e\xbf\x29\x2e\x17\x25\xb6\xcb\x4e\xc1\x15\x43\x5a\x82\x00\xbd\x32\x0c\xde\xe6\x65\xc3\xb1\xed\xd1\xf1\x74\xc4\x4b\x3c\x9e\xfe\x2b\xb3\x69\xeb\x7a\xc3\xe6\xfe\xd9\xfd\x83\xdb\xff\xaf\xfb\xef\x00\x00\x00\xff\xff\x0d\xe3\xf5\xf7\x01\x18\x00\x00")

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

	info := bindataFileInfo{name: "endpoints/endpoints.tpl", size: 6145, mode: os.FileMode(436), modTime: time.Unix(1461687162, 0)}
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

	info := bindataFileInfo{name: "general.tpl", size: 295, mode: os.FileMode(436), modTime: time.Unix(1460882529, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _restGorillaPatTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x5a\x5f\x6f\xe3\xb8\x11\x7f\xb6\x3f\x05\x4f\xc0\x15\xd2\x56\x55\xf6\x80\xa2\x0f\x2e\xb6\xc0\xde\x6e\xee\xba\xed\xfe\x71\xe3\xa4\xf7\x10\x04\x3d\x46\xa2\x6d\x35\xb2\xa8\x25\x29\x27\x41\x2e\xdf\xbd\x33\x43\x4a\xa2\x6c\x39\x76\x36\xd9\x5e\x7b\x0b\x64\x6d\x93\x9c\xe1\xfc\xfd\x71\x38\xd2\xdd\x1d\xcb\xc4\x3c\x2f\x05\x0b\x2a\x9e\x5e\xf1\x85\x08\xd8\xfd\xfd\x78\xe4\x7e\x30\x98\x4f\xa6\x57\x0b\x1c\x83\xaf\xa2\xcc\xf0\xdb\xd8\xa3\xca\x57\x95\x54\x46\x5b\xaa\x60\x91\x9b\x65\x7d\x99\xa4\x72\x75\xb4\x90\x7f\xb8\xca\xcd\x11\xfe\x01\x59\x25\xf3\xd2\x04\xe3\xd1\xd2\x98\xca\x28\x5e\x6a\xa4\x62\x3b\xd6\xb7\x0b\x8e\x70\x39\x50\x2d\x64\xad\xb2\xd4\xdc\x6c\x10\xc0\x20\xad\x4f\x65\x69\xc4\x0d\xb2\x1f\x9e\xaf\x84\x5a\xb9\xbd\xb5\x50\xeb\x3c\x15\x3b\x18\xb9\xd9\x66\xdb\x1d\x8b\x8c\x54\x82\x66\x65\xc1\xcb\x45\x22\xd5\xe2\xe8\xe6\xa8\x14\x9e\x1c\x30\x29\xca\x54\x66\x79\xb9\x38\xfa\xb7\x96\x25\xae\x9e\xaf\x48\xc0\x42\x2e\xf0\x03\x97\x37\xbb\x68\xa3\x80\x72\xed\xbe\x02\x8d\x0e\x7c\x63\xfb\xd6\x06\x9e\xd6\x41\xe3\x79\x5d\xa6\xe4\x9d\x19\x8a\x03\x43\x33\x2b\xbb\x0e\x2b\x6e\x96\x9a\x79\xca\x26\x53\x1c\x89\x59\xe3\x06\xcd\x56\xbc\x3a\xb7\x5b\x5d\x34\x83\xc9\xb1\xfb\x12\xb1\x70\xc9\xcb\xac\x10\xaa\xcf\xa4\xe1\x1f\xb1\x3b\xd0\xef\xe8\x88\xad\xb9\xca\xf9\x65\x21\x34\x33\x92\xd5\x5a\xb0\x82\x1b\xa1\xc6\xa3\x52\xd6\x25\x9b\xbc\x62\x24\x47\xf2\x11\x7e\x85\xd1\x78\x44\x56\xfb\xbb\xb8\xc5\x99\x56\x6c\xfc\x8d\x71\xb3\x10\xc6\xaa\x01\x93\xa8\x58\x88\xbe\x76\xd6\x4c\xde\xd8\x4f\x90\x4b\xb3\x17\xbe\xc6\xa0\x91\x52\xf8\x27\x15\x0a\x35\x1a\x29\x7e\x6d\xc7\x80\x0d\xed\x97\xfc\x28\x0c\xf2\x8a\x59\xb3\x3d\x48\x32\xca\xe7\xb4\xe8\x9b\x57\xac\xcc\x0b\x22\x1c\x29\x61\x6a\x55\xc2\x37\xb0\xec\x68\x04\xb6\x92\x57\xc8\x04\x18\x26\x61\x6f\x4f\x47\xff\x0d\xcc\x13\x21\x32\x02\x99\x57\x60\x3f\x94\x63\x1e\xfe\xdc\x6d\xbc\xa1\x66\xc4\x32\x09\xc6\x2a\xa5\x61\x76\x3b\x4f\x1b\xbb\xe2\xe7\x68\x43\x96\xee\xc7\xbd\xb5\xb9\x8b\x03\xf8\xe0\x75\x61\xd8\x2a\xcf\xc0\x51\xd7\x5c\x09\x3d\x1e\x81\x3f\x58\xa5\x44\x05\xbf\xde\x28\x01\xbe\x68\x1d\x9e\x7c\x68\xd7\x31\x67\xe0\xbc\x2c\x85\x62\x03\xce\xdf\x1a\xb2\x96\xb5\x02\x53\xd0\x0d\x39\x27\x06\x95\x3e\xd7\x42\x1b\x06\x14\x42\xcd\x79\x2a\xee\x40\xe1\x10\x04\xab\x24\xc4\xb1\x37\xba\xe5\x35\xd4\xab\x2a\x60\x6e\x29\x8b\x4c\xa8\x09\xe3\xe5\xad\x59\x42\x6c\xb2\x5b\x59\xb3\x6b\x0e\x22\x40\x80\x65\x92\x5d\x43\x36\x32\xb3\x44\xbd\x4c\x6e\x6e\x1d\x69\xef\xdf\xa5\x98\xa3\x9f\x78\x55\x61\xf6\x20\x19\x37\xfc\x92\x6b\x81\x8b\xfd\x68\x3e\x2e\x61\x61\xea\x0c\x15\x3a\xe1\x3d\xf3\x33\x32\x90\x8d\x1d\x6f\xf6\xde\x7a\xc2\x33\xf5\x59\x95\x7d\x35\x53\x3f\xde\xd2\x90\x85\xbb\x4d\x8d\xda\xe9\x13\xf1\x99\x22\xdb\x32\x81\xe8\xf6\xad\x72\xd2\x68\xea\x6c\x0b\x79\xd9\xec\x0f\x7c\xc1\x62\x2b\x6e\x72\x59\x92\x99\x90\x4b\x83\xcb\xc9\x5f\x4f\x4f\xa7\x8e\x18\x85\x26\x43\x42\x9a\x40\x6e\x78\x49\x36\xd2\xbd\xdc\xa4\x84\x79\x87\xd2\x96\xbc\x68\xe7\x09\x67\x84\xfa\xa0\x17\x60\xbf\x60\x95\x6b\x8d\x91\xd0\xe9\xcc\x3a\xb0\x6f\xd3\x0f\xe9\xe8\x67\x9b\x3a\x36\x8f\x47\xa2\xc0\xed\x7e\x77\x7e\x81\x89\x76\x7a\x5b\x61\x06\xdf\x61\x5a\x8d\xc8\x08\x68\x8c\xe4\x1f\xb5\x50\xb7\xbe\xc2\x24\x1d\xd9\xaa\x05\x93\x06\x9f\x7c\xdd\x36\x21\xe4\xf1\xda\x21\x6a\xcc\x2a\x40\x61\x33\x0f\x03\xf2\x12\x93\x97\x86\xe7\x25\x6a\xfc\xad\xb6\x8c\x58\xf8\xad\x8e\x82\x0e\xc0\x48\xa6\x68\x9f\xf2\xf0\x1f\x60\x04\x84\x9d\x4e\xde\x14\x52\x8b\xb0\x75\x29\xe0\x47\x46\x49\x04\xf1\xbb\xce\x65\xad\xad\x41\x21\xf8\xe4\x9c\xc6\xf3\xac\xc3\x35\x0d\xf2\x72\x95\x2e\xc3\xcf\x51\xf2\xba\x28\x42\x51\x7c\x55\xed\x35\x6d\x66\xb5\x9f\xc0\x5f\x10\x13\x29\x9d\x28\xc9\x0c\xc6\xeb\x82\xab\x30\x3a\xcc\x04\x4e\x5f\x23\x8a\xc2\xea\x85\x69\xe8\xa4\x2f\x44\x19\xbe\x00\x65\xd8\x5f\xd8\xcb\x46\x7e\x8c\x85\x69\x63\x13\x08\x1a\x5a\x70\xfe\xf2\xa2\xc7\x4d\x58\xd8\x60\x7c\xa1\x84\x58\x91\xd5\x4a\x1b\x46\x53\x7e\x5b\x48\x9e\x59\x90\x6a\x6d\x4b\x53\xc7\x2d\x5c\x0d\x20\x90\xc5\x8f\xb0\xb7\x7d\xdc\x63\xd9\xba\xee\x99\x41\xb2\xb6\xc8\xb5\x01\x92\xdb\xf0\x87\xb2\x0c\x63\xdf\xfb\x1c\x52\x72\x37\xf2\xb1\xff\x1d\xe8\x23\xc5\xec\x2a\x3b\xf3\x6a\x07\xc0\x0f\xc6\xf6\x66\x58\xad\xa1\x7e\xb2\x20\x6a\x59\x26\xa1\x57\x50\xf9\xb2\xe1\xe2\x02\x8d\x04\x8b\x91\xe8\x9c\x22\x79\x5a\xd4\x8a\x17\x61\x74\x01\xe8\xdb\x83\xa6\x46\x00\x22\x71\xc8\xf9\xcb\x2f\xec\x85\xff\xdb\x4a\xe4\x86\xa0\x90\xbb\x12\x61\x8f\x47\xcc\x5e\x46\x2d\x06\x0c\xec\x09\x44\x48\xfb\xdc\x31\x45\xf2\x34\xf6\xf0\xe2\x08\x25\x88\x51\xf0\x2e\x80\x80\xf4\x5a\x81\x01\x6d\x70\xb4\xce\xa3\x4d\x9a\xc2\xa6\x52\xd2\xc8\x54\x16\xbd\x70\x9b\xba\xc1\xff\x8f\x90\x23\xfb\xb7\x27\xc8\x13\xa2\x4d\x83\x65\xd2\x25\x5b\x27\xa1\x01\x17\xbb\x68\x4e\x21\x5d\xd9\x70\xd4\x4d\x1c\x17\x27\x61\x0b\xc9\x37\x15\x94\xf5\x27\x6e\x38\x5c\xef\x0c\xda\xc8\x9d\x1e\xe8\x88\x1d\xbc\x3e\x8a\xeb\x96\x51\x2f\xbc\x62\xb6\x8e\x3a\xc9\xfd\x4a\xd6\x79\x7e\x21\xc0\x0e\x00\x3b\xba\x73\x3c\xde\xce\xf0\x98\x07\x24\x4d\x97\x22\xbd\x02\xf7\x75\x75\xed\x78\x44\x63\x53\x58\xf3\xbd\x85\xad\x49\xe3\x69\x8f\xce\x6a\x11\x0d\x06\xc6\x56\xf9\xfa\xc8\x08\xf9\xba\xa8\x34\x5a\xd1\x55\x09\x54\xc1\x2b\xc3\x87\xfa\xa6\x2d\x31\xdc\xd9\xb6\xc2\x83\x57\x5e\xdb\xd0\xe9\x54\xee\x87\xd1\x60\x1c\x79\x81\x64\xd1\xe0\xe1\xd2\x96\xd6\x0c\x3b\xca\xaa\x77\xa8\x9f\x5e\xcf\x41\xdf\xdf\x92\x9b\xc8\x76\x7b\x53\xf9\x00\x1f\x3c\x8b\xc3\x9b\x14\x3b\x74\x3f\x2f\x79\xd7\x3e\xbe\x6c\xb9\x9c\xdc\xfe\x0a\xfe\xe1\xb5\x17\x90\x18\x7b\x0d\x64\x4e\x2c\xf9\x35\xad\x18\x8f\xec\xb0\xeb\x05\xbc\x7b\xeb\x6e\x13\x0f\x18\x9d\xd1\x05\xa3\xbd\x59\xb0\xe6\xae\xc5\x86\x2e\x1e\x5b\xf9\x91\x67\x74\xce\x26\x67\x27\xef\x6d\xa1\x1e\x46\x74\xb5\x0e\x26\x79\x16\x44\x0c\x8f\x92\x1c\xea\xbb\x74\xc9\xcb\x05\x1e\x3b\x29\x5e\x36\x27\x1e\x4c\x81\x20\x99\x06\xa2\xd7\x59\x16\x06\x40\x13\x43\x7d\x1b\x51\xa8\x59\x29\xa0\xd0\x1b\x90\x83\x8c\xe9\xbe\x4f\x98\xa2\x32\x94\xb6\x9f\x74\x9c\x1b\x71\x66\xc2\xd8\x4d\x70\xef\x28\x1e\xb8\xb7\x5b\x9b\xfd\x6d\xf6\xe9\xa3\xad\x04\x1f\x6b\x31\x7b\xea\xda\x6e\x41\x7b\xca\x6f\x18\x0a\x2c\xc1\x21\x6c\x52\x7b\x17\x75\xa7\xb4\x23\x04\x1d\xfb\xf7\x1f\xbb\xde\x79\xd8\x59\x62\x8c\x98\x9f\xa2\x68\xd8\xb0\x42\x05\xdf\xd2\xbc\x0a\x55\xf2\xbd\xcc\xa8\x71\x62\x03\x14\x96\x25\x76\xce\x49\x16\x6d\xf5\x29\xba\x50\xea\xf5\xfb\x1c\x99\xd3\xec\x07\x4c\xf3\x7c\x55\x15\x54\x48\x73\x3f\xca\x5a\xe9\x6c\x84\x11\x16\xe5\x29\x28\x78\xeb\xc6\xc1\xc9\x10\x1a\xf3\x5c\x14\x19\xad\x0e\xaf\xa5\xba\xd2\xae\xf8\xe6\x86\x61\x35\x9b\x31\x88\x1a\xa6\x64\x6d\x00\x74\x62\xfa\x82\x17\xa2\x4a\xa4\xf9\x3c\x4f\x23\x5b\x5a\xf8\xdb\xec\x93\xf5\xcb\xe2\xfc\x21\xf8\x77\x98\xb5\x9d\x54\x0e\x5e\x22\xaf\xf1\x83\x4b\xb0\xe0\x46\x49\x9d\xeb\x3e\x63\x08\x32\x38\x15\x6d\xf9\xd5\x00\xa3\xaf\x59\x43\xf1\xdf\xd6\x6d\xdc\xb5\x1a\x9e\x98\x60\x71\xd3\x95\xc3\x6a\x95\x2b\x00\x32\x8d\xad\x63\xf8\xca\x57\x82\x1a\x8e\x23\x1c\x98\x19\x65\xa1\xe2\x07\xa9\x56\xff\xe4\x45\x2d\xc2\x40\x53\x67\xda\xb5\xec\x9a\x45\x80\x97\x41\x60\xe1\x92\xe6\x2d\x5c\x50\xdb\x15\xee\xa3\x45\x6e\x42\xb7\x32\x66\x41\x4c\xc4\x23\xb4\xf0\xbf\x62\xbb\x2f\x75\x05\x01\x6c\x98\x25\xf6\x2e\x8f\x24\x6f\x32\x83\x61\xe2\xe0\x9d\xa9\xbe\xf0\x15\x5f\xf8\x0d\x0d\x5f\x0d\x39\x9f\x6b\x01\xf6\x2e\xf2\x55\x6e\x5a\x90\xd8\xb6\xbe\x84\x25\xac\x06\xdb\xff\xe9\x8f\xae\x8a\x90\x7a\x40\x7b\xcb\xce\x6a\x50\x0c\x2d\xa0\x7d\x82\xa6\x0e\x25\x1e\x9e\x71\x68\xcc\x78\xfd\x54\x6a\x52\xc3\xa5\x14\xb4\x38\x83\xcd\x43\x24\x88\xd9\x77\x2f\x63\x06\x72\xfc\xd9\xde\xaa\x7a\x67\x91\x84\xa8\x92\xa6\xab\x3f\xee\x9b\xeb\xcd\xc0\x4e\xc5\x43\x3b\x15\xfb\x76\x2a\xf0\x56\xb3\xb9\x53\x57\x7d\x86\x2a\x72\x3e\x80\x31\x95\x8b\xb5\x18\xf7\x7d\x26\xcc\x27\x32\x56\x68\x6d\x16\x6d\x4d\xbf\x47\x53\x85\x64\x30\x62\xd5\x1d\x1f\xb8\x6e\x0b\xff\x58\x07\xf9\x98\x04\x76\x52\x33\xce\xb6\x13\xcf\x2c\x01\xae\x36\xa1\x18\x38\x80\xde\xb2\x85\xf4\x8c\xad\xc4\x4a\x42\xae\x83\x21\xea\x14\x98\x09\x3f\xc7\x9b\x6d\x7e\x65\xfc\xea\x0e\xb8\x1d\xe8\xe5\x5a\xa5\x8f\xb6\x46\xa7\xa9\xe3\xf0\x2b\x81\x59\x1b\xa1\x0f\xe1\xf5\x01\x4f\x17\xbc\xde\x4e\xd3\x8c\xd8\x6d\xc1\xfd\xfc\xf6\xc4\x62\xff\x2c\x46\x1c\xf6\x1e\xe6\xf8\xe7\x2d\x3d\x59\xe8\x2a\x7a\x2c\xaa\x72\x70\x93\x9c\x63\x18\xb2\x93\xe3\xd9\xe9\xbc\xee\xee\xde\x40\xda\x3e\x26\x72\x7d\x08\xef\x3e\xf9\x62\xe0\xc9\x11\x26\x4e\x43\x72\x1e\xa4\xd4\x79\x0f\xb0\x1f\xe1\xaf\x05\xdc\x47\x2b\x38\x92\x10\xb4\xb1\x8f\x91\xba\x0b\x66\x2b\x40\xc7\x23\x1a\x62\x9c\xfc\x24\xf2\xc5\x12\xcd\xf2\xdd\xe0\xf4\x07\x61\x96\x32\x43\xe1\xcf\x2f\xac\xd4\x77\xc1\xf4\xd3\xec\x34\xb8\x1f\x5c\x6e\xa3\xcc\x85\x57\x2f\xef\x86\xb9\x77\x4f\x66\xa8\xf6\xf4\x75\xfc\xf0\x53\xd3\xc9\x88\x37\x5b\x1b\xc3\x9a\xec\x63\x46\x2c\xe2\xfe\xc3\x9f\x2f\xe1\xf4\x0e\x6f\x36\x78\xde\x6e\xdc\xb9\x43\xc7\x80\x05\xbf\xdf\x68\xc5\x46\x7d\xa7\x36\xf0\x7a\xb0\x5b\x7b\x4d\xdd\xce\xb1\x1d\x9f\x68\x98\xfd\xa0\xf7\x7e\x3c\xde\x70\x9e\xb7\x7e\xc0\x7d\x94\xb8\xbb\xd8\x3f\x83\xfb\x1e\xc1\xae\xef\x40\xac\xd9\xbe\x8c\xd3\xb6\x03\xe9\x32\x1e\xb6\x1c\xf6\x7a\xd0\x36\x85\x9f\xea\xbf\x86\x4b\x34\xc4\x7a\x38\xf3\xce\x36\x7c\xd7\xae\x1e\xf0\x9c\x3d\x06\x86\x79\x3f\x83\xe3\x0e\x66\xd6\x77\x9b\x95\xea\x4b\x38\xed\xce\x3b\xd7\xa3\xdf\xe7\x35\x2c\xff\x9f\x06\xa5\x96\x43\xb4\xcd\x74\x07\x8c\xba\xc9\x83\xd2\xd0\xad\x1d\x70\xa4\xbb\x9d\x0c\x31\x7e\x06\x3f\x1e\xc8\x6a\x4f\xf2\x1d\xc4\x65\x67\xe2\xd1\xc5\x6c\x9f\xfb\x32\x51\x88\xa7\x27\x5d\xc3\x25\x1a\x62\x3d\xe8\xa9\xb7\xc7\xef\x8f\x4f\x8f\xfb\xce\x6a\x09\x0e\x41\xcc\x8e\xfb\x33\xb8\xeb\x40\x66\xbb\x93\xc5\x32\xd8\x61\x6d\x57\x14\x41\x4d\x04\xc5\x8e\xd7\x0e\x39\xc1\xd2\xe9\x32\x2f\xb3\xe6\x79\x2b\x94\xdd\xd8\x41\xb0\xed\x82\xee\x35\x1b\x6f\x79\xa8\xe6\x3d\x3f\x9d\xd0\x52\xb4\x54\xcc\x76\xbe\x7c\x03\x13\x20\xb0\x66\x49\x92\x0c\xbd\x55\x33\xc5\x69\x5b\x6b\x16\x72\x91\x4c\xdd\x73\x51\xac\xba\x88\xa7\x7d\x1c\xca\xfa\x19\x1c\xf5\x5b\xb5\xec\xcd\xc9\xd9\x5b\xbf\x40\xeb\x5e\xf8\xf1\x5f\xba\x01\x2d\x9a\xce\xa9\xee\x82\xca\xbe\xa7\xd3\x8b\xad\xde\xc4\xae\x3d\x9b\x17\xaa\x32\xa1\x53\x95\x57\xd8\xc5\x61\x61\xe7\xc1\x78\xbb\xc8\xc7\x97\x81\x9c\xd6\x9b\x72\xf5\xdf\x61\xf2\xa2\xdb\xa3\x49\xc8\x54\xa1\xb3\x27\x98\xd3\x49\x44\x0e\x63\xf6\xd2\xd4\x48\x05\x36\x70\xe5\xf3\xe4\x55\x3b\x66\xfd\x05\x4e\x74\x57\x49\xaf\xb2\xae\x78\x99\xa7\xa1\x7d\xba\x0c\x91\x62\xdf\x74\x73\xef\x61\xfd\x27\x00\x00\xff\xff\x7c\xdf\x5e\x84\x2a\x27\x00\x00")

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

	info := bindataFileInfo{name: "rest/gorilla/pat.tpl", size: 10026, mode: os.FileMode(436), modTime: time.Unix(1463557256, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _storeUpperioTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x58\x4d\x6f\xdc\x38\x12\x3d\x77\xff\x0a\x4e\x2f\x66\x40\x05\x5e\x75\x0e\xc1\x1c\x32\xf0\x61\x36\xce\x0c\x0c\x38\x99\xac\x3d\xd9\x4b\x10\x0c\xd4\x12\xd5\x26\x2c\x89\x32\x49\xb5\x63\x18\xfd\xdf\xf7\x55\x91\x52\xcb\xdd\xf2\xd7\x20\x97\x5d\x03\x6d\x48\xfc\xa8\x2a\x56\xbd\x7a\x55\xe2\xdd\x9d\x28\x54\xa9\x1b\x25\x16\x6d\x96\x5f\x65\x6b\xb5\x10\xdb\xed\x7c\x16\x5f\x04\xe6\xd3\x4f\x57\x6b\x1a\xc3\xa3\x6a\x0a\x7a\x9a\x8f\x76\xe9\xba\x35\xd6\xbb\xb0\x6b\xb1\xd6\xfe\xb2\x5b\xa5\xb9\xa9\x97\x6b\xd3\xd9\x62\x79\xa5\xfd\xd2\x79\x63\xd5\xe2\xd1\xd9\x65\xd7\xb6\xca\x6a\x83\x55\x90\xad\x4b\x91\x9e\x16\xe9\xa9\xbb\xf0\x56\x37\xa4\x7c\xbc\xd5\x65\xd8\xa1\x21\x21\xed\x3a\x5d\x2c\x76\x66\xed\x29\xf8\x27\x49\xa7\x5f\x65\xd6\x8b\xf9\x03\x82\xb1\x49\x35\xb9\x29\xf0\xb6\x5c\x65\x4e\xfd\xfc\x26\x98\x30\x88\x2c\x6b\x4f\xa6\x6b\xb3\xd4\xa6\xf3\xba\xa2\x17\x36\x36\xc5\x50\xb1\x5a\x3c\xe0\x16\x88\x0c\x9e\x9c\x2f\x97\xec\xc4\x3f\x6f\x5b\x95\x7e\xcc\x6a\x85\xc1\x0b\x3a\xf2\x27\x6b\x36\xba\x50\x56\xc0\x83\x95\xaa\x55\xe3\x9d\x60\x5f\xa4\xbb\x99\xc6\x2b\x5b\x66\xb9\x22\x21\x6d\x18\x75\xc2\x66\x37\xd3\x12\xe7\x65\xd7\xe4\x8f\x2b\x93\x4e\x39\xb7\x93\x7b\xb7\x4d\x84\xec\xd5\xf2\xc2\x23\xa1\xac\xa5\x9f\xb1\x89\xb8\x83\xd7\x36\x99\x15\xc5\xea\x82\xb6\x15\xab\xf4\x24\xf3\x19\xb9\x29\x8c\x9b\x2b\xb1\x32\xa6\xc2\x2a\xb8\x78\x0d\x83\xdf\x1e\x0b\x3c\xa5\x1f\xd5\xcd\x99\x59\xc3\x73\x67\x3c\x2c\x83\xe7\xd2\x13\xed\xf2\xcc\x16\x09\xd6\x23\x10\x41\xe8\x11\x09\x39\x16\x64\x56\x2a\x47\x0a\x92\x5f\xc4\x0f\x98\xb9\x9b\xcf\x66\x64\xd0\xb1\x80\xb4\xf4\x3d\x99\x55\xca\x85\xfa\xd6\xaa\xdc\xab\x62\x6c\x11\x0e\xc5\x52\x8e\xc4\xda\x78\xf1\xe3\x3f\x36\x8b\x23\x7e\x4f\x20\xc1\x2a\xdf\xd9\x66\x3e\x43\x34\x66\xf0\x64\x8c\x11\x9f\x5a\x64\x88\x5d\x3f\xef\xa0\xe7\xa7\x49\xff\xdd\xf5\xd6\x86\x83\x02\x18\x71\xcf\x23\x01\x86\x7a\xbb\x41\xbc\xd6\xaa\x01\xb4\x73\xf1\xee\xf3\xf9\x89\x28\x8d\x15\x1e\x2b\x0f\xb6\x90\x98\xdf\x69\x65\x46\x07\x5b\xdd\x0a\xce\x10\xf1\xee\xec\x54\x78\xf2\xf1\xe4\xa6\xa8\xc7\xdb\x2e\xf7\xe4\xab\x93\x95\xa0\xbf\x7b\x81\x8a\xa1\xa1\xb8\x84\x70\x44\x9b\xdf\x59\x05\x55\x22\x3b\x10\x4a\x9e\xf4\x97\x4a\x14\x51\x04\x42\x54\xf2\x40\x9b\x59\xa0\x34\xa0\x0c\xa8\x79\x35\x69\x4d\x12\x05\xcb\xf9\x2c\x37\xf0\x6d\xc0\xd6\x3b\x3c\xc2\x7b\xaa\x8d\xef\xef\x1b\xaf\xfd\xed\x27\x0f\x90\xc9\x3d\xc0\xc1\xb4\xb5\xf2\x22\x37\x55\x85\x20\x6b\xd3\x90\xa0\xaa\x0a\xc0\x04\xc0\x1c\x84\x55\x95\x4c\x18\x44\x34\xf6\xc3\xb1\x68\x74\xc5\x50\xd9\x0b\x74\xd6\xb6\xd5\x2d\xf2\xa5\x29\x4c\x2d\x88\x2c\xc8\x55\x94\xf7\xde\xf4\x4f\xba\x78\x90\x18\x68\x3d\xf4\xd1\x3e\x42\xf4\x7f\xde\x90\x4e\x45\x43\xaa\x4d\x25\x4e\xbf\x3b\xfc\x76\x4b\x53\x29\x39\x04\x52\x7a\x3f\x1e\x8b\xc0\x28\xe9\x79\x76\xf3\xf9\xfc\xec\x7d\xe4\x99\x94\x1f\xd4\x9f\x26\xe8\x92\x90\xff\xe5\xed\xd7\x64\x4c\x3b\x6c\xfd\x87\xcc\xba\xcb\xac\x62\xd7\x6b\xaf\xea\x23\xb2\xb2\x35\xce\xe9\x55\xa5\x78\x85\xbc\xee\x74\x7e\x25\x4a\xfd\x8d\x71\x15\x39\x94\x98\x02\x2b\x6a\x71\x03\x36\x24\x2c\x44\x41\xca\x06\x9f\xd5\x8a\x93\x2e\x9e\xe3\xde\xfc\x2f\xa2\xcf\xb9\x36\xf8\xfb\x18\xab\xfb\xf9\x93\x7f\x91\x03\x26\xbc\x3e\xb8\x1d\x7e\x1f\x5c\x5f\x14\x6c\xb8\xe2\x40\x93\xc3\xc7\x01\x7d\xc0\xe3\x7f\xf5\x4a\x69\x6d\xfa\x2b\x8e\xd3\x14\x52\xb5\xd1\x35\x95\xe3\x3c\x99\xe9\x62\xc0\xc2\xe4\xba\xc8\xdc\x87\x76\x06\xd9\x2e\x55\x91\x46\x98\x4e\x44\x4e\x70\x25\x0b\xf6\xe1\xfc\x56\xfc\xe8\x16\xac\x2a\x10\x8f\x4c\xf6\xb9\x64\xf2\x18\x63\x53\x07\x0c\x92\x2b\xae\x14\xfb\x61\xe7\x95\x97\xa2\x29\x0e\xd0\x3a\x0c\x48\xc0\x52\x82\xc9\x7f\x7e\x93\xec\x61\xe7\x1e\x35\x5d\xa8\xcc\xe6\x97\x53\x69\x0e\x8e\xd1\xa8\x39\x94\xa5\x9a\xc2\x22\xc1\x95\x4f\xe5\x76\x90\x86\xdc\xbe\x8e\x89\xfc\xef\x4e\xd9\xdb\x24\xbe\x9c\x2b\xd7\x55\x9e\xb3\x38\xd8\xd0\x43\x92\xf2\x27\x4c\x4a\xd2\x20\x91\xf5\x56\x71\x41\x09\xa3\x7b\x55\x67\x36\x45\x02\x0f\xb1\xc0\x13\x80\x0c\xc2\x30\x60\xb5\xda\x44\xcf\x6b\xe8\x26\x86\xd5\x1b\xd5\x88\x6b\x3a\xc1\xce\x0b\x8e\x35\x81\xad\x38\xf7\xa3\xf9\x4c\x5f\xf2\x3a\xfd\x5d\xf9\xf0\x98\x44\xc5\x61\xe5\xf1\x3d\xd5\xae\x07\xf0\x6f\x1a\xb0\xa4\x85\xdb\x80\x88\xc9\x79\x96\x90\x8c\x4c\xa5\xcc\x71\x68\xa9\x98\x9d\x1a\xe4\x75\x9d\x91\x5d\x9c\xfd\x59\x73\x3b\xef\x45\xe0\x7f\x7a\x81\x75\xb2\x37\x92\x5f\xae\x93\x34\x4d\x93\x28\xea\x12\xcc\x57\x11\x6f\xaf\x21\x2c\x18\xcc\x67\xf8\xa3\x2c\x9d\xf2\x88\x02\x7c\xf6\x7a\x6c\x16\xcb\xbc\xd2\x2d\x38\xa9\xf1\xf2\xde\xda\x24\xd8\xb8\x13\x72\xa6\x6b\xfd\x90\x8c\x30\xb7\x13\x12\xd7\x26\xc3\x39\x87\x2c\x82\xa9\x01\xa7\x7f\xa0\x22\x87\x51\xc7\x29\x52\x6a\xeb\xfc\x21\x68\xe1\x8c\xfc\x52\xbd\x10\xb5\x90\x4d\xe5\xe8\x6f\xd5\xa2\x01\x39\x96\x91\xea\x44\x69\x51\x4c\x8a\x5d\x79\x25\x9c\xfc\xf4\xe5\xeb\xbe\xea\xbb\x2d\x25\x09\x21\x95\x95\x20\x01\x38\x55\x64\x92\x5e\xf4\x20\xca\x93\xd8\x8f\x74\x75\x3b\x88\x87\xcf\xc0\xe0\x86\x5b\x34\x2a\xbc\x75\xd6\x8a\x65\x2c\xf1\xf3\x81\xc2\x62\x1e\x22\xda\xbf\x22\x0b\xaa\xe7\x14\x43\x2c\x68\xd0\x19\x95\xa6\x6b\x40\xa0\x56\x51\xdb\xce\xdb\x2a\xd5\xc8\x57\x55\x42\x28\x7e\x3d\xe6\xc9\xe0\x1c\x72\xc6\x47\xe3\x7f\xa3\x6d\x13\x15\x16\xd5\x68\x1d\xda\x85\x4d\x56\x75\x8a\x2c\x0e\x69\xc5\x47\xe0\x35\x91\xf3\x42\x40\x7b\x77\x16\x61\xfd\x7c\x26\x5f\x45\xfe\xbb\xef\xbe\x04\xf6\x08\x32\xeb\xcb\xeb\xaf\x03\x9b\xe0\x64\x11\x2d\x9f\xdb\x82\x9a\x97\x03\x78\x98\xe6\x65\xc8\x08\x72\xfe\x36\x38\xbe\x57\xa3\x42\x72\x56\x23\x12\xe2\xae\x14\x8e\xa5\xee\x8e\xeb\x05\x69\x0d\x0d\xd5\x91\xf8\x8b\x2b\x1f\xa5\xd5\x87\xac\x25\xf1\x94\x77\x6f\xc7\x9c\x02\x5e\xa5\x73\x30\xb7\x24\xc9\xff\x41\x37\xd1\x85\x70\x93\xed\x81\x00\x8a\x81\xc8\xfb\x9c\x20\xda\x89\xd1\xe4\x4e\xe0\xb9\xe5\x9f\x45\x3f\xbb\xfc\xef\x37\xfe\x27\xaa\x52\xdf\x03\x87\x41\xce\x3e\x0e\xff\xd7\x31\x67\x55\x6d\x36\x4f\x87\xed\x9c\x97\x4d\x1b\x3d\x1d\xb4\x82\xdc\xf5\xf2\xa0\x8d\x08\x04\xb4\x69\xf2\x90\xdc\x22\xa3\x67\xc2\x57\x0d\x43\xd0\x0b\x10\xfa\xb3\xa6\x6f\xd2\x9e\x0a\xdd\x48\x92\x4c\xf6\x59\x83\xce\x10\x75\x1f\x7c\x58\xa2\x42\x1c\xd8\x72\xa6\xc1\x91\x8f\xda\x23\x2a\x2c\x79\x89\x51\x24\x72\xcf\x30\x1a\xda\x33\x6e\xb2\x80\x05\xf3\xce\xc0\xe6\xba\x71\xf4\xc9\xcd\xa1\x44\xbd\x58\x83\x0a\x4c\xf9\x52\xa3\x20\x48\xb6\xd5\x94\x25\x89\xe0\x46\x96\x0c\x52\x5c\x50\xdb\x0a\x35\xe1\xd0\xa6\x64\x30\x98\xd7\x4b\x2e\x5d\xaa\x42\x90\xe3\x37\x2d\xe0\x18\xbb\x08\x36\x95\xae\x4a\xfa\xcb\x9a\x71\xd2\x3c\xf9\x09\xcb\x49\x24\x24\x6d\x11\x8c\xec\x7e\xeb\x7e\xa7\x1a\x13\x89\x14\x4d\x27\x25\x81\xf7\x64\x2c\x41\x2e\xd0\xeb\xd3\xeb\x76\xbb\x78\x01\xe4\x91\x57\x0d\x09\x00\xe8\x77\x8a\xc4\x20\x2a\x00\x9f\xe8\xf3\x71\xc2\x42\xf3\x11\x6e\x01\x04\xba\xba\x10\xce\xf0\x5a\x1a\x00\x82\xde\x1f\xb9\x5a\x7a\xf4\xc3\xa0\xbf\xec\x39\xb8\x6f\x60\x3f\xb9\x34\x8e\x1f\x47\x85\x3d\xb8\xcc\x9a\x06\x42\xcf\x57\x2b\xe7\xe8\xd2\x91\x4b\x0d\x5d\xde\xd0\x11\xf1\x85\xfe\x94\x72\xf6\x94\xac\xdd\x3a\x7e\xd5\xc7\x11\xd6\xdb\x13\xe3\xae\x9b\x39\xa5\xd6\xaa\xc9\xaa\x30\x89\x56\xca\x6e\x94\xfd\x80\xcd\x28\x4f\x6e\xbd\x33\x95\xcc\x97\x8b\x70\x79\x79\x24\x16\x93\xba\x69\x22\x5a\x8d\x47\x6c\xdf\x21\x94\x84\xef\x0e\x59\x7e\xa7\x53\x96\xa3\x63\x1e\x89\x8d\x40\xb7\x7f\xef\x36\x6f\x38\x78\x6f\x45\x80\x91\xa4\xeb\xb3\x8b\x16\xbb\x3c\x4b\xc0\x56\xfa\x4e\x18\xf2\xa6\x32\xf8\x3e\xb9\x31\x5d\x55\x70\x93\x98\xf3\x7b\xdf\xe4\x0e\xd8\xa3\x22\xc1\xf4\x94\xd2\xa6\x4f\x95\xa2\xc9\xce\xa9\xbe\x66\xd1\x2e\x30\xcf\x69\x23\x73\xff\x2d\xa1\xc6\xef\xc6\xa2\x7d\xed\xda\x91\x04\x47\x5b\x35\x04\xd1\x98\x57\xdf\x9e\x71\x9f\x44\x62\xe5\xe1\xd1\x22\x99\x8f\x2e\x60\xff\x1b\x00\x00\xff\xff\x29\x41\xcf\x97\xcd\x16\x00\x00")

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

	info := bindataFileInfo{name: "store/upperio.tpl", size: 5837, mode: os.FileMode(436), modTime: time.Unix(1463553974, 0)}
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
	"general.tpl":             generalTpl,
	"rest/gorilla/pat.tpl":    restGorillaPatTpl,
	"store/upperio.tpl":       storeUpperioTpl,
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
