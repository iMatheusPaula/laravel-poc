# API App

Laravel application responsible for managing appointments and publishing domain events to Pub/Sub.

## Local scope

- App code lives entirely inside `apps/api`
- Container runtime files live in `apps/api/.docker`
- This app is intended to be buildable independently from the monorepo root
