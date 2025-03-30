# Gator - Blog Aggregator CLI

## 📖 Description

**Gator** is a feature-rich Go-based command-line tool designed to aggregate, fetch, and store RSS feed data from various sources. It provides robust capabilities for:

- Adding and managing RSS feeds.
- Fetching and storing posts from feeds.
- Scheduling automated feed scraping using a ticker-based approach.
- Handling middleware for command processing.
- Providing a structured database schema using SQLC.

This tool serves as a backend utility for collecting and managing content from multiple blogs or news sources, efficiently storing them for later retrieval or further processing.

---

## 📁 Project Structure

```
blog-aggregator-go-master/
│
├── main.go                      # Entry point of the application
├── command_list.go              # Command handling for listing feeds
├── command_middleware.go        # Middleware system for command processing
├── command_system.go            # Core command system handling user input
├── handler_aggregator.go        # Aggregation command handler for periodic scraping
├── handler_feed.go              # Feed management handler (adding, removing feeds)
├── handler_feed_follow.go       # Handles the follow functionality for feeds
├── rss_command.go               # RSS fetching commands (parsing RSS feeds)
├── go.mod / go.sum              # Go modules files
├── sqlc.yaml                    # SQLC configuration file for generating Go code from SQL
├── internal/
│   ├── config/
│   │   └── config.go            # Configuration handling (e.g., loading .env variables)
│   └── database/
│       ├── db.go                # Database connection setup
│       ├── models.go            # Data models (SQLC generated)
│       ├── feeds.sql.go         # SQLC generated code for feeds operations
│       ├── posts.sql.go         # SQLC generated code for posts operations
│       └── feed_follows.sql.go  # SQLC generated code for feed follows
├── notes.txt                    # Development notes and ongoing tasks
├── .gitignore                   # Git ignore file
├── README.md                    # Project documentation (this file)
└── gator                        # Compiled binary (if applicable)
```

---

## ⚙️ Installation

### Installing as a CLI Application

1. **Clone the repository:**
```sh
git clone https://github.com/your-username/gator.git
```

2. **Navigate to the project directory:**
```sh
cd blog-aggregator-go-master
```

3. **Install dependencies:**
```sh
go mod download
```

4. **Build the application as a CLI executable:**
```sh
go build -o gator .
```

5. **Move the executable to your system PATH for global usage:**
```sh
mv gator /usr/local/bin/
```

6. **Verify installation:**
```sh
gator --help
```

7. **Setup environment variables:**
Create a `.env` file with necessary configuration, e.g.,
```
DB_URL=postgresql://user:password@localhost:5432/database
```

8. **Generate SQLC code (if needed):**
```sh
sqlc generate
```

---

## 🚀 Usage

### Authentication / Configuration
```
gator login
```

### Adding a Feed
```
gator addfeed <URL>
```

### Listing Feeds
```
gator list
```

### Starting Aggregation
```
gator agg <duration>  # e.g., gator agg 10s, gator agg 1m, gator agg 1h
```

---

## 🛠️ Technologies Used

- **Go** (Golang) - Core language.
- **SQLC** - For generating type-safe database interactions.
- **PostgreSQL** - Primary database for storing feeds and posts.
- **godotenv** - Environment variable management.

---

## 📌 Features

- Efficient RSS feed fetching and parsing.
- Scheduled scraping using Go's `time.Ticker`.
- Support for multiple feeds.
- Middleware pattern implementation for extensibility.
- Type-safe SQL queries using SQLC.
- Error handling and graceful shutdown support.
- Improved CLI experience with subcommands (login, addfeed, list, agg).

---

## 🔒 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.