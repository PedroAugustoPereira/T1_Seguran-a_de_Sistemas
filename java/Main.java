import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;

public class Main {
    // Frequências do Português
    static final double[] portugueseFreqs = {
        0.1463, // a
        0.0104, // b
        0.0388, // c
        0.0499, // d
        0.1257, // e
        0.0102, // f
        0.0130, // g
        0.0128, // h
        0.0618, // i
        0.0040, // j
        0.0002, // k
        0.0278, // l
        0.0474, // m
        0.0505, // n
        0.1073, // o
        0.0252, // p
        0.0120, // q
        0.0653, // r
        0.0781, // s
        0.0434, // t
        0.0463, // u
        0.0167, // v
        0.0001, // w
        0.0021, // x
        0.0001, // y
        0.0047  // z
    };

    public static void main(String[] args) {
        String command = "";
        String inputFile = "";
        String chave = "";

        if (args.length >= 2) {
            command = args[0];
            inputFile = args[1];
            if (args.length >= 3) {
                chave = args[2];
            }
        } else {
            // Modo interativo para rodar direto pelo "Play" do VS Code
            java.util.Scanner scanner = new java.util.Scanner(System.in);
            System.out.println("=== SISTEMA DE CRIPTOGRAFIA VIGENÈRE ===");
            System.out.print("Digite o comando (crypt, decrypt, attack): ");
            command = scanner.nextLine().trim();
            
            System.out.print("Digite o caminho do arquivo (se estiver no VS Code, digite apenas: texto.txt): ");
            inputFile = scanner.nextLine().trim();

            if (command.equals("crypt") || command.equals("decrypt")) {
                System.out.print("Digite a senha: ");
                chave = scanner.nextLine().trim();
            }
            scanner.close();
        }

        String message = "";

        try {
            message = new String(Files.readAllBytes(Paths.get(inputFile)));
        } catch (IOException e) {
            System.out.println("Erro ao ler o arquivo " + inputFile + ": " + e.getMessage());
            return;
        }

        if (command.equals("crypt")) {
            if (chave.isEmpty()) {
                System.out.println("Erro: A senha precisa ser fornecida.");
                return;
            }

            // 1. Higienização do texto
            message = sanitizeText(message);

            // 2. Criptografia
            String cyptedMessage = crypt(message, chave);

            // 3. Salvar o resultado
            try {
                Files.write(Paths.get("texto_criptografado_java.txt"), cyptedMessage.getBytes());
                System.out.println("[+] Criptografia concluída!");
                System.out.println("[+] Arquivo salvo com sucesso: texto_criptografado_java.txt");
            } catch (IOException e) {
                System.out.println("Erro ao salvar arquivo: " + e.getMessage());
            }

        } else if (command.equals("decrypt")) {
            if (chave.isEmpty()) {
                System.out.println("Erro: A senha precisa ser fornecida.");
                return;
            }

            message = message.trim().replace("\n", "").replace("\r", "");
            String decryptedMessage = decrypt(message, chave);

            try {
                Files.write(Paths.get("texto_decifrado_java.txt"), decryptedMessage.getBytes());
                System.out.println("[+] Descriptografia concluída!");
                System.out.println("[+] Arquivo salvo com sucesso: texto_decifrado_java.txt");
            } catch (IOException e) {
                System.out.println("Erro ao salvar arquivo: " + e.getMessage());
            }

        } else if (command.equals("attack")) {
            message = message.trim().replace("\n", "").replace("\r", "");
            System.out.println("[*] Iniciando Criptoanálise do arquivo: " + inputFile);
            breakingIC(message);
        } else {
            System.out.println("Comando inválido. Use 'crypt', 'decrypt' ou 'attack'.");
        }
    }

