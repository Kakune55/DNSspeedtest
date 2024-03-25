import time, csv, tqdm, dns.resolver

domain = "python.org"
testRepeat = 5 #必须大于等于1 且为整数
csvPath = "dnslist.csv"

def delayTest(nameserver:str,repeat:int):
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
try:
    csv_reader = csv.reader(open(csvPath))
except:
    print("读取dns列表失败")
    input()
    exit()

mindelay = 2000
minDelayNameserver = ""
with tqdm.tqdm(total=len(open(csvPath).readlines())) as pbar:
    for row in csv_reader:
        for i in range(len(row)):
            if i == 0:
                print("\n",row[0])
            elif row[i] == "":
                continue
            else:
                print(row[i],end="\t\t")
                delay = delayTest(row[i],5)
                print(format(delay,'.4f'),"ms")
                if delay < mindelay:
                    mindelay = delay
                    minDelayNameserver = row[0] + " : " +row[i]
        pbar.update(1)
print("---------------")
print(f"MinDelayNameServer is {minDelayNameserver}\n{format(mindelay,'.4f')}ms")
input()


              



