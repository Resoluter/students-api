# Students-API with Golang and PostgreSQL

#### The Object 🍵
| Properties | Description | Type  |
|:----------- |:---------------|:--------|
|first_name| First name | String| 
|last_name| Last name | String |
|email| Email | String | 

#### Routes 
| Routes | HTTP Methods| Description
|:------- |:---------------|:--------------
| /api/students/     | GET                  | Displays all identity
| /api/students/      | POST               | Creates a new identity
| /api/students/{id}| GET     | Displays a specific identity, given its id
| /api/students/{id}| PUT  | Update identitiy Value
| /api/students/{id}}| DELETE | Deletes a specific identity, given its id
	
### Technologies
Project is created with:

* Golang 
* gorilla/mux 
* lib/pq  
* joho/godotenv 
* DBeaver