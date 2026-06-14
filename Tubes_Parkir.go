package main

import (
	"fmt"
	"strings"
)

const NMAX = 100
const LEBAR = 58

type Petugas struct {
	username string
	password string
}

type Admin struct {
	username string
	password string
}

type Kendaraan struct {
	platNomor string
	jenis     string
	durasi    int
	biaya     int
	status    string
}

type ArrPetugas [NMAX]Petugas
type ArrKendaraan [NMAX]Kendaraan

var (
	admin         Admin
	dataPetugas   ArrPetugas
	nPetugas      int
	dataKendaraan ArrKendaraan
	nKendaraan    int
)

func cetakBanner() {
	garisBanner := strings.Repeat("><", LEBAR/2)

	fmt.Println("\n" + garisBanner)
	fmt.Println("           SISTEM MANAJEMEN PARKIR PANANDYA MALL          ")
	fmt.Println(garisBanner)
}

func cetakGaris() {
	fmt.Println(strings.Repeat("-", LEBAR))
}

func cetakJudulMenu(judul string) {
	sisaRuang := LEBAR - len(judul) - 2

	garis := strings.Repeat("=", sisaRuang/2)

	fmt.Printf("%s %s %s\n", garis, judul, garis)
}

func cariPetugas(user, pass string) int {
	idx := -1
	i := 0
	for i < nPetugas && idx == -1 {
		if dataPetugas[i].username == user && dataPetugas[i].password == pass {
			idx = i
		}
		i++
	}
	return idx
}

func tambahPetugas() {
	if nPetugas >= NMAX {
		fmt.Println("Data petugas penuh.")
		return
	}

	var user string

	fmt.Print("Username : ")
	fmt.Scan(&user)

	i := 0
	for i < nPetugas {
		if dataPetugas[i].username == user {
			fmt.Println("Username sudah digunakan.")
			return
		}
		i++
	}

	dataPetugas[nPetugas].username = user

	fmt.Print("Password : ")
	fmt.Scan(&dataPetugas[nPetugas].password)

	nPetugas++

	fmt.Println("Petugas berhasil ditambahkan.")
}

func editPetugas() {
	var user string
	fmt.Print("Masukkan Username Petugas : ")
	fmt.Scan(&user)

	i := 0
	for i < nPetugas {
		if dataPetugas[i].username == user {
			var pilihanEdit string

			fmt.Println("\nData Petugas Ditemukan!")
			fmt.Println("1. Edit Username")
			fmt.Println("2. Edit Password")
			fmt.Print("Pilih data yang akan diubah (1/2): ")
			fmt.Scan(&pilihanEdit)

			if pilihanEdit == "1" {
				fmt.Print("Username Baru : ")
				fmt.Scan(&dataPetugas[i].username)
				fmt.Println("Username berhasil diubah.")
			} else if pilihanEdit == "2" {
				fmt.Print("Password Baru : ")
				fmt.Scan(&dataPetugas[i].password)
				fmt.Println("Password berhasil diubah.")
			} else {
				fmt.Println("Pilihan tidak valid. Batal mengedit.")
			}

			return
		}
		i++
	}

	fmt.Println("Petugas tidak ditemukan.")
}

func hapusPetugas() {
	var user string

	fmt.Print("Masukkan Username Petugas : ")
	fmt.Scan(&user)

	i := 0
	for i < nPetugas {
		if dataPetugas[i].username == user {

			j := i
			for j < nPetugas-1 {
				dataPetugas[j] = dataPetugas[j+1]
				j++
			}

			nPetugas--
			fmt.Println("Petugas berhasil dihapus.")
			return
		}
		i++
	}

	fmt.Println("Petugas tidak ditemukan.")
}

func lihatPetugas() {
	fmt.Println()
	cetakJudulMenu("DAFTAR PETUGAS")

	i := 0
	for i < nPetugas {
		fmt.Println("Username :", dataPetugas[i].username)
		fmt.Println("Password :", dataPetugas[i].password)
		cetakGaris()
		i++
	}
}

