# Problem 14 — Longest Collatz sequence
“The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.”

My solution is:

package main
import "fmt"
func main() {
numberThatGeneratesTheLongestChain := 1
chainLen := 1
for i := 2; i <= 1_000_000; i++ {
count := 1
start := i
for start != 1 {
if start % 2 == 0 {
start /= 2
} else {
start = start * 3 + 1
}
count++
}
if count > chainLen {
chainLen = count
numberThatGeneratesTheLongestChain = i
}
}
fmt.Printf("%d generates %d\n", numberThatGeneratesTheLongestChain, chainLen)

}
The solution, shown below, was again brute-forced but it works (< 0.2 seconds):

$ time ./problem14
837799 generates 525
real    0m0.243s
user    0m0.240s
sys 0m0.004s


