## How to run this app:
## Prequisite
Please make sure port 8080 is free

## Option 1
Run from docker
```
$ docker compose build --no-cache
$ docker compose up -d
```

## Option 2
Run directy with this command on termimal
```
$ make run
```

To test the api download bruno
https://www.usebruno.com/

Open bruno > Collections> Open Collection > Select folder `./amartha/api_collection`

Dont forget to select "development" environment

- Available login usernames are: user1, user2, user3, admin
- Their ids are: 1, 2, 3, 4 (for borrower/investor/reporter user id)

Application flow:
- POST {{base_url}}/account/login
- POST {{base_url}}/loans
- POST {{base_url}}/loans/:id/approve
- POST {{base_url}}/loans/:id/invest
- POST {{base_url}}/loans/:id/disburse