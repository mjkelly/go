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

You can specify as many templates as you want -- their output will be concatenated in the order they're provided. That means:
```
./buildfile -v Name="buildfile user" t1.tmpl t1.tmpl
```
outputs
```
Hello buildfile user
Hello buildfile user
```

You can also take input from a file, like this:
```
./buildfile -v Name@=namefile t1.tmpl
```

Here's an example
```
echo "my name" > namefile
./buildfile -v Name@=namefile t1.tmpl
```
output:
```
Hello my name

```

## Building

```
go build
```

It doesn't have dependencies outside the standard library.
