package core

import (
	"blblcd/client"
	"blblcd/model"
	"blblcd/utils"
	"fmt"
	"log/slog"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

func FetchVideoList(mid int64, page int, order string, cookie string) (videoList model.VideoListResponse, err error) {
	api := "https://api.bilibili.com/x/space/wbi/arc/search?"
	params := url.Values{}
	params.Set("mid", fmt.Sprint(mid))
	params.Set("order", order)
	params.Set("platform", "web")
	params.Set("pn", fmt.Sprint(page))
	params.Set("ps", "30")
	params.Set("tid", "0")

	crypedApi, err := SignAndGenerateURL(api+params.Encode(), cookie)
	if err != nil {
		return videoList, err
	}

	resp, err := client.
		Client.
		R().
		SetHeader("Origin", "https://space.bilibili.com").
		SetHeader("Referer", client.Origin+"/"+strconv.FormatInt(mid, 10)).
		SetHeader("Cookie", cookie).
		Get(crypedApi)
	if err != nil {
		slog.Error("get json error", "err", err.Error())
		return
	}
	if resp.IsErrorState() {
		slog.Error("get json state error", "body", resp.String())
		return
	}

	err = resp.Unmarshal(&videoList)
	if err != nil {
		slog.Error("parse json error", "err", err.Error())
	}

	slog.Info(fmt.Sprintf("爬取up主视频列表成功,mid:%d，第%d页", mid, page))
	return
}

func SearchVideoList4KeyWord(keyWord string, cookie string, opt model.Option) (vedioDataList []model.SearchVideoItem, err error) {
	api := "https://api.bilibili.com/x/web-interface/wbi/search/type?"
	re := regexp.MustCompile(`<em class="keyword">.*?</em>`)

	nextPage := 1 // 起始页

	for {
		params := url.Values{}
		params.Set("search_type", "video")
		params.Set("platform", "web")
		params.Set("keyword", keyWord)
		params.Set("from_source", "web_search")
		params.Set("page", fmt.Sprintf("%d", nextPage))
		// 上传时间
		var pubtimeBeginSec, pubtimeEndSec = 0, 0
		currentTime := time.Now()
		switch opt.UploadTime {
		case 0:
		case 1:
			pubtimeEndSec = int(time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location()).Unix())
			pubtimeBeginSec = pubtimeEndSec - 86399
		case 2:
			pubtimeEndSec = int(time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location()).Unix())
			pubtimeBeginSec = pubtimeEndSec - 604799
		case 3:
			pubtimeEndSec = int(time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location()).Unix())
			pubtimeBeginSec = pubtimeEndSec - 15551999
		default:
			return vedioDataList, fmt.Errorf("Erro:invaild uploadtime")

		}
		params.Set("pubtime_begin_s", strconv.Itoa(pubtimeBeginSec))
		params.Set("pubtime_end_s", strconv.Itoa(pubtimeEndSec))
		// 时长
		var duration int
		switch opt.Duration {
		case 0:
		case 1, 2, 3, 4:
			duration = opt.Duration
		default:
			return vedioDataList, fmt.Errorf("Erro:invaild duration")
		}
		if duration != 0 {
			params.Set("duration", strconv.Itoa(duration))
		}
		reqURL := api + params.Encode()
		slog.Info("请求页面", "page", nextPage, "url", reqURL)
		utils.RandomDelay(opt.MaxDelaySec)
		var respData model.SearchResp
		resp, err := client.
			Client.
			R().
			SetHeader("Cookie", cookie).
			SetHeader("origin", "https://search.bilibili.com").
			SetHeader("referer", "https://search.bilibili.com/video?"+params.Encode()).
			SetHeader("sec-fetch-dest", "empty").
			SetHeader("sec-fetch-mode", "cors").
			SetHeader("sec-fetch-site", "same-origin").
			SetHeader("accept", "application/json, text/plain, */*").
			Get(reqURL)
		if err != nil {
			slog.Error("search request failed", "err", err)
			return nil, err
		}
		if resp.IsErrorState() {
			slog.Error("search response error", "status", resp.StatusCode, "body", resp.String())
			return nil, fmt.Errorf("search request failed: status %d", resp.StatusCode)
		}
		err = resp.Unmarshal(&respData)
		if err != nil {
			return nil, fmt.Errorf("解析json失败: %w", err)
		}
		if respData.Code != 0 {
	return nil, fmt.Errorf("search api error: code %d, message %s", respData.Code, respData.Message)
}
		slog.Info(fmt.Sprintf("爬取视频列表成功,当前为%d页", respData.Data.Page))
		for _, v := range respData.Data.Result {
			v.Title = re.ReplaceAllString(v.Title, "")
			vedioDataList = append(vedioDataList, v)
		}

		if respData.Data.Next <= nextPage || (opt.Pages > 0 && nextPage >= opt.Pages) {
			slog.Info("已经是最后一页，爬取结束")
			break
		}
		nextPage++
	}
	return vedioDataList, nil
}
