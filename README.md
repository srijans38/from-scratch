# From Scratch

Learning by building low-level things from first principles.

## Philosophy

Software abstractions are powerful, they let us build quickly without understanding every layer beneath. But abstractions hide. A web framework hides HTTP. A database library hides storage engines. A container runtime hides namespaces and resource limits. For experienced engineers, there's immense value in peeling back those layers and building the thing yourself, seeing exactly what trade-offs the original designers made and why.

This repository is a collection of projects built with that philosophy: implement real systems from the ground up, with no helper libraries or convenient abstractions. The goal isn't to create production software, it's to understand how production software works by recreating its core mechanisms.

Each project is an exploration: what are the actual challenges when you remove the abstraction? What invariants must you maintain? What design decisions matter? Building forces understanding in a way that reading never can. And once you've built something, the next step is to teach it, to create visualizations that show how the system works, so others can see the internal mechanics the way you now do.

This isn't about being a purist or reinventing the wheel in production. It's about depth: taking time to understand the foundations so thoroughly that you can recognize patterns, spot bottlenecks, and design better systems in the future.

## What's in This Repo

A collection of independent projects, each exploring a different system or problem domain.

Every project includes:

- **The implementation** — built from scratch in Go(mostly), with minimal dependencies
- **Interactive visualizations** — diagrams and animations that explain how the system works, designed to be embedded in technical writing

## Projects

- `http-server/` — HTTP/1.1 server from raw TCP sockets

*More projects will be added as they're completed.*

