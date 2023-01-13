#include <string.h>
#include <WiFi.h>
#include "my_mqtt.h"
#include "my_module.h"

#define WLAN_SSID "123"
#define WLAN_PASS "12345678"
#define INS my_module::get_instance()

constexpr char MQTT_SERVER[] = "192.168.6.133";
constexpr uint16_t MQTT_SERVERPORT = 1883;
constexpr char MQTT_USERNAME[] = "test1";
constexpr char MQTT_PASSWORD[] = "a123456";
constexpr char MQTT_DATA_TOPIC[] = "pi/esp32/data";
constexpr char MQTT_CMD_TOPIC[] = "pi/esp32/cmd";

SemaphoreHandle_t mutexHandle;
mqtt_client *mymqtt;

void setup() {
  Serial.begin(115200);
  mutexHandle = xSemaphoreCreateMutex();
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

  //mqtt
  mymqtt = new mqtt_client(MQTT_SERVER, MQTT_SERVERPORT, MQTT_USERNAME, MQTT_PASSWORD, MQTT_DATA_TOPIC, MQTT_CMD_TOPIC);
  mymqtt->init();

  //mod
  INS.init();

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
    INS.module_run_once();
    xSemaphoreGive(mutexHandle);
    vTaskDelay(pdMS_TO_TICKS(1000));
  } else {
    vTaskDelay(pdMS_TO_TICKS(1000));
  }
}

void task1(void *parameter) {
  BaseType_t ret;
  string msg;
  vTaskDelay(pdMS_TO_TICKS(5000));
  while (1) {
    ret = xSemaphoreTake(mutexHandle, 1000);
    if (ret == pdPASS) {
      msg = INS.generate_data_msg();
      Serial.println(msg.c_str());
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
  while (1) {
    mymqtt->mqtt_sub(1000);
  }
  Serial.println("Ending task 2");
  vTaskDelete(NULL);
}
