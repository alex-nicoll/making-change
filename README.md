# making_change

This is an algorithm for the problem of making change for a sum while minimizing the number of coins. It seems to be faster than the dynamic programming solution that is usually presented.

The problem can be restated as:

minimize     x<sub>1</sub> + x<sub>2</sub> + ... + x<sub>n</sub>

subject to   d<sub>1</sub>x<sub>1</sub> + d<sub>2</sub>x<sub>2</sub> + ... + d<sub>n</sub>x<sub>n</sub> = sum

where d's are the coin denominations and x's are the number of coins of each denomination.

The idea behind the algorithm is as follows. First assume that the d's are decreasing. Then, set x<sub>1</sub> to it's maximum possible value, i.e., x<sub>1</sub>=floor(sum/d<sub>1</sub>). Let r = mod(sum,d<sub>1</sub>) be the remaining sum to make change for. Proceed recursively with r as the new sum and d<sub>2</sub>...d<sub>n</sub> as the new denominations, to compute x<sub>2</sub>...x<sub>n</sub>. We now have the greedy solution, which we were able to compute very quickly using division. We'll work our way towards an optimal solution. Iteratively decrement x<sub>1</sub> and add d<sub>1</sub> to r. In each iteration, x<sub>1</sub>+ceil(r/d<sub>2</sub>) is the number of coins in the best possible solution when considering only x<sub>1</sub>, d<sub>2</sub>, and unknown d<sub>3</sub>...d<sub>n</sub>. If our current best solution is better (fewer coins), then stop iterating. Otherwise, proceed recursively with r as the new sum and d<sub>2</sub>...d<sub>n</sub> as the new denominations, to compute x<sub>2</sub>...x<sub>n</sub>, and set this as the new best if it is less than the current best. The stopping condition above leads to the optimal solution because current\_best &le; x<sub>1</sub>+ceil(r/d<sub>2</sub>) &#8658; current\_best &le; (x<sub>1</sub>-1)+ceil((r+d<sub>1</sub>)/d<sub>2</sub>).
