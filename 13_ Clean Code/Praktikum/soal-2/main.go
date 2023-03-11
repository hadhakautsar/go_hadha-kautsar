package main

//Setiap huruf di awal kata dalam nama variabel dan fungsi menggunakan huruf besar (CamelCase).
type kendaraan struct {
	roda      int //variabel totalRoda diganti menjadi roda
	kecepatan int //Variabel kecepatan diubah dari kecepatanperjam
}

type mobil struct {
	kendaraan
}

func (m *mobil) berjalan() {
	m.tambahKecepatan(10) //Variabel tambahKecepatan diubah dari tambahkecepatan
}

func (m *mobil) tambahKecepatan(kecepatanBaru int) {
	m.kecepatan += kecepatanBaru //mengubah operator aritmatika
}

func main() {
	//Penggunaan CamelCase pada nama variabel mobilCepat dan mobilLamban
	mobilCepat := mobil{}
	mobilCepat.berjalan()
	mobilCepat.berjalan()
	mobilCepat.berjalan()

	mobilLamban := mobil{}
	mobilLamban.berjalan()
}
