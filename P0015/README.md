# Problem 15 — Lattice paths
“Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.


How many such routes are there through a 20×20 grid?”

This problem is called “Lattice Paths”. These paths can be calculated using central binomial coefficients, such as:



n = 20


The solution can be checked with Wolfram Alpha, which is linked below:

n = 20; (2 * n)! / (n!)^2 - Wolfram|Alpha
Uh oh! Wolfram|Alpha doesn't run without JavaScript. Please enable JavaScript. If you don't know how, you can find…
www.wolframalpha.com


As you can see, the solution involves using a factorial function. The problem is that we need to calculate a factorial for the 40 number that would overflow type. We will need to use big numbers again. Let’s write a factorial function using big numbers that will help us to calculate (2 * n)! / (n!)²:

package main
import (
"fmt"
"math/big"
)
func factorial(n int) *big.Int {
f := big.NewInt(1)
for i := 1; i <= n; i++ {
f.Mul(f, big.NewInt(int64(i)))
}
return f
}
func main() {
// n = 20; (2 * n)! / (n!)^2
n := 20
a := factorial(2  * n)
b := factorial(n)
b.Mul(b, b)
result := new(big.Int)
result.Div(a, b)
fmt.Println(result)
}
Output:

$ time ./problem15
137846528820
real    0m0.005s
user    0m0.001s
sys 0m0.005s
My Strategy To Learn the Basics of a Programming Language
Say I want to learn Ruby, or Nim, or Clojure, or Haskell, I would create an account on Project Euler and try to solve the first 15 problems.

Once I am done, I go back to the first problem and try to change the code based on the refactoring techniques I know from other programming languages.

For example, for the first problem I used the following condition:

if i % 3 == 0 || i % 5 == 0 {
sum += i
}
I could actually change it to the followiong:

package main
import "fmt"
func isDivisibleBy3Or5(n int) bool {
return n % 3 == 0 || n % 5 == 0
}
func main() {
sum := 0
for i := 1; i < 1_000; i++ {
if isDivisibleBy3Or5(i) {
sum += i
}
}
fmt.Printf("Sum: %d\n", sum)
}
And then ask myself, “Is this idiomatic?”. And then, dig about it and move forward.

Conclusion
I encourage you to try to solve a few of these problems with the programming language of your choice. The goal is to have some fun while learning!

Please let me know your thoughts in the comments. Thank you!

https://github.com/leogtzr/15_projecteuler_solutions_Golang



