package komozicop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKomoziCop(t *testing.T) {
	testcase := []struct {
		testname       string
		userpoststring string
		usekomozi      bool
	}{
		{
			testname:       "普通ののあの言葉を怒らないテスト",
			userpoststring: "noah_orbergを冠してReoYamadaを名乗るとはええ度胸やな",
			usekomozi:      false,
		},
		{
			testname:       "語尾に小文字をつける時に怒る",
			userpoststring: "のあだょ",
			usekomozi:      true,
		},
	}

	for _, tc := range testcase {
		t.Run(tc.testname, func(t *testing.T) {
			iskomoziuse, err := CheckIlligalKomoziUsing(tc.userpoststring)

			assert.NoError(t, err)
			assert.Equal(t, tc.usekomozi, iskomoziuse)
		})
	}
}
