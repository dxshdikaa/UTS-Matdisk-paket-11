// Tugas Besar MTK Diskrit
// Dika Aldiansyah & Rifqi Anwar Sidiq
// Paket 11

package main

func main() {

	// =========================
	// Materi I : Himpunan
	// Paket 11
	// =========================
	A := []int{3, 10, 12, 20}
	B := []int{5, 10, 18}
	C := []int{10, 7}

	S := []int{4, 6, 8, 9, 10, 12}

	N := 85
	K := 13

	ProsesHimpunan(A, B, C, N)
	CariPasangan(S, K)

	// =========================
	// Materi III : Deret & Rekurens
	// Paket 11
	// =========================
	Rekurens(4, -3, 9)
	DeretGeometri(1, 0.8, 8)
}