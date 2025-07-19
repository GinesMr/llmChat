# llmChat

Minimal test version of AiBasicStudio — a lightweight terminal-based LLM assistant specialized in blockchain and cryptocurrency domains.

## Overview

`llmChat` is a minimal proof-of-concept CLI application built in Go using [Cobra](https://github.com/spf13/cobra). It allows you to interact with a local LLM (such as Ollama with `gemma2:2b`) from your terminal, enforcing a focused expert persona for blockchain and DeFi queries.

This is the simplest version of what will later become part of AiBasicStudio.

## Features

- Terminal-based chat interface
- Crypto/Blockchain expert system prompt
- Streaming LLM output via `langchaingo`
- Built using Cobra CLI framework
- Easy to extend and modular

## Requirements

- Go 1.21 or higher
- Ollama installed and running locally
- Model installed locally (e.g., `gemma2:2b`)

## Installation

Clone the repository and build the binary:

```bash
git clone https://github.com/your-user/llmChat.git
cd llmChat
go build -o llmChat
