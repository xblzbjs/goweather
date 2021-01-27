# goweather

简单的命令行天气客户端使用OpenWeather API和高德地图地理编码。

## 使用

### 第一步

在`constants.go`文件中输入你的OpenWeatherApiKey（需要申请）

### 第二步

在`constants.go`文件中输入你的GaoDeApiKey（需要申请）

## 命令行参数：

```shell
Usage: goweather [ -period=current|hourly|daily ] [ -units=C(摄氏度)|F(华氏度) ] <地点>...

Usage of goweather:
  -period string
        current | hourly | daily (default "current")
  -units string
        C(摄氏度) | F(华氏度) (default "C")

```



## 例子

```shell
$ goweather -period=daily -units=C 深圳市南山区
深圳市南山区的天气:
Wednesday  1/27   High: 21.49C Low: 17.85C | Humidity: 57% | 多云
Thursday   1/28   High: 22.93C Low: 17.18C | Humidity: 45% | 晴
Friday     1/29   High: 18.79C Low: 14.46C | Humidity: 44% | 晴
Saturday   1/30   High: 18.93C Low: 13.86C | Humidity: 43% | 晴
Sunday     1/31   High: 20.98C Low: 15.37C | Humidity: 43% | 晴
Monday     2/ 1   High: 22.32C Low: 16.82C | Humidity: 48% | 多云
Tuesday    2/ 2   High: 21.63C Low: 18.15C | Humidity: 54% | 晴
Wednesday  2/ 3   High: 20.85C Low: 16.73C | Humidity: 59% | 多云
```



## TODO:

- 完善高德地图地理编码
- 高德天气查询G
- Google Map PlatForm  API整合