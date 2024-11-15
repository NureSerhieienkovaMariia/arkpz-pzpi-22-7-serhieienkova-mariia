ПЗПІ-22-7

Сергєєнкова Марія

Тема: Програмна система інтелектуальної настройки музичних інструментів

Опис теми:  
Ця система допомагає музикантам точно налаштувати різні типи гітар, забезпечуючи чистоту звучання нот і усуваючи фальшиві ноти. Програма використовує дані від датчиків для аналізу звукових частот і допомагає користувачам налаштовувати інструменти з урахуванням навколишніх умов, які можуть впливати на звучання (температура, вологість). Система також підтримує кілька режимів налаштування для різних жанрів музики та специфічних строїв інструментів.

Основна функціональність:  
Система надає зручний інтерфейс для інструментів різного типу і надає підказки щодо точної настройки. Додатково користувач може зберігати профілі налаштувань для різних інструментів та умов.

Складові проєкту:
1. Back-end
   - Обробка даних про частоту звуку, яку зчитують датчики, та аналіз їх точності.
   - Налаштування рекомендацій для кожного типу інструмента (електрична, акустична гітара, бас-гітара тощо).
   - Зберігання профілів налаштувань для різних інструментів і умов.

2. Front-end
   - Інтерфейс для користувачів, де вони можуть вибирати тип інструмента, налаштовувати параметри й отримувати підказки щодо корекції.
   - Візуалізація частот і розбіжностей у звучанні для кращого розуміння.
   - Підтримка української та англійської мов для зручності користувачів різних країн.

3. IoT
   - Підключення до IoT-пристроїв, що включають датчики часу гри та інтенсивності натискання, що визначатимуть ступінь зносу струн та надаватимуть рекомендації щодо заміни.
   - Моніторинг умов зберігання: температурні та вологісні датчики у футлярі або кімнаті контролюватимуть зовнішні умови, що впливають на налаштування та стан інструмента.
   - Сигналізація по заміні струн: система надсилатиме сповіщення, коли знос досягне критичного рівня або коли строки заміни наблизяться.
