curl -X PUT localhost:8080/events/9 -i --json '{
    "name": "SuperEvent updated!", 
    "description": "Just a test event", 
    "location": "Here", 
    "dateTime": "2006-01-02T15:04:05Z"
}' 
echo

# curl -X PUT localhost:8080/events/9 -H 'Content-Type: application/json' --data '{ "name": "SuperEvent updated", "description": "Just a test event", "location": "Here", "dateTime": "2006-01-02T15:04:05Z" }' 
