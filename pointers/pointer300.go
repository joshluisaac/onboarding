// Copyright 2009 TheAuthors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package maint

import "fmt"

func Main300() {
	customerAge := 29

	fmt.Println(customerAge)
	fmt.Println(&customerAge)

	customerAgePtr := &customerAge

	fmt.Println(customerAgePtr)  // value of pointer
	fmt.Println(&customerAgePtr) // memory address of customerAgePtr
	fmt.Println(*customerAgePtr) // memory address dereferenced

	fmt.Println("+++++Start of modification+++++")
	*customerAgePtr = *customerAgePtr + 40
	fmt.Println(customerAge)
	fmt.Println(&customerAge)
	fmt.Println(customerAgePtr) // value of pointer
	fmt.Println(&customerAgePtr)
	fmt.Println(*customerAgePtr)
	fmt.Println("+++++End of modification+++++")

}
