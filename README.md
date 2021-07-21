## 使用

### 初始化语言翻译

```
package main

import (
	"github.com/wo4zhuzi/i18n"
	"log"
)

func main() {
	i18nClient, err := i18n.NewI18n("./lang", "en")
	if err != nil {
		log.Fatalln(err)
	}

	str := i18nClient.T("人生最幸福的时期，是什么时候？")
	log.Println(str)

	str1 := i18nClient.T("%s年欧洲杯冠军是%s", "2021", "Italy")
	log.Println(str1)
}
```

### 自定义语言翻译

```
package main

import (
	"github.com/wo4zhuzi/i18n"
	"log"
)

func main() {
	i18nClient, err := i18n.NewI18n("./lang", "en")
	if err != nil {
		log.Fatalln(err)
	}

	str := i18nClient.T("人生最幸福的时期，是什么时候？")
	log.Println(str)

	str1 := i18nClient.T("%s年欧洲杯冠军是%s", "2021", "Italy")
	log.Println(str1)

	//自定义语言翻译
	str2 := i18nClient.TOption("人生最幸福的时期，是什么时候？", "ko")
	log.Println(str2)

	str3 := i18nClient.TOption("%s年欧洲杯冠军是%s", "ko", "2021", "Italy")
	log.Println(str3)

}
```