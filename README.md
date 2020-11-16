# boilr-environ

This is a [boilr template](https://github.com/tmrts/boilr) for creating an environment variable extension. Use an environment variable extension to provide custom environment variables to your pipelines. Get started by cloning the project and installing the template:

```console
$ boilr template download drone/boilr-environ drone-environ
```

create a project in directory my-environ:

```console
$ boilr template use drone-environ my-environ
```

enter the docker registry name for this project:

```text
[?] Please choose a value for "DockerRepository" [default: "foo/bar"]:
```

enter the go module name:

```text
[?] Please choose a value for "GoModule" [default: "github.com/foo/bar":
```
