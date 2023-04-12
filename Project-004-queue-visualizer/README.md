# Queue Visualizer

A simple queue program along with unicode lines to visualize the queue.

## Operations Implemented

- [x] Print queue after PUSH
- [x] Print queue after POP
- [x] Print queue

## Limits

- MAX Queue size can be 10
- Only 5 digit number is allowed to be pushed to queue

## Results

- PUSH

```
Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit 
1
Enter nummber to Push
10


   ⇓
╔═════╗
╣   10║
╚═════╝ ⇖ FI


Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit 
1
Enter nummber to Push
20


   ⇓
╔═════╗╔═════╗
║   20╠╣   10║
╚═════╝╚═════╝ ⇖ FI


Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
1
Enter nummber to Push
30


   ⇓
╔═════╗╔═════╗╔═════╗
║   30╠╣   20╠╣   10║
╚═════╝╚═════╝╚═════╝ ⇖ FI


Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
1
Enter nummber to Push
40


   ⇓
╔═════╗╔═════╗╔═════╗╔═════╗
║   40╠╣   30╠╣   20╠╣   10║
╚═════╝╚═════╝╚═════╝╚═════╝ ⇖ FI


Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
1
Enter nummber to Push
50


   ⇓
╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗
║   50╠╣   40╠╣   30╠╣   20╠╣   10║
╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝ ⇖ FI


Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
1
Enter nummber to Push
60


   ⇓
╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗
║   60╠╣   50╠╣   40╠╣   30╠╣   20╠╣   10║
╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝ ⇖ FI


Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
1
Enter nummber to Push
70


   ⇓
╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗
║   70╠╣   60╠╣   50╠╣   40╠╣   30╠╣   20╠╣   10║
╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝ ⇖ FI


Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
1
Enter nummber to Push
80


   ⇓
╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗
║   80╠╣   70╠╣   60╠╣   50╠╣   40╠╣   30╠╣   20╠╣   10║
╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝ ⇖ FI


Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
1
Enter nummber to Push
90


   ⇓
╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗
║   90╠╣   80╠╣   70╠╣   60╠╣   50╠╣   40╠╣   30╠╣   20╠╣   10║
╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝ ⇖ FI


Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
1
Enter nummber to Push
100


   ⇓
╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗
║  100╠╣   90╠╣   80╠╣   70╠╣   60╠╣   50╠╣   40╠╣   30╠╣   20╠╣   10║
╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝ ⇖ FI


Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
1
Queue is full !!!
```


- POP

```
   ⇓
╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗
║  100╠╣   90╠╣   80╠╣   70╠╣   60╠╣   50╠╣   40╠╣   30╠╣   20╠╣   10║
╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝ ⇖ FI


Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
2
Popping the start


   ⇓
╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗
║  100╠╣   90╠╣   80╠╣   70╠╣   60╠╣   50╠╣   40╠╣   30╠╣   20║
╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝ ⇖ FI


Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
2
Popping the start


   ⇓
╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗
║  100╠╣   90╠╣   80╠╣   70╠╣   60╠╣   50╠╣   40╠╣   30║
╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝ ⇖ FI


Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
2
Popping the start


   ⇓
╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗
║  100╠╣   90╠╣   80╠╣   70╠╣   60╠╣   50╠╣   40║
╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝ ⇖ FI


Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
2
Popping the start


   ⇓
╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗
║  100╠╣   90╠╣   80╠╣   70╠╣   60╠╣   50║
╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝ ⇖ FI


Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
2
Popping the start


   ⇓
╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗
║  100╠╣   90╠╣   80╠╣   70╠╣   60║
╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝ ⇖ FI


Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
2
Popping the start

   ⇓
╔═════╗╔═════╗╔═════╗╔═════╗
║  100╠╣   90╠╣   80╠╣   70║
╚═════╝╚═════╝╚═════╝╚═════╝ ⇖ FI


Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
2
Popping the start


   ⇓
╔═════╗╔═════╗╔═════╗
║  100╠╣   90╠╣   80║
╚═════╝╚═════╝╚═════╝ ⇖ FI


Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
2
Popping the start


   ⇓
╔═════╗╔═════╗
║  100╠╣   90║
╚═════╝╚═════╝ ⇖ FI


Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
2
Popping the start


   ⇓
╔═════╗
╣  100║
╚═════╝ ⇖ FI


Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
2
Popping the start
Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
2
Queue is empty !!!
```


- Print

```
Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
3

   ⇓
╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗╔═════╗
║  100╠╣   90╠╣   80╠╣   70╠╣   60╠╣   50╠╣   40╠╣   30╠╣   20╠╣   10║
╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝╚═════╝ ⇖ FI


```

---

```

Enter the operation number: 1.Push       2.Pop   3.Print Stack    4.Exit
3
Queue is empty !!!
```


