import zipfile

files = {
    "初音ミクが好きです.txt": "oo ee oo~",
    "サブフォルダ/readme.txt": "Hello, world!",
    "Shift-JISが大嫌いです.txt": "Shift-JIS stinks"
}

def main():
    with zipfile.ZipFile("shiftjis.zip", "w") as f:
        for name, content in files.items():
            info = zipfile.ZipInfo(name.encode("cp932").decode("latin1")) # Shift-JIS
            info.flag_bits &= ~0x800 # clear UTF-8 flag
            f.writestr(info, content)

if __name__ == "__main__": main()
