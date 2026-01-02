# Game Launcher

A desktop game launcher built with Wails, Svelte, and TypeScript. Handles authentication, automatic patching, and game client launching.

## Features

- **Authentication** - Login via authentication server, stores access/refresh tokens
- **Auto-Patching** - Compares local files against manifest, downloads updates from patching server
- **Game Launch** - Retrieves game credentials and pipes them to the game client
- **Profile Display** - Shows user profile and avatar from auth server

## Tech Stack

- **Wails** - Go backend with native webview frontend
- **Svelte + TypeScript** - Reactive UI components
- **Tailwind CSS** - Styling
- **Vite** - Frontend build tooling

## How It Works

1. User logs in → receives JWT access token
2. On "Play" button → fetches game credentials from `/game/credentials`
3. Launcher pipes username + API key to game client
4. Game client hashes API key and authenticates with legacy MySQL database

## Development

```bash
wails dev
```

Dev server available at http://localhost:34115 for browser debugging.

## Building

```bash
wails build
```

Produces a redistributable executable.
