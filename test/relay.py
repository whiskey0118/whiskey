import re


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

# getRelayNumber()

def relayPercent():
    with open(r'F:\我的\test\tutk\Dataflow201905.txt','r') as f:
        one = f.readlines()

    sum = 0
    for i in one:
        res = i.split()
        if str(res[5]) != "0":
            sum = int(res[4]) - int(res[3]) + sum

    print(sum)


def getStaSession(day):
    #获取connectResu中去除uid数
    with open(r'F:\我的\test\tutk\Statistics20190402.txt','r') as f:
        total = f.read()

    patt = re.compile(r"\[{}.*".format(day))
    dayListTotal = re.findall(patt,str(total))
    dayList = []
    for i in dayListTotal:
        if "3600" in i.split("|")[2]:
            dayList.append(i)

    pattTotBSend =re.compile(r'TotBSend:{\d+}')
    pattRelayBSend = re.compile(r'RelayBSend:{\d+}')
    pattNumP2PReq = re.compile(r"NumP2PReq:{\d+}")
    pattNumRelayReq = re.compile(r"NumRelayReq:{\d+}")
    pattNumber = re.compile(r"\d+")

    totalSend = 0
    totalRelaySend = 0
    totalNumP2PReq = 0
    totalNumRelayReq = 0
    # print(len(dayList))
    for i in dayList:
        totalSendList = pattTotBSend.findall(i)
        totalSendByte = pattNumber.findall(totalSendList[0])[0]
        totalSend = int(totalSendByte) + totalSend
        # print(totalSendByte)

        relayBSendList = pattRelayBSend.findall(i)
        a = pattNumber.findall(relayBSendList[0])[0]
        totalRelaySend = int(a)+totalRelaySend

        numP2PReqList = pattNumP2PReq.findall(i)
        b = pattNumber.findall(numP2PReqList[0])[1]
        totalNumP2PReq = int(b)+totalNumP2PReq

        numRelayReqList = pattNumRelayReq.findall(i)
        c = pattNumber.findall(numRelayReqList[0])[0]
        totalNumRelayReq = int(c)+totalNumRelayReq

        totalP2PSend = int(totalSend) - int(totalRelaySend)


    #一天时间发出的流量总和[0]
    #一天时间P2P发出的流量总和[1]
    #一天时间Relay发出的流量总和[2]
    #一天时间P2P请求书总和[3]
    #一天时间relay请求书总和[4]
    #一天时间请求总和[5]
    totalReq = totalNumP2PReq+totalNumRelayReq
    return [totalSend,totalP2PSend,totalRelaySend,totalNumP2PReq,totalNumRelayReq,totalReq]


def getDevUsgUid(day):
    #获取deviceUsg文件中的去重uid数
    dayList = []

    for hour in range(24):
        if hour <10 :
            fileName = r'F:\我的\test\tutk-0506\DeviceUsage{}0{}.txt'.format(day,hour)
        else:
            fileName = r'F:\我的\test\tutk-0506\DeviceUsage{}{}.txt'.format(day,hour)

        with open(fileName, 'r') as f:
            total = f.readlines()
        f.close()
        patt = re.compile('{.*}')
        uidList = []
        for i in total:
            res = i.split('|')
            uid = re.findall(patt, str(res[1]))[0].strip('{}')
            uidList.append(uid)

        newList = set(uidList)
        dayList += newList
        # 去重后的uid数
        # print("{}小时的uid数：{}".format(hour, len(newList)))
    # print("{}请求的总行数： {}".format(day,len(dayList)))
    print("{}号中发生过请求记录并去重后的uid数：{}".format(day,len(set(dayList))))
    return len(set(dayList))


def getData():
    #1000用户使用流量列表
    dataList = []

    #使用总流量
    dataTotalFlowList = []

    #一天请求数
    dataTotalReqList = []

    #一天发生请求的设备数
    dataTotalUidList = []

    for i in range(6,13):
        if i <10:
            day = "0{}".format(i)
            day2 = "2019-05-{}".format(day)
        else:
            day = i
            day2 = "2019-05-{}".format(i)

        sta = getStaSession(day2)
        devUsg = getDevUsgUid(day)
        # print("{}一天时间发出的流量总和： {} byte".format(day, sta[0]))
        # print("{}一天时间P2P发出的流量总和： {} byte".format(day, sta[1]))
        # print("{}一天时间Relay发出的流量总和： {} byte".format(day, sta[2]))
        # print("{}一天时间P2P请求数总和： {} ".format(day, sta[3]))
        # print("{}一天时间relay请求数总和： {} ".format(day, sta[4]))
        # print("{}一天时间请求数总和： {} ".format(day, sta[5]))

        # 一天时间发出的流量总和[0]
        # 一天时间P2P发出的流量总和[1]
        # 一天时间Relay发出的流量总和[2]
        # 一天时间P2P请求书总和[3]
        # 一天时间relay请求书总和[4]
        # 一天时间请求总和[5]

        # print("获取数据:")
        # print("{}每个p2p请求使用流量：{}".format(day, sta[1] / sta[3]))
        # print("{}每个relay请求使用流量：{}".format(day, sta[2] / sta[4]))
        # print("{}每个请求使用流量：{}".format(day, sta[0] / sta[5]))
        # print("{}每个uid发起请求数：{}".format(day, sta[5] / devUsg))
        print("{} 1000uid一天使用流量：{}".format(day, sta[0] / devUsg * 1000))
        print("\n")
        dataList.append(sta[0] / devUsg * 1000)
        dataTotalFlowList.append(sta[0])
        dataTotalReqList.append(sta[5])
        dataTotalUidList.append(devUsg)

    sumdata = 0
    TotalFlow=0
    TotalReq=0
    TotalUid=0

    for i in dataList:
        sumdata = i + sumdata

    for i in dataTotalFlowList:
        TotalFlow = i+TotalFlow

    for i in dataTotalReqList:
        TotalReq = i+TotalReq

    for i in dataTotalUidList:
        TotalUid = i+TotalUid

    perdata = sumdata/7
    print("5/6-5/12  1000设备平均一天使用流量：{}".format(sumdata/7))
    print("5/6-5/12  平均一天使用流量：{}".format(TotalFlow / 7))
    print("5/6-5/12  平均一天请求数：{}".format(TotalReq / 7))
    print("5/6-5/12  平均一天发生请求设备数：{}".format(TotalUid / 7))


getData()





