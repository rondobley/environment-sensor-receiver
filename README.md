# Environment Sensor Receiver
This is part of a project where I used a couple of BMP280 environmental sensors to capture
temperature data, etc. The sensors send JSON messages to this receiver, which then
processes them and inserts them into a TimeScaleDb instance.

This is not meant to be a real world production ready app. It is something I worked on as
a personal project so I can learn and get better developing projects on Go.

## Install

Create a config file, see [config-example.json](config-example.json). If running locally,
place the config file in the root dir of the project, if running in production, place the
file in `/usr/local/etc/environment-sensor-recevier`.

## Options
`-env` to set the environment to either `prod` for production or `local` for local dev.
