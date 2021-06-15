# outerloop-devfile-sample

Sample parent and child devfiles showing variable substitution

#### Flattening devfiles

The `flatten` go binary can be added to your go/bin directory and used to "flatten" a devfile (or child devfile) and render a parsed and validated version. 

Go source files are available in the `flatten-devfile` dir. Run `go install` from inside that dir to rebuild the `flatten` binary in go/bin. 