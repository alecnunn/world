# World Generator

[![Maintainability](https://api.codeclimate.com/v1/badges/47856d43d023e779baa7/maintainability)](https://codeclimate.com/github/ironarachne/world/maintainability) [![Test Coverage](https://api.codeclimate.com/v1/badges/47856d43d023e779baa7/test_coverage)](https://codeclimate.com/github/ironarachne/world/test_coverage) [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ironarachne_world&metric=alert_status)](https://sonarcloud.io/dashboard?id=ironarachne_world)

This is a tool and library to procedurally generate fantasy worlds.

The API runs on port 7531.

## Configuration

The following environment variables configure the API and are required:

* `WORLDAPI_DATA_DIRECTORY`: The absolute directory that contains data files for the application.
* `WORLDAPI_SAVE_DIRECTORY`: The absolute directory where files will be written to.
* `WORLDAPI_SAVE_TARGET`: Either `DO` to save to Digital Ocean or `filesystem` to save to the local file system.
* `WORLDAPI_WEB_DOMAIN`: The domain name (without protocol) the API serves on. For example, `www.worldapi.com`.

## Building

Run `make build`. The binary will be `build/worldapi`.

## Running in Development

Run `make run`. It will set sane defaults and use the local filesystem for saving files.
