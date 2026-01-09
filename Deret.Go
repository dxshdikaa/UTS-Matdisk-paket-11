// Tugas Besar MTK Diskrit
// Dika Aldiansyah & Rifqi Anwar Sidiq
// Paket 11

package main

import "fmt"

// Hitung relasi rekurens
// an = C1*a(n-1) + C2*a(n-2)
// a0 = 1, a1 = 1
func Rekurens(C1, C2, N int) {
	a := make([]int, N+1)

	a[0] = 1
	if N > 0 {
		a[1] = 1
	}

	fmt.Println("Relasi Rekurens")
	fmt.Printf("C1=%d, C2=%d, N=%d\n", C1, C2, N)

	for i := 2; i <= N; i++ {
		a[i] = C1*a[i-1] + C2*a[i-2]
	}

	for i := 0; i <= N; i++ {
		fmt.Printf("Suku %d : %d\n", i, a[i])
	}

	fmt.Printf("Hasil akhir suku ke-%d : %d\n\n", N, a[N])
}

// Hitung deret geometri
func DeretGeometri(a, r float64, N int) {
	total := 0.0
	suku := a

	for i := 0; i < N; i++ {
		total += suku
		suku *= r
	}

	sInf := a / (1 - r)
	persen := (total / sInf) * 100

	fmt.Println("Deret Geometri")
	fmt.Printf("a=%.2f, r=%.2f, N=%d\n", a, r, N)
	fmt.Printf("Jumlah %d suku : %.2f\n", N, total)
	fmt.Printf("Jumlah tak hingga : %.2f\n", sInf)
	fmt.Printf("Kedekatan : %.2f%%\n\n", persen)
}