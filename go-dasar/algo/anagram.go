package main

import (
	"fmt"
	"sort"
	"strings"
)

// Fungsi untuk mengecek apakah dua string merupakan anagram
func areAnagrams(str1, str2 string) bool {
	// Menghapus spasi dan mengubah huruf menjadi lowercase
	str1 = strings.ReplaceAll(str1, " ", "")
	str2 = strings.ReplaceAll(str2, " ", "")
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	// Mengubah string menjadi slice rune untuk dapat diurutkan
	runes1 := []rune(str1)
	runes2 := []rune(str2)

	// Mengurutkan slice rune
	sort.Slice(runes1, func(i, j int) bool {
		return runes1[i] < runes1[j]
	})

	sort.Slice(runes2, func(i, j int) bool {
		return runes2[i] < runes2[j]
	})

	// Membandingkan hasil pengurutan
	return string(runes1) == string(runes2)
}

// Fungsi untuk mengecek apakah semua string dalam array adalah anagram
func areAllAnagrams(strings []string) bool {
	// Mengecek setiap kombinasi string dalam array
	for i := 0; i < len(strings)-1; i++ {
		for j := i + 1; j < len(strings); j++ {
			if !areAnagrams(strings[i], strings[j]) {
				return false
			}
		}
	}
	return true
}

func main() {
	// Contoh penggunaan dengan array string
	arr := []string{"eat", "tea", "ate"}

	if areAllAnagrams(arr) {
		fmt.Println("Semua string dalam array adalah anagram.")
	} else {
		fmt.Println("Tidak semua string dalam array adalah anagram.")
	}
}
