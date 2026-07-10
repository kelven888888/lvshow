package utils

import (
	"encoding/json"
	"errors"
	"ginshop.com/global"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gofrs/uuid"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"
)

type Parms struct {
	Keys  string
	Value string
}

func Response(ctx *gin.Context, httpStatus int, code int, msg string, data any, args ...Parms) {

	language, _ := ctx.Get("Language")
	if language == nil {
		language = global.SHOP_CONFIG.System.Language
	}
	msg = Languageresponse(msg, language.(string))
	for _, v := range args {
		msg = strings.Replace(msg, v.Keys, v.Value, 1)
	}

	//arg := &args[0]
	//fmt.Println(fmt.Sprintf("%+v", arg))
	//
	//val := reflect.ValueOf(arg)
	//
	//// 检查反射值的类型是否为map
	//fmt.Println(val.Kind(), val, arg)
	//if val.Kind() != reflect.Map {
	//
	//	fmt.Println("Error: Not a map type.")
	//
	//}
	//
	//fmt.Println("--- Iterating Map ---")
	//
	//// 获取map的所有键
	//
	//keys := val.MapKeys()
	//
	//// 遍历键
	//
	//for _, key := range keys {
	//
	//	// 通过键获取对应的值
	//
	//	value := val.MapIndex(key)
	//
	//	// 将reflect.Value转换回原始接口类型，以便打印或进一步处理
	//
	//	fmt.Printf("Key: %v, Value: %v\n", key.Interface(), value.Interface())
	//
	//}
	ctx.JSON(httpStatus, gin.H{
		"code":     code,
		"msg":      msg,
		"data":     data,
		"language": language,
		"website":  global.SHOP_CONFIG.System.WebApiURL,
	})
}

func Success(ctx *gin.Context, msg string, data any, args ...Parms) {

	Response(ctx, http.StatusOK, 200, msg, data, args...)
}

func Fail(ctx *gin.Context, msg string, data any, args ...Parms) {

	Response(ctx, http.StatusOK, 400, msg, data, args...)
}

// 获取全部请求参数
func DataMapByRequest(c *gin.Context) gin.H {
	// 1. 获取路径参数 (需要知道路由定义了哪些参数，这里假设已知或手动构建)
	// 注意：Gin 没有直接导出所有 Path Params 的 Map，通常需手动指定
	pathParams := gin.H{}
	// 示例：如果路由是 /api/:version/:resource
	// pathParams["version"] = c.Param("version")
	// pathParams["resource"] = c.Param("resource")

	// 2. 获取所有查询参数
	queryParams := c.Request.URL.Query()

	// 3. 获取所有请求头
	headers := c.Request.Header

	// 4. 获取请求体参数 (根据 Content-Type 判断)
	var bodyParams interface{}
	contentType := c.GetHeader("Content-Type")

	if strings.Contains(contentType, "application/json") {
		var jsonBody map[string]interface{}
		// 使用 ShouldBindBodyWith 允许后续再次读取 Body（如果需要）
		if err := c.ShouldBindBodyWith(&jsonBody, binding.JSON); err == nil {
			bodyParams = jsonBody
		} else {
			bodyParams = "Error parsing JSON: " + err.Error()
		}
	} else if strings.Contains(contentType, "application/x-www-form-urlencoded") ||
		strings.Contains(contentType, "multipart/form-data") {
		var formBody map[string][]string
		if err := c.ShouldBindWith(&formBody, binding.Form); err == nil {
			bodyParams = formBody
		} else {
			bodyParams = "Error parsing Form: " + err.Error()
		}
	} else {
		bodyParams = "Unsupported Content-Type or empty body"
	}

	// 返回汇总结果
	return gin.H{
		"path_params":  pathParams, // 需根据实际情况填充
		"query_params": queryParams,
		"headers":      headers,
		"body_params":  bodyParams,
	}
}

// 生成指定长度的随机字符
func RandomString(n int) string {
	var letters = []byte("qwertyuioplkjhgfdsazxcvbnmMNBVCXZASDFGHJKLPOIUYTREWQ")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// 获取外网IP
func ExternalIp() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network?")
}

// 格式化 IP
func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil
	}
	return ip
}

// 过滤指定数组中的key
func ParamsFilter(isFilterStr string, params map[string]any) map[string]any {
	var data = make(map[string]any)
	for key, value := range params {
		if value != "" {
			find := strings.Contains(isFilterStr, key)
			if !find {
				data[key] = value
			}
		}
	}
	return data
}

func UUID() string {
	return uuid.Must(uuid.NewV4()).String()
}

// AnyToMap interface 转 map
func AnyToMap(v any) (map[string]any, error) {
	dataJson, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	var MapData map[string]any
	err = json.Unmarshal(dataJson, &MapData)
	if err != nil {
		return nil, err
	}
	return MapData, nil
}
