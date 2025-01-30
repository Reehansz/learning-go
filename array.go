package main

func InitArray() [4]int {
	var arr [4]int
	for i := 0; i < 4; i++ {
		arr[i] = i + 1
	}
	return arr

}

func InitString() [5]string {
	var temp [5]string
	for i := 0; i < 5; i++ {
		temp[i] = (string(int('a') + i))
	}
	return temp
}
