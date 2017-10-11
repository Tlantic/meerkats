<p align="center"><img src="http://icons.iconarchive.com/icons/shwz/disney/256/timon-icon.png" width="200px" alt="Meerkat"
         title="Meerkats Logger" align="left"/></p>

**Meerkats** <sub>Structured Logger</sub>
===
High performance, structured and flexible logging. Leverage the destination of your logs using context
and handler levels setting up different ones for each handlers (i.e. write to *stdout* only when tracing
and to *syslog* only when an error occurs)<br>

<br><br><br>
Install
---
Grab it.
```bash
go get -u github.com/Tlantic/meerkats
```

> **Note:**
> Due to the continue development of meerkats please consider using a package manager that supports semantic versioning like [Glide](https://github.com/Masterminds/glide).


<br>
Basic Usage
---
### Handlers
By default meerkats won't initialize any handler and therefore won't event write to stdout, nevertheless a basic [Writer Handler](https://github.com/Tlantic/meerkats/tree/master/handlers/writer) subpackage
is shipped with it that implements io.Writer allowing you to write entries to stdout or any other file.

In order to register a handler just use the Register() logger method.
```Go
  handler := writer.New()
  kat.Register(handler)
```

### Levels
The [Level](https://github.com/Tlantic/meerkats/blob/master/levels.go#L5) type serves a double purpose. It is used to instruct meerkats from which level should an entry be considered and passed down to its handlers.
```Go
  kat.SetLevel(kat.INFO)
  
  kat.Trace("trace entry")      // Ignored even if a handler subscribed to such level
  kat.Info("info entry")        // Dispatched to handlers
  kat.Warning("warning entry")  // Dispatched to handlers
```
It can also be used to tell handlers what entries will they dispatch by using bitwise ops.
```Go
  handler := writer.New()
  handler.SetLevel(kat.DEBUG|kat.WARNING|kat.ERROR)
  kat.Register(handler)
  
  kat.Debug("debug entry")  // Dispatched to handler
  kat.Info("info entry")    // Ignored
  kat.Error("error entry")  // Dispatched to handler
```

### Fields
Meerkats took inspiration from [Zap](https://github.com/uber-go/zap) using a more verbose method of [adding fields](https://github.com/Tlantic/meerkats/blob/master/field.go#L122) to log entries delegating the efforts of field type assertions to the developer
instead of using reflection methods during runtime. Face it, most of the fields type is known prior execution so theres no excuse to hurt performance.
```go
import (
  rand  
  kat "github.com/Tlantic/meerkats"
  "github.com/Tlantic/meerkats/handlers/writer"
)

func main() {
  kat.Trace("Init message", kat.String("foo", "bar"), kat.Int("random", rand.Int()) )
}
```

### Contexts
Logs written while bootstrapping are somehow different that those written within a HTTP request handler.
With this in mind meerkats allows you to contextualize logger instances, inheriting fields and handlers from a parent context.

```go
package main

import (
  "io"
  "net/http"
  
  kat "github.com/Tlantic/meerkats"
  "github.com/Tlantic/meerkats/handlers/writer"
)



func main() {

  // initialize root
  kat.Register(writer.New())
  kat.EmitString("service-name", "svc-test")
  
  err := http.ListenAndServe(":12345", func (w http.ResponseWriter, req *http.Request) {
    
    log := kat.Clone() // create new logger context
    log.EmitString("Origin", req.Header.Get("Origin"))
    //...
    
    io.WriteString(w, "Ok")
  }))
  panic(err)
}
```


<br>

Standard Go Log Package
---

### Interface
Some packages may not want to directly depend on meerkats to use a logger or wan't a more generic interface to interact with.
To satisfy such need Meerkats can be wrapped within a [StandardLogger](https://github.com/Tlantic/meerkats/blob/master/logger.go#L32) instance.
```Go
  logger := kat.New(writer.New(kat.Level(kat.INFO|kat.ERROR|kat.FATAL|kat.PANIC)))
  
  std := kat.NewStandardLogger(logger, kat.INFO)
```

### Intercept
Since the logger implements the io.Writer interface meerkats can be used as the standard go log package output.
```Go
  logger := kat.New()
  log.SetOutput(logger)
```


