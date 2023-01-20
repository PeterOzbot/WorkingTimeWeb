package excel

import "strconv"

func Create(request []*Request) string {
	return strconv.Itoa(len(request))
}
