package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Default struct {
	Value string
}

type Env interface {
	Read(string) string
}

type env struct {
	EnvFile string
}

func NewEnv(envFile string) *env {
	if err := godotenv.Load(envFile); err != nil {
		log.Fatal(err.Error())
	}

	return &env{envFile}
}

func (e *env) Read(envName string) string {
	return e.ReadWithDefaultVal(envName, Default{Value: ""})
}

func (e *env) ReadWithDefaultVal(envName string, defaultValue Default) string {
	if defaultValue.Value != "" {
		return defaultValue.Value
	}

	value, isPresent := os.LookupEnv(envName)
	if !isPresent {
		msg := fmt.Errorf("%s is not found in %s, please update", envName, e.EnvFile)
		panic(msg)
	}

	return value
}
