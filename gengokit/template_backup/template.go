// Code generated by go-bindata. DO NOT EDIT.
// sources:
// NAME/cmd/NAME/main.gotemplate (356B)
// NAME/handlers/handlers.gotemplate (62B)
// NAME/handlers/hooks.gotemplate (62B)
// NAME/handlers/middlewares.gotemplate (75B)
// NAME/svc/client/grpc/client.gotemplate (3.184kB)
// NAME/svc/client/http/client.gotemplate (105B)
// NAME/svc/endpoints.gotemplate (4.272kB)
// NAME/svc/server/run.gotemplate (3.259kB)
// NAME/svc/transport_grpc.gotemplate (2.962kB)
// NAME/svc/transport_http.gotemplate (106B)

package template

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _cmdNameMainGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x90\x31\x6f\x83\x30\x14\x84\x67\xde\xaf\x38\x31\xc1\x50\xbc\x77\x0d\x1d\xb2\x34\x51\x9a\x76\x77\xe0\x00\xab\xc4\x20\xdb\x10\x45\xc8\xff\xbd\x82\x44\x55\xa6\x77\x4f\x77\xba\x4f\xef\x29\x85\xdd\x50\x13\x2d\x2d\x9d\x0e\xac\x71\xb9\x23\xb8\xc9\xfb\x02\xe5\x01\x9f\x87\x33\x3e\xca\xfd\xb9\x10\xa5\x70\xa2\x9b\xac\x35\xb6\x7d\x04\x70\x33\x7d\x8f\x61\xa6\xbb\x39\x13\x88\xd0\x19\x8f\xc6\xf4\xdc\xc2\x3f\x74\xde\x0c\xf6\x1d\xcb\x52\x3c\x75\x8c\x2f\x06\x4a\x1d\xf8\xea\xae\x7b\x8c\x22\xa3\xae\x7e\x75\x4b\x5c\xb5\xb1\x22\xe6\x3a\x0e\x2e\x20\x93\x24\x6d\x7a\xdd\xa6\x22\x89\x52\x38\xaf\xa8\x2f\xba\xd9\x54\x94\x24\x5d\x96\x62\xbf\xe5\x8e\x3a\x74\x78\x8b\x11\xca\xcf\x95\xf2\x74\x33\x5d\x2a\xb9\x48\x33\xd9\x6a\x6b\xcc\x72\x2c\x5b\xc5\xf7\x58\xeb\x40\xe8\xba\x76\xf4\x9e\x1e\xa6\x41\xe8\x78\x47\xa7\x67\xe2\x42\xda\xff\xd3\x02\xed\xfa\x95\x95\xef\x25\x59\x47\x71\xd4\xce\x33\xcb\x45\x92\x07\xa4\x38\x4d\x36\x7b\xca\x92\x8d\x9e\xfa\xb0\x1b\x6c\x63\xda\x5c\xa2\xfc\x05\x00\x00\xff\xff\x49\xa3\x9c\x96\x64\x01\x00\x00")

func cmdNameMainGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_cmdNameMainGotemplate,
		"cmd/NAME/main.gotemplate",
	)
}

func cmdNameMainGotemplate() (*asset, error) {
	bytes, err := cmdNameMainGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/NAME/main.gotemplate", size: 356, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3c, 0xd, 0xff, 0x1e, 0xe5, 0xa9, 0x62, 0x3d, 0xb4, 0x16, 0xf5, 0xfd, 0x84, 0x56, 0x3e, 0xc4, 0x7d, 0x9e, 0x50, 0x8f, 0x7b, 0xc0, 0xcc, 0x27, 0x7f, 0xe7, 0x57, 0x20, 0x66, 0xd, 0x77, 0x87}}
	return a, nil
}

var _handlersHandlersGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\xc1\x0d\x80\x20\x0c\x05\xd0\xbb\x53\xf4\x4c\xa2\x1d\xc6\x09\x88\x7c\xab\x11\xa9\x69\xeb\x89\xb0\x3b\xaf\x77\x4e\xb4\x03\x24\xba\x86\xfd\xee\x2c\x68\xa2\xcf\x1d\x7c\xe5\x56\x2a\x8c\x03\xef\x57\x73\xc0\x37\x51\x3a\xd5\xe8\xd0\x02\x4a\x3c\xc6\x32\x03\x00\x00\xff\xff\xd6\x21\xab\x2e\x3e\x00\x00\x00")

func handlersHandlersGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_handlersHandlersGotemplate,
		ServerHandlerPath,
	)
}

func handlersHandlersGotemplate() (*asset, error) {
	bytes, err := handlersHandlersGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ServerHandlerPath, size: 62, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1e, 0xcb, 0xd5, 0x72, 0x80, 0xc6, 0xf9, 0x82, 0x4b, 0xe0, 0x8b, 0x90, 0xb8, 0x9b, 0xbc, 0x5d, 0x8d, 0x12, 0xd4, 0x8e, 0x54, 0xf6, 0x72, 0xcb, 0xef, 0xf5, 0x12, 0xd0, 0xe1, 0xb8, 0x41, 0xc8}}
	return a, nil
}

var _handlersHooksGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\xc1\x0d\x80\x20\x0c\x05\xd0\xbb\x53\xf4\x4c\xa2\x1d\xc6\x09\x88\x7c\xab\x11\xa9\x69\xeb\x89\xb0\x3b\xaf\x77\x4e\xb4\x03\x24\xba\x86\xfd\xee\x2c\x68\xa2\xcf\x1d\x7c\xe5\x56\x2a\x8c\x03\xef\x57\x73\xc0\x37\x51\x3a\xd5\xe8\xd0\x02\x4a\x3c\xc6\x32\x03\x00\x00\xff\xff\xd6\x21\xab\x2e\x3e\x00\x00\x00")

func handlersHooksGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_handlersHooksGotemplate,
		"handlers/hooks.gotemplate",
	)
}

func handlersHooksGotemplate() (*asset, error) {
	bytes, err := handlersHooksGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "handlers/hooks.gotemplate", size: 62, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1e, 0xcb, 0xd5, 0x72, 0x80, 0xc6, 0xf9, 0x82, 0x4b, 0xe0, 0x8b, 0x90, 0xb8, 0x9b, 0xbc, 0x5d, 0x8d, 0x12, 0xd4, 0x8e, 0x54, 0xf6, 0x72, 0xcb, 0xef, 0xf5, 0x12, 0xd0, 0xe1, 0xb8, 0x41, 0xc8}}
	return a, nil
}

var _handlersMiddlewaresGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\x41\x0e\x02\x21\x0c\x05\xd0\xbd\xa7\xe8\x9a\x44\x7b\x18\x4f\x40\xec\xb7\x12\x81\x4e\xda\x4e\x66\x41\xb8\xfb\xbc\xb5\xb8\xd0\x1b\x20\xb5\x67\xfa\x19\xc1\x8a\xa9\xf6\x6f\xc9\xbf\x3a\xa5\xc3\x83\x13\xe3\xe8\x35\x11\x3c\x9a\x48\xc7\x55\x1d\xf1\x52\xa3\xaf\x39\x7d\x4c\x40\x85\xf7\x7e\xdc\x01\x00\x00\xff\xff\xcf\x9e\xe9\x81\x4b\x00\x00\x00")

func handlersMiddlewaresGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_handlersMiddlewaresGotemplate,
		"handlers/middlewares.gotemplate",
	)
}

