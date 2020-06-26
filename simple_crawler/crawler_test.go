package simple_crawler

import (
	"fmt"
	"net/http"
	"regexp"
	"sync"
	"testing"
	"time"
)

func Test_CrawlerPlatform(t *testing.T) {
	list := []JobInfo{
		{PlatformName: "今日头条", Url: "https://www.toutiao.com/a6842087368221000200"},
		{PlatformName: "搜狐网", Url: "https://www.sohu.com/a/404042582_496421?spm=smpc.mil-home.feed.1.1593052479609Cw5rk4s"},
		{PlatformName: "网易新闻", Url: "https://news.163.com/20/0625/09/FFV6B8VQ0001899O.html"},
		{PlatformName: "新浪财经", Url: "https://finance.sina.com.cn/roll/2020-06-25/doc-iirczymk8863563.shtml"},
		{PlatformName: "新浪微博", Url: "https://s.weibo.com/weibo?q=%23%E4%BF%84%E7%BD%97%E6%96%AF%E7%BA%A2%E5%9C%BA%E9%98%85%E5%85%B5%E5%BC%8F%E7%9B%B4%E6%92%AD%23&from=default"},
		{PlatformName: "凤凰新闻", Url: "https://news.ifeng.com/c/7iyXEy1Ptaf"},
		{PlatformName: "东方财富网", Url: "http://finance.eastmoney.com/a/202006251533566237.html"},
		{PlatformName: "海峡网", Url: "http://www.hxnews.com/news/fj/fz/202006/23/1907292.shtml"},
		{PlatformName: "河北新闻网", Url: "http://hebei.hebnews.cn/2020-06/25/content_7959208.htm"},
		{PlatformName: "It之家", Url: "https://www.ithome.com/0/494/604.htm"},
		{PlatformName: "百度", Url: "https://mbd.baidu.com/newspage/data/landingsuper?context=%7B%22nid%22%3A%22news_9872723525150545336%22%7D&n_type=0&p_from=1"},
	}

	startTime := time.Now()
	CrawlerPlatform(list, printHandleResult)
	diff := time.Now().Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), " s")
}

func Benchmark_CrawlerPlatform(b *testing.B) {
	b.StopTimer()
	list := []JobInfo{
		{PlatformName: "今日头条", Url: "https://www.toutiao.com/a6842087368221000200"},
		{PlatformName: "搜狐网", Url: "https://www.sohu.com/a/404042582_496421?spm=smpc.mil-home.feed.1.1593052479609Cw5rk4s"},
		{PlatformName: "网易新闻", Url: "https://news.163.com/20/0625/09/FFV6B8VQ0001899O.html"},
		{PlatformName: "新浪财经", Url: "https://finance.sina.com.cn/roll/2020-06-25/doc-iirczymk8863563.shtml"},
		{PlatformName: "新浪微博", Url: "https://s.weibo.com/weibo?q=%23%E4%BF%84%E7%BD%97%E6%96%AF%E7%BA%A2%E5%9C%BA%E9%98%85%E5%85%B5%E5%BC%8F%E7%9B%B4%E6%92%AD%23&from=default"},
		{PlatformName: "凤凰新闻", Url: "https://news.ifeng.com/c/7iyXEy1Ptaf"},
		{PlatformName: "东方财富网", Url: "http://finance.eastmoney.com/a/202006251533566237.html"},
		{PlatformName: "海峡网", Url: "http://www.hxnews.com/news/fj/fz/202006/23/1907292.shtml"},
		{PlatformName: "河北新闻网", Url: "http://hebei.hebnews.cn/2020-06/25/content_7959208.htm"},
		{PlatformName: "It之家", Url: "https://www.ithome.com/0/494/604.htm"},
		{PlatformName: "百度", Url: "https://mbd.baidu.com/newspage/data/landingsuper?context=%7B%22nid%22%3A%22news_9872723525150545336%22%7D&n_type=0&p_from=1"},
	}
	wg := &sync.WaitGroup{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			CrawlerPlatform(list, printHandleResult)
		}()
	}
	wg.Wait()
}

func printHandleResult(workResult *WorkResult) {
	urlContentMap := make(map[string]string)
	urlContentMap["https://www.toutiao.com/a6842087368221000200"] = "新冠病毒抗体研究获重大突破"
	urlContentMap["https://www.sohu.com/a/404042582_496421?spm=smpc.mil-home.feed.1.1593052479609Cw5rk4s"] = "美国疫情依然汹涌"
	urlContentMap["https://news.163.com/20/0625/09/FFV6B8VQ0001899O.html"] = "美国医院床位告急"
	urlContentMap["https://finance.sina.com.cn/roll/2020-06-25/doc-iirczymk8863563.shtml"] = "国美电器创始人黄光裕获假释"
	urlContentMap["https://s.weibo.com/weibo?q=%23%E4%BF%84%E7%BD%97%E6%96%AF%E7%BA%A2%E5%9C%BA%E9%98%85%E5%85%B5%E5%BC%8F%E7%9B%B4%E6%92%AD%23&from=default"] = "俄罗斯红场阅兵式直播"
	urlContentMap["https://news.ifeng.com/c/7iyXEy1Ptaf"] = "国安部原副部长马建被判无期"
	urlContentMap["http://finance.eastmoney.com/a/202006251533566237.html"] = `券商股进入估值“攻势周期”`
	urlContentMap["http://www.hxnews.com/news/fj/fz/202006/23/1907292.shtml"] = "保险合同非本人签名后续"
	urlContentMap["http://hebei.hebnews.cn/2020-06/25/content_7959208.htm"] = "新冠病毒抗体研究获重大突破"
	urlContentMap["http://hebei.hebnews.cn/2020-06/25/content_7959208.htm"] = "报告新型冠状病毒肺炎"
	urlContentMap["https://www.ithome.com/0/494/604.htm"] = `苹果 iOS/iPadOS 14“新面孔”`
	urlContentMap["https://mbd.baidu.com/newspage/data/landingsuper?context=%7B%22nid%22%3A%22news_9872723525150545336%22%7D&n_type=0&p_from=1"] = "CIM新基建，CityBase的定位与腾讯的版图"

	if workResult.HttpStatus != http.StatusOK {
		fmt.Printf("%s平台，http状态为:%d,查找结果为：%v\n", workResult.Job.PlatformName, workResult.HttpStatus, false)
		return
	}

	//查找内容
	re := regexp.MustCompile(urlContentMap[workResult.Job.Url])
	findStr := re.FindString(string(workResult.Content))
	if len(findStr) > 0 {
		fmt.Printf("%s平台，http状态为:%d,查找结果为：%v\n", workResult.Job.PlatformName, workResult.HttpStatus, true)
		return
	}

	fmt.Printf("%s平台，http状态为:%d,查找结果为：%v\n", workResult.Job.PlatformName, workResult.HttpStatus, false)
	return
}
