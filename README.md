# LLM Firewall 

![License](https://img.shields.io/github/license/itsVentie/ebpf-llm-firewall)
![Go Version](https://img.shields.io/badge/go-1.22-blue)
Low-latency LLM gateway designed for production security, featuring a modular architecture with eBPF support.

## Project Structure
```text
ebpf-llm-firewall/
├── cmd/proxy/           # Application entry point
├── internal/firewall/   # Core logic: scanning & validation
├── internal/ebpf/       # eBPF kernel-level monitors
├── docker/              # Deployment configurations
└── Makefile             # Build automation
```

## Getting Started
### Prerequisites
Go 1.22+
Docker

Building & Running
Bash
# Build the proxy
make build

# Run the proxy
make run

## Features
Token Redaction: High-speed scanning for API keys and PII.

Injection Defense: Heuristic-based prompt injection detection.

Modular Design: Clear separation of concerns for easier testing and maintenance.

## Roadmap
[ ] Implement eBPF socket monitoring for kernel-level security.

[ ] Add advanced configuration via config.yaml.

[ ] Benchmarking and performance optimizations.