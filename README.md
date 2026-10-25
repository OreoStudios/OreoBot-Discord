# OreoBot

OreoBot is a feature-rich Discord bot built in Go for OreoStudios. It brings moderation, community tools, support tickets, commissions, giveaways, payments, music, and optional Minecraft integration together in one configurable application.

The codebase is organized as a **modular monolith**: the features run together as one bot, while their code and dependencies are divided into clear modules. That structure is documented and checked with **GoModulith**, my Go library for describing and verifying modular application architectures.

## Features

- **Moderation and safety:** moderation commands, warnings and logs, anti-spam settings, no-ping controls, link filtering, phishing feed checks, and optional scam-image text scanning.
- **Community tools:** welcome and leave messages, invite tracking, autoroles, counting games, custom commands, and product role verification.
- **Tickets and commissions:** ticket panels and transcripts, commission requests and quotes, invoice workflows, and payment provider integrations.
- **Payments:** PayPal, Stripe, and Coinbase integrations, with polling or webhook notifications where supported by the provider configuration.
- **Giveaways:** create and manage giveaways, collect entries, and restore giveaway timers after a restart.
- **Music:** direct playback or a Lavalink backend, with queue and playback controls.
- **Minecraft (optional):** RCON management, Discord-to-Minecraft account linking, and player profile features.
- **GitHub notifications (optional):** subscribe Discord channels to repository activity and releases.
- **Storage:** SQLite by default, optional MongoDB, and JSON-backed guild state.
- **Chat bridge (optional):** connect a Discord channel with a Minecraft server through RabbitMQ.

Most optional features are disabled in the sample configuration. Configure and enable only the integrations you use.

## Quick start

### Requirements

- Go 1.24 or newer
- A Discord application and bot token
- FFmpeg and yt-dlp if you use direct music playback
- The external services you enable, such as MongoDB, RabbitMQ, Lavalink, or payment providers

### Run from source

```sh
git clone https://github.com/OreoStudios/OreoBot-Discord.git
cd OreoBot-Discord
cp config.example.json config.json
```

Open `config.json` and set `discord.token` near the top of the file. Set `discord.guild_id` if you want application commands registered to one server while configuring the bot. Then start OreoBot:

```sh
go run . -config config.json
```

OreoBot registers its Discord application commands when it starts. Invite the bot with the `bot` and `applications.commands` scopes, and grant the Discord permissions needed by the features you enable. Enable the required privileged gateway intents in the Discord Developer Portal for features that need them.

### Build a binary

Build for the current operating system:

```sh
go build -o oreobot .
```

Build a Linux amd64 binary, for example for a Linux server or Pterodactyl:

```sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bot.bin .
```

Start a built binary with:

```sh
./bot.bin -config config.json
```

On Windows, run the generated executable from PowerShell and use `OreoBot.exe` as the output name if desired.

## Configuration

`config.example.json` documents the available settings. Copy it to `config.json`, then fill in the settings for your server and integrations. The Discord token is configured at `discord.token`. The bot exits with an error if the token is missing.

The `enabled` flags control optional features. For example, set `music.enabled` to `true` to enable music commands, or `minecraft.enabled` to `true` to enable Minecraft commands. See `modules/config/config.go` for the typed configuration model and `config.example.json` for the full set of options.

Keep real bot tokens, API keys, webhook secrets, payment credentials, and database credentials private. Do not commit live credentials to a public repository; use private deployment configuration and rotate any credential that has been exposed.

## Pterodactyl

Two egg definitions are included in `.claude/`:

- [`egg-oreo-bot-discord.json`](.claude/egg-oreo-bot-discord.json) downloads `bot.bin` from the latest GitHub release. It downloads `config.json` and `lang.yml` from the repository during installation.
- [`egg-oreo-bot-local-binary.json`](.claude/egg-oreo-bot-local-binary.json) runs a `bot.bin` you upload to the server. It still downloads `config.json` and `lang.yml` from the repository.

For the local-binary egg, upload the Linux amd64 `bot.bin` to the server’s root directory before starting the server. For either egg, configure `config.json` with your bot token and server settings after installation.

## Architecture and GoModulith

OreoBot uses [GoModulith](https://github.com/el211/GoModulith), my Go library for making modular-monolith boundaries explicit and verifiable. Each module provides metadata in `module_info.go` and package documentation in `doc.go`. The descriptors identify the module, describe its purpose, and declare its dependencies and public contracts. This is similar in spirit to using Java’s `package-info.java` to document a package, with GoModulith providing the module-level metadata and architecture checks.

GoModulith is used here to:

- describe each module and its dependencies in Go;
- check module dependency rules and detect dependency cycles;
- keep the intended architecture reviewable alongside the source code.

This is a build-time/source-analysis tool, not a runtime plugin system. The bot remains one Go application, with `main.go` wiring the feature modules together. For the module map and dependency rules, see [`ARCHITECTURE.md`](ARCHITECTURE.md).

The main feature modules are:

| Module | Responsibility |
| --- | --- |
| `antiscam` | Scam and phishing detection, feeds, and moderation actions |
| `commissions` | Commission workflows, quotes, and invoices |
| `config` | Typed configuration loading |
| `discord` | Bot startup and shared Discord handlers |
| `giveaways` | Giveaway lifecycle and interactions |
| `localization` | Language file loading and translations |
| `minecraft` | Minecraft RCON client support |
| `moderation` | Moderation commands and domain rules |
| `music` | Playback and Lavalink integration |
| `payments` | Payment gateways, polling, and webhooks |
| `persistence` | Database and state storage |
| `shared` | Shared value types and Discord helpers |
| `tickets` | Support ticket workflows and transcripts |

Feature modules use `api/` for public contracts, `domain/` for business rules where they exist, and `infrastructure/` for Discord or provider integrations. Not every module needs every layer.

### Check the architecture

The repository includes an architecture test in `architecture_test.go`. Run it with:

```sh
go test ./...
```

You can also use the GoModulith CLI to verify or inspect the module graph:

```sh
go run github.com/el211/GoModulith/cmd/gomodulith@v0.0.0-20261001095035-8bef35d50d71 verify -root . -dir modules
go run github.com/el211/GoModulith/cmd/gomodulith@v0.0.0-20261001095035-8bef35d50d71 graph -root . -dir modules
go run github.com/el211/GoModulith/cmd/gomodulith@v0.0.0-20261001095035-8bef35d50d71 docs -root . -dir modules
```

## Language files

`lang.yml` contains the default bot messages. Keep it beside the executable when running OreoBot, or in the working directory used to start it.

## Contributing

When adding a feature, place it in the module that owns that capability, update its module metadata when its dependencies or public contracts change, and keep the architecture rules in `architecture_test.go` accurate. Read [`ARCHITECTURE.md`](ARCHITECTURE.md) before changing module boundaries.

## Links

- [OreoBot repository](https://github.com/OreoStudios/OreoBot-Discord)
- [GoModulith repository](https://github.com/el211/GoModulith)
- [OreoBot wiki](oreobot-wiki.html)
