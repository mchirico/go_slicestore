package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type pdata struct {
	name       map[string]string
	formatname map[string]string
	list       []string
	raw        []DS
}

func Pdata() pdata {
	p := pdata{}
	p.name = map[string]string{}
	p.formatname = map[string]string{}
	p.list = []string{}
	return p
}

func (p *pdata) Add(data []byte) {

	ds := DS{}
	err := json.Unmarshal(data, &ds)
	if err != nil {
		log.Printf("Can't Unmarshal in pkg.process")
	}

	for _, v := range ds.ResponseData.VaultUsageReport.Vaults {
		p.name[v.Name] = ""
	}

	p.raw = append(p.raw, ds)

}

func (p *pdata) FormatNames() {

	p.list = []string{}
	for str, _ := range p.name {

		re := regexp.MustCompile("[0-9]+$")
		r := re.FindAllString(str, -1)

		if len(r) < 1 {
			p.formatname[str] = str
			p.name[str] = str
			continue
		}

		idx := strings.LastIndex(str, r[0])
		i, err := strconv.Atoi(r[0])

		if err != nil {
			p.formatname[str] = str
			p.name[str] = str
			continue
		}

		name := fmt.Sprintf("%s%06d", str[0:idx], i)
		p.formatname[name] = str
		p.name[str] = name

	}

	for k, _ := range p.formatname {
		p.list = append(p.list, k)
	}

	sort.Slice(p.list, func(i, j int) bool {
		return p.list[i] < p.list[j]
	})

}
