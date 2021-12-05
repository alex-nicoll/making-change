# making-change

This is an algorithm for the problem of making change for a sum while minimizing the number of coins. It seems to be faster than the dynamic programming solution that is usually presented.

The problem can be restated as:

minimize     x1 + x2 + ... + xn

subject to   d1\*x1 + d2\*x2 + ... + dn\*xn = sum

where d's are the coin denominations and x's are the number of coins of each denomination.

The idea behind the algorithm is as follows. First assume that the d's are decreasing. Then, set x1 to it's maximum possible value, i.e., x1=floor(sum/d1). Let r=mod(sum,d1) be the remaining sum to make change for. Proceed recursively with r as the new sum and d2...dn as the new denominations, to compute x2...xn. We now have the greedy solution, which we were able to compute very quickly using division. We'll work our way towards an optimal solution. Iteratively decrement x1 and add d1 to r. In each iteration, x1+ceil(r/d2) is the number of coins in the best possible solution when considering only x1, d2, and unknown d3...dn. If our current best solution is better (fewer coins), then stop iterating. Otherwise, proceed recursively with r as the new sum and d2...dn as the new denominations, to compute x2...xn, and set this as the new best if it is less than the current best. The stopping condition above leads to the optimal solution because current\_best <= x1+ceil(r/d2) => current\_best <= (x1-1)+ceil((r+d1)/d2).
