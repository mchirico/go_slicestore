package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
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
		log.Printf("Can't Unmarshal in pkg.process:%s\n", data)
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

type wvault struct {
	name  string
	usage int64
	start string
	end   string
}

func (p *pdata) Write(file string) error {
	data := ""

	for row, raw := range p.raw {
		wv := map[string]wvault{}
		for _, vaults := range raw.ResponseData.VaultUsageReport.Vaults {
			wv[vaults.Name] = wvault{vaults.Name, vaults.Usage,
				vaults.ReportStart, vaults.ReportEnd}
		}

		if row == 0 {
			str := "date"
			sep := ""
			for _, v := range p.list {

				sep = ","

				str += sep + wv[p.formatname[v]].name

			}
			data += str + "\n"
		}

		str := ""
		sep := ""
		for idx, v := range p.list {
			if idx == 0 {

				t, _ := time.Parse("Mon, 02 Jan 2006 15:04:05 +0000",
					wv[p.formatname[v]].start)

				str += t.Format("01-02-2006") + "," + fmt.Sprintf("%d", wv[p.formatname[v]].usage)
			} else {
				sep = ","

				str += sep + fmt.Sprintf("%d", wv[p.formatname[v]].usage)
			}

		}
		data += str + "\n"

	}

	err := ioutil.WriteFile(file, []byte(data), 0600)
	if err != nil {
		log.Printf("error in yaml write: %v", err)
	}
	return err

}
