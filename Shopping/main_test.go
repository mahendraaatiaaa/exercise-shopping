package main

import "testing"

func TestPembayaranBarang(t *testing.T) {
	type args struct {
		hargaTotal       float64
		metodePembayaran string
		dicicil          bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Case 1", args{550000, "credit", true}, false},  // Valid case
		{"Case 2", args{0, "transfer", false}, true},     // Harga nol
		{"Case 3", args{600000, "debit", true}, true},    // Cicilan tidak memenuhi syarat
		{"Case 4", args{700000, "credit", true}, false},  // Valid case
		{"Case 5", args{800000, "credit", false}, true},  // Credit harus dicicil
		{"Case 6", args{300000, "unknown", false}, true}, // Metode tidak dikenali
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PembayaranBarang(tt.args.hargaTotal, tt.args.metodePembayaran, tt.args.dicicil); (err != nil) != tt.wantErr {
				t.Errorf("PembayaranBarang() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
