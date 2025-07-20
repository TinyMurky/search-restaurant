package config

import (
	"bytes"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/TinyMurky/search-restaurant/internal/pkg/customerrors"
	"github.com/TinyMurky/search-restaurant/internal/pkg/testhelpers"
)

func TestLoadConfigFromReader(t *testing.T) {
	t.Run("should read correct yaml", func(t *testing.T) {
		want := Config{
			Server: ServerConfig{
				Port: 20777,
			},
		}

		out, err := yaml.Marshal(&want)

		testhelpers.AssertNoError(t, err)

		r := bytes.NewBuffer(out)

		cfg, err := loadConfigByReader(r)

		testhelpers.AssertNoError(t, err)

		got := *cfg

		testhelpers.AssertDeepEqual(t, got, want)
	})

	t.Run("should throw IOError if yaml is incorrect", func(t *testing.T) {
		wrongCfg := `
server:
    wrong: 123
		`

		r := strings.NewReader(wrongCfg)

		_, err := loadConfigByReader(r)

		var ioErr *customerrors.IOError
		testhelpers.AssertAsError(t, err, &ioErr)
	})
}
