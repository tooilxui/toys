#sqlplus sys/kenda491@192.1.1.188/MESDBD as sysdba
export MMSG=/tmp/$$.mail
export ADDR="saxon.pad@gmail.com"
export ORACLE_HOME=/u01/app/oracle/product/12.1.0/dbhome_1
export ORACLE_SID=orcl
file=/tmp/${$}_`date +%Y%m%d`
sqlplus -s "/as sysdba" << .eof > $file
set pages 0
select df.tablespace_name tspace,
round(sum(fs.bytes)/(df.bytes) * 100) "%_free",
round(sum(fs.bytes)/(1024*1024)) free_ts_size,
df.bytes/(1024*1024) tot_ts_size
from dba_free_space fs, (select tablespace_name, sum(bytes) bytes from dba_data_files
group by tablespace_name ) df
where fs.tablespace_name = df.tablespace_name
group by df.tablespace_name, df.bytes
having round(sum(fs.bytes)/(df.bytes) * 100) < 20;
exit;
.eof
#list all datafile below
egrep "SOE|APEX|SYSAUX|SYSTEM|TEMP|UNDOTBS1|USERS|AUDIT_TBS|APP_ENCRYPTED" /tmp/${$}_`date +%Y%m%d` >/tmp/table${ORACLE_SID}.txt
check_stat=`cat /tmp/table${ORACLE_SID}.txt|wc -l`;
oracle_num=`expr $check_stat`
if [ $oracle_num -ne 0 ]
then
echo "tablespace less than 20% for $ORACLE_SID : $oracle_num" > $MMSG
mail -s "TABLESPACE WITH LESS THAN 20% FREE SPACE" $ADDR < $MMSG
mail -s "TABLESPACE WITH LESS THAN 20% FREE SPACE IN DATABASE $ORACLE_SID" $ADDR < /tmp/table$ORACLE_SID.txt
fi
rm -f $MMSG > /dev/null 2>&1
rm $file