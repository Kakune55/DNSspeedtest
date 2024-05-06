import time, csv, tqdm, dns.resolver, requests

csv_url = "https://github.com/Kakune55/DNSspeedtest/releases/latest/download/dnslist.csv"
domain = "baidu.com"
testRepeat = 10 #必须大于等于1 且为整数
csvPath = "dnslist.csv"

def delayTest(nameserver:str,repeat:int): 
    "循环测试"
    resolver = dns.resolver.Resolver()
    resolver.nameservers = [nameserver]
    alltime = 0
    for i in range(repeat):
        try:
            start = time.time()
            resolver.resolve(qname=domain, rdtype='A')
            end = time.time()
        except:
            return 9999
        alltime += end - start
    return alltime/repeat*1000


def localRead():
    return csv.reader(open(csvPath,encoding='utf-8'))

def OnlineRead():
    return csv.reader(requests.get(csv_url).content.decode('utf-8').splitlines())

def runTest(csvPath:str,repeat:int):
    delayTable = []
    with tqdm.tqdm(total=len(open(csvPath,encoding='utf-8').readlines())) as pbar:
        for row in csv_reader:
            for i in range(len(row)):
                if i == 0:
                    pbar.desc = row[0]
                elif row[i] == "":
                    continue
                else:
                    delay = delayTest(row[i],repeat)
                    delayTable.append((row[0] + " " + row[i], delay))
            pbar.update(1)

    print("结果如下\n")
    for i in sorted(delayTable, key=lambda x:x[1]): #格式化输出
        if i[1] < 100:
            print(f"{format(i[1],'.2f')}ms\t{i[0]}")
        elif i[1] < 1000:
            print(f"{format(i[1],'.1f')}ms\t{i[0]}")
        else:
            print(f"{format(i[1],'.0f')} ms\t{i[0]}")

if __name__ == '__main__':
    print("初始化")
    try:
        print("尝试从在线加载DNS列表",end="",flush=True)
        csv_reader = OnlineRead()
        print("-成功")
    except:
        print("-失败")
        try:
            print("尝试从本地加载DNS列表",end="",flush=True)
            csv_reader = localRead()
            print("-成功")
        except:
            print("-失败")
            input()
            exit()
    
    menuSwitch = input("\n1.快速测试 2.平均值测试\n输入你的选项:")
    if menuSwitch == "1":
        print(f"开始快速测试")
        runTest(csvPath, 1)
    elif menuSwitch == "2":
        print(f"开始测试 循环次数：{testRepeat}")
        runTest(csvPath, testRepeat)
    else:
        print("未知选项")
    input("测试完成!按回车键退出")
    exit()
    