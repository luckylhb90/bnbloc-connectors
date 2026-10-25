# bnbloc-connectors

Sample connectors for [bnbloc-connections](https://github.com/luckylhb90/bnbloc-connections):
each file is one system as an organization's admin adds it on
**Components → Connectors** (paste the JSON, set its client values, enable
it).

| File | System | Connect |
|---|---|---|
| `connectors/github.json` | GitHub REST and GraphQL | a fine-grained personal access token |
| `connectors/github-oauth.json` | GitHub REST and GraphQL | an OAuth app (set `clientId`; the client secret on the admin page) |
| `connectors/gitlab.json` | GitLab REST v4 (gitlab.com) | a personal access token |
| `connectors/notion.json` | Notion pages and databases | an internal integration's secret |
| `connectors/linear.json` | Linear GraphQL | a personal API key |
| `connectors/figma.json` | Figma files and comments | a personal access token |
| `connectors/okx-demo.json` | OKX demo trading (simulated funds) | an API key, its secret and passphrase; requests signed in headers, `x-simulated-trading: 1` on each |
| `connectors/bybit-testnet.json` | Bybit testnet (test funds) | an API key and its secret; requests signed in headers |
| `connectors/google.json` | Google Calendar, Drive, Docs, Sheets, Slides and Gmail | Google OAuth (set `clientId`; the client secret on the admin page) |

Google's APIs share `www.googleapis.com`, and a host belongs to one
connector, so one Google connector covers Calendar, Drive, Docs, Sheets,
Slides and Gmail: one sign-in for all of them.

OAuth apps register the redirect URI
`<your bnbloc origin>/v1/component-callbacks/connections/oauth`.

## Google OAuth client

1. Google Cloud console → **APIs & Services → Library**: enable the Google
   Calendar, Drive, Docs, Sheets, Slides and Gmail APIs.
2. **OAuth consent screen**: user type *Internal* when everyone signs in
   with your Google Workspace; otherwise *External* in *Testing*, with each
   member's Google account added as a test user (Gmail and Drive read are
   restricted scopes; an unverified app shows a warning before consent).
   Add the scopes from `connectors/google.json`.
3. **Credentials → Create credentials → OAuth client ID**, type *Web
   application*, authorized redirect URI
   `<your bnbloc origin>/v1/component-callbacks/connections/oauth`.
4. On **Components → Connectors**, paste `connectors/google.json` with its
   `clientId` set, enter the client secret, enable it.

The connections service reaches `accounts.google.com`,
`oauth2.googleapis.com` and the connector's hosts (`www`, `gmail`, `docs`,
`sheets` and `slides.googleapis.com`)
(the sign-in, the token refreshes and, unless the connector sends from the
members' computers, the calls).

Every sample is checked against the connector rules of bnbloc-connections
(`go test ./...`). An agent drafts new ones itself: asked to connect a
system with an HTTP API, it reads the API's description and posts a draft
(`.room/connections/DRAFTING.md`), which an admin reviews on the same page.
