package main

import (
	"fmt"
)

type Transaksi struct {
	jenis  string
	jumlah int
}

type ATM struct {
	saldo   int
	riwayat []Transaksi
}

func (atm *ATM) tambahkanSaldo(jumlah int) {
	if jumlah <= 0 {
		fmt.Println("Jumlah yang ditambahkan harus positif.")
		return
	}
	atm.saldo += jumlah
	atm.riwayat = append(atm.riwayat, Transaksi{"Tambah Saldo", jumlah})
	fmt.Println("Saldo Anda Saat Ini:", atm.saldo)
}

func (atm *ATM) tarikSaldo(jumlah int) {
	if jumlah <= 0 {
		fmt.Println("Jumlah yang ditarik harus positif.")
		return
	}
	if jumlah > atm.saldo {
		fmt.Println("Saldo tidak mencukupi.")
		return
	}
	atm.saldo -= jumlah
	atm.riwayat = append(atm.riwayat, Transaksi{"Tarik Saldo", jumlah})
	fmt.Println("Saldo Anda Saat Ini:", atm.saldo)
}

func (atm *ATM) cekSaldo() {
	fmt.Println("Saldo Anda Saat Ini:", atm.saldo)
}

func (atm *ATM) riwayatTransaksi() {
	fmt.Println("Riwayat Transaksi:")
	if len(atm.riwayat) == 0 {
		fmt.Println("Belum ada transaksi.")
		return
	}
	for _, transaksi := range atm.riwayat {
		fmt.Printf("- %s sebesar Rp%d\n", transaksi.jenis, transaksi.jumlah)
	}
}

func main() {
	var username, password string
	var pilihan int

	atm := ATM{saldo: 3500000}

	correctUsername := "Dzihni"
	correctPassword := "2406421516"

	fmt.Println("==Selamat Datang di ATM==")
	fmt.Println("Silahkan login")

	fmt.Print("Input Username: ")
	fmt.Scanf("%s\n", &username)

	fmt.Print("Input Password: ")
	fmt.Scanf("%s\n", &password)

	if username == correctUsername && password == correctPassword {
		fmt.Println("Login berhasil!")

		for {
			fmt.Println("==Menu ATM==")
			fmt.Println("1. Informasi Akun")
			fmt.Println("2. Cek Saldo")
			fmt.Println("3. Tambah Saldo")
			fmt.Println("4. Tarik Saldo")
			fmt.Println("5. Riwayat Transaksi")
			fmt.Println("6. Keluar")
			fmt.Print("Pilih menu: ")
			fmt.Scan(&pilihan)

			switch pilihan {
			case 1:
				fmt.Println("==Anda Memilih Informasi Akun==")
				fmt.Println("Berikut Informasi Akun Anda: ")
				fmt.Println("Nama Lengkap = Dzihni Hanifah")
				fmt.Println("NPM = 2406421516")
				fmt.Println("Username = Dzihni")
				fmt.Println("Password = 2406421516")
				fmt.Println("Jenis Kelamin = Perempuan")
				fmt.Println("Alamat = Bogor")

			case 2:
				fmt.Println("==Anda Memilih Cek Saldo==")
				atm.cekSaldo()

			case 3:
				fmt.Println("==Anda Memilih Tambah Saldo==")
				var jumlah int
				fmt.Print("Masukkan jumlah yang ingin ditambahkan: ")
				fmt.Scan(&jumlah)
				atm.tambahkanSaldo(jumlah)

			case 4:
				fmt.Println("==Anda Memilih Tarik Saldo==")
				var jumlah int
				fmt.Print("Masukkan jumlah yang ingin ditarik: ")
				fmt.Scan(&jumlah)
				atm.tarikSaldo(jumlah)

			case 5:
				fmt.Println("==Anda Memilih Riwayat Transaksi Saldo==")
				atm.riwayatTransaksi()

			case 6:
				fmt.Println("Terima kasih telah menggunakan ATM.")
				return

			default:
				fmt.Println("Pilihan tidak valid.")
			}
		}
	} else {
		fmt.Println("Username atau Password salah.")
	}
}
