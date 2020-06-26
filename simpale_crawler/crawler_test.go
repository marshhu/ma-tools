package simpale_crawler

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_CrawlerPlatform(t *testing.T) {
	list := []JobInfo{
		//{PlatformName: "今日头条", PlatformUrl: "https://www.toutiao.com", Url: "https://www.toutiao.com/a6842087368221000200/", FindContent: "新冠病毒抗体研究获重大突破"},
		{PlatformName: "搜狐网", PlatformUrl: "https://www.sohu.com", Url: "https://www.sohu.com/a/404042582_496421?spm=smpc.mil-home.feed.1.1593052479609Cw5rk4s", FindContent: "美国疫情依然汹涌"},
		{PlatformName: "网易新闻", PlatformUrl: "https://news.163.com", Url: "https://news.163.com/20/0625/09/FFV6B8VQ0001899O.html", FindContent: "美国医院床位告急"},
		{PlatformName: "新浪财经", PlatformUrl: "https://finance.sina.com.cn", Url: "https://finance.sina.com.cn/roll/2020-06-25/doc-iirczymk8863563.shtml", FindContent: "国美电器创始人黄光裕获假释"},
		{PlatformName: "新浪微博", PlatformUrl: "https://s.weibo.com/weibo", Url: "https://s.weibo.com/weibo?q=%23%E4%BF%84%E7%BD%97%E6%96%AF%E7%BA%A2%E5%9C%BA%E9%98%85%E5%85%B5%E5%BC%8F%E7%9B%B4%E6%92%AD%23&from=default", FindContent: "俄罗斯红场阅兵式直播"},
		{PlatformName: "凤凰新闻", PlatformUrl: "https://news.ifeng.com ", Url: "https://news.ifeng.com/c/7iyXEy1Ptaf ", FindContent: "国安部原副部长马建被判无期"},
		{PlatformName: "东方财富网", PlatformUrl: "http://finance.eastmoney.com", Url: "http://finance.eastmoney.com/a/202006251533566237.html", FindContent: `券商股进入估值“攻势周期”`},
		{PlatformName: "海峡网", PlatformUrl: "http://www.hxnews.com", Url: "http://www.hxnews.com/news/fj/fz/202006/23/1907292.shtml", FindContent: "保险合同非本人签名后续"},
		{PlatformName: "河北新闻网", PlatformUrl: "http://hebei.hebnews.cn", Url: "http://hebei.hebnews.cn/2020-06/25/content_7959208.htm", FindContent: "报告新型冠状病毒肺炎"},
		{PlatformName: "It之家", PlatformUrl: "https://www.ithome.com", Url: "https://www.ithome.com/0/494/604.htm", FindContent: `苹果 iOS/iPadOS 14“新面孔”`},
		{PlatformName: "百度", PlatformUrl: "https://mbd.baidu.com/", Url: "https://mbd.baidu.com/newspage/data/landingsuper?context=%7B%22nid%22%3A%22news_9872723525150545336%22%7D&n_type=0&p_from=1", FindContent: "CIM新基建，CityBase的定位与腾讯的版图"},
	}

	startTime := time.Now()
	CrawlerPlatform(list)
	diff := time.Now().Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), " s")
}

func Benchmark_CrawlerPlatform(b *testing.B) {
	b.StopTimer()
	list := []JobInfo{
		//{PlatformName: "今日头条", PlatformUrl: "https://www.toutiao.com", Url: "https://www.toutiao.com/a6842087368221000200/", FindContent: "新冠病毒抗体研究获重大突破"},
		{PlatformName: "搜狐网", PlatformUrl: "https://www.sohu.com", Url: "https://www.sohu.com/a/404042582_496421?spm=smpc.mil-home.feed.1.1593052479609Cw5rk4s", FindContent: "美国疫情依然汹涌"},
		{PlatformName: "网易新闻", PlatformUrl: "https://news.163.com", Url: "https://news.163.com/20/0625/09/FFV6B8VQ0001899O.html", FindContent: "美国医院床位告急"},
		{PlatformName: "新浪财经", PlatformUrl: "https://finance.sina.com.cn", Url: "https://finance.sina.com.cn/roll/2020-06-25/doc-iirczymk8863563.shtml", FindContent: "国美电器创始人黄光裕获假释"},
		{PlatformName: "新浪微博", PlatformUrl: "https://s.weibo.com/weibo", Url: "https://s.weibo.com/weibo?q=%23%E4%BF%84%E7%BD%97%E6%96%AF%E7%BA%A2%E5%9C%BA%E9%98%85%E5%85%B5%E5%BC%8F%E7%9B%B4%E6%92%AD%23&from=default", FindContent: "俄罗斯红场阅兵式直播"},
		{PlatformName: "凤凰新闻", PlatformUrl: "https://news.ifeng.com ", Url: "https://news.ifeng.com/c/7iyXEy1Ptaf ", FindContent: "国安部原副部长马建被判无期"},
		{PlatformName: "东方财富网", PlatformUrl: "http://finance.eastmoney.com", Url: "http://finance.eastmoney.com/a/202006251533566237.html", FindContent: `券商股进入估值“攻势周期”`},
		{PlatformName: "海峡网", PlatformUrl: "http://www.hxnews.com", Url: "http://www.hxnews.com/news/fj/fz/202006/23/1907292.shtml", FindContent: "保险合同非本人签名后续"},
		{PlatformName: "河北新闻网", PlatformUrl: "http://hebei.hebnews.cn", Url: "http://hebei.hebnews.cn/2020-06/25/content_7959208.htm", FindContent: "报告新型冠状病毒肺炎"},
		{PlatformName: "It之家", PlatformUrl: "https://www.ithome.com", Url: "https://www.ithome.com/0/494/604.htm", FindContent: `苹果 iOS/iPadOS 14“新面孔”`},
		{PlatformName: "百度", PlatformUrl: "https://mbd.baidu.com/", Url: "https://mbd.baidu.com/newspage/data/landingsuper?context=%7B%22nid%22%3A%22news_9872723525150545336%22%7D&n_type=0&p_from=1", FindContent: "CIM新基建，CityBase的定位与腾讯的版图"},
	}
	wg := &sync.WaitGroup{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			CrawlerPlatform(list)
		}()
	}
	wg.Wait()
}
