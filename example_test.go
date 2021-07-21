package i18n

import "testing"

func TestAdd(t *testing.T) {
	i18nClient, err := NewI18n("./lang", "en")

	if err != nil {
		t.Error("err")
	}

	str := i18nClient.T("人生最幸福的时期，是什么时候")

	t.Log(str)

	str1 := i18nClient.T("小博今年%d岁了", 18)

	t.Log(str1)

	str2 := i18nClient.T("%s年欧洲杯冠军是%s", "2021", "Italy")

	t.Log(str2)

}