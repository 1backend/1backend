---
sidebar_position: 30
tags:
  - prompt-svc
  - prompts
  - ai
  - services
---

# Prompt Svc

The prompt service provides an easy to use interface to prompt LLMs and use AI models. Aims to serve humans and machines alike with its resilient queue based architecture.

> This page provides a high-level overview of `Prompt Svc`. For detailed information, refer to the [Prompt Svc API documentation](/docs/openorch/prompt).

## Responsibilities

The prompt service:

- Accepts prompts
- Maintains a list of prompts
- Processes prompts as soon as it's able to
- Streams prompt answers
- Handles retries of prompts that errored with an exponential backoff

It's able to stream back LLM responses, or it can respond syncronously if that's what the caller wants, for details see the [Add Prompt (`/prompt-svc/prompt`) Endpoint](/docs/openorch/prompt).

##

## Dependencies

- [Chat Svc](/docs/built-in-services/chat-svc) to save prompt responses to chat threads and messages
- [Model Svc](/docs/built-in-services/model-svc) to get the address of the running AI models, see their status etc.

## Current limitations

There are planned improvements for this service:

- It should manage models: start needed ones and stop unneeded ones based on the volume of prompts in the backlog
