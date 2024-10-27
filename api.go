package main

import (
	"fmt"
	"log"
)

func main() {
	// Configuração da API do LinkedIn
	client := linkedin.NewClient("seu-client-id", "seu-client-secret")
	// Configurações adicionais

	// Compartilhamento de conteúdo
	shareContent(client, "Este é um exemplo de postagem no LinkedIn!")
}

func shareContent(client *linkedin.Client, content string) {
	post := linkedin.NewPost(content)
	err := client.SharePost(post)
	if err != nil {
		log.Fatalf("Erro ao compartilhar conteúdo: %v", err)
	}
	fmt.Println("Conteúdo compartilhado com sucesso!")
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
