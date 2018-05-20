#Go Unit and Integration Test Seperation

### Option 1: Integration and Unit tests in the Same File

Inside each integration test
```golang
	if testing.Short() {
		t.Skip("Skipping itegration test")
	}

```
Using the `-short` flag, only the unit test will run:

`go test ./... -v  -short`


Without it, both Unit and the Integration test will run.


To execute only the integration test, the tests are named ending in Integration so that the `-run` flag can be used to execute only the tests ending with `Integration` 

`go test ./... -v -run Integration`

### Option 2: Unit and Integration in Separate Files.

Using the build tags

```golang
// +build integration
```

and the `-run` flag the integration test can be defined in another file and does not have to be skipped.  This time the test function ends in IT so that it is easy to determine which tst was run

`go test ./... -v -tags integration -run IT`

