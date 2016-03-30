# kafka-console
a simple kafka console consumer written in Go

## dependency

- https://github.com/Shopify/sarama

## usage

```kafka-console <broker1;broker2;...> <topic>```

## sample output

```
$ kafka-console 192.168.6.55:9092 yuankui_logsmash
2016/03/30 15:04:52 partitions: [0 1 2 3]
topic:[yuankui_logsmash] parition:[0] offset:[1638599] key:[] value:hello this is a
topic:[yuankui_logsmash] parition:[0] offset:[1638600] key:[] value:test message
```
