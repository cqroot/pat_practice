num: int = int(input())

count = 0
while num > 1:
    if num % 2 == 1:
        num = int((num * 3 + 1) / 2)
    else:
        num = int(num / 2)
    count += 1

print(count)
