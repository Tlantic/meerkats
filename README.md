#Meerkats

## Install
``sh
    go get github.com/Tlantic/meerkats
``

## Usage

### Init
``go
    sentry := NewMeerkat(MeerkatOptions{})
    defer sentry.Close()
``

### Register Handlers
``go
    out := handlers.NewWritterLogger( os.Stdout )
    sentry.RegisterHandler(LEVEL_ALL, out.HandleEntry)
``

### Logging
``go
    sentry.WithField("hello", "world").WithFields(Fields{
    	"make": "better",
    }).Error("Oops something went wrong.")
``