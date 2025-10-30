### 🗺️ **ROADMAP.md**

# 🗺️ Roadmap — ImageReportGen

## 🎯 Objetivo Geral
Transformar o antigo gerador de relatórios em Go/Fyne em uma aplicação desktop moderna, modular e extensível, capaz de gerar relatórios DOCX/PDF com templates e integração a serviços em nuvem.

---

## 🧩 Versão 1 — MVP Moderno (Wails + DOCX + LocalFS)

**Objetivo:** recriar o fluxo atual, mantendo funcionalidade e compatibilidade com templates DOCX.

### Entregas principais
- Estrutura base Wails (Go backend + React/Vue frontend)
- Módulo `DocxRenderer` (usando `github.com/gomutex/godocx` ou similar)
- Upload/seleção de imagens (drag & drop e ZIP)
- Seleção de template DOCX
- Configuração de layout padrão (Paisagem 3×1 / Retrato 2×2)
- Campos opcionais de cabeçalho (cliente, município, estado, data)
- Geração 100% offline
- Limpeza automática de temporários
- “Abrir relatório após gerar” (opcional)

---

## 🧭 Versão 2 — Recursos Avançados (UX + Configuração)

**Objetivo:** melhorar experiência do usuário e permitir personalização do layout.

### Entregas principais
1. Salvar configurações do usuário
2. Ordenação manual das fotos
3. Renomear arquivo de saída (“Salvar como…”)
4. Abrir relatório automaticamente após gerar
5. Presets de layout (padrão / compacto / detalhado)
6. Validação e recuperação de imagens corrompidas
7. Log de geração (`.log` junto ao DOCX)
8. Exportar DOCX e PDF (convertendo localmente via Office se disponível)
9. Adicionar carimbo automático de data
10. Controle de versão do template

---

## ☁️ Versão 3 — Integração com Google Drive

**Objetivo:** permitir sincronização automática de imagens e relatórios.

### Entregas principais
- Autenticação via OAuth (Google)
- Escolha de pasta origem (imagens) e destino (relatórios)
- Download de imagens e upload automático de DOCX
- Cache local para uso offline
- Status de sincronização

---

## 🧾 Versão 4 — Suporte a PDF nativo

**Objetivo:** permitir geração direta de PDF sem depender do Word.

### Entregas principais
- Novo `PdfRenderer` (HTML-to-PDF)
- Templates HTML equivalentes aos DOCX
- Configuração de layout dinâmica (linhas, colunas, margens)
- Conversão DOCX → PDF (fallback)
- Preview fiel (opcional)

---

## 🔮 Futuro

- Editor visual de template (web-based)
- Suporte a legendas por foto
- Geração por lote (multi-relatório)
- API local (para integração com outros sistemas)
- Distribuição pública / instalador

---

## 📅 Cronograma sugerido (estimado)

| Versão | Duração | Entregas principais |
|---------|----------|--------------------|
| V1 | 3-4 semanas | Base Wails + DOCX Renderer + UI mínima |
| V2 | 4-6 semanas | Configs, layouts, UX aprimorado |
| V3 | 3-5 semanas | Google Drive |
| V4 | 3-5 semanas | PDF Renderer |

---

## 🧑‍💻 Stack técnica

| Camada | Tecnologia |
|--------|-------------|
| UI | Wails + React (ou Svelte) |
| Backend | Go 1.22+ |
| Relatórios | godocx (V1–V3) / gofpdf ou HTML-to-PDF (V4) |
| Cloud | Google Drive API (V3+) |
| Build | Wails CLI |
| Configuração | JSON local (`user_config.json`) |


## 📜 Licença

MIT © Rodrigo P. Melo
