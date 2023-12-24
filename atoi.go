package piscine

// Atoi, bir string'i tamsayıya çeviren fonksiyondur.
// Fonksiyon, string içindeki rakamları bir tamsayıya dönüştürür.
// Ayrıca, baştaki boşlukları, artı ve eksi işaretleri dikkate alır.
// Eğer string geçerli bir tamsayıya dönüştürülemezse veya taşma olursa, 0 döndürülür.
func Atoi(s string) int {
	var result int  // Dönüştürülen tamsayı değeri
	multiplier := 1 // Çarpan, basamak değerini hesaplamak için kullanılır
	valid := true   // Geçerlilik kontrolü için bayrak
	sign := 1       // Artı veya eksi işaretini belirlemek için

	// String'i sondan başa doğru tarayarak tamsayıyı oluştur
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			// Rakamı tamsayıya dönüştür ve çarpan ile çarp
			digit := int(s[i] - '0')
			result += digit * multiplier
			multiplier *= 10
		} else if s[i] == '-' && i == 0 {
			// Eksi işareti bulunursa ve sıfırıncı indekste ise işaret değişkenini ayarla
			sign = -1
		} else if s[i] == '+' && i == 0 {
			// Artı işareti bulunursa ve sıfırıncı indekste ise işaret değişkenini ayarla
			sign = 1
		} else if s[i] != ' ' {
			// Rakam, boşluk, artı veya eksi dışında bir karakter varsa geçersiz durumu işaretle
			valid = false
			break
		}
	}

	// Geçersiz durum kontrolü veya taşma durumu kontrolü
	if !valid || (result == 0 && multiplier != 1) {
		return 0
	}

	// Sonucu işaretle ve döndür
	return result * sign
}
