package helpers

import (
	"bytes"
	"os"
)

func ReadFileBytes(reader *bytes.Reader, size int64) (buffer []byte, err error) {
	buffer = make([]byte, size)
	_, err = reader.Read(buffer)
	return
}

// PathExist check the path directory if exist
func PathExist(p string) bool {
	stat, err := os.Stat(p)
	return err != nil && stat.IsDir()
}
