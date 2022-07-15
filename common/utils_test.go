//
// Licensed Materials - Property of IBM
//
// 5737-I09
//
// Copyright IBM Corp. 2022 All Rights Reserved.
// US Government Users Restricted Rights - Use, duplication or
// disclosure restricted by GSA ADP Schedule Contract with IBM Corp
//
package common

import (
	"fmt"
	"testing"
)

func TestTempFile(t *testing.T) {
	newFile := TempFile([]byte("Carsten"))
	fmt.Println(newFile)
}
