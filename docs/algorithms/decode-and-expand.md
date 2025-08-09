# Decode and Expand

This is a recurring pattern in problems that involve nested structures with hierarchical rules, and it’s particularly common with strings that represent some encoded or compressed format. For example, you need to decode a string that has been compressed by providing the number of repetitions followed by the characters inside of a pair of brackets.
```
encoded = "3[a2[b]]
decoded = "abbabbabb"
```