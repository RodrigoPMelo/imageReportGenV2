# 🖼️ ImageReportGen — Desktop Report Generator (Wails)

## 📘 Visão Geral

O **ImageReportGen** é um aplicativo desktop multiplataforma (foco inicial em **Windows**) criado para gerar **relatórios fotográficos padronizados** de forma rápida, offline e personalizável.

Ele permite que o usuário selecione imagens locais ou compactadas em `.zip`, defina um **template DOCX** (folha A4 simples ou com cabeçalho institucional), e gere automaticamente um relatório formatado — organizando fotos por orientação (paisagem/retrato), ajustando dimensões e aplicando um layout profissional.

O projeto está sendo reconstruído em **[Wails](https://wails.io/)** (Go + JS/TS) para oferecer uma base moderna, escalável e pronta para evoluir com novas funcionalidades, como integrações com Google Drive e exportação em PDF.

---

## ✨ Principais Funcionalidades (Versão 1)

- **Interface desktop moderna** desenvolvida com Wails  
- **Suporte a imagens locais e arquivos .zip**
- **Template DOCX configurável** (folha A4 ou modelo institucional)
- **Layout automático**:
  - **Paisagem:** 3 imagens por página (1 coluna × 3 linhas)  
  - **Retrato:** 4 imagens por página (2 × 2)
- **Campos opcionais de cabeçalho:**
  - Nome do cliente
  - Município / Estado
  - Data de conclusão
- **Execução 100% offline**
- **Redimensionamento automático** e uniformização de imagens
- **Limpeza de arquivos temporários** após a geração

---

## 🧩 Arquitetura

O aplicativo é dividido em camadas independentes:

| Camada | Descrição |
|--------|------------|
| **UI (frontend)** | Construída com JS/TS via Wails — exibe a interface, gerencia drag & drop e configuração de layout |
| **Core (backend Go)** | Lida com geração do relatório, extração de imagens e manipulação de templates |
| **Renderer** | Responsável por exportar o relatório final (inicialmente DOCX, futuramente PDF) |
| **Storage** | Gerencia acesso a arquivos locais e (em versões futuras) integração com Google Drive |

---

## ⚙️ Requisitos

- **Windows 10+** (ou superior)  
- **Go 1.22+**  
- **Node.js 18+**  
- **Wails CLI**  
  ```bash
  go install github.com/wailsapp/wails/v2/cmd/wails@latest
  ```

---

## 🏗️ Como rodar o projeto

```bash
# Clonar o repositório
git clone https://github.com/SEU_USUARIO/ImageReportGen.git
cd ImageReportGen

# Instalar dependências
npm install
go mod tidy

# Rodar em modo desenvolvimento
wails dev

# Gerar build (.exe)
wails build
```

O executável final será gerado em:
```
build/bin/ImageReportGen.exe
```

---

## 🧭 Roadmap

As versões estão planejadas em etapas evolutivas:

| Versão | Foco | Principais Entregas |
|---------|------|----------------------|
| **V1** | Base | Wails UI + Core + DOCX Renderer + LocalFS |
| **V2** | Experiência | Configurações, ordenação, logs, presets e melhorias de UX |
| **V3** | Integração | Suporte a Google Drive (import/export) |
| **V4** | Expansão | Renderer PDF + templates HTML |

> Veja o arquivo [`ROADMAP.md`](./ROADMAP.md) para detalhes de cada versão.

---

## 🧑‍💻 Stack Técnica

| Categoria | Tecnologia |
|------------|-------------|
| Linguagem principal | Go 1.22+ |
| Framework desktop | Wails (Go + JS/TS) |
| UI | React ou Svelte |
| Relatórios | `gomutex/godocx` (DOCX) / HTML-to-PDF (futuro) |
| Armazenamento | Sistema de arquivos local / Google Drive (V3+) |
| Build | Wails CLI |
| Configurações | JSON local persistido |

---

## 🪶 Licença

MIT © Rodrigo P. Melo  
Desenvolvido como ferramenta interna com potencial público futuro.

---

## 🏛️ Versão Original

A primeira versão deste projeto foi escrita integralmente em Go com Fyne.  
Você pode consultar o código original aqui:  
👉 [ImageReportGen (versão Go/Fyne)](https://github.com/RodrigoPMelo/ImageReportGen)
