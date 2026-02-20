package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv carga las variables de entorno desde un archivo .env en caso de que no existan a nivel de ejecución.
// Es utilizada de forma transversal por los diferentes adapters y el punto de entrada de la aplicación.
func LoadEnv(envToLoad []string) {
	for _, env := range envToLoad {
		if os.Getenv(env) == "" {
			err := godotenv.Load()
			if err != nil {
				log.Fatal("Error loading .env file")
			}
			log.Println("Environment variables loaded from .env file")
			return
		}
	}
}
