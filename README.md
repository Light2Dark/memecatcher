# Meme Fetcher

A website that fetches memes from Reddit, searches through them to match your query and displays them to you.

## Installation

- Clone the repository
- Install Go and install the dependencies using `go mod download`
- Create a `.env` file in the root of the project from the `.env.example` file and replace the values with your own
- Run the server using `go run ./cmd/api`

### Tips to develop locally

- Hot reloading (https://templ.guide/commands-and-tools/hot-reload)
- Running tailwind watcher (https://tailwindcss.com/blog/standalone-cli)
- Disable the cache when developing (developer tools -> network -> tick 'disable cache' checkbox)

## Technologies Used

- HTMX: Frontend library for making AJAX requests
- Tailwind: For styling
- Go: All of the logic is written in Go, responses are sent back in HTML (using Templ)
- AlpineJS: For minor client side interactivity
- Neon.tech: Cloud Postgres DB

## TODO

- Logging is not working correctly, errors are not being printed
- I wanted to add explanation field in the db schema, but it will be costly to do so
