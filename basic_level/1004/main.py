num = int(input())

max_score = -1
min_score = 101

for i in range(num):
    line = input()

    current_name = line.split(" ")[0]
    current_id = line.split(" ")[1]
    current_score = int(line.split(" ")[2])

    if current_score > max_score:
        max_name = current_name
        max_id = current_id
        max_score = current_score

    if current_score < min_score:
        min_name = current_name
        min_id = current_id
        min_score = current_score

print(max_name, max_id)
print(min_name, min_id)
