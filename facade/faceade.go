package facade

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

func ServerUploadFile(pokemonId string, file *multipart.FileHeader) (path string, err error) {

	serverPath := "storage/image"

	serverPathFile := strings.Join([]string{serverPath, pokemonId}, "/")

	// Buat direktori untuk menyimpan file yang diunggah
	err = os.MkdirAll(serverPathFile, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("gagal membuat direktori uploads : %v", err)
	}

	// Buka file yang diunggah
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("error: Gagal membuka file")
	}
	defer src.Close()

	// Simpan file di server
	serverFilePath := strings.Join([]string{serverPath, pokemonId, file.Filename}, "/")
	dst, err := os.Create(serverFilePath)
	if err != nil {
		return "", fmt.Errorf("error: gagal menyimpan file di server")
	}
	defer dst.Close()

	// Salin isi file ke file di server
	if _, err := io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("error: gagal menyalin isi file")
	}

	return serverFilePath, nil
}

func ServerRemoveFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return fmt.Errorf("gagal menghapus file yang sudah di upload: %v", err.Error())
	}
	return nil
}
