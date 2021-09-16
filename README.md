##Микросервис для работы с балансом пользователей.


### Запуск
Бд создается при запуске 
```sh
$ docker-compose up -d
```


##Примеры запросов:
Базовый url http://localhost:7777
### GET
Получение баланса пользователя

`/get_user_balance?id=1`
 

Получение баланса пользователя в заданной валюте

`/get_user_balance?id=1&currency=USD`



### POST

Начислить средства на баланс

`/add_or_withdraw_funds`
```
{
"userId":"6",
"funds":5000
}
```

Списать средства 
```
{
    "userId":"6",
    "funds":5000,
    "isWithdraw":true
}
```


Перевести деньги другому юзеру

`/transfer_funds`
```
{
    "fromUserId":"5",
    "toUserId":"2",
    "funds":50
}
```

