# jet-helper

[Jet](https://github.com/CloudyKit/jet) is a template engine developed to be easy to use, powerful, dynamic, yet secure and very fast.
`jet-helper` is a package to utilize Jet similarly to Golang text/template and html/template.

## Usage Sample

```golang
import (
     tmpl "github.com/rosbit/jet-helper"
)

// sample 1
tmpls := tmpl.NewTempls(".")
t, err := tmpls.Lookup("index.jet")
if err != nil {
    // error processing
}
// t.Funcs(map[string]interface{}{})   // add Funcs like html/template::FuncMap
t.Execute(io.Writer, any)

// sample 2 to Parse text template
t, err := tmpl.Parse("text template as a string")
...
```
