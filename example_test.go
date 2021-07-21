package i18n

import "testing"

func TestAdd(t *testing.T) {
	i18nClient, err := NewI18n("./lang", "en")

	if err != nil {
		t.Error("err")
	}

	str := i18nClient.T("人生最幸福的时期，是什么时候？")

	t.Log(str)

	str1 := i18nClient.T("%s年欧洲杯冠军是%s", "2021", "Italy")

	t.Log(str1)

	str2 := i18nClient.TOption("人生最幸福的时期，是什么时候？", "ko")

	t.Log(str2)

	str3 := i18nClient.TOption("%s年欧洲杯冠军是%s", "ko", "2021", "이탈리아")

	t.Log(str3)

}
