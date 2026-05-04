# T1 - Cifra de Vigenere

Projeto de criptografia e criptoanalise com a Cifra de Vigenere. O objetivo e criptografar textos com uma senha e recuperar o texto original de um cifrotexto sem conhecer essa senha, assumindo idioma portugues.

## Pontos principais

- Ler um arquivo `.txt` e uma senha.
- Higienizar o texto: minusculas, sem acentos, sem espacos, sem numeros, sem pontuacao e apenas `a-z`.
- Criptografar usando Vigenere com a senha repetida ciclicamente.
- Quebrar o texto cifrado sem a senha.
- Estimar o tamanho da chave com Indice de Coincidencia.
- Descobrir a chave por analise de frequencia das letras em portugues.
- Gerar arquivos com o texto criptografado, decifrado e atacado.

## Implementacao

O projeto tem duas implementacoes: Java e Go.

Conceitualmente, a criptografia converte letras para valores de `0` a `25`, soma o deslocamento da letra correspondente da senha e aplica modulo `26`. A descriptografia faz o inverso, subtraindo o deslocamento.

O ataque divide o texto cifrado em colunas conforme possiveis tamanhos de chave, calcula o Indice de Coincidencia para escolher o tamanho mais provavel e depois testa deslocamentos em cada coluna comparando a frequencia resultante com a distribuicao esperada do portugues.

## Como rodar

A senha e os nomes dos arquivos ficam no topo do `Makefile`.

Go:

```bash
make go-crypt
make go-decrypt
make go-attack
```

Java:

```bash
make java-crypt
make java-decrypt
make java-attack
```

Arquivos principais:

- `texto.txt`: entrada original
- `go/texto_criptografado_go.txt`: saida cifrada em Go
- `go/texto_decifrado_go.txt`: saida decifrada em Go
- `go/texto_atacado_go.txt`: saida do ataque em Go
- `texto_criptografado_java.txt`: saida cifrada em Java
- `texto_decifrado_java.txt`: saida decifrada em Java
- `texto_atacado_java.txt`: saida do ataque em Java
