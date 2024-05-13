package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	dbgravacoes "github.com/SergioPedron/estudosGo_dbgravacoes"
)

//--------------------------------------------------------------------------------------------------------------------------------------------//

// Retorna JSON de albuns como resposta ao método GET do endpoint /albuns
func getAlbuns(c *gin.Context) {
	albuns, err := dbgravacoes.AlbunsTodos(db)
	if err != nil {
		fmt.Println("erro %w", err) // Rever como o gin pode tratar/exibir o erro
	}
	c.IndentedJSON(http.StatusOK, albuns) // Utilizar c.JSON para gerar compactado
}

//--------------------------------------------------------------------------------------------------------------------------------------------//

// Adiciona um album com base no JSON recebido  pelo método POST no endpoint /postAlbuns
func postAlbuns(c *gin.Context) {
	var novoAlbum dbgravacoes.Album
	// Chama BindJSON para vincular o JSON recebido na struct novoAlbum
	if err := c.BindJSON(&novoAlbum); err != nil {
		return
	}
	// Adiciona o album na tabela
	id, err := dbgravacoes.AdicionaAlbum(db, novoAlbum)
	if err != nil {
		fmt.Println("erro %w", err) // Rever como o gin pode tratar/exibir o erro
	}
	fmt.Printf("ID do novo Album cadastrado: %v\n", id)
}

//--------------------------------------------------------------------------------------------------------------------------------------------//

// Retorna um album conforme Id recebido por parâmetro
// Parâmetro enviado pelo cliente e, em seguida, retorna esse álbum como resposta.
func getAlbumPorID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 0, 64)
	if err != nil {
		//fmt.Println("erro ao converter string parâmetro para inteiro %w", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"erro": "erro ao converter string parâmetro para inteiro. erro: " + err.Error()}) // retorna JSON com a mensagem
		return
	}
	album, encontrou, err := dbgravacoes.AlbumPorID(db, id)
	if err == nil {
		if !encontrou {
			c.IndentedJSON(http.StatusNotFound, gin.H{"mensagem": "Album não encontrado"}) // retorna JSON com a mensagem
		} else {
			c.IndentedJSON(http.StatusOK, album) // Utilizar c.JSON para gerar compactado
		}
	} else {
		fmt.Println(err)                                                // Rever como o gin pode tratar/exibir o erro
		c.IndentedJSON(http.StatusNotFound, gin.H{"erro": err.Error()}) // retorna JSON com a mensagem
	}

}

//--------------------------------------------------------------------------------------------------------------------------------------------//
