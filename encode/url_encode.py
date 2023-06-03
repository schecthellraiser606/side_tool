import sys

def urlencode(string):
    urlencode = ""
    
    for char in string:
      decimal = ord(char)
      urlencode += "%" + hex(decimal)[2:]
      
    return urlencode
  
encode=urlencode(sys.argv[1])
print(encode)

# Usage: python url_encode.py "code to encode"
# python url_encode.py "ping -c 10.10.14.79 "