from konlpy.tag import Okt
import csv
 
JVM_PATH = '/Library/Java/JavaVirtualMachines/zulu-11.jdk/Contents/Home/bin/java'
okt = Okt(jvmpath=JVM_PATH)

reader = csv.reader(open('search.csv', 'r'))
writer = csv.writer(open('searchMorpheme.csv', 'w'))

for row in reader:
    row.append(okt.nouns(row[0]))
    writer.writerow(row)