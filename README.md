# httptool

It's anooying that golang http.Client add cookie and header
You need to add your header for each request
```go
  client := http.Client{}
  req, err := http.NewRequest(method, url, nil)
  req.Header.Set(YOUR_HEADER)
  req.AddCookie(&http.Cookie{Name: name, Value: value)
```

Now you can write easier
httptool will automatic process
```go
  client := httptool.New()
  client.SetHeader("User-Agent", Fake_User_Agent")
  client.SetCookie(Name, Value)
  req, err := client.Get(url, method)
```
