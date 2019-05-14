import os
import time

common1 = "sudo lsof -p 3190"
common = "sudo lsof -p 3190 | wc -l"
def test():
    result1 = os.popen(common1).read()
    result = os.popen(common).read()

    with open('/home/centos/lsof.log','w+') as f:
        for i in range(100):
            time.sleep(2)
            a=time.strftime("%Y-%m-%d %H:%M:%S") + '  ' + result
            f.write(a)
            f.write(result1)

test()