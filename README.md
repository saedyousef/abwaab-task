# Abwaab Task

## Installation


```
$ cd /to/your/go/src/ON_YOUR_MACHINE & git clone https://github.com/saedyousef/abwaab-task.git

```

---

Now lets copy our .env.example into .env.
The .env file contains our secret keys and Twitter App's keys.
```

$ cd abwaab-task & cp .env.example /path/to/project/.env

```

> We will start our app now

`$ go run main.go`

You will notice that go modules are installing.
Once you've ran the app, database connection will be initialized.

Now we're ready to go!

---


Let's Browse the APIs

## APIs

> Auth APIs

Singup API, this api accept POST method, and `Content-type : application/json` `header`
request body: 
```
{
    "username": "anyusername",
    "password": "PASSWORD",
    "password_confirm": "PASSWORD",
    "name": "Saed Yousef" 
}

```
Check sample request and response from postman in the screenshot.
![Signup](screenshots/signup.png)

---

Login API, this api accept POST method, and `Content-type : application/json` `header`
request body: 
```
{
    "username": "YOUR_USERNAME",
    "password": "YOUR_PASSWORD"
}

```
Check sample request and response from postman in the screenshot.
![Login](screenshots/login.png)

---

Refresh Token API, this api accept POST method, and `Content-type : application/json` `header`
request body: 
```
{
    "refresh_token": "YOUR_REFRESH_TOKEN"
}

```
Afte your access token is expired you can request this api to refresh your access token by provides your `refresh_token`.

Check sample request and response from postman in the screenshot.
![Refresh Token](screenshots/refresh.png)

---
### How To authorize request?

To authorize a request, in postman Click on Authorization and select Bearer Token and place your access_token that is returned from login, signup or refresh apis.

---

> Twitter Search API

Twitter Search API, this api accept GET method, and requires a string query param "query" and return the first 50 matched results, these results will be saved in our database.

Authentication `required` to make a request on this API.

Check sample request and response from postman in the screenshot.

![Twitter Search](screenshots/twitter_search.png)

---

### Tweets CRUD APIs

> Create Tweet API, this api accept POST method, and `Content-type : application/json` `header`
request body: 
```
{
    "body": "Any string Body"
}

```
Authentication `required` to make a request on this API.

Check sample request and response from postman in the screenshot.

![Create Tweet](screenshots/create_tweet.png)

---

> Update Tweet API, this api accept PUT method, you need to pass the tweet id in the url, set the `Content-type : application/json` `header`
request body: 
```
{
    "body": "Any string Body"
}

```
Authentication `required` to make a request on this API.

Check sample request and response from postman in the screenshot.

![Update Tweet](screenshots/update_tweet.png)

---

> Delete Tweet API, this api accept DELETE method, you need to pass the tweet id in the url
Authentication `required` to make a request on this API.

Check sample request and response from postman in the screenshot.

![Delete Tweet](screenshots/delete_tweet.png)

---

> Tweet Details API, this api accept GET method, you need to pass the tweet id in the url

Check sample request and response from postman in the screenshot.

![Tweet Details](screenshots/tweet_details.png)

---

> Tweets List API, this api accept GET method, will return all tweets saved by the logged in user.
Pagination available on this api, you can pass 2 string query params to navigate the results.
Params : page(page number), limit(Results per page).

Authentication `required` to make a request on this API.

Check sample request and response from postman in the screenshot.

![Delete Tweet](screenshots/tweets_list.png)

---

### Referencees

[Gorm](http://gorm.io/docs/index.html)
<br/>

[Go-gin](https://godoc.org/)
<br/>

[Paginator](https://github.com/biezhi/gorm-paginator/)
<br/>
And some Articles on [Medium](https://medium.com/)