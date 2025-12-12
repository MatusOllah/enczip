import zipfile

files = {
    "ľščťžýáíé.txt": "skúška diakritiky",
    "zložka/readme.txt": "Hello, world!",
    "kajšmentke.txt": "kožmeker"
}

def main():
    with zipfile.ZipFile("cp1250.zip", "w") as f:
        for name, content in files.items():
            info = zipfile.ZipInfo(name.encode("cp1250").decode("latin1")) # Windows-1250
            info.flag_bits &= ~0x800 # clear UTF-8 flag
            f.writestr(info, content)

if __name__ == "__main__": main()
