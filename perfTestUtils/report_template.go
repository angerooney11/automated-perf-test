// Code generated by go-bindata.
// sources:
// report/content.tmpl
// report/footer.tmpl
// report/header.tmpl
// DO NOT EDIT!

package perfTestUtils

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

var _reportContentTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x57\xdb\x6e\xe3\x36\x10\x7d\xdf\xaf\x18\xa8\x31\x6c\x03\x85\x2f\x9b\x26\x40\x55\xdb\x80\xbd\x6d\x77\xb3\x4d\x0a\x63\x9d\xf6\xa5\xe8\x03\x2d\xd1\x36\x1b\x59\x52\x29\x3a\x89\xab\xea\xdf\x3b\xd4\xcd\xba\x90\xb2\xb2\x68\x16\x2b\xc0\x80\x25\x0e\x87\x87\x33\x67\x86\x87\x61\x68\xd3\x0d\x73\x29\x18\x96\xe7\x0a\xea\x0a\x23\x8a\xde\x00\x4c\x6c\xf6\x08\x96\x43\x82\x60\x6a\x08\xcf\x5f\x10\x6e\xcc\xf0\xf3\xe9\x99\xec\xc6\xd9\xb8\x4f\x6c\x9b\xb9\x5b\x63\x16\x86\x83\x77\x9e\xbb\x61\xdb\xc1\x7c\x79\xf3\x2b\xd9\xd3\x28\x02\xd3\x84\xf9\x41\x78\x7b\x22\xa8\x0d\x4b\xca\x37\x1e\xdf\x13\xd7\xa2\x70\x4f\x03\x01\x9f\xa8\xef\x71\x21\x8d\x7a\x38\x59\x0e\xaf\x04\x11\xc1\xe0\x3d\x15\x72\xfc\x9e\x49\x1f\xfd\xc9\x70\x37\x3e\xad\x3e\x19\x22\xb6\xc2\x6b\x01\x29\xfe\xfd\x40\x49\x02\xe6\x84\x76\x22\xc8\xda\xa1\x0a\x1b\x58\x7b\xdc\xa6\x7c\x6a\x8c\x0c\x78\x62\xb6\xd8\x4d\x8d\xef\x47\x9d\xc2\xd4\x89\xe0\xe5\x5d\x97\x22\x20\xec\x6c\xd6\x95\x9c\x35\xd9\x5d\xd7\x22\xf2\xc1\xc3\x4d\x1e\x5c\x5c\x04\x04\xee\xc7\x84\x53\x88\xee\x09\xdf\x52\x21\x0d\xa2\xc8\xac\x7e\x5e\x62\x54\xa2\x08\xf7\x7d\x3d\x9b\x0c\x85\xad\x07\xd1\x00\xea\xed\x95\x06\xd4\x8a\xf2\x47\x66\xd1\xa0\x02\xcc\xa1\x2e\x14\x52\x90\x5a\x7d\xa2\x81\xef\xb9\x01\x95\xa9\x08\x5e\x0d\xd2\x19\xb7\x38\xa2\x49\x04\x8e\xc8\xe4\xea\x06\xdb\x30\x05\x3c\xd7\x72\x98\xf5\x30\x35\x76\xcc\xa6\x77\x74\xef\xf1\xe3\xdc\x25\xce\x31\x60\x41\xaf\x5f\xa5\xfd\x67\x73\xe9\xe4\x41\xb7\x95\x1a\x9f\x2e\x6b\x71\x4a\xd0\x41\x06\x0f\xc3\x76\xd9\x14\xb6\x16\x91\x9f\x3b\x8e\xf7\x84\xa5\xf9\x3b\xe1\x2c\xae\xcb\x22\x47\xe3\x41\xb9\xe1\x25\x25\x0f\xc9\xda\xb9\xdd\xbf\xe0\x73\xe6\x8a\x0d\x18\x9d\xef\x06\x6f\x37\xd8\x36\x3a\xe7\xb2\xd8\x02\x4e\x18\xb2\x0d\x0c\x6e\x82\x64\xb1\x25\x0e\x02\x92\x6e\x83\xad\x09\x2c\xcf\xf1\x30\xbe\x5b\x4e\xa9\x6b\xcc\x96\xf3\xd5\x6a\x32\x94\x03\x38\x87\x3a\x01\xad\x98\x71\x6a\x1b\xb3\x9f\xe7\x37\xb7\x27\x23\xd7\x6e\xe4\x6f\x9d\x64\x35\x72\xa9\x08\xc5\xec\xa9\xb1\x8f\xd1\x62\xd0\x04\xc1\x56\x5a\x6b\x95\xc5\x5e\x1a\x47\x33\xdb\xad\x82\x1c\x45\x7a\x65\x51\x69\x26\x54\x32\xef\x4c\xa3\xca\x5c\x8c\xaf\x47\x18\xf8\xc5\x6c\x41\x02\x0a\x32\xab\x90\x44\xda\x9c\x0c\x17\x67\xea\x1a\xdd\xc8\x0e\x2f\x67\x9e\xba\x44\xf2\x96\x91\x03\x49\x81\xa1\xb8\xf7\xee\x16\xf8\x2f\xee\xf4\x02\x07\xa2\xe8\x6e\x71\xd6\x75\x0e\xf0\x4a\x02\x5c\xcf\xe2\xe3\xa1\x0c\x70\xdd\x0e\xe0\x09\xdc\xff\x0b\x6c\x9c\x00\xeb\xe4\xa5\xd2\x0e\x12\xc4\x94\x76\x3d\x51\xa6\x75\x14\x05\xe2\xe8\xd0\xa9\x11\xf3\xd5\x94\x74\x4d\x29\x9a\xec\xa1\x5a\x6f\xb8\x2b\x0b\x4f\x67\xb2\xa5\xe5\x1d\x74\x5e\xda\x35\x95\x1d\x33\x25\xb6\x8e\xb6\x5d\x2b\xa3\x76\x57\xe1\xb0\x68\xb7\x23\x5c\x28\x6c\x72\x3b\xac\x96\xee\x2d\x3a\x7a\x97\x18\x56\x0a\xaa\x02\xe7\xdc\xa7\xc0\xe2\xcc\x17\xf5\xe9\x8f\x84\x43\xbe\xc8\xc7\x15\x4c\xc1\xba\x1c\x6c\x29\xc2\x47\x15\xd2\x0b\x6b\xf6\x36\x11\x04\xdb\x9e\x12\x35\xa6\xe7\xb0\x77\x03\x13\xfe\xd0\x25\x39\x0c\xff\x0a\x3c\x17\x93\x01\x86\xac\x06\x03\x2a\x25\x92\x9e\x28\x07\x9b\xe1\x91\xfe\x6d\x0b\x2f\x92\xfa\x46\xf1\x28\x2e\x79\x50\x3a\xf8\xb3\xf6\x55\xb1\xd2\x3f\x9e\xb7\xd7\x6d\x93\xba\x92\x14\xb6\x09\x82\x1f\x68\x1b\x67\x10\x1c\xd6\x71\xb6\x75\x1e\x83\x9d\xf7\xd4\xde\x1d\x79\x66\x81\xce\xd3\x51\x37\x20\x1f\x87\xac\xa9\x63\x42\x37\xad\xf3\xde\x2f\x8b\x7e\x57\x69\xac\x58\x54\xf1\x69\xcb\x99\xad\x5d\xee\xb9\x11\x08\x12\xae\x89\x26\xa5\x07\x4b\xfc\x23\xa6\x5b\x66\x7a\x89\x31\x64\x82\xa1\xc6\xd2\xe4\xb6\xf8\xd4\xf3\x9c\x3f\xea\xc9\xe5\xaf\x51\xff\x87\xb2\xd5\x45\xcf\xf8\x26\xaf\x14\xa3\x3f\x20\xbe\x8f\x3d\xa8\x57\x28\x9e\x01\x75\xe8\x1e\x1b\x4f\x65\xe6\x64\x58\x2d\xbe\x6a\x0b\x69\xa5\xb6\x52\x91\xf9\xb5\xca\xad\x14\x1e\x64\x22\x18\xa4\x0a\xfe\xf2\xea\x4b\x21\xc5\x5f\x59\x86\x95\xa4\x54\xa6\xc9\xe4\xba\xc9\xd1\x55\xd0\x60\x99\xf8\xca\xf5\x56\x2e\xb4\x62\xd9\xf5\x3a\x7a\x2b\x48\x02\xa2\x12\x5c\x2d\xc5\x56\x4a\xac\x73\xd4\x81\xf4\x90\x5e\x13\xeb\x61\xcb\x3d\xbc\x30\x99\xb7\x6c\xbb\x13\xef\x39\x39\x6a\xa5\x58\x1e\xde\xcb\x98\x57\x89\x96\x91\xd7\xe1\x66\xbd\x50\xcc\x4b\x3a\x51\x1e\x24\x31\xe5\x7a\x77\xcc\x71\x58\xff\xc5\x0e\xb2\x4b\xf4\x67\x3b\xe8\x3c\xa6\x44\xd3\xcf\x54\xcb\x8c\x30\xe4\xc4\x45\xb1\x72\x21\xaf\x97\x60\x4e\x8b\xe7\x59\xab\xb6\x17\x86\x17\x6b\x29\x52\x71\x2a\xc3\x7b\xea\x33\x5c\x28\xa4\xa7\xea\x92\x9a\x2c\x39\xc8\x82\xae\x77\x4f\x1e\xb7\x45\xef\xcd\x57\xdf\x8a\x57\xd0\xba\x95\x62\x8f\x56\xad\x0d\x74\xcc\xc5\x6f\xbe\xd1\x34\x8d\xfe\x0d\x31\xa6\x91\xc6\x48\xf2\x71\x47\x25\xfd\xa6\xe3\x91\xff\x3c\x4b\xf4\x6e\x75\xb7\x71\x92\x92\x21\x59\x0a\x49\x10\xc7\x23\x7a\x5d\xd6\x8d\xb9\x99\xac\xd9\x9f\x7e\xcc\x5e\xa1\xa6\x4a\x4b\x06\xda\x7b\x78\xd6\x06\xbe\x20\xf0\xd4\x4c\x46\x4c\x6f\x95\x44\xf6\x02\xbb\x57\x9a\xd4\xac\x89\xd5\x58\x92\xed\xa0\x41\x96\xe3\x52\xd8\x77\x93\x25\x13\x78\xca\xee\x7b\x2e\x50\xd2\xdb\x8b\xc6\x54\xdf\x5b\xf5\xca\x33\xea\xbd\x8d\x72\xcf\x55\xfb\x9a\x70\xbd\x68\xaf\x2e\x5e\x79\x55\x08\x75\xa9\xd1\x33\x97\xaf\x2e\xd1\x73\xbd\x85\xc9\x9f\x73\xec\xdd\x3a\x19\xad\x96\xe7\xe2\xe8\x53\x94\x98\x08\xb7\x2e\x2d\x5f\x5b\x67\xe3\xa2\x3a\x5f\x71\xab\x6e\x92\xa4\x18\x4b\xe6\x99\x30\x1a\x5c\xb5\xd0\x86\x9a\xf5\x75\xc2\x5c\x2b\xca\x33\x41\x3e\xc7\x22\x91\x1d\x14\x0a\xe7\xcf\x8a\x22\x15\xed\x40\xa1\xd0\x35\x17\xa3\x46\xc5\x9d\xa6\xc5\x42\xc6\x6c\x51\xfb\x77\xb5\x77\x2b\x48\x4d\x98\x14\xe8\x05\x2a\xa4\x2d\x41\xd6\xbf\xee\x14\x52\xc4\xa8\x51\x51\x4b\x41\x9d\xd1\xfa\xa4\xa7\x4f\x44\x57\xc8\x69\xad\x94\xce\x5f\xd7\x7c\xd6\xf4\x7b\x93\xf6\x87\xff\x02\x00\x00\xff\xff\x76\x7a\xfc\x44\xcb\x17\x00\x00")

