# golang-test-assignment

You must design a CLI tool. This CLI tool waits for CSV files with a huge amount of data (this data not ordered). This data has the next columns:
*orderID, date, userID, amount, currency*.

This tool must do next:
1) Parse file
2) Filter data (should support different filters). The current assignment should support filter by dates (should making aggregation only for the current year)
3) Calculate fees. The default rate for each day 3.2 (multiplier). Each Friday we have an additional multiplier 1.6. Each Monday we have an additional multiplier 0.7
3) Aggregate total amount by userID and currency
4) Output data with CSV file, with next columns:
*userID, amount, currency, fee*
5) Test coverage at least 55%
6) Code should pass all linters (`make lint`)
7) For delivery, make your own repository in any system (github/bitbucket/gitlab)
8) Add a description about your implementation

## Examples

This assignment contains examples generator. 
For configure them, use environment variables. For example, for generate 10000 examples, go into the cmd/generator folder and execute:
```
EXAMPLE_COUNT=10000 EXAMPLE_OUTPUT=example.csv go run main.go
```