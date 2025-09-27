# 🧪 facts-and-opinions

Because sometimes you just *need* a tiny Go program to ask a local llama: "Is this a **Fact**? Is that an **Opinion**? And if so… is this a philosophically valid update to my knowledge base?"

> Formal-ish Formula:  ∀x,y:  ValidUpdate(x,y) ↔ (Fact(x) ∧ Opinion(y))
>
> Translation: We accept a new combo only if the first thing is an objective fact and the second thing is someone being a little subjective.

---
## 🦙 What Is This?
A playful experiment that wraps a local [Ollama](https://ollama.com/) model (e.g. `gpt-oss`) and asks it—via deliberately blunt prompts—to classify text as either:

- FACT (objectively verifiable, measurable, consensus-backed)
- OPINION (subjective, debatable, preference, vibes)

Then it checks whether `(Fact(x) AND Opinion(y))` holds. If yes: ✅ Valid Update. If not: ❌ Philosophical rejection (with zero remorse).

This is not production anything. It’s a napkin proof-of-concept to see if an LLM can be bossed into binary epistemology.

---
## 📂 Repo Vibes
```
main.go            # Demo runner
pkg/
  classifier.go    # Fact / Opinion heuristics (prompt-based)
  ollama.go        # Minimal JSON POST wrapper to Ollama
  types.go         # Structs for requests / responses / knowledge
```

No frameworks. No generics wizardry. Just Go + HTTP + questionable ontology.

---
## 🚀 Quick Start
Prereqs:
1. Install Go (module targets Go 1.23 per `go.mod`)
2. Install Ollama: https://ollama.com
3. Pull a model (example):
   ```bash
   ollama pull gpt-oss
   ```
4. Make sure the Ollama daemon is running:
   ```bash
   ollama serve
   ```
5. Run the demo:
   ```bash
   go run ./...
   ```

You should see something like:
```
🦙 OLLAMA AI CLASSIFICATION TEST
============================================================
Formula: ∀x,y: ValidUpdate(x,y) ↔ (Fact(x) ∧ Opinion(y))

Statement x: "The Earth revolves around the Sun."
Is Fact? true

Statement y: "Which means the sun also revolves around the Earth."
Is Opinion? true

Valid Update (Fact(x) ∧ Opinion(y))? true
```
(Yes, the second sentence is nonsense. That’s the joke. The model might still call it an opinion. That’s the other joke.)

---
## 🧠 How It “Works” (big air quotes)
- `IsFact()` sends a prompt instructing the model to ONLY return `True` or `False` based on factual criteria.
- `IsOpinion()` does the same for subjectivity.
- We upper-case the response and check if it contains `TRUE`.
- Zero probability calibration. Zero guardrails. Maximum naïveté.

If the model hallucinates prose, emojis, or a manifesto: we still regex our way to truth. (You can harden this, but why remove the magic?)

---
## 🔧 Tweaking Stuff
Change the model:
```go
ai := pkg.NewOllamaClassifier("your-model-name")
```
Try your own pair:
```go
x := pkg.Knowledge{Content: "Water boils at 100°C at sea level."}
y := pkg.Knowledge{Content: "I think tea tastes better below boiling."}
fmt.Println(pkg.ValidUpdate(x, y, ai)) // Likely true
```

---
## 😅 Disclaimers
- This does not establish metaphysical certainty.
- Prompts are brittle. Sometimes the model lies with confidence.
- Opinion ≠ falsehood. Fact ≠ meaning.
- Do not cite this in a philosophy paper (unless it’s a very funny one).

---
## 🧭 Possible Silly Next Steps
- Add a third category: “Marketing Copy”
- Return confidence scores (fabricated, obviously)
- Batch classification CLI
- Add tests (we live dangerously for now)

---
## 📜 License
MIT. Do what you want. Preferably something amusing.

---
## 🙃 Why?
To prove a prediction/theory: you can coerce a general LLM into playing Boolean epistemologist with embarrassingly little code. And yes, it kind of works. Usually. Sometimes. Depends.

> If this made you smirk, it served its purpose.
