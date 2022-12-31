## Description
This project is used for test assignment. 
There are three endpoints that can be called. 
* /booklist , get book list from this [here](https://openlibrary.org/dev/docs/api/subjects)
* /bookschedule , book a pick up schedule
* /getallschedule , get all list of saved schedule


## Usage
Open terminal / command prompt and run command below.
```bash
go run .
```

### Book List

To get book list, open Postman, set the request method to 'GET' and fill the URL field with value below.
```bash
http://localhost:8081/booklist
```

It will return result similar to value below.
```bash
[
    {
        "title": "Romeo and Juliet",
        "authors": [
            {
                "name": "William Shakespeare"
            }
        ],
        "edition_count": 969
    },
    {
        "title": "The Republic",
        "authors": [
            {
                "name": "Πλάτων"
            },
            {
                "name": "G.M.A. Grube"
            }
        ],
        "edition_count": 349
    },
]
```

### Book Pick Up Schedule

To book pick up schedule, open Postman, set the request method to 'POST' and fill the URL field with value below.
```bash
http://localhost:8081/bookschedule
```

#### Param
Prepare the parameter consisting name of the booker and book information from booklist endpoint. Below is example of parameter for request body.
```
{
	"name" : "Riky",
    "title": "Romeo and Juliet",
    "authors": [
        {
            "name": "William Shakespeare"
        }
    ],
    "edition_count": 969
}
```

If success, it will return message "Book success"

### Pick Up Schedule List

To pick up schedule list, open Postman, set the request method to 'GET' and fill the URL field with value below.
```bash
http://localhost:8081/getallschedule
```

It will return result similar to value below.
```bash
[
    {
        "Name": "Riky",
        "Book": {
            "title": "Romeo and Juliet",
            "authors": [
                {
                    "name": "William Shakespeare"
                }
            ],
            "edition_count": 969
        },
        "PickUpDate": "2022-12-31T21:45:58.7843247+07:00"
    }
]
```