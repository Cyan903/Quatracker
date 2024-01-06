<p align="center">
    <picture>
        <source media="(prefers-color-scheme: dark)" srcset="/assets/logo/docs/logo-light.svg" />
        <source media="(prefers-color-scheme: light)" srcset="/assets/logo/docs/logo-dark.svg" />
        <img height="100" src="/assets/logo/logo-1.svg" />
    </picture>
</p>

<p align="center">
    <img src="https://img.shields.io/github/go-mod/go-version/cyan903/quatracker?style=for-the-badge" />
    <img src="https://goreportcard.com/badge/github.com/cyan903/quatracker?style=for-the-badge" />
    <img src="https://img.shields.io/github/package-json/v/cyan903/quatracker?filename=frontend%2Fpackage.json&style=for-the-badge" />
    <img src="https://img.shields.io/github/license/cyan903/quatracker?style=for-the-badge" />
</p>

<hr />

A cross-platform score tracker for the rhythm game [Quaver](https://quavergame.com/). Works as a convenient way to calculate, monitor, and track local scores. Supports both key modes along with unlimited local profiles.

## Features

- Advanced score filter
- Play history graph
- Judgement statistics
- Map statistics
- Overall statistics
- Grade tracker
- Unknown profiles
- Custom filters

## Building

Quatracker depends on the [wails](https://wails.io/docs/reference/cli/) CLI to compile.

```sh
$ git clone https://github.com/Cyan903/Quatracker.git
$ cd Quatracker
$ make update
$ make build
```

## Development

Quatracker depends on [wails](https://wails.io/) and [webkitgtk](https://webkitgtk.org/) as development dependencies.

```sh
$ make dev # run in development mode
$ make format # format and lint code
$ make update # update and validate dependencies
```

## License

[MIT](LICENSE)
