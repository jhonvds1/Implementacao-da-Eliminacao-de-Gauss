Nome: Jonatha Viegas da Silva
Matrícula: 22103000

Instruções para Uso dos Programas
Pré-requisitos
Para compilar e executar os programas em C, Go e Rust, você precisa ter os seguintes programas instalados em sua máquina:
C: GCC (GNU Compiler Collection) para compilar o código.
Go: Go runtime e compilador.
Rust: Cargo (gerenciador de pacotes e compilador).

Para compilar o código em C, você deve usar o seguinte comando no terminal: gcc gauss.c -o gauss
Este comando irá gerar um executável chamado gauss.
Depois de compilar o programa, execute-o com o seguinte comando: ./gauss <tamanho_da_matriz> [semente_opcional]
<tamanho_da_matriz>: O tamanho da matriz (valor inteiro N).
[semente_opcional]: Uma semente para o gerador de números aleatórios (opcional).

Para rodar o programa em Go, basta usar o seguinte comando (não é necessário compilar previamente): go run gauss.go <tamanho_da_matriz> [semente_opcional]
<tamanho_da_matriz>: O tamanho da matriz (valor inteiro N).
[semente_opcional]: Uma semente para o gerador de números aleatórios (opcional).

Para compilar o programa em Rust, use o comando: cargo build --release
Após a compilação, o executável estará localizado dentro da pasta target/release. Para executar o programa, use:./target/release/gauss <tamanho_da_matriz> [semente_opcional]
<tamanho_da_matriz>: O tamanho da matriz (valor inteiro N).
[semente_opcional]: Uma semente para o gerador de números aleatórios (opcional).
