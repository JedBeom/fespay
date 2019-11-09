import csv
import io  # for newline option
import uuid

d = dict()  # new dictionary
f = open("input.txt", mode="rt", encoding='UTF8')  # load input.txt
for line in f:
    if line.isspace():
        continue  # continue if there are only whitespaces
    a = line.rstrip('\n').split("\t")  # remove last \n, split by tab

    for i in [3, 2, 1]:
        if len(a) == i * 2:
            for j in range(i):
                d[a[j]] = a[j+i]
            break

output = io.open("output.csv", mode="w", encoding="UTF8", newline="\n")  # open file with newline="\n". 윈도우 망해라
cw = csv.writer(output)
cw.writerow(["id", "grade", "class", "number", "name", "card_code", "pay_code", "type"])  # header


def get_uuid():
    x = uuid.uuid4()  # generate random uuid
    return str(x)


count = 0

for card_code, v in d.items():
    attr = v.split()  # split by whitespaces

    g = None
    c = None
    num = None

    if attr[0][1:] == "학년":  # if student
        user_type = 1
        g = attr[0][0]  # grade
        c = attr[1][0]  # class
        num = ""  # number initial
        if len(attr[2]) == 3:  # two-digit number
            num = attr[2][:2]
        else:  # one-digit number
            num = attr[2][0]

        name = attr[3]  # get name

    elif attr[0] == "교사":
        user_type = 2
        name = attr[1]  # eg) "교사 범XX"
    elif attr[0] == "학부모":
        user_type = 3
        name = attr[2]  # 학부모는 "학부모 학부모 김XX" 이런 형식...
    else:
        continue

    cw.writerow([get_uuid(), g, c, num, name, card_code, card_code, user_type])
    count += 1

print(f"Wrote {count} rows.")
