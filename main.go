package main

import (
	"crypto/rand"
	"fmt"
)

// Daftar karakter yang akan digunakan untuk membuat kata sandi
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+{}[]<>?/|\\"

// Fungsi untuk menghasilkan kata sandi acak dengan panjang yang ditentukan
func generatePassword(length int) (string, error) {
	// Menentukan ukuran byte yang dibutuhkan untuk menghasilkan kata sandi dengan panjang tertentu
	byteLength := length

	// Membuat buffer byte yang sesuai dengan ukuran yang dibutuhkan
	buffer := make([]byte, byteLength)

	// Menghasilkan byte acak menggunakan crypto/rand
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	// Menghasilkan kata sandi dari buffer byte dengan karakter dari charset
	password := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex := int(buffer[i]) % len(charset)
		password[i] = charset[randomIndex]
	}

	return string(password), nil
}

func main() {
	length := 20 // Panjang kata sandi yang diinginkan
	password, err := generatePassword(length)
	if err != nil {
		fmt.Println("Gagal menghasilkan kata sandi:", err)
		return
	}

	fmt.Println("Kata Sandi Baru:", password)
}
