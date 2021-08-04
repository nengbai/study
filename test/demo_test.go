package test

import (
	"study/common"
	"testing"
)

func Json2Map_testing(t *testing.T) {
	j := `1`
	common.Json2Map(j)

}

func UpperCase_testing(t *testing.T) {
	J := `small`
	common.UpperCase(J)

}
