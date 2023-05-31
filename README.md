# DORA

Calculate DORA four-keys metrics for a git repository.

Simple CLI tool to capture a snapshot of the DORA four-keys metrics for a git repository.
You'll need to provide a personal access token with the necessary permissions to access the GitHub repository.
This project is not maintained to the DORA team. (https://dora.dev).
To know more about the four metrics, go to https://bit.ly/dora-fourkeys.


## Installation

```bash
go install github.com/yellow-canary/dora@latest
```

## Set up

```bash
export GH_TOKEN=<my-personal-access-token>
```

## Usage

```bash
dora calculate -r owner/repo
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[Apache](http://www.apache.org/licenses/LICENSE-2.0)