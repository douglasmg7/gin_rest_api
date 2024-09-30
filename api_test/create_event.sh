curl localhost:8080/events -i --json '{
    "name": "SuperEvent", 
    "description": "Just a test event", 
    "location": "Here", 
    "dateTime": "2006-01-02T15:04:05Z"
}' 
echo
