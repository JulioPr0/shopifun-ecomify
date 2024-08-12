package main

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/JulioPr0/shopifun-ecomify/tree/basic_go/model"
)

func main() {
	fmt.Println("start main")
	// basics()

	fmt.Println("sum3Number(1, 2, 3) =", sum3Number(1, 2, 3)) // 6
	fmt.Println("sum3Number(3, 5, 7) =", sum3Number(3, 5, 7)) // 15
	fmt.Println("sum3Number(-10, 6, 4) =", sum3Number(-10, 6, 4)) // 0
	fmt.Println("=====")

	fmt.Println("mean3Number(1, 2, 3) =", mean3Number(1, 2, 3))       // 2.0
	fmt.Println("mean3Number(4, 4, 5) =", mean3Number(4, 4, 5))       // 4.333
	fmt.Println("mean3Number(10, 20, 50) =", mean3Number(10, 20, 50)) // 26.667
	fmt.Println("=====")

	fmt.Println("mean(1, 2, 3) =", mean([]int{1, 2, 3}))                                    // 2.0
	fmt.Println("mean(10, 20, 30, 40, 50) =", mean([]int{10, 20, 30, 40, 50}))              // 30.0
	fmt.Println("mean(1,2,3,4,5,6,7,8,9,10) =", mean([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})) // 4.5
	fmt.Println("=====")

	fmt.Println(`isPalindrome("katak") =`, isPalindrome("katak"))
	fmt.Println(`isPalindrome("golang") =`, isPalindrome("golang"))
	fmt.Println(`isPalindrome("1234567890987654321") =`, isPalindrome("1234567890987654321")) // true
	fmt.Println(`isPalindrome("1234567890887654321") =`, isPalindrome("1234567890887654321")) // false
	fmt.Println("=====")

	ans, found := findDuplicateNumber([]int{1, 2, 3, 3, 4, 5})
	fmt.Println("findDuplicateNumber(1,2,3,3,4,5) =", ans, found) // 3, true
	findDuplicateNumber([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 5})
	fmt.Println("findDuplicateNumber(1,2,3,4,5,6,7,8,9,5) =", ans, found) // 5, true
	fmt.Println("=====")

	fmt.Println("printTypeAndValue(12.34) =", printTypeAndValue(12))    // int: 12
	fmt.Println("printTypeAndValue(12.34) =", printTypeAndValue(12.34)) // float64: 12.34
	fmt.Println(`printTypeAndValue("tes") =`, printTypeAndValue("tes")) // string: tes
	fmt.Println("printTypeAndValue(true) =", printTypeAndValue(true))   // bool: true
	fmt.Println("=====")

	students := []model.Student{
		{Name: "Erwin", Score: 90},
		{Name: "a", Score: 75},
		{Name: "Jody", Score: 90},
		{Name: "b", Score: 70},
		{Name: "Irham", Score: 80},
		{Name: "c", Score: 65},
		{Name: "d", Score: 60},
		{Name: "Fatan", Score: 80},
	}
	fmt.Println("findStudents() =", findStudents(students, 80, false)) // 4 nama tidak terurut
	fmt.Println("findStudents() =", findStudents(students, 80, true))  // 4 nama terurut
	fmt.Println("findStudents() =", findStudents(students, 70, false)) // 6 nama tidak terurut
	fmt.Println("findStudents() =", findStudents(students, 70, true))  // 6 nama terurut
	fmt.Println("=====")

	st := students[0]
	changeName(&st, "Ethan")
	fmt.Println("st.Name =", st.Name)

	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	mergedSlice := merge2slices(slice1, slice2)
	fmt.Println(mergedSlice) // [1 2 3 4 5 6]
}

// func basics() {
// 	// variable string
// 	// var title string = "golang basics"
// 	title := "golang basics"
// 	fmt.Println("Hello, welcome to", title)
// 	fmt.Println("Hello,", "welcome", "to", title)

// 	// basic data types
// 	angka := 987                   // int
// 	angkaBerkoma := 12.34567       // float64
// 	angka32 := int32(angkaBerkoma) // int32; value = 12

// 	// angka, angkaBerkoma, angka32 := 987, 12.34567, int32(123) // alternative way of assignment

// 	fmt.Printf("angka = %d %f %d\n", angka, angkaBerkoma, angka32)
// 	fmt.Printf("angkaBerkoma = %.2f %.3f %.4f %.10f\n", angkaBerkoma, angkaBerkoma, angkaBerkoma, math.Pow(angkaBerkoma, 2))
// 	// x := fmt.Sprintf("%.2f", angkaBerkoma)

