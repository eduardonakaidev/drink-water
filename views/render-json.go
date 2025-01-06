package views

import (
	"encoding/json"
	"net/http"
)
// RenderJSON envia uma resposta JSON ao cliente com o código de status fornecido.
func RenderJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	// Define o cabeçalho como JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Serializa os dados em JSON
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Erro ao gerar resposta JSON", http.StatusInternalServerError)
	}
}