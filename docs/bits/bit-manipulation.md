# Bit Manipulation
A single digit of a binary sequence is called a **bit** and can be either **0** or **1**. We say that a certain bit is **set**, if it is one, and **cleared** if it is zero. 

Bits are the smallest unit of data used by computers to write instructions. In a binary sequence such as 1100 the bits are read from *right* to *left* where the bit on the right (0) is the **Least Significant Bit (LSB)** and the bit on the left (1) is the **Most Significant Bit (MSB)**.

## Bitwise Operators
Bit manipulation is the process of manipulating bits or groups of bits ina  byte. Bit manipulation is often used to perform operations on data that are otherwise difficult or impossible to perform with traditional operators.

### & Operator
The bitwise *AND* operator compares each bit of its first operand with the corresponding bit of its second operand. If **both bits are 1**, the corresponding result bit is set to **1**. Otherwise, the corresponding result bit is set to **0**.

|  X  |  Y  |  X & Y  |
| :-: | :-: | :-----: |
|  0  |  0  |    0    |
|  0  |  1  |    0    |
|  1  |  0  |    0    |
|  1  |  1  |    1    |

```
x     = 01011000
y     = 01010111
----------------
x & y = 01010000
```

### | Operator
The bitwise inclusive *OR* operator compares each bit of its first operand with the corresponding bit of its second operand.  If **one of the two bits is 1**, the corresponding result bit is set to **1**. Otherwise, the corresponding result bit is set to **0**.

|  X  |  Y  |  X | Y  |
| :-: | :-: | :-----: |
|  0  |  0  |    0    |
|  0  |  1  |    1    |
|  1  |  0  |    1    |
|  1  |  1  |    1    |

```
x     = 01011000
y     = 01010111
----------------
x | y = 01011111
```

### ^ Operator
The bitwise exclusive OR *(XOR)* operator compares each bit of its first operand with the corresponding bit of its second operand. If **one bit is 0 and the other bit is 1**, the corresponding result bit is set to **1**. Otherwise, the corresponding result bit is set to **0**.

|  X  |  Y  |  X ^ Y  |
| :-: | :-: | :-----: |
|  0  |  0  |    0    |
|  0  |  1  |    1    |
|  1  |  0  |    1    |
|  1  |  1  |    0    |

```
x     = 01011000
y     = 01010111
----------------
n ^ y = 00001111
```

### ~ Operator
The bitwise complement *(NOT)* operator flips each bit of a number, if a bit is set the operator will clear it, if it is cleared the operator sets it.

|  X  |  ~X  |
| :-: | :--: |
|  0  |   1  |
|  1  |   0  |

```
x     = 01011000
----------------
~x    = 10100111
```

### References
- Algorithms for Competitive Programming. (2024). *Bit manipulation*. https://cp-algorithms.com/algebra/bit-manipulation.html
- educative. (2022). *The basics of bit manipulation*. https://www.educative.io/blog/bit-manipulation-algorithms