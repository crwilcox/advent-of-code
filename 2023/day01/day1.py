with open("input") as input:
    sum = 0
    for x in input.readlines():
        first = 0
        last = 0
        for i in x:
            try:
                first = int(i)
                break
            except Exception:
                pass
        for i in x[::-1]:
            try:
                last = int(i)
                break
            except Exception:
                pass
        number = first * 10 + last
        print(number)
        sum += number

    print(sum)