func handlersMiddlewaresGotemplate() (*asset, error) {
	bytes, err := handlersMiddlewaresGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "handlers/middlewares.gotemplate", size: 75, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xcc, 0xfe, 0x9d, 0x1a, 0xaf, 0x47, 0xe8, 0x97, 0x82, 0x24, 0x50, 0x17, 0xb4, 0x49, 0x73, 0x3b, 0x68, 0xb7, 0xe5, 0x3a, 0x3d, 0xb6, 0x15, 0x9d, 0xb1, 0x8f, 0xc4, 0x27, 0xaf, 0xa7, 0x3c, 0xc1}}
	return a, nil
}

var _svcClientGrpcClientGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x56\x4d\x6f\xdb\x38\x10\x3d\x8b\xbf\x62\xd6\x08\x16\x52\xa0\xd0\xf7\x2c\x7c\xa9\xd3\x2d\xba\xd8\xa6\x46\x1a\x74\x0f\x45\x51\x30\xd4\x58\x26\x2c\x93\x2a\x49\x3b\x31\x04\xfd\xf7\xc5\x90\x94\x23\x27\x8e\xdb\x43\x10\x8b\xf3\x38\x1f\xef\x0d\x39\x9c\x4e\x61\x6e\x2a\x84\x1a\x35\x5a\xe1\xb1\x82\x87\x3d\x78\xbb\x75\x8e\xc3\xcd\x67\xb8\xfd\x7c\x0f\xef\x6f\x3e\xde\x73\x36\x9d\xc2\x1d\xda\xad\xd6\x4a\xd7\x11\x00\x8f\xaa\x69\xc0\xec\xd0\x3e\x5a\xe5\x11\xfc\x4a\x39\x58\xaa\x06\x03\xf8\x2b\x5a\xa7\x8c\xbe\x86\xae\xe3\xe9\x77\xdf\x8f\x0c\x70\x23\x3c\x8e\xad\xf4\xdd\xf7\x8c\x20\x0b\x21\xd7\xa2\x46\xa8\x6d\x2b\xa1\xb5\x66\xa7\x2a\x74\x20\xa0\xbe\x5b\xcc\x41\x36\x0a\xb5\x87\xa5\xb1\xe0\x57\x48\x0e\xbe\xa0\xdd\x29\x89\xfc\x56\x6c\xb0\xef\xc1\xa5\x4f\xd6\x8e\xdc\x30\xa6\x36\xad\xb1\x1e\x72\x96\x4d\xa4\xd1\x1e\x9f\xfc\x84\x65\x93\xda\x98\xba\x41\x5e\x9b\x46\xe8\x9a\x1b\x5b\x4f\x09\xfd\xb6\x65\xba\x41\x2f\x2a\xe1\x45\x80\x28\xbf\xda\x3e\x70\x69\x36\xd3\x76\x5d\x4f\xd1\x5a\x63\xdd\x84\x1d\x5b\x6a\x73\xb5\x56\x7e\x4a\x7f\xa8\xab\xd6\x28\x4d\x81\xc9\x97\xb7\x42\xbb\x90\xd4\x1b\xf8\x03\x20\x25\xc5\xb2\xe9\x14\xee\x89\xe6\x54\x32\xcb\x26\x5d\xc7\x3f\x86\xca\x16\xc2\xaf\xe0\xaa\xef\x61\xea\x76\x54\x40\xfb\x00\x64\x5c\xbc\x3b\x36\x4f\x58\x11\x38\xbe\xc5\x47\xb0\xe8\xb7\x56\x3b\x10\x7a\x20\x0d\x1e\x84\x5c\xc7\x26\x38\xa6\x5b\x1a\xad\x51\x7a\x65\x34\x87\x8f\x1e\x94\x23\xf2\xc9\x8f\x45\xd7\x1a\xed\xd4\x83\x6a\x94\xdf\x83\x59\x06\x55\xa4\x68\x1a\xb4\xe0\x0d\x54\x4a\x34\x25\x08\x5d\x41\x23\x3c\x5a\x90\x8d\x71\x58\x46\xd0\xb3\x4f\xb6\xdc\x6a\x49\x39\xe5\xb4\x08\x97\x54\x2f\x9f\x87\xd0\x73\xa3\x75\x09\xa6\x25\x9c\x03\xce\xd3\xf2\xe7\xb0\x50\x40\xde\x3e\xf0\x57\x3d\x40\x5f\x68\x4b\x08\x8a\x14\xd0\xb1\x6c\x27\x2c\x48\x99\xaa\x99\x1b\xbd\x54\x35\x63\x19\x35\xd1\x8f\x12\x96\x70\x3d\x03\x2b\x74\x8d\x87\x38\x1d\xcb\x32\xb4\x96\x0c\xcb\xfc\x4f\x29\x0b\x96\x65\x6a\x49\x0e\xe1\x8f\x19\x68\xd5\x04\x44\x16\x19\xa4\xef\x14\xcc\xf1\xff\xac\x68\x73\xb4\xb6\x84\x89\x14\x5a\x1b\x0f\xa2\x6d\x9b\x7d\xf2\x3c\x21\x47\x3d\xcb\x7a\xc6\x32\x39\x2a\xc4\x51\xa4\x6f\xdf\x8f\xda\xe2\xa8\x52\x0a\x77\xca\xfa\x0e\x97\xc6\x62\x4e\xc9\xa4\xb6\xfe\x2a\x9a\x2d\xba\x7b\xf3\xe1\x6e\x31\xff\x94\xba\x35\x97\x92\xaf\x50\x54\x68\x5d\x51\x94\x14\x3e\xeb\xba\x2b\x78\x54\x7e\x05\x17\x1e\x29\x38\xef\x7b\x96\x8d\x56\xdb\x75\x4d\x64\x92\xe9\xc2\x23\x4f\x67\x32\xf2\x4b\xd1\x08\x19\x39\xbb\x50\x03\x68\x50\xe1\x13\xfa\x95\xa9\x5c\x04\x06\xee\xbb\xee\xde\xfc\x6b\x1e\xd1\xc2\x85\x4a\x22\xbd\x4f\xa7\x01\x86\x63\xc1\x87\x95\xb0\x2b\xf0\x4b\x61\xde\xde\x38\x83\x63\x46\x6e\xf1\x31\x92\x92\xc7\xbd\xc4\x88\x2e\xd3\xef\x49\xd7\x0d\x35\xf5\x3d\xef\xba\x71\xbe\x71\x71\x32\x86\xaa\x97\x8b\xef\xb5\x34\x15\x12\xa9\x23\xeb\x1d\xfe\xdc\xa2\xf3\x03\xe6\x06\x4f\x62\xc2\x09\xc1\x01\x14\x1a\xf6\x83\x09\xe4\x5e\x28\x3e\x98\xef\xf7\xed\x90\x48\xd7\x0f\xd8\xa3\x16\xe1\x9c\xa7\xf5\xe2\x40\x55\x5e\x84\x95\xa4\x08\xea\x2a\xa9\x98\x7e\x0d\x3f\xd8\xd0\xa9\x6e\x27\x0f\x7b\x5d\x47\x80\xb1\x86\x2f\x05\xa4\x0b\x23\xb8\x7b\xc5\xfd\x35\x00\x9c\x13\xb5\x7c\x8e\x9d\xf5\x25\x1d\x10\x16\xef\x76\x22\x07\xa2\x4a\x10\xe9\x62\xe7\x73\x88\x53\xe3\x2c\xb3\x74\x1d\x09\x38\xbe\x2d\x79\xdc\x31\x40\xfe\xa6\xfb\xc5\xaf\x44\xb8\xc9\x76\x68\xbd\x03\x41\x7e\xc3\x1d\x77\xa2\x0e\xb0\x48\x87\xd6\x1b\x10\xb0\x75\x68\xaf\x2a\xb3\x11\x4a\xbf\x01\x8d\x31\x38\x2c\xac\xda\x08\xab\x9a\x3d\xed\x59\x6e\x1b\x50\x1a\x44\xba\x74\xd2\x1d\x77\xb6\x90\xfc\x07\xa4\x43\xcc\xe7\xf1\x7f\x19\x5a\xfc\x2e\x24\xa3\xb4\x47\xbb\x14\x12\xbb\xbe\x80\x7c\xf4\x35\xbe\xe8\x62\xde\xd7\xb3\xe7\x7d\x3c\xbf\xfc\x75\xcb\x15\x87\x0e\x09\x0e\x06\xc5\x0e\xfd\xf3\x42\xb9\x78\x18\x7e\x4b\xb9\x73\xe7\xe6\xa4\x70\x71\x43\x42\xbc\xa5\xdb\xaf\x35\x89\x01\x82\x80\x67\x44\x0e\xa8\xdf\x12\xee\x5c\x1d\xa7\x74\x1b\x32\xf8\x4d\xd5\x7e\x86\x19\x94\xf2\x39\xa1\x58\x30\xbc\x21\xd8\xcf\x57\x72\x31\xbf\x6f\xf1\x68\xda\x81\xf3\x76\x2b\x3d\x05\x4b\x83\x00\xbe\x7d\x77\xde\x2a\x5d\xa7\x93\x39\x9e\x36\x51\x18\xaa\x3b\x7c\x05\x01\x36\xa6\x52\x4b\x85\x2e\xce\xee\xc3\xb3\x80\x26\x69\x88\x76\xb4\x9f\xb6\xe6\x97\xe3\x04\x8a\x58\x2e\x8b\x6c\xce\xfd\xd3\x30\xa7\xbe\xa0\xae\xf2\x35\xee\xc3\x70\x8f\x19\x15\xc7\xce\xba\x43\xad\xc1\xad\x81\x53\x8e\xc3\x40\x36\xc3\x94\x83\x19\x90\x4b\x36\x1e\xd1\x34\xf6\xfa\x14\xff\xdc\xac\x0c\xb9\x0c\xe4\x14\x70\x6a\xea\x8e\xbb\xf3\x45\x76\xd2\x3f\xbd\x6e\x86\x4d\x05\x97\xc3\xcb\x91\x7f\xba\x29\x5e\x22\x42\xf2\x34\x27\x5b\xa1\xc6\xca\x64\xc3\x13\x65\xfd\xfc\x44\x09\xe9\x85\xe9\xa8\x96\xb0\x2b\xc1\x04\x9b\xf4\x4f\x3c\x54\x93\xaf\x0b\x9e\xa7\xdc\xff\x22\x63\x1c\xa4\xd1\xf1\x8c\x1e\x23\xc4\x77\xf8\x2c\x61\x5d\xc2\x2e\x4c\x90\x3e\x3c\x4b\xe2\x23\x27\x42\xc7\xcf\x9c\xcb\x4d\x05\x33\x38\x14\xf0\x8f\x51\x3a\xbf\xdc\x54\xe5\xf3\xd2\x82\xf6\x44\xaf\x9c\xf3\xa2\x18\xdc\x25\x66\xa4\x7f\x8a\xec\xff\x1f\x00\x00\xff\xff\x00\xce\x0e\xa6\x70\x0c\x00\x00")

func svcClientGrpcClientGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_svcClientGrpcClientGotemplate,
		"svc/client/grpc/client.gotemplate",
	)
}

func svcClientGrpcClientGotemplate() (*asset, error) {
	bytes, err := svcClientGrpcClientGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "svc/client/grpc/client.gotemplate", size: 3184, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x63, 0x72, 0x1f, 0xe5, 0x3a, 0x45, 0x1, 0x91, 0xd8, 0x5b, 0xa8, 0x47, 0x45, 0x45, 0x98, 0xee, 0x0, 0xf5, 0xc1, 0x3c, 0x43, 0xf0, 0x86, 0x3c, 0xec, 0xbe, 0x2d, 0x84, 0xed, 0x1a, 0x17, 0x6c}}
	return a, nil
}

var _svcClientHttpClientGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xcb\xb1\x0d\xc2\x30\x10\x05\xd0\x3e\x53\x5c\x1d\x09\xdf\x10\x34\x29\x91\xc8\x02\x56\xf8\x98\x88\xc3\x67\x9d\x3f\x95\xe5\xdd\x69\x18\xe0\x8d\xa1\xab\xdc\x01\x29\x7e\x61\x7c\x7b\xd7\x82\x5a\xfc\x7d\x52\x5f\x64\x63\xe4\xda\x9b\x07\x95\xf8\x34\xcb\x44\x2a\x2e\x4f\x0f\x39\xfc\x01\x59\x75\xce\x65\x8c\x23\x9b\x49\xda\xf6\xfd\xb6\xc1\x1a\x22\x5d\xed\x44\xe5\xfe\x27\x92\xe6\x5c\x7e\x01\x00\x00\xff\xff\x0b\x3c\x4c\x9e\x69\x00\x00\x00")

func svcClientHttpClientGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_svcClientHttpClientGotemplate,
		"svc/client/http/client.gotemplate",
	)
}

func svcClientHttpClientGotemplate() (*asset, error) {
	bytes, err := svcClientHttpClientGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "svc/client/http/client.gotemplate", size: 105, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa1, 0xf0, 0x36, 0xf9, 0x16, 0xea, 0x9d, 0x4e, 0x73, 0x64, 0xc5, 0xad, 0xb3, 0x1b, 0x4, 0xe, 0xd8, 0xc8, 0x1e, 0xf7, 0x7a, 0x39, 0x40, 0x4c, 0xb2, 0x12, 0x83, 0x35, 0xca, 0x82, 0x6f, 0xd0}}
	return a, nil
}

