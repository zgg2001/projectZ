#include <SimpleDHT.h>

#define LED 4 //灯
#define LDR 13 //光敏
#define TH 12 //温湿度
#define FLAME 14 //火焰
#define GAS 27 //可燃气体
#define SRTRIG 18 //超声波发出
#define SRECHO 5 //超声波接收

SemaphoreHandle_t mutexHandle;
SimpleDHT11 dht11(TH);

int Intensity = 0; //光照度数值
int Temperature = 0; //温度
int Humidity = 0; //湿度
int IsFlame = 0; //火焰
int IsFlammable = 0; //可燃气体
unsigned long Distance = 0; //超声波距离

//超声波探距
unsigned long sr_ping();

void setup() {  
  pinMode(LED, OUTPUT);
  pinMode(SRTRIG, OUTPUT); 
  mutexHandle = xSemaphoreCreateMutex();
  Serial.begin(115200);
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
    Intensity = analogRead(LDR);  //读取GPIO的值，存入Intensity变量
    if(Intensity < 3000) {
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
    IsFlame = analogRead(FLAME);
    //可燃气体数据
    IsFlammable = analogRead(GAS);
    //超声波数据
    Distance = sr_ping()/58;
    xSemaphoreGive(mutexHandle);
    vTaskDelay(pdMS_TO_TICKS(1000));
  } else {
    vTaskDelay(pdMS_TO_TICKS(1000));
  }
}

void task1(void *parameter) {
  BaseType_t ret;  
  vTaskDelay(pdMS_TO_TICKS(1000));
  while (1) {
    ret = xSemaphoreTake(mutexHandle, 1000);
    if (ret == pdPASS) {
      Serial.println();      
      //光照
      Serial.print("Intensity = ");  
      Serial.println(Intensity);
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

unsigned long sr_ping() { 
  digitalWrite(SRTRIG, HIGH);
  vTaskDelay(pdMS_TO_TICKS(10));
  digitalWrite(SRTRIG, LOW);
  return pulseIn(SRECHO, HIGH);
}
