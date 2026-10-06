# messagebird-go-rest-api

> **Status: to be replaced.** This is our fork of the MessageBird Go SDK, and upstream is no longer maintained.
> The plan is to replace it with a small HTTP client inside each consumer that covers only the calls we make.
> It cannot be archived yet: notification-service and hubspot-service still import it.

## What it is

A fork of `github.com/messagebird/go-rest-api` (upstream v9.1.0, July 2022) under the module path
`github.com/Bijles-aan-Huis-B-V/messagebird-go-rest-api` (`go 1.16`). It is a library only.

Usage pattern: every API package takes a `messagebird.Client` as its first argument.

```go
client := messagebird.New(accessKey)
_, err := conversation.SendMessage(client, req)
```

Changes made in the fork (2022-11 to 2024-02):

- Module path renamed.
- `contact/` moved to the Contacts v2 API (`https://contacts.messagebird.com/v2/contacts`), with identifiers and
  `contact.Upsert` (`POST /v2/ops/contacts/upsert`).
- New `integration/` package for WhatsApp templates on `https://integrations.messagebird.com`
  (`CreateWhatsAppTemplate`, `ListWhatsAppTemplates` on `/v3/platforms/whatsapp/templates`,
  `DeleteWhatsAppTemplate`). The 2024-02-08 commit sends the list filters as query parameters instead of a GET
  body.
- `secret.go`: `messagebird.Secret{Key, Channels}`, the shape of the per-country secret notification-service
  loads.

## Why it is still around

Both consumers pin `v0.0.0-20240208154711-421306c3a4d4`, which is the newest code commit (later commits only
touch docs).

| Consumer | What it calls | Purpose |
| --- | --- | --- |
| notification-service | `messagebird.New`, `messagebird.Secret`, `conversation.SendMessage` (HSM template messages), `integration.ListWhatsAppTemplates` and its component constants | send WhatsApp template messages; sync WhatsApp templates |
| hubspot-service | `messagebird.New`, `contact.Upsert`, `contact.CreateRequest` / `Identifier` | keep MessageBird contacts in sync with users |

Both read the API keys from the Secrets Manager secret named by their `MESSAGE_BIRD` env var.

## Before it can be archived

1. notification-service: replace the conversation send and the WhatsApp template list with a local HTTP client
   (Conversations API `https://conversations.messagebird.com/v1`, Integrations API v3), keeping the per-country
   `{key, channels}` secret format or migrating it.
2. hubspot-service: replace `contact.Upsert` with a local client for the Contacts v2 upsert call.
3. Remove the module from both `go.mod` files, then archive this repo.

## Build and test

There is nothing to deploy; consumers pick up changes by bumping the pseudo-version in their `go.mod`.

```bash
go build ./...
go test ./...             # uses the fake TLS server in internal/mbtest
go test ./sms/ -run TestCreateMessage
```

CI: `.github/workflows/tests.yml` runs `go test ./...` on every push and PR with Go 1.16.x, 1.17.x and 1.18.x
(the upstream workflow, unchanged).

## Git Workflow

Branch from `master`, open a PR, merge after review. Avoid new features here; put new MessageBird calls in the
consumer instead. See [`../bah-knowledge/DEVELOPMENT.md`](../bah-knowledge/DEVELOPMENT.md).
