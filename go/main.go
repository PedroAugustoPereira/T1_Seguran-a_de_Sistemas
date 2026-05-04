package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

const (
	DECRYPT_FILE  string = "texto_decifrado_go.txt"
	ENCRYPT_FILE  string = "texto_criptografado_go.txt"
	ATTACKED_FILE string = "texto_atacado_go.txt"
	ALPHABET_SIZE int    = 26
	ASCII_LOWER_A int    = 97
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso correto:")
		fmt.Println("  Para criptografar: go run main.go dictionary.go crypt <arquivo.txt> <senha>")
		fmt.Println("  Para atacar/quebrar: go run main.go dictionary.go attack <arquivo_criptografado.txt>")
		return
	}

	command := os.Args[1]
	inputFile := os.Args[2]

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Erro ao ler o arquivo %s: %v\n", inputFile, err)
		return
	}
	message := string(content)

	if command == "crypt" {
		if len(os.Args) < 4 {
			fmt.Println("Erro: A senha precisa ser fornecida.")
			return
		}
		chave := os.Args[3]

		// 1. Higienização do texto
		message = sanitizeText(message)

		// 2. Criptografia
		cyptedMessage := crypt(message, chave)

		// 3. Salvar o resultado em arquivo
		err := os.WriteFile(ENCRYPT_FILE, []byte(cyptedMessage), 0644)
		if err != nil {
			fmt.Println("Erro ao salvar arquivo:", err)
			return
		}
		fmt.Println("[+] Criptografia concluída!")
		fmt.Println("[+] Arquivo salvo com sucesso:", ENCRYPT_FILE)

	} else if command == "attack" {
		// Remove espaços/quebras de linha que possam ter vindo do arquivo lido
		message = strings.TrimSpace(message)
		message = strings.ReplaceAll(message, "\n", "")
		message = strings.ReplaceAll(message, "\r", "")

		fmt.Println("[*] Iniciando Criptoanálise do arquivo:", inputFile)
		breakingIC(message)
	} else if command == "decrypt" {
		if len(os.Args) < 4 {
			fmt.Println("Erro: A senha precisa ser fornecida.")
			return
		}
		chave := os.Args[3]

		// Remove quebras de linha para garantir a leitura correta do cifrotexto
		message = strings.TrimSpace(message)
		message = strings.ReplaceAll(message, "\n", "")
		message = strings.ReplaceAll(message, "\r", "")

		// Executa a decifragem
		decryptedMessage := decrypt(message, chave)

		// Salva o resultado
		err := os.WriteFile(DECRYPT_FILE, []byte(decryptedMessage), 0644)
		if err != nil {
			fmt.Println("Erro ao salvar arquivo:", err)
			return
		}
		fmt.Println("[+] Descriptografia concluída!")
		fmt.Println("[+] Arquivo salvo com sucesso:", DECRYPT_FILE)
	} else {
		fmt.Println("Comando inválido. Use 'crypt', 'decrypt' ou 'attack'.")
	}
}

func crypt(message string, key string) string {
	message = strings.ToLower(message)
	message = strings.Trim(message, " ")

	lenKey := len(key)
	nIndexKey := 0
	result := ""

	for i := 0; i < len(message); i++ {
		nAsciimessage := int(message[i]) - ASCII_LOWER_A
		nAsciiKey := int(key[nIndexKey]) - ASCII_LOWER_A

		nAscii := (nAsciimessage + nAsciiKey) % ALPHABET_SIZE

		nAscii = nAscii + ASCII_LOWER_A

		result = result + string(rune(nAscii))
		nIndexKey++

		if nIndexKey == lenKey {
			nIndexKey = 0
		}
	}

	return result
}

func decrypt(message string, key string) string {
	message = strings.ToLower(message)
	message = strings.Trim(message, " ")

	lenMesage := len(message)
	lenKey := len(key)
	result := ""

	nIndexKey := 0

	for i := 0; i < lenMesage; i++ {
		nAsciiMessage := int(message[i]) - ASCII_LOWER_A
		nAsciiKey := int(key[nIndexKey]) - ASCII_LOWER_A

		nAsciiOrigin := (nAsciiMessage - nAsciiKey + ALPHABET_SIZE) % ALPHABET_SIZE
		nAsciiOrigin = nAsciiOrigin + ASCII_LOWER_A
		result = result + string(rune(nAsciiOrigin))

		nIndexKey++

		if nIndexKey == lenKey {
			nIndexKey = 0
		}
	}

	return result
}

