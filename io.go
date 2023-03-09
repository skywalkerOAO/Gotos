package Gotos

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
)

/*CopyFile
 * @Author Hex575A
 * @Description 将文件复制到另一个目录下
 * @Date 8:53 2022/8/30
 * @Param dstName 源文件绝对路径
 * @Param srcName 目标文件绝对路径
 * @return err 失败返回错误，成功无返回值。
 */
func CopyFile(dstName, srcName string) error {
	file, err := os.Open(dstName)
	filepath, err := os.Create(srcName)
	if err != nil {
		return err
	}
	//写入二进制数据
	_, err = io.Copy(filepath, file)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	err = filepath.Close()
	if err != nil {
		return err
	}
	return nil
}

/*PathExists
 * @Author Hex575A
 * @Description 判断目录或文件是否存在
 * @Date 8:51 2022/8/30
 * @Param path string 目录路径
 * @return bool 存在返回true 不存在返回false
 */
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

/*CreatePath
 * @Author Hex575A
 * @Description 自动创建目录，创建成功返回true 否返回false
 * @Date 16:11 2022/8/29
 * @Param path 目录路径
 * @Param FileMode 创建文件的权限
 * @return bool 创建成功返回true 失败返回false
 */
func CreatePath(path string, FileMode fs.FileMode) bool {
	if !PathExists(path) {
		err := os.Mkdir(path, FileMode)
		if err != nil {
			fmt.Println(fmt.Sprintf("创建目录'%s'失败", path))
			return false
		}
	}
	return true
}

/*ReceiveFile
 * @Author Hex575A
 * @Description 接收前端的文件并保存至TargetPath
 * @Date 15:58 2022/8/29
 * @Param r http请求。
 * @Param TargetPath 接收文件后储存的绝对目录。
 * @Param FileGroup 文件组名，须有前后端协调一致，若组名不对则无法获取相应文件。
 * @Param ValueTag 是否包含其他参数，若包含须在MultipartForm中的Value中，若不包含请留空或传入""。
 * @return map[int]map[string]interface{} 返回值为一个文件的详细信息，格式为
 * filelist[i]["temp_target_path"] = targetPath 文件i的路径
 * filelist[i]["file_Args"] = values            文件i的额外参数(若无则无此item)
 * filelist[i]["file_name"] = files[i].Filename 文件i的名称
 */
func ReceiveFile(r *http.Request, TargetPath string, FileGroup string, ValueTag string) map[int]map[string]interface{} {
	//设置内存缓冲区大小
	err := r.ParseMultipartForm(256 << 20)
	if err != nil {
		return nil
	}
	if len(r.MultipartForm.Value) != 0 {
		//获取上传的文件组，值为前端formdata.append的组名
		files := r.MultipartForm.File[FileGroup]
		values := r.MultipartForm.Value[ValueTag][0]
		lens := len(files)
		filelist := make(map[int]map[string]interface{})
		for i := 0; i < lens; i++ {
			//创建文件
			file, err := files[i].Open()
			targetPath := TargetPath + files[i].Filename
			filelist[i] = make(map[string]interface{})
			filelist[i]["temp_target_path"] = targetPath
			filelist[i]["file_Args"] = values
			filelist[i]["file_name"] = files[i].Filename
			filepath, err := os.Create(targetPath)
			if err != nil {
				panic(err)
			}
			//写入二进制数据
			_, err = io.Copy(filepath, file)
			if err != nil {
				panic(err)
			}
			// 结束打开的文件释放占用状态（内存）
			file.Close()
			// 结束创建的新文件的占用状态）（硬盘）
			filepath.Close()
		}

		return filelist
	}
	files := r.MultipartForm.File[FileGroup]
	lens := len(files)
	filelist := make(map[int]map[string]interface{})
	for i := 0; i < lens; i++ {
		//创建文件
		file, err := files[i].Open()
		targetPath := TargetPath + files[i].Filename
		filelist[i] = make(map[string]interface{})
		filelist[i]["temp_target_path"] = targetPath
		filelist[i]["file_name"] = files[i].Filename
		filepath, err := os.Create(targetPath)
		if err != nil {
			panic(err)
		}
		//写入二进制数据
		_, err = io.Copy(filepath, file)
		if err != nil {
			panic(err)
		}
		// 结束打开的文件释放占用状态（内存）
		file.Close()
		// 结束创建的新文件的占用状态）（硬盘）
		filepath.Close()
	}
	return filelist
}
