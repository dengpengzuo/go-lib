package fsutils

import (
	"io"
	"os"
	"syscall"
)

const (
	FILE_MODE = 0644
	DIR_MODE  = 0755
)

func IsDir(filename string) bool {
	fi, err := sysStat(filename)
	return err == nil && fileMode(fi) == os.ModeDir
}

func IsSymLink(filename string) bool {
	fi, err := sysStat(filename)
	return err == nil && fileMode(fi) == os.ModeSymlink
}

func sysStat(filename string) (*syscall.Stat_t, error) {
	var fs syscall.Stat_t
	err := syscall.Stat(filename, &fs)
	if err != nil {
		return nil, &os.PathError{"stat", filename, err}
	}
	return &fs, nil
}

func fileMode(fs *syscall.Stat_t) os.FileMode {
	switch fs.Mode & syscall.S_IFMT {
	case syscall.S_IFBLK:
		return os.ModeDevice
	case syscall.S_IFCHR:
		return os.ModeDevice | os.ModeCharDevice
	case syscall.S_IFDIR:
		return os.ModeDir
	case syscall.S_IFIFO:
		return os.ModeNamedPipe
	case syscall.S_IFLNK:
		return os.ModeSymlink
	case syscall.S_IFSOCK:
		return os.ModeSocket
	}

	/*
	   if fs.Mode&syscall.S_ISGID != 0 {
	       return os.ModeSetgid
	   }
	   if fs.Mode&syscall.S_ISUID != 0 {
	       return os.ModeSetuid
	   }
	   if fs.Mode&syscall.S_ISVTX != 0 {
	       return os.ModeSticky
	   }
	*/
	return 0
}

func WriteFile(filename string, data []byte, append bool) error {
	flag := os.O_WRONLY | os.O_CREATE
	if append {
		flag = flag | os.O_APPEND
	} else {
		flag = flag | os.O_TRUNC
	}
	f, err := os.OpenFile(filename, flag, FILE_MODE)
	if err != nil {
		return err
	}
	_, err = f.Write(data)

	if err1 := f.Close(); err == nil {
		err = err1
	}

	return err
}

func CopyRW(reader io.Reader, writer io.Writer, close bool) (uint64, error) {
	return 0, nil
}
