The suffix tree is computed from the suffix array in linear time. Although the suffix array is computed in O(n * log n) where n is the length of the string. 

There are multiple ways to make this better:
- [Here](https://arxiv.org/abs/1610.08305) is a paper which details a suffix array building algorithm which runs in linear time.
- Ukkonen's algorithm can be modified in a way that it builds the suffix tree in linear time. Both Cochamore's and Gusfield's book covers this.

A suffix automaton implementation would be great.
