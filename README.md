# GatorRSS
An RSS Aggregator Written in Go
## Dependencies
Postgresql and Go is required to run this program
## Installation
Install gator through following command:
```bash
go install GatoRSS
```
## SetUp
For the SetUp Create a .gatorconfig.json with your postgres connectionstring like this:
```json
{
    "db_url": "connection_string_goes here"
}
```
## Commands
to register a user:
```bash
GatorRSS register <username>
```
Add a feed:
```bash
GatorRSS addfeed <name> <url>
```
To fetch topics for your feeds use the command agg. It will continuously fetch topics in the specified time interval:
```bash
GatorRSS agg <time_interval>
```
To Browse through the latest posts use browse, if no number of posts is specified, the 2 latest posts will be displayed:
```bash
GatorRSS browse <number_of_posts_to_display>
```
