# loop

Seriously, I shouldn't call this test app loop. But that is basically what it does.

- `-i` number of 1 second iterations, 1 line printed per each
- `-e` exit status

It also captures signals and output a proper message to the console.

```
$loop -i 4
Iteration #1
Iteration #2
Iteration #3
Iteration #4
```
