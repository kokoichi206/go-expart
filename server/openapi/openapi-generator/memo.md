``` sh
openapi-generator generate -i ../petstore.yml -g go-gin-server -o ./gen

# local で使ってるバージョン。　
$ openapi-generator --version
openapi-generator-cli 7.0.0
  commit : c37fa8a
  built  : -999999999-01-01T00:00:00+18:00
  source : https://github.com/openapitools/openapi-generator
  docs   : https://openapi-generator.tech/
```




```
mvn -v
Apache Maven 3.9.5 (57804ffe001d7215b5e7bcb531cf83df38f93546)
Maven home: /opt/homebrew/Cellar/maven/3.9.5/libexec
Java version: 11.0.21, vendor: Amazon.com Inc., runtime: /Users/kokoichi/Library/Java/JavaVirtualMachines/corretto-11.0.21/Contents/Home
Default locale: en_JP, platform encoding: UTF-8
OS name: "mac os x", version: "13.0.1", arch: "x86_64", family: "mac"


mvn clean install
```
