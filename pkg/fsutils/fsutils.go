package fsutils

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"syscall"
)

const (
	FILE_MODE      = 0644
	DIR_MODE       = 0755
	DIR_MOUNT_MODE = 0750
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

func CopyFile(r, w string) (count int64, err error) {
	var reader *os.File
	var writer *os.File

	if reader, err = os.OpenFile(r, os.O_RDONLY, FILE_MODE); err != nil {
		return
	}
	defer reader.Close()

	if writer, err = os.OpenFile(w, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, FILE_MODE); err != nil {
		return
	}
	defer writer.Close()

	buf := make([]byte, 8192, 8192)
	for {
		rn, re := reader.Read(buf[0:])
		if rn > 0 {
			wn, we := writer.Write(buf[0:rn])
			if we != nil {
				err = we
				break
			}
			if rn != wn {
				err = errors.New("short write")
				break
			}
			count += int64(wn)
		}
		if re != nil {
			if re != io.EOF {
				err = re
			}
			break
		}
	}
	return
}

func RmFile(f string) error {
	return os.Remove(f)
}

func RmDir(d string, recursive bool) error {
	if recursive {
		return os.RemoveAll(d)
	} else {
		return os.Remove(d)
	}
}

func MvFile(o, n string) error {
	return os.Rename(o, n)
}

func JoinPath(child ...string) string {
	return filepath.Join(child)
}

func MountBind(src, target string) error {
	// mount -o bind src target
	return nil
}

func UnMount(target string) error {
	return nil
}
