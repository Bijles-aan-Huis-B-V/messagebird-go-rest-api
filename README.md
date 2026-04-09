# MessageBird Go REST API (Fork)

> **Consider replacing.** The upstream repository is no longer maintained. We may replace this with a minimal HTTP client for the specific endpoints we use (WhatsApp HSM templates via notification-service and hubspot-service).

## Overview

Fork of the MessageBird Go REST API client. Provides a Go SDK for MessageBird's SMS, Voice, Conversations, WhatsApp, and other APIs.

## Usage

```go
client := messagebird.New(accessKey)
msg, err := sms.Read(client, "message-id")
```

## Used By

- notification-service (WhatsApp messaging)
- hubspot-service (SMS delivery)

## Commands

```bash
go test ./...           # Run all tests
go build ./...          # Build (library only)
```
