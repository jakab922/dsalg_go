# Proof that building the sieve works in linear time

The *first* array holds the smallest prime divisor for all numbers from $2$ to $n$. For each number $i \in 2 \ldots n$ there is a unique prime $p_{i}$ such that $k = p_{k} * i$ where $p_{k}$ is the smallest prime divisor of $k$ and the smallest prime divisor of $i$ is at least $p_{i}$. The for cycle

```
for j := 0; j < len(primes) && primes[j] <= first[i] && i*primes[j] <= n; j++ {
    first[i*primes[j]] = primes[j]
}
```

sets the values between $i * primes[0]$ and $i * first[i]$ where $first[i]$ is also a prime. $i * primes[j_1] \neq i * primes[j_2]$ if $j_1 \neq j_2$. Same is true if the $j$ index is the same and $i$ is different. Now suppose $k = i_1 * primes[j_1] = i_2 * primes[j_2]$. We can assume without loss of generality that $primes[j_1] < primes[j_2]$. This means that $primes[j_2] > primes[j_1] \geq first[k]$ so we can't enter the loop with $primes[j_2]$ and with this we can see that we only set one number once. Note that it also proves that it sets the correct value in the first array. All that is left to see that we set them at least once. But that's easy since $k = i * primes[j]$ so we set the value when the outer loop takes value $i$. 


