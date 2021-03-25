# Board Game Top Five (bgtf)

## Core features

### Search for top five board game videos associated with board game text search

- Fuzzy search
- Get top five relevant results based off of later defined criteria
- Almost like an RSS for bg youtube?

### Simple API

- Interact with data in useful way without having to aggregate or search manually
- Useful for multiple applications:
  - CLI
  - React App frontend

### Simple interface

- Previews of videos
- Links directly to videos
  - description
  - likes, dislikes
- Ads? (monetization)
- Accounts, basic web app features

## Tech

### Frontend

- React App components
- TypeScript

### Backend

- Go API interacting with YouTube API

#### Calls

- `GET:/api/search`

  - Search for board game by name or other criteria specified in headers

- `GET:/api/top_five`

  - Return top five popular board games with individual calls to `/api/search`

- `POST:/api/write`

  - Write user data to Postgres
    - I.e. favorites to follow

### Commandline (optional)

Interact with bgtf through commandline. CLI will make calls to API server
running locally (or in Cloud Run).

- Simple evoke:
  - `bg search oath`
    - Search for board game by name. Returns top five youtube results for that video
  - `bg hot`
    - Get top five most popular games coming out based on specified criteria

## Optional features

- Cloud Run
- CLI
