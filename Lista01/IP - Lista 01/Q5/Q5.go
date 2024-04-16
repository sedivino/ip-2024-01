package main

import "fmt"

func main() {
    var numeroConta int
    var tipoConsumidor string
    var consumoAgua float64

    // Entrada de dados
    fmt.Println("Digite o número da conta do cliente:")
    fmt.Scanln(&numeroConta)

    fmt.Println("Digite o tipo de consumidor (residencial, comercial ou industrial):")
    fmt.Scanln(&tipoConsumidor)

    fmt.Println("Digite o consumo de água em metros cúbicos:")
    fmt.Scanln(&consumoAgua)

    // Cálculo da conta de água
    var valorConta float64
    switch tipoConsumidor {
    case "residencial":
        valorConta = 5.0 + (0.05 * consumoAgua)
    case "comercial":
        if consumoAgua <= 80 {
            valorConta = 500.0
        } else {
            valorConta = 500.0 + 0.25*(consumoAgua-80)
        }
    case "industrial":
        if consumoAgua <= 100 {
            valorConta = 800.0
        } else {
            valorConta = 800.0 + 0.04*(consumoAgua-100)
        }
    default:
        fmt.Println("Tipo de consumidor inválido.")
        return
    }

    // Saída de dados
    fmt.Printf("Número da conta: %d\nValor a ser pago: R$ %.2f\n", numeroConta, valorConta)
}
