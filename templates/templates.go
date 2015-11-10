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

var _endpointsEndpointsTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x58\x4b\x6f\xe3\x36\x10\x3e\xcb\xbf\x62\x56\x45\x0a\xa9\x70\xe5\x9e\x53\xf8\x50\x24\xe9\x03\xd8\xee\xa6\x9b\xec\x29\xf0\x41\x91\x46\xb2\x60\x89\x94\x49\xca\x8d\x61\xf8\xbf\x77\x86\xa4\x6d\xd9\x49\xe3\x6d\xda\x00\x1b\xd4\x07\xdb\xd2\xbc\xf8\xcd\xe3\x23\x09\xaf\x56\x90\x63\x51\x09\x84\xb0\x4d\xb3\x59\x5a\x62\x08\xeb\xf5\x20\xf0\x2f\x40\xfa\xe4\x7a\x56\xb2\x8c\x1e\x51\xe4\xfc\x34\xe8\x79\x55\x4d\x2b\x95\xd1\xce\x2b\x2c\x2b\x33\xed\xee\x93\x4c\x36\xa3\x52\x7e\x3f\xab\xcc\x88\x3f\xe4\xd6\xca\x4a\x98\x70\x10\x94\xb2\x53\x79\x66\x1e\x60\xdf\x94\x84\xd6\x32\x93\xc2\xe0\x03\x1b\x3e\xad\xd7\x46\x2a\xb4\x5a\x59\xa7\xa2\x4c\xa4\x2a\x47\x0f\x23\x81\x3d\x4f\x52\x16\x8d\x8d\x50\xcb\x32\xfc\x1b\xd4\x99\xcc\x5d\xa2\x83\xd1\xc8\xe6\x78\xc3\x81\x93\x0f\x69\x83\x24\xbd\xf2\x78\x35\x28\x34\x9d\x12\x70\xf1\xf9\xd3\x25\xe0\x56\x5a\x48\xf5\xd8\x69\x50\x74\x22\x7b\x26\x56\x24\x64\x27\x86\xc0\xdf\x2d\x68\xa3\x2a\x51\xc6\x10\xed\x82\x36\x69\x7b\xe7\xc4\x93\x8d\x30\xd9\x38\xc7\xb0\xa2\xbc\x08\xea\x22\x55\x55\x7a\x5f\xa3\x06\x23\xa1\xd3\x08\x75\x6a\x50\x0d\x82\xb4\xae\x65\x76\x25\x4c\x65\x96\x70\x3e\x06\x86\x12\xc5\xf0\x1d\xa3\xb9\x5d\xb6\x8c\x03\x56\x9b\x64\xbe\xed\x49\x57\x6b\x58\xef\x79\xbf\xaf\xb4\xe9\x47\xb8\x9b\x3c\x1d\x63\x4f\xee\xa2\x94\x68\x6c\xe6\xec\xfe\x0b\x9a\xc7\x05\x0a\x6c\xf7\xec\x1b\x99\x84\x8f\x0c\x42\x97\xa3\xb5\xea\x55\x7b\x8a\x0a\xad\xe2\xf6\xe3\xe5\xc7\x73\xaa\xd3\x12\xa6\xe9\x02\x41\xe0\x9f\x5c\xc7\x2e\x33\x5c\x0c\xeb\x35\x08\x76\x6e\x63\xb2\x9c\x61\xf4\x6c\x59\x07\x3d\x87\xbb\x30\x53\x48\xd5\x0c\x27\xe0\xd3\xe7\x31\xf5\x63\x95\x5c\xb8\xdf\x21\x20\x90\x31\xaa\x22\xcd\x70\xb5\xa6\x06\x2a\xea\x45\x4f\x42\x06\x4a\xf1\x47\x2a\xd7\x34\x06\x4e\x95\xd9\x04\x22\x5b\x9a\x9e\x26\x35\x95\x14\xa4\x54\x43\x90\x33\xae\xc6\x86\x17\xc9\xaf\xb7\xb7\xd7\x9f\x70\xde\xa1\x36\x0c\x20\x26\xa3\xaa\x80\x77\x64\xb5\xa2\xc7\x40\x73\x78\xb2\xb7\xe9\x26\x57\xbc\xce\x6f\xbc\xba\x48\xeb\x8d\x3a\xb9\x41\xb5\x40\xf5\xbb\x2e\x29\x91\xb0\xa9\xb4\xa6\xe4\xa9\x71\x36\x26\xad\x0f\x3b\x92\x05\x01\x87\xa3\x68\xf4\xc3\x6f\xae\xbb\xf4\xb4\xee\x21\xf7\x95\x0d\xb4\xcb\x8d\xc1\xfa\x4e\x47\xca\xc3\x63\xf9\xbb\x31\x88\xaa\x7e\x19\x4a\x22\x6c\x72\xd3\x52\x93\x4c\x11\x85\xb6\x78\x20\xef\x4d\x5a\x09\x46\x7e\xa6\xfd\x48\x44\x67\x3a\x0e\x87\xb0\x9d\x22\x8b\x27\x7e\x36\x8b\x20\x20\xc6\xa3\x02\x9d\x5c\xd4\x52\x63\x14\xfb\xbc\x5c\xab\x69\xc8\x78\xe6\x49\x44\x5b\x45\x72\xed\xd7\x77\xba\x73\x38\xfb\x66\x41\xab\x21\x2f\xe0\xe3\x27\x17\x56\x15\x51\x9e\x5e\xf1\x1f\xe7\xce\x06\x81\x2f\x80\x85\xe1\xf2\xb7\x58\x86\x1e\xed\x06\x98\xdb\x4e\x38\x9a\x5d\x23\x8a\x3d\xa4\x63\x2d\x45\xc1\x9b\x1f\xcd\x83\x6e\xa5\xd0\xdc\x58\x1e\xe1\x71\x7f\xff\xe9\xcd\xb3\x4d\xc9\xee\x59\xe7\x70\xb7\xa7\xc0\xf5\xd0\x97\xd8\x2f\x31\xb0\x4b\xf4\xf8\x44\x72\x55\xe1\xe2\x18\xa3\x76\x93\xf9\x8f\x78\x35\xe7\x22\x7b\xdf\x24\x72\xd5\xfe\xa3\x43\xb5\x8c\x4f\xac\xfb\x4a\x59\x67\xcf\x19\xe6\x5d\x83\x8d\x54\x4b\x7b\x8e\x6e\x4f\x34\xa6\x59\xcd\xf0\x0f\x4e\xa3\xad\xf7\x66\x9c\x7a\x7c\xbc\xc1\x54\x65\xd3\x68\x1e\x27\x3f\xd5\x75\x84\xf5\xeb\x15\x46\xdb\x95\xb6\x74\xd4\xe1\xd0\x72\x75\x4b\xc2\x97\x31\x8f\xb0\xea\xe4\x3d\x0a\x86\x0e\xe3\x31\xfc\xe0\xf0\xfa\x38\x3b\xb4\x1f\xa4\xf9\x99\x56\xca\x1f\x87\xfd\x52\xee\x62\x7d\x8c\xad\x35\xd5\xfa\xc4\xd4\x13\x53\xff\x2d\x53\xdf\x20\x37\x5f\x4c\xa2\x43\x0e\x75\x6d\x7e\xfc\x06\xf9\x42\x16\x11\x29\xda\x3d\x22\x3d\x8d\x37\xde\x12\x8e\xf5\xe1\x9c\x79\x16\x4e\x0e\x68\x47\xe5\xd8\x99\xb8\x6b\x45\x38\x21\x29\x61\xcd\x59\x31\x4f\xe8\x02\x4f\x80\x73\x1d\x9d\x38\xfa\xb5\x72\xd4\x0d\xdb\xee\x0e\xeb\x73\x60\xee\x7d\xb6\xaa\x88\xdb\xc9\xb7\xc2\x1f\x5f\xed\x92\x6a\x8f\x34\x9f\x20\x1f\x20\xaf\x4a\xbf\x23\xf7\xcf\x43\x32\xe6\x58\xe3\x6b\x91\xf1\x4b\xb6\xc8\x83\x1d\x32\x8a\x8f\xde\x70\xdc\x88\x58\xca\x82\x2c\xa0\xca\x9f\x3f\x3c\x4f\x7c\x7d\x53\x7c\x2d\x2a\x6a\x96\x99\xa2\xab\x85\x30\xdc\x62\x7e\xb5\x6d\x7e\x83\xa7\x26\xa5\xe4\x28\xf6\xe4\x16\x74\x69\x55\x76\x0b\xfa\x3f\xed\x3f\xfe\xd9\xfd\xdb\xe8\xff\x78\xfc\x2b\x00\x00\xff\xff\x87\x18\x1f\x1a\x64\x15\x00\x00")

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

	info := bindataFileInfo{name: "endpoints/endpoints.tpl", size: 5476, mode: os.FileMode(436), modTime: time.Unix(1447127629, 0)}
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

