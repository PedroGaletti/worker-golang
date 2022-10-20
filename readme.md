## Worker - Golang

**Warning**: This project is a basic worker in Golang

## Stack and Packages

- [Golang](https://go.dev) - Build fast, reliable, and efficient software at scale
- [Docker](https://www.docker.com) - Accelerate how you build, share, and run modern applications
- [Go Cron](https://github.com/go-co-op/gocron) - Gocron is a job scheduling package which lets you run Go functions at pre-determined intervals
- [Zero Log](https://github.com/rs/zerolog) - Zerolog package provides a fast and simple logger dedicated to JSON output

## Project structure

```
$PROJECT_ROOT
├── cmd              # Image files
│   ├── configs      # Configs of application
│   └── jobs         # Cron logic
└── main.go          # Entry point
```

## Environment Variables

```
Variable                         | Type    | Description                       | Default
-------------------------------- | ------- | --------------------------------- | --------------------------------------------------------------------
CRON_TIME                        | string  | Interval time of the cron         | */1 * * * *
LOCALIZATION                     | string  | Timezone of the cron              | America/Sao_Paulo
LOG_LEVEL                        | string  | Leveled Logging                   | info
```

## Make commands:

Assuming that you have already cloned the project and the [Go](https://golang.org/doc/install) is installed, ensure that all dependencies are vendored in the project:

`make install`

To build the application:

`make build`

To run the application local:

`make run`
