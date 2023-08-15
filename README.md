# Notify
Simple alert service that other applications can use to email/text in an event. Made to be used from as a cli tool, easy to call from other applications through standard input. It allows you to send text messages, attach files, and specify a custom smtp server, default being gmail.


# Notation
```golang
    //required flags
    nofity <username> <password> -to <sender>

    //optional flags
    notify <username> <password> -to <sender> -a <attachment> -smtp <server>
```