    public static String sanitizeText(String text) {
        text = text.toLowerCase();
        
        // Substituindo caracteres acentuados
        text = text.replace("á", "a").replace("à", "a").replace("ã", "a").replace("â", "a");
        text = text.replace("é", "e").replace("ê", "e");
        text = text.replace("í", "i");
        text = text.replace("ó", "o").replace("õ", "o").replace("ô", "o");
        text = text.replace("ú", "u");
        text = text.replace("ç", "c");

        // Mantendo apenas letras a-z
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < text.length(); i++) {
            char c = text.charAt(i);
            if (c >= 'a' && c <= 'z') {
                sb.append(c);
            }
        }
        return sb.toString();
    }

    public static String crypt(String message, String key) {
        message = sanitizeText(message);
        key = sanitizeText(key);

        int lenMessage = message.length();
        int lenKey = key.length();
        StringBuilder result = new StringBuilder();

        int nIndexKey = 0;
        for (int i = 0; i < lenMessage; i++) {
            int nAsciiMessage = message.charAt(i) - 97;
            int nAsciiKey = key.charAt(nIndexKey) - 97;

            int nAscii = (nAsciiMessage + nAsciiKey) % 26;
            nAscii = nAscii + 97;

            result.append((char) nAscii);
            nIndexKey++;

            if (nIndexKey == lenKey) {
                nIndexKey = 0;
            }
        }
        return result.toString();
    }

    public static String decrypt(String message, String key) {
        message = sanitizeText(message); // Remove possible spaces
        key = sanitizeText(key);

        int lenMessage = message.length();
        int lenKey = key.length();
        StringBuilder result = new StringBuilder();

        int nIndexKey = 0;
        for (int i = 0; i < lenMessage; i++) {
            int nAsciiMessage = message.charAt(i) - 97;
            int nAsciiKey = key.charAt(nIndexKey) - 97;

            int nAsciiOrigin = (nAsciiMessage - nAsciiKey + 26) % 26;
            nAsciiOrigin = nAsciiOrigin + 97;
            result.append((char) nAsciiOrigin);

            nIndexKey++;

            if (nIndexKey == lenKey) {
                nIndexKey = 0;
            }
        }
        return result.toString();
    }

    public static void breakingIC(String message) {
        int keyLength = GetKeyLength(message);
        System.out.println("\n[!] Discovered Key Length: " + keyLength);

        String password = DiscoverKey(message, keyLength);
        System.out.println("[!] Discovered Password: " + password);

        String originalText = decrypt(message, password);
        
        try {
            Files.write(Paths.get("texto_atacado_java.txt"), originalText.getBytes());
            System.out.println("[+] Ataque concluído com sucesso!");
            System.out.println("[+] O texto atacado foi salvo no arquivo: texto_atacado_java.txt");
        } catch (IOException e) {
            System.out.println("[!] Erro ao salvar o texto atacado: " + e.getMessage());
        }
    }

    public static String DiscoverKey(String message, int keyLength) {
        char[] keyFound = new char[keyLength];

        for (int col = 0; col < keyLength; col++) {
            StringBuilder columnText = new StringBuilder();
            for (int i = 0; i < message.length(); i++) {
                if (i % keyLength == col) {
                    columnText.append(message.charAt(i));
                }
            }

            int bestShift = 0;
            double maxScore = 0.0;

            for (int shift = 0; shift < 26; shift++) {
                int[] decryptedFreqs = new int[26];

                for (int i = 0; i < columnText.length(); i++) {
                    char c = columnText.charAt(i);
                    int decryptedChar = (c - 'a' - shift + 26) % 26;
                    decryptedFreqs[decryptedChar]++;
                }

                double score = 0.0;
                for (int charIndex = 0; charIndex < 26; charIndex++) {
                    score += decryptedFreqs[charIndex] * portugueseFreqs[charIndex];
                }

                if (score > maxScore) {
                    maxScore = score;
                    bestShift = shift;
                }
            }

            keyFound[col] = (char) (bestShift + 'a');
        }

        return new String(keyFound);
    }

    public static double CalculateIC(String text) {
        int n = text.length();
        if (n <= 1) return 0.0;

        int[] freqs = new int[26];
        for (int i = 0; i < n; i++) {
            char c = text.charAt(i);
            if (c >= 'a' && c <= 'z') {
                freqs[c - 'a']++;
            }
        }

        double ic = 0.0;
        for (int f : freqs) {
            ic += f * (f - 1);
        }

        ic = ic / ((double) n * (n - 1));
        return ic;
    }

    public static int GetKeyLength(String message) {
        int maxKeyLength = 20;
        double[] icResults = new double[maxKeyLength + 1];

        for (int l = 1; l <= maxKeyLength; l++) {
            StringBuilder[] columnBytes = new StringBuilder[l];
            for (int i = 0; i < l; i++) {
                columnBytes[i] = new StringBuilder();
            }

            for (int i = 0; i < message.length(); i++) {
                int targetColumn = i % l;
                columnBytes[targetColumn].append(message.charAt(i));
            }

            double sumIC = 0.0;
            for (int column = 0; column < l; column++) {
                sumIC += CalculateIC(columnBytes[column].toString());
            }

            icResults[l] = sumIC / l;
        }

        System.out.println("=== INDEX OF COINCIDENCE RESULTS ===");
        int bestKeyLength = 1;
        double maxIC = 0.0;

        for (int i = 1; i <= maxKeyLength; i++) {
            System.out.printf("Length %d \t-> Average IC: %.5f\n", i, icResults[i]);
            if (icResults[i] > maxIC + 0.002) {
                maxIC = icResults[i];
                bestKeyLength = i;
            }
        }

        System.out.println("\n[+] The most probable key length is: " + bestKeyLength);
        return bestKeyLength;
    }
}
