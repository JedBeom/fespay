import csv
import uuid

d = dict()
f = open("student.txt", encoding='utf-8-sig')
for line in f:
    if line.isspace(): continue
    a = line.rstrip('\n').split("\t")

    for i in [3,2,1]:
        if len(a) == i * 2:
            for j in range(i):
                d[a[j]] = a[j+i]
            break

output = open("output.csv", mode="w")
cw = csv.writer(output)
cw.writerow(["id", "grade", "class", "number", "name", "card_code", "type"])

def get_uuid():
    x = uuid.uuid4()
    return str(x)

for k, v in d.items():
    attr = v.split()

    if len(attr) > 1: # if student
        g = attr[0][0] # grade
        c = attr[1][0] # class
        num = "" # number initial
        if len(attr[2]) == 3: # two-digit number
            num = attr[2][:2]
        else: # one-digit number
            num = attr[2][0]

        name = attr[3] # get name
        usertype = 1
        cw.writerow([get_uuid(), g,c,num,name,k,usertype])

    else:
        name = attr[0]
        usertype = 2
        cw.writerow([get_uuid(), 0, 0, 0, name,usertype])

print(200)
