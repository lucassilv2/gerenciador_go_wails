# 🧠 Gerenciador de Tarefas com Wails + Svelte

Um gerenciador de tarefas simples, leve e multiplataforma, construído com [Wails](https://wails.io/) no backend (Go) e [Svelte](https://svelte.dev/) no frontend. Permite listar, pausar, retomar e encerrar processos do sistema com uma interface semelhante ao terminal.

---

## ✨ Funcionalidades

- 🔍 Lista todos os processos do sistema (PID, usuário, uso de CPU, tempo, comando e estado)
- ⏸️ Pausar um processo
- ▶️ Retomar um processo pausado
- 🗑️ Finalizar um processo
- 🔁 Atualização automática a cada 1s

---

## 🛠️ Tecnologias Utilizadas

- [Go (Golang)](https://golang.org/)
- [Wails](https://wails.io/)
- [Svelte](https://svelte.dev/)

---

## 🚀 Como executar

### Pré-requisitos

- [Go](https://golang.org/doc/install) >= 1.20
- [Node.js](https://nodejs.org/) >= 16
- [Wails CLI](https://wails.io/docs/gettingstarted/installation)

### Clonando o projeto

```bash
git clone https://github.com/lucassilv2/gerenciador_go_wails.git
cd gerenciador-go-wails
wails dev