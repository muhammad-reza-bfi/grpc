package file

import (
	"fmt"
	"os"

	errorc "github.com/elangreza14/grpc/internal/error"
)

type File struct {
	Name string
	Data []byte
}

func Read(location string) (*File, error) {
	file, err := os.Open(location)
	if err != nil {
		// handle the error here
		return nil, err
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return nil, err
	}

	return &File{
		Name: stat.Name(),
		Data: bs,
	}, err
}

func Write(file *File, location string) error {

	f, err := os.Create(fmt.Sprintf("%s/%s", location, file.Name))
	errorc.CheckErr(err)
	defer f.Close()

	_, err = f.Write(file.Data)
	errorc.CheckErr(err)

	return nil
}
