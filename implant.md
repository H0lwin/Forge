# پرامپت نهایی — توصیف خروجی و تجربه کاربری

---

```
You are an expert Go developer. Build a complete, production-grade CLI tool 
called "forge" written in Go.

Do NOT describe what you will do. Write the complete, runnable source code 
for every file. No placeholders. No TODOs. Fully working.

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
WHAT THE USER SEES — EXACT TERMINAL EXPERIENCE
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

When user types: forge

┌─────────────────────────────────────────────────────┐
│                                                     │
│    ███████╗ ██████╗ ██████╗  ██████╗ ███████╗      │
│    ██╔════╝██╔═══██╗██╔══██╗██╔════╝ ██╔════╝      │
│    █████╗  ██║   ██║██████╔╝██║  ███╗█████╗        │
│    ██╔══╝  ██║   ██║██╔══██╗██║   ██║██╔══╝        │
│    ██║     ╚██████╔╝██║  ██║╚██████╔╝███████╗      │
│    ╚═╝      ╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚══════╝      │
│                                                     │
│         ⚡ Professional Project Scaffolding          │
│              Build faster. Ship better.             │
│                    v1.0.0                           │
└─────────────────────────────────────────────────────┘

  Available Commands:
  
  ● new          Create a new project interactively
  ● add          Add a feature to existing project
  ● doctor       Check your development environment
  ● config       Manage forge settings
  ● templates    Browse and manage project templates
  ● version      Show version info

  Run 'forge new' to get started


━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
WHEN USER RUNS: forge new
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

The terminal clears and shows a full interactive TUI wizard.
User navigates with arrow keys, enters text, toggles checkboxes.
Everything is colored, animated, and responsive.

─────────────────────────────────────
STEP 1 of 6 ░░░░░░░░░░░░░░░░░░░░  0%
─────────────────────────────────────

  What type of project are you building?

  ❯ 🌐  Frontend       Websites, web apps, SPAs
    🔧  Backend        APIs, servers, microservices  
    📱  Mobile         iOS, Android, cross-platform
    🔀  Fullstack      Frontend + Backend together


─────────────────────────────────────
STEP 2 of 6 ████░░░░░░░░░░░░░░░░  20%
─────────────────────────────────────

  Choose your framework:

  ❯ ⚡  Next.js         React framework with SSR/SSG
    🔷  Vite            Fast build tool (React/Vue/Vanilla)
    🚀  Astro           Content-focused static sites


[User selects Backend → shows:]

  ❯ 🐍  Django          Batteries-included Python framework
    ⚡  FastAPI          Modern, fast Python API framework
    🌶️  Flask           Lightweight Python microframework
    🟢  Express         Minimal Node.js web framework
    🏠  NestJS          Structured Node.js framework


─────────────────────────────────────
STEP 3 of 6 ████████░░░░░░░░░░░░  40%
─────────────────────────────────────

  Project details:

  Project name  ›  my-awesome-api
                   ✓ Valid name

  Project path  ›  ~/projects/my-awesome-api
                   ✓ Path is available

  Python ver.   ›  3.11  ▼


─────────────────────────────────────
STEP 4 of 6 ████████████░░░░░░░░  60%
─────────────────────────────────────

  Choose your environment manager:

  ❯ 📦  venv      Built-in Python virtual environment
    📜  poetry    Dependency management + packaging
    ⚡  uv        Extremely fast Python package installer


─────────────────────────────────────
STEP 5 of 6 ████████████████░░░░  80%
─────────────────────────────────────

  Select extras  (space to toggle, enter to confirm)

  Core
  ◉  Git init + initial commit
  ◉  README.md template
  ◉  .editorconfig
  ○  Makefile with common commands

  Database
  ◉  PostgreSQL (docker-compose)
  ○  MySQL
  ○  SQLite only

  Django Extras
  ◉  Django REST Framework
  ○  Celery + Redis (async tasks)
  ○  Django Channels (WebSockets)
  ○  Simple JWT authentication
  ○  django-cors-headers

  Dev Tools
  ◉  Docker + docker-compose
  ◉  GitHub Actions CI/CD
  ○  Pre-commit hooks
  ○  Pytest + coverage setup
  ○  Sentry error tracking config


─────────────────────────────────────
STEP 6 of 6 ████████████████████  100%
─────────────────────────────────────

  Review your project before creation:

  ╔═══════════════════════════════════════════════════╗
  ║  📋 Project Summary                               ║
  ╠═══════════════════════════════════════════════════╣
  ║  Name       my-awesome-api                        ║
  ║  Framework  Django 5.x                            ║
  ║  Language   Python 3.11                           ║
  ║  Env Mgr    venv                                  ║
  ║  Path       ~/projects/my-awesome-api             ║
  ╠═══════════════════════════════════════════════════╣
  ║  Extras                                           ║
  ║  ✓ Git init      ✓ README       ✓ Docker          ║
  ║  ✓ DRF           ✓ PostgreSQL   ✓ GitHub Actions  ║
  ╚═══════════════════════════════════════════════════╝

  ❯ 🚀  Create Project
    ←   Go Back  
    ✕   Cancel


━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
DURING PROJECT CREATION — LIVE PROGRESS
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

After user confirms, the terminal shows live animated progress:

  ⚡ Forging your project...
  ────────────────────────────────────────────────

  ✓  [ 1/12]  Creating project directory
  ✓  [ 2/12]  Initializing Python 3.11 virtual environment
  ✓  [ 3/12]  Installing Django 5.0
  ✓  [ 4/12]  Installing Django REST Framework
  ✓  [ 5/12]  Scaffolding project structure
  ✓  [ 6/12]  Configuring settings (base/dev/prod)
  ✓  [ 7/12]  Generating .env + .env.example
  ✓  [ 8/12]  Running initial migrations
  ⠸  [ 9/12]  Setting up Docker & docker-compose...
  ░  [10/12]  Creating GitHub Actions workflow
  ░  [11/12]  Generating README.md
  ░  [12/12]  Git init + initial commit

  ────────────────────────────────────────────────
  ⏱  Elapsed: 00:23

[If a step fails:]

  ✗  [ 9/12]  Setting up Docker — docker not found

  What would you like to do?
  ❯ Skip this step and continue
    Retry
    Abort setup


━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
AFTER SUCCESSFUL CREATION
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

  ╔═════════════════════════════════════════════════════╗
  ║                                                     ║
  ║   ✅  Project created successfully!                  ║
  ║                                                     ║
  ╠═════════════════════════════════════════════════════╣
  ║                                                     ║
  ║   📁  ~/projects/my-awesome-api                     ║
  ║   ⏱  Built in 31 seconds                           ║
  ║   📦  12 steps completed                            ║
  ║                                                     ║
  ╠═════════════════════════════════════════════════════╣
  ║   Project structure                                 ║
  ║                                                     ║
  ║   my-awesome-api/                                   ║
  ║   ├── config/          Django project config        ║
  ║   │   ├── settings/    base / dev / prod            ║
  ║   │   ├── urls.py                                   ║
  ║   │   └── wsgi.py                                   ║
  ║   ├── apps/            Your Django apps here        ║
  ║   │   └── core/                                     ║
  ║   ├── .venv/           Virtual environment          ║
  ║   ├── .env             Your secrets (git ignored)   ║
  ║   ├── .env.example     Commit this one              ║
  ║   ├── requirements.txt                              ║
  ║   ├── Dockerfile                                    ║
  ║   ├── docker-compose.yml                            ║
  ║   └── README.md                                     ║
  ║                                                     ║
  ╠═════════════════════════════════════════════════════╣
  ║   🚀  Next steps                                    ║
  ║                                                     ║
  ║   cd my-awesome-api                                 ║
  ║   source .venv/bin/activate                         ║
  ║   python manage.py runserver                        ║
  ║                                                     ║
  ║   Or with Docker:                                   ║
  ║   docker compose up --build                         ║
  ║                                                     ║
  ╠═════════════════════════════════════════════════════╣
  ║   💡  Useful commands                               ║
  ║                                                     ║
  ║   forge add celery       Add Celery to this project ║
  ║   forge add auth         Add JWT authentication     ║
  ║   forge doctor           Check your environment     ║
  ╚═════════════════════════════════════════════════════╝

  Open in VS Code? [y/N]  ›


━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
WHEN USER RUNS: forge doctor
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

  🩺 Checking your development environment...

  ┌──────────────┬──────────────┬────────────┬───────┐
  │ Tool         │ Status       │ Version    │ Note  │
  ├──────────────┼──────────────┼────────────┼───────┤
  │ git          │ ✓ Installed  │ 2.43.0     │       │
  │ python3      │ ✓ Installed  │ 3.11.7     │       │
  │ node         │ ✓ Installed  │ 20.11.0    │       │
  │ pnpm         │ ✓ Installed  │ 8.15.0     │       │
  │ docker       │ ✓ Running    │ 25.0.2     │       │
  │ flutter      │ ✗ Missing    │ —          │ ⚠    │
  │ uv           │ ✗ Missing    │ —          │       │
  │ poetry       │ ✓ Installed  │ 1.7.1      │       │
  └──────────────┴──────────────┴────────────┴───────┘

  2 tools missing. Install them?

  ❯ Install uv (fast Python installer)
    Install flutter
    Skip for now


━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
WHEN USER RUNS: forge add
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

  [Inside an existing Django project]

  Detected project: Django  ✓

  What would you like to add?

  ❯ 🔐  auth          JWT authentication (Simple JWT)
    📬  celery        Async task queue with Redis
    🗄️  cache         Redis caching layer
    📧  email         Email backend (SendGrid/SMTP)
    🐳  docker        Dockerfile + docker-compose
    🔄  ci            GitHub Actions pipeline
    📊  sentry        Error monitoring
    🧪  pytest        Testing setup with coverage


━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
WHEN USER RUNS: forge config
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

  ⚙️  Forge Configuration  (~/.forge/config.yaml)

  ┌─────────────────────────────────────────────────┐
  │  User                                           │
  │  name      Ali Ahmadi                           │
  │  email     ali@example.com                      │
  │  github    alidev                               │
  ├─────────────────────────────────────────────────┤
  │  Defaults                                       │
  │  git_init          true                         │
  │  open_editor       false                        │
  │  editor            code                         │
  ├─────────────────────────────────────────────────┤
  │  Python                                         │
  │  default_version   3.11                         │
  │  env_manager       venv                         │
  ├─────────────────────────────────────────────────┤
  │  Node                                           │
  │  package_manager   pnpm                         │
  └─────────────────────────────────────────────────┘

  ❯ Edit a setting
    Reset to defaults
    Export config
    Import config


━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
SPECIAL FLAGS — NON-INTERACTIVE MODE
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

For CI/CD and scripting, everything works without interaction:

forge new \
  --name my-api \
  --framework django \
  --path ~/projects \
  --python-version 3.11 \
  --env-manager uv \
  --extras docker,ci,drf,postgres \
  --no-interactive

forge new --name my-app --framework nextjs --extras tailwind,docker --no-interactive

forge new --name my-app --framework django --dry-run
→ Shows exactly what would happen, creates nothing

forge new --name my-app --framework fastapi --verbose
→ Shows every command output in real time


━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
SUPPORTED FRAMEWORKS (all must be fully implemented)
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Frontend:
  next        Next.js (App Router, TS, Tailwind, ESLint, Prettier, Husky)
  vite        Vite (react-ts / vue-ts / vanilla-ts, Tailwind, Router)
  astro       Astro (React/Vue integration, Tailwind, TS)

Backend:
  django      Django (split settings, DRF, venv/poetry/uv, migrations)
  fastapi     FastAPI (folder structure, Alembic, pydantic settings)
  flask       Flask (app factory, blueprints, SQLAlchemy optional)
  express     Express (TS optional, MVC structure, nodemon)
  nestjs      NestJS (Swagger, Prisma optional, JWT optional)

Mobile:
  reactnative React Native (Expo or Bare, Navigation, NativeWind)
  flutter     Flutter (go_router, riverpod/bloc, flavors, dotenv)


━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
TECHNICAL REQUIREMENTS
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Stack:
  Go 1.22+
  github.com/spf13/cobra          CLI commands
  github.com/charmbracelet/huh    Interactive forms
  github.com/charmbracelet/bubbletea  TUI engine
  github.com/charmbracelet/lipgloss   Styling
  github.com/charmbracelet/bubbles    Spinner, progress, table
  github.com/charmbracelet/log        Logging
  github.com/spf13/viper          Config management

Architecture rules:
  - Each framework is a separate file implementing Generator interface
  - Generator interface: Name() / Category() / Steps() / PreCheck() / PostMessage()
  - Runner executes shell commands with streaming output, timeout, retry
  - All template files embedded via go:embed
  - Zero global state — dependency injection only
  - Every error wrapped with context
  - Cross-platform: Linux, macOS, Windows
  - --dry-run flag supported everywhere
  - --no-interactive flag supported everywhere
  - --verbose flag shows raw command output

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
DELIVERABLE
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Write the complete, fully runnable Go source code for every file.
Zero placeholders. Zero TODOs. Every file complete and working.
Start with go.mod → main.go → cmd/ → internal/ → templates/
End with exact build and run instructions.
```