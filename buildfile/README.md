# Buildfile

This is a utility to help in building text artifacts. It's designed to be
called from a Makefile.

## Examples

An invocation like this:
```
./buildfile -v Name="buildfile user" t1.tmpl 
```

Reads t1.tmpl and provides the argument `Name=buildfile user`.

If this is the content of `t1.tmpl`:
```
$ cat t1.tmpl 
Hello {{.Name}}
```

Then the output is:
```
Hello buildfile user
```

## Building

```
go build
```

It doesn't have dependencies outside the standard library.
