package pattern

/*Паттерн Команда помогает нам, когда у нас есть множество схожих элементов, выполняющих разные действия.
Допустим у нас есть структуры, отражающие какие-то управляющие элементы(рычаги, кнопки и т.д.).
Каждый элемент может выполнять различную задачу и если мы попробуем решить это наследованием,
то получим иерархию классов, которая будет только расти по мере добавления новых действий,
а управлять ей будет сложно. Вместо этого можно выносить новые действия в служебные структуры Команд,
которые будут определять, какие действия нужно выполнить(как правило, переадресовывая их некому получателю).
Тогда у нас получится вызывающий код, который привязывает Конкретные команды к конкретным управляющим элементам.
Управляющие элементы в свою очередь, при обращении к ним, будут вызывать Команду, которая уже вызовет конкретную логику.

Так же мы можем выполнять какие-либо вспомогательные служебные действия при вызове команд. Т.к. вызывающий код
привязывает Команду к управляющему элементу, этот самый вызывающий код, между делом, может добавить логирование,
ведение истории(и отмены операций), нотификации.*/