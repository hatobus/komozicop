package komozicop

import (
	"github.com/ikawaha/kagome/tokenizer"
)

var t = tokenizer.New()

func CheckIlligalKomoziUsing(userpost string) (bool, error) {
	tokens := t.Tokenize(userpost)
	lastwordinfo := tokens[len(tokens)-2]

	nountype := lastwordinfo.Features()[0]

	if nountype == "名詞" {
		lastwordfeaturelength := len(lastwordinfo.Features())
		inflexions := lastwordinfo.Features()[lastwordfeaturelength-4 : lastwordfeaturelength-1]
		komoziflag := true
		for _, c := range inflexions {
			if c == "*" {
				continue
			} else {
				komoziflag = false
				break
			}
		}
		return komoziflag, nil
	}

	return false, nil
}
