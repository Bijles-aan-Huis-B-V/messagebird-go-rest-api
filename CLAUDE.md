@../bah-knowledge/TECHNICAL.md
@../bah-knowledge/PRODUCT.md
@../bah-knowledge/DEVELOPMENT.md

# CLAUDE.md

This file provides guidance to Claude Code when working in this repository. Platform-wide patterns (service
layout, communication, infrastructure, deployment) live in `bah-knowledge/TECHNICAL.md`; this file covers only
what is specific to messagebird-go-rest-api.

## Status

**To be replaced.** Fork of the unmaintained upstream MessageBird Go SDK. Do not add features; new MessageBird
calls belong in a small client inside the consuming service. It cannot be archived while notification-service and
hubspot-service import it. `README.md` has the replacement checklist.

## Project Overview

- Module `github.com/Bijles-aan-Huis-B-V/messagebird-go-rest-api`, `go 1.16`. Library only, no deploy, no tags.
- Forked from upstream v9.1.0 (2022-07). Fork-only commits: `d5316d0` module rename, `f07ae79` + `2a4b5c3`
  Contacts v2 API and `contact.Upsert`, `c67565d` + `23679fe` `Secret` type and `integration/` (WhatsApp
  templates), `421306c` (2024-02-08) list filters as query params.
- Consumers pin `v0.0.0-20240208154711-421306c3a4d4` (= newest code commit):
  - notification-service (`internal/messageBird/service.go`, `internal/whatsapp/service.go`,
    `cmd/notification-server/main.go`): `messagebird.New`, `messagebird.Secret`, `conversation.SendMessage`
    with `conversation.HSM` content, `integration.ListWhatsAppTemplates`, `integration.WhatsAppComponent*`.
  - hubspot-service (`internal/messagebird/service.go`, `cmd/server/server.go`): `messagebird.New`,
    `contact.Upsert`, `contact.CreateRequest`, `contact.Identifier`.
  - Nothing else in the platform imports `sms`, `voice`, `mms`, `verify`, `lookup`, `hlr`, `number`, `group`,
    `balance`, `partner_accounts`, `voicemessage` or the signature packages.

## Build / Run / Test

```bash
go build ./...
go test ./...
go test ./sms/ -run TestCreateMessage
```

CI: `.github/workflows/tests.yml` (upstream's), `go test ./...` on every push/PR, Go 1.16.x / 1.17.x / 1.18.x,
`actions/checkout@v2`, `actions/setup-go@v2`. `vendor/` is git-ignored.

## Layout and client model

- `client.go`: `Client` interface with one method, `Request(v, method, path, data)`; `DefaultClient` (`New(key)`)
  adds `Authorization: AccessKey ...`, JSON decoding and error mapping. A path that does not start with
  `http(s)://` is prefixed with `Endpoint = "https://rest.messagebird.com"`.
- `prepareRequestBody`: `nil` = no body, `string` = form-encoded, anything else = JSON.
- API roots used by packages: `conversation/` `https://conversations.messagebird.com/v1`, `voice/`
  `https://voice.messagebird.com/v1`, `integration/` `https://integrations.messagebird.com`, `contact/`
  `https://contacts.messagebird.com/v2/contacts` (+ `/v2/ops/contacts` for upsert).
- `integration/api.go` declares `version = "v2"`, but `ListWhatsAppTemplates` builds `/v3/platforms/whatsapp/templates`
  inline.
- `error.go`: `ErrorResponse` (code, description, parameter). `SetErrorReader` stores one global custom reader;
  both `voice/` and `partner_accounts/` register one in `init()`, so importing both means last one wins.
- `api.go`: `PaginationRequest` + `DefaultPagination`.
- `secret.go`: `Secret{Key string; Channels map[string]string}`; notification-service unmarshals the
  Secrets Manager secret named by its `MESSAGE_BIRD` env var into `map[country]Secret`. hubspot-service uses its
  own config type for the same env var.
- `signature/` is deprecated upstream; `signature_jwt/` is the replacement.

## Tests (`internal/mbtest`)

1. `TestMain` calls `mbtest.EnableServer(m)` (fake TLS server).
2. `mbtest.WillReturnTestdata(t, "file.json", status)` sets the canned response; fixtures live in
   `<package>/testdata/`.
3. `mbtest.Client(t)` returns a client pointed at the fake server.
4. Assert with `mbtest.AssertEndpointCalled`, `AssertTestdata`, `AssertTestdataJson`.
5. `mbtest.MockClient()` gives a testify mock; `mbtest.HTTPTestTransport(handler)` for custom handlers.

## Replacing it (when asked)

1. In the consumer, write a minimal client for only the calls listed above (same JSON shapes; copy the request
   and response structs you need).
2. Keep the Secrets Manager secret format (`MESSAGE_BIRD`) unless the infra change is planned with it.
3. Remove the `require` line, run `go mod tidy`, deploy the consumer.
4. When neither consumer imports this module any more, archive the repo.

## Gotchas

- A merge here has no effect until a consumer bumps its pseudo-version.
- `go 1.16` in `go.mod` and a 1.16-1.18 CI matrix: CI does not test the Go versions the consumers build with.
- `UPGRADING.md` is upstream's upgrade guide for SDK major versions and does not describe this fork.
