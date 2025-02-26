Workout Bot
===========

Описание проекта
----------------

Workout Bot — это телеграм-бот, созданный для управления расписанием тренировок. Бот позволяет пользователям добавлять, удалять и просматривать свои занятия, а также получать уведомления о предстоящих тренировках.

Стек технологий
---------------

*   **Golang** : Язык программирования, используемый для реализации бизнес-логики бота.
    
*   **SQLite3** : Легковесная база данных для хранения информации о пользователях и их расписаниях тренировок.
    
*   **Telegram Bot API** : Интерфейс для взаимодействия с Telegram.
    
*   **GORM** : ORM (Object-Relational Mapping) для работы с базой данных SQLite3.
    

Функциональность
----------------

1.  **Добавление тренировок** : Пользователи могут добавлять тренировки в свое расписание с указанием даты, времени и типа тренировки.
    
2.  **Просмотр расписания** : Пользователи могут посмотреть все запланированные тренировки.
    
3.  **Удаление тренировок** : Возможность удалить конкретную тренировку из расписания.
    
4.  **Уведомления** : Бот отправляет напоминания за определенное время до начала тренировки.
    
5.  **Персонализация** : Каждый пользователь имеет свой уникальный набор тренировок.