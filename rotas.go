package main

// Define os endpoints e os métodos. Passar o NOME da função que irá retornar a resposta
func definerotas() {
	rotas.GET("/albuns", getAlbuns)        // método GET endpoint /albuns retorna JSON de todos os álbuns
	rotas.POST("/novoalbum", postAlbuns)   // métodos POST endpoint /novoalbum recebe um JSON e inclui um álbum
	rotas.GET("/album/:id", getAlbumPorID) // método GET end point /album recebe um id no parÂmetro e retorna o respectivo álbum
}
