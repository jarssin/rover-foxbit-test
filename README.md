# Rover Foxbit Test

Este projeto simula o controle de robôs exploradores (rovers) em um plateau.

## Como funciona

- O plateau é definido por um grid de coordenadas.
- Cada rover recebe uma posição inicial (X, Y, Direção) e uma sequência de comandos.
- Os comandos podem ser:
  - `L`: girar à esquerda
  - `R`: girar à direita
  - `M`: mover para frente
- O sistema valida limites do plateau e comandos inválidos.

## Exemplo de entrada

```
5 5
1 2 N
LMLMLMLMM
3 3 E
MMRMMRMRRM
```

## Exemplo de saída esperada

```
1 3 N
5 1 E
```

## Como rodar

2. Clone o repositório e acesse a pasta do projeto.
3. Execute:

```bash
go run ./cmd/main.go < stdin
```

Ou compile e rode:

```bash
go build -o roverfoxbit ./cmd/main.go
./roverfoxbit < stdin
```

## Estrutura do projeto

- `cmd/`: ponto de entrada da aplicação
- `internal/utils/`: utilitários de leitura de entrada
- `pkg/rover/`: lógica do rover
- `pkg/plateau/`: lógica do plateau

Desenvolvido por Jarssin para o desafio Foxbit.
