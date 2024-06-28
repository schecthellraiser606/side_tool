import base64
import chardet

def detect_encoding(file_path):
    with open(file_path, 'rb') as file:
        raw_data = file.read()
        result = chardet.detect(raw_data)
        return result['encoding']

def encode_file_lines_to_base64(input_file_path, output_file_path):
    try:
        encoding = detect_encoding(input_file_path)
        with open(input_file_path, 'r', encoding=encoding) as input_file:
            with open(output_file_path, 'w', encoding='utf-8') as output_file:
                for line in input_file:
                    encoded_line = base64.b64encode(line.encode(str(encoding))).decode('utf-8')
                    output_file.write(encoded_line + '\n')
        print(f"Successfully encoded and saved to {output_file_path}")
    except Exception as e:
        print(f"An error occurred: {e}")

# 使用例
input_file_path = '/usr/share/wordlists/rockyou.txt'
output_file_path = '/root/work/passlist.txt'
encode_file_lines_to_base64(input_file_path, output_file_path)