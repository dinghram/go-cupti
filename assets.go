// Code generated by fileb0x at "2019-06-19 12:19:53.898377 -0500 CDT m=+0.019651412" from config file "b0x.yml" DO NOT EDIT.
// modification hash(9cfbed4d61be049656758162530f90e8.90b27eb7f67e71e30d95593145408bb3)

package cupti

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"net/http"
	"os"
	"path"

	"golang.org/x/net/webdav"
)

var (
	// CTX is a context for webdav vfs
	CTX = context.Background()

	// FS is a virtual memory file system
	FS = webdav.NewMemFS()

	// Handler is used to server files through a http handler
	Handler *webdav.Handler

	// HTTP is the http file system
	HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct {
	// Prefix allows to limit the path of all requests. F.e. a prefix "css" would allow only calls to /css/*
	Prefix string
}

// FileEventMappingJSON is "/event_mapping.json"
var FileEventMappingJSON = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x9c\x5f\x6f\xdb\x38\x12\xc0\x9f\x37\x9f\x62\xe0\x87\x7b\x4a\x03\x51\xb2\x65\xa7\x6f\x41\x36\xd7\x0b\xd0\xb4\x8b\xd6\x7b\xc0\xe2\x70\x10\x68\x89\xb6\x85\x4a\xa2\x97\xa4\x9b\xf8\x0e\xfb\xdd\x0f\xfa\xe3\x5a\x4e\x3d\x92\x28\x91\x76\x0b\x5c\x1f\xab\x21\xe3\xf9\xcd\x68\x38\x43\x8e\xf8\xaf\xab\xff\x5e\xfd\x32\x62\x5f\x59\xa6\x46\x6f\x81\x5c\x5f\xfd\x32\x8a\xa3\x11\x94\xff\xde\x02\x71\x1c\xdf\xf7\xbc\x5b\x27\x7f\x10\xf1\x94\xc6\x59\x50\x7f\xee\x3a\xb7\x53\x32\x71\xdd\xfc\x71\x46\x53\x36\x7a\x5b\x3d\x1a\xd1\x50\xc5\x5f\x59\x10\xee\xc2\x84\xc9\x60\x93\x8e\x72\x11\xb9\xe6\x42\x05\x11\x93\xa1\x88\x37\x2a\xe6\xd9\xe8\xed\x5e\x12\x4a\x49\xa8\x24\x13\x9e\xad\x5e\x0b\x7e\xd8\xa6\x0b\x26\x80\x2f\xf7\xb2\x14\xd2\x6d\xa2\xe2\x8d\xe0\x21\x93\x92\x0b\x58\x53\x09\x54\x41\xc2\xa8\x54\xc0\x33\x06\xd5\xdc\xcf\x54\x6c\x6e\x8a\x79\x43\xaa\xd8\x8a\x8b\xdd\x28\xff\xf5\xa3\xfb\xdf\x7f\x9b\x3f\x06\x0f\xff\x7c\xf8\x30\x0f\xee\xef\xe6\x0f\xef\x3e\x7e\xfa\x23\x78\xfc\xf0\x79\xfe\xe9\xf7\xfb\xf9\xe3\xc7\x0f\xa3\xab\xbf\xae\x8f\x00\xb9\x18\x20\xd2\x17\x50\xfe\xd3\x3a\xf1\x29\x04\x9b\xf0\xdc\x85\xe1\x36\xdd\x26\x54\xb1\x08\xb2\x6f\xa8\x8e\x47\x33\x51\xb2\xbb\x81\xbf\x73\x01\xec\x2b\x13\xbb\xf2\x3f\x20\x56\x10\x67\xa1\x60\x29\xcb\x94\x84\xc5\x0e\xd4\x9a\x61\xd3\xc4\x59\xf1\xb4\x1c\xf8\xbc\x8e\xc3\x35\x84\x34\x83\x05\xdb\x3f\x11\x34\x5b\x31\x70\x40\x71\xf0\xc7\x46\xc0\x7b\x08\xf8\xd9\x44\x1b\xbc\x5c\x53\xc1\xa2\x20\x89\x02\x25\x68\x26\x73\xcd\x78\x26\x71\xfe\x0d\xf2\xcd\x5e\x5a\x17\x87\x25\x17\x50\xce\x04\x09\xa7\x11\xd0\x30\xf7\x59\x26\x6f\xe0\x89\xbe\xc4\xe9\x36\xad\x8b\x83\x8c\xff\x53\xc0\xfc\xca\x13\x45\x21\x96\x40\xdc\x19\x2c\x76\x8a\xc9\x6b\xa0\xd9\xae\xb0\x43\x35\x45\x9c\xad\x20\xe5\x82\x81\x5a\xd3\xec\x20\x07\xcf\x71\x92\x40\x48\xb7\x92\x55\x6f\x49\xc2\xbe\xff\x45\xf4\xe8\x37\xc5\x99\x54\x62\x5b\x3c\xbe\x81\xf9\x3a\x96\x40\x13\xc9\x73\xbf\x48\xb6\x11\x93\xc0\x5e\x94\xa0\xc7\x93\x14\x7f\x20\xca\xfd\xa5\x9a\x68\x41\xb3\x2f\x10\xf2\x6c\x99\xc4\xa1\x92\x9d\x6d\x7f\x7f\x77\xff\x8f\x87\xef\xac\x3e\xc6\xac\xee\xf7\xb5\xba\x54\x7a\x56\x3f\x25\xdf\xcf\xea\x52\xe5\x46\xfa\xc1\xcc\x5e\xfe\xa8\x1f\xce\xee\x13\xc4\xee\x63\xfd\xb7\x9d\x25\x74\x23\x59\xb4\x5f\x88\x64\x43\xa0\xad\x44\x21\x4c\x78\xf8\x05\x37\xf5\x09\xb1\x81\xc1\xcd\xc7\x56\x15\x77\xa8\xba\x9b\xce\xea\x02\xcf\xe0\xf3\x53\x47\xa5\x6b\xc2\x03\x55\x9f\x62\xaa\x7b\xda\xaa\xe7\x5e\x1c\xb0\x17\x16\x6e\x15\x8b\x82\xe5\x86\xf8\xc1\x26\xde\xb0\x40\x3a\x38\x83\x96\x31\xcd\x6f\x7a\xf1\x36\xd6\xde\x9d\xfc\x45\x29\x67\xca\xdf\x8b\x7c\x32\xc8\x27\x2b\x61\x81\xdc\x2e\xde\x6c\xa8\x50\x71\xf1\xa6\x3b\x46\x56\xc5\x19\x46\x6f\x6c\x8e\x1e\xe9\x41\x8f\x58\xa6\x47\x4c\xc0\xbb\xc5\xe0\x4d\xcc\xc1\x73\x7b\xc0\x73\x2d\xc3\x73\x4d\xc0\x23\x0e\x46\xcf\x37\x47\xcf\xeb\x41\xcf\xb3\x4c\xcf\x33\x42\x0f\x2d\xb4\xa6\x83\xe9\xf9\x63\xfd\xb0\xf7\x6a\xcc\x30\x7a\xfe\xb8\x21\xec\x19\xa1\x87\x56\x61\x33\x73\xf4\x48\x0f\x7a\xc4\x32\x3d\x23\x61\x8f\x60\xa5\xd4\xed\xad\x39\x7a\x6e\x0f\x7a\xae\x65\x7a\x66\xe2\x1e\x52\x92\x8c\x1d\xc7\x1c\x3d\xaf\x07\x3d\xcf\x32\x3d\x33\x71\x6f\x82\xd1\x23\x43\xe9\xa5\x54\x3b\xec\xbd\x1a\x32\x88\x5d\x4a\x6d\x07\x3d\x1f\x43\xe7\x1a\x43\x47\xf4\xd1\x11\xbb\xe8\xcc\x44\xbc\x29\x86\xce\x33\x86\xce\xd5\x47\xe7\xda\x45\x67\x26\xdc\xcd\x30\x74\x63\x63\xe8\x3c\x7d\x74\x9e\x5d\x74\x66\x62\x1d\x56\x5f\xcc\xa6\x03\xb6\x2c\x17\x34\xfb\x12\xec\xb7\x59\x70\x70\x9f\x6b\xbb\x7a\x47\x3b\x33\x1d\xc0\x49\x6c\x2c\xac\x58\xc6\x44\xb1\xbf\xfc\xbc\x66\xe5\x4e\x2f\x8d\x22\x51\x6c\x68\x15\xfb\x4a\xea\x99\x03\x17\xe5\xbe\x54\x35\x4d\xca\x52\x2e\x76\xe5\x6c\x82\xfd\xb9\x65\x52\x49\x58\xd2\x24\xd9\xef\x15\x4b\x9a\xb2\xbd\x54\xfe\xe7\x06\x6e\x1d\xb9\x58\x61\x32\x9b\x0d\xd8\x33\xd4\xc3\x5e\xee\xaa\xf5\xe4\x7e\x62\xf0\x10\xf0\xe5\x74\x67\x21\x4f\xb0\x58\xa1\x5f\x50\x2b\x96\x49\x2e\xca\xb7\xfd\xf8\x24\xa9\x69\x7d\x6f\x1d\xd5\xe9\x54\xe9\x79\xcd\x04\x83\x72\xae\x32\x46\x3c\x53\xb9\x3f\x00\xe1\x19\x58\x5a\xe0\x5d\x17\xc3\xe7\x9b\xc4\x47\x7a\xe1\x23\x36\xf1\x19\x59\xe4\x5d\x0f\xc3\x37\x35\x89\xcf\xed\x85\xcf\xb5\x89\xcf\xc8\x42\xef\xa2\x75\xcd\xcc\x24\x3e\xaf\x17\x3e\xcf\x26\x3e\x33\x8b\x3d\xb2\xe8\xb4\xd4\x35\xe5\x79\xc6\x31\xbd\xe5\x22\x90\xdb\xc5\xc6\x09\x04\xa3\x51\x20\x59\xa8\xb8\x68\x38\xa7\x5a\x2e\xa0\x10\x87\x5c\x1c\xea\xe2\xcd\xbc\x7e\xfd\x74\xf7\x54\x8e\xf9\xb6\x36\x28\x9e\x4f\x05\xb5\xb8\x76\xfd\xea\x4c\x98\x14\x8b\x8d\xe7\x16\x67\x4d\xd5\x41\x54\xe7\x55\xe3\xe9\xe1\xe9\xe3\xa7\x3f\x3a\x1f\xa9\xb7\xd4\x35\x0d\xe8\x88\x1e\x3a\x62\x03\x1d\x39\x07\x3a\x6c\x27\xa7\xa5\xae\x69\xf4\xba\x67\x11\x2b\xa6\xe1\x76\x85\xbc\x2e\xbc\x72\xd0\x85\x1d\x0f\x3b\x5c\x6e\x29\x6d\x1a\x1d\x4f\x8f\x1e\xb1\x43\xef\x2c\xbe\x87\x45\xbc\x69\xf3\x3e\xd8\xc4\x6b\x69\x15\xc2\xa1\xdd\xd5\xfb\x84\x2c\x35\x09\xd9\xec\x11\x9a\x92\xbe\x64\x8a\xae\x9b\x56\x30\x07\xa9\xff\x77\x07\x55\xc8\x27\xbe\x36\xf2\x8d\xe0\xcb\x40\x89\x78\xb5\x62\x22\x70\x1a\x8a\x8d\x4d\x9a\x4b\x75\xa9\x2b\x4a\x49\x64\x2f\xa2\x4c\x58\x1e\xd3\x94\xf8\xf0\x37\x70\x5e\x1c\xc7\x31\x92\x0f\x63\xd1\x6d\x32\x1d\x88\x84\xb4\x21\x21\x16\x90\x18\xc9\x71\xb1\xae\x92\xc9\x6c\x20\x12\xb7\x0d\x89\x6b\x01\xc9\xd8\x66\xe7\xc9\xe4\x76\x20\x12\xaf\x0d\x89\x67\x01\xc9\xcc\x66\x47\x8a\xef\x0c\x44\x32\x6e\x43\x32\x36\x8f\x84\x38\x36\xdb\x4c\x7c\x32\x10\xc9\xa4\x0d\xc9\xc4\x3c\x12\xd7\xb1\xd9\x3c\xe2\xbb\x03\x91\xf8\x6d\x48\x7c\xf3\x48\xc6\x8e\xd5\x96\x10\xdf\x1b\xc8\x64\xda\xc6\x64\x6a\x9e\xc9\xcc\xb1\xda\xe8\xe1\x4f\xb4\x99\x14\x87\x21\xb1\x94\x5b\x16\x35\xa4\x25\x8f\x07\xfd\x20\xe3\x0a\xca\x01\xdd\xd3\x65\xb5\xa6\x0a\xa2\x38\x3a\x0c\x2e\x3a\x4f\x6b\xd8\x8e\xaa\x8b\x3c\x67\x34\xd6\x5d\x8f\x36\x76\xf8\xfe\x10\x5a\xa4\x1b\xad\x5e\xa4\xca\x41\x20\xe3\x6c\x95\xb0\x33\x51\xc2\xb2\x5d\x7f\xda\x8f\xd2\xfe\x2d\xe8\x86\xe9\x48\xba\x19\xd4\xe9\x77\xcd\x2c\x0c\x2c\xcf\xf5\xf5\x93\x3a\xb5\x2e\xf6\x8d\x3a\x32\x99\x1f\xa4\x8d\xa1\xa9\x0a\xab\xaa\x9c\x2a\x7f\x8f\xac\xaa\x31\x1a\xae\xeb\xa3\xbe\x2f\xc9\x0e\xe5\x58\x35\xf0\xba\xea\xe0\x8e\xb3\x15\x6c\x04\x8b\xe2\x9c\x75\xf4\x86\x2f\x6b\x12\x85\x13\x57\x7f\xbf\xf8\xdb\x47\xbd\xe0\x8f\x0a\x22\xce\x64\x19\x0b\xca\x6e\x70\x10\x6c\x93\xd0\x9d\xbc\xb1\xd9\x0f\xe2\xf9\xfa\xf9\x67\xc6\x55\x70\xd0\x32\xe0\xcb\x65\xa0\x67\xd0\x0f\x5c\xd5\x30\xc1\x01\x13\x58\xb4\x31\xcd\xca\x48\x7b\xf2\x0f\xf7\x30\xfd\xd1\xb4\xd9\x9b\xd3\xd3\x5e\xda\xea\x58\xd5\x31\xd5\x0f\x60\xc5\xc9\x68\x1c\x06\x09\xa7\x0d\xa6\x7d\x57\x4a\xc1\x37\xa9\x66\xc3\xbd\xff\xf5\xb4\xed\x6c\x36\xa3\x78\xd3\x59\x6f\xe5\x8b\x33\xde\x76\xed\x0f\x62\xcd\xea\x7f\x9e\x5b\x54\x1f\xab\x25\x66\xfa\x49\x22\x55\x3c\x0d\x42\xbe\xcd\x1a\x4e\xe4\xef\x14\x4f\xe1\x20\xd3\xac\xf8\xdd\xfc\xe3\x93\x45\xd5\xd1\x86\x10\xfd\x60\xb7\x4a\xf8\x82\x26\x41\x49\x80\x36\xec\x0d\xbe\x2b\x04\x21\x17\xbc\x09\xa9\xd4\x42\x71\x73\x7f\xf7\xd9\x1e\x0e\xbc\x51\x63\xac\x8f\x43\xb0\xa8\xcd\x13\xde\x09\x16\x75\x56\xff\xd3\x83\xc5\x10\xe0\xa2\x7b\xe7\xbd\x1d\xa1\x25\xfc\x95\x4e\xd0\x39\xfa\xbd\xb3\xa8\x3b\x96\xe2\xcf\x9c\xbe\xba\xb7\x45\xbf\x52\xf9\xee\xc1\xcf\xa6\xf6\xe8\x67\xac\xfa\x3b\x29\x09\x0f\x5b\x0d\xff\x3e\x97\xe9\x6e\xf7\xf7\x16\x35\x47\x3f\xe5\x74\x7b\x6a\xde\x62\xf6\x52\xf5\xee\x56\xb7\xa9\x3b\x96\xe5\x4e\xf5\x75\xaf\x5a\xd2\xf2\x78\xde\xda\x86\xf6\x4d\xa8\x53\xdf\x59\x1e\xf1\x6d\x86\x7b\x34\xe9\x1b\x0f\x81\xd0\xbc\xfa\xd5\x40\xc0\x5e\xb0\x3b\x0c\xcb\xeb\x1f\x9a\x08\x4e\xfa\x02\x69\x0e\x06\xb5\x9e\xd0\x4e\xd1\xc0\xa6\xee\x58\x16\x38\xf5\xfb\xea\xde\x12\x0e\xea\x9d\x99\x9d\xe2\x81\x4d\xed\xd1\xcd\x63\xfd\x57\xa1\xbc\x3e\x22\xa1\xdb\x2c\x5c\x37\x15\xb7\x27\xe4\xda\xdb\xa7\x25\x1c\x0d\x18\xd0\xa8\x82\xe5\x7a\x53\xfd\xac\x5f\xa6\x41\xa8\x68\x07\x95\x4f\x09\xb6\x6c\xec\x29\x6a\x4c\x65\xb4\x23\xac\xf9\xed\x1e\x4f\x4f\xac\x77\xee\x89\xde\x9c\x20\x8d\xa5\x6c\x6a\x97\x48\xdc\x13\x1d\x3a\x50\x1b\xd5\x62\xff\xe3\x26\x13\xc9\x32\x05\x8a\x97\x0d\x28\x4b\xc1\x53\x90\x49\x1c\x32\x70\x8a\x58\xe1\x42\x48\xc3\x35\xab\x2e\x0f\x38\xd5\x7c\x52\xec\x61\x78\xee\x9b\x3e\x1d\x28\x48\xbb\x31\x06\xd8\xef\x0b\x98\xf4\x02\x4c\xec\x02\x26\x17\x03\x8c\x76\x97\xdd\xf6\xf7\xe0\x5a\x63\x5e\x77\x07\xae\xb5\xe7\x75\xc7\x7b\xdc\x9f\xf7\xe3\xb9\x2f\x96\x0f\x13\xa7\xbf\xfb\xf6\xa0\x4b\x6c\xd2\xbd\x9c\xef\x62\x19\x37\xf1\x06\xfa\xae\x62\x2f\x7b\xc2\x7f\x6e\x99\x88\xbb\x3b\xb0\x62\x2f\x7b\xcc\xf5\x91\x3a\x9c\x0b\xb4\x73\xf6\xa2\xb6\x82\x95\x50\x8b\xe6\xbf\x0b\x7b\x32\x96\xd8\x93\xe9\xd0\x95\xae\x17\xec\x32\xae\xf6\xa2\xfd\x2a\x24\xff\x98\xb8\xb1\xb2\xc1\x25\x06\x5c\x7b\x1d\xab\xf6\xee\xd5\xef\xfd\x7a\x1d\x2b\x8d\x26\xd6\x76\xa7\x5e\x53\x55\xcc\x19\x67\x17\xc7\x8d\x55\x2a\xee\xc4\x84\x77\xeb\xf1\x3e\xb8\xb6\x1e\xf0\x0e\x7e\xfd\x03\x11\xc7\xaa\x23\x6f\x3a\x34\xef\xd8\xc9\x94\xa5\x1a\x11\xa5\xdb\xe0\x96\x0d\x85\x9d\x54\x2c\xdd\x7f\x66\xf7\xfd\x57\x03\x17\x86\x8d\x36\xed\x8c\xc9\xe0\x3a\xa5\x27\xee\xb6\xd1\x3a\xbc\x4f\xb4\xca\x5f\x1a\x38\x56\x18\xba\x03\xd2\x6a\xc5\x15\x4d\x8e\xd2\xbf\xee\xb8\x9b\xc7\x9e\x82\x3d\xcf\x47\x68\xba\x72\x79\x91\xdc\x71\x08\x82\xf7\xe4\xfa\x38\x0e\x5d\x1f\x9b\xef\xbc\x86\xc1\x0a\x4a\xcf\x1b\x6a\x98\xa3\xb2\x52\xd7\x32\xe8\x60\xdc\x34\x9a\x5e\xff\x13\xd8\x06\xab\x45\xc9\x78\x60\xb5\xd4\x23\xc5\x24\x67\xc9\xe7\x2f\x57\x3b\xa1\x2d\x55\x64\x36\x74\x67\xa5\x17\xed\xf3\x24\xf4\x17\xe4\x8d\x15\xab\xae\x6b\xc0\xb9\x75\x32\x4c\x72\xc6\x8c\xfe\x82\xbc\xb1\x82\xd5\xf5\x4d\xf8\xb7\x1e\xf0\x73\xa6\xf4\x17\x44\x8e\x15\xad\xde\x6c\xe8\x6e\x57\xaf\x24\x93\x9c\x21\xa7\xbf\x20\x6d\xac\x66\x1d\xbb\x83\xb7\xc6\x7b\xe2\x3e\x47\x4e\x7f\x41\xe0\x68\xc9\x3a\x60\x33\xb7\x7f\x4e\x4f\x4c\xe7\xf4\xe4\x67\xcd\x1b\xd1\x1e\x33\x6f\x3c\xd4\x30\xbd\x72\x7a\x62\x3e\xa7\xff\x69\x6d\x83\xde\x16\xd4\x52\x07\x3b\x27\xbe\x16\x0a\x63\x16\xa8\x97\xfd\xb5\x21\x9b\x6d\x22\x1b\x5a\x01\x50\xe9\x96\xef\x87\x72\xb9\xf2\xb6\xa5\xf2\xa3\xe6\x88\xaa\xea\x52\xf4\x34\x96\x32\xe6\x19\xa8\xb5\xe0\xdb\xd5\x1a\x7e\xbb\x7f\xac\xdf\x3b\x62\xe9\x86\x8c\x71\xdb\x51\x11\x46\x4a\x68\x91\x12\x86\x48\x09\x16\xb2\xea\xab\x06\x1b\x90\x30\x77\x22\x2d\x6b\x5e\x87\x2b\xe3\x77\xb2\xf5\x12\xf5\xd7\xb2\x4d\x57\xa8\x57\xaf\x5c\x75\x7d\x7c\x67\x63\x13\x6d\x3d\x5e\xdd\xa2\xd3\xa4\xc6\x69\xd1\x8e\x9f\x41\xb1\xbd\x4a\xf1\xb7\x9b\x76\x9e\x63\xb5\x86\x67\x2e\xbe\x14\xdd\x46\x57\x7f\x5d\xfd\xfb\xea\x7f\x01\x00\x00\xff\xff\x06\x7b\xf4\x43\x6a\x66\x00\x00")

func init() {
	err := CTX.Err()
	if err != nil {
		panic(err)
	}

	var f webdav.File

	var rb *bytes.Reader
	var r *gzip.Reader

	rb = bytes.NewReader(FileEventMappingJSON)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "/event_mapping.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	Handler = &webdav.Handler{
		FileSystem: FS,
		LockSystem: webdav.NewMemLS(),
	}

}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {
	path = hfs.Prefix + path

	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(f)
	return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// WalkDirs looks for files in the given dir and returns a list of files in it
// usage for all files in the b0x: WalkDirs("", false)
func WalkDirs(name string, includeDirsInList bool, files ...string) ([]string, error) {
	f, err := FS.OpenFile(CTX, name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	for _, info := range fileInfos {
		filename := path.Join(name, info.Name())

		if includeDirsInList || !info.IsDir() {
			files = append(files, filename)
		}

		if info.IsDir() {
			files, err = WalkDirs(filename, includeDirsInList, files...)
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}