func breakingIC(message string) {
	message = strings.ToLower(message)
	message = strings.ReplaceAll(message, " ", "")
	message = strings.ReplaceAll(message, "\n", "")
	message = strings.ReplaceAll(message, "\r", "")

	// Step 1: Find the key length
	keyLength := GetKeyLength(message)
	fmt.Printf("\n[!] Discovered Key Length: %d\n", keyLength)

	// Step 2: Break the cipher using Frequency Analysis
	password := DiscoverKey(message, keyLength)
	fmt.Printf("[!] Discovered Password: %s\n", password)

	// Step 3: Decrypt the text to prove it works
	originalText := decrypt(message, password)

	err := os.WriteFile(ATTACKED_FILE, []byte(originalText), 0644)
	if err != nil {
		fmt.Println("[!] Erro ao salvar o texto decifrado:", err)
		return
	}
	fmt.Println("[+] Ataque concluído com sucesso!")
	fmt.Println("[+] O texto atacado foi salvo no arquivo:", ATTACKED_FILE)
}

// DiscoverKey extracts the exact password using Frequency Analysis
func DiscoverKey(message string, keyLength int) string {
	var wg sync.WaitGroup

	keyFound := make([]byte, keyLength)

	for col := 0; col < keyLength; col++ {
		wg.Add(1)

		go func(currentColumn int) {
			defer wg.Done()

			var columnText []byte
			for i := 0; i < len(message); i++ {
				if i%keyLength == currentColumn {
					columnText = append(columnText, message[i])
				}
			}

			bestShift := 0
			maxScore := 0.0

			for shift := 0; shift < ALPHABET_SIZE; shift++ {
				decryptedFreqs := make([]int, ALPHABET_SIZE)

				for _, char := range columnText {
					decryptedChar := (int(char-'a') - shift + ALPHABET_SIZE) % ALPHABET_SIZE
					decryptedFreqs[decryptedChar]++
				}

				score := 0.0
				for charIndex := 0; charIndex < ALPHABET_SIZE; charIndex++ {
					score += float64(decryptedFreqs[charIndex]) * portugueseFreqs[charIndex]
				}

				if score > maxScore {
					maxScore = score
					bestShift = shift
				}
			}

			keyFound[currentColumn] = byte(bestShift + 'a')
		}(col)
	}

	wg.Wait()
	return string(keyFound)
}

func GetKeyLength(message string) int {
	var wg sync.WaitGroup
	maxKeyLength := 10

	icResults := make([]float64, maxKeyLength+1)

	for keyLength := 1; keyLength <= maxKeyLength; keyLength++ {
		wg.Add(1)

		go func(l int) {
			defer wg.Done()

			//criamos um array de colunas
			columnBytes := make([][]byte, l)

			//populanmos o array de colunas com base em l
			for i := 0; i < len(message); i++ {
				targetColumn := i % l
				columnBytes[targetColumn] = append(columnBytes[targetColumn], message[i])
			}

			var sumIC float64 = 0

			//calculamos a ic de cada coluna, passamos um cast do arry inteiro de caracteres da coluna
			for column := 0; column < l; column++ {
				text := string(columnBytes[column])
				sumIC += CalculateIC(text)
			}

			//adicionamos a média de todas as colunas no array
			icResults[l] = sumIC / float64(l)
		}(keyLength)
	}

	wg.Wait()

	fmt.Println("=== INDEX OF COINCIDENCE RESULTS ===")
	bestKeyLength := 1
	maxIC := 0.0

	for i := 1; i <= maxKeyLength; i++ {
		fmt.Printf("Length %d \t-> Average IC: %.5f\n", i, icResults[i])
		// Adicionando uma margem para sempre preferir o período menor!
		// Como a chave se repete, "pedropedro" dá o mesmo IC que "pedro".
		// O '+ 0.002' garante que não vamos trocar para um tamanho maior a menos que seja genuinamente mais provável.
		if icResults[i] > maxIC+0.002 {
			maxIC = icResults[i]
			bestKeyLength = i
		}
	}

	fmt.Printf("\n[+] The most probable key length is: %d\n", bestKeyLength)
	return bestKeyLength
}

func CalculateIC(text string) float64 {
	totalChars := len(text)

	if totalChars <= 1 {
		return 0.0
	}

	frequencies := make([]int, ALPHABET_SIZE)

	for i := 0; i < totalChars; i++ {
		char := text[i]
		frequencies[char-'a']++
	}

	sum := 0.0
	for _, f := range frequencies {
		sum += float64(f * (f - 1))
	}

	ic := sum / float64(totalChars*(totalChars-1))
	return ic
}

func sanitizeText(text string) string {
	text = strings.ToLower(text)

	// Removendo conforme o enunciado
	replacer := strings.NewReplacer(
		"á", "a", "à", "a", "ã", "a", "â", "a", "ä", "a",
		"é", "e", "è", "e", "ê", "e", "ë", "e",
		"í", "i", "ì", "i", "î", "i", "ï", "i",
		"ó", "o", "ò", "o", "õ", "o", "ô", "o", "ö", "o",
		"ú", "u", "ù", "u", "û", "u", "ü", "u",
		"ç", "c", "ñ", "n",
	)
	text = replacer.Replace(text)

	var clean strings.Builder
	for _, char := range text {
		if char >= 'a' && char <= 'z' {
			clean.WriteRune(char)
		}
	}
	return clean.String()
}
