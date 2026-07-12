---
title: Ebpf Llm Firewall
emoji: 🛡️
colorFrom: pink
colorTo: purple
sdk: docker
pinned: false
license: apache-2.0
---

# eBPF Contextual LLM Firewall & Safe-Guard Proxy

A high-performance, zero-allocation proxy engineered in Go to intercept, inspect, and sanitize LLM prompts and responses in real-time. 

### Core Features
* **Token/Credential Masking:** High-speed scanning and redaction of API keys, private tokens, and sensitive PII before they hit the LLM context.
* **Prompt Injection Defense:** Structural and heuristic validation of incoming JSON objects.
* **Low Latency Engine:** Minimal memory footprints and sub-millisecond processing times.