package utils

import "os"

func FileClear(f *os.File) error {
	_, err := f.Seek(0, 0)
	if err != nil {
		return err
	}

	err = f.Truncate(0)
	if err != nil {
		return err
	}

	return nil
}
