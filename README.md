# country_population
quick go server app to look up population by country

## Setup
1) Ensure your go environment path variables are set
2) run `go get github.com/mgsterling11/country_population`
3) cd into src/github/mgsterling11/country_population
4) run `go build`
5) run `./country_population` to start server
6) in another terminal window, run `curl http://localhost:8080/test-server` to confirm the server is running on port 8080
7) `curl http://localhost:8080/country/COUNTRY NAME` to find current populations for any country. For example, `curl http://localhost:8080/country/japan` or `curl http://localhost:8080/country/france`. For countries with more than one word in the name, use `_` to indicate spaces: `curl http://localhost:8080/country/united_kingdom` or `curl http://localhost:8080/country/united_arab_emirates`
