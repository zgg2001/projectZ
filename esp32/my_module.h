#ifndef _MY_MODULE_H_
#define _MY_MODULE_H_

#include <sstream>
#include <SimpleDHT.h>

using std::string;

class my_module {
public:
  static constexpr int ESP32_ID = 1;
  static constexpr int PIN_LED = 4;      //灯
  static constexpr int PIN_LDR = 27;     //光敏
  static constexpr int PIN_TH = 25;      //温湿度
  static constexpr int PIN_FLAME = 14;   //火焰
  static constexpr int PIN_GAS = 13;     //可燃气体
  static constexpr int PIN_SRTRIG = 5;   //超声波发出
  static constexpr int PIN_SRECHO = 18;  //超声波接收
  static constexpr int PIN_BUZZER = 26;  //蜂鸣器
  static constexpr int PIN_SERVO = 32;   //舵机A

public:
  static my_module& get_instance() {
    static my_module ins;
    return ins;
  }

  void init();
  unsigned long sr_ping();  //超声波探距
  void servo_lock();        //舵机上锁模式
  void servo_lock_down();   //舵机下锁模式
  void module_run_once();
  string generate_data_msg();

private:
  my_module();
  virtual ~my_module();
  my_module(const my_module&);
  my_module& operator=(const my_module&);

private:
  int _IsDark = 0;              //光照 - LED 0灭1亮
  int _Temperature = 0;         //温度
  int _Humidity = 0;            //湿度
  int _IsFlame = 1;             //火焰 - 0报警
  int _IsFlammable = 1;         //可燃气体 - 0报警
  unsigned long _Distance = 0;  //超声波距离
  int _IsWarn = 0;              //是否蜂鸣警告中

  SimpleDHT11* _dht11;
};

#endif
