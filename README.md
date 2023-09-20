# inline-python-mod

Dagger module for inlining Python code in a pipeline

## Requirements

You need to run a [Dagger with module support](https://github.com/shykes/dagger/tree/zenith-functions/zenith#project-zenith)

## Examples

### Get time and time on my local timezone

```sh
dagger query --progress=plain -m "github.com/samalba/inline-python-mod@main" date < examples.gql
```

### Get my public IP

```sh
dagger query --progress=plain -m "github.com/samalba/inline-python-mod@main" myIP < examples.gql
```
