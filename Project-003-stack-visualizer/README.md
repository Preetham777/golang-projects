# Stack Visualizer

A simple stack program along with unicode lines to visualize the stack.

## Operations Implemented

- [x] Print stack after PUSH
- [x] Print stack after POP
- [x] Print stack

## Limits

- MAX Stack size can be 10
- Only 5 digit number is allowed to be pushed to stack

## Results

- PUSH

```
Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit 
1
Enter nummber to Push
10
╔═════╗
║   10║ ⟸   top
╚═════╝
Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
1
Enter nummber to Push
20
╔═════╗
║   20║ ⟸   top
╠═════╣
║   10║
╚═════╝
Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
1
Enter nummber to Push
30
╔═════╗
║   30║ ⟸   top
╠═════╣
║   20║
╠═════╣
║   10║
╚═════╝
Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
1
Enter nummber to Push
40
╔═════╗
║   40║ ⟸   top
╠═════╣
║   30║
╠═════╣
║   20║
╠═════╣
║   10║
╚═════╝
Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
1
Enter nummber to Push
50
╔═════╗
║   50║ ⟸   top
╠═════╣
║   40║
╠═════╣
║   30║
╠═════╣
║   20║
╠═════╣
║   10║
╚═════╝
Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
1
Enter nummber to Push
60
╔═════╗
║   60║ ⟸   top
╠═════╣
║   50║
╠═════╣
║   40║
╠═════╣
║   30║
╠═════╣
║   20║
╠═════╣
║   10║
╚═════╝
Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
1
Enter nummber to Push
70
╔═════╗
║   70║ ⟸   top
╠═════╣
║   60║
╠═════╣
║   50║
╠═════╣
║   40║
╠═════╣
║   30║
╠═════╣
║   20║
╠═════╣
║   10║
╚═════╝
Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
1
Enter nummber to Push
80
╔═════╗
║   80║ ⟸   top
╠═════╣
║   70║
╠═════╣
║   60║
╠═════╣
║   50║
╠═════╣
║   40║
╠═════╣
║   30║
╠═════╣
║   20║
╠═════╣
║   10║
╚═════╝
Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
1
Enter nummber to Push
90
╔═════╗
║   90║ ⟸   top
╠═════╣
║   80║
╠═════╣
║   70║
╠═════╣
║   60║
╠═════╣
║   50║
╠═════╣
║   40║
╠═════╣
║   30║
╠═════╣
║   20║
╠═════╣
║   10║
╚═════╝
Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
1
Enter nummber to Push
100
╔═════╗
║  100║ ⟸   top
╠═════╣
║   90║
╠═════╣
║   80║
╠═════╣
║   70║
╠═════╣
║   60║
╠═════╣
║   50║
╠═════╣
║   40║
╠═════╣
║   30║
╠═════╣
║   20║
╠═════╣
║   10║
╚═════╝
Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit 
1
Stack is full !!!
Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
```


- POP

```
Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
3
╔═════╗
║   30║ ⟸   top
╠═════╣
║   20║
╠═════╣
║   10║
╚═════╝
Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
2
Popping the top
╔═════╗
║   20║ ⟸   top
╠═════╣
║   10║
╚═════╝
Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
2
Popping the top
╔═════╗
║   10║ ⟸   top
╚═════╝
Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
2
Popping the top
╔═════╗
╚═════╝
Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
2
Stack is empty !!!
Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
4
```

- Only 5-digit number is allowed

```
Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit 
1
Enter nummber to Push
37494748
Max 5 digit number is allowed
Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
```