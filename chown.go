// +build !linux

package logrotate

import (
	"os"
)

func chown(_ string, _ os.FileInfo) error {
	return nil
}
