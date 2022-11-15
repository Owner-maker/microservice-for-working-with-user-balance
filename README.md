# microservice-for-working-with-user-balance
This service is a task for the Avito internship. Provides you a REST API to work with user balances (crediting funds, debiting funds, transferring funds from user to user, as well as a method for obtaining a user's balance).

## Вопросы / примечания по ТЗ
1. Резервирование средств: должно ли присутствовать отдельное поле как "резервный" счет пользователя?
  - Я реализовал это следующим образом - как таковое резервирование средств, насколько я понимаю, это списание денежных средств с "основного" счета пользователя (поле "balance" сущности "user") и начисление их на "резервный" (отдельный) счет пользователя, поэтому при покупке услуги, с поля "balance" сущности "user" списываются средства в размере стоимости услуги и создается новая "транзакция", которая уже хранит в себе информацию о т.н. зарезервированных средствах. Такая транзакция обладает статусом "active", а после того, как услуга будет выполнена "транзакция" получает статус "completed" и заполняется поле "timestamp" (время выполнения услуги), если услуга была отменена/не выполнена, "зарезервированные" средства возвращаются на "основной" счет пользователя (поле "balance"), а транзакция удаляется из БД.
