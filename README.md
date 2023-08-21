To Run,
in seperate terminals cd into `api` and `gen` and run `go run main.go` then in the third terminal curl the first address (curl http://localhost:8080) and this generates the token you then curl another e.g(curl http://localhost:9001 --header 'Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJiaWxsaW5nLmp3dGdvLmlvIiwiYXV0aG9yaXplZCI6dHJ1ZSwiY2xpZW50Ijoia29kZWZvcmNlIiwiZXhwIjoxNjkyNTk4MzA0LCJpc3MiOiJqd3Rnby5pbyJ9.bBgjkeMYAdosB3Q8jvk5x16I-4KKjLZzJrV-HQ1faP4')

There's only one minute from generating to checking although that can be modified in the code to your liking.

Cheers! to more learning