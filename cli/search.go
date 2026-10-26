package cli

import (
	"blblcd/core"
	"blblcd/utils"
	"blblcd/model"
	"fmt"
	"sync"
	"github.com/spf13/cobra"
)


func init() {
	rootCmd.AddCommand(searchCmd)
}


var searchCmd = &cobra.Command{
	Use: "search [key-word]",
	Example: "blblcd search  宝可梦",
	Short:   "根据关键词搜索站内视频",
	Run: func(cmd *cobra.Command, args []string) {
		if keyWord == "" {
			fmt.Println("please provide key word")
			return
		}
		fmt.Printf("爬取关于%s的视频\n",keyWord)
		cookie, err := utils.ReadTextFile(cookieFile)
		if err != nil {
			fmt.Println(err)
			return
		}

		utils.PrintLogo()
		opt := model.Option{
			Corder:      corder,
			Mapping:     mapping,
			Cookie:      cookie,
			Output:      output,
			ImgDownload: imgDownload,
			MaxTryCount: maxTryCount,
			MaxDelaySec: maxDelaySec,
			Duration: duration,
			UploadTime: uploadTime,
		}
		videoList,err :=core.SearchVideoList4KeyWord(keyWord,cookie,opt)
		if err != nil {
			fmt.Println(fmt.Errorf("Erro:%s",err))
			return 
		}
		fmt.Println(videoList)
		// 初始化struct的chan
		sem := make(chan struct{}, workers)
		wg := sync.WaitGroup{}
		for _,v := range videoList {
			bvid := v.Bvid
			opt.Bvid=bvid
			sem <- struct{}{}
			wg.Add(1)
			go func() {
				defer wg.Done()
				defer func() {
					<-sem
				}()
				core.FindComment(core.Bvid2Avid(bvid), &opt)
			}()
		}
		wg.Wait()
	},

}

