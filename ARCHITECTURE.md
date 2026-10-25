# Architecture

OreoBot is a modular monolith. Each top-level directory in `modules/` owns a
capability; package imports within that module are implementation details.
`main.go` is the composition root and wires the modules together.

Each module root has `doc.go` and `module_info.go`. Its `Descriptor`
implements GoModulith's `module.Provider` and reports an ID, Go import path,
description, version, dependencies, and exported contracts. This is the Go
equivalent of the feature declaration in Java's `package-info.java`.

## Modules

| Module | Responsibility | Allowed module dependencies |
| --- | --- | --- |
| `antiscam` | Message risk scoring, OCR, phishing feeds, and moderation actions | `config`, `localization`, `persistence`, `shared` |
| `config` | Configuration loading and typed settings | none |
| `commissions` | Commission requests, quotes, and invoices | `config`, `localization`, `payments`, `persistence`, `tickets` |
| `giveaways` | Giveaway creation, entries, timed endings, and rerolls | `config`, `localization`, `persistence`, `shared` |
| `localization` | Language file loading and translation access | none |
| `persistence` | JSON state and database access | `config` |
| `shared` | Framework independent value helpers shared by features | none |
| `minecraft` | RCON client | none |
| `moderation` | Member moderation commands, warning records, and mod logs | `config`, `localization`, `persistence`, `shared` |
| `tickets` | Support ticket commands, panels, channels, and transcripts | `config`, `localization`, `persistence` |
| `music` | Playback and Lavalink integration | `config` |
| `payments` | Payment orchestration, provider gateways, polling, and webhooks | `config`, `persistence` |
| `discord` | Discord bot startup and remaining feature handlers | `antiscam`, `commissions`, `config`, `giveaways`, `localization`, `minecraft`, `moderation`, `music`, `payments`, `persistence`, `shared`, `tickets` |

The dependency allowlist and cycle check live in `architecture_test.go` and
use GoModulith's source analyzer. A module can only import other modules on its
allowlist. Moderation, tickets, commissions, giveaways, and anti-scam now own
their Discord adapters and command/event routing. Modules use `api/` for
feature contracts, `domain/` for isolated business rules where they exist, and
`infrastructure/` for Discord, storage, and provider integration. The remaining
Discord handlers are shared features still awaiting extraction.

## Useful GoModulith commands

```sh
go run github.com/el211/GoModulith/cmd/gomodulith@v0.0.0-20261001095035-8bef35d50d71 verify -root . -dir modules
go run github.com/el211/GoModulith/cmd/gomodulith@v0.0.0-20261001095035-8bef35d50d71 graph -root . -dir modules
go run github.com/el211/GoModulith/cmd/gomodulith@v0.0.0-20261001095035-8bef35d50d71 docs -root . -dir modules
```