// 	var semangatBelajar bool // bool = false (zero-value) "", 0, 0.0, nil
// 	fmt.Printf("semangatBelajar = %t %T\n", semangatBelajar, semangatBelajar)
// 	semangatBelajar = true
// 	fmt.Printf("semangatBelajar = %t %T\n", semangatBelajar, semangatBelajar)

// 	// 2 way to write constants
// 	const PI = 3.14
// 	const ROLE_OWNER = "owner"
// 	const (
// 		STATUS_200 = "OK"
// 		STATUS_400 = "Bad Request"
// 	)
// 	fmt.Printf("PI = %f; STATUS = %s\n", PI, STATUS_200)

// 	fmt.Printf("title = %v %T\n", title, title) // print value, type
// 	title2 := []byte(title)     
// 	fmt.Printf("%v\n", title2)          			 
// 	fmt.Printf("%v\n", []byte(title))           // print []byte

// 	fmt.Println("=====")


// 	// convert int to string
// 	angkaString := strconv.Itoa(angka)
// 	// angkaString := fmt.Sprintf("%d", angka)
// 	fmt.Printf("angkaString = %s\n", angkaString)

// 	// convert string to int
// 	// angkaString = "qwerty"
// 	angka, err := strconv.Atoi(angkaString)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// convert string to float64
// 	angkaString = "123.456"
// 	angkaBerkoma, err = strconv.ParseFloat(angkaString, 64)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("angka = %d %f\n", angka, angkaBerkoma)
// 	fmt.Println("=====")
// }


// sum3Number adalah jumlah dari 3 angka (basic function)
func sum3Number(a int, b int, c int) int {
	// write code here
	return a + b + c
}

// mean3Number adalah nilai rata-rata dari 3 angka (parameter & return berbeda)
func mean3Number(a, b, c int) float64 {
	// write code here
	return float64((a + b + c)) / 3
}

// Cari nilai rata-rata data arr [1,2,3,4,5] = 3 (for-range)
func mean(arr []int) float64 {
	// write code here
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return float64(sum) / float64(len(arr))
}

// isPalindrome check wether str is palindrome or not. "katak" = true. (for-i, if)
func isPalindrome(str string) bool {
	for idx := range str {
		if str[idx] != str[len(str)-idx-1] {
			return false
		}
	}
	return true // contoh: "string: Hello" | "int: 123"
}

// which number is duplicate? [1,2,3,4,2,5] = 2 (map, for-range-map)
func findDuplicateNumber(arr []int) (int, bool) {
	m := map[int]bool{}
	for _, v := range arr {
			if m[v] {
					return v, true
			}
			m[v] = true
	}
	return -1, false
}


// print Type and Value of data (learn interface{} as generic)
func printTypeAndValue(data interface{}) string {
	// write code here
	// type inference struct {}
	// if value, ok := data.(int); ok {
	// 	return fmt.Sprintf("int: %d", value)
	// }
	// if value, ok := data.(float64); ok {
	// 	return fmt.Sprintf("float64: %f", value)
	// }
	// if value, ok := data.(string); ok {
	// 	return fmt.Sprintf("string: %s", value)
	// }
	// if value, ok := data.(bool); ok {
	// 	return fmt.Sprintf("int: %v", value)
	// }

	// package reflect
	t := reflect.TypeOf(data)
	switch t.String() {
	case "int":
			return fmt.Sprintf("int: %d", data)
	case "float64":
			return fmt.Sprintf("float64: %f", data)
	case "string":
			return fmt.Sprintf("string: %s", data)
	case "bool":
			return fmt.Sprintf("bool: %v", data)
	}
	return fmt.Sprintf("tipe_data: %v", data) // contoh: "string: Hello" | "int: 123"

	// return fmt.Sprint("%T: %v", data, data)
}

// Sort & filter Students. exam score >= minScore, sort by name (using struct)
func findStudents(students []model.Student, minScore float64, isSortByName bool) []model.Student {
	// write filter code here
	result := []model.Student{}
	for _, student := range students {
		if student.Score >= minScore {
			result = append(result, student)
		}
	}

	if isSortByName {
		// write sorting code here
		sort.Slice(result, func(i, j int) bool {
			return result[i].Name < result[j].Name
		})
	}

	return result
}

// mutate Student data (pointer)
func changeName(student *model.Student, name string) {
	// write code here
	student.Name = name
}

// jika waktu masih ada, bahas dibawah ini
func merge2slices(slice1 []int, slice2 []int) []int {
	// append 2 slices
	slice3 := append(slice1, slice2...)
	return slice3
}