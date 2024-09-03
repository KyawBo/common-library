package secret

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

type TestUnmarsahl struct {
	Secrets Secrets `mapstructure:"secrets"`
}

type Secrets struct {
	TestGetSecret string `mapstructure:"test-get-secret"`
}

func TestGetSecretValue(t *testing.T) {

	var testUnmarsahl TestUnmarsahl

	os.Setenv("SECRET_test-get-secret", "test_get_secret_value")

	GetSecretValue()

	err := viper.Unmarshal(&testUnmarsahl)
	assert.NoError(t, err)

	assert.Equal(t, testUnmarsahl.Secrets.TestGetSecret, "test_get_secret_value")
}
