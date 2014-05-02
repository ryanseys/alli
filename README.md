# alli

Lists all open issues in your GitHub repositories

## Usage

1. (Optional **) Set GitHub personal authentication token ([create one here](https://github.com/settings/applications)) as `GH_TOKEN_ALLI` in your environment variables.
2. Build and run: `go run alli.go`
3. Enter in username (and optionally save username to `~/.alli` for quicker execution later)

** Optional because you may continue unauthenticated but will be rate limited much more heavily by GitHub.

### Example output
```bash
# ~/alli [git:master o] $ go run alli.go
Yay! Using authentication token!
Using saved username: ryanseys

ryanseys/alli
#5 Add explanation of this repo
#1 Use go-github library
```

## License

MIT &copy; Ryan Seys
