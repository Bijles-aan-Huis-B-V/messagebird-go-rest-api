# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

Fork of the MessageBird Go REST API client. Module path: `github.com/Bijles-aan-Huis-B-V/messagebird-go-rest-api` (differs from upstream `github.com/messagebird/go-rest-api`). Provides a Go SDK for MessageBird's SMS, Voice, Conversations, MMS, WhatsApp, and other APIs. Used by `notification-service` (WhatsApp) and `hubspot-service` (SMS).

## Commands

```bash
# Run all tests
go test ./...

# Run tests for a specific package
go test ./sms/
go test ./voice/

# Run a single test
go test ./sms/ -run TestCreateMessage

# Build (library only, no main package)
go build ./...
```

## Architecture

**Client pattern**: `messagebird.Client` is an interface with a single `Request(v, method, path, data)` method. `DefaultClient` implements it, handling auth, HTTP transport, and JSON marshalling. Each API domain package accepts a `Client` as its first argument (not methods on the client).

```go
// Usage pattern - all packages follow this:
client := messagebird.New(accessKey)
msg, err := sms.Read(client, "message-id")
```

**API domain packages**: Each subdirectory (`sms/`, `voice/`, `conversation/`, `contact/`, `balance/`, etc.) is a self-contained package with types, request builders, and API functions. They call `client.Request()` with the appropriate path.

**Multiple API roots**: The default endpoint is `rest.messagebird.com`. Packages with different API roots prepend their own base URL to bypass the default:
- `voice/` uses `https://voice.messagebird.com/v1`
- `conversation/` uses `https://conversations.messagebird.com/v1`
- `integration/` uses `https://integrations.messagebird.com`

**Error handling divergence**: The `voice` package has its own `ErrorResponse`/`Error` types (with `Code` + `Message` fields), registered via `messagebird.SetErrorReader()` in an `init()` function. All other packages use `messagebird.ErrorResponse` (with `Code` + `Description` + `Parameter` fields). Note: `SetErrorReader()` stores a single global `errorReader` — only one custom error reader can be active at a time (last-writer-wins).

**Request body encoding**: `prepareRequestBody` in `client.go` switches on type: `nil` sends no body, `string` sends as `application/x-www-form-urlencoded`, anything else is JSON-marshalled.

**Integration package API versions**: `integration/` declares `version = "v2"` as default, but `ListWhatsAppTemplates` in `integration/whatsapp.go` hardcodes `"v3"` inline. Be aware of this inconsistency when adding new integration endpoints.

**Pagination**: `api.go` defines `PaginationRequest` (with `Limit`/`Offset`) and `DefaultPagination`. Sub-packages that need pagination should embed this type.

**Signature validation**: `signature/` is deprecated. Use `signature_jwt/` (`NewValidator` + `ValidateSignature`) for webhook validation.

## Testing

Tests use a shared TLS test server from `internal/mbtest`. The pattern:

1. Each test package has `TestMain` calling `mbtest.EnableServer(m)` to start a fake HTTPS server.
2. Tests call `mbtest.WillReturnTestdata(t, "filename.json", statusCode)` to set canned responses.
3. Tests use `mbtest.Client(t)` to get a client wired to the fake server.
4. Assertions use `mbtest.AssertEndpointCalled(t, method, path)` and `mbtest.AssertTestdata`/`AssertTestdataJson` for request body validation.
5. Fixture files live in `<package>/testdata/*.json`.

For unit tests that don't need the TLS server, use `mbtest.MockClient()` (returns a `testify/mock`-based no-op client). `mbtest.HTTPTestTransport(handler)` is also available for tests needing a custom HTTP handler (e.g., signature validation tests).

Uses `github.com/stretchr/testify` for assertions.
