# Blog Aggregator CLI - Go Application

## 📖 Description

The **Blog Aggregator CLI** is a Go-based command-line tool designed to aggregate, fetch, and store RSS feed data from various sources. It provides capabilities for:

- Adding and managing RSS feeds.
- Fetching and storing posts from feeds.
- Scheduling automated feed scraping using a ticker-based approach.
- Providing middleware and command parsing for extensibility.

This tool is intended to be a backend utility for collecting and managing content from multiple blogs or news sources.

---

## 📁 Project Structure

```
blog-aggregator-go-master/
│
├── main.go                      # Entry point of the application
├── command_list.go              # Command handling for listing feeds
├── command_middleware.go        # Middleware system for command processing
├── command_system.go            # Core command system handling user input
├── handler_aggregator.go        # Aggregation command handler
├── handler_feed.go              # Feed management handler (add, remove feeds)
├── handler_feed_follow.go       # Feed following logic
├── rss_command.go               # RSS fetching commands
├── go.mod / go.sum              # Go modules files
├── sqlc.yaml                    # SQLC configuration file for generating Go code from SQL
├── internal/
│   ├── config/
│   │   └── config.go            # Configuration handling (e.g., loading .env variables)
│   └── database/
│       ├── db.go                # Database connection setup
│       ├── models.go            # Data models
│       ├── feeds.sql.go         # SQLC generated code for feeds
│       ├── posts.sql.go         # SQLC generated code for posts
│       └── feed_follows.sql.go  # SQLC generated code for feed follows
├── notes.txt                    # Development notes
└── .gitignore                   # Git ignore file
```

---

## ⚙️ Installation

1. **Clone the repository:**
```sh
git clone https://github.com/your-username/blog-aggregator-go.git
```

2. **Navigate to the project directory:**
```sh
cd blog-aggregator-go-master
```

3. **Install dependencies:**
```sh
go mod download
```

4. **Setup environment variables:**
Create a `.env` file with necessary configuration, e.g.,
```
DB_URL=postgresql://user:password@localhost:5432/database
```

5. **Generate SQLC code (if needed):**
```sh
sqlc generate
```

6. **Run the application:**
```sh
go run main.go
```

---

## 🚀 Usage

### Adding a Feed
```
/rss <URL>
```

### Listing Feeds
```
/list
```

### Starting Aggregation
```
/agg <duration>  # e.g., /agg 10s, /agg 1m, /agg 1h
```

---

## 🛠️ Technologies Used

- **Go** (Golang)
- **SQLC** (for generating type-safe database queries)
- **PostgreSQL** (Database)
- **godotenv** (for environment variable management)

---

## 🔒 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

---

## 📌 Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