var _svcEndpointsGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\x4d\x6f\xdc\x38\x12\x3d\x4b\xbf\xa2\x22\x78\xe1\xee\x40\xa6\xef\x0e\x7c\xd8\x4d\xbc\xbb\x06\x36\x1f\x88\xbd\xbb\x87\x20\x08\xd8\x52\xa9\x45\x98\x22\x19\x92\xea\x76\x8f\xa0\xff\x3e\x28\x52\x52\xab\xd3\x4a\x26\x33\xc7\xc1\x1c\x0c\xdb\xfc\x78\x7c\xf5\xaa\xea\x91\xba\xbe\x86\xd7\xba\x44\xd8\xa2\x42\xcb\x3d\x96\xb0\x39\x80\xb7\xad\x73\x0c\xde\xbc\x87\x77\xef\x1f\xe1\xee\xcd\xfd\x23\x4b\xaf\xaf\xe1\x23\xda\x56\x29\xa1\xb6\x71\x01\xec\x85\x94\xa0\x77\x68\xf7\x56\x78\x04\x5f\x0b\x07\x95\x90\x18\x16\xff\x0f\xad\x13\x5a\xdd\x40\xd7\xb1\xe1\xef\xbe\x9f\x4d\xc0\x1b\xee\x71\x3e\x4b\xff\xf7\x7d\x9a\x1a\x5e\x3c\xf1\x2d\x82\xdb\x15\x29\xad\x7f\x1c\x61\xa1\xd0\xca\x73\xa1\x1c\x34\xe8\x6b\x5d\x3a\xf0\x1a\x1a\xfe\x84\x20\x54\x29\x76\xa2\x6c\xb9\x04\x54\xa5\xd1\x42\x79\x07\x95\xd5\x0d\x38\xb4\x3b\x51\xa0\xcb\x09\xc9\xe2\xd7\x16\x9d\x07\xae\x4a\xb0\xe8\x8c\x56\x0e\xc1\x1f\x0c\x06\x24\x5a\x4a\x41\x68\x87\x47\x94\x1c\xb8\x83\x3d\x4a\x49\xbf\x51\x15\xba\x44\xeb\x08\x80\xf0\x4a\x1c\xfe\xaf\xb4\x1d\x36\x06\xb4\x3c\x0c\x70\x12\xa7\x02\xdd\x5a\x70\xad\x31\xda\x92\xb8\xde\x72\xe5\xe8\x6f\x3a\x4e\x70\x29\x7e\xe1\x5e\x68\x45\x68\x95\xb6\x0d\xf7\x8e\xa5\xa9\x68\xc2\x8a\x55\x9a\x64\x55\xe3\xb3\x34\xc9\x28\x72\x7c\xf6\x59\x9a\x26\xd9\x56\xf8\xba\xdd\xb0\x42\x37\xd7\x5b\x7d\xf5\x24\xfc\x35\xfd\x8c\x8c\x69\x89\xd9\x40\xd6\x75\xec\xc3\x3f\xee\x03\xd0\x07\xee\x6b\xb8\xea\xfb\x2c\x5d\x07\x41\xef\x26\x89\x0a\x2d\x25\x16\xde\x8d\x5c\x7d\x3d\x0b\x1d\x7c\xcd\x3d\x14\xba\x31\x14\x18\x57\xc0\xcb\x72\xd4\x93\xc1\xbd\xbf\x74\x04\xd6\x20\x57\x9e\xe4\xdb\x20\xb4\x0e\x4b\xd2\x89\x43\x8d\xd2\xa0\x05\xe7\x6d\x5b\xf8\x9c\xa6\x87\xa3\x96\x4f\x12\xca\x6b\xe0\x04\xe7\x84\xda\x4a\x04\xc3\x2d\x6f\xd0\xa3\xa5\x52\xa2\xf1\x7b\x05\x3c\x66\xc8\xe6\x20\xfc\xa5\xa3\xc3\xaa\x56\x06\xa5\xab\x56\x15\xa4\xe2\x40\x59\x21\x09\xad\x41\x9b\x50\xd1\xa0\x69\xaf\x41\x7b\x35\x1e\x48\x80\x1b\xee\x84\x63\xf0\x4f\x6d\x01\x9f\x79\x63\x24\xe6\x70\xd0\x2d\x34\x62\x5b\x7b\x30\xdc\x51\x96\x67\x52\x11\xc1\xe9\xa0\x78\x8e\xb1\xba\x6c\x0b\x0c\x32\x70\x05\xb5\xf7\x86\xfd\x9b\xab\x52\x12\xc7\xbd\xf0\x35\x20\x2f\xea\xa1\x58\x61\x35\x9e\xbe\x86\xbd\xb0\x58\x42\x6b\x22\xa8\x33\x58\x88\x4a\x14\x60\xb8\xaf\x19\xac\xee\x03\x3f\xe1\x08\x7f\xc3\x37\xf2\x00\x1c\x1a\xe1\x7c\x2c\x74\x28\xd1\x89\xad\xa2\xad\x42\xed\xf4\x13\x06\x29\x1f\x62\x5a\xa6\xc6\x08\x14\xf1\x34\xd9\x31\x19\x04\x31\x2a\xc9\xd6\x73\x75\x0b\x29\x50\xf9\x53\x75\x67\x89\x3b\xf6\x98\x3c\x50\x27\x46\x38\x2c\x7f\x94\x46\xea\x86\xa8\x95\x20\x85\x1b\x8c\x65\x75\xe4\x2b\x94\x47\x5b\x71\x2a\xa8\xe5\x4c\x10\xd8\x74\xd8\x72\x9f\xb7\x2e\x3a\xd2\xd0\x58\xd7\x21\x0f\xef\x70\xff\x7a\x88\xa7\xd0\xcd\x46\xa8\xa0\x53\x33\x50\x9c\x25\x36\x1f\xdc\xc0\xb7\x56\x81\x08\x95\x4c\x04\x0b\x2e\x25\xda\x58\xcc\x03\x59\x96\x86\x70\xce\x04\xed\xd2\xae\xb3\x5c\x6d\x11\x2e\x04\xdc\xdc\x02\x1b\xd7\xbf\x8d\xc9\xe8\xfb\x34\xe9\xba\x0b\xc1\xde\xf1\x06\xfb\x7e\xdc\x0f\x00\x53\x10\x6c\x1c\x4c\xbb\xee\x8a\x46\xfb\x3e\xed\x4f\x7b\xf5\x27\x0e\xa1\xea\x84\xd5\x8c\xe1\x1a\x66\xe7\xae\x0a\xff\x0c\x83\x8f\xb0\xd7\xf1\x77\x4e\xd5\xf0\xd2\x6c\x58\xd7\xfd\x4b\xd3\x32\xb8\x10\xec\x63\x74\xc9\xc7\x83\xc1\x61\xeb\x1a\x56\xe7\x8b\xa2\x7d\xce\x56\xe5\x80\xd6\x6a\xbb\x86\x2e\x4d\x92\xd1\x5e\xc3\x20\x11\x46\xb6\xa0\x01\x71\x22\x0e\xeb\x34\x49\x44\x15\x96\xbe\xb8\x05\x25\x64\xc0\x48\x86\xac\x28\x21\x03\x4c\x9a\x24\x7d\x3a\x8d\x8e\x27\xb0\x9f\xe1\xb6\xce\x09\x25\x4d\xfa\xb4\xeb\xa2\xbc\x24\xee\x5b\x6a\xa9\xb9\xc2\xa1\x69\x2f\x3c\x06\x85\x63\xde\xe6\xa2\x5f\x78\x5c\xd2\x3d\x0a\x4f\x60\x4b\x21\x3a\x08\xf4\xe6\x7b\xe3\x8a\x87\xd0\x83\xeb\xf3\x22\x38\x09\x9e\xb0\x97\x53\x37\xde\x66\x53\x0f\x75\x94\xa8\xe9\x5e\x9b\x0d\xc7\x24\xcc\xb2\x43\xe8\x5f\x29\xa2\x01\x63\x49\xc3\xb3\x22\x08\xfb\x76\x53\x42\x1d\xfb\xa6\xb8\x02\xa3\xb8\x6a\x21\x97\x4b\xd9\x8c\xf9\x9c\x66\x76\x43\x92\xe2\x70\x50\x3f\xe6\x6a\x9e\xb3\xff\x5b\x6e\xfe\x2e\xe5\xdd\x73\x81\xc6\xc3\xde\x72\xe3\xa2\xcd\x4e\xea\x55\x02\x65\x49\x77\xcc\xd0\x9f\xc7\x86\x0d\xe9\x0d\xfe\xb4\x70\x71\xb2\xb7\xa2\x2c\x25\xee\xb9\x8d\xef\x97\xff\xba\xf1\x45\x43\x77\xb9\x31\xf2\x40\x36\x43\xd6\xe9\x09\xbc\x99\x56\x87\xbb\x01\x77\x68\x0f\x53\x2a\xa9\xad\xc8\x45\xc6\xdb\x92\xf0\xde\x1b\xba\x39\xc8\x3d\xf3\x99\x79\x15\x5c\xd1\xcd\x49\xf7\x0d\x96\xb4\x6d\x73\x00\x45\x39\x88\x37\x2a\x3e\x17\xb2\x2d\xb1\x8c\x8f\x99\x0d\x12\x05\x8a\xd9\x60\xc9\xce\xd4\x58\x1d\x39\xe5\x90\x3d\x78\xee\x5b\x97\xe5\x90\x7d\x10\x6a\x9b\xad\xd3\xd1\x1e\x5e\xce\xfc\xe1\x7b\xfb\x61\x41\x95\xfc\xc8\x86\x31\xe6\xbc\x15\x6a\x1b\xca\x49\xa8\x61\xf8\xe6\x16\x1a\x6e\x3e\xc5\xa9\xcf\x51\xfe\xae\xa7\xf4\x93\xad\xfd\x96\x7d\x25\x49\x36\xab\xa8\xec\x06\x26\x80\x3e\x1f\x20\x62\x19\x24\x7d\x9a\x26\x94\x95\x2f\x44\x29\x94\x71\x80\x9e\xe8\x75\xd1\x4e\xbe\xe4\xa0\x9f\x68\x7a\x24\xf8\x09\x9f\x3f\xbf\x82\x17\xfa\x29\x96\xa4\xe1\x4a\x14\xab\xaa\xf1\xec\xc1\x58\xa1\x7c\xb5\xca\xee\x46\x88\x29\x93\x97\x7f\x73\x97\x50\x6a\x74\xa0\xb4\x07\x7c\x16\xce\xbf\x02\x87\x38\x2f\x80\xa9\x86\x1c\xdb\xea\x8c\x48\xad\xd7\x83\x59\x95\x28\xd1\xe3\x6a\x64\x10\xe6\x8e\x01\x08\x55\xe4\xf0\xe5\x18\xc1\xa4\xe4\xcf\x6b\x26\x2a\xda\x05\xb7\xb7\x70\xa2\xde\xd0\x74\x8b\xae\x0b\xb7\x33\xf2\xab\xc5\x25\xeb\xb1\x0b\x4f\x54\x8f\x1d\xf8\x1f\xbe\x41\x89\xe5\xb1\x30\xe2\x77\xc0\x16\xfd\x58\xc6\xf3\xc7\x5d\xac\xe6\x7d\x8d\x6a\x9a\xd5\xb3\xca\x1d\xc0\x62\x01\xe6\xb1\xe1\x86\x9e\x68\xe3\x62\x88\x1f\x17\x3c\x7e\xa1\x88\x82\xde\x38\x56\x14\xf1\xf1\x39\xe3\x50\x8b\xa2\x0e\x5b\x1d\xaa\x25\x0a\xc3\xc5\x3e\xec\x1e\x9f\x35\xda\x0e\xd7\xfa\x79\x54\xc1\x79\x63\x2d\xe7\xe7\x26\xbd\xe0\xdb\xe9\xf7\xe2\xfa\xc3\x36\x75\x46\x2a\x1f\xe2\x0c\x8a\x5b\x2c\x50\xec\xe2\x03\x30\x84\xf8\xcd\xbb\x9a\xc1\x03\xe2\x04\x33\x43\x09\x13\xe3\xbb\x74\xb2\x00\x22\x4a\x45\x59\xa2\xe7\x42\x86\x37\xe4\xd8\x51\xe1\xf3\x64\x78\xfb\x72\x29\xfc\x81\xfd\xc8\x4d\x4e\x62\x9f\x9b\xca\xef\x56\xf4\x2f\xcb\xf9\x53\x5a\xce\xc9\xb6\x7c\xf9\x69\xf8\x3d\x07\xfa\x35\x00\x00\xff\xff\x9a\x16\xde\x59\xb0\x10\x00\x00")

func svcEndpointsGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_svcEndpointsGotemplate,
		endpointsGotemplate,
	)
}

func svcEndpointsGotemplate() (*asset, error) {
	bytes, err := svcEndpointsGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: endpointsGotemplate, size: 4272, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x60, 0x30, 0x93, 0xa5, 0x9f, 0x65, 0x43, 0xe2, 0x5b, 0x4b, 0xe, 0x33, 0x90, 0xa8, 0x3b, 0x7b, 0x3f, 0xbe, 0xbc, 0x3f, 0x60, 0x73, 0x30, 0x1, 0x4, 0x4c, 0x5f, 0x3d, 0x95, 0x45, 0x5f, 0x38}}
	return a, nil
}

var _svcServerRunGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x56\x4d\x6f\xdb\x38\x13\x3e\x8b\xbf\x62\x2a\xf4\x7d\x21\x03\xae\x54\x60\xb7\x7b\xc8\xd6\x87\x36\x4e\xdb\x00\x4d\x6a\x38\x6e\xf7\xb8\xa0\xa5\x91\x4c\x54\x22\xb5\x43\xca\x6e\x20\xf8\xbf\x2f\x86\x92\x6c\xb9\x4d\xdc\x66\x73\x88\x25\xce\xcc\x33\x0f\xe7\x53\x49\x02\x97\x26\x43\x28\x50\x23\x49\x87\x19\xac\xef\xc1\x51\x63\x6d\x0c\xf3\x4f\x70\xfb\x69\x05\x57\xf3\xeb\x55\x2c\x92\x04\x96\x48\x8d\xd6\x4a\x17\x9d\x02\xec\x54\x59\x82\xd9\x22\xed\x48\x39\x04\xb7\x51\x16\x72\x55\xa2\x57\xfe\x82\x64\x95\xd1\x17\xd0\xb6\x71\xff\xbc\xdf\x8f\x04\x30\x97\x0e\xc7\x52\x7e\xdf\xef\x85\xa8\x65\xfa\x55\x16\x08\x16\x69\x8b\x24\x84\xaa\x6a\x43\x0e\x22\x01\xfd\x5f\x98\x97\xb2\x08\x8f\xaf\xc6\x8e\x5e\xf2\xca\x85\x22\x08\x4b\x53\xf0\x8f\x46\xd7\xff\x24\x1b\xe7\xea\xf1\x73\x52\xd7\x64\xf2\x50\x88\x20\x49\xe0\xb7\x0c\x16\x92\xdc\xbd\x08\xc2\xc2\x98\xa2\xc4\xb8\x30\xa5\xd4\x45\x6c\xa8\x48\x0a\xaa\xd3\x5e\x6f\xc5\x57\xbc\x43\xda\xaa\x14\x45\x50\xaf\x21\x6c\xdb\x78\xf1\xf6\xda\x53\x5c\x48\xb7\x81\x17\xfb\x3d\x7b\x69\xdb\xf8\xf4\x10\x12\xbb\x4d\x1f\x91\x6c\xa4\xce\x4a\x24\x1b\x8a\x89\x10\x5b\x49\x30\xc7\x5c\x36\xa5\xbb\x34\x3a\x57\x05\x74\x3f\x42\xe4\x8d\x4e\x41\x69\xe5\xa2\x09\xb4\x22\xe0\x28\xc4\x77\x8e\x94\x2e\xbe\x48\x8a\xfe\x7f\x62\x14\xcf\x71\xdd\x14\x6f\xb2\x8c\xa6\x10\x66\xfc\x1c\xcb\x2c\xa3\x70\x0a\xe1\xc5\xab\x97\x7f\xbc\xe4\x07\xaf\x02\x52\x67\x50\xa1\x23\x95\x5a\x28\x95\x75\xa8\x81\x35\xd1\xda\x70\xf2\x33\x27\x1f\x56\xab\x45\xef\x83\x43\x3a\x76\xf1\xca\xbb\x60\x85\x27\xa3\xbe\x5f\x2e\x2e\x7b\x54\x0e\xfd\x18\xf5\x77\x8f\x5a\x2c\x17\x97\x10\x31\xf6\xe4\x47\x70\x9f\xa7\xcf\x16\x01\xf5\x56\x91\xd1\x15\x6a\x07\x5b\x49\x4a\xae\x4b\xb4\x53\x50\x39\x58\x74\x31\xbc\x2b\x65\x61\x61\x23\xb7\x08\x35\x29\x43\xca\xdd\xfb\x5a\x86\x2b\xbd\x65\x7d\x1b\x8b\x40\xe5\x1e\x18\x2e\x66\x60\x6c\xfc\x1e\x1d\xea\x6d\x14\xce\xaf\xde\x7e\x7e\xff\xf7\x9b\xf9\x7c\x19\x4e\xfe\xec\x14\x9e\xcd\x20\x0c\x39\x29\xc1\x23\x59\x80\x99\x57\x14\xc1\xde\xa3\xfa\x8a\x3e\x45\x5d\x7c\x5a\xae\x18\xcf\x8b\x1e\xc3\x1b\x02\x0e\x33\xc8\x2b\x17\xdf\xd5\xa4\xb4\xcb\xa3\xf0\xe2\x7f\x36\x9c\x7a\xd3\xc9\xe0\xe2\x01\xe2\x6c\xfd\x6b\xbc\x47\x7e\xc6\xb4\x1f\xc0\xe4\x64\xfd\x1a\xe6\x90\xd6\x11\xe6\x5e\x08\x3f\x7f\x7c\x99\xa7\x46\x3b\xa9\xb4\x05\xb7\x41\x20\xfc\xa7\x51\x84\x19\xe4\x0a\xcb\xcc\x42\x6e\x08\x86\xe1\x23\x87\xd1\xe0\xee\x6b\x1c\xac\xad\xa3\x26\x75\xec\xf6\xc0\xdd\xfa\x02\x13\xc1\x31\x09\xc3\xc9\x81\x4a\x7f\xb0\xef\x9b\xeb\x16\x77\x57\x3a\xab\x8d\xd2\xce\x46\x13\xb0\xdb\x34\x3e\xbc\x33\x74\x92\xc0\xdb\xc6\x2a\x8d\xd6\x42\x66\x2a\xa9\x74\x2c\x02\x6e\x58\xdb\x8d\x04\xa8\xd7\x71\xdb\xc6\xfd\x80\x88\x6f\x65\x85\xfb\xfd\x5d\x47\x36\xe0\x88\x0c\x7a\x33\x18\x7a\x3e\xbe\xc5\x5d\xaf\x1f\x4d\x44\xc0\x2e\xfe\x22\x59\x0f\x43\x06\x76\xca\x6d\xa0\x52\x59\x56\xe2\x4e\x12\xda\x18\xee\x10\x0f\xd6\xc9\x58\x52\x98\x87\x3d\x30\xde\xe0\xa2\x17\xfb\x2a\xf1\xf7\x19\xee\x77\x7a\x9f\x48\x04\x6d\x4b\x52\x17\x08\xcf\x15\x67\xfc\x70\xa7\x1b\x74\x1b\x93\x59\x1e\x5c\x22\x08\xda\x76\x65\x3e\x9a\x1d\x12\x3c\x57\xfd\x75\x0f\x80\x33\x1f\xbf\x1b\xf9\x15\xdb\xf6\x07\xe9\x88\x48\xdb\xa2\xce\x18\x8d\x7b\x17\x0f\xe1\xbe\x98\x9d\xc6\xbf\xfd\x65\x4a\x3f\x38\xbb\xe0\xcd\x70\x86\xea\x74\x44\xa2\x0f\x8b\xcf\x81\xc5\x12\x53\x5e\x89\xc7\x22\x78\x62\x3a\x8e\xd7\xf9\x2e\x1f\xc7\x32\x3b\xa8\xf0\xf5\x09\x5d\x43\x1a\x0e\x67\x7d\x87\x2c\x1b\x0d\xd6\x49\x72\x16\x24\x68\xdc\x01\x4f\xdb\xbe\x0b\xa6\xe0\xc7\xe1\xf0\xc2\xe3\x5c\x82\x9f\xf8\xfd\x59\xc7\xd9\x6d\x90\x91\x6a\x69\x2d\x66\xdc\x6b\xdc\x34\xac\x5c\x9a\xa2\x40\xea\xea\x7f\xd9\xe8\x28\xcd\x87\x8d\xe3\xb7\xcc\x49\x42\x4e\xfb\xa3\x8b\xd4\x0d\xa6\x1b\xa9\x55\x2a\xcb\x63\x09\x21\x51\xca\xfa\x95\xfc\x8a\x11\x8b\x01\x89\x0c\xf5\x16\xd7\xda\x21\x51\x53\xbb\x21\x24\xb1\x08\x0a\x73\x8c\xcf\x41\xfe\xa1\x3b\x89\x18\xae\xb7\xed\x76\x56\x37\xf4\x07\x43\xa6\xde\xad\xc4\xa0\x34\x45\xbc\xe0\x99\x58\xea\x28\x74\x24\xb5\xe5\x99\x18\x0e\x3b\x90\x1f\xfa\x6d\x92\xe6\xa3\xe9\xcc\xe0\x41\xc5\x8c\xfd\x1a\xeb\x5b\x12\x6f\x9a\x6f\xbe\x27\xab\xb8\x63\x12\x85\x89\x87\xe9\x3e\x1d\x92\x70\xda\xa9\xf7\x34\xdf\x31\x0d\x2f\x89\xaf\x75\x86\xdf\x26\x67\x4c\xd3\x2a\x2b\x95\xc6\xc7\x11\x2e\x3b\x85\x73\x18\xfc\x4f\x95\x67\x30\x16\x9d\xc2\x39\x0c\x7b\x5f\xad\x4d\xf9\x38\xc4\x9d\x97\x9f\x43\x70\x24\xd3\x33\x1c\x56\x2c\x9e\xf8\xf8\xfa\xa2\x78\xfd\xa2\xd3\xfc\xe8\x33\xf8\x46\x67\x3e\xd0\xd1\x49\x36\xa6\x50\xf1\x7c\x1a\x0a\xcc\x7f\x43\x1c\x72\xf9\x84\x94\xb3\xe1\x77\x19\x1f\x76\x03\x5f\x68\x33\x8c\x18\x1e\x51\x2c\x38\x94\xdb\xb1\x23\x7f\xce\xfa\xf8\x09\xb4\x19\x93\xf6\x3d\xf9\x5f\x48\xb3\x61\x38\x1d\x73\x1e\x96\x15\xb3\x29\xf5\x94\x5b\x89\x99\x6b\x74\x3d\x9f\x28\x74\x69\xfd\x80\xb2\xca\xbd\xee\xb3\x19\x68\x55\x7a\xb7\x87\xdb\x20\x11\xbf\x76\xe3\x46\x04\x7e\xea\x05\x96\xb6\xe3\x90\x30\x54\xb7\xbb\x4e\x23\xe2\x27\x81\xff\x30\x1b\xfa\x84\x7c\x97\xd4\xeb\x78\x89\x05\x33\xa2\x47\x76\x60\x64\xa7\x60\x69\x7b\x52\x0d\x36\xee\x62\x59\xea\x71\xf8\x96\x8d\x7e\x26\x4e\xa3\x84\xdf\x14\x07\xe8\xf5\x8b\x6e\x1a\xec\x85\xf8\x37\x00\x00\xff\xff\x1a\x36\xad\xf0\xbb\x0c\x00\x00")

func svcServerRunGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_svcServerRunGotemplate,
		"svc/server/run.gotemplate",
	)
}

func svcServerRunGotemplate() (*asset, error) {
	bytes, err := svcServerRunGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "svc/server/run.gotemplate", size: 3259, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x72, 0x35, 0x88, 0xf5, 0x53, 0xf0, 0x1, 0x9d, 0xff, 0x9a, 0x9e, 0xfb, 0x16, 0xf4, 0xb8, 0xd3, 0xb3, 0xb0, 0xd4, 0x2, 0xbd, 0x9b, 0x29, 0xa1, 0xc, 0x74, 0xc1, 0x5, 0xb0, 0x77, 0xa0, 0x7c}}
	return a, nil
}

var _svcTransport_grpcGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x56\x4d\x6f\xdb\x38\x10\x3d\x8b\xbf\x62\xd6\x28\x16\x56\xe1\xd0\x7b\x0e\x90\x4b\x93\x6e\x5b\xec\xe6\x03\x59\xa3\x7b\x28\x8a\x82\x96\xc6\x12\x61\x89\x54\x48\xca\x89\x97\xd0\x7f\x5f\x0c\xf5\x61\x39\x76\x1c\x1f\x0c\x58\xe4\xe3\xcc\x9b\xf7\x86\x23\xcd\xe7\x70\xad\x53\x84\x0c\x15\x1a\xe1\x30\x85\xe5\x16\x9c\xa9\xad\xe5\x70\x73\x0f\x77\xf7\x0b\xf8\x7c\xf3\x6d\xc1\xd9\x7c\x0e\x8f\x68\x6a\xa5\xa4\xca\x5a\x00\x3c\xcb\xa2\x00\xbd\x41\xf3\x6c\xa4\x43\x70\xb9\xb4\xb0\x92\x05\x06\xf0\x77\x34\x56\x6a\x75\x09\xde\xf3\xee\x7f\xd3\x8c\x36\xe0\x46\x38\x1c\xef\xd2\x73\xd3\x30\x56\x89\x64\x2d\x32\x04\xbb\x49\x18\xe1\x17\x7d\x58\xa8\x8c\xde\xc8\x14\x2d\x58\x34\x1b\x34\x17\x56\xa6\x08\x4b\xa9\x52\xa9\x32\x0b\x2b\x6d\xc0\xe5\x08\xd9\xe3\xc3\x35\x38\x23\x94\xad\xb4\x71\x81\xcb\x37\x07\xb5\x93\x85\xfc\x0f\x6d\x80\x0c\xbb\xf3\xcc\x54\x09\xff\x27\x84\xe3\x8c\xc9\x92\x16\x61\xca\xa2\x89\x42\x37\xcf\x9d\xab\x26\x2c\x9a\x24\x5a\x39\x7c\x71\x13\xc6\xa2\x49\xa6\x75\x56\x20\xcf\x74\x21\x54\xc6\xb5\xc9\x42\x88\x79\x89\x4e\xa4\xc2\x09\xc2\xd0\xc2\x90\x01\x26\x99\x74\x79\xbd\xe4\x89\x2e\xe7\x99\xbe\x58\x4b\x37\xa7\xdf\x3e\x05\x3a\xd6\x97\x4a\x6c\x64\x82\x2c\xaa\x96\x30\xf1\x9e\x3f\x7c\xfa\x16\x68\x3d\x08\x97\xc3\x45\xd3\x4c\x58\x1c\x74\xb9\x15\x6b\xfc\xf2\xf8\x70\xdd\xb2\x87\x52\xac\xd1\x82\x00\x8b\x0e\xf4\x0a\x50\xa5\x95\x96\xca\x59\x10\x1b\x21\x0b\xb1\x2c\x10\x04\xed\x07\x79\xbc\xe7\x5d\x1a\x7e\x27\x4a\x6c\x9a\x5e\x82\x55\xad\x92\x57\x91\xa7\xbb\x50\x9f\xfb\x7f\x33\xd0\x95\x93\x5a\x59\xe0\x9c\xef\xd5\xdb\x89\x79\x1f\xb6\x63\xa8\x96\xfc\x8d\x5c\xe0\x59\x64\x47\x58\x0b\x97\x57\xf0\xe3\xe7\xdb\xc1\x3c\x8b\xa2\x63\xbb\x9f\x70\xa5\x0d\x4e\x7b\x07\x16\xfa\xba\xb5\x2b\x9e\xb1\xa8\x79\x9d\xe3\x0a\x44\x55\xa1\x4a\xa7\x7b\xcb\x43\x39\x9c\xf3\x98\x45\x06\x5d\x6d\x14\xfc\x4e\xd9\xda\x1c\x3e\xd8\xe3\x3d\x2c\xf4\xdf\xfa\x19\x0d\xec\x95\x04\x4d\xc3\x22\xef\x8d\x50\x19\xc2\x07\x49\x85\x0c\xfb\xb7\xe8\x72\x9d\x5a\x42\x44\xde\xf7\xc7\x3f\xc8\x4e\x8b\x4b\xd8\x2f\xe9\x0e\x9f\x3b\xd5\x59\x14\x45\x83\xf2\xdc\xfb\xe1\x48\x6f\xc2\x8c\x10\x37\x98\xe8\x34\x98\x35\x42\x3c\xe2\x53\x8d\xb6\x05\x7c\x56\x47\x01\xb6\xd2\xca\x62\x40\xec\x29\xc1\x39\xa7\x45\xd2\xce\xfb\x0b\xea\x22\x62\xde\xb0\x26\xb4\xdc\x4e\x10\x90\x65\x55\x60\x89\xd4\x15\x74\xa3\xbc\xff\xa2\x83\x14\xc7\xbd\x96\xca\xa1\x59\x89\x04\x99\xdb\x56\x38\x8e\x63\x9d\xa9\x13\x07\x9e\xbd\xaf\xdf\x11\xf9\x00\x5e\xe9\xf7\x55\xa8\xb4\x40\xc3\x76\xe4\x5b\xe6\x5d\x98\x30\x24\x46\xd9\x9d\xde\x15\x72\x7e\x0d\xef\x52\x0d\xb7\x68\x6a\xe1\xe3\x2e\x55\xbc\x0b\x3f\xb0\x9f\x26\xee\x05\xba\xe1\xc2\xbb\xae\x9d\x81\xc1\x27\xf8\x18\xee\xcd\x0e\xdf\x39\xba\xd8\x56\x3d\xa9\x18\xa6\x87\xa0\xd6\xd5\x11\x6a\x06\x68\x8c\xa6\xe4\x2c\xfa\x45\xa1\xab\xb0\x42\xb4\xa9\xa7\x0e\xf4\x6c\xaf\x14\x75\x0b\x71\x0b\x5c\x62\x16\xc9\x55\x38\xf4\xdb\x15\x28\x59\x50\xa8\xfe\x86\x28\x59\x84\x78\xe1\xa2\x75\x6b\x06\x2b\x7e\x0e\xb5\x78\x46\xc7\x59\xc3\xbc\x6f\x8d\x22\x9b\x3a\xa9\xdb\xae\x7e\x5f\xe7\xf9\x1c\x4e\x5d\x00\x90\x34\xf0\x5e\x0d\xfb\xf6\x40\x87\xf8\x93\x8c\x72\xb9\x70\x64\xc3\x06\x0d\x8d\xcb\xd0\xe8\xed\x90\x3c\xec\x37\xd3\x45\x76\x1a\x04\xd4\x16\xcd\x45\xaa\x4b\x21\xd5\x29\x30\x87\x07\x23\x4b\x61\x64\xb1\xa5\x23\xab\xba\x00\xa9\xc2\xa4\x1e\xcd\xdc\x53\x75\x4c\x7f\x1d\x76\x09\xd5\xf2\x88\x4f\xbb\xae\xf4\xd4\x12\xa3\xa7\xb1\xf5\xd4\x52\x97\x57\xfd\x99\x63\xf6\x1c\xb4\xd7\xc8\xcf\xa7\x13\x4e\xb5\xe3\xe5\x2c\xa7\x4e\x4e\xa2\xa3\x56\xb5\x27\x7a\xc8\x5b\x5e\xbd\xef\x42\x97\x22\x78\x76\xc2\xd9\xaa\xd8\x9e\x65\xd5\xc9\x42\x8e\x79\x35\x30\x38\xd3\x2c\x5b\x91\x8a\xfd\xa9\xf3\x6e\xd3\xc8\x2f\x5b\x1d\x33\xec\x2b\x16\x15\x1a\xcb\xda\x1a\x0e\xde\x96\xc7\x67\x51\x99\x0e\x48\x7e\x7b\x13\xbf\x06\x10\x5d\x9a\xa8\xeb\x19\x6c\x02\xe5\xd0\x04\x65\x1a\x66\x84\x5c\xc1\x66\x3c\x33\xda\x0f\x1c\x84\x35\x6e\x83\xdb\x69\x4a\x1f\x9b\xda\xe5\x24\x71\x9f\x85\x06\x74\x29\x1c\x4c\xd7\x31\x3c\xe7\x32\xc9\x03\xb4\x28\xa0\x20\xbb\xba\x28\x42\xa5\xe1\xa5\x43\xdf\x67\xfc\x5a\x28\xad\x64\x22\x8a\xaf\x28\x52\x34\x7f\xe1\x96\x3e\x7f\x5c\x97\xc8\xea\xb6\x65\xa4\x83\x44\x28\x58\x62\x1f\x22\x49\xd0\x5a\x4c\x29\x37\x4a\x97\xa3\xe9\x32\xd3\x3e\x49\x71\x35\xd4\xfa\xaf\x74\xf9\x77\x51\xd4\xd8\x8e\x44\xaa\xf5\xc7\x1f\x3f\xe3\x77\x81\x6f\xb0\x9b\xae\xe3\x5d\x84\xf0\x6e\x1d\xac\x4b\xdc\x0b\x6b\xd8\xff\x01\x00\x00\xff\xff\x71\x92\xdd\x9a\x92\x0b\x00\x00")

