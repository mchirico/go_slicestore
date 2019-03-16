package main

import (
	"fmt"
	"github.com/mchirico/go_slicestore/pkg"
)

func main() {
	d := pkg.DS{}
	q := pkg.Query{}
	d.GetStruct(q.Query())
	fmt.Printf("%v", d.ResponseData.VaultUsageReport.Vaults[0].Name)

}