var _restGorillaPatTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x5a\xeb\x6f\xe3\xc6\x11\xff\x2c\xfd\x15\x13\x02\x29\xc8\x8b\x42\x5d\x81\xa2\x1f\x5c\xb8\x40\xce\xa7\x5c\x52\xdc\xc3\xf5\xa3\xfd\x60\x18\x05\x4d\xae\x24\xd6\x2b\x2e\xbd\xa4\x64\x1b\x8e\xfe\xf7\xce\x63\x49\x2e\xf5\xf0\xd9\x4d\x8c\x3a\x35\x60\x99\xdc\x9d\x9d\x9d\x99\xdf\xcc\xec\xcc\xca\x0f\x0f\x90\xa9\x69\x5e\x28\x08\xca\x24\xbd\x4e\x66\x2a\x80\xf5\x7a\x38\x70\x2f\x80\xf3\xf1\xf1\xf5\x8c\xc6\xf0\x51\x15\x19\x3d\x0d\xbd\x55\xf9\xa2\x34\xb6\xae\x64\x55\x30\xcb\xeb\xf9\xf2\x2a\x4e\xcd\x62\x3c\x33\xdf\x5f\xe7\xf5\x98\x7e\x71\x59\x69\xf2\xa2\x0e\x86\x83\x79\x5d\x97\xb5\x4d\x8a\x8a\x56\xc1\x1e\xfa\x96\x60\x4c\xe4\xc1\x26\x5f\x9b\x6b\x9d\x8c\xcb\x84\xf8\xcd\xcc\xd2\x66\x69\x7d\xb7\xc1\x0a\x07\x99\x53\x6a\x8a\x5a\xdd\xd5\x5b\x2c\x9a\xf9\x52\xd9\xc5\xde\xc9\xaa\x36\x56\xf1\xac\xd1\x49\x31\x8b\x8d\x9d\x8d\xef\xc6\x85\xf2\xd8\xe2\xa4\x2a\x52\x93\xe5\xc5\x6c\xfc\xef\xca\x14\x44\x3d\x5d\xf0\x7e\xda\xcc\xe8\x0f\x91\x37\x5a\x54\xb5\xc5\x95\x2b\xf7\x88\x6b\xaa\x60\x8f\x55\x91\xa5\x00\x31\x1c\x8f\x19\x83\x53\x92\x05\x07\x4e\x54\x55\xc3\x55\x5e\x64\x15\xb0\x78\x50\x1b\x40\x4b\x80\x35\xcb\x5a\xd9\xe1\x74\x59\xa4\x5b\xf4\xa1\x85\x37\x48\x13\x9f\x30\xcd\x08\x4a\x20\xb5\xe3\x4f\xcb\xbb\x11\x5c\x25\x95\x1a\x41\x61\x96\x85\x7c\x96\x20\x92\x45\xf0\x80\xca\xe1\xe6\xab\xc4\xe6\xc9\x95\x56\x15\xed\xb4\xac\x14\xe8\x84\x36\x1a\x24\x5a\x9b\x74\x52\xd4\x79\x7d\x0f\x07\x87\x40\x1b\x87\x11\xbc\xa1\xbd\xcf\xee\x4b\xda\x1a\x1e\xc0\xaa\x7a\x69\x0b\xf8\x83\x37\xfa\xb0\x86\x75\x6f\xf5\xc7\x1c\x35\xf2\x38\x5c\x5c\xee\xe6\xd1\x1b\x17\x2e\x33\x55\x8b\x9e\xb8\xfc\x83\xaa\x7d\xbd\x87\x03\x36\xcf\xe7\x64\xc1\xb3\x81\x3f\x17\x88\x66\xaa\x98\x1a\x9b\x2a\xfc\xcb\x4a\x94\xd6\xa0\x59\xf0\xe1\x4a\x4d\x89\x2e\xb5\x0a\x35\x1d\x0e\xd0\x00\x38\xa7\xca\xc4\xaa\x23\x1e\x82\xc6\x9d\xe3\x4f\x79\x96\x69\x75\x8b\x33\xe0\xe4\xcf\x8b\x42\xd9\x8e\x60\xe2\x1e\xa2\xed\x21\xb4\xef\x60\xe0\x74\x63\xd4\x42\xf2\x62\xe7\x58\xf1\x91\xfc\x1d\xa1\xf6\x37\x4b\x82\x1c\x57\x28\x3b\x4d\x52\xf5\xb0\x8e\x20\xb4\x0a\xc3\x03\xbd\xc6\x1b\x1d\x81\xb2\x96\x7e\x8d\x8d\x98\x37\xa9\x58\x6a\x9c\x9b\x1b\x9d\x29\x7b\x00\x49\x71\x5f\xcf\x11\x5a\xb8\x37\x4b\xb8\x4d\x50\x04\x44\x34\x33\x70\x8b\xbe\x0f\xf5\xbc\x31\x84\x5b\xda\xfb\x71\x26\x49\xca\x92\x7c\x95\x96\x25\x75\x42\xae\x43\xc4\x4e\x09\x56\x9d\x94\x68\x85\x8e\x70\x16\x81\x58\x3f\xc9\xdc\xcb\x32\xdb\x34\xf7\x39\x0f\xbd\x8c\xb9\x9f\x6f\x6d\x74\xfd\xfd\xe6\x66\x3b\x2c\x92\x92\x7c\xcd\x31\x89\x43\x7c\xbf\x90\x68\xba\xf4\x19\x0e\x9d\x85\xd1\x7b\x1b\x09\x90\x33\x1a\x61\x91\xd4\xb9\x29\x98\xd5\x08\xcc\x35\xf1\x6a\xd2\x5b\xfc\xd3\xd9\xd9\xf1\x89\x30\x26\xd1\xc9\xb4\x83\x7c\x0a\xdf\x20\x19\x83\x3d\xa8\x48\x1e\x5c\xc1\x6e\x1f\x4f\x48\xb0\x9f\x69\xd3\x22\xd1\xed\x7c\x7c\xaa\xec\x4a\xd9\x4f\xd5\x0c\x0d\x18\x2c\xf2\xaa\x22\x77\xe8\x94\x86\x2e\x5d\xe2\x0f\x31\x44\x7e\xf8\x87\x5f\xc5\x78\xf4\xb8\x66\x0d\x94\xa6\xed\x36\x02\x39\x64\xc1\x6e\xd8\x0c\xa4\x7e\x80\xac\xed\x7d\x70\x19\x87\x22\xd7\xdf\xe9\xb5\x67\x01\x1e\xa7\xf7\x4a\x4c\x4a\x4a\xbb\xb0\x0e\x6d\xa3\x26\x4d\x7c\x73\x08\x45\xae\xff\x5b\x6d\x31\x29\xc7\xa7\x25\x62\x51\x4f\xc3\x80\x61\x03\x73\x55\x27\x79\x41\x16\xf8\xb6\x49\xa6\xe1\xb7\x55\x14\x8c\xa0\x4d\x1d\x2c\x52\xf4\x35\x6b\xe0\x07\xe6\x6d\x74\xc4\x2a\x3e\xd2\xa6\x52\x61\xab\x20\x26\xf3\x8c\x43\x8b\x0d\x8b\x5e\x68\xa6\xfc\x9a\x67\xc3\x8e\x25\xca\x99\xd8\x74\x1e\xde\x44\xf1\x0f\x5a\x87\x4a\xbf\xa8\xd6\x15\x6f\x26\x5a\x1f\xe0\x6f\x30\xe2\xa5\x03\x39\x02\x9e\xa2\xae\xd3\xad\x56\x5a\x8b\x32\x14\x84\x4e\x64\xad\x8a\xf0\x0d\x6a\x00\x7f\x85\xb7\x4e\x68\xf1\x03\x0c\xe9\x55\x70\x89\x2c\x79\xfa\xe2\xed\x65\x8f\xd7\x6f\x9c\xa7\x24\x97\x3c\x21\x4f\xa1\x68\xfd\x24\xe5\x32\x0f\x68\xf4\xe5\x5e\x2e\xe2\x53\x6a\x7f\x26\x82\xd7\x93\x8a\x58\x53\xa1\x92\x99\xc3\x3d\xa9\x79\xa7\x8b\x6d\x02\xbd\x6a\x93\x9a\xb0\xdc\x9f\xd5\x90\x58\xbb\xa3\x9c\x16\x5d\x70\x35\x81\x81\xdf\x3f\xd0\x9b\x8d\x99\xf4\x50\x76\xfe\xe5\x17\x78\xe3\xbf\x8b\x24\x6e\x08\x16\xc9\xb5\x0a\x7b\x3c\x46\xf0\x36\x6a\x23\xcf\xdb\x0b\x89\x19\xb7\xdf\xd8\xa9\x58\x8e\x46\x7f\xcf\x91\x68\xe7\x11\x09\xdc\xf3\xa0\x5b\x8b\x06\x13\x67\x68\xc1\xe2\x4d\x30\x45\x24\x4b\x5d\xd3\xc1\x57\x9b\xd4\xe8\x9e\x7b\x1d\xbb\xc1\xdf\x87\x8b\xb1\xdd\xdb\x74\xfd\x2b\xbc\xab\x42\xcb\xa4\x73\x58\xc5\x61\x8d\xd0\x3a\xef\x4d\x31\x5e\x61\xb7\x97\x1d\x38\x2e\x4e\xc2\x36\x13\xde\x95\x49\x91\x9d\xb8\xe1\x70\xb5\xd7\x49\x23\x97\xab\x09\x88\x3d\xbc\x3e\xab\xdb\x96\x11\xbb\xd5\x08\x56\x51\x27\x71\xab\x41\x87\xf8\x4c\xa1\xfe\x98\x6f\xaa\xd6\xa0\x54\x5e\xd3\xe1\x6a\xf0\x44\x9d\xab\xf4\x1a\x51\x5b\xb4\x70\x52\xd9\x5a\xb8\xa3\xfc\x18\x09\x8f\x1c\xc5\x41\x03\xb3\xb7\xba\xa9\xc5\x77\x79\xc5\x56\xfd\xf8\x4c\xf7\x78\xd9\x14\x34\x58\x90\x42\xdc\x67\x60\x65\x8e\xad\x46\x5b\xb5\xb8\xa3\x65\x41\x87\x9d\xb9\x15\xbf\xe9\x54\xee\xfb\xd0\x4e\x27\xf2\xbc\xc8\x81\xf2\x78\x09\xca\x79\x62\x37\x5c\x4e\x93\xa7\xe0\x25\xa4\xff\xa7\x80\xb1\x11\xbf\x1a\xd1\x4f\x45\xe3\xd7\x42\xdf\x44\xdc\xd3\xd1\x6f\x63\x78\xe5\xa7\x99\x2d\xec\x19\xff\x43\xfc\xe1\x17\x7e\xcb\x14\xb5\xda\x3f\xbf\xc7\x80\x14\xcf\xc8\x53\x2c\x69\xef\xdd\x78\x06\x07\x79\x86\x55\x9c\xd2\x19\x53\x87\xb7\xc6\x5e\x57\x92\xcf\xa9\xeb\xa6\xd2\x22\x83\xf3\x93\x8f\xdc\x7f\x13\xf2\x92\xd5\x7d\xae\xbd\x1b\x8f\xf8\x3d\xcf\xb8\xf0\xff\x91\xdc\xc0\xb5\x32\xd8\xa3\x13\x65\xec\xa6\x18\xb7\x2d\x34\xb7\xe2\x0c\xa5\xa3\xc3\x39\x46\x11\xa4\xb6\x0e\x23\x32\x7a\x18\xa0\xdc\x41\x04\x74\x1e\xe5\x58\xa6\xa5\xf3\xa4\x98\xd1\xd9\x95\x52\xd7\x78\xe0\xe5\x3a\xf4\x9d\xac\xc2\x45\x3f\x64\x59\x18\xe0\x9a\x11\xd6\xa6\x11\x3b\xaa\x6c\xee\x91\x36\xfc\x4f\x55\x2d\xab\x88\x59\x34\xec\x10\x59\xfb\x26\xa5\x82\x89\xd4\x97\x37\xe0\x3e\x00\xb0\x3a\x93\xe3\xb4\xf1\x70\xdf\x5c\xcd\x8a\x17\x34\xd8\xd0\xf5\x26\x9b\x3a\xd1\x04\x55\x0c\x89\x45\x2f\xaa\xe8\x66\x0a\x1f\xb1\xfe\xe7\x9b\x8e\x01\x0d\x9c\xd6\x56\x2c\xfd\x23\xf6\x69\xff\x48\xf4\x52\x85\x41\xc5\x17\x5f\x64\x01\x74\xd4\x86\x08\x9d\x35\x08\xc4\x57\x79\x5e\xb6\xe3\xcb\x1e\x2c\xc5\x75\x5e\x87\x8e\x72\x04\xc1\x88\x17\x0f\xc8\x2a\xff\x1a\xc9\xbe\xb4\x07\x61\x05\xb2\x58\x7c\xfe\x26\x3e\xc5\x37\x5e\xe8\x25\x34\x5f\xe6\x32\x99\xf9\xad\x9c\x2f\xbd\x99\x4e\x2b\x85\x31\xa5\xf3\x45\xde\x5d\xb4\x6c\x5b\xcf\x20\x09\x2c\xd1\x76\x7f\xfe\x93\xcb\xe1\xa6\xda\xa1\xb4\xb0\x13\xc1\xf5\x2e\x02\xde\x27\x68\x4a\x00\xe6\xe1\xd9\x84\xc7\xea\x36\xd3\xb8\x1b\xb1\xf8\x98\xb4\x38\xc7\xcd\x43\x5a\x30\x82\x3f\xbe\x1d\x01\xca\xf1\x17\x29\x60\x7b\xf1\x6f\xd0\x01\x4c\xed\xe2\xbf\xf9\xa0\x8a\x72\xc7\x4e\xfa\xb1\x9d\xf4\xd7\x76\xd2\x54\x50\x6e\xee\xd4\x15\x00\xd4\x9e\x0a\x06\x38\x66\x73\xb5\xa2\x00\xbb\xa1\xf0\xf8\xc2\x36\x0a\xc5\x54\x51\x33\xfa\x91\x0c\x13\xb2\x79\x78\x61\x17\x62\x37\xfb\x82\xe8\x6f\xa7\x5f\x3e\x53\x48\xc8\x64\x05\x09\x6c\x87\x43\x3d\xc7\x44\xe4\xa2\xcc\xb1\x64\x0e\xa8\xa2\x91\xf6\x1c\x0f\xbb\x0c\x16\x6a\x61\x30\x02\x51\xe7\x65\x8a\xcc\x94\x1f\x79\xcd\x36\x2f\x9a\xaa\x50\xa4\x46\x9a\xae\xd8\xee\x6c\xe0\x5d\x24\xb4\xf1\xb8\xa9\x15\x96\x6e\x29\x61\x49\x37\xad\x14\xbd\x22\xa1\x0d\x6d\xfc\xce\x64\xf7\x64\x68\x39\x60\x90\xcc\x49\x1f\x7a\xe7\xd7\x4e\x0b\xbb\x3b\xa6\x67\x1b\xb8\x33\x9e\xe3\xf0\xa2\xb6\x6b\xae\x96\xb8\x17\x7a\xe4\x5a\x49\x7a\x6d\xb1\x6e\x70\xd9\xb4\x7f\x3d\x90\xe5\x52\x65\xc7\xc9\xea\x55\xb6\x2d\x27\x77\x7b\xd3\x67\xc4\xc7\xda\xb3\xd8\xb4\xcd\x3f\x35\x4b\x3d\xd4\x69\x7e\x1b\x1a\xbe\x52\x77\x02\x57\xa5\x7b\xad\x60\x96\xaf\x54\x01\x7e\x39\x63\x80\x88\x86\x83\x8d\x05\x4d\x92\xbb\x05\x67\x6b\x59\xf1\x4f\x9b\xf3\xfd\xf7\xae\x82\x08\x81\xd8\xb0\x39\xf2\xf4\x7d\x6d\x52\x88\xaf\xdd\x76\x6e\x86\x14\xb1\x0c\xb7\x35\xd6\x0e\x3f\x23\xa9\xf8\xaa\xc6\x71\x00\xc5\xdd\x4a\xa3\x8e\xdc\xca\xa0\x26\x7c\xe5\x35\x91\xb7\x39\x4d\x30\x39\xb4\x3a\x6e\xf1\xf9\x8a\x96\x9b\x07\x1f\x8a\x72\xb3\xcc\xd3\x6b\x2c\x67\xee\xf8\x20\x9e\x99\xeb\x9c\x4a\x98\xac\x3d\x36\xa8\x77\x2d\xe9\x1c\xc1\x1e\xf5\x4a\xab\xc5\xb0\xed\xd0\xe8\x62\xc9\xeb\xd1\xb8\x45\xeb\x7b\xfc\xbb\x24\x73\x3e\xcd\x42\x1e\x74\x17\x5c\xbc\xf6\x51\xe2\x88\xee\xb2\x1a\x97\x79\x96\xdd\xfd\xfe\x8f\x59\x11\x8a\xd4\xe6\x79\x05\xdf\xf7\xf4\xe3\x57\x7c\xfc\x15\x0b\x96\x6f\xf3\x6a\xc8\xd7\x5e\xc7\xf8\x48\x5b\x52\x35\x07\xdf\x41\x30\x0e\xf0\x93\x26\xf8\xe5\x21\xcf\xd6\xc1\x70\x50\xea\xa5\x4d\xf4\xe7\xc7\xc8\xd1\x93\xb5\x99\xc5\xc7\xee\xbe\xed\x64\x72\x7a\xc6\xdb\xc8\x35\x1b\xf4\x39\x50\xbc\x36\x55\x10\x17\x09\xfe\x37\x14\x4d\xdd\x5f\x85\xde\x37\x33\x91\xc8\x7f\xe4\xbe\x97\x28\xf1\xa4\xfc\x09\xd5\xd6\xe2\x0a\x7d\x03\xa3\xe1\xe4\x1e\x30\x44\xab\xb5\x97\xc8\x64\xce\x45\x49\x09\x76\xc4\x86\x76\x5d\xc6\xd1\x3c\xc9\x0b\x22\xec\x28\x8f\xb4\x4a\xec\x07\xf9\xa6\x8d\x2f\x08\xb9\x9c\x3f\xaf\x14\x95\xf3\x65\x24\x43\xfd\xab\x0b\x7f\x4c\x64\xe4\x91\x9d\xcd\x6e\x18\xc8\xb7\x2b\x10\x7c\x47\xaa\x45\x51\xd8\x9a\xe2\xc2\x4d\x05\x97\xbc\x4b\x2f\x77\x89\xd4\x7e\xa8\xd3\x48\x5f\x73\x51\xfb\x1d\x5f\x04\x86\xad\x3e\x28\xb9\xd3\x26\xda\xb7\xc4\x8f\xad\x70\x33\xd8\xc8\xa5\xac\x14\xc9\xc7\x5f\x4e\xcf\xb6\xd0\xc4\xf7\x0e\x0e\x07\xd4\x89\xab\x0a\x80\x6e\xd9\x35\xf7\x91\xaf\x16\x30\xaa\xbd\x3b\xb8\xb6\x7a\xdd\x30\x68\x4a\x9c\x9d\x88\x35\x93\x3d\xcc\xf8\x98\x78\x15\x88\x7d\x98\x10\x60\x45\x0b\x55\x87\xc3\x26\x52\x72\xed\x4b\x9f\xbf\x5b\xa0\xb8\xbd\xda\x05\x12\x4d\xf4\x00\x72\xfd\xd6\x2b\x82\x68\x33\xa6\x3c\x20\x1c\x52\xae\xd8\x6a\x22\xaa\x5c\xbe\x5a\xa0\x44\xd2\xc7\x52\xa0\xfb\x96\x62\x17\x56\x32\xd5\x43\xab\xe3\xf7\x3f\x07\xeb\xf8\xbc\x1f\x4f\x1d\x0a\x0e\xa5\xf7\x4a\x2b\x0f\xa5\x8c\x5f\x5f\x1d\x50\x7b\x60\x11\x69\x77\xc2\x22\x53\xaf\x34\xcb\xbd\x9f\x7c\x9c\x9c\x4d\x7a\xc0\xf4\x0c\x8f\xd8\xc8\xff\x98\xb8\x7f\x37\xf9\x4f\x00\x00\x00\xff\xff\x2d\x61\x4d\x01\xfa\x23\x00\x00")

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

	info := bindataFileInfo{name: "rest/gorilla/pat.tpl", size: 9210, mode: os.FileMode(436), modTime: time.Unix(1447127619, 0)}
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

