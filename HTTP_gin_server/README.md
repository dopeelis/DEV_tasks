## Самый простой HTTP сервер с использованием gin
#

Методы:

GET `'api/v1/about'` Ответ: 200 + json     
GET `'api/v1/user/id'` Ответ: 200 + json с id (Если id < 0: 406 + json с описанием ошибки)