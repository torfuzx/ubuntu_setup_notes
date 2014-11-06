Source Code Pro
---------------

```bash
mkdir /tmp/adodefont
cd /tmp/adodefont
wget http://downloads.sourceforge.net/project/sourcecodepro.adobe/SourceCodePro_FontsOnly-1.017.zip
unzip SourceCodePro_FontsOnly-1.017.zip
mkdir -p ~/.fonts
cp SourceCodePro_FontsOnly-1.017/OTF/*.otf ~/.fonts
fc-cache -f -v
```


Inconsolata
-----------

```
sudo apt-get install fonts-inconsolata
```

Monaco
------

```
curl -kL https://raw.github.com/cstrap/monaco-font/master/install-font-ubuntu.sh | bash
```


Microsoft Yahei
---------------

```
  wget -O get-fonts.sh.zip http://files.cnblogs.com/DengYangjun/get-fonts.sh.zip
  unzip -o get-fonts.sh.zip 1>/dev/null
  chmod a+x get-fonts.sh
  ./get-fonts.sh
```
