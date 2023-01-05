#include <string.h>
#include <SimpleDHT.h>
#include <WiFi.h>
#include "WiFiClient.h"
#include "Adafruit_MQTT.h"
#include "Adafruit_MQTT_Client.h"

#define WLAN_SSID "123"
#define WLAN_PASS "12345678"

#define MQTT_SERVER "192.168.53.133"
#define MQTT_SERVERPORT 1883

#define MQTT_USERNAME "test1"
#define MQTT_PASSWORD "a123456"

#define LED 4      //灯
#define LDR 27     //光敏
#define TH 25      //温湿度
#define FLAME 14   //火焰
#define GAS 13     //可燃气体
#define SRTRIG 18  //超声波发出
#define SRECHO 5   //超声波接收
#define BUZZER 19  //蜂鸣器
#define SERVO 32   //舵机A

SemaphoreHandle_t mutexHandle;
WiFiClient client;
Adafruit_MQTT_Client mqtt(&client, MQTT_SERVER, MQTT_SERVERPORT, MQTT_USERNAME, MQTT_PASSWORD);
Adafruit_MQTT_Publish pub = Adafruit_MQTT_Publish(&mqtt, "my/mqtt/topic");
SimpleDHT11 dht11(TH);

int IsDark = 0;              //光照 - LED 0灭1亮
int Temperature = 0;         //温度
int Humidity = 0;            //湿度
int IsFlame = 1;             //火焰 - 0报警
int IsFlammable = 1;         //可燃气体 - 0报警
unsigned long Distance = 0;  //超声波距离
int IsWarn = 0;              //是否蜂鸣警告中

//MQTT connect
void MQTT_connect();
//超声波探距
unsigned long sr_ping();
//舵机上锁模式
void servo_lock();
//舵机下锁模式
void servo_lock_down();

void setup() {
  pinMode(LED, OUTPUT);
  pinMode(SRTRIG, OUTPUT);
  pinMode(BUZZER, OUTPUT);
  pinMode(SERVO, OUTPUT);
  mutexHandle = xSemaphoreCreateMutex();
  Serial.begin(115200);

  WiFi.begin(WLAN_SSID, WLAN_PASS);
  vTaskDelay(pdMS_TO_TICKS(1000));
  Serial.print("Connecting to ");
  Serial.println(WLAN_SSID);
  while (WiFi.status() != WL_CONNECTED) {
    delay(1000);
    Serial.print(".");
  }
  Serial.println("WiFi connected");
  Serial.println("IP address: ");
  Serial.println(WiFi.localIP());

  vTaskDelay(pdMS_TO_TICKS(1000));
  xTaskCreate(
    task1,   /* Task function. */
    "Task1", /* String with name of task. */
    10000,   /* Stack size in bytes. */
    NULL,    /* Parameter passed as input of the task */
    1,       /* Priority of the task. */
    NULL);   /* Task handle. */

  xTaskCreate(
    task2,   /* Task function. */
    "Task2", /* String with name of task. */
    10000,   /* Stack size in bytes. */
    NULL,    /* Parameter passed as input of the task */
    1,       /* Priority of the task. */
    NULL);   /* Task handle. */
}

void loop() {
  BaseType_t ret;
  ret = xSemaphoreTake(mutexHandle, 1000);
  if (ret == pdPASS) {
    //光敏数据
    IsDark = digitalRead(LDR);  //读取GPIO的值，存入Intensity变量
    if (IsDark == 0) {
      digitalWrite(LED, LOW);
    } else {
      digitalWrite(LED, HIGH);
    }
    //温湿度数据
    byte temperature = 0;
    byte humidity = 0;
    int err = SimpleDHTErrSuccess;
    if ((err = dht11.read(&temperature, &humidity, NULL)) == SimpleDHTErrSuccess) {
      Temperature = (int)temperature - 1;
      Humidity = (int)humidity + 8;
    }
    //火焰数据
    IsFlame = digitalRead(FLAME);
    //可燃气体数据
    IsFlammable = digitalRead(GAS);
    //报警
    if (IsFlame == 0 || IsFlammable == 0) {
      if (IsWarn == 0) {
        tone(BUZZER, 330);
        IsWarn = 1;
      }
    } else if (IsWarn == 1) {
      noTone(BUZZER);
      IsWarn = 0;
    }
    //超声波数据
    Distance = sr_ping() / 58;
    xSemaphoreGive(mutexHandle);
    vTaskDelay(pdMS_TO_TICKS(1000));
  } else {
    vTaskDelay(pdMS_TO_TICKS(1000));
  }
}

void task1(void *parameter) {
  BaseType_t ret;
  char buf[15];  //xx-xxx-xxx-x-x\0
  vTaskDelay(pdMS_TO_TICKS(5000));
  while (1) {
    ret = xSemaphoreTake(mutexHandle, 1000);
    if (ret == pdPASS) {
      Serial.println();
      //光照
      Serial.print("IsDark = ");
      Serial.println(IsDark);
      //温湿度
      Serial.print("Temperature(°C) = ");
      Serial.println(Temperature);
      Serial.print("Humidity(H) = ");
      Serial.println(Humidity);
      //火焰
      Serial.print("IsFlame = ");
      Serial.println(IsFlame);
      //可燃气体
      Serial.print("IsFlammable = ");
      Serial.println(IsFlammable);
      //超声波
      Serial.print("Distance(cm) = ");
      Serial.println(Distance);
      snprintf(buf, 15, "%02d:%03d:%03d:%01d:%01d", 1, Temperature, Humidity, IsFlame, IsFlammable);
      //MQTT connect and send
      MQTT_connect();
      pub.publish(buf);
      xSemaphoreGive(mutexHandle);
      vTaskDelay(pdMS_TO_TICKS(5000));
    } else {
      vTaskDelay(pdMS_TO_TICKS(1000));
    }
  }
  Serial.println("Ending task 1");
  vTaskDelete(NULL);
}

void task2(void *parameter) {
  //Serial.println("Ending task 2");
  vTaskDelete(NULL);
}

void MQTT_connect() {
  int8_t ret;
  if (mqtt.connected()) {
    return;
  }
  while ((ret = mqtt.connect()) != 0) {
    mqtt.disconnect();
    delay(5000);
  }
}

unsigned long sr_ping() {
  digitalWrite(SRTRIG, HIGH);
  vTaskDelay(pdMS_TO_TICKS(10));
  digitalWrite(SRTRIG, LOW);
  return pulseIn(SRECHO, HIGH);
}

void servo_lock() {
  for(int i = 0; i<100; i++)
  {
    digitalWrite(SERVO,HIGH);
    delayMicroseconds(500);//1.5ms
    digitalWrite(SERVO,LOW);
    delayMicroseconds(19500);//18.5ms
  }
}

void servo_lock_down() {
  for(int i = 0; i<100; i++)
  {
    digitalWrite(SERVO,HIGH);
    delayMicroseconds(1500);//1.5ms
    digitalWrite(SERVO,LOW);
    delayMicroseconds(18500);//18.5ms
  }
}
