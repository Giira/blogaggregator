# RSS Feed Aggregator - Gator

## Prerequisites
- Go 1.23.3 or later
- PostgreSQL database
- A Unix style environment

## Installation
1. Clone the repository:
```
git clone https://github.com/Giira/blogaggregator
cd blogaggregator
```
2. Install dependencies:
```
go mod download
```
3. Setup the PostgreSQL database:
- Install PostgreSQL
- Create a new database
- Perform the database migration:
```
psql -d your_db -f sql/schema/001_users.sql
```
4. Setup configuration file:
- Create .gatorconfig.json in your home directory
- It should contain:
```
{
    "db_url": "postgres://username:password@localhost:5432/your_db"
}
```
5. Build:
```
go build
```

## Usage
### Users
```
. register <username>   # Create new user
. login <username>      # Login as user
. users                 # List registered users
. reset                 # Delete all users
```
### Feeds
```
. addfeed <name> <url>  # Add new RSS feed
. feeds                 # List feeds in database
. follow <url>          # Follow a feed as logged in user
. unfollow <url>        # Unfollow a feed
. following             # List logged in users followed feeds
```
### Functions
```
. browse [n]            # Show n latest posts - default 2
. agg <interval>        # Aggregate feeds - Make interval long enough to avoid DOSing feeds