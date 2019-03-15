package pkg

import (
	"encoding/json"
	"log"
	"sync"
)

type Vaults struct {
	Id          int    `json:"id"`
	Uuid        string `json:"uuid"`
	Name        string `json:"name"`
	StoragePool string `json:"storagePool"`
	Usage       int64  `json:"usage"`
	ReportStart string `json:"usageReportingStart"`
	ReportEnd   string `json:"usageReportingEnd"`
}

type DateRange struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type VaultRec struct {
	DateRange DateRange `json:"dateRange"`
	Vaults    []Vaults  `json:"vaults"`
}

type VaultUsageReport struct {
	VaultUsageReport VaultRec `json:"vaultUsageReport"`
}

type ResponseHeader struct {
	Now       int64  `json:"now"`
	Status    string `json:"status"`
	RequestId string `json:"requestId"`
}

type DS struct {
	sync.Mutex
	ResponseStatus string           `json:"responseStatus"`
	ResponseHeader ResponseHeader   `json:"responseHeader"`
	ResponseData   VaultUsageReport `json:"responseData"`
}

func (d *DS) GetStruct(input []byte) error {
	d.Lock()
	defer d.Unlock()

	err := json.Unmarshal(input, d)
	if err != nil {
		log.Fatalf("We can't read json")
	}
	return err
}
