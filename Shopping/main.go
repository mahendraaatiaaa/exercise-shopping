package main

import (
	"errors"
	"fmt"
)

func PembayaranBarang(hargaTotal float64, metodePembayaran string, dicicil bool) error {
	// Check if hargaTotal is greater than 0
	if hargaTotal <= 0 {
		return errors.New("harga tidak bisa nol")
	}

	// Check if metodePembayaran is valid
	validMetode := map[string]bool{"cod": true, "transfer": true, "debit": true, "credit": true, "gerai": true}
	if !validMetode[metodePembayaran] {
		return errors.New("metode tidak dikenali")
	}

	// Check if dicicil is true and metodePembayaran is credit and hargaTotal >= 500.000
	if dicicil {
		if metodePembayaran != "credit" || hargaTotal < 500000 {
			return errors.New("cicilan tidak memenuhi syarat")
		}
	} else {
		// Check if dicicil is false and metodePembayaran is credit
		if metodePembayaran == "credit" {
			return errors.New("credit harus dicicil")
		}
	}

	return nil
}

func main() {
	// Example usage
	err := PembayaranBarang(550000, "credit", true)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Pembayaran berhasil.")
	}
}
