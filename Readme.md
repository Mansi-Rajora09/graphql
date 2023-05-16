# README

## Quick start
1. [Initialise a new go module](https://golang.org/doc/tutorial/create-module)

       mkdir example
       cd example
       go mod init example

2. Add `github.com/99designs/gqlgen` to your [project's tools.go](https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module)

       printf '// +build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go

       go mod tidy

3. Initialise gqlgen config and generate models

       go run github.com/99designs/gqlgen init

4. Start the graphql server

       go run server.go

More help to get started:
 - [Getting started tutorial](https://gqlgen.com/getting-started/) - a comprehensive guide to help you get started
 - [Real-world examples](https://github.com/99designs/gqlgen/tree/master/_examples) show how to create GraphQL applications
 - [Reference docs](https://pkg.go.dev/github.com/99designs/gqlgen) for the APIs


###  Some Useful Links & Commands

- Install graphql Library:
  - Command: `go install github.com/99designs/gqlgen@v0.17.25`
- See gqlgen version
  - Command `go run github.com/99designs/gqlgen@v0.17.25 version`
- Generate the schema:
  - Command: `go run github.com/99designs/gqlgen@v0.17.25 generate`. 


## Questions that might arise in your mind

### How do I prevent fetching child objects that might not be used?

When you have nested or recursive schema like this:

```graphql
type Todo {
  id: ID!
  text: String!
  done: Boolean!
  user: User!
}
```

You need to tell gqlgen that it should only fetch friends if the user requested it. There are two ways to do this;

- #### Using Custom Models

Write a custom model that omits the friends field:

```go
type User struct {
  ID int
  Name string
}
```

And reference the model in `gqlgen.yml`:

```yaml
# gqlgen.yml
models:
  User:
    model: github.com/you/pkg/model.User # go import path to the User struct above
```

- #### Using Explicit Resolvers

If you want to Keep using the generated model, mark the field as requiring a resolver explicitly in `gqlgen.yml` like this:

```yaml
# gqlgen.yml
models:
  Todo:
    fields:
      user:
        resolver: true # force a resolver to be generated
```

After doing either of the above and running generate we will need to provide a resolver for friends:

```go
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}
```

### Can I change the type of the ID from type String to Type Int?

Yes! You can by remapping it in config as seen below:

```yaml
models:
  ID: # The GraphQL type ID is backed by
    model:
      - github.com/99designs/gqlgen/graphql.IntID # a go integer
      - github.com/99designs/gqlgen/graphql.ID # or a go string
```

This means gqlgen will be able to automatically bind to strings or ints for models you have written yourself, but the
first model in this list is used as the default type and it will always be used when:

- Generating models based on schema
- As arguments in resolvers

There isn't any way around this, gqlgen has no way to know what you want in a given context.


### Screenshot for Reference

<img width="1790" alt="Screenshot 2023-05-16 at 5 27 17 PM" src="https://github.com/Mansi-Rajora09/graphql/assets/41550139/c6100718-3d73-4876-90fd-df1c9bb098e1">


<img width="1093" alt="Screenshot 2023-05-16 at 5 27 33 PM" src="https://github.com/Mansi-Rajora09/graphql/assets/41550139/fa8a56a0-8fbb-476e-9cba-ae39c5bdbd85">

<img width="1716" alt="Screenshot 2023-05-16 at 5 27 51 PM" src="https://github.com/Mansi-Rajora09/graphql/assets/41550139/197933da-2bd5-43b8-b49c-66cd48fb14c8">
