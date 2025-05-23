package s3browser

import (
	"html/template"
	"path/filepath"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
)

type TmplData struct {
	CDNURL string
	Directory
}

type Directory struct {
	Path    string
	Folders []Folder
	Files   []File
	README  string
}

type Folder struct {
	Name string
}

type File struct {
	Folder string
	Bytes  int64
	Name   string
	Date   time.Time
}

type Config struct {
	account       Account
	Prefix        string
	Refresh       string
	Debug         bool
	CSSCDN        string
	prefixes      []string
	accountGetter func(string) (Account, error)
	accountID     string
}

func (c *Config) GetAccount() (Account, error) {
	account := c.account
	if len(c.accountID) > 0 && c.accountGetter != nil {
		a, err := c.accountGetter(c.accountID)
		if err != nil {
			return account, err
		}
		account = a
	}
	return account, nil
}

type Node struct {
	Link         string
	ReadableName string
}

// HumanSize returns the size of the file as a human-readable string
// in IEC format (i.e. power of 2 or base 1024).
func (f File) HumanSize() string {
	return humanize.IBytes(uint64(f.Bytes))
}

// HumanModTime returns the modified time of the file as a human-readable string.
func (f File) HumanModTime(format string) string {
	return f.Date.Format(format)
}

func (d Directory) ReadableName() string {
	return cleanUp(d.Path)
}

func (d Directory) RenderREADME() template.HTML {
	return template.HTML(d.README)
}

func (f Folder) ReadableName() string {
	return cleanUp(f.Name)
}

func cleanUp(s string) string {
	dir := filepath.Dir(s)
	return filepath.Base(dir)
}

func (d Directory) Breadcrumbs() []Node {
	var nodes []Node

	tempDir := strings.Split(d.Path, "/")
	built := ""
	for _, tempFolder := range tempDir {
		if len(tempFolder) < 1 {
			continue
		}
		if len(built) < 1 {
			built = tempFolder + "/"
		} else {
			built = built[:len(built)-1] + "/" + tempFolder + "/"
		}
		nodes = append(nodes, Node{Link: built, ReadableName: tempFolder})
	}
	return nodes
}