func svcTransport_grpcGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_svcTransport_grpcGotemplate,
		"svc/transport_grpc.gotemplate",
	)
}

func svcTransport_grpcGotemplate() (*asset, error) {
	bytes, err := svcTransport_grpcGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "svc/transport_grpc.gotemplate", size: 2962, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa2, 0xc7, 0xff, 0xf, 0x16, 0xfd, 0xd8, 0x43, 0xfc, 0x25, 0x2b, 0x9c, 0xfe, 0x3c, 0xf, 0x3a, 0x3e, 0xc7, 0x21, 0xdf, 0x32, 0x41, 0x25, 0x39, 0x7d, 0x4b, 0xb6, 0xc2, 0xc, 0x27, 0x81, 0x3c}}
	return a, nil
}

var _svcTransport_httpGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xcb\xb1\x0d\xc2\x30\x10\x05\xd0\x3e\x53\x5c\x1d\x09\xdf\x1a\x29\x91\xe2\x05\xac\xf0\x31\x08\x93\xb3\xce\x1f\x1a\xeb\x76\xa7\x61\x80\x37\xa7\xae\xb2\x03\x52\xed\x42\xff\x8c\xa1\x15\x67\xb5\xd7\x93\xfa\x20\x3b\xbd\x9c\xa3\x9b\x53\x89\x77\x6f\x85\x18\xa9\x9a\xdc\xcd\xe5\xb0\x1b\x64\xd5\x88\x65\xce\xa3\xb4\x26\x69\xcb\xf9\xba\xa1\x75\x78\xda\xe1\x5f\x78\xfe\x1b\x49\x11\xcb\x2f\x00\x00\xff\xff\xdd\x3a\x4a\x8f\x6a\x00\x00\x00")

