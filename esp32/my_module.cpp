#include "my_module.h"

my_module::my_module() {
}

my_module::~my_module() {
  delete _dht11;
}

void my_module::init() {
  _dht11 = new SimpleDHT11(PIN_TH);
  pinMode(PIN_LED, OUTPUT);
  pinMode(PIN_SRTRIG, OUTPUT);
  pinMode(PIN_BUZZER, OUTPUT);
  pinMode(PIN_SERVO, OUTPUT);
}

unsigned long my_module::sr_ping() {
  digitalWrite(PIN_SRTRIG, HIGH);
  vTaskDelay(pdMS_TO_TICKS(10));
  digitalWrite(PIN_SRTRIG, LOW);
  return pulseIn(PIN_SRECHO, HIGH);
}

void my_module::servo_lock() {
  for (int i = 0; i < 100; i++) {
    digitalWrite(PIN_SERVO, HIGH);
    vTaskDelay(pdMS_TO_TICKS(500));  //0.5ms
    digitalWrite(PIN_SERVO, LOW);
    vTaskDelay(pdMS_TO_TICKS(19500));  //19.5ms
  }
}

void my_module::servo_lock_down() {
  for (int i = 0; i < 100; i++) {
    digitalWrite(PIN_SERVO, HIGH);
    vTaskDelay(pdMS_TO_TICKS(1500));  //1.5ms
    digitalWrite(PIN_SERVO, LOW);
    vTaskDelay(pdMS_TO_TICKS(18500));  //18.5ms
  }
}

void my_module::module_run_once() {
  //光敏数据
  _IsDark = digitalRead(PIN_LDR);
  if (_IsDark == 0) {
    digitalWrite(PIN_LED, LOW);
  } else {
    digitalWrite(PIN_LED, HIGH);
  }
  //温湿度数据
  byte temperature = 0;
  byte humidity = 0;
  int err = SimpleDHTErrSuccess;
  if ((err = _dht11->read(&temperature, &humidity, NULL)) == SimpleDHTErrSuccess) {
    _Temperature = (int)temperature - 1;
    _Humidity = (int)humidity + 8;
  }
  //火焰数据
  _IsFlame = digitalRead(PIN_FLAME);
  //可燃气体数据
  _IsFlammable = digitalRead(PIN_GAS);
  //报警
  if (_IsFlame == 0 || _IsFlammable == 0) {
    if (_IsWarn == 0) {
      tone(PIN_BUZZER, 330);
      _IsWarn = 1;
    }
  } else if (_IsWarn == 1) {
    noTone(PIN_BUZZER);
    _IsWarn = 0;
  }
  //超声波数据
  _Distance = sr_ping() / 58;
}

string my_module::generate_data_msg() {
  std::ostringstream buffer;
  buffer << ESP32_ID << ":"
         << _Temperature << ":"
         << _Humidity << ":"
         << _IsFlame << ":"
         << _IsFlammable;
  return buffer.str();
}
