num_list = ["ling", "yi", "er", "san", "si", "wu", "liu", "qi", "ba", "jiu"]

numstr = input()

count: int = 0
for num in numstr:
    count += int(num)

if count < 10:
    print(num_list[count])
elif count < 100:
    print(num_list[int(count/10)], num_list[count % 10])
else:
    print(num_list[int(count/100)],
          num_list[int(count/10 % 10)],
          num_list[count % 10])

