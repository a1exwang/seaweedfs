// +build linux

package storage

import (
	"github.com/ncw/directio"
	"os"
	"syscall"

	"github.com/chrislusf/seaweedfs/weed/glog"
)

func createVolumeFile(fileName string, preallocate int64, useDirectIO bool) (file *os.File, e error) {
	if useDirectIO {
		file, e = directio.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	} else {
		file, e = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	}
	if preallocate != 0 {
		syscall.Fallocate(int(file.Fd()), 1, 0, preallocate)
		glog.V(0).Infof("Preallocated %d bytes disk space for %s", preallocate, fileName)
	}
	return file, e
}
