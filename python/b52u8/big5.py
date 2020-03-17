import os
import time
import logging
file_filter = 'apscau.unl'
input_dir = 'd:\\convert\\big5\\'
output_dir = 'd:\\convert\\utf8\\'
input_file = ''
output_file = ''
logging.basicConfig(filename= input_dir + '\\log.txt', level=logging.INFO,  filemode='w') 

if not os.path.isdir(output_dir) :
    os.mkdir(output_dir)
for exp_name in os.listdir(input_dir):
  if 'erp.exp' in exp_name:
    out_exp_dir = output_dir + exp_name + '\\'
    in_exp_dir = input_dir + exp_name + '\\'
    if not os.path.isdir(out_exp_dir) :
        os.mkdir(out_exp_dir)
    for filename in os.listdir(in_exp_dir):
      if file_filter in filename :        
        localtime = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime()) 
        input_file = in_exp_dir + filename
        output_file = out_exp_dir + filename
        #print ('Loading: %s' % input_file)
        print (localtime + ' : Loading: ' + input_file)
        logging.info(localtime + ' : Loading: ' + input_file)
        
        # errors='replace' replaces bad sequences with whitespace sequence
        # errors='ignore' removes bad sequences completely
        # if errors param is not set, the byte location of the first bad sequence will be printed
        """with open(input_file, encoding='big5', errors='replace') as f:
          lines = f.readlines()        
          converted = [x.encode('utf8', 'strict') for x in lines]        
        with open(output_file, 'wb') as f:
          for x in converted:
            #f.write(x.replace('\r',''))
            f.write(x)"""
        with open(input_file, encoding="big5", errors="replace") as fin:
          with open(output_file, "wb") as fout:
            for x in fin:
              #x = x.replace('\r','') 
              fout.write(x.encode("utf8", "strict"))
