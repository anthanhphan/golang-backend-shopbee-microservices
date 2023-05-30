### Please make sure that you have make and golang compiler!
Run developer mode:
```make dev```

Run production mode:
```make pro```

API:
#### Sign in 
+ Route ```http://localhost:3000/v1/sign-in```
+ Body parse
```
{
    "email": "thanhanphan18@gmail.com",
    "password": "12345"
}
```

#### Sign up 
+ Route: ```http://localhost:3000/v1/sign-up```
+ Body parse
```
{
    "fullname": "Phan Thanh An",
    "email": "thanhanphan18@gmail.com",
    "phonenumber": "0325364474",
    "password": "12345",
    "sex": "M"
}
```
