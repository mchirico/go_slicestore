package pkg

import (
	"encoding/json"
	"log"
)


type ProcessStruct struct {
	origname string
	formatname string
	order  int
}


type Pdata struct {
	name []ProcessStruct
	raw []DS
}



func (p *Pdata) Add(data []byte)   {

	ds := DS{}
	err := json.Unmarshal(data, &ds)
	if err != nil {
		log.Printf("Can't Unmarshal in pkg.process")
	}

	ps := ProcessStruct{}
	for _,v := range ds.ResponseData.VaultUsageReport.Vaults {
		ps.origname = v.Name
	}

	p.name = append(p.name,ps)
	p.raw = append(p.raw, ds)



}


func SortNames() {

}