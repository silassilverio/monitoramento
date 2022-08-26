# Monitoramento de Sites

## Descrição
Este é um projeto feito no curso da Alura sobre a linguagem golang. Basicamente é um monitorador de sites, que faz uma requisição http e checa se o site está disponivel ou não, logando a data/hora da requisição, o site e o status code.

## Como executar

```bash
# Clone este repositório
$ git clone <https://github.com/silassilverio/monitoramento-Go.git>

# Rode o programa
$ go run main.go

# siga as instruções do menu
1- Moritorar
2- Exibir logs
3- Sair do programa
```
## Informações extra
Existem duas constantes para definir o tempo entre os monitoramentos (delay) e também a quantidade de vezes que a checagem será feita em cada site (monitoramentos).

Após a execução um arquivo de log.txt será criado na raiz do projeto, mas também é possivel visualiza-los atraés do comando 2