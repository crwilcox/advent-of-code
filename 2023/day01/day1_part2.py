with open("input") as input:
    sum = 0
    spelled = [
        'zero', 'one', 'two', 'three', 'four',
        'five', 'six', 'seven', 'eight', 'nine']
    for x in input.readlines():
        first = 0
        last = 0
        for k, i in enumerate(x):
            try:
                first = int(i)
                break
            except Exception:
                found = False
                for idx, spell in enumerate(spelled):
                    if x[k:].startswith(spell):
                        first = idx
                        found = True
                        break
                if found:
                    break

        for k, i in enumerate(x[::-1]):
            k = len(x)-k-1
            try:
                last = int(i)
                break
            except Exception:
                found = False
                for idx, spell in enumerate(spelled):
                    if x[k:].startswith(spell):
                        last = idx
                        found = True
                        break
                if found:
                    break
                

        number = first * 10 + last
        sum += number

print(f"🎄 Part 2 🎁: {sum}")
