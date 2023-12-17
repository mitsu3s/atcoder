n = int(input())
a = list(map(int, input().split()))

r = 0

while True:
    for i in range(n):
        if a[i] % 2 == 0:
            a[i] = a[i] // 2
        else:
            print(r)
            exit()
    r += 1
