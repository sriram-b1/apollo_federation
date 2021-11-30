# GQLGEN Apollo Federation

## How to run

- Run `yarn` or `npm install`, which would install required dependencies
- Run the `server.go` files in all the graphql services
For example, to run the __Metro__ service we have to do like this

```bash
$cd Metro
$go run server.go
```

- Run `node Gateway/index`, which would launch the federation server
- Headover to [http://localhost:4000](http://localhost:4000)
