package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/urfave/cli"
)

var (
	depth     int
	rootPath  string
	exts      []string
	excludes  []string
	showDraft bool
)

const (
	Ignore = iota
	Directory
	Draft
	Private
	Public
)

// Node .
type Node struct {
	Type     int
	Depth    int
	Name     string
	Title    string
	RelPath  string
	LinkName string
	Size     int64
	ModTime  time.Time
	Children []*Node
}

type FileList []os.FileInfo

func (fl FileList) Less(i, j int) bool {
	return fl[i].Name() < fl[j].Name()
}

func (fl FileList) Len() int {
	return len(fl)
}

func (fl FileList) Swap(i, j int) {
	fl[i], fl[j] = fl[j], fl[i]
}

// Summary .
func (n *Node) Summary(prefix string) {
	for _, c := range n.Children {
		if c.Type == Directory && len(c.Children) == 0 {
			continue
		} else if c.LinkName == "README.md" {
			fmt.Println(fmt.Sprintf("%s - [%s](%s)", prefix, c.Title, c.LinkName))
			fmt.Println("* * *")
		} else if c.Name == "README.md" || c.Name == "SUMMARY.md" {
			continue
		} else if c.LinkName == "" {
			fmt.Println(fmt.Sprintf("%s - %s", prefix, c.Name))
		} else if c.Type == Private {
			fmt.Println(fmt.Sprintf("%s - [[P] %s](%s)", prefix, c.Title, c.LinkName))
		} else if c.Type == Draft {
			fmt.Println(fmt.Sprintf("%s - [[D] %s](%s)", prefix, c.Title, c.LinkName))
		} else {
			fmt.Println(fmt.Sprintf("%s - [%s](%s)", prefix, c.Title, c.LinkName))
		}
		if len(c.Children) > 0 { // dir
			c.Summary(prefix + "  ")
		}
	}
}

// Tree .
func (n *Node) Tree(prefix string) {
	l := len(n.Children) - 1
	for i, c := range n.Children {
		p := fmt.Sprintf("%s├──", prefix)
		if i == l {
			p = fmt.Sprintf("%s└──", prefix)
		}
		fmt.Println(p, c.Name)
		if len(c.Children) > 0 { // dir
			if i == l {
				c.Tree(prefix + "    ")
			} else {
				c.Tree(prefix + "│   ")
			}
		}
	}
}

// Title .
func Title(data string) string {
	re := regexp.MustCompile(`[#]+\s*(.+)`)
	res := re.FindStringSubmatch(string(data))
	if len(res) > 1 {
		return res[1]
	}
	return ""
}

// Filter .
func (n *Node) Filter(files []os.FileInfo) {
	sort.Sort(FileList(files))
	for _, file := range files {
		c := &Node{
			Name:    file.Name(),
			ModTime: file.ModTime(),
			RelPath: filepath.Join(n.RelPath, file.Name()),
		}
		if file.IsDir() { // directory
			excluded := false
			for _, e := range excludes { // excluded
				if e == c.Name {
					excluded = true
				}
			}
			if !excluded {
				c.Depth = n.Depth + 1
				c.Type = Directory
				c.Read()
			}
		} else { // file
			ext := path.Ext(c.Name)
			c.Type = Ignore
			for _, e := range exts {
				if e == ext {
					switch c.Name[0] {
					case '_':
						c.Type = Draft
					case '.':
						c.Type = Private
					default:
						c.Type = Public
					}
					if c.Name == "README.md" {
						n.LinkName = filepath.Join(n.RelPath, c.Name)
					}
					c.LinkName = filepath.Join(n.RelPath, c.Name)
				}
			}
		}
		switch c.Type {
		case Ignore:
			continue
		default:
			data, err := ioutil.ReadFile(filepath.Join(rootPath, c.LinkName))
			if err == nil {
				c.Title = Title(string(data))
			}
			if c.Title == "" {
				c.Title = c.LinkName
			}
		}
		n.Children = append(n.Children, c)
	}
}

// Read parse directory recursively, filter unmatched files and empty directories
func (n *Node) Read() {
	if n.Depth > depth {
		return
	}
	files, err := ioutil.ReadDir(filepath.Join(rootPath, n.RelPath))
	if err != nil {
		return
	}
	n.Filter(files)
}

func main() {
	app := cli.NewApp()

	app.Name = "gitbook-gen"
	app.Usage = "generate SUMMARY.md for Gitbook"
	app.Description = "gitbook directory, current directory for default"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "depth, d",
			Value: 2,
			Usage: "depth of TOC",
		},
		cli.StringFlag{
			Name:  "extensions, e",
			Value: ".md,.markdown",
			Usage: "separated by commas, NO spaces",
		},
		cli.StringFlag{
			Name:  "excludes",
			Value: "_book",
			Usage: "exclude directories",
		},
		cli.BoolTFlag{
			Name:  "show-draft",
			Usage: "show draft or NOT",
		},
	}
	app.Action = func(c *cli.Context) error {
		rootPath = "."
		if c.NArg() > 0 {
			rootPath = c.Args().First()
		}
		r := &Node{
			Type:  Directory,
			Depth: 0,
		}

		depth = c.Int("d")
		exts = strings.Split(c.String("e"), ",")
		excludes = strings.Split(c.String("excludes"), ",")
		showDraft = c.Bool("show-draft")
		r.Read()
		//fmt.Println(r.RelPath)
		//r.Tree("")
		r.Summary("")
		return nil
	}
	app.Run(os.Args)
}
