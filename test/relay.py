


def realy():
    with open('Relay2019052122.txt','r') as f:
        one = f.readlines()

    sum = 0
    for i in one:
        res = i.split()
        sum = int(res[5]) + sum
        print(res[5])

    print(sum)


# realy()


def getRelayNumber():
    with open('Relay2019052122.txt','r') as f:
        one = f.readlines()

    sum = 0
    uidList = []
    for i in one:
        res = i.split()
        sum = int(res[5]) + sum
        if str(res[5]) != "0":
            # print(res[5])
            uidList.append(res[2])

    newList = set(uidList)
    # print("去重后的uid数：{}".format(len(newList)))
    # print(newList)
    print("产生流量的session的uid数{}".format(len(newList)))
    print("产生的session数为：{}".format(len(one)))
    print("流入流出产生的总流量为：{} byte".format(sum))

getRelayNumber()