func svcTransport_httpGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_svcTransport_httpGotemplate,
		"svc/transport_http.gotemplate",
	)
}

func svcTransport_httpGotemplate() (*asset, error) {
	bytes, err := svcTransport_httpGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "svc/transport_http.gotemplate", size: 106, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x52, 0x57, 0x56, 0xc6, 0xb4, 0xe5, 0x1f, 0xf4, 0x1d, 0xa5, 0xda, 0x23, 0xea, 0x8f, 0xfb, 0xff, 0xae, 0x4b, 0x12, 0xe4, 0xf6, 0xbf, 0x11, 0xa6, 0x4, 0x83, 0x53, 0xfd, 0xbf, 0xce, 0x4a, 0x47}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"cmd/NAME/main.gotemplate":          cmdNameMainGotemplate,
	ServerHandlerPath:                   handlersHandlersGotemplate,
	"handlers/hooks.gotemplate":         handlersHooksGotemplate,
	"handlers/middlewares.gotemplate":   handlersMiddlewaresGotemplate,
	"svc/client/grpc/client.gotemplate": svcClientGrpcClientGotemplate,
	"svc/client/http/client.gotemplate": svcClientHttpClientGotemplate,
	endpointsGotemplate:          svcEndpointsGotemplate,
	"svc/server/run.gotemplate":         svcServerRunGotemplate,
	"svc/transport_grpc.gotemplate":     svcTransport_grpcGotemplate,
	"svc/transport_http.gotemplate":     svcTransport_httpGotemplate,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"cmd": &bintree{nil, map[string]*bintree{
		"NAME": &bintree{nil, map[string]*bintree{
			"main.gotemplate": &bintree{cmdNameMainGotemplate, map[string]*bintree{}},
		}},
	}},
	"handlers": &bintree{nil, map[string]*bintree{
		"handlers.gotemplate":    &bintree{handlersHandlersGotemplate, map[string]*bintree{}},
		"hooks.gotemplate":       &bintree{handlersHooksGotemplate, map[string]*bintree{}},
		"middlewares.gotemplate": &bintree{handlersMiddlewaresGotemplate, map[string]*bintree{}},
	}},
	"svc": &bintree{nil, map[string]*bintree{
		"client": &bintree{nil, map[string]*bintree{
			"grpc": &bintree{nil, map[string]*bintree{
				"client.gotemplate": &bintree{svcClientGrpcClientGotemplate, map[string]*bintree{}},
			}},
			"http": &bintree{nil, map[string]*bintree{
				"client.gotemplate": &bintree{svcClientHttpClientGotemplate, map[string]*bintree{}},
			}},
		}},
		"endpoints.gotemplate": &bintree{svcEndpointsGotemplate, map[string]*bintree{}},
		"server": &bintree{nil, map[string]*bintree{
			"run.gotemplate": &bintree{svcServerRunGotemplate, map[string]*bintree{}},
		}},
		"transport_grpc.gotemplate": &bintree{svcTransport_grpcGotemplate, map[string]*bintree{}},
		"transport_http.gotemplate": &bintree{svcTransport_httpGotemplate, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
