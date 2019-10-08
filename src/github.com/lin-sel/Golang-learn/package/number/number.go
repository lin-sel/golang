package number

/*

if you want to export any package function to another package then function first Letter should be Capital letter.

*/

// Isprime return boolean
func Isprime(num int32) bool {

	if num%2 == 0 || num%3 == 0 || num%5 == 0 {
		return false
	}

	return num > 1
}
