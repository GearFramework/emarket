package auth

import (
	"flag"
	"fmt"
	"github.com/imdario/mergo"
	"github.com/joho/godotenv"
	"log"
)

// Flags структура консольных флагов
type Flags struct {
	Addr string
}

func NewEnv(envFile string) error {
	return GetEnv(envFile)
}

// Получаем рабочее окружение
func GetFlags(defaultFlags *Flags) *Flags {
	_ = mergo.Merge(defaultFlags, GetDefaultFlags())
	conf := &Flags{
		Addr: *flag.String("a", defaultFlags.Addr, fmt.Sprintf("Адрес %s", defaultFlags.Addr)),
	}
	flag.Parse()
	return conf
}

func GetDefaultFlags() *Flags {
	return &Flags{
		Addr: ":8080",
	}
}

func GetEnv(envFile string) error {
	if err := godotenv.Overload(envFile); err != nil {
		log.Printf("Error env file %s, read from env %s\n", envFile, err.Error())
		return err
	}
	return nil
}
