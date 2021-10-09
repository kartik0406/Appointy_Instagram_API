# Appointy_Instagram_API
Task  | Instagram Backend API: 
An Instagram RESTFUL API created in Go language and mongoDB Database

# Code Execution
![](https://github.com/kartik0406/Appointy_Instagram_API/blob/master/src/screenshots/img1.png)

# Create an User
![](https://github.com/kartik0406/Appointy_Instagram_API/blob/master/src/screenshots/img2.png)
![](https://github.com/kartik0406/Appointy_Instagram_API/blob/master/src/screenshots/img3.png)

# Get a user using id
![](https://github.com/kartik0406/Appointy_Instagram_API/blob/master/src/screenshots/img4.png)
 
# Create a Post
![](https://github.com/kartik0406/Appointy_Instagram_API/blob/master/src/screenshots/img5.png)
![](https://github.com/kartik0406/Appointy_Instagram_API/blob/master/src/screenshots/img6.png)
![](https://github.com/kartik0406/Appointy_Instagram_API/blob/master/src/screenshots/img7.png)
  
# Get a post using id
![](https://github.com/kartik0406/Appointy_Instagram_API/blob/master/src/screenshots/img8.png)
  
# List all posts of a user
![](https://github.com/kartik0406/Appointy_Instagram_API/blob/master/src/screenshots/img9.png)


## Constraints:

1. Complete API has been developed using Go and mongodb
2. MongoDB Cluster has been used to storage
3. Only standard libraries and given libraries have been used

## Functionalities:

1. Create a User
* POST request
* JSON request body used
* Added data to the URL ‘/users'

2. Display details of all users
* GET request
* Display the user (unique) id, name, email and hash of the password
* URL: "/users"

3. Fetch details of the user using id
* GET request
* Displays user id, name, email and hash of the password
* URL: "/users/_id_"

4. Create a Post
* POST request
* JSON request body used
* Added data to the URL ‘/posts'

5. Display details of all posts (like scrolling through the feed)
* GET request
* Display the post (unique) id, caption, imageURL and timestamp
* URL: "/posts"

6. Fetch details of all the posts of a particular user using id
* GET request
* Display the post (unique) id, caption, imageURL and timestamp (for all posts of that user)
* URL: "/posts/_id_"

7. Implemented an approach to Pagination in handler/getUserPost.go

8. Tried to Develop Unit testing to test request and response in test_endpoints.go

## Attributes:

1. User:
* ID
* Name
* Email
* Password (saved as string as md5 hash)

2. Post:
* Id
* Caption
* Image URL
* Posted Timestamp
* userId

## Added functionalities:

1. Passwords have been encrypted and they can't be reverse engineered (using md5).
2. Unit Testing code is developed to test the request and response routes. 

## Instructions to Run and Setup:
- Install go from https://golang.org/dl/
- Install mongo dependencies (run go get go.mongodb.org/mongo-driver/mongo)
- Fork the project and download as a zip.
- Now extract and open the folder.
- Now open the folder with terminal 
- cd src 
- Run commands: go run main.go
- The server started locally at the port 9000.
