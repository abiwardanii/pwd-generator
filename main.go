package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// Fungsi untuk menghasilkan kata sandi acak dengan panjang yang ditentukan
func generatePassword(length int) (string, error) {
	// Menentukan ukuran byte yang dibutuhkan untuk menghasilkan kata sandi dengan panjang tertentu
	byteLength := (length * 3) / 4

	// Membuat buffer byte yang sesuai dengan ukuran yang dibutuhkan
	buffer := make([]byte, byteLength)

	// Menghasilkan byte acak menggunakan crypto/rand
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	// Mengonversi byte acak menjadi format base64 yang aman
	password := base64.URLEncoding.EncodeToString(buffer)[:length]

	return password, nil
}

func main() {
	length := 16 // Panjang kata sandi yang diinginkan
	password, err := generatePassword(length)
	if err != nil {
		fmt.Println("Gagal menghasilkan kata sandi:", err)
		return
	}

	fmt.Println("Kata Sandi Baru:", password)
}
