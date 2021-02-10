package enderecos

import "strings"

// TipoDeEndereco verifica se o tipo do endereco é valido ou não
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia", "travessa"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo Inválido"
}
