# すべて取らないといけないとき、例えばN=15は[1, 2, 3, 4, 5]で5になる。このあと16~(15+6-1)=20までは6でいけるやろ、そういうことじゃ

x = int(input())
i = 0
s = 0
while True:
    if s >= x:
        break
    i += 1
    s += i
print(i)
