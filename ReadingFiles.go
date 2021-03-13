package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var arquivos []string

	raiz := "."
	err := filepath.Walk(raiz, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".go" { //Filtrar por extensão
			arquivos = append(arquivos, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	for _, arquivo := range arquivos {
		fmt.Println(arquivo)
	}
}

func MoverArquivos(origem string, destino string) {
	err := os.Rename(origem, destino)
	if err != nil {
		log.Fatal(err)
	}
}
