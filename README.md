# go weather

## Overview

A simple command line weather client uses Gaode Weather Api

## Installting

```bash

```

## Getting Start

1. Apply for [Gaode Api](https://lbs.amap.com/dev/key/app) and [OpenweatherMap Api](https://openweathermap.org/api)

2. Create a `config.yml` file in the project path

   ```bash
   goweather
   ├── README.md
   ...
   ├── config.yml		# Create Here
   ...
   ├── sample_config.yml # Module
   ```

   

3. 

## 

## Example

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

- OpenWeather Api with Google Map PlatForm
- Translate