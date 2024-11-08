import time, os, csv, tqdm, dns.resolver, requests, numpy, prettytable

csv_url = (
    "https://github.com/Kakune55/DNSspeedtest/releases/latest/download/dnslist.csv"
)
domain = "baidu.com"
testRepeat = 10  # 必须大于等于1 且为整数
csvPath = "dnslist.csv"


def delayTest(nameserver: str, repeat: int) -> list:
    "循环测试"
    resolver = dns.resolver.Resolver()
    resolver.nameservers = [nameserver]
    resultList = []
    for i in range(repeat):
        try:
            start = time.time()
            resolver.resolve(qname=domain, rdtype="A")
            end = time.time()
        except:
            return 9999
        resultList.append((end - start) * 1000)
    return resultList


def localRead():
    return csv.reader(open(csvPath, encoding="utf-8"))


def OnlineRead():
    return csv.reader(requests.get(csv_url,timeout=1).content.decode("utf-8").splitlines())


def formatTime(input: float) -> str:
    if input < 100:
        return f"{format(input,'.2f')}"
    elif input < 1000:
        return f"{format(input,'.1f')}"
    else:
        return f"{format(input,'.0f')}"


def runTest(csvPath: str, repeat: int):
    delayTable = []
    with tqdm.tqdm(total=len(open(csvPath, encoding="utf-8").readlines())) as pbar:
        for row in csv_reader:
            for i in range(len(row)):
                if i == 0:
                    pbar.desc = row[0]
                elif row[i] == "":
                    continue
                else:
                    delay = delayTest(row[i], repeat)
                    delayTable.append(
                        [
                            numpy.mean(delay),
                            numpy.std(delay),
                            numpy.amax(delay),
                            numpy.amin(delay),
                            row[0],
                            row[i],
                        ]
                    )
            pbar.update(1)

    print("结果如下\n")
    resultTable = prettytable.PrettyTable(
        ["平均值", "标准差", "最大值", "最小值", "ID", "IP"]
    )
    # resultTable = [["平均值\t标准差\t最大值\t最小值\t单位ms"]]
    # for i in sorted(delayTable, key=lambda x: x[0]):
    #     resultTable.append(i)
    # for i in resultTable:  # 格式化输出
    #     for j in i:
    #         if isinstance(j, str):
    #             print(j)
    #         else:
    #             print(formatTime(j), end="\t")
    # return resultTable
    for i in sorted(delayTable, key=lambda x: x[0]):
        resultTable.add_row(
            [
                formatTime(i[0]),
                formatTime(i[1]),
                formatTime(i[2]),
                formatTime(i[3]),
                i[4],
                i[5],
            ]
        )
    return resultTable


if __name__ == "__main__":
    print("初始化")
    try:
        print("尝试从在线加载DNS列表", end="", flush=True)
        csv_reader = OnlineRead()
        print("-成功")
    except:
        print("-失败")
        try:
            print("尝试从本地加载DNS列表", end="", flush=True)
            csv_reader = localRead()
            print("-成功")
        except:
            print("-失败")
            input()
            exit()

    menuSwitch = input("\n1.快速测试\n2.平均值测试\n\n输入你的选项:")
    if menuSwitch == "1":
        print(f"开始快速测试")
        resultTable = runTest(csvPath, 1)
        print(resultTable)
    elif menuSwitch == "2":
        print(f"开始测试 循环次数：{testRepeat}")
        resultTable = runTest(csvPath, testRepeat)
        print(resultTable)
    else:
        print("未知选项")
    if input("测试完成!按回车键退出 输入 s 将结果保存为csv文件:\n").upper() == "S":
        try:
            if input("使用UTF-8编码?默认使用GBK Y/(N)").upper() == "Y":
                with open(
                    f"{os.path.join(os.path.expanduser('~'),'Desktop')}/result.csv",
                    "wt",
                    newline="",
                    encoding="utf-8",
                ) as f:
                    f.write(resultTable.get_csv_string())
                    print(
                        "文件已保存至 ",
                        f"{os.path.join(os.path.expanduser('~'),'Desktop')}/result.csv",
                    )
            else:
                with open(
                    f"{os.path.join(os.path.expanduser('~'),'Desktop')}/result.csv",
                    "wt",
                    newline="",
                    encoding="gbk",
                ) as f:
                    f.write(resultTable.get_csv_string())
        except Exception as e:
            print("保存文件时发生错误")
            print(e)    
        input()
    exit()