func reportContentTmplBytes() ([]byte, error) {
	return bindataRead(
		_reportContentTmpl,
		"report/content.tmpl",
	)
}

func reportContentTmpl() (*asset, error) {
	bytes, err := reportContentTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "report/content.tmpl", size: 6091, mode: os.FileMode(420), modTime: time.Unix(1455026397, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _reportFooterTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\x4a\xcb\xcf\x2f\x49\x2d\x52\xaa\xad\xe5\x52\xb0\xd1\x4f\xca\x4f\xa9\xb4\xe3\xb2\xd1\xcf\x28\xc9\xcd\xb1\xe3\xaa\xae\x4e\xcd\x4b\xa9\xad\x05\x04\x00\x00\xff\xff\x70\x9a\x96\xda\x2c\x00\x00\x00")

func reportFooterTmplBytes() ([]byte, error) {
	return bindataRead(
		_reportFooterTmpl,
		"report/footer.tmpl",
	)
}

func reportFooterTmpl() (*asset, error) {
	bytes, err := reportFooterTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "report/footer.tmpl", size: 44, mode: os.FileMode(420), modTime: time.Unix(1455026397, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _reportHeaderTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x57\x5b\x6f\xe2\x38\x14\x7e\xe7\x57\x78\xb2\x2f\x1d\x69\x9d\x0b\xa1\x0c\x61\x29\x52\x3b\x45\xdb\x4a\xdd\x9d\x6a\xcb\x4b\x9f\x56\x4e\x7c\x4c\xcc\x24\x76\xd6\x36\x74\x18\xc4\x7f\x5f\x27\x74\x85\x12\xd2\x16\x7a\x91\xd6\x52\xa4\x38\x3e\xfe\xbe\xef\x5c\x72\xe2\xac\xd7\x14\x18\x17\x80\x9c\x14\x08\x05\xe5\x6c\x36\x9d\xd1\xa7\xcb\x6f\x5f\xa7\xf7\xb7\x13\x94\x9a\x3c\x1b\x77\x3a\xc8\x8e\xd1\xd5\xe4\xfc\x72\x5c\xdd\x56\xd3\xe9\xf5\xf4\x66\x32\x5e\xaf\xdd\xaf\x52\x30\x3e\x73\xcf\x6f\xaf\xff\x24\x39\x6c\x36\xe8\x7c\x61\x64\x4e\x0c\x50\x74\x0b\x8a\x49\x95\x13\x91\x00\x9a\x82\x36\xe8\x2f\x28\xa4\x32\x23\x6f\xbb\xb9\xd3\xd9\xc1\xe9\x44\xf1\xc2\x20\xad\x92\x33\x27\x35\xa6\x18\x7a\x5e\x12\xce\xb5\x2b\xd5\xcc\x9b\x6b\x6f\xfe\xcf\x02\xd4\x0a\x07\x6e\x10\xb8\xbe\x9b\x73\x81\xa3\xc0\xf7\xfb\x7d\x16\xbb\x73\xed\x20\xb3\x2a\xe0\xcc\x31\xf0\xc3\x78\x73\xb2\x24\x5b\x30\x67\x3c\xf2\xb6\x77\xe3\x83\x79\x68\x88\x43\xf7\xf4\x91\x22\x8c\x4f\xfb\xbd\x30\xea\xbe\x2f\x45\x12\x56\xe0\x41\x37\x0a\xba\x2c\xee\x1f\x09\xfe\x09\xe3\x97\x08\x98\x5c\x08\x4a\x0c\x97\x62\x4b\x44\x19\x0c\x82\xc0\x3f\x88\x08\xe3\xe3\xb8\x52\x3e\x4b\x33\x7b\x19\xb7\x20\xc9\x77\xdc\x23\xec\xd4\xf7\x7b\xf4\x43\xb8\x96\x20\xa8\x54\x5e\x2e\x6d\x99\x0a\xfe\x53\xe1\xae\xdb\x77\x83\xca\xc7\xfe\x80\xd1\x24\x89\xa2\x83\x79\x77\xc4\x19\x17\xdf\x51\xaa\x80\xed\x73\x26\x5a\x7b\xa2\xac\xdf\x8c\xff\x04\x0c\xbd\xfe\x69\x12\x0f\xfa\xae\x7d\xec\xa0\x1c\x28\x27\x67\x8e\x85\x04\x10\x0e\x52\x90\xd9\x89\x59\x65\xa0\x53\x00\x53\x13\x51\xd9\x7b\xe3\x83\x29\x1b\xf9\x8b\xbe\x0c\x68\x2f\x81\xc1\x87\xf3\x26\x21\x8e\xfd\x30\xe8\x9e\x32\xf2\xe1\x5c\x15\x0c\x8e\x22\x16\x0f\xa2\x41\xf4\x76\xba\xb2\x7c\x5e\xa2\x24\x79\x61\xb1\xbc\x24\x25\xca\xfc\xad\x0b\x6b\x0e\x98\x92\x30\x22\x21\xc0\x6b\x15\xec\xd5\xf0\x0b\x22\x6c\x67\x94\x4a\xc9\x07\x4c\xbf\x24\xcc\x8f\xba\xc1\x5b\x88\x77\xcc\x77\xd3\xfb\x9b\x49\xd3\x6a\xa7\xac\x1c\x17\xdf\x2e\xef\xd1\xba\xf6\xa8\x1c\xb1\x7d\x6d\x67\xaa\xac\xb8\x21\x7a\x48\xb9\x81\xdf\xf6\x4c\x12\x99\x49\x35\x44\x71\x66\x4d\x6b\x8b\x9b\x4e\x6d\xea\x1a\x59\x5c\x10\xb5\xcf\x51\x10\x4a\xb9\x98\x61\xbb\x3e\x44\x38\xf0\x8b\x1f\xfb\x24\x3b\x1d\xf8\x91\xef\xa6\x6c\x2b\x77\x99\xfd\x90\xfc\xae\xc8\xaa\x65\x87\x54\xb6\x11\xe0\x2a\x4c\x43\x64\xf7\xca\x65\x43\x7c\x43\x5f\x1a\xb8\x8f\x4a\x5a\xe2\xf0\x9f\xc6\x0c\x98\x19\xa2\x41\xab\x46\x26\x85\x5d\x0b\xed\x1a\x22\x8a\x93\xec\x57\xa4\x89\xd0\x58\x83\xe2\xec\xc9\xb0\x3d\x11\xd4\x32\x4b\x58\xa7\x84\xca\x87\x21\x0a\x2c\x62\x79\x55\x2e\xef\x7b\xdb\xf4\x23\x7c\xa3\x1f\xef\xa9\xac\x26\xac\xff\x3f\x12\x56\x2f\x4d\xca\x97\x57\xf6\x7c\x63\x05\x3c\xf7\x0a\xd4\x4a\xef\x22\x5b\xb4\x88\xa8\x57\x9d\x90\xe2\xf9\x9a\x73\x0d\x89\x33\xb8\x7d\x75\x54\x1a\x70\xbf\xe4\x60\xfb\xc7\xca\x1e\xb9\x0c\xb1\xed\x4b\xb5\x20\xa6\x50\x8a\x1f\xf6\xfc\xe7\xb0\x46\x5e\xd5\x31\xf6\x0e\x2d\xf5\x96\xb1\x24\x0a\x71\xfd\x47\x45\x79\x2e\x48\xb6\xd2\x5c\x5f\x71\x4a\x41\xa0\x33\xc4\x48\xa6\x1b\xbe\xb3\x85\x48\xca\xcf\x16\x4a\x39\x85\xfa\xb6\x93\xcf\xfb\x4a\x39\x3b\x69\x47\x6f\xb1\x2d\x07\x95\xc9\x22\x07\x61\xdc\x19\x98\x49\x06\xe5\xed\xc5\xea\x9a\x9e\x38\x8d\xa0\x38\x9f\xdd\x2a\x3d\x36\xe9\xb6\xcb\x93\x95\x15\xeb\x70\x51\xb6\x7b\x67\x3f\x9f\x95\x90\x23\x7c\xac\x02\x09\xf6\xf1\xfb\x6b\x2c\x8b\xe9\x58\x85\x46\xb5\x15\x69\xa3\x6a\x1a\xd3\x6d\x5a\xef\x40\x2d\x79\x02\x47\xe7\xb5\xb1\xef\xc9\xc4\xb6\xe2\x1f\x9b\x59\xbd\x05\x79\x4b\x6a\x8f\xf0\xb3\x8a\xd5\x2b\x72\x7b\x88\xca\x67\x93\xfb\x94\xc6\x63\xb3\xbb\xfb\x5f\xd8\xfe\xb2\x79\xdb\x7f\xb6\xed\x24\x96\x74\x35\xee\xac\xd7\xf6\x14\xbd\xd9\xfc\x1b\x00\x00\xff\xff\x71\xac\x47\x54\xf7\x0d\x00\x00")

func reportHeaderTmplBytes() ([]byte, error) {
	return bindataRead(
		_reportHeaderTmpl,
		"report/header.tmpl",
	)
}

func reportHeaderTmpl() (*asset, error) {
	bytes, err := reportHeaderTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "report/header.tmpl", size: 3575, mode: os.FileMode(420), modTime: time.Unix(1455026397, 0)}
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
	"report/content.tmpl": reportContentTmpl,
	"report/footer.tmpl":  reportFooterTmpl,
	"report/header.tmpl":  reportHeaderTmpl,
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
	"report": &bintree{nil, map[string]*bintree{
		"content.tmpl": &bintree{reportContentTmpl, map[string]*bintree{}},
		"footer.tmpl":  &bintree{reportFooterTmpl, map[string]*bintree{}},
		"header.tmpl":  &bintree{reportHeaderTmpl, map[string]*bintree{}},
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