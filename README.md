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
| `connectors/google.json` | Google Calendar, Drive and Gmail | Google OAuth (set `clientId`; the client secret on the admin page) |

Google's APIs share `www.googleapis.com`, and a host belongs to one
connector, so one Google connector covers Calendar, Drive and Gmail.

OAuth apps register the redirect URI
`<your bnbloc origin>/v1/component-callbacks/connections/oauth`.

Every sample is checked against the connector rules of bnbloc-connections
(`go test ./...`). An agent drafts new ones itself: asked to connect a
system with an HTTP API, it reads the API's description and posts a draft
(`.room/connections/DRAFTING.md`), which an admin reviews on the same page.
