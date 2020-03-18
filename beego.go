package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//获取指定目录下的所有文件,包含子目录下的文件
func getAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			getAllFiles(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".go")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := getAllFiles(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}

func fixControllers(path string) {
	controllersPath := filepath.Join(path, "controllers")
	if fileList, err := getAllFiles(controllersPath); err == nil {
		for _, fileName := range fileList {
			if data, err := ioutil.ReadFile(fileName); err == nil {
				newData := strings.ReplaceAll(string(data), "c.Data[\"json\"] = l", "c.Data[\"json\"] = utils.PageUtil(count, offset, limit, l)")
				newData = strings.ReplaceAll(newData, "l, err := models.", "l, count, err := models.")
				ioutil.WriteFile(fileName, []byte(newData), 0644)
			}
		}
	}
}

func fixModels(path string) {
	modelsPath := filepath.Join(path, "models")
	if fileList, err := getAllFiles(modelsPath); err == nil {
		for _, fileName := range fileList {
			if data, err := ioutil.ReadFile(fileName); err == nil {
				newData := strings.ReplaceAll(string(data), "(ml []interface{}, err error)", "(ml []interface{}, count int64, err error)")
				newData = strings.ReplaceAll(newData, "return nil, errors.New", "return nil, 0, errors.New")
				oldStr :=
					`return ml, nil
	}
	return nil, err`
				newStr :=
					`return ml, count, nil
	}
	return nil, 0, err`
				newData = strings.ReplaceAll(newData, oldStr, newStr)
				if !strings.Contains(newData, "if offset == 0 {") {
					newStr :=
						`	if offset == 0 {
		count, err = qs.Count()
	}

	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {`
					newData = strings.ReplaceAll(newData, "if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {", newStr)
				}
				ioutil.WriteFile(fileName, []byte(newData), 0644)
			}
		}
	}
}

func RunBeego(path string) {
	fixControllers(path)
	fixModels(path)
}