func binarySearchKendaraan(plat string) int {
	kiri := 0
	kanan := nKendaraan - 1
	idx := -1

	for kiri <= kanan && idx == -1 {
		tengah := (kiri + kanan) / 2
		if dataKendaraan[tengah].platNomor == plat {
			idx = tengah
		} else if dataKendaraan[tengah].platNomor < plat {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return idx
}

func sortPlatAsc() {
	i := 0
	for i < nKendaraan-1 {
		minIdx := i
		j := i + 1
		for j < nKendaraan {
			if dataKendaraan[j].platNomor < dataKendaraan[minIdx].platNomor {
				minIdx = j
			}
			j++
		}
		temp := dataKendaraan[i]
		dataKendaraan[i] = dataKendaraan[minIdx]
		dataKendaraan[minIdx] = temp
		i++
	}
}

func sortBiayaAsc() {
	i := 0
	for i < nKendaraan-1 {
		min := i
		j := i + 1
		for j < nKendaraan {
			if dataKendaraan[j].biaya < dataKendaraan[min].biaya {
				min = j
			}
			j++
		}
		temp := dataKendaraan[i]
		dataKendaraan[i] = dataKendaraan[min]
		dataKendaraan[min] = temp
		i++
	}
}

func sortBiayaDesc() {
	i := 1
	for i < nKendaraan {
		key := dataKendaraan[i]
		j := i - 1

		for j >= 0 && dataKendaraan[j].biaya < key.biaya {
			dataKendaraan[j+1] = dataKendaraan[j]
			j--
		}

		dataKendaraan[j+1] = key
		i++
	}
}

func kendaraanMasuk() {
	if nKendaraan < NMAX {
		var plat, jenis string
		fmt.Print("Masukkan Plat Nomor (Tanpa Spasi, ex: BK1234AB) : ")
		fmt.Scan(&plat)
		fmt.Print("Masukkan Jenis (Mobil/Motor)                    : ")
		fmt.Scan(&jenis)

		dataKendaraan[nKendaraan].platNomor = plat
		dataKendaraan[nKendaraan].jenis = jenis
		dataKendaraan[nKendaraan].status = "PARKIR"
		dataKendaraan[nKendaraan].biaya = 0
		nKendaraan++

		fmt.Println("Berhasil! Palang terbuka. Kendaraan berhasil masuk.")
	} else {
		fmt.Println("Mohon maaf, kapasitas parkir penuh!")
	}
}

func kendaraanKeluar() {
	var plat string
	fmt.Print("Masukkan Plat Nomor Kendaraan Keluar : ")
	fmt.Scan(&plat)

	sortPlatAsc()
	idx := binarySearchKendaraan(plat)

	if idx != -1 && dataKendaraan[idx].status == "PARKIR" {
		var durasi int
		fmt.Print("Berapa jam kendaraan terparkir?      : ")
		fmt.Scan(&durasi)

		if durasi <= 0 {
			durasi = 1
		}

		dataKendaraan[idx].durasi = durasi
		dataKendaraan[idx].status = "KELUAR"

		if dataKendaraan[idx].jenis == "Mobil" {
			dataKendaraan[idx].biaya = 5000 + (durasi * 2000)
		} else {
			dataKendaraan[idx].biaya = 2000 + (durasi * 1000)
		}

		cetakGaris()
		fmt.Println("Rincian Pembayaran:")
		fmt.Printf("Plat Nomor  : %s\n", dataKendaraan[idx].platNomor)
		fmt.Printf("Total Biaya : Rp%d\n", dataKendaraan[idx].biaya)
		fmt.Println("Pembayaran sukses. Palang terbuka.")
	} else {
		fmt.Println("Kendaraan tidak ditemukan atau sudah keluar.")
	}
}

func cariKendaraan() {
	var plat string

	fmt.Print("Masukkan Plat Nomor : ")
	fmt.Scan(&plat)

	sortPlatAsc()
	idx := binarySearchKendaraan(plat)

	if idx != -1 {
		fmt.Println("\nData Kendaraan Ditemukan")
		fmt.Println("Plat   :", dataKendaraan[idx].platNomor)
		fmt.Println("Jenis  :", dataKendaraan[idx].jenis)
		fmt.Println("Durasi :", dataKendaraan[idx].durasi)
		fmt.Println("Biaya  :", dataKendaraan[idx].biaya)
		fmt.Println("Status :", dataKendaraan[idx].status)
	} else {
		fmt.Println("Kendaraan tidak ditemukan.")
	}
}

func editKendaraan() {
	var plat string

	fmt.Print("Masukkan Plat Nomor yang akan diedit : ")
	fmt.Scan(&plat)

	sortPlatAsc()
	idx := binarySearchKendaraan(plat)

	if idx != -1 {
		fmt.Print("Plat Baru                : ")
		fmt.Scan(&dataKendaraan[idx].platNomor)

		fmt.Print("Jenis Baru (Mobil/Motor) : ")
		fmt.Scan(&dataKendaraan[idx].jenis)

		fmt.Print("Durasi Baru              : ")
		fmt.Scan(&dataKendaraan[idx].durasi)

		if dataKendaraan[idx].durasi <= 0 {
			dataKendaraan[idx].durasi = 1
		}

		if dataKendaraan[idx].jenis == "Mobil" {
			dataKendaraan[idx].biaya = 5000 + dataKendaraan[idx].durasi*2000
		} else {
			dataKendaraan[idx].biaya = 2000 + dataKendaraan[idx].durasi*1000
		}

		fmt.Println("\nData berhasil diubah.")
		fmt.Printf("Total Biaya yang baru: Rp%d\n", dataKendaraan[idx].biaya)
	} else {
		fmt.Println("Kendaraan tidak ditemukan.")
	}
}

func cetakLaporanPendapatan() {
	var pilih, urut string

	fmt.Println()
	cetakJudulMenu("FILTER KENDARAAN")
	fmt.Println("1. Semua")
	fmt.Println("2. Mobil")
	fmt.Println("3. Motor")
	fmt.Print("Pilih : ")
	fmt.Scan(&pilih)

	fmt.Println()
	cetakJudulMenu("PENGURUTAN BIAYA")
	fmt.Println("1. Terbesar ke Terkecil (Insertion Sort)")
	fmt.Println("2. Terkecil ke Terbesar (Selection Sort)")
	fmt.Print("Pilih : ")
	fmt.Scan(&urut)

	if urut == "1" {
		sortBiayaDesc()
	} else if urut == "2" {
		sortBiayaAsc()
	} else {
		fmt.Println("Pilihan tidak valid!")
		return
	}

	fmt.Println()
	cetakGaris()
	fmt.Printf("%-12s %-8s %-8s %-10s\n", "PLAT", "JENIS", "DURASI", "BIAYA")
	cetakGaris()

	i := 0
	for i < nKendaraan {
		if pilih == "1" ||
			(pilih == "2" && dataKendaraan[i].jenis == "Mobil") ||
			(pilih == "3" && dataKendaraan[i].jenis == "Motor") {

			fmt.Printf("%-12s %-8s %-8d %-10d\n",
				dataKendaraan[i].platNomor,
				dataKendaraan[i].jenis,
				dataKendaraan[i].durasi,
				dataKendaraan[i].biaya)
		}
		i++
	}

	cetakGaris()
}

func loginAdmin() {
	var user, pass string

	fmt.Print("Username Admin : ")
	fmt.Scan(&user)

	fmt.Print("Password Admin : ")
	fmt.Scan(&pass)

	if user == admin.username && pass == admin.password {
		menuAdmin := ""

		for menuAdmin != "5" {
			fmt.Println()
			cetakGaris()
			cetakJudulMenu("MENU ADMIN")
			fmt.Println("1. Tambah Petugas")
			fmt.Println("2. Edit Petugas")
			fmt.Println("3. Hapus Petugas")
			fmt.Println("4. Lihat Petugas")
			fmt.Println("5. Logout")
			fmt.Print("Pilih Menu (1-5): ")
			fmt.Scan(&menuAdmin)

			if menuAdmin == "1" {
				tambahPetugas()
			} else if menuAdmin == "2" {
				editPetugas()
			} else if menuAdmin == "3" {
				hapusPetugas()
			} else if menuAdmin == "4" {
				lihatPetugas()
			} else if menuAdmin == "5" {
				fmt.Println("Logout Admin.")
			} else {
				fmt.Println("Pilihan tidak valid.")
			}
		}
	} else {
		fmt.Println("Username atau Password Admin salah!")
	}
}

func main() {
	admin = Admin{"admin", "sipalingadmin"}
	pilihanUtama := ""

	for pilihanUtama != "3" {
		cetakBanner()
		fmt.Println("1. Login Admin")
		fmt.Println("2. Login Petugas")
		fmt.Println("3. Tutup Aplikasi")
		fmt.Print("PILIH MENU (1-3): ")
		fmt.Scan(&pilihanUtama)

		if pilihanUtama == "1" {
			loginAdmin()
		} else if pilihanUtama == "2" {
			var user, pass string

			fmt.Print("\nUsername Petugas : ")
			fmt.Scan(&user)
			fmt.Print("Password Petugas : ")
			fmt.Scan(&pass)

			idxPetugas := cariPetugas(user, pass)

			if idxPetugas != -1 {
				fmt.Println("\nSelamat datang,", dataPetugas[idxPetugas].username)
				menuPetugas := ""

				for menuPetugas != "6" {
					fmt.Println()
					cetakGaris()
					cetakJudulMenu("MENU PETUGAS")
					fmt.Println("1. Kendaraan Masuk")
					fmt.Println("2. Kendaraan Keluar")
					fmt.Println("3. Edit Kendaraan")
					fmt.Println("4. Cari Kendaraan")
					fmt.Println("5. Cetak Laporan")
					fmt.Println("6. Logout")
					fmt.Print("PILIH MENU (1-6): ")
					fmt.Scan(&menuPetugas)

					if menuPetugas == "1" {
						kendaraanMasuk()
					} else if menuPetugas == "2" {
						kendaraanKeluar()
					} else if menuPetugas == "3" {
						editKendaraan()
					} else if menuPetugas == "4" {
						cariKendaraan()
					} else if menuPetugas == "5" {
						cetakLaporanPendapatan()
					} else if menuPetugas == "6" {
						fmt.Println("Logout berhasil. Sampai jumpa!")
					} else {
						fmt.Println("Pilihan tidak valid.")
					}
				}
			} else {
				fmt.Println("Username atau Password salah!")
			}
		} else if pilihanUtama == "3" {
			fmt.Println("\nSistem ditutup. Terima kasih!")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
