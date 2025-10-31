import { useState } from "react";
import "./style.css";
import { GenerateReport } from "../wailsjs/go/app/ImageReportApp";

function App() {
  const [templatePath, setTemplatePath] = useState("");
  const [imagesRaw, setImagesRaw] = useState("");
  const [output, setOutput] = useState("");
  const [error, setError] = useState("");

  const generate = async () => {
    setError("");
    try {
      const images = imagesRaw
        .split("\n")
        .map((s) => s.trim())
        .filter(Boolean);

      const res = await GenerateReport({
        templatePath,
        outputPath: "",
        clientName: "",
        city: "",
        state: "",
        includeClient: false,
        includeLocation: false,
        includeDate: false,
        completionDate: "",
        images,
      });

      setOutput(res);
    } catch (err: any) {
      setError(err?.toString() ?? "erro ao gerar relatório");
    }
  };

  return (
    <main className="app">
      <h1>ImageReportGen V2</h1>
      <p>Relatório fotográfico em DOCX.</p>

      <label>Template DOCX</label>
      <input
        value={templatePath}
        onChange={(e) => setTemplatePath(e.target.value)}
        placeholder="C:\\modelos\\template.docx"
      />

      <label>Imagens (uma por linha)</label>
      <textarea
        rows={6}
        value={imagesRaw}
        onChange={(e) => setImagesRaw(e.target.value)}
        placeholder="C:\\fotos\\img1.jpg
C:\\fotos\\img2.jpg"
      />

      <button onClick={generate}>Gerar relatório</button>

      {output && (
        <p className="success">
          Relatório gerado em: <strong>{output}</strong>
        </p>
      )}

      {error && <p className="error">{error}</p>}
    </main>
  );
}

export default App;
