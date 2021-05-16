# URL Shortener w/ Database

This is a URL Shortener written in Go. The mapping between URLs and paths are stored in a mySQL database.

The mapping can be add using a HTTP GET request, with url and path as a query parameters.

# Usage

1. Run `docker-compose up` in project directory
2. Add mapping to database by going to `http://localhost:8000/url?url=sampleURL&path=/samplePath`
3. Use the new mapping by going to `http://localhost:8000/samplePath`
