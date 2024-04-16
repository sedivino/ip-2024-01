package main

import "fmt"

func main() {
    // Preços por categoria
    pp := 1 // preço da categoria populular
    pg := 5 // preço da categoria geral
    pa := 10 // preço da categoria arquibancada
    pc := 20 // preço da categoria cadeira

    var quantidadedejogos int
    fmt.Println("Digite a quantidade de jogos:")
    fmt.Scanln(&quantidadedejogos)

    n := 0
    for n < quantidadedejogos {
        var pt, ct, cg, ca, cc int
        fmt.Println("Digite o público total do jogo de futebol:")
        fmt.Scanln(&pt)

        fmt.Println("Digite a porcentagem de pessoas na categoria popular:")
        fmt.Scanln(&ct)

        fmt.Println("Digite a porcentagem de pessoas na categoria geral:")
        fmt.Scanln(&cg)

        fmt.Println("Digite a porcentagem de pessoas na categoria arquibancada:")
        fmt.Scanln(&ca)

        fmt.Println("Digite a porcentagem de pessoas na categoria cadeiras:")
        fmt.Scanln(&cc)

        // Calculando a quantidade de pessoas em cada categoria
        np := pt * ct / 100
        ng := pt * cg / 100
        na := pt * ca / 100
        nc := pt * cc / 100

        // Calculando a renda de cada categoria
        rp := np * pp
        rg := ng * pg
        ra := na * pa
        rc := nc * pc

        // Calculando a renda total do jogo
        rt := rp + rg + ra + rc

        // Imprimindo a renda do jogo
        fmt.Printf("A renda do jogo é: %d\n", rt)

        n = n + 1
    }
}
