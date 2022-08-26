package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	exibeIntroducao()
	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciaMonitoramento()
		case 2:
			imprimeLogs()
		case 3:
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conhecço esse comando!")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	nome := "Silas"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão ", versao)
}

func exibeMenu() {
	fmt.Println("Escolha um comando: ")
	fmt.Println(`1- Moritorar`)
	fmt.Println(`2- Exibir logs`)
	fmt.Println(`3- Sair do programa`)
}

func leComando() int {

	var comandoLido int

	fmt.Scan(&comandoLido) //comando "scan" faz a infernêcia do tipo automaticamente
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido
}

func iniciaMonitoramento() {
	fmt.Println("Monitorando...")
	// sites := []string{
	// 	"https://random-status-code.herokuapp.com/",
	// 	"https://alura.com.br",
	// 	"https://gorila.com.br"}

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {

			fmt.Println("testando site", site)
			testaSite(site)

		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")

}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt") // Retorna um endereço de memória do arquivo
	//arquivo, err := ioutil.ReadFile("sites.txt") //Retorna um array de bytes

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) // cria arquivo com permissoes

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	// existe uma convenção de formatação de datas em go, para saber mais elr documentação
	arquivo.WriteString(time.Now().Format("02/01/2006 - 15:04:05") + " - " + site + " - online " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	fmt.Println("Exibindo Logs...")
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	fmt.Println(string(arquivo))
}
