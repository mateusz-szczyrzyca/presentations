#!/usr/bin/env python3

def fib(num):
  a,b = 0, 1
  for i in range(0, num):
    a, b = b, a + b

  return a

print(fib(50))
