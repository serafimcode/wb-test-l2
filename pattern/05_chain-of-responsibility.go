package pattern

/*Цепочка обязанностей позволяет нам динамически создавать и менять набор обработчиков команд. К примеру,
http запрос. Его можно пропустить через несколько "перехватчиков", каждый из которых будет выполнять
одну конкретную операцию с запросом(в идеале, не мутирую запрос). Таким образом, вместо того, чтобы складывать
все обработчики в одно место, создавая "спагетти" из условных операторов, мы можем разбить эти операции
по отдельным структурам и в вызывающем коде выбирать те из них, которые нам нужны и устанавливать порядок выполнения
на случай, если нам нужно остановить обработку на каком то шаге.*/
