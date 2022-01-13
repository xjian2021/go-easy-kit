package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

type Watch struct {
	watch *fsnotify.Watcher
}

//AddDir 实现递归监控
func (w *Watch) AddDir(dir string) {
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			if err = w.watch.Add(path); err != nil {
				log.Fatalf("w.watch.Add fail err:%s", err.Error())
				return err
			}
			log.Printf("添加监控：%s\n", info.Name())
		}
		return nil
	})
	if err != nil {
		log.Fatalf("filepath.Walk fail err:%s", err.Error())
	}
}

func (w *Watch) StartWatch() {
	fmt.Println("开始监视文件夹...")
	for {
		select {
		case event := <-w.watch.Events:
			switch {
			case event.Op == fsnotify.Create:
				//log.Println(event.String())
				log.Println("创建文件:", event.Name)
				stat, err := os.Stat(event.Name)
				if err == nil && stat.IsDir() {
					if err = w.watch.Add(event.Name); err != nil {
						log.Printf("添加监控:%s 失败\n", event.Name)
					}
					log.Println("添加监控:", event.Name)
				}
			case event.Op == fsnotify.Write:
				//log.Println(event.String())
				log.Println("写入文件:", event.Name)
			case event.Op == fsnotify.Rename:
				//log.Println(event.String())
				log.Println("重命名文件:", event.Name)
			case event.Op == fsnotify.Chmod:
				//log.Println(event.String())
				log.Println("修改权限:", event.Name)
			case event.Op == fsnotify.Remove:
				//log.Println(event.String())
				log.Println("删除文件:", event.Name)
			}
		case err := <-w.watch.Errors:
			log.Fatalf("watcher.Errors:%s", err.Error())
		}
	}
}

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	dirName := "./tmp"
	w := &Watch{watch: watcher}
	w.AddDir(dirName)
	w.StartWatch()
	//if err = watcher.Add("./tmp"); err != nil {
	//	log.Fatal(err)
	//}
	//go func() {
	//	fmt.Println("开始监视文件夹...")
	//	for {
	//		select {
	//		case event := <-watcher.Events:
	//			switch {
	//			case event.Op == fsnotify.Create:
	//				//log.Println(event.String())
	//				log.Println("创建文件:", event.Name)
	//				stat, err := os.Stat(event.Name)
	//				if err == nil && stat.IsDir() {
	//					if err = watcher.Add(event.Name); err != nil {
	//						log.Printf("添加监控:%s 失败\n", event.Name)
	//					}
	//					log.Println("添加监控:", event.Name)
	//				}
	//			case event.Op == fsnotify.Write:
	//				//log.Println(event.String())
	//				log.Println("写入文件:", event.Name)
	//			case event.Op == fsnotify.Rename:
	//				//log.Println(event.String())
	//				log.Println("重命名文件:", event.Name)
	//			case event.Op == fsnotify.Chmod:
	//				//log.Println(event.String())
	//				log.Println("修改权限:", event.Name)
	//			case event.Op == fsnotify.Remove:
	//				//log.Println(event.String())
	//				log.Println("删除文件:", event.Name)
	//			}
	//		case err = <-watcher.Errors:
	//			log.Fatalf("watcher.Errors:%s", err.Error())
	//		}
	//	}
	//}()
	//
	//select {}

}
