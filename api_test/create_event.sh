curl localhost:8080/events -i \
    -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAdGVzdC5jb20iLCJleHAiOjE3Mjc5MjA4NzYsInVzZXJJZCI6MX0.8pX9A7d4wcDjBGslz5wqYfh2b0BY162KjOcWLg68lJE' \
    --json '{
    "name": "SuperEvent", 
    "description": "Just a test event", 
    "location": "Here", 
    "dateTime": "2006-01-02T15:04:05Z"
}' 
echo
