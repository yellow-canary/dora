# FourKeys CLI

FourKeys is a command-line interface (CLI) application written in Go that calculates a snapshot of the DORA four keys metrics of a GitHub repository. This tool helps you assess the performance of your software development and delivery process by calculating the following metrics:
* Deployment Frequency, 
* Lead Time to Change, 
* Change Failure Rate, and 
* Time to Restore

You'll need to provide a personal access token with the necessary permissions to access the GitHub repository.

This project is not maintained to the DORA team. (https://dora.dev).
To know more about the four metrics, go to https://bit.ly/dora-fourkeys.

## Installation

To use FourKeys CLI, you need to have Go installed on your machine. If you don't have Go installed, you can download it from the official Go website: https://golang.org/dl/

Once Go is installed, you can install FourKeys CLI by running the following command:
```bash
go install github.com/yellow-canary/fourkeys@latest
```

## Usage

To calculate the 4 DORA metrics for a GitHub repository, follow these steps:

1. Generate a personal access token from GitHub. You can create a token by going to your GitHub account settings -> Developer Settings -> Personal access tokens.
2. Export the generated access token as an environment variable:
```bash
export GH_TOKEN=<my-personal-access-token>
```

3. Run the FourKeys CLI command, providing the GitHub repository URL as an argument:

```bash
fourkeys calculate -r owner/repo
```
4. FourKeys will retrieve the necessary data from the GitHub API and perform the calculations. Once completed, it will display the calculated DORA metrics for the given repository.

```
Deployment Frequency: 2.1 releases per week
Lead Time to Change: 17.5 hours
Change Failure Rate: 9.4%
Time to Restore Service: 0.5 hours
```
## Understanding the four key metrics

The DORA Four Keys metrics are calculated based on the data obtained from a GitHub repository. Here's an explanation of how each metric is calculated:

### Deployment Frequency:
Deployment Frequency measures how often code changes are deployed to production. A new deployment is measured by a new release tag in Github.

### Lead Time to Change:
Lead Time to Change measures the time it takes for a code change to go from the start to being deployed. A change is measured from the moment a Pull Request is open to it's conclusion.

### Change Failure Rate: 
Change Failure Rate measures the percentage of deployments that result in a incident, failure or require remediation (e.g., rollbacks, hotfixes, etc.). A incident is measured by an Github Issue with the label 'bug'. 

### Time to Restore Services
Time to Restore measures the time it takes to recover from a failure or incident. It is calculated by finding the time to resolve an incident (closing a Github Issue with the label 'bug'). 

## Contributing

Contributions to FourKeys CLI are welcome! If you encounter any issues or have suggestions for improvements, please open an issue first to discuss what you would like to change.

If you'd like to contribute code, you can fork the repository, create a new branch, make your changes, and submit a pull request. Please ensure that your code follows the established coding style and includes or update tests as appropriate.

## License

[Apache](http://www.apache.org/licenses/LICENSE-2.0)

Thank you for using FourKeys CLI! If you have any questions or need further assistance, please don't hesitate to reach out.