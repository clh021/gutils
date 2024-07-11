package browser

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/clh021/gutils/ufunc/envutil"
)

type ConfigWithLock struct {
	sync.Mutex
	Items       []BrowserItem
	AgentGetCnt int
	ItemCnt     int
}

func GetBrowsers() *[]BrowserItem {
	items := filterConf(getConf())
	confWithLock := &ConfigWithLock{Items: items, AgentGetCnt: 0, ItemCnt: len(items)}

	port, err := envutil.GetFreePort()
	if err != nil {
		log.Fatalln(err)
	}
	defaultUrl := fmt.Sprintf("http://127.0.0.1:%d?b=defaultbrowser", port)
	UserAgentServe(port, &confWithLock.Items, &confWithLock.AgentGetCnt)

	waitSecond := 330
	timeoutTimer := time.NewTimer(time.Duration(waitSecond) * time.Second)
	defer timeoutTimer.Stop()

	for i := range confWithLock.Items {
		go sendUserAgentRequest(port, &confWithLock.Items[i])
		getVersion(&confWithLock.Items[i])
	}

	checkUserAgent := func() {
		for {
			select {
			case <-timeoutTimer.C:
				log.Printf("浏览器内核信息采集超时：%s\n", defaultUrl)
				return
			default:
				log.Println("浏览器内核信息采集中...", confWithLock.AgentGetCnt, "/", confWithLock.ItemCnt)
				if confWithLock.AgentGetCnt == confWithLock.ItemCnt {
					log.Println("浏览器内核信息采集完成")
					return
				}
				time.Sleep(6 * time.Second)
			}
		}
	}
	checkUserAgent()
	checkDefault(&confWithLock.Items)
	return &confWithLock.Items
}

func getVersion(b *BrowserItem) (e error) {
	out, e := execGetCmdStdoutWithEnv(b.Bin, "--version")
	if e != nil {
		return
	}
	b.CmdVer = strings.TrimSpace(string(out))

	cmdVer := strings.TrimPrefix(b.CmdVer, "360Browser") // 针对360浏览器的数字调整

	b.Version, e = regVer(cmdVer, b.CmdReg)
	return
}

func regVer(str, reg string) (string, error) {
	pattern := regexp.MustCompile(reg)
	// 查找匹配项
	matches := pattern.FindStringSubmatch(str)
	if len(matches) < 2 {
		return "", fmt.Errorf("无法提取 版本号")
	}
	// 提取版本号（matches[1] 是第一个括号内的匹配内容）
	ver := matches[1]
	return ver, nil
}

func filterConf(items []BrowserItem) []BrowserItem {
	validItems := []BrowserItem{}
	for _, item := range items {
		if item.Bin != "" {
			if _, err := exec.LookPath(item.Bin); err == nil {
				// fmt.Println("append:", item.Bin)
				validItems = append(validItems, item)
			}
		}
	}
	return validItems
}

func execGetCmdStdoutWithEnv(pathBin string, args ...string) ([]byte, error) {
	cmd := exec.Command(pathBin, args...)
	cmd.Stdin = os.Stdin
	return cmd.Output()
}

func checkDefault(Conf *[]BrowserItem) (e error) {
	conf := *Conf
	out, e := execGetCmdStdoutWithEnv("xdg-mime", "query", "default", "x-scheme-handler/http")
	if e != nil {
		return
	}
	for i, c := range conf {
		if c.Desktop == strings.TrimSpace(string(out)) {
			conf[i].IsDefault = true
		}
	}
	return
}
