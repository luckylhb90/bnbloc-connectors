# bnbloc-connectors

Sample connectors for [bnbloc-connections](https://github.com/luckylhb90/bnbloc-connections):
each file is one system as an organization's admin adds it on
**Components → Connectors** (paste the JSON, set its client values, enable
it).

| File | System | Connect |
|---|---|---|
| `connectors/github.json` | GitHub REST and GraphQL | a fine-grained personal access token |
| `connectors/github-oauth.json` | GitHub REST and GraphQL | an OAuth app (set `clientId`; the client secret on the admin page) |
| `connectors/google-calendar.json` | Google Calendar API v3 | Google OAuth (set `clientId`; the client secret on the admin page) |

OAuth apps register the redirect URI
`<your bnbloc origin>/v1/component-callbacks/connections/oauth`.

Every sample is checked against the connector rules of bnbloc-connections
(`go test ./...`). An agent drafts new ones itself: asked to connect a
system with an HTTP API, it reads the API's description and posts a draft
(`.room/connections/DRAFTING.md`), which an admin reviews on the same page.
