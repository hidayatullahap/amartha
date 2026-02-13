How to run:

run this command on termimal
```
$ make run
```

To test the api download bruno
https://www.usebruno.com/

Open bruno > Collections> Open Collection > Select folder `./amartha/api_collection`

- Available login usernames are: user1, user2, user3, admin
- Their ids are: 1, 2, 3, 4 (for borrower/investor/reporter user id)

Application flow:
- POST {{base_url}}/account/login
- POST {{base_url}}/loans
- POST {{base_url}}/loans/:id/approve
- POST {{base_url}}/loans/:id/invest
- POST {{base_url}}/loans/:id/disburse