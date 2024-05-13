package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	dbgravacoes "github.com/SergioPedron/estudosGo_dbgravacoes"
)

var db *sql.DB        // manipulador do banco
var rotas *gin.Engine // manipulador das rotas

// init é executado antes de qualquer coisa e aqui será utilizado para inicializar o manipulador do banco e definição das rotas
func init() {
	// Utilizar o Conectadbgravacoes no init ou no início do main pois se der erro o mesmo será tratado com log.fatal e irá interromper o fluxo do programa. Rever tratamento de erros
	db = dbgravacoes.Conectadbgravacoes() //  Inicializa banco
	rotas = gin.Default()                 //  Inicializa roteador gin
}

func main() {
	definerotas()               // define endpoints e métodos
	rotas.Run("localhost:8080") // Inicia o servidor
}
