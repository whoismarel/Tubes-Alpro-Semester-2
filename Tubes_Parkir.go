package main
import (
	"fmt"
	"strings"
)
const MAX = 100
type Petugas struct {
	username string
	password string
	nama     string
}
type Kendaraan struct {
	platNomor string
	jenis     string
	durasi    int
	biaya     int
	status    string
}
type ArrPetugas [MAX]Petugas
type ArrKendaraan [MAX]Kendaraan
var (
	dataPetugas   ArrPetugas
	nPetugas      int
	dataKendaraan ArrKendaraan
	nKendaraan    int
)

func cetakBanner() {
	fmt.Println("\n><><><><><><><><><><><><><><><><><><><><><><><><><><><><><")
	fmt.Println("           SISTEM MANAJEMEN PARKIR PANANDYA MALL     ")
	fmt.Println("><><><><><><><><><><><><><><><><><><><><><><><><><><><><><")
}

func cetakGaris() {
	fmt.Println(strings.Repeat("-", 58))
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
	if nKendaraan < MAX {
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

func cetakLaporanPendapatan() {
	sortBiayaDesc()
	
	fmt.Println("\n LAPORAN KENDARAAN")
	cetakGaris()
	fmt.Printf("%-12s | %-8s | %-8s | %-15s\n", "PLAT NOMOR", "JENIS", "DURASI", "TOTAL BIAYA (Rp)")
	cetakGaris()
	
	i := 0
	for i < nKendaraan {
		fmt.Printf("%-12s | %-8s | %-8d | %-15d\n", 
			dataKendaraan[i].platNomor, 
			dataKendaraan[i].jenis, 
			dataKendaraan[i].durasi, 
			dataKendaraan[i].biaya)
		i++
	}
	cetakGaris()
}

func main(){
	dataPetugas[0] = Petugas{"Poltak", "poltak123", "Pak Poltak"}
	nPetugas = 1

	pilihanUtama := -1
	for {
		cetakBanner()
		fmt.Println("1. Login Petugas")
		fmt.Println("2. Tutup Aplikasi")
		fmt.Print("PILIH MENU (1/2): ")
		fmt.Scan(&pilihanUtama)

		if pilihanUtama == 1 {
			var user, pass string
			fmt.Print("\nUsername : ")
			fmt.Scan(&user)
			fmt.Print("Password : ")
			fmt.Scan(&pass)

			idxPetugas := cariPetugas(user, pass)
			
			if idxPetugas != -1 {
				fmt.Println("><><><><><><><><><><><><><><><><><><><><><><><><><><><><><")
				fmt.Println("               Selamat datang", dataPetugas[idxPetugas].nama, ":)")
				
				menuPetugas := -1
				for menuPetugas != 4 {
					cetakGaris()
					fmt.Println("                    MENU PETUGAS TIKET                 ")
					fmt.Println("><><><><><><><><><><><><><><><><><><><><><><><><><><><><><")
					fmt.Println("1. Kendaraan Masuk")
					fmt.Println("2. Kendaraan Keluar (Hitung Biaya)")
					fmt.Println("3. Cetak Laporan Pendapatan")
					fmt.Println("4. Logout")
					fmt.Print("PILIH MENU (1-4): ")
					fmt.Scan(&menuPetugas)

					if menuPetugas == 1 {
						kendaraanMasuk()
					} else if menuPetugas == 2 {
						kendaraanKeluar()
					} else if menuPetugas == 3 {
						cetakLaporanPendapatan()
					} else if menuPetugas == 4 {
						fmt.Println("Log out berhasil. Sampai jumpa!")
					} else {
						fmt.Println("Pilihan tidak valid.")
					}
				}
			} else {
				fmt.Println(" Username atau Password salah!")
			}
		} else if pilihanUtama == 2 {
			fmt.Println("\nSistem ditutup. Terima kasih!")
			break
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
