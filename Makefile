PASSWORD := senhaforte
TEXT_FILE := texto.txt
GO_ENCRYPTED_FILE := texto_criptografado_go.txt
JAVA_ENCRYPTED_FILE := texto_criptografado_java.txt

.PHONY: go-crypt go-decrypt go-attack java-crypt java-decrypt java-attack

go-crypt:
	@echo "Crypting $(TEXT_FILE)..."
	@cd go && go run . crypt ../$(TEXT_FILE) $(PASSWORD)

go-decrypt:
	@echo "Decrypting $(GO_ENCRYPTED_FILE)..."
	@cd go && go run . decrypt $(GO_ENCRYPTED_FILE) $(PASSWORD)

go-attack:
	@echo "Attacking $(GO_ENCRYPTED_FILE)..."
	@cd go && go run . attack $(GO_ENCRYPTED_FILE)

java-crypt:
	@echo "Crypting $(TEXT_FILE)..."
	@javac java/Main.java
	@java -cp java Main crypt $(TEXT_FILE) $(PASSWORD)

java-decrypt:
	@echo "Decrypting $(JAVA_ENCRYPTED_FILE)..."
	@javac java/Main.java
	@java -cp java Main decrypt $(JAVA_ENCRYPTED_FILE) $(PASSWORD)

java-attack:
	@echo "Attacking $(JAVA_ENCRYPTED_FILE)..."
	@javac java/Main.java
	@java -cp java Main attack $(JAVA_ENCRYPTED_FILE)
