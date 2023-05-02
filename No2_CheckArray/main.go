package main

func main() {
	var Array1 = [3]string{"c", "b", "c"}
	var Array2 = [3]string{"a", "d", "c"}

	println("Array 1 :")
	for i := 0; i < len(Array1); i++ {
		print(Array1[i], " ")
	}

	println("\nArray 2 :")
	for i := 0; i < len(Array1); i++ {
		print(Array2[i], " ")
	}
	println("\nCek Array : ")
	for i := 0; i < len(Array1); i++ {
		if Array1[i] != Array2[i] {
			println("Indexs Ke - ", i, "Tidak Sama")
		}
	}
}
