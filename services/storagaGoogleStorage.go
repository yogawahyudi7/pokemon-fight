package services

import (
	"context"
	"fmt"
	"io"
	"os"
	"pokemon-fight/configs"
	"pokemon-fight/constants"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type ServiceInterface interface {
	UploadImagePokemonGCS(pokemonId string, filePtch string) error
}

type Services struct {
	config *configs.ServerConfig
}

func NewServices(config *configs.ServerConfig) *Services {
	return &Services{config: config}
}

func (s *Services) UploadImagePokemonGCS(pokemonId string, filePtch string) error {

	bucketName := "poke_fight_club"
	pokemonIdImagePath := "image/pokemonId"
	credentials := "credential/crested-return-398003-dade99e8c333.json"

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentials))
	if err != nil {
		return fmt.Errorf("gagal membuat klien gcs: %v", err)
	}
	defer client.Close()

	// Buka bucket
	bucket := client.Bucket(bucketName)

	buketPath := strings.Join([]string{pokemonIdImagePath, pokemonId}, "/")

	t := time.Now()
	timeFile := t.Format("20060102150405")
	filenameJoin := []string{constants.TypeIMG, "-", timeFile}
	filename := strings.Join(filenameJoin, "")

	objectName := fmt.Sprintf("%v/%v", buketPath, filename)

	// Buka file lokal
	file, err := os.Open(filePtch)
	if err != nil {
		return fmt.Errorf("gagal membuka file lokal: %v", err)
	}
	defer file.Close()

	// Buat objek di google storage
	obj := bucket.Object(objectName)

	// Buat penulis objek
	writer := obj.NewWriter(ctx)

	// Salin isi file lokal ke objek google storage
	if _, err := io.Copy(writer, file); err != nil {
		return fmt.Errorf("gagal mengunggah file: %v", err)
	}

	// Tutup penulis objek
	if err := writer.Close(); err != nil {
		return fmt.Errorf("gagal menutup penulis objek: %v", err)
	}

	return nil
}
