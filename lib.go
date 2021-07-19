// +build linux,cgo,!nogpu

package cupti

/*
#cgo CFLAGS: -I${SRCDIR}/cbits -I${SRCDIR}/csrc -O3 -Wall -g
#cgo CFLAGS: -std=c11 -Wno-unused-variable -DCUPTI_ENABLED
#cgo CXXFLAGS: -I${SRCDIR}/csrc
#cgo CXXFLAGS: -std=c++11 -DCUPTI_ENABLED
#cgo CFLAGS: -I .
#cgo CFLAGS: -I /usr/local/cuda/include
#cgo CFLAGS: -I /usr/local/cuda/extras/CUPTI/include
#cgo CXXFLAGS: -I /usr/local/cuda/include
#cgo CXXFLAGS: -I /usr/local/cuda/extras/CUPTI/include
#cgo CFLAGS: -DFMT_HEADER_ONLY
#cgo LDFLAGS: -L . -lcuda -lcudart -lcupti -lnvperf_host -lnvperf_target -Wl,-rpath,$ORIGIN
#cgo darwin,amd64 LDFLAGS: -L /usr/local/cuda/lib -L /usr/local/cuda/extras/CUPTI/lib
#cgo linux,amd64 linux,ppc64le LDFLAGS: -L /usr/local/cuda/lib64 -L /usr/local/cuda/extras/CUPTI/lib64
*/
import "C"
