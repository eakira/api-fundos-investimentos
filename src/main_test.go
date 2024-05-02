package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	t.Run("when_not_dot_env_return_error", func(t *testing.T) {
		err := os.Rename(".env", "Renomeando.env")
		assert.Nil(t, err)

		defer func() {
			err := os.Rename("Renomeando.env", ".env")
			if err != nil {
				t.Fatalf("Erro ao restaurar o nome do arquivo original: %s", err)
			}
		}()

		err = LoadDotEnv()
		assert.NotNil(t, err)
	})

	t.Run("when_dot_env_return_sucess", func(t *testing.T) {

		err := LoadDotEnv()
		assert.Nil(t, err)
	})

}
