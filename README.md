# Calebs Company — Polyglot Monorepo

A full-stack platform that spans the entire software engineering stack, using each language for what it does best.

## Architecture

```
calebs_company/
├── web/            TypeScript · Next.js 14 (React) — frontend
├── api/            Python · FastAPI — REST backend, business logic, DB access
├── auth/           Go — JWT issuance & verification service
├── realtime/       Go — WebSocket server, pub/sub rooms
├── crypto/         Rust — Argon2 password hashing, SHA-256, secure token gen
├── engine/         C++ — pricing calculator, statistical compute
├── data-pipeline/  Python — ETL / data processing (future)
├── infra/          Terraform, Kubernetes manifests, Docker configs
└── docker-compose.yml
```

## Services & Ports

| Service    | Language      | Port  | Description                        |
|------------|---------------|-------|------------------------------------|
| `web`      | TypeScript    | 3000  | Next.js frontend                   |
| `api`      | Python        | 8000  | FastAPI REST backend               |
| `auth`     | Go            | 8001  | JWT auth service                   |
| `realtime` | Go            | 9000  | WebSocket pub/sub                  |
| `postgres` | —             | 5432  | PostgreSQL database                |

## Quick Start

```bash
# Start all services
docker-compose up --build

# Or run individually for development:

# Frontend
cd web && npm install && npm run dev

# API
cd api && pip install -r requirements.txt && uvicorn app.main:app --reload

# Auth service
cd auth && go run ./cmd/server

# Realtime service
cd realtime && go run ./cmd/server

# Rust crypto (build)
cd crypto && cargo build --release

# C++ engine (build)
cd engine && cmake -B build && cmake --build build
```

## Language Rationale

- **TypeScript** — type-safe React frontend; industry standard for web UIs
- **Python** — ergonomic for APIs, ORM, data science, scripting
- **Go** — concurrency model ideal for WebSockets and high-throughput auth
- **Rust** — memory-safe crypto primitives; zero-cost abstractions
- **C++** — maximum performance for numerical compute

## Topics Covered

- [ ] REST API design (Python/FastAPI)
- [ ] WebSocket real-time communication (Go)
- [ ] JWT authentication (Go)
- [ ] Password hashing & cryptography (Rust)
- [ ] Pricing / stats computation engine (C++)
- [ ] PostgreSQL + async ORM (Python/SQLAlchemy)
- [ ] Docker containerisation
- [ ] docker-compose local orchestration
- [ ] Kubernetes manifests (`infra/k8s/`)
- [ ] Cloud infrastructure with Terraform (`infra/terraform/`)
- [ ] CI/CD pipelines (`.github/workflows/`)
- [ ] Security: secrets management, SAST, dependency scanning
