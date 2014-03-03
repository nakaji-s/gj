gj
--

golang json formatter

* example

```
$ cat sample.json 
{"object":{"buffer_size": 10,"Databases":[{"host": "localhost","user": "root","pass": "","type": "mysql","name": "go","Tables":[{"name": "testing","statment": "teststring","regex": "teststring ([0-9]+) ([A-z]+)","Types":[{"id": "int","value": "string"}]}]}]}}

$ gj sample.json 
{
  "object": {
    "buffer_size": 10,
    "Databases": [
      {
        "host": "localhost",
        "user": "root",
        "pass": "",
        "type": "mysql",
        "name": "go",
        "Tables": [
          {
            "name": "testing",
            "statment": "teststring",
            "regex": "teststring ([0-9]+) ([A-z]+)",
            "Types": [
              {
                "id": "int",
                "value": "string"
              }
            ]
          }
        ]
      }
    ]
  }
}
```

happy :)

