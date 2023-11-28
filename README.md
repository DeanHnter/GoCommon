# GoCommon
This is my personal project for common reusable functions, data-structures etc I find useful for multiple projects. Its not a consistent framework etc, just a bunch of things that I have implemented to solve problems.

## OS
### Environment Variables
```go
	envs := gocmnos.ReadEnvironmentVariables()
	if _, exists := envs.Exists("ENV_VARIABLE"); !exists {
		panic("Environment variable ENV_VARIABLE doesent exist")
	}
```
