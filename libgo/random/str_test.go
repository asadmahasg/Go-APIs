package random_test

import (
	"fmt"
	"project/libgo/random"
)

func ExampleHexStr() {
	for i := 0; i < 5; i++ {
		token := random.HexStr(16)
		fmt.Println(token)
	}
	// Output:
	// 67aab2d956bd7cc621af22cfb169cba8
	// 226eeb52947edbf3e97d1e6669e212c2
	// 5f3615e95d103d14ffb5b655aa0eec1e
	// ff3ab4efbd74025b87b14b59422d304c
	// a6705813c174ca73ed795ea0bab12726
}

func ExampleString() {
	for i := 0; i < 5; i++ {
		token := random.String(16)
		fmt.Println(token)
	}
	// Output:
	// 7EbxkrHc1l3Ahmyr
	// I5XH2gc1EEHgbmGI
	// GlCycMpsxGkn9cDQ
	// U2OfBDQoak0z8FwV
	// kDX1m81u14YwEiCY
}
