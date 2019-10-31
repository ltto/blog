package conf

import (
	"blog/system"
	"blog/tools"
	"blog/tools/git"
	"bufio"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var (
	Nav      = Navs{Dir: system.Conf.Blog.Dir}
	AllBlog  []*BlogInfo
	YearBlog map[string][]*BlogInfo
	MMP      = map[string]string{
		"1":  "Jan",
		"2":  "Feb",
		"3":  "Mar",
		"4":  "Apr",
		"5":  "May",
		"6":  "Jun",
		"7":  "Jul",
		"8":  "Aug",
		"9":  "Sept",
		"10": "Oct",
		"11": "Nov",
		"12": "Dec",
	}
)

func init() {
	err := git.Pull(Nav.Dir)
	if err != nil {
		panic(err)
	}
	err = Nav.Reload()
	if err != nil {
		panic(err)
	}
}

type BlogInfo struct {
	Change   int
	Nav      string
	FileName string
	Desc     []tools.String
	Title    string
	M        string
	Y        string
	D        string
	Time     string
}
type Navs struct {
	nav   []string
	links map[string][]*BlogInfo
	Dir   string
}

func (n *Navs) Navs() []string {
	return n.nav
}

func (n *Navs) Links(nav string) []*BlogInfo {
	return n.links[nav]
}
func (n *Navs) LinksAll() map[string][]*BlogInfo {
	return n.links
}

func (n *Navs) Reload() error {
	navs, err := ioutil.ReadDir(n.Dir)
	if err != nil {
		return err
	}
	n.nav = make([]string, 0, len(navs)+2)
	linksAll := make([]*BlogInfo, 0, 16)
	linksYear := make(map[string][]*BlogInfo)
	n.links = make(map[string][]*BlogInfo, len(navs))
	for _, nav := range navs {
		if nav.IsDir() && !strings.HasPrefix(nav.Name(), ".") {
			n.nav = append(n.nav, nav.Name())
			links, err := ioutil.ReadDir(path.Join(n.Dir, nav.Name()))
			if err != nil {
				return err
			}
			n.links[nav.Name()] = make([]*BlogInfo, 0, len(links))
			for _, link := range links {
				if !link.IsDir() && !strings.HasPrefix(link.Name(), ".") && (strings.HasSuffix(link.Name(), ".md") || strings.HasSuffix(link.Name(), ".MD")) {
					log, err := git.Log(path.Join(n.Dir, nav.Name()), link.Name())
					if err != nil {
						return err
					}
					desc, err := getDesc(path.Join(n.Dir, nav.Name(), link.Name()))
					if err != nil {
						return err
					}
					blogInfo := BlogInfo{
						Change:   len(log),
						Nav:      nav.Name(),
						FileName: log[0].File.Name(),
						Desc:     desc,
						D:        log[0].Day,
						M:        log[0].Month,
						Y:        log[0].Year,
						Time:     log[0].Time,
						Title:    tools.PathToName(log[0].File.Name())}
					if _, ok := linksYear[blogInfo.Y]; !ok {
						linksYear[blogInfo.Y] = make([]*BlogInfo, 0, 16)
					}
					linksYear[blogInfo.Y] = append(linksYear[blogInfo.Y], &blogInfo)
					n.links[nav.Name()] = append(n.links[nav.Name()], &blogInfo)
					linksAll = append(linksAll, &blogInfo)
				}
			}
		}
	}
	AllBlog = linksAll
	YearBlog = linksYear
	return nil
}

func getDesc(file string) ([]tools.String, error) {
	open, err := os.Open(file)
	if err != nil {
		return make([]tools.String, 0), err
	}
	defer open.Close()
	scan := bufio.NewScanner(open)
	desc := make([]tools.String, 3)
	for i := 0; scan.Scan(); i++ {
		text := strings.TrimSpace(scan.Text())
		if text == "" || strings.HasPrefix(text, "![") {
			i--
		} else {
			s := tools.String(text)
			desc[i] = *s.TrimLeft("####").TrimLeft("##").TrimLeft("#").TrimLeft("+").TrimLeft(">")
		}
		if i == 2 {
			break
		}
	}
	return desc, nil
}
