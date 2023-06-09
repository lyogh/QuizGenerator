# Quiz Generator
## Назначение

Проект предназначен для генерации карточек вопросов тестов на основе правдивых фактов. Карточки генерируются случайным образом с различными вариантами ответов.
Результат генерации используется для проведения тестирования утилитой QuizGen или выгружается в форматы тестов Aiken, GIFT и т.д. 

### Пример 1: Генерация карточек тестовых вопросов на основе географических фактов.

Имеем список фактов представленный в формате YAML (примеры см. в папке assets):
```
# Географические факты
---
- Все факты:
  - Москва:
    - город
    - река
    - столица
    - столица России
    - город в России
    - город в России c 10 млн. жителей
  - Берлин:
    - город
    - город в Германии
    - столица
    - столица Германии
  - Россия:
    - страна
  - Лена:
    - река
  - Волга:
    - река
  - Волгоград:
    - город
    - город в России
```
После генерации получаем список тестовых вопросов в форме Aiken:
```
Верно ли утверждение: Россия - город в России
A. Нет
B. Да
ANSWER: A

Верно ли утверждение: Волга - река
A. Да
B. Нет
ANSWER: A

Выберите верные утверждения
A. Правильные ответы: D, C
B. Москва - столица Германии
C. Берлин - река
D. Лена - город в Германии
E. Берлин - город
ANSWER: E

Верно ли утверждение: Берлин - город в Германии
A. Да
B. Нет
ANSWER: A

Верно ли утверждение: Волгоград - столица Германии
A. Да
B. Нет
ANSWER: B
```
## Утилита QuizGen

### Режим тестирования
Запуск тестирования по фактам описанным в файле `data.yaml`: 
```
QuizGen.exe data.yaml
```

### Режим выгрузки
Выгрузка вопросов теста по фактам описанным в файле `data.yaml` в файл формата GIFT с названием `out.txt`: 
```
QuizGen.exe -f gift data.yaml out.txt
```
