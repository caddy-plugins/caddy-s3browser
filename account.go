package s3browser

import "strings"

type Account struct {
	Key      string
	Bucket   string
	Secret   string
	Region   string
	CDNURL   string
	Endpoint string
	Secure   bool
}

func (a *Account) setDefaults() {
	if len(a.CDNURL) > 0 {
		a.CDNURL = strings.TrimSuffix(a.CDNURL, `/`)
	}
}

var AccountGetter func(arg string) (Account, error)
