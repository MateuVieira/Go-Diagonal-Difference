# Go-Diagonal-Difference


<h1 align="center">
  <img alt="TackCash" title="TackCash" src="https://user-images.githubusercontent.com/1194257/65596422-1cef2080-df97-11e9-9abb-a225204d1805.png" width="300px"/>
</h1>

### Description

> This problem was taken from Hackerrank for the Simple Array Sum - Patrice Algorithms.

Given a square matrix, calculate the absolute difference between the sums of its diagonals.

For example, the square matrix  *arr* is shown below:

```bash
1 2 3
4 5 6
9 8 9  
```
The left-to-right diagonal: *1+5+9=15*. The right to left diagonal: *3+5+9=17* . Their absolute difference is *|15-17| = 2*.

Function description

Complete the **diafonalDifference** function in the editor below.

diagonalDifference takes the following parameter:

 - int arr[n][m]: an array of integers
 
Return

 - int: the absolute diagonal difference
 
Input Format

The first line contains a single integer,*n*, the number of rows and columns in the square matrix *arr*.
Each of the next *n* lines describes a row, *arr[i]*, and consists of *n* space-separated integers *arr[i][j]*.

Output Format

Return the absolute difference between the sums of the matrix's two diagonals as a single integer.

Sample Input
```bash
3
11 2 4
4 5 6
10 8 -12
```

Sample Output
```bash
15
```

Explanation

The primary diagonal is:

```bash
11
   5
     -12
```

Sum across the primary diagonal: *11 + 5 - 12 = 4*

The secondary diagonal is:

```bash
     4
   5
10
```

Sum across the secondary diagonal: *4 + 5 + 10 = 19*
Difference: *|4 - 19| = 15*

Note: *|x|* is the absolute value of *x*

### Motivation
I'm starting to learn to program in Go and I believe that one of the best ways is coding.
